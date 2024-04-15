package main

import (
    "context"
    "api/internal/presenter"
)

func main() {
	srv := presenter.NewServer()
	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
