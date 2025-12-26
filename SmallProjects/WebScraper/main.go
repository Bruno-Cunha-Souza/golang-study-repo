package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const (
	baseURL       = "http://books.toscrape.com"
	outputFile    = "books.json"
	productPod    = ".product_pod"
	titleSelector = "h3 a"
	priceSelector = ".price_color"
)

type Book struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func main() {
	books, err := scrapeBooks(baseURL)
	if err != nil {
		log.Fatalf("Erro ao fazer scraping dos livros: %v", err)
	}

	if err := saveBooksToFile(books, outputFile); err != nil {
		log.Fatalf("Erro ao salvar livros em arquivo: %v", err)
	}

	fmt.Printf("Livros coletados e salvos em %s\n", outputFile)
}

func scrapeBooks(url string) ([]Book, error) {
	var books []Book

	c := colly.NewCollector()

	c.OnHTML(productPod, func(e *colly.HTMLElement) {
		book := Book{
			Title: e.ChildText(titleSelector),
			Price: e.ChildText(priceSelector),
		}
		books = append(books, book)
		fmt.Printf("Livro encontrado: %s -> %s\n", book.Title, book.Price)
	})

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("erro ao visitar URL: %w", err)
	}

	return books, nil
}

func saveBooksToFile(books []Book, filename string) error {
	file, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao criar JSON: %w", err)
	}

	if err := os.WriteFile(filename, file, 0644); err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	return nil
}
