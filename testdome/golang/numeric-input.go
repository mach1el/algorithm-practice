package main
import (
	"fmt"
	"unicode"
)
type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (char *NumericInput) Add(input rune) {
	if unicode.IsDigit(input) {
		char.input += string(input)
	}
}

func (char *NumericInput) GetValue() string {
	return char.input
}

func main() {
  var input UserInput = &NumericInput{}
  input.Add('1')
  input.Add('a')
  input.Add('0')
  fmt.Println(input.GetValue())
}