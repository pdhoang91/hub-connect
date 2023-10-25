package http

import (
	"encoding/base64"
	"fmt"
	"hub-connect/internal/delivery/http/model"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	ErrInvalidHeader = "invalid header format"
	ErrNoCredentials = "need to include Authorization in the header"
	ErrInvalidCreds  = "invalid username and password"
)

// Define a map to store username-password pairs
var credentialsMap = map[string]string{
	"pdhoang91@gmail.com": "cubicasa",
}

func ConfigureCORS(r *gin.Engine) {

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// Configure CORS
	r.Use(cors.Default()) // Use the default CORS configuration
}

func AuthMiddleware(ctx *gin.Context) {
	// Check if the user is authenticated (you can use a token, session, or any other method here)

	if isAuthenticated, err := isAuthenticated(ctx); err != nil {
		response := model.ErrorResponse{
			Status:  "error",
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusUnauthorized, response)
		ctx.Abort()
		return
	} else {
		if !isAuthenticated {
			response := model.ErrorResponse{
				Status:  "error",
				Code:    http.StatusUnauthorized,
				Message: ErrInvalidCreds,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}
	}

	ctx.Next()
}
func isAuthenticated(ctx *gin.Context) (bool, error) {
	authHeader := ctx.GetHeader("Authorization")

	//Actually, here we can use gin's BasicAuth. But I found that gin's parseBasicAuth function returned errors that were not detailed,
	//So I wrote a separate code to catch more detailed errors.
	//ctx.Request.BasicAuth()
	//parseBasicAuth

	// Check if the header is empty
	if authHeader == "" {
		// No credentials provided
		return false, fmt.Errorf(ErrNoCredentials)
	}

	// Extract and decode credentials
	username, password, err := decodeCredentials(authHeader)
	if err != nil {
		// Error decoding credentials
		return false, err
	}

	// Verify the username and password
	return checkCredentials(username, password), nil
}

func decodeCredentials(authHeader string) (string, string, error) {
	// The header should have a format like "Basic base64-encoded-credentials"
	// We need to extract the base64-encoded-credentials part and decode it
	credentials := strings.Split(authHeader, " ")
	if len(credentials) != 2 || credentials[0] != "Basic" {
		// Invalid header format
		return "", "", fmt.Errorf(ErrInvalidHeader)
	}

	// Decode the base64-encoded credentials
	decoded, err := base64.StdEncoding.DecodeString(credentials[1])
	if err != nil {
		// Error decoding credentials
		return "", "", err
	}

	decodedStr := string(decoded)
	parts := strings.Split(decodedStr, ":")
	if len(parts) != 2 {
		// Invalid credentials format
		return "", "", fmt.Errorf(ErrInvalidHeader)
	}

	return parts[0], parts[1], nil
}

func checkCredentials(username, password string) bool {
	//username:pdhoang91@gmail.com
	//password:cubicasa
	//Basic Authentication: Basic cGRob2FuZzkxQGdtYWlsLmNvbTpjdWJpY2FzYQ==

	if storedPassword, ok := credentialsMap[username]; ok {
		return storedPassword == password
	}
	return false
}
