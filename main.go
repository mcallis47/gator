package main

import (
	"fmt"
	"os"

	"github.com/mcallis47/gator/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func main() {
	cur_config := config.Read()
	s := &state{Config: &cur_config}
	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no command provided")
		os.Exit(1)
	}
	cmd := command{Name: args[0], Args: args[1:]}
	err := cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login requires 1 argument")
	}
	s.Config.SetUser(cmd.Args[0])
	return nil
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.handlers[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command %q", cmd.Name)
	}
	return handler(s, cmd)
}
