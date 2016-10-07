package handler

import (
	"fmt"
	"net/http"
	// "github.com/toolkits/str"
	// "github.com/toolkits/web"
	// "http/render"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test dddddddd")
}
