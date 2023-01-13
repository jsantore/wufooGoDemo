package main

import (
	"fmt"
)

type WuFooData struct {
	EntryID         int
	Prefix          string `json:"Field2"`
	FirstName       string `json:"Field3"`
	LastName        string `json:"Field4"`
	Title           string `json:"Field6"`
	Org             string `json:"Field7"`
	Email           string `json:"Field8"`
	Website         string `json:"Field9"`
	CourseProject   string `json:"Field11"`
	GuestSpeaker    string `json:"Field12"`
	SiteVisit       string `json:"Field13"`
	JobShadow       string `json:"Field14"`
	Internship      string `json:"Field15"`
	CareerPanel     string `json:"Field16"`
	NetworkingEvent string `json:"Field17"`
	SubjectArea     string `json:"Field113"`
	Description     string `json:"Field112"`
	Funding         string `json:"Field115"`
	CreateDate      string `json:"DateCreated"`
	CreatedBy       string `json:"CreatedBy"`
}

type WufooResponse struct {
	Entries []WuFooData `json:"Entries"`
}

func (datum WuFooData) prettyPrint() {
	printableString := fmt.Sprintf("\nWuFoo Entry:\n%s %s\nArea: %s\nDescription: %s",
		datum.FirstName, datum.LastName, datum.SubjectArea, datum.Description)
	fmt.Println(printableString)
}
