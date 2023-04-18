package cmd

import "testing"

func TestCommand(t *testing.T) {
	tv := &TV{}

	onCmd := &OnCommand{tv}
	offCmd := &OffCommand{tv}

	onBtn := &Button{onCmd}
	onBtn.Press()

	offBtn := &Button{offCmd}
	offBtn.Press()
}
