package httpstats

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"fmt"
	"strings"
)

var headers = map[string]string{
	"count": "Count",
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
	"method": "Method",
	"uri": "Uri",
}

var defaultHeaders = []string{
	"Count", "Min", "Max", "Sum", "Avg",
	"P1", "P50", "P99", "Stddev",
	"Min(Body)", "Max(Body)", "Sum(Body)", "Avg(Body)",
	"Method", "Uri",
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
			fmt.Sprint(s.Count()), round(s.MinResponseTime()), round(s.MaxResponseTime()),
			round(s.SumResponseTime()), round(s.AvgResponseTime()),
			round(s.P1ResponseTime()), round(s.P50ResponseTime()), round(s.P99ResponseTime()),
			round(s.StddevResponseTime()),
			round(s.MinResponseBodySize()), round(s.MaxResponseBodySize()), round(s.SumResponseBodySize()), round(s.AvgResponseBodySize()),
			s.Method(), s.Uri()}
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
			fmt.Sprint(s.Count()), round(s.MinResponseTime()), round(s.MaxResponseTime()),
			round(s.SumResponseTime()), round(s.AvgResponseTime()),
			round(s.P1ResponseTime()), round(s.P50ResponseTime()), round(s.P99ResponseTime()),
			round(s.StddevResponseTime()),
			round(s.MinResponseBodySize()), round(s.MaxResponseBodySize()), round(s.SumResponseBodySize()), round(s.AvgResponseBodySize()),
			s.Method(), s.Uri(),
		}
		fmt.Println(strings.Join(data, `\t`))
	}
}
