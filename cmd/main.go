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
	case "produce":
		if len(os.Args) < 4 {
			fmt.Println("Por favor, especifique a mensagem a ser enviada.")
			return
		}
		message := os.Args[3]
		err := kviewer.Produce(topic, message)
		if err != nil {
			fmt.Println("Erro ao produzir mensagem:", err)
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
	fmt.Println("produce <meu-topico> <mensagem>- Consumir mensagens de um tópico específico.")
}
