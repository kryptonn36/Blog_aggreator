package main

import (
	"context"
	"Blog_aggreator/internal/config"
	"Blog_aggreator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
	"database/sql"
	"github.com/joho/godotenv"
)

// this will tell the handlers about the status of application state
type State struct{
	db *database.Queries
	Config *config.Config
}

func main() {
	godotenv.Load()
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.Db_url)
	if err!= nil{
		log.Fatalf("Error connecting to database: %v",err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &State{
		db: dbQueries,
		Config: &cfg,
	}
	c := Commands{
		registeredcommands:make(map[string]func (*State,Command) error),
	}

	// Commands that are executed
	c.Register("login", HandlerLogin)
	c.Register("register", handlerRegister)
	c.Register("reset",HandlerReset)
	c.Register("users", users)
	c.Register("agg", handlerAgg)
	c.Register("addfeed",middlewareLoggedIn(handlerAddfeed))
	c.Register("feeds", middlewareLoggedIn(Feed_list))
	c.Register("follow", middlewareLoggedIn(HandlerFollow))
	c.Register("following", middlewareLoggedIn(HandlerFollowing))
	c.Register("unfollow", middlewareLoggedIn(HandlerUnfollow))
	c.Register("browse", middlewareLoggedIn(handlerBrowse)) 

	if len(os.Args) < 2{
		log.Fatal("no command provided")
	}
	arguments := os.Args
	name := arguments[1]
	slice := arguments[2:]

	cmd := Command{
		Name:name,
		Args:slice,
	}
	
	if err := c.Run(s, cmd); err!=nil{
		log.Fatal(err)
	}

}


func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.db.GetUserByName(context.Background(), s.Config.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
