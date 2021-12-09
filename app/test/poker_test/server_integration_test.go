package poker_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/sudobrendan/gowithtests/app/internal/poker"
)

func TestRecordingWinsAndRetrievingThemInMemory(t *testing.T) {
	// SETUP
	db, rmDb := poker.CreateTempFileStore(t, `[]`)
	defer rmDb()
	store, err := poker.NewFileSystemPlayerStore(db)
	poker.AssertNoError(t, err)
	server := poker.NewPlayerServer(store)
	name := "Pepper"
	numWins := 3
	for i := 0; i < numWins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), poker.GetPostPlayerScoreRequest(name))
	}

	// SUITE
	t.Run("get score", func(t *testing.T) {
		// Given
		response := httptest.NewRecorder()
		request := poker.GetGetPlayerScoreRequest(name)

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseCode(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response, strconv.Itoa(numWins))
	})

	t.Run("get league", func(t *testing.T) {
		// Given
		response := httptest.NewRecorder()
		request := poker.GetGetLeagueRequest()
		wantLeague := []poker.Player{
			{Name: name, Wins: numWins},
		}

		// When
		server.ServeHTTP(response, request)

		// Then
		poker.AssertResponseCode(t, response, http.StatusOK)
		poker.AssertResponseContentType(t, response, poker.JsonContentType)
		poker.AssertResponseLeague(t, response, wantLeague)
	})
}
