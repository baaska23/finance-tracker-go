// package middleware

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	"secret_room_backend_v2/internal/platform/types"

// 	"github.com/gin-gonic/gin"
// )

// func DeviceAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		basicResponse := types.NewAPIResponse()

// 		macAddress := c.GetHeader("X-Device-MAC")
// 		serialNumber := c.GetHeader("X-Device-Serial")
// 		clientSignature := c.GetHeader("X-Device-Signature")
// 		timestamp := c.GetHeader("X-Timestamp")

// 		if macAddress == "" || serialNumber == "" || clientSignature == "" || timestamp == "" {
// 			basicResponse.Message = "Missing authentication headers"
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, basicResponse)
// 			return
// 		}

// 		if !validateTimestamp(timestamp) {
// 			basicResponse.Message = "Invalid or expired timestamp"
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, basicResponse)
// 			return
// 		}

// 		expectedSignature := generateDeviceSignature(macAddress, serialNumber, timestamp)

// 		if clientSignature != expectedSignature {
// 			basicResponse.Message = "Invalid device signature"
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, basicResponse)
// 			return
// 		}

// 		c.Set("device_mac", macAddress)
// 		c.Set("device_serial", serialNumber)

// 		c.Next()
// 	}
// }

// func generateDeviceSignature(mac, serial, timestamp string) string {
// 	secret := os.Getenv("DEVICE_SECRET_KEY")
// 	message := fmt.Sprintf("%s:%s:%s:%s", mac, serial, timestamp, secret)

// 	hash := sha256.Sum256([]byte(message))
// 	return hex.EncodeToString(hash[:])
// }

// func validateTimestamp(timestamp string) bool {
// 	t, err := time.Parse("2006-01-02 15", timestamp)
// 	if err != nil {
// 		return false
// 	}

// 	now := time.Now().UTC()
// 	currentHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)

// 	diff := currentHour.Sub(t).Hours()
// 	return diff >= -1 && diff <= 1
// }
