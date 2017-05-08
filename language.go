package main

import (
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/i18n"
)

var (
	T i18n.TranslateFunc
)

func LoadLanguage(lang string) {
	var err error
	folder := Config.GetString("templateFolder") + "/" + Config.GetString("translateFolder")
	translFile := fmt.Sprintf("%s/%s.flat.yaml", folder, lang)
	if _, err = os.Stat(translFile); os.IsNotExist(err) {
		Logger.Fatalf("[LANG] File '%s' does not exists : %s", translFile, err)
	}
	i18n.MustLoadTranslationFile(translFile)
	T, err = i18n.Tfunc(lang)
	if err != nil {
		Logger.Fatalf("[LANG] Can't load language %s : %s", lang, err)
	}
	funcMap["T"] = T
}
