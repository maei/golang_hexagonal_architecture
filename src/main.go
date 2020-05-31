package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_hexagonal_architecture/src/controller"
	mr "github.com/maei/golang_hexagonal_architecture/src/repository/mongodb"
	"github.com/maei/golang_hexagonal_architecture/src/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	repo := repo()
	service := service.NewSumService(repo)
	handler := controller.NewSumController(service)

	router := echo.New()

	router.POST("/sum", handler.NewCompute)
	router.GET("/sum/:code", handler.FindResult)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- router.Start(httpPort())

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
	repo.Disconnect(context.Background())

}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func repo() service.SumRepositoryInterface {
	//mongoURL := os.Getenv("MONGO_URL")
	//mongodb := os.Getenv("MONGO_DB")
	//mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	mongoURL := "mongodb://localhost:27017"
	mongodb := "sum"
	mongoTimeout := 30
	repo, err := mr.NewMongoSumRepository(mongoURL, mongodb, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}
	return repo

}
