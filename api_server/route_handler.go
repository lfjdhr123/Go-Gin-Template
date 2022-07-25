package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"template/api_server/request"
	"template/api_server/response"
)

func loadRouterHandler(router *gin.Engine, routerHandler *RouterHandler) {
	router.POST("/samples", routerHandler.CreateSample)
	router.GET("/samples/:name", routerHandler.GetSampleByName)
	router.GET("/samples", routerHandler.GetSampleList)
}

func (routerHandler *RouterHandler) CreateSample(c *gin.Context) {
	var sampleInfo request.Sample
	if err := c.BindJSON(&sampleInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
		return
	}
	if err := routerHandler.controller.CreateSample(sampleInfo.ToController()); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("sample create failed due to: %v", err))
		return
	}
	c.AbortWithStatus(http.StatusOK)
}

func (routerHandler *RouterHandler) GetSampleByName(c *gin.Context) {
	name := c.Param("name")
	sample, err := routerHandler.controller.GetSampleByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("fail to get sample by name [%s] due to: %v", name, err))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, response.InitSample(sample))
}

func (routerHandler *RouterHandler) GetSampleList(c *gin.Context) {
	samples, err := routerHandler.controller.GetSampleList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("fail to get sample list due to: %v", err))
		return
	}
	res := make([]*response.Sample, 0)

	for _, sample := range samples {
		res = append(res, response.InitSample(sample))
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}
