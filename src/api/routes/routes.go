// Package routes has the main configuration for the notify REST API endpoints.
package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Router object.
var Frouter *fiber.App

func init() {
	// Start fiber router.
	Frouter = fiber.New()

	// Default config.
	Frouter.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1,http://localhost,http://localhost:8080",
		AllowMethods:     "POST,GET,DELETE,OPTIONS,PUT,PATCH,ACL,CANCELUPLOAD,CHECKIN,CHECKOUT,COPY,HEAD,LOCK,MKCALENDAR,MKCOL,MOVE,PROPFIND,PROPPATCH,REPORT,SEARCH,UNCHECKOUT,UNLOCK,VERSION-CONTROL",
		AllowHeaders:     "Origin,Content-Length,Content-Type,X-Requested-With,Cache-Control,Accept-Encoding,X-CSRF-Token,Accept,Cookie,Set-Cookie,authorization,Keep-Alive,X-PINGARUNER,X-PINGOTHER,Overwrite,Destination,Depth,User-Agent,Translate,Range,Content-Range,Timeout,X-File-Size,If-Modified-Since,X-File-Name,Location,Lock-Token,If,Connection,User-Agent,Host,Access-Control-Allow-Origin",
		ExposeHeaders:    "Content-Length",
	}))

	// Run any endpoints configured in notification_sink.go and listen to port 8080.
	notification_sink()
	log.Fatal(Frouter.Listen(":" + "8080"))
}
