package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sudobrendan/gowithtests/app/internal/poker"
)

func TestGETPlayerScores(t *testing.T) {
	store := poker.NewStubPlayerStore(
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		[]poker.Player{},
	)
	server := poker.NewPlayerServer(store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		// Given
		request := poker.GetGetPlayerScoreRequest("Pepper")
		response := httptest.NewRecorder()
		wantBody := "20"
		wantCode := 200

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseBody(t, response, wantBody)
		poker.AssertResponseCode(t, response, wantCode)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// Given
		request := poker.GetGetPlayerScoreRequest("Floyd")
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseBody(t, response, "10")
		poker.AssertResponseCode(t, response, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		// Given
		request := poker.GetGetPlayerScoreRequest("fizz")
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseCode(t, response, http.StatusNotFound)
	})
}

func TestPOSTWin(t *testing.T) {
	store := poker.NewStubPlayerStore(
		map[string]int{},
		[]string{},
		[]poker.Player{},
	)
	server := poker.NewPlayerServer(store)

	t.Run("records win on POST", func(t *testing.T) {
		// Given
		request := poker.GetPostPlayerScoreRequest("Pepper")
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseCode(t, response, http.StatusAccepted)
		poker.AssertPlayerStoreWinCalls(t, *store, []string{"Pepper"})
	})
}

func TestGETLeague(t *testing.T) {

	t.Run("returns the league table as JSON", func(t *testing.T) {
		// Given
		request := poker.GetGetLeagueRequest()
		response := httptest.NewRecorder()
		wantLeague := []poker.Player{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}
		store := poker.NewStubPlayerStore(nil, nil, wantLeague)
		server := poker.NewPlayerServer(store)

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseCode(t, response, http.StatusOK)
		poker.AssertResponseContentType(t, response, poker.JsonContentType)
		poker.AssertResponseLeague(t, response, wantLeague)
	})
}
