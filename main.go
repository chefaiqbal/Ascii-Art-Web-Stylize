package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Struct to hold the data for rendering the main page
type MainPageData struct {
	Message string
}

// Struct to hold the data for rendering the ASCII art result page
type ResultPageData struct {
	Input       string
	BannerStyle string
	ASCIIArt    string
}

// Function to load the banner file
func loadBanner(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error loading banner file: %v", err)
	}
	return string(data), nil
}

// Function to generate ASCII art from the input string and banner
func generateASCIIArt(input string, banner string) string {
	var output strings.Builder

	// Split the input into lines
	lines := strings.Split(input, "\n")

	// Split the banner into lines
	bannerLines := strings.Split(banner, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		// Generate the ASCII art representation for the line
		letters := strings.Split(line, "")
		for j := 1; j < 9; j++ {
			str := ""
			for _, letter := range letters {
				char := []rune(letter)
				if len(char) > 0 && char[0] >= 32 && char[0] <= 126 {
					index := (int(char[0]) - 32) * 9
					str += bannerLines[index+j]
				}
			}
			output.WriteString(str)
			output.WriteByte('\n')
		}
	}

	return output.String()
}

// Handler for the main page
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}

	// Render the main page template
	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		http.Error(w, "Error:500\nInternal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Function to validate the input string for non-Basic Latin characters
func validateInput(input string) bool {
	for _, char := range input {
		// Check if the character is outside the Basic Latin range (32 to 126) and not a new line character
		if char != '\n' && (char < 32 || char > 126) {
			return false
		}
	}
	return true
}

// Handler for the ASCII art generation
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Get the input string and banner style from the form data
	input := strings.ReplaceAll(r.Form.Get("input"), "\r\n", "\n")
	bannerStyle := r.Form.Get("banner")

	// Validate the input for non-Basic Latin characters
	if !validateInput(input) {
		http.Error(w, "Error:400\nInvalid input characters\nOnly Latin characters allowed", http.StatusBadRequest)
		return
	}

	// Load the banner file
	banner, err := loadBanner("banners/" + bannerStyle + ".txt")
	if err != nil {
		http.Error(w, "Error:404\nBanner file not found: "+bannerStyle, http.StatusNotFound)
		return
	}

	// Generate the ASCII art representation
	asciiArt := generateASCIIArt(input, banner)

	// Render the result page template
	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	data := ResultPageData{Input: input, BannerStyle: bannerStyle, ASCIIArt: asciiArt}
	tmpl.Execute(w, data)
}

// Handler for page not found
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error:404\nPage not found", http.StatusNotFound)
}

func main() {
	// Register the request handlers
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/favicon.ico", http.NotFound) // Ignore favicon requests

/*Keeping for stylize Project
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	*/

	// Handle page not found
	http.HandleFunc("/404", notFoundHandler)

	// Start the HTTP server in a separate goroutine
	go func() {
		fmt.Println("Server running on http://localhost:8080 \nTo stop the server press Ctrl+C")
		http.ListenAndServe(":8080", nil)
	}()

	// Wait for a termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("Server stopped")
}


//TODO
//400 BAD REQUEST FOR NON LATIN CHAR
