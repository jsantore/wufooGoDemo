package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //import for side effects
	"log"
	"os"
)

//func main() {
//	myDatabase := OpenDataBase("./prototypeDB.db")
//	defer myDatabase.Close()
//	create_tables(myDatabase)
//}

func OpenDataBase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func create_tables(database *sql.DB) {
	createStatement1 := "CREATE TABLE IF NOT EXISTS WuFooData( " +
		"entryID INTEGER PRIMARY KEY," +
		"prefix TEXT NOT NULL," +
		"first_name TEXT NOT NULL," +
		"last_name TEXT NOT NULL," +
		"title TEXT," +
		"org TEXT," +
		"email TEXT," +
		"website TEXT," +
		"course_project BOOLEAN," +
		"guest_speaker BOOLEAN," +
		"site_visit BOOLEAN," +
		"job_shadow BOOLEAN," +
		"internship BOOLEAN," +
		"career_panel BOOLEAN," +
		"networking_event BOOLEAN," +
		"subject_area TEXT NOT NULL," +
		"description TEXT NOT NULL," +
		"funding BOOLEAN," +
		"created_date TEXT," +
		"created_by TEXT" +
		");"
	_, err := database.Exec(createStatement1)
	if err != nil {
		log.Println(err)
	}
}

func insertWufooData(database *sql.DB, wufooData []WuFooData) {
	insertStatement := "INSERT INTO WuFooData (prefix, first_name, last_name, title, org, email, website," +
		"course_project, guest_speaker, site_visit, job_shadow, internship, career_panel, networking_event, " +
		"subject_area, description, funding, created_date, created_by) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	preppedStatement, err := database.Prepare(insertStatement)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	for _, entry := range wufooData {

		var projectBool, speakerBool, visitBool, jobshadowBool, internshipBool, careerPanelBool, NetworkingEventBool, FundingBool bool
		if len(entry.CourseProject) > 0 {
			projectBool = true
		}
		if entry.Funding == "yes" {
			FundingBool = true
		}
		if len(entry.GuestSpeaker) > 0 {
			speakerBool = true
		}
		if len(entry.SiteVisit) > 0 {
			visitBool = true
		}
		if len(entry.JobShadow) > 0 {
			jobshadowBool = true
		}
		if len(entry.Internship) > 0 {
			internshipBool = true
		}
		if len(entry.CareerPanel) > 0 {
			careerPanelBool = true
		}
		if len(entry.NetworkingEvent) > 0 {
			NetworkingEventBool = true
		}
		_, err = preppedStatement.Exec(entry.Prefix, entry.FirstName, entry.LastName, entry.Title, entry.Org,
			entry.Email, entry.Website, projectBool, speakerBool, visitBool, jobshadowBool, internshipBool, careerPanelBool, NetworkingEventBool,
			entry.SubjectArea, entry.Description, FundingBool, entry.CreateDate, entry.CreatedBy)
		if err != nil {
			log.Println(err)
		}
	}
}

func getData() []WuFooData {
	db := OpenDataBase("prototypeDB.db")
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM WuFooData").Scan(&count)
	if err != nil {
		log.Println("Unable to get data")
	}
	data := make([]WuFooData, count)
	selectStatement := "SELECT * FROM WuFooData"
	wufooRows, err := db.Query(selectStatement)
	defer wufooRows.Close()
	rowNum := 0
	for wufooRows.Next() {
		currentItem := WuFooData{}
		wufooRows.Scan(&currentItem.EntryID, &currentItem.Prefix, &currentItem.FirstName, &currentItem.LastName, &currentItem.Title, &currentItem.Org,
			&currentItem.Email, &currentItem.Website, &currentItem.CourseProject, &currentItem.GuestSpeaker, &currentItem.SiteVisit, &currentItem.JobShadow,
			&currentItem.Internship, &currentItem.CareerPanel, &currentItem.NetworkingEvent, &currentItem.SubjectArea, &currentItem.Description,
			&currentItem.Funding, &currentItem.CreateDate, &currentItem.CreatedBy)
		data[rowNum] = currentItem
		rowNum++
	}
	return data
}
