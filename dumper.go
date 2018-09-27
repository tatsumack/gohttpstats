package httpstats

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	)

func (hs *HTTPStats) DumpStats(filename string) error {
	buf, err := yaml.Marshal(&hs.stats)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, buf, os.ModePerm)

	return err
}