package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/syhbt/fynd/pkg/matcher" // Ubah sesuai dengan path package matcher yang sebenarnya
)

// Run melakukan proses berdasarkan input flag yang diberikan.
func Run(keyword, keywordList string) {
	if keyword != "" {
		// Jika menggunakan flag -q, ambil input langsung.
		processInput(keyword)
	} else if keywordList != "" {
		// Jika menggunakan flag -l, ambil input dari file.
		processFile(keywordList)
	} else {
		// Jika tidak memasukkan flag -q atau -l, atau dari inputan pipeline, tampilkan error.
		fmt.Println("Masukan input")
	}
}

func processInput(keyword string) {
	// Lakukan pengecekan dengan menggunakan matcher dari package matcher.
	matchResult := matcher.Match(keyword)

	// Tampilkan hasil match.
	fmt.Println(matchResult)
}

func processFile(keywordListFile string) {
	// Buka file untuk membaca daftar kata kunci.
	file, err := os.Open(keywordListFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Membaca setiap baris dari file sebagai kata kunci dan melakukan proses pada setiap kata kunci.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyword := strings.TrimSpace(scanner.Text())
		if keyword != "" {
			// Lakukan pengecekan dengan menggunakan matcher dari package matcher untuk setiap kata kunci.
			matchResult := matcher.Match(keyword)
			// Tampilkan hasil match.
			fmt.Println(matchResult)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

