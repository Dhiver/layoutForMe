package main

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"gopkg.in/yaml.v2"
)

type Accountable struct {
	Name string `yaml:"name"`
	Mail string `yaml:"mail"`
}

type Version struct {
	Number  string   `yaml:"number"`
	Date    string   `yaml:"date"`
	Author  string   `yaml:"author"`
	Section []string `yaml:"section"`
	Log     string   `yaml:"log"`
}

type Output struct {
	Name      string `yaml:"name"`
	Extention string `yaml:"extention"`
	Template  string `yaml:"template"`
}

type Metadata struct {
	Institute      string        `yaml:"institute"`
	Title          string        `yaml:"title"`
	Author         []string      `yaml:"author"`
	Subject        string        `yaml:"subject"`
	Abstract       string        `yaml:"abstract"`
	Keywords       []string      `yaml:"keywords"`
	Lang           string        `yaml:"lang"`
	DateLayout     string        `yaml:"dateLayout"`
	Date           string        `yaml:"date"`
	Header         string        `yaml:"header"`
	Footer         string        `yaml:"footer"`
	Accountables   []Accountable `yaml:"accountable",flow`
	Versions       []Version     `yaml:"version",flow`
	IncludeOrder   []string      `yaml:"includeOrder"`
	Outputs        []Output      `yaml:"output",flow`
	HighlightStyle string        `yaml:"highlightStyle"`
	ExtraData      interface{}   `yaml:"extraData"`

	Font           string   `yaml:"font"`
	FontSize       string   `yaml:"fontSize"`
	DocumentClass  string   `yaml:"documentClass"`
	ClassOption    []string `yaml:"classOption"`
	Toc            bool     `yaml:"toc"`
	TocDepth       int      `yaml:"tocDepth"`
	NumberSections bool     `yaml:"numberSections"`
	PaperSize      string   `yaml:"paperSize"`
	LinkColor      string   `yaml:"linkColor"`
	LineStretch    float32  `yaml:"lineStretch"`
	Thanks         string   `yaml:"thanks"`
}

func (m *Metadata) Read() {
	data, err := ioutil.ReadFile(Config.GetString("metaFile"))
	if err != nil {
		Logger.Fatalf("[META] Could not read document metadata : %s", err)
	}
	err = yaml.Unmarshal(data, m)
	if err != nil {
		Logger.Fatalf("[META] Could not parse document metadata : %s", err)
	}
	if m.Lang == "" {
		m.Lang = "en-US"
	}
}

func (m *Metadata) Interpret() {
	if m.DateLayout == "" {
		m.DateLayout = DEFAULT_DATE_LAYOUT
	}
	if m.Date == "" {
		m.Date = Date(m.DateLayout)
	}
	if m.Abstract == "" {
		m.Abstract = FileInclude("README.md")
	}
	if len(m.Versions) == 0 {
		m.Versions = Versions(m.DateLayout)
	}
	en := display.English.Languages()
	m.Lang = strings.ToLower(en.Name(language.MustParse(m.Lang)))
}

func (m *Metadata) Sanitize() {
	m.Institute = LatexEscape(m.Institute)
	m.Title = LatexEscape(m.Title)
	for k, v := range m.Author {
		m.Author[k] = LatexEscape(v)
	}
	m.Subject = LatexEscape(m.Subject)
	m.Abstract = LatexEscape(m.Abstract)
	for k, v := range m.Keywords {
		m.Keywords[k] = LatexEscape(v)
	}
	m.Date = LatexEscape(m.Date)
	m.Header = LatexEscape(m.Header)
	m.Footer = LatexEscape(m.Footer)
	for k, v := range m.Accountables {
		m.Accountables[k].Name = LatexEscape(v.Name)
		m.Accountables[k].Mail = LatexEscape(v.Mail)
	}
	for k, v := range m.Versions {
		m.Versions[k].Number = LatexEscape(v.Number)
		m.Versions[k].Date = LatexEscape(v.Date)
		m.Versions[k].Author = LatexEscape(v.Author)

		for i, j := range m.Versions[k].Section {
			m.Versions[k].Section[i] = LatexEscape(j)
		}
		m.Versions[k].Log = LatexEscape(v.Log)

	}
}
