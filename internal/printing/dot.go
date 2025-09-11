package printing

import "github.com/jedib0t/go-pretty/v6/text"

type DotColor string

const (
	DotGreen  DotColor = "green"
	DotGray   DotColor = "gray"
	DotYellow DotColor = "yellow" // “orange” approché
	DotRed    DotColor = "red"
	DotBlue   DotColor = "blue"
)

func StatusDot(c DotColor) string {
	switch c {
	case DotGreen:
		return text.FgGreen.Sprint("●")
	case DotGray:
		return text.FgHiBlack.Sprint("●")
	case DotYellow:
		return text.FgYellow.Sprint("●")
	case DotRed:
		return text.FgRed.Sprint("●")
	case DotBlue:
		return text.FgCyan.Sprint("●")
	default:
		return "●"
	}
}
