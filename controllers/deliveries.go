package controllers

import (
	"go_test/pb"
	
	// "fmt"
	"net/http"
	"log"

	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
)

// GET /deliveries/invalid_combination
// Get all invalid delivery combination
func FindInvalidDeliveries(c *gin.Context) {
	// Connect to gcs service
	conn, err := grpc.Dial("gcs-service:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	gcsClient := pb.NewInvalidDeliveriesServiceClient(conn)

	// Call gcs service
	req := &pb.InvalidDeliveriesRequest{}
	if res, err := gcsClient.FindInvalidDeliveries(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": res.GetDeliveries(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}