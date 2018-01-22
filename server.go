package firewood

import (
	"net/http"

	"github.com/tabalt/gracehttp"
	"github.com/urfave/negroni"
)

type Server struct {
	n    *negroni.Negroni
}


func RunServer(addr string, handler http.Handler) error {
	n := negroni.New()
	n.Use(NewLogger())
	n.UseHandler(handler)
	err := gracehttp.ListenAndServe(addr, n)

	return err
}
