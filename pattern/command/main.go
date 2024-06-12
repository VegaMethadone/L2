package main

import "fmt"

/*
Паттерн "Команда" (Command) позволяет инкапсулировать запрос на выполнение определенного действия в виде отдельного объекта. Этот объект запроса на действие и называется командой. При этом объекты, инициирующие запросы на выполнение действия, отделяются от объектов, которые выполняют это действие.

Команды могут использовать параметры, которые передают ассоциированную с командой информацию. Кроме того, команды могут ставиться в очередь и также могут быть отменены.

Когда использовать команды?
Когда надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие действия. То есть когда необходимы функции обратного действия в ответ на определенные действия.

Когда необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.

Когда надо поддерживать логгирование изменений в результате запросов. Использование логов может помочь восстановить состояние системы - для этого необходимо будет использовать последовательность запротоколированных команд.
*/

type Command interface {
	Execute()
}

type Light struct{}
type LightOnCommand struct{ light *Light }
type LightOffCommand struct{ light *Light }
type RemoteControl struct{ command Command }

func (l *Light) On()  { fmt.Println("Light is On") }
func (l *Light) Off() { fmt.Println("Light is Off") }

func (c *LightOnCommand) Execute()  { c.light.On() }
func (c *LightOffCommand) Execute() { c.light.Off() }

func (r *RemoteControl) SetCommand(command Command) { r.command = command }
func (r *RemoteControl) PressButton()               { r.command.Execute() }

func main() {
	light := &Light{}

	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()
}
