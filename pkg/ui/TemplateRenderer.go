package ui

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/adampresley/logrusviewer/cmd/logrusviewer/www"
	"github.com/labstack/echo"
)

var templates map[string]*template.Template
var pageList = []string{
	"viewer.gohtml",
	"selectLogFile.gohtml",
}

/*
TemplateRenderer describes a handlers for rendering layouts/pages
*/
type TemplateRenderer struct {
	templates *template.Template
}

/*
NewTemplateRenderer creates a new struct
*/
func NewTemplateRenderer(debugMode bool) *TemplateRenderer {
	result := &TemplateRenderer{}
	result.LoadTemplates(debugMode)

	return result
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	var tmpl *template.Template
	var ok bool

	if tmpl, ok = templates[name]; !ok {
		return fmt.Errorf("Cannot find template %s", name)
	}

	return tmpl.ExecuteTemplate(w, "layout", data)
}

func (t *TemplateRenderer) LoadTemplates(debugMode bool) {
	templates = make(map[string]*template.Template)

	for _, fileName := range pageList {
		trimmedName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		templates["mainLayout:"+trimmedName], _ = template.Must(
			template.New("layout").Parse(www.FSMustString(debugMode, "/www/logrusviewer/layouts/mainLayout.gohtml")),
		).Parse(www.FSMustString(debugMode, "/www/logrusviewer/pages/"+fileName))
	}
}
