package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//only supports linux for now

type NewTransactionRequest struct {
	Amount float64 `json:"amount"`
}
type NewTransactionResponse struct {
	UUID    string  `json:"uuid"`
	Amount  float64 `json:"amount"`
	Address string  `json:"address"`
}

// Callback for /new gin endpoint, takes a gin context.
// The endpoint expects a JSON body with the following format: {'amount': float},
// and returns a JSON body with the following format: {'uuid': string', amount': float, 'address': string}
func postNew(c *gin.Context) {
	var newTransactionRequest NewTransactionRequest
	if err := c.ShouldBindJSON(&newTransactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionResponse := NewTransactionResponse{
		UUID:    "1258598732",
		Amount:  newTransactionRequest.Amount,
		Address: "2345egf34t897yfv897g3278gf8g73487g",
	}
	c.JSON(http.StatusOK, transactionResponse)
}

func main() {
	router := gin.Default()
	router.POST("/api/transactions/new", postNew)

	router.Run("localhost:8080")
}
