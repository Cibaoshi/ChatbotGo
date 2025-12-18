package main

import (
	"html/template"
	"net/http"
	"os"
)

func renderTemplate(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка сервера: не удалось загрузить шаблон", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, PageData{})
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userPrompt := r.FormValue("prompt")
	if userPrompt == "" {
		renderTemplate(w, PageData{Error: "Вопрос не может быть пустым."})
		return
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		renderTemplate(w, PageData{Prompt: userPrompt, Error: "GEMINI_API_KEY не установлен на сервере."})
		return
	}

	responseText, err := sendToGemini(userPrompt, apiKey)
	if err != nil {
		renderTemplate(w, PageData{Prompt: userPrompt, Error: err.Error()})
		return
	}

	renderTemplate(w, PageData{Prompt: userPrompt, Response: responseText})
}
