package main

import (
	"context"
	"fmt"
	"html/template"
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
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(writer, nil)
	})

	http.HandleFunc("/go2", func(writer http.ResponseWriter, req *http.Request) {
		url := req.FormValue("url")
		log.Println("요청 URL: ", url)

		shortCode := utils.ShortCode(url)
		shortURL := fmt.Sprintf("http://localhost:8080/go2/%s", shortCode)
		log.Print("생성 URL: ", shortURL, "\n\n")

		utils.SetKey(&ctx, dbClient, shortCode, url, 0)
		fmt.Fprintf(writer,
			`<p class="mt-4 text-green-600">짧은 URL: <a href="/go2/%s" class="underline">%s</a></p>`, shortCode, shortURL)
	})

	http.HandleFunc("/go2/{code}", func(writer http.ResponseWriter, req *http.Request) {
		code := req.PathValue("code")
		if code == "" {
			http.Error(writer, "URL이 없습니다.", http.StatusBadRequest)
			return
		}

		originURL, err := utils.GetOriginURL(&ctx, dbClient, code)
		if err != nil {
			log.Println("err: ", err)
			http.Error(writer, "해당 URL은 없습니다.", http.StatusNotFound)
			return
		}

		log.Println("originURL: ", originURL)
		http.Redirect(writer, req, originURL, http.StatusPermanentRedirect)
	})

	// 서버 시작
	log.Println("http://localhost:8080 서버 온")
	log.Println("===========================\n")
	http.ListenAndServe(":8080", nil)
}
