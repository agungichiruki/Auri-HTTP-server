package main

import (
	"fmt"
	"net/http"
	"auri/handler"
	"auri/middleware"
	"flag"
)

func main() {
	//fs := http.FileServer(http.Dir("./html"))

	port := flag.Int("port", 8080, "[--port] The port used to run Auri. By default, Auri will run on port 8080.")
	mode := flag.String("mode", "", "[--mode] By default, Auri will act as an HTTP server, serving the index.html file. Use the [--mode directory-listing] option, to run Auri as a directory listing server.")
	root := flag.String("root", ".", "[--root] The directory where files will be served by the Auri HTTP Server. By default, Auri will read files in the same location as the Auri HTTP Server.")
	compression := flag.String("compression", "", "[--compression] Compression to use. [--compression brotli] to use brotli compression, [--compression gzip] to use gzip compression. By default, Auri does not use any compression.")
	flag.Usage = func() {
		fmt.Println(`Auri HTTP Server - A simple, lightweight, and portable HTTP server with a variety of functional features.
Licensed under the BSD 3-Clause License
`)
		fmt.Println("Usage: auri [OPTIONS]")
		fmt.Println("")
		fmt.Println("Available options:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println(`Example: auri --port 80 --mode directory-listing --root /var/www/html/"
Or run Auri with default settings (no need to include any options).
Example: auri`)
	}
	
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)

	if *mode == "directory-listing" {
		http.HandleFunc("/", handler.CustoFileServer(*root))
	} else {
		fs := http.FileServer(http.Dir(*root))
		http.Handle("/", middleware.CompressionMiddleware(fs, *compression))
	}
	
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
	fmt.Printf("[info] Running on http://localhost%s\n", addr)
	if *compression == "" {
		fmt.Println("[info] Compression: None")
	} else {
		fmt.Println("[info] Compression:", *compression)
	}
	fmt.Println("")
	fmt.Println("Use [ctrl + c], to shutdown Auri.")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}