package main

import (
    "github.com/samuel/go-zookeeper/zk"
    "time"
    "github.com/urfave/cli"
    "os"
    "fmt"
    "path/filepath"
    "runtime"
    "encoding/json"
    "io/ioutil"
)

var (
    app = cli.NewApp()
    conn *zk.Conn
    confFile string
    config struct {
        Addrs []string
    }
)

type nullLogger struct{}

func (nullLogger) Printf(format string, a ...interface{}) {}

func main() {
    dir := os.Getenv("HOME")
    if dir == "" && runtime.GOOS == "windows" {
        dir = os.Getenv("APPDATA")
        if dir == "" {
            dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", "zkcli")
        }
        dir = filepath.Join(dir, "zkcli")
    } else {
        dir = filepath.Join(dir, ".config", "zkcli")
    }
    if err := os.MkdirAll(dir, 0700); err != nil {
        println(err.Error())
        os.Exit(1)
    }
    confFile = filepath.Join(dir, "config.json")

    app.Name = "zkcli"
    app.Usage = "zookeeper client tool"
    app.Author = "sijie.li"
    app.Email = "lsj86@qq.com"
    app.Version = "0.0.1"
    app.EnableBashCompletion = true
    app.Before = initialize
    app.After = func(c *cli.Context) error {
        if conn != nil {
            conn.Close()
        }
        return nil
    }
    app.Setup()
    app.Run(os.Args)
}

func initialize(ctx *cli.Context) error {
    if ctx.Args().First() == "" || ctx.Args().First() == "init" {
        return nil
    }

    b, err := ioutil.ReadFile(confFile)
    if err != nil {
        return err
    }
    err = json.Unmarshal(b, &config)
    if err != nil {
        return fmt.Errorf("could not unmarshal %v: %v", confFile, err)
    }

    zk.DefaultLogger = nullLogger{}
    conn, _, err = zk.Connect(config.Addrs, time.Second)
    return err
}

func nodeCompletion(ctx *cli.Context) {
    path := ctx.Args().First()
    children, _, _ := conn.Children(path)
    if len(children) > 0 {
        for _, item := range children {
            fmt.Fprintln(ctx.App.Writer, item)
        }
    }
}