package main

import (
	"context"
	"fmt"
	"github.com/sonereker/distributed-go-sandbox/basic/registry"
	"log"
	"net/http"
)

func main() {
	registry.SetupRegistryService()

	http.Handle("/services", &registry.Service{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		_, _ = fmt.Scanln(&s)
		_ = srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
