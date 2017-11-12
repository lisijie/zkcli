package main

import (
    "github.com/urfave/cli"
    "fmt"
    "path/filepath"
)

func init() {
    command := cli.Command{
        Name: "rm",
        Usage: "Delete path",
        ArgsUsage: "path",
        Flags: []cli.Flag{
            cli.BoolFlag{
                Name: "r",
                Usage: "Delete the supplied path and all child nodes.",
            },
        },
        Action: func(ctx *cli.Context) {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return
            }
            r := ctx.Bool("r")
            path := ctx.Args().First()

            if exist, _, err := conn.Exists(path); !exist || err != nil {
                if err != nil {
                    fmt.Printf("Delete %s error: %s\n", path, err)
                    return
                }
                if !exist {
                    fmt.Printf("Delete %s error: %s\n", path, "path not exists.")
                    return
                }
            }
            removeDeep(path, r)
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}

func removeDeep(path string, recursive bool) {
    children, _, _ := conn.Children(path)
    if len(children) > 0 && recursive {
        for _, node := range children {
            p := filepath.Join(path, node)
            removeDeep(p, recursive)
        }
    }
    err := conn.Delete(path, -1)
    if err != nil {
        fmt.Printf("Delete %s error: %s\n", path, err)
        return
    } else {
        fmt.Println("Deleted", path)
    }
}
