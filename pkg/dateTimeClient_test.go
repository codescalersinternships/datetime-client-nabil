package datetimeclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestGetCurrentDate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path == "/datetime" && r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("2021-12-12 12:12:12"))
			assertNotError(t, err)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	t.Run("test valid date case", func(t *testing.T) {
		myClient := NewClient(server.URL, time.Duration(1)*time.Second)
		got, err := myClient.GetCurrentDate()
		assertNotError(t, err)
		want := currDateRes{"2021-12-12 12:12:12"}
		assertDate(t, got, want)
	})
	t.Run("test json date case", func(t *testing.T) {
		myClient := NewClient(server.URL, time.Duration(1)*time.Second)
		got, err := myClient.GetCurrentDate()
		assertNotError(t, err)
		var resParsed string
		err = json.Unmarshal([]byte(strconv.Quote(got.Date)), &resParsed)
		assertNotError(t, err)
		want := "2021-12-12 12:12:12"
		assertDateStrings(t, resParsed, want)

	})
	t.Run("invalid url", func(t *testing.T) {
		myClient := NewClient(server.URL+"/wrong", time.Duration(1)*time.Second)
		_, err := myClient.GetCurrentDate()
		assertError(t, err)

	})
	// t.Run("verify timeout constraint", func(t *testing.T) {
	// 	myClient := NewClient("http://localhost:8090", time.Duration(1)*time.Second)
	// 	_, err := myClient.GetCurrentDate()
	// 	assertNotError(t, err)
	// 	time.Sleep(time.Duration(3) * time.Second)
	// 	if err == nil {
	// 		t.Errorf("expected timeout error but got nil")
	// 	}
	// })

}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("error isn't nill error: %q", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("error is nill error: %q", err)
	}
}

func assertDate(t *testing.T, got, want currDateRes) {
	t.Helper()
	gotDate := got.Date
	wantDate := want.Date

	if gotDate[:len(gotDate)-2] != wantDate[:len(wantDate)-2] {
		t.Errorf("got %q, want %q", got, want)
		return
	}
	gotSec, err := strconv.Atoi(gotDate[len(gotDate)-2:])
	if err != nil {
		t.Errorf("can't parse seconds %q", err.Error())
	}
	wantSec, err := strconv.Atoi(wantDate[len(wantDate)-2:])
	if err != nil {
		t.Errorf("can't parse seconds %q", err.Error())
	}
	if gotSec <= wantSec-2 || gotSec >= wantSec+1 {
		t.Errorf("got %q, want %q", gotDate, wantDate)
	}
}

func assertDateStrings(t *testing.T, got, want string) {
	t.Helper()
	if got[:len(got)-2] != want[:len(want)-2] {
		t.Errorf("got %q, want %q", got, want)
		return
	}
	gotSec, err := strconv.Atoi(got[len(got)-2:])
	if err != nil {
		t.Errorf("can't parse seconds %q", err.Error())
	}
	wantSec, err := strconv.Atoi(want[len(want)-2:])
	if err != nil {
		t.Errorf("can't parse seconds %q", err.Error())
	}
	if gotSec <= wantSec-2 || gotSec >= wantSec+1 {
		t.Errorf("got %q, want %q", got, want)
	}
}
