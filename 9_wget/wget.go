package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := flag.String("url", "", "URL сайта для загрузки")
	flag.Parse()

	if *url == "" {
		fmt.Println("Пожалуйста, уточните URL адрес контента для скачивания, используя -url флаг")
		return
	}

	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println("Ошибка адресации URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка получения тела запроса:", err)
		return
	}

	// Extract the base URL
	baseURL := *url
	if strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL[:len(baseURL)-1]
	}
	lastIndex := strings.LastIndex(baseURL, "/")
	baseURL = baseURL[:lastIndex+1]

	// Создание папки, в которой будет лежать контент
	directory := "downloaded"

	err = os.Mkdir(directory, 0755)
	if err != nil {
		fmt.Println("Ошибка создания директории:", err)
		return
	}

	err = ioutil.WriteFile(directory+"/index.html", body, 0644)
	if err != nil {
		fmt.Println("Ошибка сохранения HTML:", err)
		return
	}

	fmt.Println("Веб-сайт успешно загружен в ", directory)
}
