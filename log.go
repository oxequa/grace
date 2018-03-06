package grace

import (
	"io/ioutil"
)

const fname = "errors.log"

type logger struct {
	err    Error
	enable bool
	fname  string
}

func (l *logger) state() bool {
	return l.enable
}

func (l *logger) name() string {
	return l.fname
}

func (l *logger) save(data []byte) error {
	err := ioutil.WriteFile(l.fname, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
