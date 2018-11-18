package parsers

import (
	"fmt"
	"io"
	"net/url"

	"github.com/najeira/ltsv"
	"github.com/tkuchiki/gohttpstats"
)

type LTSVParser struct {
	reader      *ltsv.Reader
	label       *LTSVLabel
	strictMode  bool
	queryString bool
}

type LTSVLabel struct {
	Uri     string
	Apptime string
	Reqtime string
	Size    string
	Status  string
	Method  string
	Time    string
}

func NewLTSVLabel(uri, apptime, reqtime, size, status, method, time string) *LTSVLabel {
	return &LTSVLabel{
		Uri:     uri,
		Apptime: apptime,
		Reqtime: reqtime,
		Size:    size,
		Status:  status,
		Method:  method,
		Time:    time,
	}
}

func NewLTSVParser(r io.Reader, l *LTSVLabel, query bool) *LTSVParser {
	return &LTSVParser{
		reader:      ltsv.NewReader(r),
		label:       l,
		queryString: query,
	}
}

func (l *LTSVParser) Parse() (*HTTPStat, error) {
	parsedValue, err := l.reader.Read()
	if err != nil && l.strictMode {
		return &HTTPStat{}, err
	} else if err == io.EOF {
		return &HTTPStat{}, err
	}

	u, err := url.Parse(parsedValue[l.label.Uri])
	if err != nil {
		return &HTTPStat{}, errSkipReadLine(l.strictMode, err)
	}
	var uri string
	if l.queryString {
		v := url.Values{}
		values := u.Query()
		for q := range values {
			v.Set(q, "xxx")
		}
		uri = fmt.Sprintf("%s?%s", u.Path, v.Encode())
	} else {
		uri = u.Path
	}

	resTime, err := httpstats.StringToFloat64(parsedValue[l.label.Apptime])
	if err != nil {
		var reqTime float64
		reqTime, err = httpstats.StringToFloat64(parsedValue[l.label.Reqtime])
		if err != nil {
			return &HTTPStat{}, errSkipReadLine(l.strictMode, err)
		}

		resTime = reqTime
	}

	bodySize, err := httpstats.StringToFloat64(parsedValue[l.label.Size])
	if err != nil {
		return &HTTPStat{}, errSkipReadLine(l.strictMode, err)
	}

	status, err := httpstats.StringToInt(parsedValue[l.label.Status])
	if err != nil {
		return &HTTPStat{}, errSkipReadLine(l.strictMode, err)
	}

	method := parsedValue[l.label.Method]
	timestr := parsedValue[l.label.Time]

	return NewHTTPStat(uri, method, timestr, resTime, bodySize, status), nil
}
