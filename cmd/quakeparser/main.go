package main

import (
	"fmt"
	"quake_log_parser/internal/app/parser"
	"quake_log_parser/internal/reports"
)

func main() {
	logFilePath := "logs/qgames.log"
	matches, err := parser.ParseLogFile(logFilePath)

	if err != nil {
		fmt.Printf("Erro ao analisar o arquivo de log: %v\n", err)
		return
	}

	for _, match := range matches {
		fmt.Println(reports.GenerateReport(match))
	}
}
