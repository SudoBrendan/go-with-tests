package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThemInMemory(t *testing.T) {
	// SETUP
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	name := "Pepper"
	numWins := 3
	for i := 0; i < numWins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), getPostPlayerScoreRequest(name))
	}

	// SUITE
	t.Run("get score", func(t *testing.T) {
		// Given
		response := httptest.NewRecorder()
		request := getGetPlayerScoreRequest(name)

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseCode(t, response, http.StatusOK)
		assertResponseBody(t, response, strconv.Itoa(numWins))
	})

	t.Run("get league", func(t *testing.T) {
		// Given
		response := httptest.NewRecorder()
		request := getGetLeagueRequest()
		wantLeague := []Player{
			{name, numWins},
		}

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseCode(t, response, http.StatusOK)
		assertResponseContentType(t, response, jsonContentType)
		assertResponseLeague(t, response, wantLeague)
	})
}
