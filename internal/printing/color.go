package printing

import "github.com/jedib0t/go-pretty/v6/text"

type Color string

const (
	ColorGreen  Color = "green"
	ColorGray   Color = "gray"
	ColorYellow Color = "yellow"
	ColorRed    Color = "red"
	ColorBlue   Color = "blue"
)

func StatusSymbol(symbol string, c Color) string {
	switch c {
	case ColorGreen:
		return text.FgGreen.Sprint(symbol)
	case ColorGray:
		return text.FgHiBlack.Sprint(symbol)
	case ColorYellow:
		return text.FgYellow.Sprint(symbol)
	case ColorRed:
		return text.FgRed.Sprint(symbol)
	case ColorBlue:
		return text.FgCyan.Sprint(symbol)
	default:
		return symbol
	}
}
