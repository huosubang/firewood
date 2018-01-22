package firewood

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/mholt/binding"
)


// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Offset int `url:"offset,omitempty"`
	// For paginated result sets, the number of results to include per page.
	Limit int `url:"limit,omitempty"`
}

type T interface {
	FieldMap(req *http.Request) binding.FieldMap
	Do(res http.ResponseWriter, req *http.Request) (interface{}, int, error)
}

func New(res http.ResponseWriter, req *http.Request, r T, methods ...string) {
	methodValid := false
	for _, v := range methods {
		if v == req.Method {
			methodValid = true
			break
		}
	}
	if !methodValid {
		Write(res, req, http.StatusMethodNotAllowed, nil)
		return
	}

	req.ParseForm()
	if err := binding.Bind(req, r); err != nil {
		logrus.Warnf("[%s] err:%s", req.RequestURI, err.Error())
		Write(res, req, http.StatusBadRequest, nil)
		return
	}

	d, code, err := r.Do(res, req)
	if err != nil {
		logrus.Warnf("[%s] err:%s", req.RequestURI, err.Error())
	}
	Write(res, req, code, d)

	return
}

func Write(res http.ResponseWriter, req *http.Request, code int, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Service-Code", fmt.Sprintf("%d", code))

	resp, err := json.Marshal(data)
	if err != nil {
		res.WriteHeader(500)
	} else {
		res.WriteHeader(code)
	}

	logrus.Debugf("[%s] resp:%s", req.RequestURI, string(resp))

	res.Write(resp)
}
