package config

import (
	"errors"
	"fmt"
)

// this will tell the handlers about the status of application state
type State struct{
	Config *Config
}

type Command struct{
	Name string
	Args []string
}

// this is the function signature of all command handlers
func HandlerLogin(s *State, cmd Command) error{
	if len(cmd.Args)==0{
		return errors.New("the login handler expects a single argument,the username")
	}
	s.Config.Current_user_name = cmd.Args[0]

	if err := s.Config.SetUser(cmd.Args[0]); err!= nil{
		return err
	}
	fmt.Printf("The user has been set")
	return nil
}

type Commands struct{
	Handlers map[string]func (*State, Command) error
}

func  (c *Commands) Run(s *State, cmd Command) error{
	_,ok := c.Handlers[cmd.Name]
	if !ok {
		 return errors.New("The command Does not exist")
	}
	run_command := c.Handlers[cmd.Name]
	err := run_command(s,cmd)
	if err != nil{
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State,Command) error){
	if c.Handlers==nil{
		c.Handlers = make(map[string]func(*State, Command) error)
	}
	_,ok := c.Handlers[name]
	if ok {
		 fmt.Errorf("The %v command exist",name)
	}
	c.Handlers[name] = f 
}