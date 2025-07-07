package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateeusferro/fineval/internal/auditly"
	"github.com/mateeusferro/fineval/internal/domain"
	"github.com/mateeusferro/fineval/internal/usecase"
)

func Routes(router *gin.Engine) {
	router.GET("/ping", serverCheck)
	router.POST("/evaluate", handleEvaluate)
}

func serverCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "pong")
}

func handleEvaluate(context *gin.Context) {
	evaluationInput := domain.EvaluationInput{}
	err := context.ShouldBindJSON(&evaluationInput)

	if err != nil {
		log.Fatalf("Incorrect input: %v", err)
	}

	result := usecase.Evaluate(evaluationInput)

	auditData := domain.AuditlyInput{
		Actor:    "fineval",
		Action:   "READ_RESOURCE",
		Resource: "evaluate",
		Metadata: map[string]interface{}{
			"score":   result.Score,
			"summary": result.Summary,
		},
	}

	auditly.Send(auditData)

	context.JSON(http.StatusOK, result)
}
