package httpstats

import (
	"fmt"
	"io"
	"math"
	"net/url"
	"regexp"
	"sync"

	"github.com/tkuchiki/gohttpstats/options"
	"github.com/tkuchiki/gohttpstats/parsers"
)

type hints struct {
	values map[string]int
	len    int
	mu     sync.RWMutex
}

func newHints() *hints {
	return &hints{
		values: make(map[string]int),
	}
}

func (h *hints) loadOrStore(key string) int {
	h.mu.Lock()
	defer h.mu.Unlock()
	_, ok := h.values[key]
	if !ok {
		h.values[key] = h.len
		h.len++
	}

	return h.values[key]
}

type HTTPStats struct {
	hints                         *hints
	stats                         httpStats
	useResponseTimePercentile     bool
	useRequestBodySizePercentile  bool
	useResponseBodySizePercentile bool
	printOptions                   *PrintOptions
	filter                        *Filter
	parser                        parsers.Parser
	options                       *stats_options.Options
	uriCapturingGroups            []*regexp.Regexp
}

func NewHTTPStats(useResTimePercentile, useRequestBodySizePercentile, useResponseBodySizePercentile bool, po *PrintOptions) *HTTPStats {
	return &HTTPStats{
		hints: newHints(),
		stats: make([]*httpStat, 0),
		useResponseTimePercentile:     useResTimePercentile,
		useResponseBodySizePercentile: useResponseBodySizePercentile,
		printOptions:                   po,
	}
}

func (hs *HTTPStats) Set(uri, method string, status int, restime, resBodySize, reqBodySize float64) {
	if len(hs.uriCapturingGroups) > 0 {
		for _, re := range hs.uriCapturingGroups {
			if ok := re.Match([]byte(uri)); ok {
				pattern := re.String()
				uri = pattern
			}
		}
	}

	key := fmt.Sprintf("%s_%s", method, uri)

	idx := hs.hints.loadOrStore(key)

	if idx >= len(hs.stats) {
		hs.stats = append(hs.stats, newHTTPStat(uri, method, hs.useResponseTimePercentile, hs.useRequestBodySizePercentile, hs.useResponseBodySizePercentile))
	}

	hs.stats[idx].Set(status, restime, resBodySize, reqBodySize)
}

func (hs *HTTPStats) Stats() []*httpStat {
	return hs.stats
}

func (hs *HTTPStats) CountUris() int {
	return hs.hints.len
}

func (hs *HTTPStats) SetOptions(options *stats_options.Options) {
	hs.options = options
}

func (hs *HTTPStats) SetURICapturingGroups(groups []string) error {
	uriGroups, err := CompileUriGroups(groups)
	if err != nil {
		return err
	}

	hs.uriCapturingGroups = uriGroups

	return nil
}

func (hs *HTTPStats) InitFilter(options *stats_options.Options) error {
	hs.filter = NewFilter(options)
	return hs.filter.Init()
}

func (hs *HTTPStats) DoFilter(uri, status, timestr string) bool {
	err := hs.filter.Do(uri, status, timestr)
	if err == SkipReadLineErr || err != nil {
		return false
	}

	return true
}

func (hs *HTTPStats) InitParser(parserType string, r io.Reader) error {
	switch parserType {
	case "ltsv":
		hs.parser = parsers.NewLTSVParser(r)
	default:
		return fmt.Errorf("Parser Not Supproted: %s", parserType)
	}

	return nil
}

func (hs *HTTPStats) Parse() (string, string, string, float64, float64, int, error) {
	line, err := hs.parser.Read()
	if err != nil {
		return "", "", "", 0, 0, 0, err
	}

	u, err := url.Parse(line[hs.options.UriLabel])
	if err != nil {
		return "", "", "", 0, 0, 0, err
	}
	var uri string
	if hs.options.QueryString {
		v := url.Values{}
		values := u.Query()
		for q := range values {
			v.Set(q, "xxx")
		}
		uri = fmt.Sprintf("%s?%s", u.Path, v.Encode())
	} else {
		uri = u.Path
	}

	resTime, err := StringToFloat64(line[hs.options.ApptimeLabel])
	if err != nil {
		var reqTime float64
		reqTime, err = StringToFloat64(line[hs.options.ReqtimeLabel])
		if err != nil {
			return "", "", "", 0, 0, 0, SkipReadLineErr
		}

		resTime = reqTime
	}

	bodySize, err := StringToFloat64(line[hs.options.SizeLabel])
	if err != nil {
		return "", "", "", 0, 0, 0, SkipReadLineErr
	}

	status, err := StringToInt(line[hs.options.StatusLabel])
	if err != nil {
		return "", "", "", 0, 0, 0, SkipReadLineErr
	}

	method := line[hs.options.MethodLabel]
	timestr := line[hs.options.TimeLabel]

	return uri, method, timestr, resTime, bodySize, status, nil
}

func (hs *HTTPStats) SortWithOptions() {
	hs.Sort(hs.options.Sort, hs.options.Reverse)
}

type httpStat struct {
	Uri              string        `yaml:uri`
	Cnt              int           `yaml:cnt`
	Status1xx        int           `yaml:status1xx`
	Status2xx        int           `yaml:status2xx`
	Status3xx        int           `yaml:status3xx`
	Status4xx        int           `yaml:status4xx`
	Status5xx        int           `yaml:status5xx`
	Method           string        `yaml:method`
	ResponseTime     *responseTime `yaml:response_time`
	RequestBodySize  *bodySize     `yaml:request_body_size`
	ResponseBodySize *bodySize     `yaml:response_body_size`
}

type httpStats []*httpStat

func newHTTPStat(uri, method string, useResTimePercentile, useRequestBodySizePercentile, useResponseBodySizePercentile bool) *httpStat {
	return &httpStat{
		Uri:              uri,
		Method:           method,
		ResponseTime:     newResponseTime(useResTimePercentile),
		RequestBodySize:  newBodySize(useRequestBodySizePercentile),
		ResponseBodySize: newBodySize(useResponseBodySizePercentile),
	}
}

func (hs *httpStat) Set(status int, restime, reqBodySize, resBodySize float64) {
	hs.Cnt++
	hs.setStatus(status)
	hs.ResponseTime.Set(restime)
	hs.RequestBodySize.Set(reqBodySize)
	hs.ResponseBodySize.Set(resBodySize)
}

func (hs *httpStat) setStatus(status int) {
	if status >= 100 && status <= 199 {
		hs.Status1xx++
	} else if status >= 200 && status <= 299 {
		hs.Status2xx++
	} else if status >= 300 && status <= 399 {
		hs.Status3xx++
	} else if status >= 400 && status <= 499 {
		hs.Status4xx++
	} else if status >= 500 && status <= 599 {
		hs.Status5xx++
	}
}

func (hs *httpStat) StrStatus1xx() string {
	return fmt.Sprint(hs.Status1xx)
}

func (hs *httpStat) StrStatus2xx() string {
	return fmt.Sprint(hs.Status2xx)
}

func (hs *httpStat) StrStatus3xx() string {
	return fmt.Sprint(hs.Status3xx)
}

func (hs *httpStat) StrStatus4xx() string {
	return fmt.Sprint(hs.Status4xx)
}

func (hs *httpStat) StrStatus5xx() string {
	return fmt.Sprint(hs.Status5xx)
}

func (hs *httpStat) Count() int {
	return hs.Cnt
}

func (hs *httpStat) StrCount() string {
	return fmt.Sprint(hs.Cnt)
}

func (hs *httpStat) MaxResponseTime() float64 {
	return hs.ResponseTime.Max
}

func (hs *httpStat) MinResponseTime() float64 {
	return hs.ResponseTime.Min
}

func (hs *httpStat) SumResponseTime() float64 {
	return hs.ResponseTime.Sum
}

func (hs *httpStat) AvgResponseTime() float64 {
	return hs.ResponseTime.Avg(hs.Cnt)
}

func (hs *httpStat) P1ResponseTime() float64 {
	return hs.ResponseTime.P1(hs.Cnt)
}

func (hs *httpStat) P50ResponseTime() float64 {
	return hs.ResponseTime.P50(hs.Cnt)
}

func (hs *httpStat) P90ResponseTime() float64 {
	return hs.ResponseTime.P90(hs.Cnt)
}

func (hs *httpStat) P99ResponseTime() float64 {
	return hs.ResponseTime.P99(hs.Cnt)
}

func (hs *httpStat) StddevResponseTime() float64 {
	return hs.ResponseTime.Stddev(hs.Cnt)
}

// request
func (hs *httpStat) MaxRequestBodySize() float64 {
	return hs.RequestBodySize.Max
}

func (hs *httpStat) MinRequestBodySize() float64 {
	return hs.RequestBodySize.Min
}

func (hs *httpStat) SumRequestBodySize() float64 {
	return hs.RequestBodySize.Sum
}

func (hs *httpStat) AvgRequestBodySize() float64 {
	return hs.RequestBodySize.Avg(hs.Cnt)
}

func (hs *httpStat) P1RequestBodySize() float64 {
	return hs.RequestBodySize.P1(hs.Cnt)
}

func (hs *httpStat) P50RequestBodySize() float64 {
	return hs.RequestBodySize.P50(hs.Cnt)
}

func (hs *httpStat) P90RequestBodySize() float64 {
	return hs.RequestBodySize.P90(hs.Cnt)
}

func (hs *httpStat) P99RequestBodySize() float64 {
	return hs.RequestBodySize.P99(hs.Cnt)
}

func (hs *httpStat) StddevRequestBodySize() float64 {
	return hs.RequestBodySize.Stddev(hs.Cnt)
}

// response
func (hs *httpStat) MaxResponseBodySize() float64 {
	return hs.RequestBodySize.Max
}

func (hs *httpStat) MinResponseBodySize() float64 {
	return hs.RequestBodySize.Min
}

func (hs *httpStat) SumResponseBodySize() float64 {
	return hs.RequestBodySize.Sum
}

func (hs *httpStat) AvgResponseBodySize() float64 {
	return hs.RequestBodySize.Avg(hs.Cnt)
}

func (hs *httpStat) P1ResponseBodySize() float64 {
	return hs.RequestBodySize.P1(hs.Cnt)
}

func (hs *httpStat) P50ResponseBodySize() float64 {
	return hs.RequestBodySize.P50(hs.Cnt)
}

func (hs *httpStat) P90ResponseBodySize() float64 {
	return hs.RequestBodySize.P90(hs.Cnt)
}

func (hs *httpStat) P99ResponseBodySize() float64 {
	return hs.RequestBodySize.P99(hs.Cnt)
}

func (hs *httpStat) StddevResponseBodySize() float64 {
	return hs.RequestBodySize.Stddev(hs.Cnt)
}

func percentRank(l int, n int) int {
	pLen := (l * n / 100) - 1
	if pLen < 0 {
		pLen = 0
	}

	return pLen
}

type responseTime struct {
	Max           float64
	Min           float64
	Sum           float64
	usePercentile bool
	Percentiles   []float64
}

func newResponseTime(usePercentile bool) *responseTime {
	return &responseTime{
		usePercentile: usePercentile,
		Percentiles:   make([]float64, 0),
	}
}

func (res *responseTime) Set(val float64) {
	if res.Max < val {
		res.Max = val
	}

	if res.Min >= val || res.Min == 0 {
		res.Min = val
	}

	res.Sum += val

	if res.usePercentile {
		res.Percentiles = append(res.Percentiles, val)
	}
}

func (res *responseTime) Avg(cnt int) float64 {
	return res.Sum / float64(cnt)
}

func (res *responseTime) P1(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 1)
	return res.Percentiles[plen]
}

func (res *responseTime) P50(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 50)
	return res.Percentiles[plen]
}

func (res *responseTime) P90(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 90)
	return res.Percentiles[plen]
}

func (res *responseTime) P99(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 99)
	return res.Percentiles[plen]
}

func (res *responseTime) Stddev(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	var stdd float64
	avg := res.Avg(cnt)
	n := float64(cnt)

	for _, v := range res.Percentiles {
		stdd += (v - avg) * (v - avg)
	}

	return math.Sqrt(stdd / n)
}

type bodySize struct {
	Max           float64
	Min           float64
	Sum           float64
	usePercentile bool
	percentiles   []float64
}

func newBodySize(usePercentile bool) *bodySize {
	return &bodySize{
		usePercentile: usePercentile,
		percentiles:   make([]float64, 0),
	}
}

func (body *bodySize) Set(val float64) {
	if body.Max < val {
		body.Max = val
	}

	if body.Min >= val || body.Min == 0.0 {
		body.Min = val
	}

	body.Sum += val

	if body.usePercentile {
		body.percentiles = append(body.percentiles, val)
	}
}

func (body *bodySize) Avg(cnt int) float64 {
	return body.Sum / float64(cnt)
}

func (body *bodySize) P1(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 1)
	return body.percentiles[plen]
}

func (body *bodySize) P50(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 50)
	return body.percentiles[plen]
}

func (body *bodySize) P90(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 90)
	return body.percentiles[plen]
}

func (body *bodySize) P99(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := percentRank(cnt, 99)
	return body.percentiles[plen]
}

func (body *bodySize) Stddev(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	var stdd float64
	avg := body.Avg(cnt)
	n := float64(cnt)

	for _, v := range body.percentiles {
		stdd += (v - avg) * (v - avg)
	}

	return math.Sqrt(stdd / n)
}
