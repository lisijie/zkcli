package main

import (
    "github.com/urfave/cli"
    "github.com/samuel/go-zookeeper/zk"
)

func init() {
    command := cli.Command{
        Name: "create",
        Usage: "Create path",
        ArgsUsage: "path",
        Action: func(ctx *cli.Context) {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return
            }
            path := ctx.Args().First()
            data := ctx.Args().Get(1)
            res, err := conn.Create(path, []byte(data), 0, zk.WorldACL(zk.PermAll))
            if err != nil {
                println("Error:", err.Error())
                return
            }
            println("Created", res)
            println("")
            _, stat, _ := conn.Get(path)
            println("STAT:")
            printObj(stat)
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}
