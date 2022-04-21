package standart

import (
	"fmt"
	"lesson11/Lookup"
	"net/http"
	"net/http/pprof"
	"runtime"
	"sync/atomic"
	"time"
)

var count int64

func main() {
	/*
		arg := os.Args
		host := arg[1]
		records, err := Lookup.LookupDnsRecords(host)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, ip := range records {
			fmt.Println(ip)
		}
	*/
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	PORT := ":8888"

	//http.HandleFunc("/", defaultHandler)
	//http.HandleFunc("/resolver", resolverHandler)
	//err := http.ListenAndServe(PORT, nil)
	r := http.NewServeMux()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/resolver", resolverHandler)

	r.HandleFunc("/debug/", pprof.Index)
	r.HandleFunc("/debug/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/profile", pprof.Profile)
	r.HandleFunc("/debug/symbol", pprof.Symbol)
	r.HandleFunc("/debug/trace", pprof.Trace)

	server := &http.Server{
		Addr:         PORT,
		Handler:      r,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	//err := http.ListenAndServe(PORT, r)
	//err := server.ListenAndServe()
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println(err)
	}
}

func resolverHandler(w http.ResponseWriter, r *http.Request) {
	records, err := Lookup.LookupDnsRecords("stemply.ru")
	if err != nil {
		w.WriteHeader(404)
		return
	}

	for _, record := range records {
		fmt.Fprintf(w, "<p>%s</p>", record)
		fmt.Println(record)
	}
	temp := atomic.LoadInt64(&count)
	fmt.Fprintf(w, "Count: %v", temp)
	fmt.Println("Count: ", temp)
	return
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Printf("Request from client %s\n", r.RemoteAddr)
	atomic.AddInt64(&count, 1)
}
