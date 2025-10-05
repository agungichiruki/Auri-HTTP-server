package main

import (
	"auri/handler"
	"auri/middleware"
	"fmt"
	"net/http"
)

type ServerConfig struct {
	Port int
	Mode string
	Root string
	Compression string
}

func StartServer(cfg ServerConfig) {
	addr := fmt.Sprintf(":%d", cfg.Port)

	mux := http.NewServeMux()

	var mainHandler http.Handler

	if cfg.Mode == "directory-listing" {
		mainHandler = http.HandlerFunc(handler.DirectoryListingServer(cfg.Root))
	} else {
		mainHandler = http.FileServer(http.Dir(cfg.Root))
	}

	mainHandler = middleware.CompressionMiddleware(mainHandler, cfg.Compression)

	mux.Handle("/", mainHandler)
	
	fmt.Println(`
###########################################################################################
#  Auri HTTP Server                                                                       #
#  a simple, lightweight, and portable HTTP server with a variety of functional features  #
#                                                                                         #
#  Version 0.3~Beta                                                                       #
#  Copyright (c) 2025 Mohammad Agung                                                      #
#  https://github.com/agung/ichiruki                                                      #
#                                                                                         #
#  Licensed under the BSD 3-Clause License                                                #
#  <https://opensource.org/licenses/BSD-3-Clause>                                         #
###########################################################################################

`)
	valid := map[string]bool{"brotli": true, "gzip": true, "": true}
	if !valid[cfg.Compression] {
		fmt.Printf("[warning] Unknown compression type: %s (defaulting to none)\n", cfg.Compression)
		cfg.Compression = ""
	}

	fmt.Printf("[info] Running on http://localhost%s\n", addr)
	if cfg.Compression == "" {
		fmt.Println("[info] Compression: None")
	} else {
		fmt.Println("[info] Compression:", cfg.Compression)
	}
	fmt.Println("")
	fmt.Println("Use [ctrl + c], to shutdown Auri.")
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("[Error]", err)
	}
}