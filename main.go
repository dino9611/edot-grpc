package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tes-edot/cmd/grpcclient"
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

	ctx.JSON(200, gin.H{"message": "success", "data": result})
}

func main() {
	r := gin.Default()

	chartService := grpcclient.CharService()
	handler := &Handler{
		chartService: chartService,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/chart", handler.GetChart)
	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}
