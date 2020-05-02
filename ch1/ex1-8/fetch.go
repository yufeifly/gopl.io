package main

import (
    "fmt"
    //"io/ioutil"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url,"http://") {
            url = "http://" + url
            fmt.Println(url)
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        //var b []byte
        //b, err := ioutil.ReadAll(resp.Body)
        io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        //fmt.Println("\nstatus code :", resp.Status)
        /*if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)*/
    }
}