package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"golang.org/x/sync/errgroup"

)

func helloServer1(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"hello,world!")
}

func StartHttpServer(server *http.Server) error {
	http.HandleFunc("/hello", helloServer1)
	fmt.Println("HttpServer start")
	err := server.ListenAndServe()
	return err
}

func  main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr:":9091",Handler:http.DefaultServeMux}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func()error {
		<- errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})

	channel := make(chan os.Signal, 1)
	signal.Notify(channel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-channel:
				cancel()
			}
		}
		return nil
	})
	if err := group.Wait();err !=nil {
		fmt.Println("group error:", err)
	}
	fmt.Println("all group done!")

}