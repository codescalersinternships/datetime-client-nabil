package datetimeclient

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestGetCurrentDate(t *testing.T) {
	t.Run("test valid date case", func(t *testing.T) {
		myClient := NewClient("http://localhost:8090", time.Duration(1)*time.Second)
		got, err := myClient.GetCurrentDate()
		assertNotError(t, err)
		want := currDateRes{time.Now().Format("2006-01-02 15:04:05")}
		assertDate(t, got, want)
	})
	t.Run("test json date case", func(t *testing.T) {
		myClient := NewClient("http://localhost:8090", time.Duration(1)*time.Second)
		got, err := myClient.GetCurrentDate()
		assertNotError(t, err)
		var resParsed string
		err = json.Unmarshal([]byte(strconv.Quote(got.Date)), &resParsed)
		assertNotError(t, err)
		want := time.Now().Format("2006-01-02 15:04:05")
		assertDateStrings(t, resParsed, want)

	})

}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("error isn't nill error: %q", err)
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
