package handlers

import (
	pp "github.com/Frontware/promptpay"
	"github.com/gin-gonic/gin"
)

type PromptPay struct {
	PromptPayID string  `json:"promptpay_id"`
	Amount      float64 `json:"amount"`
}

func CreatePromptPay(c *gin.Context) {
	var req PromptPay
	c.BindJSON(&req)
	promptPay := pp.PromptPay{
		PromptPayID: req.PromptPayID,
		Amount:      req.Amount,
	}

	QRcode, err := promptPay.Gen()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "promptpay_id or amount was wrong",
		})
	} else {
		c.JSON(200, gin.H{
			"PromptPay": QRcode,
		})
	}
}
