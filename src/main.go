/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 * 
 * Filename: main.go
 * Author: ahaoozhang
 * Date: 2020-01-16 00:13:17 (Thursday)
 * Describe: 
 ******************************************************************/
package main

import (
        "github.com/gin-gonic/gin"
        log "github.com/sirupsen/logrus"
)

func main() {
    log.WithFields(log.Fields{
            "animal": "walrus",
        }).Info("A walrus appears")
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "Blog":"www.flysnow.org",
            "wechat":"flysnow_org",
        })
    })
   r.Run(":8080")
}
