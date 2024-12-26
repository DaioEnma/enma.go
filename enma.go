package enma

import (
	"github.com/DaioEnma/enma.go/src/providers/anime/hianime"
	"github.com/DaioEnma/enma.go/src/providers/manga/mangareader"
)

func NewHiAnime() *hianime.Scraper {
	return hianime.New()
}

func NewMangaReader() *mangareader.Scraper {
	return &mangareader.Scraper{}
}
