# WebScraper

A web scraping utility for extracting book data from websites, with support for pagination.

## Features

- Extract book titles and prices from web pages
- Optional pagination support to scrape multiple pages
- Configurable timeout for slow websites
- Custom URL and output file support
- JSON output format
- Built with Colly web scraping framework

## Usage

### Basic Usage

```bash
# Scrape first page only (default: books.toscrape.com)
go run main.go
```

### With Pagination

```bash
# Scrape all pages
go run main.go -paginate
```

### Custom Options

```bash
# Full options
go run main.go -url "http://example.com" -output results.json -paginate -timeout 60
```

## Command Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-url` | http://books.toscrape.com | Base URL to scrape |
| `-output` | books.json | Output JSON file |
| `-paginate` | false | Follow pagination links |
| `-timeout` | 30 | Request timeout in seconds |

## Output Format

The scraper outputs a JSON array of books:

```json
[
  {
    "title": "A Light in the Attic",
    "price": "£51.77"
  },
  {
    "title": "Tipping the Velvet",
    "price": "£53.74"
  }
]
```

## Project Structure

```
WebScraper/
├── main.go       # Main application with scraping logic
├── go.mod
├── go.sum
└── README.md
```

## Examples

### Scrape Books to Scrape (all pages)

```bash
go run main.go -paginate
# Output: 1000 livros coletados e salvos em books.json
```

### Scrape with Custom Timeout

```bash
go run main.go -timeout 60 -paginate
```

### Scrape to Custom File

```bash
go run main.go -output my_books.json -paginate
```

## How It Works

1. Creates a Colly collector with configured timeout
2. Registers HTML callbacks for product elements
3. Extracts title (from `title` attribute or text) and price
4. If pagination is enabled, follows "next" links automatically
5. Saves all collected books to JSON file

## Dependencies

- [Colly](https://github.com/gocolly/colly) - Web scraping framework
- [goquery](https://github.com/PuerkitoBio/goquery) - HTML parsing (via Colly)

## Notes

- The default target (books.toscrape.com) is a sandbox site for web scraping practice
- Respect robots.txt and rate limits when scraping other websites
- Some websites may block automated requests
