package output

import (
	"fmt"
	"os"

	"github.com/kyokomi/emoji"
)

const (
	Green Color = "\033[32m"
	Red         = "\033[31m"
)

type Color string

func Print(message string, color ...Color) {
	emojiMsg := emoji.Sprint(message)
	if len(color) > 0 {
		fmt.Println(color[0], emojiMsg)
	} else {
		fmt.Println(emojiMsg)
	}
}

func PrintErrorAndExit(err error) {
	fmt.Println(Red, err.Error())
	os.Exit(1)
}

