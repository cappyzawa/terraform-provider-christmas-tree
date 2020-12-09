package provider

import (
	"errors"
	"strings"

	cl "github.com/gookit/color"
)

// Colorize colorizes str
// arg: color based on https://pkg.go.dev/gopkg.in/gookit/color.v1#pkg-constants
func Colorize(basetext string, char string, color string) (string, error) {

	c, ok := cl.FgColors[color]
	if !ok {
		return "", errors.New(color + " is unsupported")
	}

	return strings.ReplaceAll(basetext, char, c.Sprint(char)), nil
}
