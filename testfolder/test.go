package testfolder

import "net/http"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It works!"))
}
