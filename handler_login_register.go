package main

import (
	"Blog_aggreator/internal/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

// this is the function signature of all command handlers
func HandlerLogin(s *State, cmd Command) error{
	if len(cmd.Args)==0{
		return errors.New("the login handler expects a single argument,the username")
	}
	s.Config.Current_user_name = cmd.Args[0]
	_, err := s.db.GetUserByName(context.Background(), cmd.Args[0])
	if err!=nil{
		fmt.Printf("The user %v does not exist\n",cmd.Args[0])
		os.Exit(1)
	}

	if err := s.Config.SetUser(cmd.Args[0]); err!= nil{
		return err
	}
	fmt.Printf("The user switched successfully")
	return nil
}

func handlerRegister(s *State, cmd Command) error{
	if len(cmd.Args)<1{
		return  errors.New("Name field is required")
	}

	username := cmd.Args[0]
	_, err := s.db.GetUserByName(context.Background(), username)
	if err==nil{
		os.Exit(1)
	}
	if err!=sql.ErrNoRows{
		return err
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	})
	if err!=nil{
		return err
	}

	if err := s.Config.SetUser(cmd.Args[0]); err!=nil{
		return err
	}

	fmt.Printf("The user %v was created\n",cmd.Args[0])
	log.Printf("Created User: %+v\n",user)
	return nil
}


func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
