package main

import (
	"fmt"
	"net/http"
	"os"
	"test-gokit/account"
	"test-gokit/middleware"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	fmt.Println("Start")
	r := mux.NewRouter()
	metricsMiddleware := middleware.NewMetricsMiddleware()
	db := account.GetDBconn()

	var svc account.AccountService
	// svc = accountservice{}
	{
		repository, err := account.NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = account.NewService(repository, logger)
	}
	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreateAccountHandler := httptransport.NewServer(
		account.MakeCreateCustomerEndpoint(svc),
		account.DecodeCreateCustomerRequest,
		account.EncodeResponse,
	)
	GetByCustomerIdHandler := httptransport.NewServer(
		account.MakeGetCustomerByIdEndpoint(svc),
		account.DecodeGetCustomerByIdRequest,
		account.EncodeResponse,
	)
	GetAllCustomersHandler := httptransport.NewServer(
		account.MakeGetAllCustomersEndpoint(svc),
		account.DecodeGetAllCustomersRequest,
		account.EncodeResponse,
	)
	DeleteCustomerHandler := httptransport.NewServer(
		account.MakeDeleteCustomerEndpoint(svc),
		account.DecodeDeleteCustomerRequest,
		account.EncodeResponse,
	)
	UpdateCustomerHandler := httptransport.NewServer(
		account.MakeUpdateCustomerendpoint(svc),
		account.DecodeUpdateCustomerRequest,
		account.EncodeResponse,
	)

	http.Handle("/", r)
	r.Handle("/metrics", promhttp.Handler())

	http.Handle("/account", CreateAccountHandler)
	http.Handle("/account/update", UpdateCustomerHandler)
	r.Handle("/account/getAll", GetAllCustomersHandler).Methods("GET")
	r.Handle("/account/{customerid}", GetByCustomerIdHandler).Methods("GET")
	r.Handle("/account/{customerid}", DeleteCustomerHandler).Methods("DELETE")
	r.Use(metricsMiddleware.Metrics)

	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
