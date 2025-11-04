package render

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

type Renderer struct {
	root  string
	cache map[string]*template.Template
}

// New parses all templates under root once at startup
func New(root string) (*Renderer, error) {
	r := &Renderer{root: root, cache: make(map[string]*template.Template)}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if filepath.Ext(path) != ".html" || filepath.Base(path) == "layout.html" {
			return nil
		}

		files := collectLayouts(path)
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return err
		}
		key, _ := filepath.Rel(root, path)
		key = filepath.ToSlash(key)
		r.cache[key] = tmpl
		return nil
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data any) {
	tmpl, ok := r.cache[name]
	if !ok {
		http.Error(w, "template not found: "+name, http.StatusNotFound)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func collectLayouts(page string) []string {
	var layouts []string
	dir := filepath.Dir(page)
	for {
		layoutPath := filepath.Join(dir, "layout.html")
		if fileExists(layoutPath) {
			layouts = append([]string{layoutPath}, layouts...)
		}
		parent := filepath.Dir(dir)
		if parent == dir || !strings.Contains(parent, "templates") {
			break
		}
		dir = parent
	}
	layouts = append(layouts, page)
	return layouts
}

func fileExists(path string) bool {
	matches, _ := filepath.Glob(path)
	return len(matches) > 0
}
