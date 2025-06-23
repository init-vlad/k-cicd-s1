// ping.go
package main

import "net/http"

func main() {
	// clientID := os.Getenv("CLIENT_ID")
	// if clientID == "" {
	// 	clientID = "unknown"
	// }
	// for {
	// 	log.Printf("PING 2.0 %s", clientID)
	// 	time.Sleep(10 * time.Second)
	// }

	serv := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service 1 alive: ci cd edition:))"))
	})

	http.ListenAndServe(":8080", serv)
}
