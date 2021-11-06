package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var path = filepath.FromSlash("")


func reg(w http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Fprint(w, "Activation link has been sent")

	log.Printf(
		"first-name: %s\t last-name: %s\t email: %s\t password: %s\n",
		req.Form.Get("first-name"),
		req.Form.Get("last-name"),
		req.Form.Get("email"),
		req.Form.Get("password"),
		)

}

func main() {

	fs := http.FileServer(http.Dir(path))

	http.Handle("/", fs)
	http.HandleFunc("/reg", reg)

	log.Println("Server is starting on port 9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
