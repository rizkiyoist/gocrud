package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkiyoist/gocrud/app/config"
	"github.com/rizkiyoist/gocrud/library/middleware"
	"github.com/rizkiyoist/gocrud/provider"
)

func main() {
	var err error
	roles := make(map[string][]string)
	roles[config.ADMIN] = []string{config.ADMIN}

	middleware.InitRole(roles)

	gin.SetMode(gin.DebugMode)

	p = provider.BuildAll()

	if err != nil {
		return
	}
}
