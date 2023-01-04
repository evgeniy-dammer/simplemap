package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	addr       = flag.String("addr", "0.0.0.0:8082", "TCP address to listen to")
	dir        = flag.String("dir", "./tiles/", "Directory to serve images from")
	defaultImg = flag.String("defaultImg", "./tiles/default.png", "Default image")
)

func main() {
	flag.Parse()

	r := router.New()
	r.GET("/tiles/{z}/{y}/{x}", Tiles)

	log.Printf("Server started on %s", *addr)

	if err := fasthttp.ListenAndServe(*addr, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func Tiles(ctx *fasthttp.RequestCtx) {
	path := *dir +
		ctx.UserValue("z").(string) + "/" +
		ctx.UserValue("y").(string) + "/" +
		ctx.UserValue("x").(string) + ".png"

	ctx.Response.Header.Set("Content-Type", "image/png")

	if _, err := os.Stat(path); err == nil {
		err = ctx.Response.SendFile(path)
		if err != nil {
			log.Printf("Unable to send image on path: %s", path)
		}
		log.Printf("Send image on path: %s", path)
	} else if errors.Is(err, os.ErrNotExist) {
		err = ctx.Response.SendFile(*defaultImg)
		if err != nil {
			log.Println("Unable to send default image")
		}
		log.Println("Send default image")
	}
}
