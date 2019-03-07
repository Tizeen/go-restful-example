package main

import (
	"github.com/Tizeen/go-restful-example/router"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

func main() {

	// 创建一个不包含任何中间件的router
	g := gin.New()

	// 定义一个空的gin.HandlerFunc切片
	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,

		// 将已有的切片传入可变参数
		middlewares...,
	)

	// 通过goroutine自检服务
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":9999")
	log.Printf(http.ListenAndServe(":9999", g).Error())
}

func pingServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://127.0.0.1:9999" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
