package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/Footters/hex-footters/pkg/auth/storage/redisdb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-redis/redis"
)

func main() {

	// Logging
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		// logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		// logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	// Business domain
	uRepo := redisdb.NewRedisUserRepository(redisConnection("redis:6379"))
	svc := auth.NewService(uRepo)
	svc = auth.NewLogginMiddleware(logger, svc)

	//Endpoint
	registerEndpoint := auth.MakeRegisterEndpoint(svc)
	loginEndpoint := auth.MakeLoginEndpoint(svc)

	// Transport
	mux := http.NewServeMux()

	registerHandler := httptransport.NewServer(
		registerEndpoint,
		auth.DecodeRegisterRequest,
		auth.EncodeResponse,
	)

	loginHandler := httptransport.NewServer(
		loginEndpoint,
		auth.DecodeLoginRequest,
		auth.EncodeResponse,
	)

	mux.Handle("/register", registerHandler)
	mux.Handle("/login", loginHandler)

	// Go!
	logger.Log("transport", "HTTP", "addr", ":8081")
	logger.Log("exit", http.ListenAndServe(":8081", mux))
}

func redisConnection(url string) *redis.Client {
	fmt.Println("Connecting to Redis DB")
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Ping().Err()

	if err != nil {
		panic(err)
	}
	return client
}
