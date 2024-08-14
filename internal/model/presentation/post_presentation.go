package presentation

import (
	"html/template"
)

type Post struct {
	Id         int
	GUID       string
	User       string
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	Comments   []Comment
	Session    Session
}

func (p Post) TruncatedText() string {
	chars := 0
	for i := range p.RawContent {
		chars++
		if chars > 150 {
			return p.RawContent[:i] + `...`
		}
	}
	return p.RawContent
}