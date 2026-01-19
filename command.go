package main

import(
	"errors"
)
type Command struct{
	Name string
	Args []string
}

type Commands struct{
	registeredcommands map[string]func (*State, Command) error
}

func  (c *Commands) Run(s *State, cmd Command) error{
	_,ok := c.registeredcommands[cmd.Name]
	if !ok {
		 return errors.New("The command Does not exist")
	}
	run_command := c.registeredcommands[cmd.Name]
	err := run_command(s,cmd)
	if err != nil{
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State,Command) error){
	c.registeredcommands[name] = f 
}
