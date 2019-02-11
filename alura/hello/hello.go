package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeNomes()
	exibeIntroducao()

	for {
		exibeMenu()
		registraLog("site-false", false)
		comando := lerMenu()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("opção escolhida 2")
			imprimeLogs()
		case 3:
			fmt.Println("opção escolhida 3")
			limparArquivo()
		case 0:
			fmt.Println("opção escolhida 0")
			os.Exit(0)
		default:
			fmt.Println("Comando selecionado inválido")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	fmt.Println("Grupo de estudo GoLang")

	nome := "Douglas"
	idade := 24.1

	fmt.Println("olá ", nome, "sua idade é ", idade)

	fmt.Println("olá ", reflect.TypeOf(nome), "sua idade é ", reflect.TypeOf(idade))

}

func lerMenu() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi ", comando)

	return comando
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Limpar arquivo")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://random-status-code.herokuapp.com", "https://www.goblockchain.io", "https://www.caelum.com.br"}

	sites := lerSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {

			fmt.Println("Testando site ", i, ":", site)
			testaSites(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSites(site string) {
	resp, error := http.Get(site)

	if ocorreuErro(error) {
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " Foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, " Não esta rodando o site!")
		registraLog(site, false)
	}
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Henrique"}
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println(len(nomes))
	nomes = append(nomes, "novo valor")
	fmt.Println(cap(nomes))
}

func lerSitesArquivo() []string {
	var sites []string

	arquivo, error := os.Open("sites.txt")

	// arquivo, err := ioutil.ReadFile("sites.txt")
	// fmt.Println(string(arquivo))

	if ocorreuErro(error) {
		panic(fmt.Sprintln("O c"))
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err == io.EOF {
			break
		}
		sites = append(sites, linha)
	}
	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {
	arquivo, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if ocorreuErro(error) {
		return
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(arquivo)
	defer arquivo.Close()
}

func imprimeLogs() {
	arquivo, error := ioutil.ReadFile("log.txt")
	if ocorreuErro(error) {
		return
	}

	fmt.Println("Contéudo arquivo ", string(arquivo))
}

func limparArquivo() {

	os.Remove("log.txt")

	arquivo, error := os.OpenFile("log.txt", os.O_RDWR, 0666)
	if ocorreuErro(error) {
		return
	}
	arquivo.WriteString("")

}

func ocorreuErro(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
