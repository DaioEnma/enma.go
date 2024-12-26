package enma_test

import (
	"fmt"
	"testing"

	"github.com/DaioEnma/enma.go"
	"github.com/go-resty/resty/v2"
	"github.com/gocolly/colly/v2"
)

func TestEnma(t *testing.T) {
	fmt.Println("testing enma imports")

	hianime := enma.NewHiAnime()
	mangareader := enma.NewMangaReader()

	fmt.Println(hianime, mangareader, colly.Collector{}, resty.Client{})
}
