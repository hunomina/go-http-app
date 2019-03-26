package controller

import (
	"io"
	"io/ioutil"
	"net/http"
)

type HomeController struct {

}

func (c HomeController) Hello(w http.ResponseWriter, r *http.Request){
	var content []byte
	var err error
	content, err = ioutil.ReadFile("web/html/hello.html")
	if err == nil {
		_, _ = io.WriteString(w, string(content))
	} else {
		_, _ = io.WriteString(w, err.Error())
	}
}