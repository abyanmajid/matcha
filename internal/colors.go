package internal

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Cyan    = "\033[36m"
	Magenta = "\033[35m"
	White   = "\033[37m"
)

// PrintRed prints text in red
func PrintRed(text string) {
	fmt.Println(Red + text + Reset)
}

// PrintGreen prints text in green
func PrintGreen(text string) {
	fmt.Println(Green + text + Reset)
}

// PrintYellow prints text in yellow
func PrintYellow(text string) {
	fmt.Println(Yellow + text + Reset)
}

// PrintBlue prints text in blue
func PrintBlue(text string) {
	fmt.Println(Blue + text + Reset)
}

// PrintCyan prints text in cyan
func PrintCyan(text string) {
	fmt.Println(Cyan + text + Reset)
}

// PrintMagenta prints text in magenta
func PrintMagenta(text string) {
	fmt.Println(Magenta + text + Reset)
}

// PrintWhite prints text in white
func PrintWhite(text string) {
	fmt.Println(White + text + Reset)
}

// PrintColored prints text in a specified color
func PrintColored(text string, color string) {
	fmt.Println(color + text + Reset)
}

type IMatchaArt struct {
	Addr string
}

func MatchaArt(addr string) *IMatchaArt {
	return &IMatchaArt{
		Addr: addr,
	}
}

func (m *IMatchaArt) PrintIntro() {
	matchaFigure := figure.NewColorFigure("matcha", "roman", "green", true)
	matchaFigure.Print()

	msg := fmt.Sprintf("✨ A Matcha web server has been initiated! ✨\n")

	PrintYellow(msg)
}
