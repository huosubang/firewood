// Copyright 2018 The QOS Authors

package firewood

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HttpTest struct {
	t *testing.T
	h http.Handler

	req *http.Request
}

func NewHttpTest(t *testing.T, req *http.Request) *HttpTest {
	var ht HttpTest
	ht.t = t
	ht.req = req

	return &ht
}

func (ht *HttpTest) Do(r http.Handler, v interface{}) (*httptest.ResponseRecorder, error) {
	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, ht.req)

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, rw.Body)
	} else {
		err := json.NewDecoder(rw.Body).Decode(v)
		if err != nil {
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			} else {
				return rw, err
			}
		}
	}

	return rw, nil
}
