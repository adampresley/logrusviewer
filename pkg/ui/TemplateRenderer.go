package ui

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mailslurper/mailslurper/cmd/mailslurper/www"

	"github.com/labstack/echo"

	"github.com/logrusviewer/cmd/logrusviewer/www"
)

var templates map[string]*template.Template

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
	var files []os.FileInfo
	var err error

	templates = make(map[string]*template.Template)

	if files, err = ioutil.ReadDir("/www/logrusviewer/pages"); err != nil {
		panic("Error reading pages directory")
	}

	for _, file := range files {
		if !file.IsDir() {
			basename := file.Name()
			trimmedName := strings.TrimSuffix(basename, filepath.Ext(basename))

			templates["mainLayout:"+trimmedName], _ = template.Must(
				template.New("layout").Parse(www.FSMustString(debugMode, "/www/logrusviewer/layouts/mainLayout.gohtml")),
			).Parse(www.FSMustString(debugMode, "/www/logrusviewer/pages/"+basename))
		}
	}
}
