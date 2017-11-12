package main

import (
    "github.com/urfave/cli"
    "fmt"
    "path/filepath"
)

func init() {
    command := cli.Command{
        Name: "tree",
        Usage: "List nodes in a tree-like format",
        ArgsUsage: "path",
        Flags: []cli.Flag{
            cli.IntFlag{
                Name: "l",
                Usage: "Max display depth of the tree.",
            },
        },
        Action: func(ctx *cli.Context) {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return
            }
            maxLevel := ctx.Int("l")
            root := ctx.Args().First()
            fmt.Println(root)
            listTree(root, "", 0, maxLevel)
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}

func listTree(path string, pre string, level int, maxLevel int) {
    if maxLevel > 0 && level >= maxLevel {
        return
    }
    children, _, err := conn.Children(path)
    if err != nil {
        return
    }
    count := len(children)
    for k, item := range children {
        if k + 1 == count {
            fmt.Printf("%s%s%s\n", pre, "└── ", item)
            listTree(filepath.Join(path, item), pre + "    ", level + 1, maxLevel)
        } else {
            fmt.Printf("%s%s%s\n", pre, "├── ", item)
            listTree(filepath.Join(path, item), pre + "│   ", level + 1, maxLevel)
        }
    }
}