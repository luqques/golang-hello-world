package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeMenu()
	comando := lerComando()

	nome, idade := devolveNomeEIdade()
	fmt.Println("Ola", nome, "sua idade é", idade)

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido...")
		os.Exit(-1)
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Lucas"
	idade := 21
	return nome, idade
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Print("Escolha uma opção: ")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	var sites [5]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://www.caelum.com.br"
	sites[2] = "https://www.google.com.br"
	sites[3] = "https://www.github.com"
	sites[4] = "https://www.stackoverflow.com"

	fmt.Println(sites)

	// site com URL inexistente
	site := "https://httpbin.org/status/404" // ou 200
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
