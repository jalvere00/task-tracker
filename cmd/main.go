package main

import (
    "github.com/jalvere00/task-tracker/internal/app"
    "log"
)

func main() {
    a := app.NewApp()
    if err := a.Run(); err != nil {
        log.Fatalf("Error running the application: %v", err)
    }
}