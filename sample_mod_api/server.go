package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// func loadClientCAs() *x509.CertPool {
// 	clientCAs := x509.NewCertPool()
// 	caCert, err := os.ReadFile("cert.pem")
// 	if err != nil {
// 		log.Fatalln("Could not load client CA:", err)
// 	}
// 	clientCAs.AppendCertsFromPEM(caCert)
// 	return clientCAs
// }

func main() {

	//need to handle incoming requests at certain endpoints at certain routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	//Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	//Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// ClientAuth: tls.RequireAndVerifyClientCert,
		// ClientCAs:  loadClientCAs(),
	}

	//Configure a custom Server
	//note: &http.Server to create a server construct
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	//Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}

	//HTTP 1.1 without TLS
	// fmt.Println("Server is running on port:", port)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start server", err)
	// }

}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown Version"
	}
}

//http://127.0.0.1:3000/users
//http:localhost:3000/users
