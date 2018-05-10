// +build OMIT

// Package main provides ...
package main // import "github.com/bantana/golabs/contextlabs/cancelshutdown"

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 建立一个长度为 1 的 chan 用于接收 os.Signal 信号
	stop := make(chan os.Signal, 1)
	// os.Interrupt signal 丢给 quit chan
	signal.Notify(stop, os.Interrupt)
	signal.Ignore(syscall.SIGINT)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: logHandler(http.DefaultServeMux),
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	go func() {
		<-stop
		proc, err := os.FindProcess(os.Getpid())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Shutting down server process %s pid %d", os.Args[0], proc.Pid)
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("could not shutdow %d", proc.Pid)
		}
	}()

	http.HandleFunc("/", indexH)

	log.Printf("listen server on port: %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	log.Println("Server gracefully stopped")
}

func indexH(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
