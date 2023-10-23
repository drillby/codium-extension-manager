package cli

import "fmt"

type Cli struct {
	Commands []Comand
}

type Comand struct {
	Name string
	Help string
	Run  func(args []string) error
}

func (c *Cli) AddComand(comand Comand) {
	c.Commands = append(c.Commands, comand)
}

func (c *Cli) Run(args []string) error {
	if len(args) < 1 {
		return c.printHelp()
	}

	for _, command := range c.Commands {
		if command.Name == args[0] {
			return command.Run(args[1:])
		}
	}

	return c.printHelp()
}

func (c *Cli) printHelp() error {
	println("Usage: cem <command> [arguments]")
	for _, command := range c.Commands {
		fmt.Printf("%-10s %s\n", command.Name, command.Help)
	}

	return nil
}
