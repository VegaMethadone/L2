package server

import (
	"calendar/internal/server/handlers"
	"calendar/internal/structs/conf"
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func NewServer(options *conf.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", handlers.Home())

	mux.HandleFunc("/create_event", handlers.NewEvent())
	mux.HandleFunc("/update_event", handlers.UpdateEvent())
	mux.HandleFunc("/delete_event", handlers.DeleteEvent())

	mux.HandleFunc("/events_for_day", handlers.GetEventDay())
	mux.HandleFunc("/events_for_week", handlers.GetEventWeek())
	mux.HandleFunc("/events_for_month", handlers.GetEventMonth())

	newMux := http.NewServeMux()
	newMux.Handle("/", mux)

	handler := LoggerMiddleware(newMux)
	server := &http.Server{
		Addr:         fmt.Sprintf("%s%s", options.Network.Address, options.Network.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(options.Network.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(options.Network.WriteTimeout) * time.Second,
	}

	return server
}
