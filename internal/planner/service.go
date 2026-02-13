package planner

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func daysInMonth(year int, month time.Month, loc *time.Location) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, loc).Day()
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	loc := now.Location()

	data := strings.Split(strings.TrimSpace(repeat), " ")
	if len(data) == 0 {
		return "", errors.New("Некорректный паттер repeat")
	}

	var period = strings.TrimSpace(data[0])

	dateFrom, err := time.Parse("20060102", dstart)

	if err != nil {
		return "", errors.New("Некорректный формат даты")
	}

	switch period {
	case "d":
		if len(data) < 2 {
			return "", errors.New("Некорректный формат даты")
		}

		rate := strings.TrimSpace(data[1])
		shift, err := strconv.Atoi(rate)
		if err != nil || shift > 365 {
			return "", errors.New("Некорректная дата")
		}

		dateFrom = dateFrom.AddDate(0, 0, shift)
		for !dateFrom.After(now) {
			dateFrom = dateFrom.AddDate(0, 0, shift)
		}

		return dateFrom.Format("20060102"), nil
	case "y":
		if dateFrom.Before(now) {
			date := time.Date(now.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, loc)
			if date.Before(now) || date.Equal(now) {
				date = date.AddDate(1, 0, 0)
			}
			return date.Format("20060102"), nil
		}
		return dateFrom.AddDate(1, 0, 0).Format("20060102"), nil
	case "w":
		if len(data) != 2 || strings.TrimSpace(data[1]) == "" {
			return "", errors.New("Некорректный формат repeat")
		}
		rate := strings.TrimSpace(data[1])

		var numbersOfWeekday []int
		days := strings.Split(rate, ",")
		for _, day := range days {
			shift, err := strconv.Atoi(day)
			if err != nil {
				return "", errors.New("Некорректная дата")
			}

			numbersOfWeekday = append(numbersOfWeekday, shift)
		}

		datesForAlert := make([]time.Time, len(numbersOfWeekday))
		for ind, numberOfWeekday := range numbersOfWeekday {
			date := now
			for {
				shiftedWeekday := int(date.Weekday())
				if shiftedWeekday == 0 {
					shiftedWeekday = 7
				} else {
					shiftedWeekday -= 1
				}

				if shiftedWeekday == numberOfWeekday && date.After(now) {
					datesForAlert[ind] = date
					break
				}
				date = date.AddDate(0, 0, 1)
			}
		}

		lowestDay := datesForAlert[0]

		for i := 1; i < len(datesForAlert); i++ {
			if datesForAlert[i].Before(lowestDay) {
				lowestDay = datesForAlert[i]
				break
			}
		}

		return lowestDay.Format("20060102"), nil
	case "m":
		numbersOfMonths := "1,2,3,4,5,6,7,8,9,10,11,12"
		if len(data) > 2 {
			numbersOfMonths = strings.TrimSpace(data[2])
		}

		monthOfNumberOfMonths := strings.Split(numbersOfMonths, ",")

		numberOfMounthsForAlert := make([]int, len(monthOfNumberOfMonths))
		for ind, month := range monthOfNumberOfMonths {
			shift, err := strconv.Atoi(month)
			if err != nil {
				return "", errors.New("Некорректная дата")
			}

			numberOfMounthsForAlert[ind] = shift
		}

		dayRate := strconv.Itoa(now.Day())
		if len(data) > 1 && data[1] != "x" {
			dayRate = strings.TrimSpace(data[1])
		}

		daysStingFromRate := strings.Split(dayRate, ",")
		daysFromRate := make([]int, len(daysStingFromRate))
		for ind, day := range daysStingFromRate {
			shift, err := strconv.Atoi(day)
			if err != nil {
				return "", errors.New("Некорректная дата")
			}

			daysFromRate[ind] = shift
		}

		datesForAlert := make([]time.Time, len(daysStingFromRate))

		for ind, day := range daysFromRate {
			for _, monthNumber := range numberOfMounthsForAlert {
				if daysInMonth(now.Year(), now.Month(), loc) < day {
					return "", errors.New("Некорректная дата")
				}

				year := now.Year()
				month := time.Month(monthNumber)
				normalizedDay := day
				if normalizedDay < 0 {
					normalizedDay = daysInMonth(year, month, loc) + day + 1
				}

				date := time.Date(year, month, day, 0, 0, 0, 0, loc)
				for date.Before(now) && !date.Equal(now) {
					date = date.AddDate(1, 0, 0)
					continue
				}
				datesForAlert[ind] = date
			}

		}

		lowestDay := datesForAlert[0]

		for i := 1; i < len(datesForAlert); i++ {
			if datesForAlert[i].Before(lowestDay) {
				lowestDay = datesForAlert[i]
				break
			}
		}

		return lowestDay.Format("20060102"), nil
	}
	return "", errors.New("Некорректный формат переноса")
}
