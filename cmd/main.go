package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecthomas/kong"

	internalHttp "github.com/azhinu/kinklist/internal/http"
	"github.com/azhinu/kinklist/internal/storage"
)

var Version = "dev" // будет переписываться в GitHub Actions

var cli struct {
    Version bool   `name:"version" short:"v" help:"Print version information and quit"`
    DBPath  string `name:"path" short:"p" env:"KL_DB_PATH" default:"./kinklist.db" help:"Path to DB file" type:"path"`
    Port    int    `name:"port" short:"P" env:"KL_PORT" default:"8080" help:"Listening port"`
}

func main() {
    kong.Parse(&cli)

		fmt.Println(cli.DBPath)
    if cli.Version {
        fmt.Println("kinklist version:", Version)
        return
    }

    db, err := storage.NewDB(cli.DBPath)
    if err != nil {
        log.Fatalf("DB error: %v", err)
    }
    defer db.Close()

    router := internalHttp.NewRouter(db)

    addr := fmt.Sprintf(":%d", cli.Port)
    log.Printf("Starting server at %s", addr)

    if err := http.ListenAndServe(addr, router); err != nil {
        log.Fatalf("Server error: %v", err)
    }
}
