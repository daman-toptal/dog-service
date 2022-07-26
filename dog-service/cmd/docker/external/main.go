package main

import (
	"bufio"
	"dog-service/internal/command"
	"dog-service/util/config"
	"dog-service/util/logging"
	"fmt"
	"os"
	"strings"
)

func main() {
	config.SetupConfig()

	logging.SetupLogging(config.GetString("log.level"))

	reader := bufio.NewReader(os.Stdin)
	//read user input and execute corresponding command
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		text = strings.Replace(text, "\n", "", -1)

		srv, err := command.NewCommand(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		srv.Execute()
	}
}
