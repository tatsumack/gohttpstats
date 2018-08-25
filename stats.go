package httpstats

import (
	"math"
	"fmt"
	"sync"
	)

type hints struct {
	values map[string]int
	len int
	mu sync.RWMutex
}

func newHints() *hints {
	return &hints{
		values: make(map[string]int),
	}
}

func (h *hints) loadOrStore(key string) int {
	_, ok := h.values[key]
	if !ok {
		h.mu.Lock()
		h.values[key] = h.len
		h.len++
		h.mu.Unlock()
	}

	return h.values[key]
}

type HTTPStats struct {
	hints *hints
	stats []*httpStat
	useResponseTimePercentile bool
	useRequestBodySizePercentile bool
	useResponseBodySizePercentile bool
	printOption *PrintOption
}

func NewHTTPStats(useResTimePercentile, useRequestBodySizePercentile, useResponseBodySizePercentile bool, po *PrintOption) *HTTPStats {
	return &HTTPStats{
		hints: newHints(),
		stats: make([]*httpStat, 0),
		useResponseTimePercentile: useResTimePercentile,
		useResponseBodySizePercentile: useResponseBodySizePercentile,
		printOption: po,
	}
}

func (hs *HTTPStats) Set(uri, method string, status int, restime, resBodySize, reqBodySize float64) {
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

type httpStat struct {
	uri         string
	cnt         int
	status1xx int
	status2xx int
	status3xx int
	status4xx int
	status5xx int
	method      string
	responseTime *responseTime
	requestBodySize *bodySize
	responseBodySize *bodySize

}

func newHTTPStat(uri, method string, useResTimePercentile, useRequestBodySizePercentile, useResponseBodySizePercentile bool) *httpStat {
	return &httpStat{
		uri: uri,
		method: method,
		responseTime: newResponseTime(useResTimePercentile),
		requestBodySize: newBodySize(useRequestBodySizePercentile),
		responseBodySize: newBodySize(useResponseBodySizePercentile),
	}
}

func (hs *httpStat) Set(status int, restime, reqBodySize, resBodySize float64) {
	hs.cnt++
	hs.setStatus(status)
	hs.responseTime.Set(restime)
	hs.requestBodySize.Set(reqBodySize)
	hs.responseBodySize.Set(resBodySize)
}

func (hs *httpStat) setStatus(status int) {
	if status >= 100 && status <= 199 {
		hs.status1xx++
	} else if status >= 200 && status <= 299 {
		hs.status2xx++
	} else if status >= 300 && status <= 399 {
		hs.status3xx++
	} else if status >= 400 && status <= 499 {
		hs.status4xx++
	} else if status >= 500 && status <= 599 {
		hs.status5xx++
	}
}

func (hs *httpStat) Status1xx() int {
	return hs.status1xx
}

func (hs *httpStat) StrStatus1xx() string {
	return fmt.Sprint(hs.status1xx)
}

func (hs *httpStat) Status2xx() int {
	return hs.status2xx
}

func (hs *httpStat) StrStatus2xx() string {
	return fmt.Sprint(hs.status2xx)
}

func (hs *httpStat) Status3xx() int {
	return hs.status3xx
}

func (hs *httpStat) StrStatus3xx() string {
	return fmt.Sprint(hs.status3xx)
}

func (hs *httpStat) Status4xx() int {
	return hs.status4xx
}

func (hs *httpStat) StrStatus4xx() string {
	return fmt.Sprint(hs.status4xx)
}

func (hs *httpStat) Status5xx() int {
	return hs.status5xx
}

func (hs *httpStat) StrStatus5xx() string {
	return fmt.Sprint(hs.status5xx)
}

func (hs *httpStat) Count() int {
	return hs.cnt
}

func (hs *httpStat) StrCount() string {
	return fmt.Sprint(hs.cnt)
}

func (hs *httpStat) Uri() string {
	return hs.uri
}

func (hs *httpStat) Method() string {
	return hs.method
}

func (hs *httpStat) MaxResponseTime() float64 {
	return hs.responseTime.Max()
}

func (hs *httpStat) MinResponseTime() float64 {
	return hs.responseTime.Min()
}

func (hs *httpStat) SumResponseTime() float64 {
	return hs.responseTime.Sum()
}

func (hs *httpStat) AvgResponseTime() float64 {
	return hs.responseTime.Avg(hs.cnt)
}

func (hs *httpStat) P1ResponseTime() float64 {
	return hs.responseTime.P1(hs.cnt)
}

func (hs *httpStat) P50ResponseTime() float64 {
	return hs.responseTime.P50(hs.cnt)
}

func (hs *httpStat) P90ResponseTime() float64 {
	return hs.responseTime.P90(hs.cnt)
}

func (hs *httpStat) P99ResponseTime() float64 {
	return hs.responseTime.P99(hs.cnt)
}

func (hs *httpStat) StddevResponseTime() float64 {
	return hs.responseTime.Stddev(hs.cnt)
}

// request
func (hs *httpStat) MaxRequestBodySize() float64 {
	return hs.requestBodySize.Max()
}

func (hs *httpStat) MinRequestBodySize() float64 {
	return hs.requestBodySize.Min()
}

func (hs *httpStat) SumRequestBodySize() float64 {
	return hs.requestBodySize.Sum()
}

func (hs *httpStat) AvgRequestBodySize() float64 {
	return hs.requestBodySize.Avg(hs.cnt)
}

func (hs *httpStat) P1RequestBodySize() float64 {
	return hs.requestBodySize.P1(hs.cnt)
}

func (hs *httpStat) P50RequestBodySize() float64 {
	return hs.requestBodySize.P50(hs.cnt)
}

func (hs *httpStat) P90RequestBodySize() float64 {
	return hs.requestBodySize.P90(hs.cnt)
}

func (hs *httpStat) P99RequestBodySize() float64 {
	return hs.requestBodySize.P99(hs.cnt)
}

func (hs *httpStat) StddevRequestBodySize() float64 {
	return hs.requestBodySize.Stddev(hs.cnt)
}

// response
func (hs *httpStat) MaxResponseBodySize() float64 {
	return hs.requestBodySize.Max()
}

func (hs *httpStat) MinResponseBodySize() float64 {
	return hs.requestBodySize.Min()
}

func (hs *httpStat) SumResponseBodySize() float64 {
	return hs.requestBodySize.Sum()
}

func (hs *httpStat) AvgResponseBodySize() float64 {
	return hs.requestBodySize.Avg(hs.cnt)
}

func (hs *httpStat) P1ResponseBodySize() float64 {
	return hs.requestBodySize.P1(hs.cnt)
}

func (hs *httpStat) P50ResponseBodySize() float64 {
	return hs.requestBodySize.P50(hs.cnt)
}

func (hs *httpStat) P90ResponseBodySize() float64 {
	return hs.requestBodySize.P90(hs.cnt)
}

func (hs *httpStat) P99ResponseBodySize() float64 {
	return hs.requestBodySize.P99(hs.cnt)
}

func (hs *httpStat) StddevResponseBodySize() float64 {
	return hs.requestBodySize.Stddev(hs.cnt)
}

func lenPercentile(l int, n int) int {
	pLen := (l * n / 100) - 1
	if pLen < 0 {
		pLen = 0
	}

	return pLen
}

type responseTime struct {
	max float64
	min float64
	sum float64
	cnt int
	usePercentile bool
	percentiles []float64
}

func newResponseTime(usePercentile bool) *responseTime {
	return &responseTime{
		usePercentile: usePercentile,
		percentiles: make([]float64, 0),
	}
}

func (res *responseTime) Set(val float64) {
	if res.max < val {
		res.max = val
	}

	if res.min >= val || res.min == 0 {
		res.min = val
	}

	res.sum += val

	if res.usePercentile {
		res.percentiles = append(res.percentiles, val)
	}
}

func (res *responseTime) Max() float64 {
	return res.max
}

func (res *responseTime) Min() float64{
	return res.min
}

func (res *responseTime) Sum() float64 {
	return res.sum
}

func (res *responseTime) Avg(cnt int) float64 {
	return res.sum / float64(cnt)
}

func (res *responseTime) P1(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 1)
	return res.percentiles[plen]
}

func (res *responseTime) P50(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 50)
	return res.percentiles[plen]
}

func (res *responseTime) P90(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 90)
	return res.percentiles[plen]
}

func (res *responseTime) P99(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 99)
	return res.percentiles[plen]
}

func (res *responseTime) Stddev(cnt int) float64 {
	if !res.usePercentile {
		return 0.0
	}

	var stdd float64
	avg := res.Avg(cnt)
	n := float64(cnt)

	for _, v := range res.percentiles {
		stdd += (v - avg) * (v - avg)
	}

	return math.Sqrt(stdd / n)
}

type bodySize struct {
	max float64
	min float64
	sum float64
	usePercentile bool
	percentiles []float64
}

func newBodySize(usePercentile bool) *bodySize {
	return &bodySize{
		usePercentile: usePercentile,
		percentiles: make([]float64, 0),
	}
}

func (body *bodySize) Set(val float64) {
	if body.max < val {
		body.max = val
	}

	if body.min >= val || body.min == 0.0 {
		body.min = val
	}

	body.sum += val

	if body.usePercentile {
		body.percentiles = append(body.percentiles, val)
	}
}

func (body *bodySize) Max() float64 {
	return body.max
}

func (body *bodySize) Min() float64{
	return body.min
}

func (body *bodySize) Sum() float64 {
	return body.sum
}

func (body *bodySize) Avg(cnt int) float64 {
	return body.sum / float64(cnt)
}

func (body *bodySize) P1(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 1)
	return body.percentiles[plen]
}

func (body *bodySize) P50(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 50)
	return body.percentiles[plen]
}

func (body *bodySize) P90(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 90)
	return body.percentiles[plen]
}

func (body *bodySize) P99(cnt int) float64 {
	if !body.usePercentile {
		return 0.0
	}

	plen := lenPercentile(cnt, 99)
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