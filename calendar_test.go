package gofin

import (
	"testing"
	"time"
)

func TestIsTradingDayUS(t *testing.T) {
	calendar, _ := NewTradingCalendarUS()

	// Test again holiday Juneteenth
	testTime := time.Date(2023, 6, 19, 10, 0, 0, 0, time.UTC)
	if calendar.IsTradingDay(testTime) {
		t.Errorf("IsTradingDay() failed, expected false, got true")
	}

	// Check saturday
	testTime = time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC)
	if calendar.IsTradingDay(testTime) {
		t.Errorf("IsTradingDay() failed, expected false, got true")
	}
	// Check sunday
	testTime = time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC)
	if calendar.IsTradingDay(testTime) {
		t.Errorf("IsTradingDay() failed, expected false, got true")
	}
	// Check regular day
	testTime = time.Date(2023, 6, 20, 10, 0, 0, 0, time.UTC)
	if !calendar.IsTradingDay(testTime) {
		t.Errorf("IsTradingDay() failed, expected true, got false")
	}
}

func TestNextDayOnOpen(t *testing.T) {
	calendar, _ := NewTradingCalendarUS()

	location, _ := time.LoadLocation("Europe/Berlin")

	testTime := time.Date(2023, 6, 28, 0, 0, 0, 0, location)
	nextDayWindow := calendar.NextDayOnOpen(testTime)

	if nextDayWindow.Start.Day() != 29 {
		t.Errorf("NextDayOnClose() Start failed, expected 29, got %d", nextDayWindow.Start.Day())
	}

	if nextDayWindow.End.Day() != 29 {
		t.Errorf("NextDayOnClose() End failed, expected 29, got %d", nextDayWindow.End.Day())
	}

	if nextDayWindow.Location != location {
		t.Errorf("NextDayOnClose() Location failed, expected %s, got %s", location, nextDayWindow.Location)
	}

	// compare starting hour and minute
	if calendar.OnOpen.Start.Hour() != nextDayWindow.Start.In(calendar.OnOpen.Location).Hour() {
		t.Errorf("NextDayOnClose() Start Hour failed, expected %d, got %d", calendar.OnOpen.Start.Hour(),
			nextDayWindow.Start.In(calendar.OnOpen.Location).Hour())
	}

	if calendar.OnOpen.Start.Minute() != nextDayWindow.Start.In(calendar.OnOpen.Location).Minute() {
		t.Errorf("NextDayOnClose() Start Minute failed, expected %d, got %d", calendar.OnOpen.Start.Minute(),
			nextDayWindow.Start.In(calendar.OnOpen.Location).Minute())
	}

	// compare ending hour and minute
	if calendar.OnOpen.End.Hour() != nextDayWindow.End.In(calendar.OnOpen.Location).Hour() {
		t.Errorf("NextDayOnClose() End Hour failed, expected %d, got %d", calendar.OnOpen.End.Hour(),
			nextDayWindow.End.In(calendar.OnOpen.Location).Hour())
	}

	if calendar.OnOpen.End.Minute() != nextDayWindow.End.In(calendar.OnOpen.Location).Minute() {
		t.Errorf("NextDayOnClose() End Minute failed, expected %d, got %d", calendar.OnOpen.End.Minute(),
			nextDayWindow.End.In(calendar.OnOpen.Location).Minute())
	}
}

func TestNextDayOnClose(t *testing.T) {
	calendar, _ := NewTradingCalendarUS()

	location, _ := time.LoadLocation("Europe/Berlin")

	testTime := time.Date(2023, 6, 28, 0, 0, 0, 0, location)
	nextDayWindow := calendar.NextDayOnClose(testTime)

	if nextDayWindow.Start.Day() != 29 {
		t.Errorf("NextDayOnClose() Start failed, expected 29, got %d", nextDayWindow.Start.Day())
	}

	if nextDayWindow.End.Day() != 29 {
		t.Errorf("NextDayOnClose() End failed, expected 29, got %d", nextDayWindow.End.Day())
	}

	if nextDayWindow.Location != location {
		t.Errorf("NextDayOnClose() Location failed, expected %s, got %s", location, nextDayWindow.Location)
	}

	// compare starting hour and minute
	if calendar.OnClose.Start.Hour() != nextDayWindow.Start.In(calendar.OnClose.Location).Hour() {
		t.Errorf("NextDayOnClose() Start Hour failed, expected %d, got %d", calendar.OnClose.Start.Hour(),
			nextDayWindow.Start.In(calendar.OnClose.Location).Hour())
	}

	if calendar.OnClose.Start.Minute() != nextDayWindow.Start.In(calendar.OnClose.Location).Minute() {
		t.Errorf("NextDayOnClose() Start Minute failed, expected %d, got %d", calendar.OnClose.Start.Minute(),
			nextDayWindow.Start.In(calendar.OnClose.Location).Minute())
	}

	// compare ending hour and minute
	if calendar.OnClose.End.Hour() != nextDayWindow.End.In(calendar.OnClose.Location).Hour() {
		t.Errorf("NextDayOnClose() End Hour failed, expected %d, got %d", calendar.OnClose.End.Hour(),
			nextDayWindow.End.In(calendar.OnClose.Location).Hour())
	}

	if calendar.OnClose.End.Minute() != nextDayWindow.End.In(calendar.OnClose.Location).Minute() {
		t.Errorf("NextDayOnClose() End Minute failed, expected %d, got %d", calendar.OnClose.End.Minute(),
			nextDayWindow.End.In(calendar.OnClose.Location).Minute())
	}
}

func TestNextBusinessDayUS(t *testing.T) {

	calendar, _ := NewTradingCalendarUS()
	// Normal Business Day
	normalDay := time.Date(2023, 6, 20, 21, 30, 0, 0, time.UTC)
	_, _, nextday := calendar.NextBusinessDay(normalDay).Date()
	if nextday != 21 {
		t.Errorf("NextBusinessDay() failed, expected 21, got %v", nextday)
	}

	// Friday
	normalDay = time.Date(2023, 6, 23, 21, 30, 0, 0, time.UTC)
	_, _, nextday = calendar.NextBusinessDay(normalDay).Date()
	if nextday != 26 {
		t.Errorf("NextBusinessDay() failed, expected 26, got %v", nextday)
	}

	// US Holiday on weekend
	normalDay = time.Date(2022, 12, 30, 21, 30, 0, 0, time.UTC)
	nextYear, nextMonth, nextDay := calendar.NextBusinessDay(normalDay).Date()
	if nextYear != 2023 || nextMonth != 1 || nextDay != 3 {
		t.Errorf("NextBusinessDay() failed, expected year: 2023/January/2, got %v/%v/%v", nextYear, nextMonth, nextDay)
	}

}

func TestIsOnCloseUS(t *testing.T) {
	calendar, _ := NewTradingCalendarUS()

	location, _ := time.LoadLocation("Europe/Berlin")

	testTime := time.Date(2023, 6, 18, 21, 35, 0, 0, location)
	if !calendar.IsOnClose(testTime) {
		t.Errorf("IsOnClose() failed, expected true, got false")
	}

	testTime = time.Date(2023, 6, 18, 21, 0, 0, 0, location)
	if calendar.IsOnClose(testTime) {
		t.Errorf("IsOnClose() failed, expected false, got true")
	}
}

func TestIsOnOpenUS(t *testing.T) {
	calendar, _ := NewTradingCalendarUS()

	location, _ := time.LoadLocation("Europe/Berlin")

	testTime := time.Date(2023, 6, 18, 15, 46, 0, 0, location)
	if !calendar.IsOnOpen(testTime) {
		t.Errorf("IsOnOpen() failed, expected true, got false")
	}

	testTime = time.Date(2023, 6, 18, 17, 0, 0, 0, location)
	if calendar.IsOnOpen(testTime) {
		t.Errorf("IsOnOpen() failed, expected false, got true")
	}
}
