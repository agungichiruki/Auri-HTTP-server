package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"html/template"
)

type FileEntry struct {
	Name string
	Size string
	ModTime string
	Icon template.HTML
}

func DirectoryListingServer(root string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(root, r.URL.Path)

		info, err := os.Stat(path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				http.Error(w, "Tidak bisa membaca direktori", http.StatusInternalServerError)
				return
			}

			files := []FileEntry{}
			for _, e := range entries {
				fi, _ := e.Info()
				size := "-"
				if !fi.IsDir() {
					size = fmt.Sprintf("%d bytes", fi.Size())
				}
				files = append(files, FileEntry{
					Name: e.Name(),
					Size: size,
					ModTime: fi.ModTime().Format(time.RFC1123),
					Icon: getFileIcon(e.Name(), fi.IsDir()),
				})
			}

			parent := ""
			if r.URL.Path != "/" {
				parent = filepath.Dir(r.URL.Path)
				if parent != "/" {
					parent += "/"
				}
			}

			dirTemplate.Execute(w, map[string]any {
				"Path": r.URL.Path,
				"Files": files,
				"Parent": parent,
			})
			return
		}

		http.ServeFile(w, r, path)
	}
}