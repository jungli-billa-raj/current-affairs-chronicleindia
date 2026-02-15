package main

// These imports will be used later on the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	// "log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	flag "github.com/spf13/pflag"
)

func main() {
	month := flag.StringP("month", "m", "", "Month Name. (january/jan, short name allowed)")
	year := flag.IntP("year", "y", 2026, "Year 2022 to 2026")
	flag.Parse()
	if *month == "" || (*year < 2022 || *year > 2026) {
		flag.Usage() // Prints the help menu automatically
		fmt.Println("Note: --month or -m is required.")
		os.Exit(1) // Exit with an error code
	}
	monthMap := map[string]string{
		"jan": "january", "january": "january",
		"feb": "february", "february": "february",
		"mar": "march", "march": "march",
		"apr": "april", "april": "april",
		"may": "may",
		"jun": "june", "june": "june",
		"jul": "july", "july": "july",
		"aug": "august", "august": "august",
		"sep": "september", "september": "september",
		"oct": "october", "october": "october",
		"nov": "november", "november": "november",
		"dec": "december", "december": "december",
	}

	m := strings.ToLower(*month)
	fullMonth, ok := monthMap[m]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}
	p := tea.NewProgram(initialModel(fullMonth, *year), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ohhoo there has been an Error: %v", err)
		os.Exit(1)
	}
}

// articles, err := scrape(fullMonth, *year)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Print(articles)
