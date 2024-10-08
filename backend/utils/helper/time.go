package helper

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/rchmachina/bpjs-tk/dto"
)

func CreateRandomDates(format string) (string, string) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random start date within a random range (e.g., 5 years in the past to the current date)
	randomDaysStart := rand.Intn(365*5) * -1 // Range: -5 years to today
	startDate := time.Now().AddDate(0, 0, randomDaysStart)

	// Generate a random end date that is no more than 5 months after the start date
	randomMonths := rand.Intn(5)        // Random months to add (up to 5 months)
	randomDaysEnd := rand.Intn(31)      // Random days to add (up to 31 days)

	// Add up to 5 months and random days to the start date to get the end date
	endDate := startDate.AddDate(0, randomMonths, randomDaysEnd)

	// Format the dates
	start := startDate.Format(format)
	end := endDate.Format(format)

	return start, end
}


func RandomSaldo(min, max float64) (float64, error) {
	if min >= max {
		return 0, fmt.Errorf("saldo max must be greater than min")
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random float value between min and max
	randomValue := min + rand.Float64()*(max-min)

	// Round the result to the nearest whole number and force `.00` in the last digits
	roundedValue := math.Floor(randomValue)

	return roundedValue, nil
}

// function to count total day we use

// Helper function to generate the monthly date ranges
func GenerateMonthlyRanges(startDate, endDate time.Time) []dto.DateRange {
	var dateRanges []dto.DateRange

	// Iterate over each month between startDate and endDate
	current := startDate

	for current.Before(endDate) || current.Equal(endDate) {
		// Get the first and last day of the current month
		firstOfMonth := FirstDayOfMonth(current)
		lastOfMonth := LastDayOfMonth(current)

		// Adjust the start date for the first month and end date for the last month
		start := firstOfMonth
		if current.Month() == startDate.Month() && current.Year() == startDate.Year() {
			start = startDate // Use the exact start date for the first month
		}

		end := lastOfMonth
		if current.Month() == endDate.Month() && current.Year() == endDate.Year() {
			end = endDate // Use the exact end date for the last month
		}

		// Append the DateRange to the array
		dateRanges = append(dateRanges, dto.DateRange{StartDate: start, EndDate: end})

		// Move to the next month
		current = firstOfMonth.AddDate(0, 1, 0)
	}

	return dateRanges
}
// count total day
func CountDays(startDate, endDate time.Time) int {
	return int(endDate.Sub(startDate).Hours()/24) + 1
}

// get first day of month
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// get last day of month
func LastDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, t.Location())
}
