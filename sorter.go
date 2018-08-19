package httpstats

import "sort"

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

func (hs *HTTPStats) SortMaxBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxBodySize() > hs.stats[j].MaxBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MaxBodySize() < hs.stats[j].MaxBodySize()
		})
	}
}

func (hs *HTTPStats) SortMinBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinBodySize() > hs.stats[j].MinBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].MinBodySize() < hs.stats[j].MinBodySize()
		})
	}
}

func (hs *HTTPStats) SortSumBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumBodySize() > hs.stats[j].SumBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].SumBodySize() < hs.stats[j].SumBodySize()
		})
	}
}

func (hs *HTTPStats) AvgBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgBodySize() > hs.stats[j].AvgBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].AvgBodySize() < hs.stats[j].AvgBodySize()
		})
	}
}

func (hs *HTTPStats) SortP1BodySize(reverse bool)  {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1BodySize() > hs.stats[j].P1BodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P1BodySize() < hs.stats[j].P1BodySize()
		})
	}
}

func (hs *HTTPStats) SortP50BodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50BodySize() > hs.stats[j].P50BodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P50BodySize() < hs.stats[j].P50BodySize()
		})
	}
}

func (hs *HTTPStats) SortP90BodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90BodySize() > hs.stats[j].P90BodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P90BodySize() < hs.stats[j].P90BodySize()
		})
	}
}

func (hs *HTTPStats) SortP99BodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99BodySize() > hs.stats[j].P99BodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].P99BodySize() < hs.stats[j].P99BodySize()
		})
	}
}

func (hs *HTTPStats) SortStddevBodySize(reverse bool) {
	if reverse {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevBodySize() > hs.stats[j].StddevBodySize()
		})
	} else {
		sort.Slice(hs.stats, func(i, j int) bool {
			return hs.stats[i].StddevBodySize() < hs.stats[j].StddevBodySize()
		})
	}
}
