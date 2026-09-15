package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	// "time"

	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repositories/sqlconnect"
)

func main() {
	_, err := sqlconnect.ConnectDB("dbeaver_testdb")
	if err != nil {
		fmt.Println(err)
		return
	}
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	router := router.Router()

	// rl := mw.NewRateLimiter(5, time.Minute)
	//
	// hppOPtions := mw.HppOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www.form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(
	// 	rl.Throttle(
	// 		mw.ResponseTimeMiddleware(
	// 			mw.SecurityHeaders(
	// 				mw.Compression(
	// 					mw.Hpp(hppOPtions)(mux)))),
	// 	))
	// secureMux := utils.ApplyMiddlewares(
	// 	mux,
	// 	mw.Hpp(hppOPtions),
	// 	mw.Compression,
	// 	mw.SecurityHeaders,
	// 	mw.ResponseTimeMiddleware,
	// 	rl.Throttle,
	// 	mw.Cors,
	// )

	secureMux := mw.SecurityHeaders(router)

	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on: ", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err)
	}
}
