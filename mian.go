package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	port := "8080"

	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("LogFile: %v", err)
	}
	log.SetOutput(logFile)
	log.Printf("Date: %s", time.Now().Format(time.RFC1123))
	log.Printf("Author: Dzmitry Markevich")
	log.Printf("Listening: %s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Ip error", http.StatusInternalServerError)
			log.Printf("Error ip define: %v", err)
			return
		}

		loc, err := time.LoadLocation("Local")
		if err != nil {
			loc = time.UTC
		}
		timeInZone := time.Now().In(loc)

		response := fmt.Sprintf(
			"<html><body><h1>Info</h1><p>Ip: %s</p><p>Date: %s</p></body></html>",
			clientIP,
			timeInZone.Format("2006-01-02 15:04:05 MST"),
		)

		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(response))
	})

	log.Printf("Port: %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
