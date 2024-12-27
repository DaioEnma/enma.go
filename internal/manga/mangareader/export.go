package mangareader

type MangaReaderExport struct{}

func (*MangaReaderExport) NewScraper() *Scraper {
	return New()
}

var Export = &MangaReaderExport{}
