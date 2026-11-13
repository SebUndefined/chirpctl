package printing

type FormatBoolInput struct {
	Value       bool
	Styled      bool
	Colored     bool
	TextTrue    string
	TextFalse   string
	ColorTrue   Color
	ColorFalse  Color
	SymbolTrue  string
	SymbolFalse string
}

func FormatBool(in FormatBoolInput) string {
	if !in.Styled {
		if in.Value {
			return in.TextTrue
		}
		return in.TextFalse
	}

	if in.Value {
		symbol := in.SymbolTrue
		if in.Colored {
			symbol = StatusCheck(ColorGreen)
		}
		return symbol + " " + in.TextTrue
	}

	symbol := in.SymbolFalse
	if in.Colored {
		symbol = StatusSymbol(symbol, ColorGray)
	}
	return symbol + " " + in.TextFalse
}
