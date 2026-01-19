package main

import (
	"Blog_aggreator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddfeed(s *State, cmd Command, user database.User) error{
	if len(cmd.Args)!=2{
		return fmt.Errorf("This command requires two input name and url of the feed")
	}
	username := cmd.Args[0]
	feed ,err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID        : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		Name      : username,
		Url       : cmd.Args[1] ,
		UserID    : user.ID ,
	})
	if err!= nil{
		return fmt.Errorf("Error in creating feed: %v",err)
	}

	follower, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID    		: uuid.New(),
		CreatedAt	: time.Now(),
		UpdatedAt	: time.Now(),
		UserID		: user.ID,
		FeedID		: feed.ID,
	})
	if err!=nil{
		return fmt.Errorf("Error in creating feed: %v",err)
	}
	// fmt.Print(feed)
	fmt.Printf("%v",follower.FeedName)
	return nil
}