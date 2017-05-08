package main

import (
	"io/ioutil"
)

func FileInclude(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		Logger.Errorf("[FILE] Reading file error : %s", err)
	}
	return string(data)
}
