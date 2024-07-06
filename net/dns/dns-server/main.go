package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
    m := new(dns.Msg)
    m.SetReply(r)
    m.Authoritative = true
    domain := m.Question[0].Name
    switch r.Question[0].Qtype {
    case dns.TypeA:
        rr, err := dns.NewRR(fmt.Sprintf("%s A 127.0.0.1", domain))
        if err == nil {
            m.Answer = append(m.Answer, rr)
        }
    }
    w.WriteMsg(m)
}

func main() {
    // attach request handler func
    dns.HandleFunc(".", handleDNS)

    // start server
    port := "53"
    server := &dns.Server{Addr: ":" + port, Net: "udp"}
    log.Printf("Starting at %s\n", port)
    err := server.ListenAndServe()
    defer server.Shutdown()
    if err != nil {
        log.Fatalf("Failed to start server: %s\n ", err.Error())
    }
}
