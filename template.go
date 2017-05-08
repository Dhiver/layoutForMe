package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	replacer *strings.Replacer
	funcMap  map[string]interface{}
)

func init() {
	replacer = strings.NewReplacer(
		"%", "\\%",
		"$", "\\$",
		"{", "\\{",
		"_", "\\_",
		"|", "\\textbar",
		">", "\\textgreater",
		"#", "\\#",
		"&", "\\&",
		"}", "\\}",
		"\\", "\\textbackslash",
		"<", "\\textless",
	)
	funcMap = template.FuncMap{
		"Date":          Date,
		"FileInclude":   FileInclude,
		"Join":          strings.Join,
		"LatexEscape":   LatexEscape,
		"ToString":      ToString,
		"ToStringSlice": ToStringSlice,
	}
}

func ToStringSlice(input []interface{}) []string {
	var out []string
	for _, v := range input {
		out = append(out, v.(string))
	}
	return out
}

func ToString(input interface{}) string {
	if input == nil {
		return ""
	}
	return input.(string)
}

func LatexEscape(text string) string {
	return replacer.Replace(text)
}

func templatingTemplates(templateName string, data interface{}) {
	files, err := ioutil.ReadDir(Config.GetString("templateFolder"))
	if err != nil {
		Logger.Fatalf("[TEMPLATE] Can't list files : %s", err)
	}
	var filenames []string
	for _, v := range files {
		if v.Mode().IsRegular() {
			Logger.Debugf("Adding template : %s", v.Name())
			filenames = append(filenames, Config.GetString("templateFolder")+"/"+v.Name())
		}
	}
	t, err := template.New("").Funcs(funcMap).ParseFiles(filenames...)
	if err != nil {
		Logger.Fatalf("[TEMPLATE] Can't create template : %s", err)
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		Logger.Fatalf("[TEMPLATE] Execute templating error : %s", err)
	}

	// Store
	path := Config.GetString("buildFolder") + "/"
	filename := strings.TrimSuffix(templateName, filepath.Ext(templateName))
	err = ioutil.WriteFile(path+filename, buf.Bytes(), 0644)
	if err != nil {
		Logger.Fatalf("[TEMPLATE] Can't write file : %s", err)
	}
}
