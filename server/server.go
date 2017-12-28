package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/prestonp/boop/deploy"
)

// New creates a new service handler
func New(d deploy.Deployer) http.Handler {
	r := chi.NewRouter()
	tmpl := template.Must(template.ParseFiles("./templates/list.tmpl"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		deployments, err := d.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
		w.Header().Set("content-type", "text/html")
		tmpl.ExecuteTemplate(w, "list.tmpl", deployments)
	})

	r.Get("/logs/{idx}", func(w http.ResponseWriter, r *http.Request) {
		idxParam := chi.URLParam(r, "idx")
		idx, err := strconv.ParseInt(idxParam, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}

		dep, err := d.Get(int(idx))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
		dep.File.Seek(0, 0)
		_, err = io.Copy(w, dep.File)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
	})

	r.Get("/deploy", func(w http.ResponseWriter, r *http.Request) {
		err := d.Deploy()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}

		w.Header().Set("location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprintln(w, "ok")
	})

	return r
}
