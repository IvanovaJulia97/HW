package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

//хендлер для корневого эндпоинта "/"

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

// хендлер для модификаций с файлом эндпоинт "/upload"
func EditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не может быть использован",
			http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка чтения файла",
			http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileWord := make([]byte, 1024)
	v, err := file.Read(fileWord)
	if err != nil && err.Error() != "EOF" {
		http.Error(w, "Ошибка при чтении данных",
			http.StatusInternalServerError)
		return
	}
	fileWord = fileWord[:v]

	convert := service.AutoDetectAndConvert(string(fileWord))

	myLocalFile := time.Now().UTC().Format("2025-03-27T22-12-03") + ".txt"

	err = os.MkdirAll("upload", os.ModePerm)
	if err != nil {
		http.Error(w, "Ошибка создания контейнера для загрузки"+err.Error(), http.StatusInternalServerError)
		return
	}

	failik, err := os.Create("upload/" + myLocalFile)
	if err != nil {
		http.Error(w, "Ошибка создания локального файла",
			http.StatusInternalServerError)
		return
	}
	defer failik.Close()

	_, err = failik.Write([]byte(convert))
	if err != nil {
		http.Error(w, "Ошибка записи в файл",
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convert))

}
