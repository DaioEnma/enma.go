package enma_test

import (
	"fmt"
	"testing"

	"github.com/DaioEnma/enma.go"
	"github.com/go-resty/resty/v2"
)

func TestEnma(t *testing.T) {
	fmt.Println("testing enma imports")

	hianime := enma.HiAnime.NewScraper()
	mangareader := enma.MangaReader.NewScraper()

	fmt.Println(hianime, mangareader, resty.Client{})
}
