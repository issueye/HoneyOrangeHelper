package main

import (
	"HoneyOrangeHelper/internal/initialize"
	"HoneyOrangeHelper/pages/home"

	"github.com/ying32/govcl/vcl"
)

func main() {
	initialize.Init()
	// vcl.RunApp(&home.FrmHome)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&home.FrmHome)
	vcl.Application.Run()
}
