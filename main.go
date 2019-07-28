package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	logPath := "req.log"
	httpPort := 9000

	// openLogFile(logPath)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", rootHandler)

	fmt.Printf("listening on %v\n", httpPort)
	fmt.Printf("Logging to %v\n", logPath)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1><div>Welcome to whereever you are</div>")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// func openLogFile(logfile string) {
// 	if logfile != "" {
// 		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

// 		if err != nil {
// 			log.Fatal("OpenLogfile: os.OpenFile:", err)
// 		}

// 		log.SetOutput(lf)
// 	}
// }
