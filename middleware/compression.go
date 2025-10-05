package middleware

import (
	"compress/gzip"
	"net/http"
	"github.com/andybalholm/brotli"
)

type brotliResponseWriter struct {
	http.ResponseWriter
	Writer *brotli.Writer
}

func (w *brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func CompressionMiddleware(next http.Handler, compressionType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch compressionType {
			case "brotli":
				w.Header().Set("Content-Encoding", "br")
				br := brotli.NewWriter(w)
				defer br.Close()
				bw := &brotliResponseWriter{ResponseWriter: w, Writer: br}
				next.ServeHTTP(bw, r)
			case "gzip":
				w.Header().Set("Content-Encoding", "gzip")
				gz := gzip.NewWriter(w)
				defer gz.Close()
				gw := &gzipResponseWriter{ResponseWriter: w, Writer: gz}
				next.ServeHTTP(gw, r)
			default:
			next.ServeHTTP(w, r)
		}
	})
}