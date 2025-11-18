package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeMenu()
	comando := lerComando()

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

	sliceSites := criarSitesSlice()
	fmt.Println(sliceSites)

	for i := range sliceSites {
		monitorarSite(sliceSites[i])
	}
}

func monitorarSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ao acessar o site:", site)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
		} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		}
	}
}

func criarSitesSlice() []string {
	sites := []string{"https://www.alura.com.br", "https://www.google.com.br", "https://www.github.com", "https://www.stackoverflow.com"}
	sites = append(sites, "https://www.udemy.com")

	return sites
}
