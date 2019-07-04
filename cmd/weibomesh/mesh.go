package main

import (
	"github.com/idevz/kwm/cmd/weibomesh/app"
	"os"
)

func main()  {
	command := app.NewWeiboMeshCommand()

	if err:= command.Execute();err != nil {
		os.Exit(1)
	}
}