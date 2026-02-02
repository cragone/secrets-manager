package actions

import (
	"fmt"
	"os/exec"
)

type Command string

const (
	CreateCommand Command = "create"
	ReadCommand   Command = "read"
	HelpCommand   Command = "help"
)

type Action func(args []string) error

var Commands = map[Command]Action{}

func Init() {
	Commands[HelpCommand] = Help
	Commands[CreateCommand] = Create
}

func Help(args []string) error {
	if len(args) != 0 {
		fmt.Printf("%s is invalid\n", args)
	}

	fmt.Println("Try one of the following commands")
	for command := range Commands {
		fmt.Println(command)
	}
	return nil
}

func Create(args []string) error {
	secretName := args[0]
	fmt.Printf("creating kubernetes secret: %s\n", secretName)

	cmd := exec.Command("echo", secretName)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not create secret %q: %w", secretName, err)
	}

	return nil
}
