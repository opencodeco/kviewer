package main

import (
	"fmt"
	"kviewer/pkg/kviewer"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Programa iniciado")
	setupConfig()
	if len(os.Args) < 3 {
		printHelp()
		return
	}

	action := os.Args[1]
	topic := os.Args[2]

	var maxMsgs int
	if len(os.Args) > 3 {
		maxMsgs, _ = strconv.Atoi(os.Args[3]) // Converta o terceiro argumento para um int.
	}

	switch action {
	case "consume":
		err := kviewer.Consume(topic, maxMsgs)
		if err != nil {
			fmt.Println("Erro ao consumir tópico:", err)
		}
	default:
		fmt.Println("Ação desconhecida")
	}
}

func printHelp() {
	fmt.Println("Uso:")
	fmt.Println("./kviewer <comando> [opções]")
	fmt.Println("")
	fmt.Println("Comandos disponíveis:")
	fmt.Println("consume <meu-topico> - Consumir mensagens de um tópico específico.")
}
