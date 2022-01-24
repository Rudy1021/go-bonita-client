package main

import (
	"fmt"

	bpm "github.com/kuochaoyi/go-bonita-client/bpm"
)

func main() {
	bpm.Login("isabelle_wu")
	body := bpm.StartForm("8759976868088592450", `{
		"modelInput":
		{
				"assistant":"choc",
				"recipient":"kevin_lin"
		}
}`)
	fmt.Println(body)
}
