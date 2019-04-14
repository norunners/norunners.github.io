package app

import "github.com/norunners/vue"

//go:generate vueg

type Data struct {
	Message string
}

func NewComp() *vue.Comp {
	return vue.Component(
		vue.Template(appTmpl),
		vue.Data(Data{Message: "Hello WebAssembly!"}),
	)
}
