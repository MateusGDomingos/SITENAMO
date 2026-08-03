package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Declaracao struct {
	Titulo string
	Texto  string
}

type Foto struct {
	Src string
	Alt string
}

type Dados struct {
	Nome1       string
	Nome2       string
	DataInicio  time.Time
	Declaracoes []Declaracao
	Fotos       []Foto
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; "+
				"style-src 'self' https://fonts.googleapis.com 'unsafe-inline'; "+
				"font-src 'self' https://fonts.gstatic.com; "+
				"img-src 'self' data:; "+
				"script-src 'self'; "+
				"connect-src 'self'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

func indexHandler(tmpl *template.Template, dados Dados) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, dados); err != nil {
			log.Printf("Erro ao renderizar template: %v", err)
			http.Error(w, "Erro interno", http.StatusInternalServerError)
		}
	}
}

func main() {
	dados := Dados{
		Nome1:      "Isabela",
		Nome2:      "Mateus",
		DataInicio: time.Date(2026, 05, 15, 19, 00, 00, 0, time.UTC),
		Declaracoes: []Declaracao{
			{
				Titulo: "O Início",
				Texto:  "Foi num dia comum que tudo mudou. Um olhar, um sorriso, e de repente o mundo ganhou cores que eu nunca tinha visto antes. Você apareceu e a vida fez sentido.",
			},
			{
				Titulo: "Por Que Você",
				Texto:  "Não foi planejado, não foi previsto, mas foi exatamente como deveria ser. Você chegou devagarinho e ficou pra sempre. Agora eu não consigo imaginar meus dias sem os seus.",
			},
			{
				Titulo: "Cada Dia",
				Texto:  "Cada dia ao seu lado é um presente que eu não sabia que queria. Cada risada sua é música pra mim. Cada abraço teu é onde eu quero morar. Você é meu lar.",
			},
			{
				Titulo: "Pra Sempre",
				Texto:  "Se amar fosse um lugar, eu teria mudado pra lá no primeiro dia. Se amar fosse um tempo, eu ia querer a eternidade. Você é o melhor que me aconteceu, e eu quero te fazer o mais feliz do mundo. Pra sempre, eu te amo.",
			},
		},
		Fotos: []Foto{
			{Src: "/static/img/foto-1.jpeg", Alt: "Isabela e Mateus - foto 1"},
			{Src: "/static/img/foto-2.jpeg", Alt: "Isabela e Mateus - foto 2"},
			{Src: "/static/img/foto-3.jpeg", Alt: "Isabela e Mateus - foto 3"},
			{Src: "/static/img/foto-4.jpeg", Alt: "Isabela e Mateus - foto 4"},
			{Src: "/static/img/foto-5.jpeg", Alt: "Isabela e Mateus - foto 5"},
			{Src: "/static/img/foto-6.jpeg", Alt: "Isabela e Mateus - foto 6"},
		},
	}

	tmpl := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"formatoData": func(t time.Time) string {
			return t.Format("02/01/2006")
		},
		"formatoDataISO": func(t time.Time) string {
			return t.Format("2006-01-02T15:04:05")
		},
		"anoAtual": func() int {
			return time.Now().Year()
		},
	}).ParseFiles("templates/index.html"))

	mux := http.NewServeMux()
	mux.Handle("/", securityHeaders(indexHandler(tmpl, dados)))
	mux.Handle("/static/", securityHeaders(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
