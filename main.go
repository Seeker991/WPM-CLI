package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Word list
var words = []string{
	"hello", "great", "God", "good", "bad", "go", "predict", "life", "awesome", "catch", "lovely",
	"watch",
}

func main() {
	fmt.Println("=============================")
	fmt.Println("    I am Speed Typing Test     ")
	fmt.Println("=============================")
	fmt.Printf("Total words to type: %d\n", len(words))
	fmt.Println("Instructions:")
	fmt.Println("• Type each word exactly as shown")
	fmt.Println("• Press Enter after each word")
	fmt.Println("• Press Ctrl+C to exit anytime")
	fmt.Println("=============================")
	fmt.Print("Press Enter to start the test...")
	
	// press Enter
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Start the game
	stats := runTypingTest()

	// Display results
	displayResults(stats)
}

// function to  hold stats of a typing session
type gameStats struct {
	wordsTyped    int
	charsCorrect  int
	totalChars    int
	duration      time.Duration
	startTime     time.Time
}

// function for main game loop
func runTypingTest() gameStats {
	reader := bufio.NewReader(os.Stdin)
	stats := gameStats{
		startTime: time.Now(),
	}

	fmt.Printf("Test started at: %s\n", stats.startTime.Format("2006-01-02 15:04:05"))
	fmt.Println("-----------------------------")

	for i, word := range words {
		// Display current stats before each word
		elapsed := time.Since(stats.startTime)
		currentWPM := calculateWPM(stats.charsCorrect, elapsed)
		
		fmt.Printf("\n--- Word #%d of %d ---\n", i+1, len(words))
		fmt.Printf("Time: %.1fs | Current WPM: %.1f | Accuracy: %.1f%%\n", 
			elapsed.Seconds(), currentWPM, calculateAccuracy(stats.charsCorrect, stats.totalChars))
		fmt.Printf("Type: %s\n> ", word)
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nAn error occurred. Exiting.", err)
			os.Exit(1)
		}
		
		// Trim newline/space characters from input
		input = strings.TrimSpace(input)
		
		// Update stats
		stats.wordsTyped++
		stats.totalChars += len(word)
		
		// Count correct characters
		correctChars := 0
		for j := 0; j < len(input) && j < len(word); j++ {
			if input[j] == word[j] {
				correctChars++
			}
		}
		stats.charsCorrect += correctChars
		
		// Show immediate feedback
		wordAccuracy := (float64(correctChars) / float64(len(word))) * 100
		if input == word {
			fmt.Printf("✓ Correct! (%.0f%% accuracy)\n", wordAccuracy)
		} else {
			fmt.Printf("✗ Expected: %s | You typed: %s (%.0f%% accuracy)\n", word, input, wordAccuracy)
		}
	}
	
	stats.duration = time.Since(stats.startTime)
	return stats
}

func calculateWPM(correctChars int, elapsed time.Duration) float64 {
	minutes := elapsed.Minutes()
	if minutes == 0 || correctChars == 0 {
		return 0
	}
	return (float64(correctChars) / 5.0) / minutes
}

func calculateAccuracy(correctChars, totalChars int) float64 {
	if totalChars == 0 {
		return 0
	}
	return (float64(correctChars) / float64(totalChars)) * 100
}

func displayResults(stats gameStats) {
	finalWPM := calculateWPM(stats.charsCorrect, stats.duration)
	accuracy := calculateAccuracy(stats.charsCorrect, stats.totalChars)

	fmt.Println("\n=============================")
	fmt.Println("       TEST COMPLETE!        ")
	fmt.Println("=============================")
	fmt.Printf("Start time: %s\n", stats.startTime.Format("15:04:05"))
	fmt.Printf("End time: %s\n", stats.startTime.Add(stats.duration).Format("15:04:05"))
	fmt.Printf("Total time: %.2f seconds\n", stats.duration.Seconds())
	fmt.Printf("Words completed: %d/%d\n", stats.wordsTyped, len(words))
	fmt.Printf("Characters typed correctly: %d/%d\n", stats.charsCorrect, stats.totalChars)
	fmt.Printf("Final accuracy: %.2f%%\n", accuracy)
	fmt.Printf("Final WPM: %.2f\n", finalWPM)
	fmt.Println("=============================")
}