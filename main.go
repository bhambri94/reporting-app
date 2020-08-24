package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bhambri94/reporting-app/configs"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()
	sugar     = logger.Sugar()
)

func main() {
	configs.SetConfig()
	sugar.Infof("starting reporting app server......")
	defer logger.Sync() // flushes buffer, if any

	router := fasthttprouter.New()
	router.POST(configs.Configurations.AddResultsXMLPath+"/user=:userName/password=:password", addReportingResultsXML)

	log.Fatal(fasthttp.ListenAndServe(":8010", router.Handler))
}

func addReportingResultsXML(ctx *fasthttp.RequestCtx) {
	sugar.Infof("creating xml allure report......")
	ctx.Response.Header.Set("Content-Type", "application/json")
	userName := ctx.UserValue("userName")
	password := ctx.UserValue("password")
	if userName == configs.Configurations.AppUsername && password == configs.Configurations.AppPassword {
		ctx.Response.SetStatusCode(201)
		a := ctx.Request.Body()
		f, err := os.Create("allure-results/" + generateUUID() + "-testsuite.xml")
		if err != nil {
			fmt.Println(err)
			ctx.Response.SetStatusCode(500)
			successResponse := "{\"success\":false,\"response\":\"Unable to add report\"}"
			ctx.Write([]byte(successResponse))
		}
		fmt.Println(string(a))
		fmt.Fprintln(f, string(a))
		defer f.Close()
		out, err := exec.Command("allure", "generate", "allure-results", "--clean", "-o", "allure-report").Output()
		fmt.Println(string(out))
		if err != nil {
			fmt.Println(err)
			ctx.Response.SetStatusCode(500)
			successResponse := "{\"success\":false,\"response\":\"Unable to generate new report\"}"
			ctx.Write([]byte(successResponse))
			return
		}
		successResponse := "{\"success\":true,\"response\":\"Added results to Allure reporter\"}"
		ctx.Write([]byte(successResponse))
	} else {
		ctx.Response.SetStatusCode(401)
		successResponse := "{\"success\":false,\"response\":\"Unauthorized\"}"
		ctx.Write([]byte(successResponse))
	}

}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
