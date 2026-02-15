package main

// These imports will be used later on the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {
	month := flag.StringP("month", "m", "", "Month Name")
	year := flag.IntP("year", "y", 2026, "Year 2022 to 2026")
	if *month == "" {
		fmt.Println("Error: --month or -m is required.")
		flag.Usage() // Prints the help menu automatically
		os.Exit(1)   // Exit with an error code
	}
	// p := tea.NewProgram(initialModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Ohhoo there has been an Error: %v", err)
	// 	os.Exit(1)
	// }
	articles, err := scrape(*month, *year)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(articles)
}
