package httpstats

import "errors"

var (
	SortCount                  = "Count"
	SortUri                    = "Uri"
	SortMethod                 = "Method"
	SortMaxResponseTime        = "MaxResponseTime"
	SortMinResponseTime        = "MinResponseTime"
	SortSumResponseTime        = "SumResponseTime"
	SortAvgResponseTime        = "AvgResponseTime"
	SortP1ResponseTime         = "P1ResponseTime"
	SortP50ResponseTime        = "P50ResponseTime"
	SortP90ResponseTime        = "P90ResponseTime"
	SortP99ResponseTime        = "P99ResponseTime"
	SortStddevResponseTime     = "StddevResponseTime"
	SortMaxRequestBodySize     = "MaxRequestBodySize"
	SortMinRequestBodySize     = "MinRequestBodySize"
	SortSumRequestBodySize     = "SumRequestBodySize"
	SortAvgRequestBodySize     = "AvgRequestBodySize"
	SortP1RequestBodySize      = "P1RequestBodySize"
	SortP50RequestBodySize     = "P50RequestBodySize"
	SortP90RequestBodySize     = "P90RequestBodySize"
	SortP99RequestBodySize     = "P99RequestBodySize"
	SortStddevRequestBodySize  = "StddevRequestBodySize"
	SortMaxResponseBodySize    = "MaxResponseBodySize"
	SortMinResponseBodySize    = "MinResponseBodySize"
	SortSumResponseBodySize    = "SumResponseBodySize"
	SortAvgResponseBodySize    = "AvgResponseBodySize"
	SortP1ResponseBodySize     = "P1ResponseBodySize"
	SortP50ResponseBodySize    = "P50ResponseBodySize"
	SortP90ResponseBodySize    = "P90ResponseBodySize"
	SortP99ResponseBodySize    = "P99ResponseBodySize"
	SortStddevResponseBodySize = "StddevResponseBodySize"

	SkipReadLineErr            = errors.New("Skip read line")
)
