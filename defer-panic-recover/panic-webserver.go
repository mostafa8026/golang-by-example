package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		// it is on to you to decide to see whether the err.Error() is a problem or not!
		panic(err.Error())
	}
}
