package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// STUBS

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

// TESTS

func TestGETPlayerScores(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		[]Player{},
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		// Given
		request := getGetPlayerScoreRequest("Pepper")
		response := httptest.NewRecorder()
		wantBody := "20"
		wantCode := 200

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseBody(t, response, wantBody)
		assertResponseCode(t, response, wantCode)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// Given
		request := getGetPlayerScoreRequest("Floyd")
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseBody(t, response, "10")
		assertResponseCode(t, response, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		// Given
		request := getGetPlayerScoreRequest("fizz")
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseCode(t, response, http.StatusNotFound)
	})
}

func TestPOSTWin(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		[]Player{},
	}
	server := NewPlayerServer(&store)

	t.Run("records win on POST", func(t *testing.T) {
		// Given
		request := getPostPlayerScoreRequest("Pepper")
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseCode(t, response, http.StatusAccepted)
		assertPlayerStoreWinCalls(t, store, []string{"Pepper"})
	})
}

func TestGETLeague(t *testing.T) {

	t.Run("returns the league table as JSON", func(t *testing.T) {
		// Given
		request := getGetLeagueRequest()
		response := httptest.NewRecorder()
		wantLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{nil, nil, wantLeague}
		server := NewPlayerServer(&store)

		// When
		server.ServeHTTP(response, request)

		// Then
		assertResponseCode(t, response, http.StatusOK)
		assertResponseContentType(t, response, jsonContentType)
		assertResponseLeague(t, response, wantLeague)
	})
}

// HELPERS

func getGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func getPostPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func getGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertResponseBody(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Body.String()
	if got != want {
		t.Errorf("got response body %q, wanted %q", got, want)
	}
}

func assertResponseCode(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Code
	if got != want {
		t.Errorf("got response code %d, wanted %d", got, want)
	}
}

func assertPlayerStoreWinCalls(t testing.TB, store StubPlayerStore, want []string) {
	t.Helper()
	got := store.winCalls
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got winCalls %+v, wanted %+v", got, want)
	}
}

func assertResponseLeague(t testing.TB, response *httptest.ResponseRecorder, want []Player) {
	t.Helper()
	got := getLeagueFromResponse(t, response.Body)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse league response from server %q into []Player: %v", body, err)
	}
	return
}

func assertResponseContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Result().Header.Get("content-type")
	if got != want {
		t.Errorf("response did not have correct content-type, got %q want %q", got, want)
	}
}
