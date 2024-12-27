package hianime

func (hianime *Scraper) GetHomePage() string {
	return hianime.client.BaseURL
}
