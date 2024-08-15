package middleware

import (
	"context"
	"fmt"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/web"
	"net/http"
	"strings"
)
type AuthMiddleware struct {
    Handler         http.Handler
    ExcludedRoutes  map[string]struct{} // Use a map for faster lookups
}

func NewAuthMiddleware(handler http.Handler, excludedRoutes []string) *AuthMiddleware {
    excludedRoutesMap := make(map[string]struct{}, len(excludedRoutes))
    for _, route := range excludedRoutes {
        excludedRoutesMap[route] = struct{}{}
    }
    return &AuthMiddleware{
        Handler:        handler,
        ExcludedRoutes: excludedRoutesMap,
    }
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    fmt.Printf("Requested path: %s\n", request.URL.Path)

    // Check if the path is in excluded routes
    if _, exists := middleware.ExcludedRoutes[request.URL.Path]; exists {
        fmt.Println("Path is in excluded routes")
        middleware.Handler.ServeHTTP(writer, request)
        return
    }

    // Authentication logic
    authHeader := request.Header.Get("Authorization")
    if authHeader == "" {
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusUnauthorized)
        webResponse := web.WebResponse{
            Code:   http.StatusUnauthorized,
            Status: "UNAUTHORIZED",
        }
        helper.WriteToResponseBody(writer, webResponse)
        return
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    claims, err := helper.ValidateJWT(tokenString)
    if err != nil {
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusUnauthorized)
        webResponse := web.WebResponse{
            Code:   http.StatusUnauthorized,
            Status: "UNAUTHORIZED",
        }
        helper.WriteToResponseBody(writer, webResponse)
        return
    }

    // Add user ID to context
    ctx := context.WithValue(request.Context(), "userID", claims.Issuer)
    request = request.WithContext(ctx)

    // Call the next handler
    middleware.Handler.ServeHTTP(writer, request)
}

// type AuthMiddleware struct {
// 	Handler http.Handler
// }

// func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
// 	return &AuthMiddleware{Handler: handler}
// }

// func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// if "RAHASIA" == request.Header.Get("X-API-Key") {
	// 	//ok
	// 	middleware.Handler.ServeHTTP(writer, request)
	// }else{
	// 	writer.Header().Set("Content-Type", "application/json")
	// 	writer.WriteHeader(http.StatusUnauthorized)
		
	// 	webResponse := web.WebResponse{
	// 		Code: http.StatusUnauthorized,
	// 		Status: "UNAUTHORIZED",
	// 	}
	// 	helper.WriteToResponseBody(writer, webResponse)
	// }

// }