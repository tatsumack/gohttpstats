package httpstats

import "sort"

func (hs *HTTPStats) SortCount(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].Count() > hs.stats[j].Count()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].Count() < hs.stats[j].Count()
		})
	}
}

func (hs *HTTPStats) SortMaxResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxResponseTime() > hs.stats[j].MaxResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxResponseTime() < hs.stats[j].MaxResponseTime()
		})
	}
}

func (hs *HTTPStats) SortMinResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinResponseTime() > hs.stats[j].MinResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinResponseTime() < hs.stats[j].MinResponseTime()
		})
	}
}

func (hs *HTTPStats) SortSumResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumResponseTime() > hs.stats[j].SumResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumResponseTime() < hs.stats[j].SumResponseTime()
		})
	}
}

func (hs *HTTPStats) AvgResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgResponseTime() > hs.stats[j].AvgResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgResponseTime() < hs.stats[j].AvgResponseTime()
		})
	}
}

func (hs *HTTPStats) SortP1ResponseTime(reverse bool)  {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1ResponseTime() > hs.stats[j].P1ResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1ResponseTime() < hs.stats[j].P1ResponseTime()
		})
	}
}

func (hs *HTTPStats) SortP50ResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50ResponseTime() > hs.stats[j].P50ResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50ResponseTime() < hs.stats[j].P50ResponseTime()
		})
	}
}

func (hs *HTTPStats) SortP90ResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90ResponseTime() > hs.stats[j].P90ResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90ResponseTime() < hs.stats[j].P90ResponseTime()
		})
	}
}

func (hs *HTTPStats) SortP99ResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99ResponseTime() > hs.stats[j].P99ResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99ResponseTime() < hs.stats[j].P99ResponseTime()
		})
	}
}

func (hs *HTTPStats) SortStddevResponseTime(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevResponseTime() > hs.stats[j].StddevResponseTime()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevResponseTime() < hs.stats[j].StddevResponseTime()
		})
	}
}

// request
func (hs *HTTPStats) SortMaxRequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxRequestBodySize() > hs.stats[j].MaxRequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxRequestBodySize() < hs.stats[j].MaxRequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortMinRequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinRequestBodySize() > hs.stats[j].MinRequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinRequestBodySize() < hs.stats[j].MinRequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortSumRequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumRequestBodySize() > hs.stats[j].SumRequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumRequestBodySize() < hs.stats[j].SumRequestBodySize()
		})
	}
}

func (hs *HTTPStats) AvgRequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgRequestBodySize() > hs.stats[j].AvgRequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgRequestBodySize() < hs.stats[j].AvgRequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortP1RequestBodySize(reverse bool)  {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1RequestBodySize() > hs.stats[j].P1RequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1RequestBodySize() < hs.stats[j].P1RequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortP50RequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50RequestBodySize() > hs.stats[j].P50RequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50RequestBodySize() < hs.stats[j].P50RequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortP90RequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90RequestBodySize() > hs.stats[j].P90RequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90RequestBodySize() < hs.stats[j].P90RequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortP99RequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99RequestBodySize() > hs.stats[j].P99RequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99RequestBodySize() < hs.stats[j].P99RequestBodySize()
		})
	}
}

func (hs *HTTPStats) SortStddevRequestBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevRequestBodySize() > hs.stats[j].StddevRequestBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevRequestBodySize() < hs.stats[j].StddevRequestBodySize()
		})
	}
}


// response
func (hs *HTTPStats) SortMaxResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxResponseBodySize() > hs.stats[j].MaxResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxResponseBodySize() < hs.stats[j].MaxResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortMinResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinResponseBodySize() > hs.stats[j].MinResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinResponseBodySize() < hs.stats[j].MinResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortSumResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumResponseBodySize() > hs.stats[j].SumResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumResponseBodySize() < hs.stats[j].SumResponseBodySize()
		})
	}
}

func (hs *HTTPStats) AvgResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgResponseBodySize() > hs.stats[j].AvgResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgResponseBodySize() < hs.stats[j].AvgResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortP1ResponseBodySize(reverse bool)  {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1ResponseBodySize() > hs.stats[j].P1ResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1ResponseBodySize() < hs.stats[j].P1ResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortP50ResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50ResponseBodySize() > hs.stats[j].P50ResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50ResponseBodySize() < hs.stats[j].P50ResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortP90ResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90ResponseBodySize() > hs.stats[j].P90ResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90ResponseBodySize() < hs.stats[j].P90ResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortP99ResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99ResponseBodySize() > hs.stats[j].P99ResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99ResponseBodySize() < hs.stats[j].P99ResponseBodySize()
		})
	}
}

func (hs *HTTPStats) SortStddevResponseBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevResponseBodySize() > hs.stats[j].StddevResponseBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevResponseBodySize() < hs.stats[j].StddevResponseBodySize()
		})
	}
}
