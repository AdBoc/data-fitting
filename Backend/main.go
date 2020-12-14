package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type FitArgs struct {
	Arg1 []json.Number `json:"arg1"`
	Arg2 []json.Number `json:"arg2"`
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode) //ReleaseMode

	router.GET("/", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	router.POST("/", func(context *gin.Context) {
		var requestData FitArgs
		body, _ := ioutil.ReadAll(context.Request.Body)
		if err := json.Unmarshal(body, &requestData); err != nil {
			panic(err)
		}

		var arg1, arg2 string
		for _, s := range requestData.Arg1 {
			arg1 += s.String() + ","
		}
		for _, s := range requestData.Arg2 {
			arg2 += s.String() + ","
		}
		arg1 = strings.TrimSuffix(arg1, ",")
		arg2 = strings.TrimSuffix(arg2, ",")

		c := exec.Command("C:\\Users\\adria\\AppData\\Local\\Programs\\Python\\Python39\\python.exe", "../PyScript/main.py", arg1, arg2)
		stderr := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		c.Stderr = stderr
		c.Stdout = stdout
		if err := c.Run(); err != nil {
			fmt.Println("Error: ", err, "|", stderr.String())
		}

		context.JSON(http.StatusOK, gin.H{
			"xData": "xData",
			"yData": "yData",
			"response": stdout.String(),
		})
	})

	if err := router.Run(":8000"); err != nil {
		//os.Exit(0) //should it be here?
		panic(err)
	}
}