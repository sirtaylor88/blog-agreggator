package main

import (
	"fmt"
	"os"

	"github.com/sirtaylor88/go-blog-agreggator/internal/config"
)

type state struct {
	config config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login comamnd needs an username.")
	}
	if err := config.SetUser(cmd.args[0], s.config); err != nil {
		return err
	}
	fmt.Println("The user is set.")
	return nil
}

type commands struct {
	command_by_name map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.command_by_name[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.command_by_name[cmd.name]
	if !ok {
		return fmt.Errorf("Command do not exist.")
	}
	if err := f(s, cmd); err != nil {
		return fmt.Errorf("Error when running command: %w", err)
	}
	return nil
}

func main() {
	cfg, _ := config.Read()
	s := &state{
		config: cfg,
	}

	var c commands
	c.command_by_name = make(map[string]func(*state, command) error)
	c.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Missing argument for command.")
		os.Exit(1)
	}
	cmd := command{
		name: args[1],
		args: args[2:],
	}
	if err := c.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
