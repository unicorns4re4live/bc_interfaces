package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *cli) ReadCommands() {
	for cmd := range c.commands {
		fmt.Println("CMD ID:", cmd.id)
		fmt.Println("ARGS:", cmd.args)
		switch cmd.id {
		case "quit":
			exit()
		case "auth":
			auth(c,cmd.args)
			fmt.Println("AUTH")
		case "templates":
			get_template_names(c)
			fmt.Println("TEMPLATES")
		case "template":
			get_template_by_name(c, cmd.args)
			fmt.Println("TEMPLATE BY NAME")
		case "userlist": {
			get_userlist(c)
			fmt.Println("USERLIST")
		}
		case "user": {
			get_user(c, cmd.args)
			fmt.Println("USER")
		}
		case "create": {
			create_user(c, cmd.args)
			fmt.Println("USER CREATED")

		}

		case "transfer": {
			create_transfer(c, cmd.args)
			fmt.Println("TRANSFER CREATED")
		}


		case "admin_amount":
			admin_amount(c)
			fmt.Println("ADMIN AMOUNT")

		default:
			fmt.Println("Комманда отсутствует")
		}
		fmt.Println("Новая команда:")
	}
}

func (c *cli) PushCommands() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Вводите комманды в формете \\commandName")
	for {
		commandData, err := consoleReader.ReadString('\n')

		if err != nil{
			fmt.Println("Проблемы с вводом - попробуйте ещё")
			continue
		}
		commandDataSplited := strings.Split(commandData, " ")

		commandId := strings.Replace(commandDataSplited[0], "\\", "", -1)
		commandId = strings.TrimSpace(commandId)
		c.commands <- command{id: commandId, args: commandDataSplited[1:]}
	}
}

func NewCli() *cli {
	return &cli{
		commands: make(chan command),
	}
}
