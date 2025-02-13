package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func parseFlags() (string, string, string) {
	filePath := flag.String("file", os.Getenv("LOG_ANALYZER_FILE"), "Path to the log file (required)")
	logLevel := flag.String("level", getEnv("LOG_ANALYZER_LEVEL", "info"), "Log level to analyze")
	outputPath := flag.String("output", os.Getenv("LOG_ANALYZER_OUTPUT"), "Output file for statistics (optional)")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Error: log file path is required")
		flag.Usage()
		os.Exit(1)
	}

	return *filePath, *logLevel, *outputPath
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func AnalyzeLog(filePath, logLevel string) (map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := make(map[string]int)
	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(fmt.Sprintf(`(?i)\b%s\b`, regexp.QuoteMeta(logLevel)))

	for scanner.Scan() {
		line := scanner.Text()
		if pattern.MatchString(line) {
			stats[logLevel]++
		}
	}

	scanErr := scanner.Err()
	if scanErr != nil {
		return nil, scanErr
	}

	return stats, nil
}

func WriteOutput(stats map[string]int, outputPath string) error {
	output := fmt.Sprintf("Log level statistics: %v\n", stats)

	if outputPath != "" {
		file, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, writeErr := file.WriteString(output)
		return writeErr
	}

	fmt.Print(output)
	return nil
}

func main() {
	filePath, logLevel, outputPath := parseFlags()
	stats, err := AnalyzeLog(filePath, logLevel)
	if err != nil {
		fmt.Printf("Error analyzing log: %v\n", err)
		os.Exit(1)
	}

	outputErr := WriteOutput(stats, outputPath)
	if outputErr != nil {
		fmt.Printf("Error writing output: %v\n", outputErr)
		os.Exit(1)
	}
}
