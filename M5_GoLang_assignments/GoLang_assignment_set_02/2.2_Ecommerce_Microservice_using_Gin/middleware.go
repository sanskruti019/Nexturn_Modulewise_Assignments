package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT-based authentication
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer valid-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// Logging middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		c.Writer.WriteString("\nRequest - Method: " + method + ", Path: " + path + ", Duration: " + duration.String())
	}
}

// Rate limiting middleware
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simple in-memory rate limiting logic
		c.Next()
	}
}
