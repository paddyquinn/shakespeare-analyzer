package router

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/shakespeare-analyzer/analyzer"
)

const link = "link"

func Run() {
  router := gin.Default()
  router.StaticFile("/", "html/index.html")
  router.StaticFile("/index.html", "html/index.html")
  router.GET("/analyze", func(c *gin.Context) {
    lnk := c.Query(link)
    a := analyzer.NewAnalyzer()
    characters, err := a.Analyze(lnk)
    if err != nil {
      c.JSON(http.StatusBadRequest, err.Error())
      return
    }
    c.JSON(http.StatusOK, characters)
  })
  router.Run()
}