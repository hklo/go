package main

import (
	"fmt"
	"github.com/seekasia/common"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You visited %s in `%s`!", r.URL.Path, common.AppName)
}

func main() {
	common.StartServer("8080", handle)
}
