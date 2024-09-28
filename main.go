package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	QUIZ_FILE    string = "quiz.csv"
	DURA_DEFAULT int    = 30
	MAX_DURATION int    = 300
)

type partida struct {
	questions map[string]string
	duracao   int
	corretas  int
}

func fazerPergunta(p string, r string) bool {
	var a string
	fmt.Printf("%s = ", p)
	fmt.Scan(&a)
	return strings.TrimSpace(a) == r
}

func loopPerguntas(m *partida) {

	countdown_ticker := time.NewTicker(time.Second)
	defer countdown_ticker.Stop()

	done := make(chan bool)

	go func() {
		done <- true
	}()
}

func finalizaJogo(m partida) {
	fmt.Print("\nParábens, fim do jogo!!\n")
	fmt.Printf("Você acertou %d de um total de %d!\n\n", m.corretas, len(m.questions))
	os.Exit(0)
}

func startGame(m partida) {

	fmt.Printf("Inicio da partida, você tem %d segundos para finalizar as perguntas.\n\n", m.duracao)

	timer := time.NewTimer(time.Second * time.Duration(m.duracao))

	go func() {
		for k, v := range m.questions {
			if p := fazerPergunta(k, v); p {
				m.corretas++
			}
		}
		finalizaJogo(m)
	}()

	<-timer.C
	finalizaJogo(m)
}

func criaPartida(pr [][]string) partida {
	m := partida{
		questions: make(map[string]string),
		corretas:  0,
	}

	flag.IntVar(&m.duracao, "d", DURA_DEFAULT, "Use d para alterar a duração default")
	flag.Parse()

	for _, j := range pr {
		a := strings.TrimSpace(j[1])
		m.questions[j[0]] = a
	}
	return m
}

func carregaPerguntas(qf string) [][]string {

	f, err := os.OpenFile(qf, os.O_RDWR, 0700)
	if err != nil {
		fmt.Printf("Erro ao carregar o arquivo CSV: %v\n", err)
	}

	reader := csv.NewReader(f)
	ra, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Erro na leitura dos dados do CSV: %v\n", err)
	}
	return ra
}

func welcomeMenu() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("++++              Bem vindo ao quiz             ++++")
	fmt.Println("++++              ENTER para iniciar            ++++")
	fmt.Println("++++              CTRL + C para sair            ++++")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")

	in := bufio.NewReader(os.Stdin)
	_, err := in.ReadBytes('\n')
	if err != nil {
		fmt.Println("erro:", err)
	}
}

func createGame(qf string) {
	welcomeMenu()

	perguntas := carregaPerguntas(qf)
	match := criaPartida(perguntas)

	startGame(match)
}

func main() {
	createGame(QUIZ_FILE)
}
