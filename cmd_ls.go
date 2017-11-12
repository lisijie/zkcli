package main

import (
    "github.com/urfave/cli"
    "path"
    "time"
    "fmt"
)

func init() {
    command := cli.Command{
        Name: "ls",
        Usage: "List child nodes",
        ArgsUsage: "path",
        Action: func(ctx *cli.Context) {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return
            }

            root := ctx.Args().First()
            children, _, err := conn.Children(root)
            if err != nil {
                println(err.Error())
                return
            }

            data := make([]map[string]string, 0, len(children))
            columns := []string{"Node", "Children", "Data", "Ctime", "Mtime", "Version"}
            for _, v := range children {
                cn := path.Join(root, v)
                d, st, err := conn.Get(cn)
                if err != nil {
                    println(err.Error())
                    continue
                }
                if len(d) > 50 {
                    d = d[:50]
                    d[47], d[48], d[49] = '.', '.', '.'
                }
                info := make(map[string]string)
                info["Node"] = cn
                info["Data"] = string(d)
                info["Ctime"] = time.Unix(int64(st.Ctime / 1000), st.Ctime % 1000).Format("2006-01-02 15:04:05")
                info["Mtime"] = time.Unix(int64(st.Mtime / 1000), st.Ctime % 1000).Format("2006-01-02 15:04:05")
                info["Children"] = fmt.Sprintf("%d", st.NumChildren)
                info["Version"] = fmt.Sprintf("%d", st.Version)
                data = append(data, info)
            }

            printTable(columns, data)
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}


