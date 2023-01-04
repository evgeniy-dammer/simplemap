package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	addr               = flag.String("addr", "0.0.0.0:8080", "TCP address to listen to")
	byteRange          = flag.Bool("byteRange", false, "Enables byte range requests if set to true")
	compress           = flag.Bool("compress", false, "Enables transparent response compression if set to true")
	dir                = flag.String("dir", "./html", "Directory to serve static files from")
	generateIndexPages = flag.Bool("generateIndexPages", true, "Whether to generate directory index pages")
)

func main() {
	flag.Parse()

	fs := &fasthttp.FS{
		Root:               *dir,
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: *generateIndexPages,
		Compress:           *compress,
		AcceptByteRange:    *byteRange,
	}

	fsHandler := fs.NewRequestHandler()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		default:
			fsHandler(ctx)
		}
	}

	if len(*addr) > 0 {
		log.Printf("Starting front on %q", *addr)
		go func() {
			if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServe: %v", err)
			}
		}()
	}

	select {}
}
