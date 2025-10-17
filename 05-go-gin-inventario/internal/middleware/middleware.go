package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// CORS configura middleware de CORS
func CORS() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

// Logger configura middleware de logging personalizado
func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output:    gin.DefaultWriter,
		SkipPaths: []string{"/health", "/metrics"},
	})
}

// RequestID adiciona um ID único para cada requisição
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		
		c.Header("X-Request-ID", requestID)
		c.Set("RequestID", requestID)
		c.Next()
	}
}

// Recovery configura middleware de recuperação de panic
func Recovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Printf("Panic recovered: %s", err)
		}
		
		c.JSON(500, gin.H{
			"erro":      "Erro interno do servidor",
			"codigo":    "INTERNAL_SERVER_ERROR",
			"timestamp": time.Now(),
		})
	})
}

// RateLimiter implementa um rate limiter simples (em produção usaria Redis)
func RateLimiter() gin.HandlerFunc {
	type client struct {
		count     int
		lastReset time.Time
	}
	
	clients := make(map[string]*client)
	const maxRequests = 100 // 100 requests por minuto
	
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()
		
		if clientData, exists := clients[clientIP]; exists {
			// Reset contador se passou 1 minuto
			if now.Sub(clientData.lastReset) > time.Minute {
				clientData.count = 0
				clientData.lastReset = now
			}
			
			if clientData.count >= maxRequests {
				c.JSON(429, gin.H{
					"erro":      "Muitas requisições",
					"codigo":    "RATE_LIMIT_EXCEEDED",
					"timestamp": time.Now(),
				})
				c.Abort()
				return
			}
			
			clientData.count++
		} else {
			clients[clientIP] = &client{
				count:     1,
				lastReset: now,
			}
		}
		
		c.Next()
	}
}

// Security adiciona headers de segurança
func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	}
}