package util

import (
	"log"
	"net/http"
	"net/http/pprof"
	cpuprof "runtime/pprof"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	p := cpuprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func ProfilingMemory(memPort string) {
	if len(memPort) > 0 {
		s := []string{"localhost", ":", memPort}
		h := strings.Join(s, "")
		http.HandleFunc("/", handler)
		go func() {
			log.Println(http.ListenAndServe(h, nil))
		}()
	}
}

func ProfilingCPU(cpuPort string) {
	if len(cpuPort) > 0 {
		r := http.NewServeMux()

		// Register pprof handlers
		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)

		s := []string{"localhost", ":", cpuPort}
		h := strings.Join(s, "")
		go func() {
			log.Println(http.ListenAndServe(h, r))
		}()
	}
}
