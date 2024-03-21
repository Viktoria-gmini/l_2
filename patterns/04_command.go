//паттерн команда

package main

import "fmt"

// Интерфейс команды
type Command interface {
	Execute()
}

// Конкретная реализация команды
type ConcreteCommand struct {
	receiver Receiver
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Получатель команды
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver is performing an action")
}

// Инвокер команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.command = cmd
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := Receiver{}
	command := &ConcreteCommand{receiver: receiver}

	invoker := Invoker{}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
