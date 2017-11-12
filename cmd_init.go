package main

import (
    "github.com/urfave/cli"
    "strings"
    "encoding/json"
    "fmt"
    "io/ioutil"
)

func init() {
    command := cli.Command{
        Name: "init",
        Usage: "Initialize application",
        ArgsUsage: "addrs",
        Action: func(ctx *cli.Context) {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return
            }
            addrs := strings.Split(ctx.Args().First(), ",")
            config.Addrs = addrs
            data, err := json.Marshal(config)
            if err != nil {
                fmt.Println(err)
                return
            }
            if err := ioutil.WriteFile(confFile, data, 0700); err != nil {
                fmt.Println(err)
            }
        },
    }
    app.Commands = append(app.Commands, command)
}


