package views

import (
	"html/template"
	"net/http"

	"github.com/indaco/predix-ec-configurator/helpers"
)

func NewMarkdownView(layout string, files ...string) *MarkdownView {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &MarkdownView{
		Template: t,
		Layout:   layout,
		//Body:   helpers.ParseMarkdown(markdownFile),
	}
}

type MarkdownView struct {
	Template *template.Template
	Layout   string
	Body     template.HTML
}

func (v *MarkdownView) AddMarkdownContent(markdownFile string) {
	v.Body = helpers.ParseMarkdown(markdownFile)
}

func (v *MarkdownView) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *MarkdownView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, v); err != nil {
		panic(err)
	}
}
