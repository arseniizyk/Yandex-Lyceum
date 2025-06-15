package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var days = map[string]string{
	"Monday":    "Понедельник",
	"Tuesday":   "Вторник",
	"Wednesday": "Среда",
	"Thursday":  "Четверг",
	"Friday":    "Пятница",
	"Saturday":  "Суббота",
	"Sunday":    "Воскресенье",
}

var ErrInvalidTime = errors.New("исправь свой ответ, а лучше ложись поспать")

func CheckCurrentDayOfTheWeek(answer string) bool {
	answer = strings.ToLower(answer)
	dayOfWeek := currentDayOfTheWeek()

	return strings.EqualFold(answer, dayOfWeek)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, ErrInvalidTime
	}
	time := dayOrNight()

	return strings.EqualFold(answer, time), nil
}

func currentDayOfTheWeek() string {
	now := TimeNow()
	dayOfWeek := now.Format("Monday")
	return days[dayOfWeek]
}

func dayOrNight() string {
	now := TimeNow()
	hours := now.Hour()
	if hours >= 22 || hours <= 10 {
		return "Ночь"
	}

	return "День"
}

func nextFriday() int {
	now := TimeNow()
	weekday := now.Weekday()
	daysUntilFriday := (time.Friday - weekday + 7) % 7
	return int(daysUntilFriday)
}
