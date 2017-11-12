package main

import (
    "fmt"
    "strings"
    "reflect"
)

func printObj(obj interface{}) {
    t := reflect.TypeOf(obj).Elem()
    v := reflect.ValueOf(obj).Elem()
    for i := 0; i < t.NumField(); i++ {
        fmt.Println(t.Field(i).Name, "=", v.Field(i).Interface())
    }
}

func printTable(columns []string, data []map[string]string) {
    if len(data) == 0 {
        return
    }
    fields := make(map[string]int)
    for _, row := range data {
        for k, v := range row {
            if l, ok := fields[k]; !ok {
                if len(k) > len(v) {
                    fields[k] = len(k)
                } else {
                    fields[k] = len(v)
                }
            } else {
                if len(v) > l {
                    fields[k] = len(v)
                }
            }
        }
    }
    space := 2
    max := 0
    for _, name := range columns {
        l := fields[name]
        fmt.Print(name)
        fmt.Print(strings.Repeat(" ", l - len(name) + space))
        max = max + l + space
    }
    fmt.Print("\n")
    fmt.Println(strings.Repeat("-", max))

    for _, row := range data {
        for _, name := range columns {
            l := fields[name]
            v, ok := row[name]
            if ok {
                fmt.Print(v)
                fmt.Print(strings.Repeat(" ", l - len(v) + space))
            } else {
                fmt.Print(strings.Repeat(" ", l + space))
            }
        }
        fmt.Print("\n")
    }
}

