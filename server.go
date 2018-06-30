package firewood

import (
	"net/http"

	"github.com/urfave/negroni"
)

type Server struct {
	n *negroni.Negroni
}

func RunServer(addr string, handler http.Handler) error {
	n := negroni.New()
	n.Use(NewLoggerHandler())
	n.UseHandler(handler)
	err := http.ListenAndServe(addr, n)
	return err
}
