package httpstats

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func (hs *HTTPStats) LoadStats(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var stats []*httpStat
	err = yaml.Unmarshal(buf, &stats)
	hs.stats = stats

	return err
}

