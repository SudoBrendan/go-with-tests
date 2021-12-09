package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

// STUBS

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func NewStubPlayerStore(scores map[string]int, winCalls []string, league []Player) *StubPlayerStore {
	return &StubPlayerStore{
		scores:   scores,
		winCalls: winCalls,
		league:   league,
	}
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

type StubBlindAlert struct {
	scheduledAt time.Duration
	amount      int
}

func NewStubBlindAlert(duration time.Duration, amount int) StubBlindAlert {
	return StubBlindAlert{
		scheduledAt: duration,
		amount:      amount,
	}
}

func (s *StubBlindAlert) GetScheduledAt() time.Duration { return s.scheduledAt }
func (s *StubBlindAlert) GetAmount() int                { return s.amount }

type SpyBlindAlerter struct {
	alerts []StubBlindAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, StubBlindAlert{
		scheduledAt: duration,
		amount:      amount,
	})
}

func (s *SpyBlindAlerter) GetAlerts() []StubBlindAlert { return s.alerts }

// HELPERS

func GetGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func GetPostPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func GetGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func AssertResponseBody(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Body.String()
	if got != want {
		t.Errorf("got response body %q, wanted %q", got, want)
	}
}

func AssertResponseCode(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Code
	if got != want {
		t.Errorf("got response code %d, wanted %d", got, want)
	}
}

func AssertPlayerStoreWinCalls(t testing.TB, store StubPlayerStore, want []string) {
	t.Helper()
	got := store.winCalls
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got winCalls %+v, wanted %+v", got, want)
	}
}

func AssertResponseLeague(t testing.TB, response *httptest.ResponseRecorder, want []Player) {
	t.Helper()
	got := GetLeagueFromResponse(t, response.Body)
	AssertLeague(t, got, want)
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse league response from server %q into []Player: %v", body, err)
	}
	return
}

func AssertLeague(t testing.TB, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func AssertResponseContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Result().Header.Get("content-type")
	if got != want {
		t.Errorf("response did not have correct content-type, got %q want %q", got, want)
	}
}

func CreateTempFileStore(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func AssertScoresEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertBlindsEqual(t testing.TB, got, want StubBlindAlert) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
