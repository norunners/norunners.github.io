package main

import (
	"github.com/norunners/norunners.github.io/vue/app"
	"github.com/norunners/vue"
)

type Data struct {
	Message string
}

func main() {
	vue.New(
		vue.El("#app"),
		vue.Template("<app/>"),
		vue.Sub("app", app.NewComp()),
	)

	select {}
}
