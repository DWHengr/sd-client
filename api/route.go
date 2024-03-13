package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sd-client/config"
	pkglogger "sd-client/logger"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,aurora-token")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PATCH,PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}

type Router struct {
	c      *config.Config
	engine *gin.Engine
}

var routers = []func(engine *gin.Engine){
	ServiceDiscovery,
	ServiceDiscoveryApi,
}

func NewRouter(c *config.Config) (*Router, error) {
	engine, err := newRouter(c)
	engine.LoadHTMLGlob("templates/*")
	engine.Use(Cors())
	if err != nil {
		return nil, err
	}
	for _, f := range routers {
		f(engine)
	}
	return &Router{
		c:      c,
		engine: engine,
	}, nil
}

func newRouter(c *config.Config) (*gin.Engine, error) {

	engine := gin.New()
	engine.Use(pkglogger.GinLogger())

	return engine, nil
}

// Run router.
func (r *Router) Run() {
	r.engine.Run(r.c.Port)
}

// Close router.
func (r *Router) Close() {
}

func (r *Router) router() {
}
