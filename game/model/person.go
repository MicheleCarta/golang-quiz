package model

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
}

// Print Data
func (h *Person) Print(output string) {
	fmt.Println(output)
}

// Input Data
func (h *Person) Input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
