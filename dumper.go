package httpstats

import (
	"gopkg.in/yaml.v2"
	"io"
)

func (hs *HTTPStats) DumpStats(w io.Writer) error {
	buf, err := yaml.Marshal(&hs.stats)
	if err != nil {
		return err
	}

	_, err = w.Write(buf)

	return err
}
