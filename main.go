package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	Router(r)

	log.Println("Server started")
	r.Run()
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
}

func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
