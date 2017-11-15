package main

import (
    "github.com/urfave/cli"
    "github.com/samuel/go-zookeeper/zk"
)

func init() {
    command := cli.Command{
        Name: "getacl",
        Usage: "Get ACL list",
        ArgsUsage: "path",
        Action: func(ctx *cli.Context) {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return
            }

            path := ctx.Args().First()
            acls, _, err := conn.GetACL(path)
            if err != nil {
                println(err.Error())
                return
            }

            data := make([]map[string]string, 0, len(acls))
            columns := []string{"Scheme", "ID", "Perms"}
            for _, acl := range acls {
                info := make(map[string]string)
                info["Scheme"] = acl.Scheme
                info["ID"] = acl.ID
                perms := make([]byte, 0)
                if acl.Perms & zk.PermCreate > 0 {
                    perms = append(perms, 'c')
                }
                if acl.Perms & zk.PermDelete > 0 {
                    perms = append(perms, 'd')
                }
                if acl.Perms & zk.PermRead > 0 {
                    perms = append(perms, 'r')
                }
                if acl.Perms & zk.PermWrite > 0 {
                    perms = append(perms, 'w')
                }
                if acl.Perms & zk.PermAdmin > 0 {
                    perms = append(perms, 'a')
                }
                info["Perms"] = string(perms)
                data = append(data, info)
            }

            printTable(columns, data)
        },
    }
    app.Commands = append(app.Commands, command)
    app.BashComplete = nodeCompletion
}
