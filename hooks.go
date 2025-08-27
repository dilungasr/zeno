package zeno

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

// IsPaymentCompleted checks if a hook data indicates a completed payment
func IsPaymentCompleted(data HookData) bool {
	return data.PaymentStatus == "COMPLETED"
}

// ParseHookHTTP parses the data received from the Zeno API's webhook using net/http
func ParseHookHTTP(w http.ResponseWriter, r *http.Request) (data HookData, ok bool) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		zLog(err.Error())
		return data, false
	}

	return data, IsPaymentCompleted(data)
}

// ParseHookGin parses the data received from the Zeno API's webhook using Gin
func ParseHookGin(c *gin.Context) (data HookData, ok bool) {
	if err := c.ShouldBindJSON(&data); err != nil {
		zLog(err.Error())
		return data, false
	}

	return data, IsPaymentCompleted(data)
}

// ParseHookEcho parses the data received from the Zeno API's webhook using Echo
func ParseHookEcho(c echo.Context) (data HookData, ok bool) {
	if err := c.Bind(&data); err != nil {
		zLog(err.Error())
		return data, false
	}

	return data, IsPaymentCompleted(data)
}

// ParseHookFiber parses the data received from the Zeno API's webhook using Fiber
func ParseHookFiber(c *fiber.Ctx) (data HookData, ok bool) {
	if err := c.BodyParser(&data); err != nil {
		zLog(err.Error())
		return data, false
	}

	return data, IsPaymentCompleted(data)
}
