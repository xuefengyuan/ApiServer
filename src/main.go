package main

import (
    "ApiServer/src/config"
    "ApiServer/src/router"
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/lexkong/log"
    "github.com/spf13/pflag"
    "github.com/spf13/viper"
    "net/http"
    "time"
)

var (
    cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)


func main()  {

    pflag.Parse()

    // init config
    if err := config.Init(*cfg); err != nil {
        panic(err)
    }


    // Create the Gin engine.
    g := gin.New()

    middlewares := []gin.HandlerFunc{}

    // Routes.
    router.Load(
        // Cores.
        g,

        // Middlwares.
        middlewares...,
    )

    // 开启一个协程，启动服务器健康自检
    //go func() {
    //    if err := pingServer(); err != nil {
    //        log.Fatal("The router has no response, or it might took too long to start up.", err)
    //    }
    //    log.Print("The router has been deployed successfully.")
    //}()

    addr := viper.GetString("addr")

    log.Infof("Start to listening the incoming requests on http address: %s", addr)
    log.Info(http.ListenAndServe(addr, g).Error())

}
// pingServer pings the http server to make sure the router is working.
// 服务器健康自检函数
func pingServer() error {
    for i := 0; i < viper.GetInt("max_ping_count"); i++ {
        // Ping the server by sending a GET request to `/health`.
        resp, err := http.Get(viper.GetString("url")+ "/sd/health")
        if err == nil && resp.StatusCode == 200 {
            return nil
        }

        // Sleep for a second to continue the next ping.
        log.Info("Waiting for the router, retry in 1 second.")
        time.Sleep(time.Second)
    }
    return errors.New("Cannot connect to the router.")
}