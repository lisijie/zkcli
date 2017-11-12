package main

import (
    "github.com/urfave/cli"
)

func init() {
    command := cli.Command{
        Name: "set",
        Usage: "Set path data",
        ArgsUsage: "path data",
        Action: func(ctx *cli.Context) {
            if len(ctx.Args()) < 2 {
                cli.ShowSubcommandHelp(ctx)
                return
            }
            path, data := ctx.Args().Get(0), ctx.Args().Get(1)
            stat, err := conn.Set(path, []byte(data), -1)
            if err != nil {
                println("Error:", err.Error())
                return
            }
            println("Set", path)
            println("")
            println("STAT:")
            printObj(stat)
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}

