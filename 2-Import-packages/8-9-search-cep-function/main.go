package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}


func searchCEP(cep string) (*ViaCEP, error)  {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() 

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var address ViaCEP
	err = json.Unmarshal(body, &address)

	return &address, nil
}

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
	cep, err := searchCEP(cepParam)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(cep)
}

func main()  {
	http.HandleFunc("/", handleSearchCEP)
	http.ListenAndServe(":8080", nil)
}
