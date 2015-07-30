package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		bind      = fs.String("bind", "0.0.0.0:9999", "IP address and port to bind to")
		resp_code = fs.Int("response", 200, "HTTP response code")
	)
	flag.Usage = fs.Usage
	fs.Parse(os.Args[1:])

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("********************************************************************************")
		log.Printf("%s %s %s\n", req.Proto, req.Method, req.RequestURI)
		for header, values := range req.Header {
			for _, value := range values {
				log.Printf("%s: %s\n", header, value)
			}
		}
		if body, err := ioutil.ReadAll(req.Body); err == nil && len(body) > 0 {
			log.Println("")
			log.Printf("%s\n", string(body))
		}
		http.Error(w, ``, *resp_code)
	})

	log.Printf("Listening on %s", *bind)
	err := http.ListenAndServe(*bind, nil)
	if err != nil {
		log.Fatal(err)
	}
}
