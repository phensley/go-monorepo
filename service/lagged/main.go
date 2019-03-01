package main

import "github.com/phensley/go-monorepo/pkg/logz"

// example of a service that stays fixed on a specific version
// of the monorepo's root module.

func main() {
	svc, _ := logz.NewService()
	svc.Info("hi from the past")
}
