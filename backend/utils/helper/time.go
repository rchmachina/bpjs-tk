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

	// Generate a random start date within a random range (e.g., 5 years in the past to 5 years in the future)
	randomDaysStart := rand.Intn(365*10) - 365*5 // Range: -5 years to +5 years
	startDate := time.Now().AddDate(0, 0, randomDaysStart)

	// Generate a random end date that is after the start date
	// Randomize whether the end date difference is in days, months, or years
	randomYears := rand.Intn(5)  // Random years to add (up to 5 years)
	randomMonths := rand.Intn(12) // Random months to add (up to 12 months)
	randomDaysEnd := rand.Intn(31) // Random days to add (up to 31 days)

	endDate := startDate.AddDate(randomYears, randomMonths, randomDaysEnd)

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
func GenerateMonthlyRanges(startDate, endDate time.Time) []dto.DateRange {
	var dateRanges []dto.DateRange

	// Iterate over each month between startDate and endDate
	current := startDate

	for current.Before(endDate) || current.Equal(endDate) {

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
