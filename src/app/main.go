package main

import (
	"fmt"
	"net/http"
	"github.com/seekasia/common"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You visited 1 %s in `%s`!", r.URL.Path, common.AppName)
}

func main() {
	common.StartServer("8080", handle)
}
