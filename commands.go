package main

type Command string

const (
	CreateCommand Command = "create"
	ReadCommand   Command = "read"
	HelpCommand   Command = "help"
)

type Action func() error

var Commands map[Command]Action = {
	HelpCommand: Help(),
}

func Help() error{
	for command := range Commands{
		fmt.Println("Options:", command)
	}
		return nil

}
