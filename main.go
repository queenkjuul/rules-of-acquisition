package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
)

//go:embed ferengi.json
var f embed.FS
var unmapped int64 = -2
var pattern = regexp.MustCompile(`Acquisition #(\d+)`)

func main() {
	serve := flag.Bool("serve", false, "set to true to run a REST API server")
	address := flag.String("address", ":8080", "address to listen on (if in serve mode)")
	tls := flag.Bool("tls", false, "enable TLS (https), must have cert.pem and key.pem")
	cert := flag.String("cert", "cert.pem", "certificate file to use (if TLS enabled)")
	key := flag.String("key", "key.pem", "key file to use (if TLS enabled)")
	id := flag.Int64("id", -1, "number of the specific rule you'd like to retrieve")
	flag.Parse()

	data, err := f.ReadFile("ferengi.json")
	if err != nil {
		log.Fatal("Couldn't read embedded data!")
	}
	var rules []string
	err = json.Unmarshal([]byte(data), &rules)
	if err != nil {
		log.Fatal("Couldn't parse embedded data!")
	}

	// we want to build a map of Rule Number to Rule Text
	// array index doesn't map to Rule Number
	// Rule Number is contained within the text ("Rule of Acquisition #i")
	m := make(map[int64]string)
	for _, rule := range rules {
		match := pattern.FindStringSubmatch(rule)
		if match != nil {
			i, _ := strconv.ParseInt(match[1], 10, 0)
			m[i] = rule
		} else {
			m[unmapped] = rule
			unmapped--
		}
	}

	var i int64

	if !*serve {
		// no -serve flag, CLI mode
		if *id != -1 {
			i = *id
		} else {
			i = rand.Int63n(int64(len(m)))
		}
		fmt.Println(m[i])
	} else {
		// -serve flag, REST API mode
		mux := http.NewServeMux()

		mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			i = rand.Int63n(int64(len(m)))
			w.Write([]byte(rules[i]))
		})
		mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
			val, err := strconv.Atoi(r.PathValue("id"))
			if err != nil {
				http.Error(w, "Error processing provided ID", http.StatusInternalServerError)
				return
			}
			i = int64(val)
			if v, ok := m[i]; ok {
				w.Write([]byte(v))
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}
		})

		if *tls {
			fmt.Printf("HTTPS server listening on %v\n", *address)
			log.Fatal(http.ListenAndServeTLS(*address, *cert, *key, mux))
		} else {
			fmt.Printf("HTTP server listening on %v\n", *address)
			log.Fatal(http.ListenAndServe(*address, mux))
		}
	}
}
