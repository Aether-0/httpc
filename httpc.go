package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Default HTTP methods to test
var defaultMethods = []string{"GET", "PUT", "POST", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE", "CONNECT"}

// Configuration options with default values
var (
	concurrency = 10
	timeout     = 5 * time.Second
	retries     = 1
	userAgent   = "Go-HTTP-Client/1.1"
	silent      = false
	verbose     = false
	outputFile  = ""
	methods     = strings.Join(defaultMethods, ",")
	singleURL   = ""   // New flag for a single URL
	statusCodes = ""   // New flag for filtering specific status codes
	statusMap   = make(map[int]bool)
	// Colors
	colorReset   = "\033[0m"
	colorGreen   = "\033[32m"
	colorBlue    = "\033[34m"
	colorRed     = "\033[31m"
	colorYellow  = "\033[33m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
)

// Function to print a stylized banner
func printBanner() {
	fmt.Println(colorCyan + `
   __   __  __       _____
  / /  / /_/ /____  / ___/
 / _ \/ __/ __/ _ \/ /__  
/_//_/\__/\__/ .__/\___/  
            /_/           

HTTP Method Tester
` + colorReset)
	fmt.Println("Default Methods:", colorMagenta+"GET, PUT, POST, DELETE, PATCH, HEAD, OPTIONS, TRACE, CONNECT"+colorReset)
	fmt.Println()
}

// Log functions for styled output
func logInfo(message string) {
	if !silent {
		fmt.Printf("%s[INFO]%s %s\n", colorCyan, colorReset, message)
	}
}

func logError(message string) {
	fmt.Printf("%s[ERROR]%s %s\n", colorRed, colorReset, message)
}

func logSuccess(message string) {
	fmt.Printf("%s[SUCCESS]%s %s\n", colorGreen, colorReset, message)
}

// Normalizes URL by adding "https://" if missing
func normalizeURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "https://" + strings.TrimSpace(url)
	}
	return strings.TrimSpace(url)
}

// Checks if a status code should be printed based on the filter
func shouldPrintStatus(statusCode int) bool {
	if len(statusMap) == 0 { // If no filter is set, print all statuses
		return true
	}
	return statusMap[statusCode]
}

// Prints colored output based on HTTP status code
func printColoredOutput(statusCode int, method string, url string) {
	var color string
	switch {
	case statusCode >= 200 && statusCode < 300:
		color = colorGreen // Green for success
	case statusCode >= 300 && statusCode < 400:
		color = colorBlue // Blue for redirection
	case statusCode >= 400 && statusCode < 500:
		color = colorMagenta // Magenta for client errors
	case statusCode >= 500:
		color = colorYellow // Yellow for server errors
	default:
		color = colorWhite // White for other statuses
	}
	fmt.Printf("%s[%s] %s - %d%s\n", color, method, url, statusCode, colorReset)
}

// Sends an HTTP request with retries and custom headers
func sendRequest(url, method string) (int, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", userAgent)

	for i := 0; i < retries+1; i++ {
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			defer resp.Body.Close()
			return resp.StatusCode, nil
		}
		logError(fmt.Sprintf("Retrying %s [%d/%d]", url, i+1, retries+1))
		time.Sleep(1 * time.Second)
	}
	return 0, fmt.Errorf("failed after %d retries", retries)
}

// Handle a single URL with concurrent requests for each HTTP method
func handleURL(url string, methods []string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()
	url = normalizeURL(url)

	for _, method := range methods {
		statusCode, err := sendRequest(url, method)
		if err != nil {
			logError(fmt.Sprintf("Error with %s [%s]: %v", url, method, err))
			continue
		}
		if shouldPrintStatus(statusCode) {
			output := fmt.Sprintf("[%s] %s - %d", method, url, statusCode)
			printColoredOutput(statusCode, method, url)
			if outputFile != "" {
				resultChan <- output
			}
		}
	}
}

// Writes results to output file if specified
func writeResults(resultChan chan string, doneChan chan struct{}) {
	if outputFile == "" {
		close(doneChan)
		return
	}

	file, err := os.Create(outputFile)
	if err != nil {
		logError(fmt.Sprintf("Could not create output file: %v", err))
		close(doneChan)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for result := range resultChan {
		writer.WriteString(result + "\n")
	}
	writer.Flush()
	close(doneChan)
}

func parseStatusCodes() {
	if statusCodes == "" {
		return
	}
	codes := strings.Split(statusCodes, ",")
	for _, code := range codes {
		status, err := strconv.Atoi(strings.TrimSpace(code))
		if err == nil {
			statusMap[status] = true
		} else {
			logError(fmt.Sprintf("Invalid status code filter: %s", code))
		}
	}
}

func main() {
	// Command-line argument parsing
	flag.IntVar(&concurrency, "c", concurrency, "Set concurrency level")
	flag.DurationVar(&timeout, "timeout", timeout, "Request timeout")
	flag.IntVar(&retries, "retries", retries, "Number of retries")
	flag.StringVar(&userAgent, "ua", userAgent, "User-Agent header")
	flag.BoolVar(&silent, "silent", silent, "Silent mode (only output results)")
	flag.BoolVar(&verbose, "v", verbose, "Verbose mode")
	flag.StringVar(&outputFile, "o", outputFile, "Output file")
	flag.StringVar(&methods, "m", methods, "Comma-separated HTTP methods to use")
	flag.StringVar(&singleURL, "url", "", "Test a single URL")
	flag.StringVar(&statusCodes, "status", "", "Comma-separated status codes to filter (e.g., 200,404)")
	flag.Parse()

	// Display banner
	printBanner()

	// Parse HTTP methods
	httpMethods := strings.Split(methods, ",")
	parseStatusCodes()

	if verbose {
		logInfo("Starting HTTP Method Tester Tool...")
		logInfo(fmt.Sprintf("Concurrency Level: %d, Timeout: %s, Retries: %d", concurrency, timeout, retries))
	}

	// Determine URLs to test
	var urls []string
	if singleURL != "" {
		urls = append(urls, singleURL)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	}

	// Concurrent handling of URLs
	var wg sync.WaitGroup
	resultChan := make(chan string, len(urls))
	doneChan := make(chan struct{})

	// Start result writer if output file specified
	go writeResults(resultChan, doneChan)

	// Start URL handlers
	sem := make(chan struct{}, concurrency)
	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{} // acquire semaphore
		go func(url string) {
			defer func() { <-sem }() // release semaphore
			handleURL(url, httpMethods, &wg, resultChan)
		}(url)
	}

	// Wait for all goroutines and result writing to complete
	wg.Wait()
	close(resultChan)
	<-doneChan

	if verbose {
		logSuccess("All requests completed.")
	}
}
