package httpstats

import (
	"io"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func splitCSV(val string) []string {
	strs := strings.Split(val, ",")
	if len(strs) == 1 && strs[0] == "" {
		return []string{}
	}

	trimedStrs := make([]string, 0, len(strs))

	for _, s := range strs {
		trimedStrs = append(trimedStrs, strings.Trim(s, " "))
	}

	return trimedStrs
}

type Options struct {
	File              string   `yaml:"file"`
	Sort              string   `yaml:"sort"`
	Reverse           bool     `yaml:"reverse"`
	QueryString       bool     `yaml:"query_string"`
	Tsv               bool     `yaml:"tsv"`
	NoHeaders         bool     `yaml:no_headers`
	ApptimeLabel      string   `yaml:"apptime_label"`
	ReqtimeLabel      string   `yaml:"reqtime_label"`
	StatusLabel       string   `yaml:"status_label"`
	SizeLabel         string   `yaml:"size_label"`
	MethodLabel       string   `yaml:"method_label"`
	UriLabel          string   `yaml:"uri_label"`
	TimeLabel         string   `yaml:"time_label"`
	Limit             int      `yaml:"limit"`
	Includes          []string `yaml:"includes"`
	Excludes          []string `yaml:"excludes"`
	IncludeStatuses   []string `yaml:"include_statuses"`
	ExcludeStatuses   []string `yaml:"exclude_statuses"`
	Aggregates        []string `yaml:"aggregates"`
	StartTime         string   `yaml:"start_time"`
	EndTime           string   `yaml:"end_time"`
	StartTimeDuration string   `yaml:"start_time_duration"`
	EndTimeDuration   string   `yaml:"end_time_duration"`
}

func NewDefaultOptions() *Options {
	return &Options{
		Sort:         "max",
		ApptimeLabel: "apptime",
		ReqtimeLabel: "reqtime",
		StatusLabel:  "status",
		SizeLabel:    "size",
		MethodLabel:  "method",
		UriLabel:     "uri",
		TimeLabel:    "time",
		Limit:        5000,
	}
}

func NewCliOptions(cli *Options, includes, excludes, includeStatuses, excludeStatuses, aggregates string) *Options {
	i := splitCSV(includes)
	if len(i) > 0 {
		cli.Includes = i
	}

	e := splitCSV(excludes)
	if len(e) > 0 {
		cli.Excludes = e
	}

	is := splitCSV(includeStatuses)
	if len(is) > 0 {
		cli.IncludeStatuses = is
	}

	es := splitCSV(excludeStatuses)
	if len(es) > 0 {
		cli.ExcludeStatuses = es
	}

	a := splitCSV(aggregates)
	if len(a) > 0 {
		cli.Aggregates = a
	}

	return cli
}

// cli options > config file options > default options
func MergeOptions(cli, file, def *Options) *Options {
	options := &Options{}

	if cli.File != "" {
		options.File = cli.File
	} else if file.File != "" {
		options.File = file.File
	} else {
		options.File = def.File
	}

	if cli.Sort != "" {
		options.Sort = cli.Sort
	} else if file.Sort != "" {
		options.Sort = file.Sort
	} else {
		options.Sort = def.Sort
	}

	if cli.Reverse {
		options.Reverse = true
	} else if file.Reverse {
		options.Reverse = true
	} else {
		options.Reverse = def.Reverse
	}

	if cli.QueryString {
		options.QueryString = true
	} else if file.QueryString {
		options.QueryString = true
	} else {
		options.QueryString = def.QueryString
	}

	if cli.Tsv {
		options.Tsv = true
	} else if file.Tsv {
		options.Tsv = true
	} else {
		options.Tsv = def.Tsv
	}

	if cli.NoHeaders {
		options.NoHeaders = true
	} else if file.NoHeaders {
		options.NoHeaders = true
	} else {
		options.NoHeaders = def.NoHeaders
	}

	if cli.ApptimeLabel != "" {
		options.ApptimeLabel = cli.ApptimeLabel
	} else if file.ApptimeLabel != "" {
		options.ApptimeLabel = file.ApptimeLabel
	} else {
		options.ApptimeLabel = def.ApptimeLabel
	}

	if cli.ReqtimeLabel != "" {
		options.ReqtimeLabel = cli.ReqtimeLabel
	} else if file.ReqtimeLabel != "" {
		options.ReqtimeLabel = file.ReqtimeLabel
	} else {
		options.ReqtimeLabel = def.ReqtimeLabel
	}

	if cli.StatusLabel != "" {
		options.StatusLabel = cli.StatusLabel
	} else if file.StatusLabel != "" {
		options.StatusLabel = file.StatusLabel
	} else {
		options.StatusLabel = def.StatusLabel
	}

	if cli.SizeLabel != "" {
		options.SizeLabel = cli.SizeLabel
	} else if file.SizeLabel != "" {
		options.SizeLabel = file.SizeLabel
	} else {
		options.SizeLabel = def.SizeLabel
	}

	if cli.MethodLabel != "" {
		options.MethodLabel = cli.MethodLabel
	} else if file.MethodLabel != "" {
		options.MethodLabel = file.MethodLabel
	} else {
		options.MethodLabel = def.MethodLabel
	}

	if cli.UriLabel != "" {
		options.UriLabel = cli.UriLabel
	} else if file.UriLabel != "" {
		options.UriLabel = file.UriLabel
	} else {
		options.UriLabel = def.UriLabel
	}

	if cli.TimeLabel != "" {
		options.TimeLabel = cli.TimeLabel
	} else if file.TimeLabel != "" {
		options.TimeLabel = file.TimeLabel
	} else {
		options.TimeLabel = def.TimeLabel
	}

	if cli.Limit != 0 {
		options.Limit = cli.Limit
	} else if file.Limit != 0 {
		options.Limit = file.Limit
	} else {
		options.Limit = def.Limit
	}

	if len(cli.Includes) > 0 {
		options.Includes = cli.Includes
	} else if len(file.Includes) > 0 {
		options.Includes = file.Includes
	} else {
		options.Includes = def.Includes
	}

	if len(cli.Excludes) > 0 {
		options.Excludes = cli.Excludes
	} else if len(file.Excludes) > 0 {
		options.Excludes = file.Excludes
	} else {
		options.Excludes = def.Excludes
	}

	if len(cli.IncludeStatuses) > 0 {
		options.IncludeStatuses = cli.IncludeStatuses
	} else if len(file.IncludeStatuses) > 0 {
		options.IncludeStatuses = file.IncludeStatuses
	} else {
		options.IncludeStatuses = def.IncludeStatuses
	}

	if len(cli.ExcludeStatuses) > 0 {
		options.ExcludeStatuses = cli.ExcludeStatuses
	} else if len(file.ExcludeStatuses) > 0 {
		options.ExcludeStatuses = file.ExcludeStatuses
	} else {
		options.ExcludeStatuses = def.ExcludeStatuses
	}

	if len(cli.Aggregates) > 0 {
		options.Aggregates = cli.Aggregates
	} else if len(file.Aggregates) > 0 {
		options.Aggregates = file.Aggregates
	} else {
		options.Aggregates = def.Aggregates
	}

	if cli.StartTime != "" {
		options.StartTime = cli.StartTime
	} else if file.StartTime != "" {
		options.StartTime = file.StartTime
	} else {
		options.StartTime = def.StartTime
	}

	if cli.EndTime != "" {
		options.EndTime = cli.EndTime
	} else if file.EndTime != "" {
		options.EndTime = file.EndTime
	} else {
		options.EndTime = def.EndTime
	}

	if cli.StartTimeDuration != "" {
		options.StartTimeDuration = cli.StartTimeDuration
	} else if file.StartTimeDuration != "" {
		options.StartTimeDuration = file.StartTimeDuration
	} else {
		options.StartTimeDuration = def.StartTimeDuration
	}

	if cli.EndTimeDuration != "" {
		options.EndTimeDuration = cli.EndTimeDuration
	} else if file.EndTimeDuration != "" {
		options.EndTimeDuration = file.EndTimeDuration
	} else {
		options.EndTimeDuration = def.EndTimeDuration
	}

	return options
}

func LoadOptionsFromReader(r io.Reader) (*Options, error) {
	var opts *Options
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return opts, err
	}

	err = yaml.Unmarshal(buf, opts)

	return opts, err
}
