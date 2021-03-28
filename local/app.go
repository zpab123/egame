package main

import (
	"fmt"
	"time"
)

var (
	app *App = newApp()
)

type App struct {
	Ens []*Entity
}

func (a *App) NewEntity() {
	e := Entity{
		ID: id,
	}

	id++

	a.Ens = append(a.Ens, &e)
}

func (a *App) Run() {
	for {
		a.Update()
		time.Sleep(3 * time.Second)
	}
}

func (a *App) Update() {
	a.NewEntity()
	var out string
	for _, e := range a.Ens {
		out = fmt.Sprintf("entity id=%d", e.ID)
		fmt.Println(out)
	}
}

func newApp() *App {
	ap := App{
		Ens: make([]*Entity, 0),
	}

	return &ap
}
