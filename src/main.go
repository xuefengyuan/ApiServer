package main

import (
    "ApiServer/src/router"
    "errors"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

func main()  {
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

    log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
    log.Printf(http.ListenAndServe(":8080", g).Error())

}
// pingServer pings the http server to make sure the router is working.
// 服务器健康自检函数
func pingServer() error {
    for i := 0; i < 2; i++ {
        // Ping the server by sending a GET request to `/health`.
        resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
        if err == nil && resp.StatusCode == 200 {
            return nil
        }

        // Sleep for a second to continue the next ping.
        log.Print("Waiting for the router, retry in 1 second.")
        time.Sleep(time.Second)
    }
    return errors.New("Cannot connect to the router.")
}