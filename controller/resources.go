package controller

import (
	"io"
	"io/ioutil"
	"net/http"
)

type ResourcesController struct {

}

func (c ResourcesController) GetCssFile(w http.ResponseWriter, r *http.Request){
	var content []byte
	var err error
	content, err = ioutil.ReadFile("web" + r.RequestURI)
	if err == nil {
		w.Header().Add("Content-Type", "text/css")
		_, _ = io.WriteString(w, string(content))
	}
}