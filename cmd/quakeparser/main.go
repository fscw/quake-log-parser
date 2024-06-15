package main

import (
	"fmt"
	"os"
	"path/filepath"
	"quake_log_parser/internal/app/parser"
	"quake_log_parser/internal/reports"
)

func main() {
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Erro ao obter o diretório atual: %v\n", err)
		return
	}

	logFilePath := filepath.Join(currentDir, "../../logs/qgames.log")

	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		fmt.Printf("Arquivo de log não encontrado no caminho: %s\n", logFilePath)
		return
	}

	matches, err := parser.ParseLogFile(logFilePath)
	if err != nil {
		fmt.Printf("Erro ao analisar o arquivo de log: %v\n", err)
		return
	}

	for _, match := range matches {
		fmt.Println(reports.GenerateReport(match))
	}
}
