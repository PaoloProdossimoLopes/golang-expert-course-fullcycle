package main

type Pessoa interface {
	digaNome()
}

type Client struct {
	name string
}

// SAME AS CONSTRUCTOR IN OTHER LENGUAGES
func newCliente() *Client {
	return &Client{}
}

func (client Client) digaNome() {
	println("meu nome é:", client.name)
}

func pessoaDizNome(pessoa Pessoa) {
	pessoa.digaNome()
}

func main() {
	pessoaDizNome(Client{
		"Paolo",
	})
}

type People struct {
	name string
}

//// Copy value
//func (p People) fale() {
//
//}
//
////reference value
//func (p *People) fale() {
//
//}
