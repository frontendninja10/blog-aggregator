package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/app/auth"
	"github.com/frontendninja10/blog-aggregator/internal/app/feeds"
	"github.com/frontendninja10/blog-aggregator/internal/app/users"
	"github.com/frontendninja10/blog-aggregator/internal/config"
	"github.com/frontendninja10/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)


func middlewareLoggedIn(handler func(s *app.State, cmd app.Command, user database.User) error) func(*app.State, app.Command) error {
	return func(s *app.State, cmd app.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUsername)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}
		return handler(s, cmd, user)
	}
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	dbQueries := database.New(db)

	appState := app.NewState(&cfg, dbQueries)

	cmds := app.NewCommands()
	cmds.Register("login", auth.Login)
	cmds.Register("register", users.Register)
	cmds.Register("reset", users.Reset)
	cmds.Register("users", users.ListUsers)
	cmds.Register("aggregate", middlewareLoggedIn(feeds.AggregateFeeds))
	cmds.Register("feeds", feeds.ListFeeds)
	cmds.Register("addfeed", middlewareLoggedIn(feeds.AddFeed))
	cmds.Register("follow", middlewareLoggedIn(feeds.Follow))
	cmds.Register("unfollow", middlewareLoggedIn(feeds.Unfollow))
	cmds.Register("following", middlewareLoggedIn(feeds.Following))

	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		log.Fatal("Usage: gator <command> [args...]")
	}
	fmt.Println(os.Args)


	cmd := app.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(appState, cmd); err != nil {
		log.Fatal(err)
	}
}

