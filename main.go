package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tizeen/go-restful-example/config"
	"github.com/Tizeen/go-restful-example/model"
	v "github.com/Tizeen/go-restful-example/pkg/version"
	"github.com/Tizeen/go-restful-example/router"
	"github.com/Tizeen/go-restful-example/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

var (
	// 定义命令行参数，返回指针
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	// 设置gin的运行模式
	gin.SetMode(viper.GetString("runmode"))

	//for {
	//	log.Info("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	//	time.Sleep(100 * time.Millisecond)
	//}

	// 创建一个不包含任何中间件的router
	g := gin.New()

	// 加载路由
	router.Load(
		g,

		// 将已有的切片传入可变参数
		middleware.Logging(),
		middleware.RequestId(),
	)

	// 通过goroutine自检服务是否成功启动
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	log.Info(cert)
	log.Info(key)
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	// 启动服务
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
