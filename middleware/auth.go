package middleware

import (
	"encoding/json"
	. "github.com/ash822/goweb/entity"
	"net/http"
)

const ApiKey = "topgun"

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resw http.ResponseWriter, req *http.Request) {
		resw.Header().Set("Content-type", "application/json")

		addCorsHeader(resw)

		if req.Method == "OPTIONS" {
			resw.WriteHeader(http.StatusOK)
			return
		}

		apiKey := req.Header.Get("X-API-Key")

		if apiKey != ApiKey {
			rejectRequestHandler(resw)
		} else {
			next.ServeHTTP(resw, req)
		}
	})
}

func rejectRequestHandler(resw http.ResponseWriter) {
	resw.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(resw).Encode(ServiceError{Error: "unable to authenticate, API-Key provided is invalid"})
}

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, X-API-Key")
	headers.Add("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")
}
