package httpstats

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (hs *HTTPStats) LoadStats(r io.Reader) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var stats []*httpStat
	err = yaml.Unmarshal(buf, &stats)
	hs.stats = stats

	return err
}
