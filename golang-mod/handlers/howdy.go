package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Howdy(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "howdy~\n%d\n", time.Now().Unix())
}
