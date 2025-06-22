package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type EchoResponse struct {
	Method     string              `json:"method"`
	URL        string              `json:"url"`
	Headers    map[string][]string `json:"headers"`
	Body       interface{}         `json:"body"`
	RemoteAddr string              `json:"remote_addr"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var bodyJson interface{}
		if len(body) > 0 {
			if err := json.Unmarshal(body, &bodyJson); err != nil {
				bodyJson = fmt.Sprintf("%v", string(body))
			}
		}

		response := EchoResponse{
			Method:     r.Method,
			URL:        r.URL.String(),
			Headers:    r.Header,
			Body:       bodyJson,
			RemoteAddr: r.RemoteAddr,
		}

		log.Println(response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	port := ":8080"
	fmt.Printf("Echo server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
