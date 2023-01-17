package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data []WuFooData
var (
	prefixDisplay     = widget.NewEntry()
	fNameDisplay      = widget.NewEntry()
	lNameDisplay      = widget.NewEntry()
	titleDisplay      = widget.NewEntry()
	orgDisplay        = widget.NewEntry()
	emailDisplay      = widget.NewEntry()
	descDisplay       = widget.NewEntry()
	isProjectDisplay  = widget.NewCheck("Course Project", func(bool) {})
	isGuestDisplay    = widget.NewCheck("Guest Speaker", func(bool) {})
	isSiteVisit       = widget.NewCheck("Site Visit", func(b bool) {})
	isJobShadow       = widget.NewCheck("Job Shadow", func(b bool) {})
	isInternShip      = widget.NewCheck("Internship", func(b bool) {})
	isCareerPanel     = widget.NewCheck("Career Panel", func(b bool) {})
	isNetworkingEvent = widget.NewCheck("Networking Event", func(b bool) {})
)

func windowMain() {
	displayApp := app.New()
	mainWindow := displayApp.NewWindow("Cubes Project List")
	displayList := createList()
	displayList.Resize(fyne.NewSize(250, 800))
	dataPanel := makeDataPanel()
	windowContainer := container.NewGridWithRows(1)
	windowContainer.Add(displayList)
	windowContainer.Add(dataPanel)
	windowContainer.Resize(fyne.NewSize(1000, 800))
	mainWindow.SetContent(windowContainer)
	mainWindow.Resize(fyne.NewSize(1000, 800))
	mainWindow.ShowAndRun()
}

func makeDataPanel() *fyne.Container {

	//four columns, two for labels and two for widgets
	dataPanel := container.New(layout.NewGridLayout(4))
	//first row
	leftLabel := widget.NewLabel("Position:")
	rightLabel := widget.NewLabel("Prefix:")
	titleDisplay.SetText("                                           .")
	titleDisplay.Resize(fyne.NewSize(400, 100))
	prefixDisplay.SetText("                                          .")
	dataPanel.Add(leftLabel)
	dataPanel.Add(titleDisplay)
	dataPanel.Add(rightLabel)
	dataPanel.Add(prefixDisplay)
	//Second row
	leftLabel = widget.NewLabel("First Name:")
	rightLabel = widget.NewLabel("Last Name:")
	dataPanel.Add(leftLabel)
	dataPanel.Add(fNameDisplay)
	dataPanel.Add(rightLabel)
	dataPanel.Add(lNameDisplay)
	//Third Row
	leftLabel = widget.NewLabel("Organization:")
	rightLabel = widget.NewLabel("Email:")
	dataPanel.Add(leftLabel)
	dataPanel.Add(orgDisplay)
	dataPanel.Add(rightLabel)
	dataPanel.Add(emailDisplay)

	label := widget.NewLabel("CUBES Project Description:")
	descDisplay.MultiLine = true
	descDisplay.Wrapping = fyne.TextWrapBreak
	dataPanel.Resize(fyne.NewSize(750, 400))
	descriptionPanel := container.NewVBox(isProjectDisplay, isGuestDisplay, isSiteVisit, isJobShadow,
		isInternShip, isCareerPanel, isNetworkingEvent, label, descDisplay)
	completeDataPanel := container.NewVBox(dataPanel, descriptionPanel)
	descriptionPanel.Resize(fyne.NewSize(750, 300))
	completeDataPanel.Resize(fyne.NewSize(750, 800))
	return completeDataPanel
}

func createList() *widget.List {
	data = getData()
	listOfCubesProjects := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("No Data Yet.................................................")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			displayText := fmt.Sprintf("%s : %s", data[i].Org, data[i].SubjectArea)
			o.(*widget.Label).SetText(displayText)
		},
	)
	listOfCubesProjects.OnSelected = showData
	return listOfCubesProjects
}

func showData(id widget.ListItemID) {
	selectedData := data[id]
	prefixDisplay.SetText(selectedData.Prefix)
	fNameDisplay.SetText(selectedData.FirstName)
	lNameDisplay.SetText(selectedData.LastName)
	titleDisplay.SetText(selectedData.Title)
	orgDisplay.SetText(selectedData.Org)
	emailDisplay.SetText(selectedData.Email)
	descDisplay.SetText(selectedData.Description)
	if selectedData.CourseProject == "true" {
		isProjectDisplay.SetChecked(true)
	} else {
		isProjectDisplay.SetChecked(false)
	}
	if selectedData.GuestSpeaker == "true" {
		isGuestDisplay.SetChecked(true)
	} else {
		isGuestDisplay.SetChecked(false)
	}
	if selectedData.SiteVisit == "true" {
		isSiteVisit.SetChecked(true)
	} else {
		isSiteVisit.SetChecked(false)
	}
	if selectedData.JobShadow == "true" {
		isJobShadow.SetChecked(true)
	} else {
		isJobShadow.SetChecked(false)
	}
	if selectedData.Internship == "true" {
		isInternShip.SetChecked(true)
	} else {
		isInternShip.SetChecked(false)
	}
	if selectedData.CareerPanel == "true" {
		isCareerPanel.SetChecked(true)
	} else {
		isCareerPanel.SetChecked(false)
	}
	if selectedData.NetworkingEvent == "true" {
		isNetworkingEvent.SetChecked(true)
	} else {
		isNetworkingEvent.SetChecked(false)
	}
}
