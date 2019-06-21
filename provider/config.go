package provider

import (
	"github.com/KulinaID/mark-one/app/config"
	"github.com/KulinaID/mark-one/libraries/httpresponse"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// 
func NewBuildContainer() *dig.Container {
	c := dig.New()
	c = container.BuildConfigProvider(c)
	c = container.BuildRepositoryProvider(c)
	c = container.BuildUsecaseProvider(c)
	c = container.BuildHttpHandlerProvider(c)

	return c
}

// Builds all providers in container package
func BuildAll(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(func() []byte {
		return []byte("go-crud")
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func() config.Config {
		return config.NewConfig()
	}); err != nil {
		panic(err)
	}

	// if err = container.Provide(func(cfg config.Config) (error, *gorm.DB) {
	// 	return config.NewMysqlConfig(cfg)
	// }); err != nil {
	// 	panic(err)
	// }

	// if err = container.Provide(func(cfg config.Config) (error, *mgo.Database) {
	// 	return config.NewMongoDBConfig(cfg)
	// }); err != nil {
	// 	panic(err)
	// }

	if err = c.Provide(func() httpresponse.KulinaRequestReader {
		return httpresponse.NewKulinaRequestReader()
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func() *gin.Engine {
		return gin.Default()
	}); err != nil {
		panic(err)
	}

	return c
}
