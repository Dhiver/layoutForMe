package main

import (
	"time"
)

const (
	DEFAULT_DATE_LAYOUT string = "01/02/2006"
)

func Date(layout string) string {
	t := time.Now()
	return t.Format(layout)
}
