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
	//recovery := negroni.NewRecovery()
	//recovery.PanicHandlerFunc = func(d *negroni.PanicInformation){fmt.Println("----paic")}
	//n.Use(recovery)
	n.UseHandler(handler)
	err := http.ListenAndServe(addr, n)
	return err
}
