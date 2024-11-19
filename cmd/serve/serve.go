package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/synic/adamthings.me/internal/middleware"
	"github.com/synic/adamthings.me/internal/model"
	"github.com/synic/adamthings.me/internal/route"
	"github.com/synic/adamthings.me/internal/store"
)

var (
	// will maybe be set to `true` in `debug.go`
	isDebugBuild = false
	//go:embed articles/*
	articleFiles embed.FS
	//go:embed assets/*
	assetFiles embed.FS
)

func main() {
	var articles []*model.Article

	entries, err := articleFiles.ReadDir("articles")

	if err != nil {
		log.Fatal(err)
	}

	begin := time.Now()
	for _, entry := range entries {
		var article model.Article

		name := path.Join("articles", entry.Name())

		if filepath.Ext(name) != ".json" {
			continue
		}

		data, err := articleFiles.ReadFile(name)

		if err != nil {
			log.Fatalf("error reading article %s: %v", name, err)
		}

		err = json.Unmarshal(data, &article)

		if err != nil {
			log.Fatalf("error unmarshaling article %s: %v", name, err)
		}

		articles = append(articles, &article)
	}

	log.Printf("Read %d articles in %s", len(articles), time.Since(begin))

	repo, err := store.NewArticleRepository(articles)

	if err != nil {
		log.Fatal(err)
	}

	bind := os.Getenv("BIND")
	if bind == "" {
		bind = "0.0.0.0:3000"
	}

	sub, err := fs.Sub(assetFiles, "assets")

	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServeMux()
	static := http.StripPrefix("/static/", http.FileServer(http.FS(sub)))
	server.Handle("GET /static/", static)
	route.NewArticleRouter(repo).Mount(server)

	log.Printf("Serving on %s...", bind)

	wrapped := middleware.Wrap(
		server,
		middleware.CacheStaticFiles(
			assetFiles,
			"js/htmx.min.js",
			"css/syntax.min.css",
			"css/main.css",
		),
		middleware.LoggingMiddleware(log.Default()),
		middleware.IsHtmxPartialMiddleware(),
		middleware.GzipMiddleware(),
	)

	if err = http.ListenAndServe(bind, wrapped); err != nil {
		log.Fatal(err)
	}
}
