package http

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//"github.com/xpharos/web/http/cookie"
	// "http/middleware"
	// "http/render"

	"conf"
)

func Start() {
	// render.Init()
	//cookie.Init()

	r := mux.NewRouter().StrictSlash(false)
	ConfigRouter(r)

	n := negroni.New()
	// n.Use(middleware.NewRecovery())
	n.UseHandler(r)
	n.Run(conf.GlobalConfig.HttpConfig.Listen)
}
