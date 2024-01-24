package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	id    string
	name  string
	saldo float64 `json:saldo_conta`
}

func main() {
	conta := Conta{
		id:    "fsldkj",
		saldo: 32.1,
		name:  "Paolo",
	}
	res, err := json.Marshal(conta)
	fmt.Printf("res=%v err=%v\n", res, err)

	err = json.NewEncoder(os.Stdout).Encode(conta)
	fmt.Printf("err=%v\n", err)

	jsonPuro := []byte(`{"id":"21", "name":"Leticia", "saldo_conta":32.2}`)
	var contaLeticia Conta
	err = json.Unmarshal(jsonPuro, &contaLeticia)
	fmt.Printf("err=%v\n", err)
}
