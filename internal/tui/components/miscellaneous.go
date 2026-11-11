package components

import (
	"fmt"
	"os"
	"strings"
)

func IsTTY() bool {
	fi, _ := os.Stdout.Stat()
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func IsUTF8Locale() bool {
	// crude but effective
	env := strings.ToLower(os.Getenv("LC_ALL") + " " + os.Getenv("LANG"))
	return strings.Contains(env, "utf-8") || strings.Contains(env, "utf8")
}

var DisableEmoji bool

func Emoji(s string) string {
	if DisableEmoji || !IsTTY() || !IsUTF8Locale() {
		return "" // fallback silencieux
	}
	return s
}

func PrintlnEmoji(emoji string, parts ...string) {
	msg := strings.Join(parts, " ")
	if e := Emoji(emoji); e != "" {
		fmt.Println(e, msg)
		return
	}
	fmt.Println(msg)
}
