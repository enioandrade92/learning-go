package main

import "net/http"

type blog struct {
	title string
}

func(b blog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(b.title))
}

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/blog", blog{title: "my"})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	http.ListenAndServe(":8080", mux)
}
