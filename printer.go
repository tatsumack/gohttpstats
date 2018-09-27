package httpstats

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"fmt"
	"strings"
)

var headers = map[string]string{
	"count": "Count",
	"method": "Method",
	"uri": "Uri",
	"status_1xx": "Status_1xx",
	"status_2xx": "Status_2xx",
	"status_3xx": "Status_3xx",
	"status_4xx": "Status_4xx",
	"status_5xx": "Status_5xx",
	"min": "Min",
	"max": "Max",
	"sum": "Sum",
	"avg": "Avg",
	"p1": "P1",
	"p50": "P50",
	"p99": "P99",
	"stddev": "Stddev",
	"min_body": "Min(Body)",
	"max_body": "Max(Body)",
	"sum_body": "Sum(Body)",
	"avg_body": "Avg(Body)",
}

var defaultHeaders = []string{
	"Count", "Method", "Uri", "1xx", "2xx", "3xx", "4xx", "5xx",
	"Min", "Max", "Sum", "Avg",
	"P1", "P50", "P99", "Stddev",
	"Min(Body)", "Max(Body)", "Sum(Body)", "Avg(Body)",
}

type PrintOption struct {
	format string
	noHeaders bool
	headers []string
}

func NewPrintOption() *PrintOption {
	return &PrintOption{
		format: "table",
		headers: defaultHeaders,
	}
}

func (hs *HTTPStats) Print() {
	switch hs.printOption.format {
	case "table":
		hs.printTable()
	case "tsv":
		hs.printTSV()
	}
}

func round(num float64) string {
	return fmt.Sprintf("%.3f", num)
}

func (hs *HTTPStats) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(hs.printOption.headers)
	for _, s := range hs.stats {
		data := []string{
			s.StrCount(), s.Method, s.Uri,
			s.StrStatus1xx(), s.StrStatus2xx(), s.StrStatus3xx(), s.StrStatus4xx(), s.StrStatus5xx(),
			round(s.MinResponseTime()), round(s.MaxResponseTime()),
			round(s.SumResponseTime()), round(s.AvgResponseTime()),
			round(s.P1ResponseTime()), round(s.P50ResponseTime()), round(s.P99ResponseTime()),
			round(s.StddevResponseTime()), round(s.MinResponseBodySize()), round(s.MaxResponseBodySize()), round(s.SumResponseBodySize()),round(s.AvgResponseBodySize()),
		}
		table.Append(data)
	}
	table.Render()
}

func (hs *HTTPStats) printTSV() {
	if !hs.printOption.noHeaders {
		fmt.Println(strings.Join(hs.printOption.headers, `\t`))
	}
	for _, s := range hs.stats {
		data := []string{
			s.StrCount(), s.Method, s.Uri,
			s.StrStatus1xx(), s.StrStatus2xx(), s.StrStatus3xx(), s.StrStatus4xx(), s.StrStatus5xx(),
				round(s.MinResponseTime()), round(s.MaxResponseTime()),
				round(s.SumResponseTime()), round(s.AvgResponseTime()),
				round(s.P1ResponseTime()), round(s.P50ResponseTime()), round(s.P99ResponseTime()),
				round(s.StddevResponseTime()), round(s.MinResponseBodySize()), round(s.MaxResponseBodySize()), round(s.SumResponseBodySize()), round(s.AvgResponseBodySize()),
		}
		fmt.Println(strings.Join(data, `\t`))
	}
}
