package main

import (
	"Blog_aggreator/internal/database"
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error{
	followings, err:= s.db.GetFeedFollowForUser(context.Background(),user.ID)
	if err!=nil{
		return fmt.Errorf("Error in getting feed follow: %v",err)
	}
	for _,following := range followings{
		fmt.Printf("%v\n",following.FeedName)
	}
	return nil
}