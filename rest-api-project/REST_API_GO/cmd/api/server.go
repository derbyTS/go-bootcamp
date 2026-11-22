package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// "time"

	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/pkg/utils"

	"github.com/joho/godotenv"
)

//go:embed .env
var envFile embed.FS

func loadEnvFromEmbeddedFile() {
	content, err := envFile.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Create a temp file to load env vars
	tempfile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tempfile.Name())

	// Write content of the embedded .env file to the file
	_, err = tempfile.Write(content)
	if err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}

	err = tempfile.Close()
	if err != nil {
		log.Fatalf("Error closing temp file: %v", err)
	}

	// Load env variables
	err = godotenv.Load(tempfile.Name())
	if err != nil {
		log.Fatalf("Error loading temp file: %v", err)
	}
}

func main() {
	// This use in development only
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Load the varialbes form embedded .env file
	loadEnvFromEmbeddedFile()

	fmt.Println("Environment variable CERT_FILE:", os.Getenv("CERT_FILE"))

	port := os.Getenv("SERVER_PORT")

	cert := os.Getenv("CERT_FILE")
	key := os.Getenv("KEY_FILE")

	tlsConfig := &tls.Config{
		// MinVersion: tls.VersionTLS12,
		MinVersion: tls.VersionTLS10,
	}

	router := router.MainRouter()

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOPtions := mw.HppOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www.form-urlencoded",
		Whitelist: []string{
			"sortBy",
			"sortOrder",
			"name",
			"age",
			"class",
			"page",
			"limit",
		},
	}

	// secureMux := mw.Cors(
	// 	rl.Throttle(
	// 		mw.ResponseTimeMiddleware(
	// 			mw.SecurityHeaders(
	// 				mw.Compression(
	// 					mw.Hpp(hppOPtions)(mux)))),
	// 	))

	jwtMiddleware := mw.MiddlwaresExcludePaths(
		mw.JWTMiddleware,
		"/execs/login",
		"/execs/forgotpassword",
		"/execs/resetpassword/reset/",
	)
	secureMux := utils.ApplyMiddlewares(
		router,
		mw.SecurityHeaders,
		mw.Compression,
		mw.Hpp(hppOPtions),
		mw.XSSMiddleware,
		jwtMiddleware,
		mw.ResponseTimeMiddleware,
		rl.Throttle,
		mw.Cors,
	)
	// secureMux := jwtMiddleware(mw.SecurityHeaders(router))
	// secureMux := mw.XSSMiddleware(router)

	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err)
	}
}
