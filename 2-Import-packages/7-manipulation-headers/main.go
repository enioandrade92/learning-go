package main

import "net/http"


func handleSearchCEP(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := req.URL.Query().Get("cep")

	if cepParam == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Hello World!"))
}

func main()  {
	http.HandleFunc("/", handleSearchCEP)
	http.ListenAndServe(":8080", nil)
}
