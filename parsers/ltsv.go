package parsers

import (
	"io"

	"github.com/najeira/ltsv"
)

type LTSVParser struct {
	reader *ltsv.Reader
}

func NewLTSVParser(r io.Reader) Parser {
	return &LTSVParser{
		reader: ltsv.NewReader(r),
	}
}

func (l *LTSVParser) Read() (map[string]string, error) {
	return l.reader.Read()
}
