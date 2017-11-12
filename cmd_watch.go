package main

import (
    "github.com/urfave/cli"
    "fmt"
    "github.com/samuel/go-zookeeper/zk"
    "time"
)

func init() {
    command := cli.Command{
        Name: "watch",
        Usage: "Watch path",
        ArgsUsage: "path",
        Action: func(ctx *cli.Context) {
            if len(ctx.Args()) < 1 {
                cli.ShowSubcommandHelp(ctx)
                return
            }
            path := ctx.Args().First()
            fmt.Printf("watching %s ...\n", path)
            events := make(chan zk.Event)
            quit := make(chan bool)
            go func() {
                for {
                    // 当新增、删除子节点，或者自身数据有修改时，会收到事件通知
                    _, _, ec, err := conn.ChildrenW(path)
                    if err != nil {
                        fmt.Println(err)
                        quit <- true
                        return
                    }
                    events <- (<-ec)
                }
            }()

            for {
                select {
                case ev := <-events:
                    fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), ev.Type)
                case <-quit:
                    return
                }
            }
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}
