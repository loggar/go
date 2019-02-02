package basics

import "net/http"

// App ...
type App struct {
	Message string
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(a.Message))
}
