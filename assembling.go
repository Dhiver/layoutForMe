package main

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func assembling(metadata Metadata) string {
	var textBuf bytes.Buffer
	for _, file := range metadata.IncludeOrder {
		data, err := ioutil.ReadFile(Config.GetString("textFolder") + "/" + file)
		if err != nil {
			Logger.Fatalf("[ASSEMBLING] Can't read a file : %s", err)
		}
		textBuf.WriteString(string(data) + "\n")
	}
	var metaBuf bytes.Buffer
	metaBuf.WriteString("---\n")
	buf, err := yaml.Marshal(metadata)
	if err != nil {
		Logger.Fatalf("[ASSEMBLING] Can't marshal metadata : %s", err)
	}
	metaBuf.WriteString(string(buf))
	metaBuf.WriteString("...\n")
	textBuf.Write(metaBuf.Bytes())
	return textBuf.String()
}
