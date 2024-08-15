package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies Running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2024-08-01")
	deadpoolW := models.Movie{
		ID:          1,
		Title:       "Deadpool e Wolverine",
		ReleaseDate: rd,
		AgeRating:   "18+",
		RunTime:     120,
		Description: "This description is big",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, deadpoolW)

	rd2, _ := time.Parse("2006-01-02", "2024-08-08")
	borderlands := models.Movie{
		ID:          2,
		Title:       "Borderlands",
		ReleaseDate: rd2,
		AgeRating:   "12+",
		RunTime:     120,
		Description: "Worst Movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, borderlands)
	
	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
