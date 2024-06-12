package requests

import (
	"calendar/internal/database/connection"
	"calendar/internal/structs/cal"
	"log"
	"testing"
	"time"
)

func TestNewEvent(t *testing.T) {
	store := &DBEventSotre{DB: connection.DB}

	dateString1 := "2024-06-10"
	dateString2 := "2024-06-13"
	dateString3 := "2024-06-27"
	dateFormat := "2006-01-02"
	date1, err := time.Parse(dateFormat, dateString1)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}
	date2, err := time.Parse(dateFormat, dateString2)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}
	date3, err := time.Parse(dateFormat, dateString3)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}

	testCases := []*cal.Event{
		{ID: 1,
			UserID: 1,
			Date:   date1,
			Title:  "Birthday",
			Body:   "Buy new PC"},
		{ID: 2,
			UserID: 1,
			Date:   date1,
			Title:  "Play pc",
			Body:   "Download Dota2"},
		{ID: 3,
			UserID: 1,
			Date:   date2,
			Title:  "Hang out",
			Body:   "Hang out with friends"},
		{ID: 4,
			UserID: 1,
			Date:   date3,
			Title:  "Die",
			Body:   "roffle, don't die"},
	}

	for _, testCase := range testCases {
		err = store.NewEvent(testCase)
		if err != nil {
			t.Fatalf("Error: %v\n", err)
		}
	}
}

func TestUpdateEvent(t *testing.T) {
	store := &DBEventSotre{DB: connection.DB}

	dateString := "2024-06-10"
	dateFormat := "2006-01-02"
	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}

	testCase := &cal.Event{
		ID:     1,
		UserID: 1,
		Date:   date,
		Title:  "Colonization",
		Body:   "Colonize Mars by 2050",
	}

	err = store.UpdateEvent(testCase)
	if err != nil {
		t.Fatalf("Error: %v\n", err)
	}
}

func TestGetEventsForDay(t *testing.T) {
	store := &DBEventSotre{DB: connection.DB}

	dateString := "2024-06-10"
	dateFormat := "2006-01-02"
	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}

	userId := 1

	arr, err := store.GetEventsForDay(userId, date)
	if err != nil {
		t.Fatalf("Error: %v\n", err)
	}

	for _, value := range arr {
		log.Printf("event_id: %d, user_id: %d, date: %v, title: %s, body: %s\n", value.ID, value.UserID, value.Date, value.Title, value.Body)
	}
}

func TestGetEventsForWeek(t *testing.T) {
	store := &DBEventSotre{DB: connection.DB}

	dateString := "2024-06-10"
	dateFormat := "2006-01-02"
	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}

	userId := 1

	arr, err := store.GetEventsForWeek(userId, date)
	if err != nil {
		t.Fatalf("Error: %v\n", err)
	}

	for _, value := range arr {
		log.Printf("event_id: %d, user_id: %d, date: %v, title: %s, body: %s\n", value.ID, value.UserID, value.Date, value.Title, value.Body)
	}
}

func TestGetEventsForMonth(t *testing.T) {
	store := &DBEventSotre{DB: connection.DB}

	dateString := "2024-06-10"
	dateFormat := "2006-01-02"
	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		t.Fatalf("Error parsing date: %v\n", err)
		return
	}

	userId := 1

	arr, err := store.GetEventsForMonth(userId, date)
	if err != nil {
		t.Fatalf("Error: %v\n", err)
	}

	for _, value := range arr {
		log.Printf("event_id: %d, user_id: %d, date: %v, title: %s, body: %s\n", value.ID, value.UserID, value.Date, value.Title, value.Body)
	}
}

func TestDelete(t *testing.T) {
	store := &DBEventSotre{DB: connection.DB}
	events := [4]int{1, 2, 3, 4}

	for _, eventID := range events {
		err := store.DeleteEvent(eventID)
		if err != nil {
			t.Fatalf("Error: %v\n", err)
		}
	}
}
