package handlers

import (
	cntx "api/pkg/context"
	"net/http"
	"time"
)

var nonFunctionalPaths = []string{
	"/",
}

func withRequestedAt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		if contains(nonFunctionalPaths, req.URL.Path) {
			next.ServeHTTP(wr, req)
			return
		}
		ctx := cntx.WithRequestedAt(req.Context(), time.Now())
		req = req.WithContext(ctx)
		next.ServeHTTP(wr, req)
	})
}

func contains(arr []string, item string) bool {
	for _, arrItem := range arr {
		if item == arrItem {
			return true
		}
	}

	return false
}

// func isUserAuthorised(userID, scopes, flow)bool{
// this api will check the Authorization Bearer token and is valid or not
// if valid then further check the scope of the user is eligible to make particular api call
// if not return false else return true
//}
