package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloat(hint string) (float64, error) {
	text, err := getText(hint)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(strings.TrimSpace(text), 64)
}

var userI float64

func getText(hint string) (string, error) {
	fmt.Print(hint)
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
