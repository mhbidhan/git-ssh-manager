package utils

import (
	"fmt"

	"github.com/mgutz/ansi"
	"github.com/mhbidhan/git-ssh-manager/src/constants"
)

func PrintAllCommands() {
	fmt.Printf(constants.Commands)
}

func PrintSuccessMessage(message string) {
	fmt.Print(ansi.ColorCode("green+h:black"), message+"\n", ansi.ColorCode("reset"))
}

func PrintErrorMessage(message string) {
	fmt.Print(ansi.ColorCode("red+h:black"), message+"\n", ansi.ColorCode("reset"))
}

