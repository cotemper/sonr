package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed nextjs/dist/_next/static/css/*.css
//go:embed nextjs/dist/_next/static/media/*.png
//go:embed nextjs/dist/_next/static/chunks/pages/*.js
//go:embed nextjs/dist/_next/static/*/*.js
var nextFS embed.FS

func Start() {
	// Root at the `dist` folder generated by the Next.js app.
	distFS, err := fs.Sub(nextFS, "nextjs/dist")
	if err != nil {
		log.Fatal(err)
	}

	// The static Next.js app will be served under `/`.
	http.Handle("/", http.FileServer(http.FS(distFS)))
	// The API will be served under `/api`.

	// Start HTTP server at :8080.
	log.Println("Starting HTTP server at http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}