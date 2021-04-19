package main

import (
	"context"
	"fmt"
	"github.com/sonereker/distributed-go-sandbox/basic/grades"
	"github.com/sonereker/distributed-go-sandbox/basic/log"
	"github.com/sonereker/distributed-go-sandbox/basic/registry"
	"github.com/sonereker/distributed-go-sandbox/basic/service"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(), host, port, r, grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Log service found at %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
