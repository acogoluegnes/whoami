package main
import (
	"net/http"
	"fmt"
	"log"
	"os"
	"net"
	"sort"
)


func main() {
	http.HandleFunc("/",whoami)
	http.HandleFunc("/env",env)
	fmt.Println("Start listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000",nil))
}

func whoami(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintln(w,"hostname:",hostname)
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Fprintln(w, "IP:", ip)
		}
	}
	r.Write(w)
}

func env(w http.ResponseWriter, r *http.Request) {
	env := os.Environ()
	sort.Strings(env)
	for _, e := range env {
		fmt.Fprintln(w,e)
	}
}
