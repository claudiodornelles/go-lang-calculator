package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

type operation struct {
	FirstValue  float64 `json:"first_value"`
	SecondValue float64 `json:"second_value"`
	Type        string  `json:"type"`
	Result      float64 `json:"result"`
}

var operationHistory []operation

func getHistory(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, operationHistory)
}

func calculate(context *gin.Context) {
	firstValue, firstConversionError := strconv.ParseFloat(context.Param("firstValue"), 64)
	secondValue, secondConversionError := strconv.ParseFloat(context.Param("secondValue"), 64)

	if firstConversionError != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "cannot convert firstValue to a number"})
	} else if secondConversionError != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "cannot convert secondValue to a number"})
	} else {
		switch strings.ToLower(context.Param("operation")) {
		case "sum":
			context.IndentedJSON(http.StatusOK, sum(firstValue, secondValue))
		case "sub":
			context.IndentedJSON(http.StatusOK, sub(firstValue, secondValue))
		case "mul":
			context.IndentedJSON(http.StatusOK, mul(firstValue, secondValue))
		case "div":
			result, err := div(firstValue, secondValue)
			if err != nil {
				context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			} else {
				context.IndentedJSON(http.StatusOK, result)
			}
		default:
			context.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"Error": "Operation '" + context.Param("operation") + "' not allowed."})
		}
	}
}

func sum(firstValue float64, secondValue float64) operation {
	var newOperation operation
	newOperation.Type = "sum"
	newOperation.FirstValue = firstValue
	newOperation.SecondValue = secondValue
	newOperation.Result = firstValue + secondValue

	operationHistory = append(operationHistory, newOperation)

	return newOperation
}

func sub(firstValue float64, secondValue float64) operation {
	var newOperation operation
	newOperation.Type = "sub"
	newOperation.FirstValue = firstValue
	newOperation.SecondValue = secondValue
	newOperation.Result = firstValue - secondValue

	operationHistory = append(operationHistory, newOperation)

	return newOperation
}

func mul(firstValue float64, secondValue float64) operation {
	var newOperation operation
	newOperation.Type = "mul"
	newOperation.FirstValue = firstValue
	newOperation.SecondValue = secondValue
	newOperation.Result = firstValue * secondValue

	operationHistory = append(operationHistory, newOperation)

	return newOperation
}

func div(firstValue float64, secondValue float64) (operation, error) {

	if secondValue == 0 {
		var nullOperation operation
		return nullOperation, errors.New("math: cannot divide number by zero")
	} else {
		var newOperation operation
		newOperation.Type = "div"
		newOperation.FirstValue = firstValue
		newOperation.SecondValue = secondValue
		newOperation.Result = firstValue / secondValue
		operationHistory = append(operationHistory, newOperation)
		return newOperation, nil
	}
}

func health(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"status": "up"})
}

func main() {
	router := gin.Default()
	router.GET("/calc/history", getHistory)
	router.GET("/calc/:operation/:firstValue/:secondValue", calculate)
	router.GET("/health", health)
	err := router.Run("0.0.0.0:8090")
	if err != nil {
		logger.Println(err.Error())
		return
	}
}
