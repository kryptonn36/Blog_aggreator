package main

import (
	"context"
	"fmt"

)

func HandlerReset(s *State, cmd Command) error{
	err := s.db.DeleteTable(context.Background())
	if err!= nil{
		return err
	}
	fmt.Printf("Database reset successfully")
	return nil
}