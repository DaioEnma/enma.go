package hianime

import (
	"github.com/DaioEnma/enma.go/internal/util"
	"github.com/go-resty/resty/v2"
	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	colly  *colly.Collector
	client *resty.Client
}

func New() *Scraper {
	return &Scraper{
		colly: colly.NewCollector(
			colly.AllowedDomains(domain),
		),
		client: resty.New().
			SetBaseURL(src_base_url).
			SetHeaders(map[string]string{
				"Referer":         src_base_url + src_home_path,
				"Accept":          util.ACCEPT_HEADER,
				"User-Agent":      util.USER_AGENT_HEADER,
				"Accept-Encoding": util.ACCEPT_ENCODING_HEADER,
			}),
	}
}

const domain string = "hianime.to"

const (
	src_base_url    string = "https://" + domain
	src_home_path   string = "/home"
	src_ajax_path   string = "/ajax"
	src_search_path string = "/search"
)

type searchPageFilters struct {
	GENRES_ID_MAP   map[string]int
	TYPE_ID_MAP     map[string]int
	STATUS_ID_MAP   map[string]int
	RATED_ID_MAP    map[string]int
	SCORE_ID_MAP    map[string]int
	SEASON_ID_MAP   map[string]int
	LANGUAGE_ID_MAP map[string]int
	SORT_ID_MAP     map[string]string
}

var SEARCH_PAGE_FILTERS = searchPageFilters{
	GENRES_ID_MAP: map[string]int{
		"action":        1,
		"adventure":     2,
		"cars":          3,
		"comedy":        4,
		"dementia":      5,
		"demons":        6,
		"drama":         8,
		"ecchi":         9,
		"fantasy":       10,
		"game":          11,
		"harem":         35,
		"historical":    13,
		"horror":        14,
		"isekai":        44,
		"josei":         43,
		"kids":          15,
		"magic":         16,
		"martial-arts":  17,
		"mecha":         18,
		"military":      38,
		"music":         19,
		"mystery":       7,
		"parody":        20,
		"police":        39,
		"psychological": 40,
		"romance":       22,
		"samurai":       21,
		"school":        23,
		"sci-fi":        24,
		"seinen":        42,
		"shoujo":        25,
		"shoujo-ai":     26,
		"shounen":       27,
		"shounen-ai":    28,
		"slice-of-life": 36,
		"space":         29,
		"sports":        30,
		"super-power":   31,
		"supernatural":  37,
		"thriller":      41,
		"vampire":       32,
	},
	TYPE_ID_MAP: map[string]int{
		"all":     0,
		"movie":   1,
		"tv":      2,
		"ova":     3,
		"ona":     4,
		"special": 5,
		"music":   6,
	},
	STATUS_ID_MAP: map[string]int{
		"all":              0,
		"finished-airing":  1,
		"currently-airing": 2,
		"not-yet-aired":    3,
	},
	RATED_ID_MAP: map[string]int{
		"all":   0,
		"g":     1,
		"pg":    2,
		"pg-13": 3,
		"r":     4,
		"r+":    5,
		"rx":    6,
	},
	SCORE_ID_MAP: map[string]int{
		"all":         0,
		"appalling":   1,
		"horrible":    2,
		"very-bad":    3,
		"bad":         4,
		"average":     5,
		"fine":        6,
		"good":        7,
		"very-good":   8,
		"great":       9,
		"masterpiece": 10,
	},
	SEASON_ID_MAP: map[string]int{
		"all":    0,
		"spring": 1,
		"summer": 2,
		"fall":   3,
		"winter": 4,
	},
	LANGUAGE_ID_MAP: map[string]int{
		"all":       0,
		"sub":       1,
		"dub":       2,
		"sub-&-dub": 3,
	},
	SORT_ID_MAP: map[string]string{
		"default":          "default",
		"recently-added":   "recently_added",
		"recently-updated": "recently_updated",
		"score":            "score",
		"name-a-z":         "name_az",
		"released-date":    "released_date",
		"most-watched":     "most_watched",
	},
}

var AZ_LIST_SORT_OPTIONS = map[string]bool{
	"all":   true,
	"other": true,
	"0-9":   true,
	"a":     true,
	"b":     true,
	"c":     true,
	"d":     true,
	"e":     true,
	"f":     true,
	"g":     true,
	"h":     true,
	"i":     true,
	"j":     true,
	"k":     true,
	"l":     true,
	"m":     true,
	"n":     true,
	"o":     true,
	"p":     true,
	"q":     true,
	"r":     true,
	"s":     true,
	"t":     true,
	"u":     true,
	"v":     true,
	"w":     true,
	"x":     true,
	"y":     true,
	"z":     true,
}
