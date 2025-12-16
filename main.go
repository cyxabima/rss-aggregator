package main

import (
	// "fmt"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cyxabima/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // this nasty thing must be imported for sqlc
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	// rssFeed, err := urlToFeed("https://podcastfeeds.nbcnews.com/HL4TzgYC")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(rssFeed)

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not Found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB URL is not Found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	db := database.New(conn)

	apiConf := apiConfig{
		DB: db,
	}

	// * running a separate go routine which will run in background non blocking
	go startScrapping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/healthz", handlerReadinessHandler)
	v1Router.Get("/err", handlerErrorHandler)

	v1Router.Get("/user", apiConf.middlewareAuth(apiConf.handlerGetUser))
	v1Router.Post("/user", apiConf.handlerCreateUser)

	v1Router.Get("/feeds", apiConf.handlerGetFeeds)
	v1Router.Post("/feeds", apiConf.middlewareAuth(apiConf.handlerCreateFeed))

	v1Router.Get("/feed_follows", apiConf.middlewareAuth(apiConf.handlerGetFeedFollow))
	v1Router.Post("/feed_follows", apiConf.middlewareAuth(apiConf.handlerCreateFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiConf.middlewareAuth(apiConf.handlerDeleteFeedFollow))

	svr := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = svr.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
