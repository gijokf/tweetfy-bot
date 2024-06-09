package services

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func TweetFile() (string, error) {
	file, err := os.Open("resources/tweets.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: ", err)
	}
	defer file.Close()

	lineNumber, err := countLines(file)
	if err != nil {
		return "", err
	}

	// Gerar um número aleatório entre 0 e numLinhas
	rand.NewSource(time.Now().UnixNano())
	randomLine := rand.Intn(lineNumber) + 1 // +1 para garantir que o número aleatório seja entre 1 e numLinhas

	// Ler o arquivo novamente e encontrar a linha escolhida aleatoriamente
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	actualLine := 0

	for scanner.Scan() {
		actualLine++
		if actualLine == randomLine {
			return scanner.Text(), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("linha aleatória não encontrada")
}

func countLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}
