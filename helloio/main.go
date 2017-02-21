package main

import "github.com/gopherjs/jquery"

const (
	INPUT  = "input#name"
	OUTPUT = "span#output"
)

func main() {
	// Convenience.
	jq := jquery.NewJQuery
	// On page load.
	jq().Ready(func() {
		// Show jQuery Version on console.
		print("Your current jQuery version is: " + jq().Jquery)

		// Catch keyup events on input#name element.
		jq(INPUT).On(jquery.KEYUP, func(e jquery.Event) {

			name := jq(e.Target).Val()
			name = jquery.Trim(name)

			if len(name) > 0 {
				jq(OUTPUT).SetText("Welcome to hello IO, " + name + "!")
			} else {
				jq(OUTPUT).SetText("Welcome to hello IO!")
			}
		})
	})
}
