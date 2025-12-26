package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

const (
	defaultURL    = "http://books.toscrape.com"
	defaultOutput = "books.json"
	productPod    = ".product_pod"
	titleSelector = "h3 a"
	priceSelector = ".price_color"
	nextSelector  = ".next a"
)

type Book struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func main() {
	url := flag.String("url", defaultURL, "URL base para scraping")
	output := flag.String("output", defaultOutput, "Arquivo de saída JSON")
	paginate := flag.Bool("paginate", false, "Seguir links de paginação")
	timeout := flag.Int("timeout", 30, "Timeout em segundos para requisições")
	flag.Parse()

	books, err := scrapeBooks(*url, *paginate, time.Duration(*timeout)*time.Second)
	if err != nil {
		log.Fatalf("Erro ao fazer scraping: %v", err)
	}

	if len(books) == 0 {
		fmt.Println("Nenhum livro encontrado.")
		return
	}

	if err := saveBooksToFile(books, *output); err != nil {
		log.Fatalf("Erro ao salvar arquivo: %v", err)
	}

	fmt.Printf("%d livros coletados e salvos em %s\n", len(books), *output)
}

func scrapeBooks(url string, paginate bool, timeout time.Duration) ([]Book, error) {
	var books []Book

	c := colly.NewCollector(
		colly.MaxDepth(50),
	)
	c.SetRequestTimeout(timeout)

	c.OnHTML(productPod, func(e *colly.HTMLElement) {
		book := Book{
			Title: e.ChildAttr(titleSelector, "title"),
			Price: e.ChildText(priceSelector),
		}
		if book.Title == "" {
			book.Title = e.ChildText(titleSelector)
		}
		books = append(books, book)
		fmt.Printf("Livro encontrado: %s -> %s\n", book.Title, book.Price)
	})

	if paginate {
		c.OnHTML(nextSelector, func(e *colly.HTMLElement) {
			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
			c.Visit(nextPage)
		})
	}

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Erro ao acessar %s: %v\n", r.Request.URL, err)
	})

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("erro ao visitar URL: %w", err)
	}

	c.Wait()
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
