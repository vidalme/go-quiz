package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var QUIZ_FILE string = "quiz.csv"

type match struct {
	questions map[string]int
	rodadas   int
	corretas  int
}

func readCSV(f string) *csv.Reader {
	raw, err := os.OpenFile(f, os.O_RDWR, 0700)
	if err != nil {
		fmt.Printf("Erro ao carregar o arquivo CSV: %v\n", err)
	}
	return csv.NewReader(raw)
}

func loadQuestions(r *csv.Reader) match {

	m := match{
		questions: make(map[string]int),
		rodadas:   0,
		corretas:  0,
	}

	ra, err := r.ReadAll()

	if err != nil {
		fmt.Printf("Erro na leitura dos dados do CSV: %v\n", err)
	}

	for _, j := range ra {
		a, err := strconv.Atoi(strings.TrimSpace(j[1]))
		if err != nil {
			fmt.Printf("Erro na conversao de strings para Ints nas respostas do quiz: %v\n", err)
		}
		m.questions[j[0]] = a
	}

	return m

}

func welcome() {
	fmt.Println("+++++++++++++++++++++++++++++")
	fmt.Println("++++  Bem vindo ao quiz  ++++")
	fmt.Println("+++++++++++++++++++++++++++++")

}

func fazPergunta(p string, r int) bool {
	var a string
	fmt.Println("Quanto é", p, "?")
	fmt.Scan(&a)

	na, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("erro")
	}
	if na == r {
		return true
	}
	return false

}

func matchLoop(m *match) {
	// loop := true

	for k, v := range m.questions {
		if p := fazPergunta(k, v); p {
			m.corretas++
		}
	}
	fmt.Println(m)
}

func startGame(m match) {
	matchLoop(&m)
}

func createGame() {

	welcome()
	startGame(loadQuestions(readCSV(QUIZ_FILE)))
}

func main() {
	createGame()
}
