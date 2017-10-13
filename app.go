package main

import (
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"time"
	"syscall"

	"github.com/alvintzz/nyanyangku/ui"
)

func PongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	log.Println("DSA")
}

func initHandler(mux *http.ServeMux) {
	mux.HandleFunc("/ping", PongHandler)
	mux.HandleFunc("/", IndexHandler)
}

func main() {
	//Create a mux for routing incoming requests
	mux := http.NewServeMux()
	initHandler(mux)
	ui.InitHandler(mux, config.Settings.TemplateDir)
	ui.InitDb(masterDB)

	//Handle Static file (css/js) request
	fs := http.FileServer(http.Dir(config.Settings.PublicDir))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//Create server to serve all incoming request with registered mux
	server := &http.Server{
		Addr:    config.Settings.SelfPort,
		Handler: mux,
	}

	//Register stop signal and channel that wait for the signal. If it get the signal,
	//then it will be inserted into the channel
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	//Run the API using go func while the channel waiting for signal
	go func() {
		log.Printf("Listening to %s", config.Settings.SelfURL)
		log.Fatal(server.ListenAndServe())
	}()
	<-stop

	//The apps will go here only if there STOP signal like ctrl+c or log.Fatal
	log.Print("Shutting down server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)

	log.Print("Server gracefully stopped")
}
