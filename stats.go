package httpstats

import (
	"math"
	"fmt"
	"sync"
	)

type Hints struct {
	values map[string]int
	len int
	mu sync.RWMutex
}

func NewHints() *Hints {
	return &Hints{
		values: make(map[string]int),
	}
}

func (h *Hints) LoadOrStore(key string) int {
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
	hints *Hints
	stats []*HTTPStat
	useResponseTimePercentile bool
	useBodySizePercentile bool
}

func NewHTTPStats(useResTimePercentile, useBodySizePercentile bool) *HTTPStats {
	return &HTTPStats{
		hints: NewHints(),
		stats: make([]*HTTPStat, 0),
		useResponseTimePercentile: useResTimePercentile,
		useBodySizePercentile: useBodySizePercentile,
	}
}

func (hs *HTTPStats) Set(uri, method string, restime, body float64) {
	key := fmt.Sprintf("%s_%s", method, uri)

	idx := hs.hints.LoadOrStore(key)

	if idx >= len(hs.stats) {
		hs.stats = append(hs.stats, NewHTTPStat(uri, method, hs.useResponseTimePercentile, hs.useBodySizePercentile))
	}

	hs.stats[idx].Set(restime, body)
}

func (hs *HTTPStats) Stats() []*HTTPStat {
	return hs.stats
}

type HTTPStat struct {
	Uri         string
	Cnt         int
	Method      string
	responseTime *responseTime
	bodySize *bodySize
}

func NewHTTPStat(uri, method string, useResTimePercentile, useBodySizePercentile bool) *HTTPStat {
	return &HTTPStat{
		Uri: uri,
		Method: method,
		responseTime: newResponseTime(useResTimePercentile),
		bodySize: newBodySize(useBodySizePercentile),
	}
}

func (hs *HTTPStat) Set(restime, bodysize float64) {
	hs.Cnt++
	hs.responseTime.Set(restime)
	hs.bodySize.Set(bodysize)
}

func (hs *HTTPStat) Count() int {
	return hs.Cnt
}

func (hs *HTTPStat) MaxResponseTime() float64 {
	return hs.responseTime.Max()
}

func (hs *HTTPStat) MinResponseTime() float64 {
	return hs.responseTime.Min()
}

func (hs *HTTPStat) SumResponseTime() float64 {
	return hs.responseTime.Sum()
}

func (hs *HTTPStat) AvgResponseTime() float64 {
	return hs.responseTime.Avg(hs.Cnt)
}

func (hs *HTTPStat) P1ResponseTime() float64 {
	return hs.responseTime.P1(hs.Cnt)
}

func (hs *HTTPStat) P50ResponseTime() float64 {
	return hs.responseTime.P50(hs.Cnt)
}

func (hs *HTTPStat) P90ResponseTime() float64 {
	return hs.responseTime.P90(hs.Cnt)
}

func (hs *HTTPStat) P99ResponseTime() float64 {
	return hs.responseTime.P99(hs.Cnt)
}

func (hs *HTTPStat) StddevResponseTime() float64 {
	return hs.responseTime.Stddev(hs.Cnt)
}

func (hs *HTTPStat) MaxBodySize() float64 {
	return hs.bodySize.Max()
}

func (hs *HTTPStat) MinBodySize() float64 {
	return hs.bodySize.Min()
}

func (hs *HTTPStat) SumBodySize() float64 {
	return hs.bodySize.Sum()
}

func (hs *HTTPStat) AvgBodySize() float64 {
	return hs.bodySize.Avg(hs.Cnt)
}

func (hs *HTTPStat) P1BodySize() float64 {
	return hs.bodySize.P1(hs.Cnt)
}

func (hs *HTTPStat) P50BodySize() float64 {
	return hs.bodySize.P50(hs.Cnt)
}

func (hs *HTTPStat) P90BodySize() float64 {
	return hs.bodySize.P90(hs.Cnt)
}

func (hs *HTTPStat) P99BodySize() float64 {
	return hs.bodySize.P99(hs.Cnt)
}

func (hs *HTTPStat) StddevBodySize() float64 {
	return hs.bodySize.Stddev(hs.Cnt)
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