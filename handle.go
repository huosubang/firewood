package firewood

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mholt/binding"
	"github.com/sirupsen/logrus"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Offset int64 `url:"offset,omitempty"`
	// For paginated result sets, the number of results to include per page.
	Limit int64 `url:"limit,omitempty"`
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
	logrus.Debugf("[%s] r:%+v", req.RequestURI, r)

	d, code, err := r.Do(res, req)
	logrus.Debugf("[%s] resp:%v, code:%d, err:%v", req.RequestURI, d, code, err)

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
