package internal

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func StartServer() {

	http.HandleFunc("/", sayHello)
	fmt.Println(http.ListenAndServe(":8080", nil))

}