package app

import "fmt"

type App struct {
    Name string
}

func NewApp(name string) *App {
    return &App{Name: name}
}

func (a *App) Run() {
    fmt.Printf("Running application: %s\n", a.Name)
}