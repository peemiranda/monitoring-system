package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"io/ioutil"
)

const monitoring = 3
const delay = 1

func main() {

	showIntro()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do sistema...")
			os.Exit(0)
		default:
			fmt.Println("This command is not allowed")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Pedro"
	version := 2.4
	fmt.Println("hello mr.", name)
	fmt.Println("we are on version:", version)
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("the comand used is", commandRead)
	fmt.Println(" ")

	return commandRead
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	// sites := []string{
	// 	"https://autovist.com.br/dash/home/",
	// 	"https://app.tabelasomos.com.br",
	// 	"https://app.checkview.com.br",
	// 	"https://app.safysign.com.br",
	// 	"https://sqad.com.br",
	// }

	sites = readSitesOnArchive()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("")
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Minute)
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func readSitesOnArchive()[]string{
	var sites [] string

	// archive, err := os.Open("sites.txt")
	archive, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	fmt.Println(string(archive))
	return sites
}