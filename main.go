package main

import (
	"os"
	"github.com/urfave/cli"
)

var Version string = "0.1"

func main() {
	app := cli.newApp()

	app.Name = "poem"
	app.Usage = "Be Easy to Write Diaries"
	app.Version = Version
	app.Auther = "shamisonn"
	app.Email = "sh4m1s0n@gmail.com"
	app.Commnads = Commnads
	app.Run(os.Args)
}



