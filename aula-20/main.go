package main

func main() {
	//comandos
	// go build -> gera o binario (para o computador que esta rodando)
	// go COOS=windows buiold -> gera o binario para o windows
	// go COOS=darwin GOARCH=arm64 -> gera o binario para mac com arquitetura dos apple silicons rodar
	// go env GOOS GOARCH -> Lista qual a config de OS e ARCH esta rodando as variaveis de ambiente

	//Quando se tem go.mod no momento do build ele vai usar o nome do modulo para nomear o binario
	//Caso queira trocar tem que passar uma flag (-o, por causa de output) ex>
	// go build -o xpto
}
