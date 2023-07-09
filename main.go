package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// gin.Default().SetTrustedProxies([]string{"192.168.1.0/24"})
	r.Use(CountRequests)
	Router(r)

	log.Println("Server started")
	// Port
	r.Run(fmt.Sprintf(":%s", os.Getenv("EXPOSED_PORT")))
}

func Router(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		sess := session.Must(session.NewSession())

		svc := sts.New(sess)

		input := &sts.GetCallerIdentityInput{}
		result, err := svc.GetCallerIdentity(input)
		if err != nil {
			log.Println("Error retrieving caller identity:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve caller identity"})
			return
		}

		response := make(map[string]string)
		response["Account"] = getStringValue(result.Account)
		response["Arn"] = getStringValue(result.Arn)
		response["UserId"] = getStringValue(result.UserId)

		c.JSON(http.StatusOK, response)
	})

	r.GET("/counter", func(c *gin.Context) {
		response := map[string]int{
			"count": counter,
		}

		c.JSON(http.StatusOK, response)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/switch", func(c *gin.Context) {
		mutex.Lock()
		switchStatus = !switchStatus
		mutex.Unlock()

		c.JSON(http.StatusOK, gin.H{
			"status": switchStatus,
		})
	})

	r.GET("/liveness", func(c *gin.Context) {
		if switchStatus {
			c.String(http.StatusOK, "Liveness Probe")
		} else {
			c.String(http.StatusServiceUnavailable, "Liveness Probe")
		}
	})

	r.GET("/readiness", func(c *gin.Context) {
		if switchStatus {
			c.String(http.StatusOK, "Readiness Probe")
		} else {
			c.String(http.StatusServiceUnavailable, "Readiness Probe")
		}
	})
}

func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

var counter int
var counterMutex sync.Mutex

func CountRequests(c *gin.Context) {
	counterMutex.Lock()
	defer counterMutex.Unlock()

	counter++
	c.Next()
}

var (
	switchStatus bool = true
	mutex        sync.Mutex
)
