package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Home page
func home(context *gin.Context) {
	message :=
		"This is an integer calculator using Rest API in Golang\n" +
			"The supported operations are add, sub, mul, div\n" +
			"Query examples:\n" +
			"\tlocalhost:9091/calculate?operation=add&value1=1&value2=9\t Result: 1 + 9 = 10\n" +
			"\tlocalhost:9091/calculate?operation=mul&value1=1&value2=-3\t Result: 1 * (-3) = -3\n" +
			"You can also create a request using terminal\n"

	context.String(http.StatusOK, message)
}

// Calculation page
func calculate(context *gin.Context) {
	// Getting parameters from GET
	operation := context.Query("operation")
	value1, err := strconv.Atoi(context.Query("value1"))
	value2, err := strconv.Atoi(context.Query("value2"))
	var result int

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "value1 and value2 should be integers"})
		return
	}

	// Doing operation
	var strOp string
	switch operation {
	case "add":
		result = value1 + value2
		strOp = "+"
	case "sub":
		result = value1 - value2
		strOp = "-"
	case "mul":
		result = value1 * value2
		strOp = "*"
	case "div":
		if value2 != 0 {
			result = value1 / value2
		}
		strOp = "/"
	default:
		context.String(http.StatusBadRequest, "Invalid operation")
		return
	}

	format := "%v %v %v = %v"
	if value2 < 0 {
		format = "%v %v (%v) = %v"
	}

	show := fmt.Sprintf(format, value1, strOp, value2, result)
	context.String(http.StatusCreated, show)
}

func main() {
	router := gin.Default()

	router.GET("/", home)

	router.GET("/calculate", calculate)

	err := router.Run("localhost:9091")
	if err != nil {
		return
	}
}
