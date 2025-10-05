package main

import (
	"fmt"
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

	config := ServerConfig{
		Port: *port,
		Mode: *mode,
		Root: *root,
		Compression: *compression,
	}

	StartServer(config)
}