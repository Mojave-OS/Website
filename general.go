package main

import (
	"net/http"
	"os"
)

const (
	IndexFile = "./views/general/index.html"
)

func (gs *GeneralServer) ServeIndex(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile(IndexFile);
	if err != nil {
		panic(err)
	}

	_, err = w.Write(dat)
	if err != nil {
		panic(err)
	}
}