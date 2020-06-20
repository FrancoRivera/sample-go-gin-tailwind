package internal

import ( "github.com/gin-gonic/gin")

var Router *gin.Engine

func Initialize() {
	Router = gin.Default()
	Router.LoadHTMLGlob("./views/templates/*")
}
