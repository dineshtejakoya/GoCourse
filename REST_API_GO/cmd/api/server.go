package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	_, err = sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println("Error----: ", err)
		return
	}

	port := os.Getenv("API_PORT")
	// port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	//for tls we create a custom server, for creating custom server we use tlsconfig
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//Rate Limiter
	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	//secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddlware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddlware, rl.Middleware, mw.Cors)
	//secureMux := mw.SecurityHeaders(mux)
	router := router.Router()
	secureMux := mw.SecurityHeaders(router)
	//Create custom server
	server := &http.Server{
		Addr:    port,
		Handler: secureMux,
		//mw.Hpp(hppOptions)(rl.Middleware(mw.Compression(mw.ResponseTimeMiddlware(mw.SecurityHeaders(mw.Cors(mux)))))),
		//Handler: middlewares.Cors(mux),
		//Handler: mw.SecurityHeaders(mux),
		//Handler:   mux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port", port)
	err = server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
