// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Server!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3000", nil)
}
