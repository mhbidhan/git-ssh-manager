package utils

import (
	"fmt"

	"github.com/mgutz/ansi"
	"github.com/mhbidhan/git-ssh-manager/internal/constants"
)

// PrintAllCommands outputs all available commands.
func PrintAllCommands() {
	fmt.Print(constants.Commands)
}

// PrintSuccessMessage prints a success message.
func PrintSuccessMessage(message string) {
	fmt.Print(ansi.ColorCode("green+h:black"), message+"\n", ansi.ColorCode("reset"))
}

// PrintErrorMessage prints an error message.
func PrintErrorMessage(message string) {
	fmt.Print(ansi.ColorCode("red+h:black"), message+"\n", ansi.ColorCode("reset"))
}
