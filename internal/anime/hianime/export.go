package hianime

type HiAnimeExport struct {
	SEARCH_PAGE_FILTERS  searchPageFilters
	AZ_LIST_SORT_OPTIONS map[string]bool
}

func (*HiAnimeExport) NewScraper() *Scraper {
	return New()
}

var Export = &HiAnimeExport{
	SEARCH_PAGE_FILTERS:  SEARCH_PAGE_FILTERS,
	AZ_LIST_SORT_OPTIONS: AZ_LIST_SORT_OPTIONS,
}
