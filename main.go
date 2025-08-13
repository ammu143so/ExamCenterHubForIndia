package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exam-center-assignment/internal/handler"
)

func main() {
	fmt.Println("=== Welcome to ExamCenterHub ===")
	fmt.Println("Indian Examination Center Assignment System")
	fmt.Println()

	examHandler := handler.NewExamCenterHandler()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n📋 Main Menu:")
		fmt.Println("1. Basic Exam Center Assignment")
		fmt.Println("2. Advanced Assignment with Preferences")
		fmt.Println("3. View Available Exam Types")
		fmt.Println("4. View Registration Summary")
		fmt.Println("5. Exit")
		fmt.Print("\nSelect an option (1-5): ")

		if !scanner.Scan() {
			fmt.Println("\nInput error. Exiting...")
			return
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			_ = examHandler.ProcessExamCenterAssignment()
		case "2":
			_ = examHandler.ProcessAdvancedExamAssignment()
		case "3":
			examHandler.DisplayExamTypes()
		case "4":
			examHandler.ShowRegistrationSummary()
		case "5":
			fmt.Println("\n👋 Thank you for using ExamCenterHub!")
			fmt.Println("Good luck with your exams! 🎯")
			return
		default:
			fmt.Println("\n❌ Invalid option. Please select 1-5.")
		}

		// Wait for user to press Enter before showing menu again
		if choice != "5" {
			fmt.Print("\nPress Enter to continue...")
			scanner.Scan()
		}
	}
} 