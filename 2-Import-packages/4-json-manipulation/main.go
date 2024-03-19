package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo int
}

func main()  {
	conta := Conta{Numero: 1, Saldo: 100}
	// salvar em uma variavel
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero":1,"saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}

	println(contaX.Numero)
	println(contaX.Saldo)
}
