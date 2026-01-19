package main

import (
	"context"
	"fmt"
)


func users(s *State, cmd Command) error{
	str, err:= s.db.GetUsers(context.Background())
	if err!= nil{
		return fmt.Errorf("Error in Getting the list of users: %v",err)
	}
	current_user := s.Config.Current_user_name
	for _,user := range(str){
		if user == current_user{
			fmt.Printf("%v (current)\n",user)
		}else{
			fmt.Printf("%v\n",user)
		}
	}
	return nil
}