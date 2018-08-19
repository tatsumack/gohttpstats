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
		defer h.mu.Unlock()
		h.mu.Lock()
		h.values[key] = h.len
		h.len++
	}

	return h.values[key]
}

type HTTPStats struct {
	hints *hints
	stats []*httpStat
	useResponseTimePercentile bool
	useBodySizePercentile bool
	printOption *PrintOption
}

func NewHTTPStats(useResTimePercentile, useBodySizePercentile bool, po *PrintOption) *HTTPStats {
	return &HTTPStats{
		hints: newHints(),
		stats: make([]*httpStat, 0),
		useResponseTimePercentile: useResTimePercentile,
		useBodySizePercentile: useBodySizePercentile,
		printOption: po,
	}
}

func (hs *HTTPStats) Set(uri, method string, restime, body float64) {
	key := fmt.Sprintf("%s_%s", method, uri)

	idx := hs.hints.loadOrStore(key)

	if idx >= len(hs.stats) {
		hs.stats = append(hs.stats, newHTTPStat(uri, method, hs.useResponseTimePercentile, hs.useBodySizePercentile))
	}

	hs.stats[idx].Set(restime, body)
}

func (hs *HTTPStats) Stats() []*httpStat {
	return hs.stats
}

type httpStat struct {
	uri         string
	cnt         int
	method      string
	responseTime *responseTime
	bodySize *bodySize
}

func newHTTPStat(uri, method string, useResTimePercentile, useBodySizePercentile bool) *httpStat {
	return &httpStat{
		uri: uri,
		method: method,
		responseTime: newResponseTime(useResTimePercentile),
		bodySize: newBodySize(useBodySizePercentile),
	}
}

func (hs *httpStat) Set(restime, bodysize float64) {
	hs.cnt++
	hs.responseTime.Set(restime)
	hs.bodySize.Set(bodysize)
}

func (hs *httpStat) Count() int {
	return hs.cnt
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

func (hs *httpStat) MaxBodySize() float64 {
	return hs.bodySize.Max()
}

func (hs *httpStat) MinBodySize() float64 {
	return hs.bodySize.Min()
}

func (hs *httpStat) SumBodySize() float64 {
	return hs.bodySize.Sum()
}

func (hs *httpStat) AvgBodySize() float64 {
	return hs.bodySize.Avg(hs.cnt)
}

func (hs *httpStat) P1BodySize() float64 {
	return hs.bodySize.P1(hs.cnt)
}

func (hs *httpStat) P50BodySize() float64 {
	return hs.bodySize.P50(hs.cnt)
}

func (hs *httpStat) P90BodySize() float64 {
	return hs.bodySize.P90(hs.cnt)
}

func (hs *httpStat) P99BodySize() float64 {
	return hs.bodySize.P99(hs.cnt)
}

func (hs *httpStat) StddevBodySize() float64 {
	return hs.bodySize.Stddev(hs.cnt)
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