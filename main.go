package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frontendninja10/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}

	appState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(s *state, cmd command) error),
	}
	cmds.register("login", loginHandler)

	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		log.Fatal("Usage: cli <command> [args...]")
	}
	fmt.Println(os.Args)


	var cmd command

	cmd.name = os.Args[1]
	cmd.args = os.Args[2:]

	if err := cmds.run(appState, cmd); err != nil {
		log.Fatal(err)
	}
}

