package handlers

import (
	"calendar/internal/logic"
	"calendar/internal/structs/cal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RespSuc struct {
	Result interface{} `json:"result,omitempty"`
}
type RespErr struct {
	Error *string `json:"error,omitempty"`
}

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		result := "Home is working"
		response := RespSuc{
			Result: result,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func NewEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Method not allowed")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		var data cal.Event
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Println("Failed to decode request body:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
			return
		}
		defer r.Body.Close()

		err := logic.AddEvent(&data)
		if err != nil {
			log.Println("Failed to add event:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to add event"})
			return
		}

		log.Printf("Event %d added successfully\n", data.ID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"result": "Event added successfully"})
		log.Println("Response sent successfully")

	}
}

func UpdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Method not allowed")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		var data cal.Event
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Println("Failed to decode request body:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
			return
		}

		defer r.Body.Close()

		err := logic.UpdateEvent(&data)
		if err != nil {
			log.Println("Failed to update event:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update event"})
			return
		}

		log.Printf("Event %d updated successfully\n", data.ID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("Event %d updated successfully", data.ID)})
		log.Println("Response sent successfully")
	}
}

func DeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Method not allowed")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		var data cal.Event
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Println("Failed to decode request body:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
			return
		}
		defer r.Body.Close()

		if err := logic.DeleteEvent(&data); err != nil {
			log.Println("Failed to delete event:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete event"})
			return
		}

		log.Printf("Event %d deleted successfully\n", data.ID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("Event %d deleted successfully", data.ID)})
		log.Println("Response sent successfully")
	}
}

func GetEventDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Method not allowed")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		userIDStr := r.URL.Query().Get("user_id")
		dateStr := r.URL.Query().Get("date")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			log.Println("Failed to convert userID to int:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid userID"})
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Println("Failed to parse date:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}

		events, err := logic.GetEventsForDay(userID, date)
		if err != nil {
			log.Println("Failed to get events for day:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get events for day"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string][]*cal.Event{"result": events})
		log.Println("Response sent successfully")
	}
}

func GetEventWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Method not allowed")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		userIDStr := r.URL.Query().Get("user_id")
		dateStr := r.URL.Query().Get("date")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			log.Println("Failed to convert userID to int:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid userID"})
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Println("Failed to parse date:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}

		events, err := logic.GetEventsForWeek(userID, date)
		if err != nil {
			log.Println("Failed to get events for week:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get events for week"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string][]*cal.Event{"result": events})
		log.Println("Response sent successfully")
	}
}

func GetEventMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Method not allowed")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		userIDStr := r.URL.Query().Get("user_id")
		dateStr := r.URL.Query().Get("date")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			log.Println("Failed to convert userID to int:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid userID"})
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Println("Failed to parse date:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}

		events, err := logic.GetEventsForMonth(userID, date)
		if err != nil {
			log.Println("Failed to get events for month:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get events for month"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string][]*cal.Event{"result": events})
		log.Println("Response sent successfully")
	}
}
