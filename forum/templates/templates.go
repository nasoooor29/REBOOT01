package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed *
var templates embed.FS

func getTemplatePaths() []string {
	files := []string{}
	fs.WalkDir(templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(path, "pages/") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files

}

func LoadPage(w http.ResponseWriter, templatePath string, data any) error {
	l := getTemplatePaths()
	l = append(l, templatePath)
	t, err := template.ParseFS(templates, l...)
	if err != nil {
		panic(err)
	}
	return t.ExecuteTemplate(w, filepath.Base(templatePath), data)
}
