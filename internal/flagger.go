package flagger

import (
	"flag"
)

// Flags adalah struktur yang menyimpan nilai dari setiap flag yang diberikan.
type Flags struct {
	Keyword     string
	KeywordList string
	OnlyFound   bool
	Info        bool
	Link        bool
	Silent      bool
}

// ParseFlags menguraikan flag yang diberikan dan mengembalikan nilai Flags yang sesuai.
func ParseFlags() Flags {
	// Membuat variabel untuk menyimpan nilai flag.
	var keyword, keywordList string
	var onlyFound, info, link, silent bool

	// Menguraikan flag yang diberikan.
	flag.StringVar(&keyword, "q", "", "Single keyword")
	flag.StringVar(&keywordList, "l", "", "Keyword list from file.txt")
	flag.BoolVar(&onlyFound, "of", false, "Display only found")
	flag.BoolVar(&info, "info", false, "Display additional info")
	flag.BoolVar(&link, "link", false, "Display links")
	flag.BoolVar(&silent, "silent", false, "Silent mode")

	flag.Parse()

	// Mengembalikan nilai Flags yang sesuai dengan flag yang diberikan.
	return Flags{
		Keyword:     keyword,
		KeywordList: keywordList,
		OnlyFound:   onlyFound,
		Info:        info,
		Link:        link,
		Silent:      silent,
	}
}
