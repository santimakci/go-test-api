package utils

import "github.com/fatih/color"

// Imprime la ruta en consola con color cyan
func PrintRoute(route string) {
	c := color.New(color.FgCyan)
	c.DisableColor()
	c.Print("Route: ")
	c.EnableColor()
	c.Println(route)
}
