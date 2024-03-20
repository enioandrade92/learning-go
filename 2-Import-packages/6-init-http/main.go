package main

import "net/http"


func searchCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main()  {
	http.HandleFunc("/", searchCEP)
	http.ListenAndServe(":8080", nil)
}
