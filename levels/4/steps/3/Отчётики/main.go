package main

import "math/rand"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	ReportID int
	Date     string
	User
}

func CreateReport(user User, reportDate string) Report {
	return Report{
		ReportID: rand.Int(),
		User:     user,
		Date:     reportDate,
	}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0, len(users))

	for _, u := range users {
		report := CreateReport(u, reportDate)
		reports = append(reports, report)
	}

	return reports
}
