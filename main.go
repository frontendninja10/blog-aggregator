package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/frontendninja10/blog-aggregator/internal/config"
	"github.com/frontendninja10/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db *database.Queries
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	dbQueries := database.New(db)

	appState := &state{
		cfg: &cfg,
		db: dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(s *state, cmd command) error),
	}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", getUsers)

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

