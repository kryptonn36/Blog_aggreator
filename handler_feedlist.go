package main

import (
	"Blog_aggreator/internal/database"
	"context"
	"fmt"
)

func Feed_list(s *State, cmd Command, user database.User) error{
	if len(cmd.Args)!=0{
		return fmt.Errorf("This command does not take any argument")
	}
	feeds, err := s.db.GetFeed(context.Background())
	if err != nil{
		return fmt.Errorf("Error in Getting feed: %v",err)
	}
	for _,feed := range feeds{

		fmt.Printf("Name: %v\n",feed.Name)
		fmt.Printf("URL: %v\n", feed.Url)
		fmt.Printf("Created By: %v\n", user.Name)
	}
	return nil
}