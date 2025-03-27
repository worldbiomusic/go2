package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/worldbiomusic/go2/utils"
)

var ctx = context.Background()

func main() {
	log.Println("[Go2]: URL 단축생성기")

	// DB connection 만들기
	dbClient := utils.NewRedisClient()
	if dbClient == nil {
		log.Println("Redis 연결 실패")
		return
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, "URL 단축 사이트 Go2 입니다.")
	})

	http.HandleFunc("/go2", func(writer http.ResponseWriter, req *http.Request) {
		url := req.FormValue("url")
		log.Println("Payload: ", url)

		shortCode := utils.ShortCode(url)
		shortURL := fmt.Sprintf("http://localhost:8080/g/%s", shortCode)
		log.Println("짧은 URL 생성: ", shortURL)

		utils.SetKey(&ctx, dbClient, shortURL, url, 0)
		fmt.Fprintf(writer,
			`<p class="mt-4 text-green-600">Shortened URL: <a href="/r/%s" class="underline">%s</a></p>`, shortURL, shortURL)
	})

	log.Println("http://localhost:8080 서버 온")
	http.ListenAndServe(":8080", nil)
}
