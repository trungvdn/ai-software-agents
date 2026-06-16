package main

import (
	"fmt"
	"log"

	"ai-software-agents/agents/bugfix"
	"ai-software-agents/shared/tools"
)

func main() {
	agent := bugfix.New(
		tools.NewReadFileTool(),
		tools.NewSearchCodeTool("."),
	)

	resp, err := agent.Run(
		"GetUser",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
