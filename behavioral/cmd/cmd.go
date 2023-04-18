package cmd

type Command interface {
	Execute()
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}
