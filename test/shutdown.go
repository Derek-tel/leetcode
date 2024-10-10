package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	userService := UserService{
		db:   db{},
		amqp: amqp{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/user", func(rw http.ResponseWriter, req *http.Request) {
		name := "some name" // let's imagine we got it from the request
		if err := userService.RegisterUser(req.Context(), name); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)

	srv := http.Server{}
	srv.Addr = ":8080"
	srv.Handler = r

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(c); err != nil { // Use here context with a required timeout
		log.Printf("server shutdown returned an err: %v\n", err)
	}
	log.Printf("server shutdown returned")

	//userService.Stop(context.TODO()) // Use here context with a required timeout

	log.Println("final")
}

func ranges(input ...interface{}) {
	for _, v := range input {
		fmt.Println(v)
	}
}

type UserService struct {
	amqp amqp
	db   db

	doneWG sync.WaitGroup
}

func (s *UserService) Stop(ctx context.Context) {
	log.Println("waiting for user service to finish")
	doneChan := make(chan struct{})
	go func() {
		s.doneWG.Wait()
		close(doneChan)
	}()
	select {
	case <-ctx.Done():
		log.Println("context done earlier then user service has stopped")
	case <-doneChan:
		log.Println("user service finished")
	}
}

func (s *UserService) RegisterUser(ctx context.Context, name string) error {
	log.Println("start user registration")

	_, err := s.db.InsertUser(ctx, name)
	if err != nil {
		return fmt.Errorf("db insertion failed: %v", err)
	}

	//s.publishUserInserted(ctx, userID)
	log.Println("end user registration")
	return nil
}

func (s *UserService) publishUserInserted(ctx context.Context, userID int) {
	s.doneWG.Add(1)
	go func() {
		defer s.doneWG.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Printf("publishUserInserted recovered panic: %v\n", err)
			}
		}()
		s.amqp.PublishUserInserted(ctx, userID)
	}()
}

type db struct{}

func (d db) InsertUser(ctx context.Context, name string) (int, error) {
	time.Sleep(time.Second * 20)
	log.Println("user insert")
	return 1, nil
}

type amqp struct{}

func (a amqp) PublishUserInserted(ctx context.Context, id int) {
	time.Sleep(time.Second * 20)
	log.Println("message publish")
}
