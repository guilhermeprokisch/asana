package commands

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/thash/asana/api"
)

func Done(c *cli.Context) {
	task := api.Update(api.FindTaskId(c.Args().First(), false), "completed", "true")
	fmt.Println("DONE! : " + task.Name)
}
