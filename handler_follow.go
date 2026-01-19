package main

import (
	"Blog_aggreator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, user database.User) error{
	if len(cmd.Args)==0{
		return fmt.Errorf("Follow command expecting at least 1 command")
	}
	current_feed, err:= s.db.FeedByUrl(context.Background(),cmd.Args[0])
	if err!=nil{
		return fmt.Errorf("Error in getting feed by url: %v",err)
	}

	follower, err:= s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID         : uuid.New(),
		CreatedAt  : time.Now(),
		UpdatedAt  : time.Now(),
		UserID     : user.ID,
		FeedID     : current_feed.ID,
	})
	if err != nil{
		return fmt.Errorf("Error in creating feed: %v", err)
	}
	fmt.Printf("Current User Name: %v\nFeed name: %v\n",follower.UserName, follower.FeedName)
	return nil
}


func HandlerUnfollow(s *State, cmd Command, user database.User) error{
	if len(cmd.Args)==0{
		return fmt.Errorf("This command must requires url of the feed")
	}
	current_feed, err := s.db.FeedByUrl(context.Background(), cmd.Args[0])
	if err != nil{
		return fmt.Errorf("Error in getting Feed by Url: %v",err)
	}
	err = s.db.UnfollowFeed(context.Background(),database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: current_feed.ID,
	})
	if err!= nil {
		return fmt.Errorf("Error in Unfollow feed function: %v", err)
	}
	return nil
}