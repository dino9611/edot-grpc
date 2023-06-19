package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tes-edot/gateway/grpcclient"
	charpb "github.com/tes-edot/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	chartService charpb.ChartClient
}

func (h *Handler) GetChart(ctx *gin.Context) {
	result, err := h.chartService.GetChart(ctx, &emptypb.Empty{})
	fmt.Println(result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "err": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success", "data": result.Data})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	chartService := grpcclient.CharService()
	handler := &Handler{
		chartService: chartService,
	}
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/chart", handler.GetChart)
	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}
