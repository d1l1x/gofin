package gofin

import (
	"testing"
	"time"
)

func TestNewTradingCalendar(t *testing.T) {
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 18, 9, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 18, 16, 0, 0, 0, time.UTC)

	// Create a new trading calendar
	calendar, _ := NewTradingCalendar(TradingWindow{Start: start, End: end}, TradingWindow{Start: start, End: end})

	// Check if the calendar was created correctly
	if calendar.OpenWindow.Start != start || calendar.OpenWindow.End != end {
		t.Errorf("NewTradingCalendar() failed, expected %v and %v, got %v and %v", start, end, calendar.OpenWindow.Start, calendar.OpenWindow.End)
	}
}

func TestNewTradingCalendarOpeningWindowWrongLocation(t *testing.T) {
	// Test Opening window
	start := time.Date(2023, 6, 18, 9, 30, 0, 0, time.UTC)

	location, _ := time.LoadLocation("America/New_York")
	end := time.Date(2023, 6, 18, 16, 0, 0, 0, location)

	_, err := NewTradingCalendar(TradingWindow{Start: start, End: end}, TradingWindow{})

	if err == nil {
		t.Errorf("NewTradingCalendar() expected an error for opening window, got nil")
	}
}

func TestNewTradingCalendarClosingWindowWrongLocation(t *testing.T) {
	// Test Closing window
	start := time.Date(2023, 6, 18, 9, 30, 0, 0, time.UTC)

	location, _ := time.LoadLocation("America/New_York")
	end := time.Date(2023, 6, 18, 16, 0, 0, 0, location)

	_, err := NewTradingCalendar(TradingWindow{}, TradingWindow{Start: start, End: end})

	if err == nil {
		t.Errorf("NewTradingCalendar() expected an error for closing window, got nil")
	}
}

func TestNewTradingCalendarOpeningClosingWindowWrongLocation(t *testing.T) {
	// Test Opening-Closing window
	startOpen := time.Date(2023, 6, 18, 9, 30, 0, 0, time.UTC)
	endOpen := time.Date(2023, 6, 18, 9, 30, 0, 0, time.UTC)

	location, _ := time.LoadLocation("America/New_York")
	startClose := time.Date(2023, 6, 18, 16, 0, 0, 0, location)
	endClose := time.Date(2023, 6, 18, 16, 0, 0, 0, location)

	_, err := NewTradingCalendar(TradingWindow{Start: startOpen, End: endOpen},
		TradingWindow{Start: startClose, End: endClose})

	if err == nil {
		t.Errorf("NewTradingCalendar() expected an error opening-closing window, got nil")
	}
}

func TestIsTradingDay(t *testing.T) {
	calendar, _ := NewTradingCalendar(TradingWindow{}, TradingWindow{})

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
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 19, 9, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 19, 16, 0, 0, 0, time.UTC)

	// Create a new trading calendar with just an Open window
	calendar, _ := NewTradingCalendar(TradingWindow{Start: start, End: end}, TradingWindow{})

	// Get the next trading window
	nextDayWindow, _ := calendar.NextDayOnOpen(start)

	// Check if the next trading window is correct
	if !nextDayWindow.Start.Equal(calendar.OpenWindow.Start.AddDate(0, 0, 1)) {
		t.Errorf("NextDayOnOpen() failed, Start stamp not correct")
	}
	if !nextDayWindow.End.Equal(calendar.OpenWindow.End.AddDate(0, 0, 1)) {
		t.Errorf("NextDayOnOpen() failed, End stamp not correct")
	}
}

func TestNextDayOnOpenWrongLocation(t *testing.T) {
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 19, 9, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 19, 16, 0, 0, 0, time.UTC)

	// Create a new trading calendar with just an Open window
	calendar, _ := NewTradingCalendar(TradingWindow{Start: start, End: end}, TradingWindow{})

	location, _ := time.LoadLocation("America/New_York")
	testTime := time.Date(2023, 6, 18, 10, 0, 0, 0, location)
	// Get the next trading window
	_, err := calendar.NextDayOnOpen(testTime)

	if err == nil {
		t.Errorf("NextDayOnOpen() expected an error, got nil")
	}
}

func TestNextDayOnClose(t *testing.T) {
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 19, 9, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 19, 16, 0, 0, 0, time.UTC)

	// Create a new trading calendar with just an Open window
	calendar, _ := NewTradingCalendar(TradingWindow{}, TradingWindow{Start: start, End: end})

	// Get the next trading window
	nextDayWindow, _ := calendar.NextDayOnClose(start)

	// Check if the next trading window is correct
	if !nextDayWindow.Start.Equal(calendar.CloseWindow.Start.AddDate(0, 0, 1)) {
		t.Errorf("NextDayOnOpen() failed, Start stamp not correct")
	}
	if !nextDayWindow.End.Equal(calendar.CloseWindow.End.AddDate(0, 0, 1)) {
		t.Errorf("NextDayOnOpen() failed, End stamp not correct")
	}

}

func TestNextDayOnCloseWrongLocation(t *testing.T) {
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 19, 9, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 19, 16, 0, 0, 0, time.UTC)

	// Create a new trading calendar with just an Open window
	calendar, _ := NewTradingCalendar(TradingWindow{Start: start, End: end}, TradingWindow{})

	location, _ := time.LoadLocation("America/New_York")
	testTime := time.Date(2023, 6, 18, 10, 0, 0, 0, location)
	// Get the next trading window
	_, err := calendar.NextDayOnClose(testTime)

	if err == nil {
		t.Errorf("NextDayOnOpen() expected an error, got nil")
	}
}

func TestNextBusinessDay(t *testing.T) {

	calendar, _ := NewTradingCalendar(TradingWindow{}, TradingWindow{})
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

func TestIsOnClose(t *testing.T) {
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 18, 21, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 18, 21, 45, 0, 0, time.UTC)

	// Create a new trading calendar
	calendar, _ := NewTradingCalendar(TradingWindow{}, TradingWindow{Start: start, End: end})

	testTime := time.Date(2023, 6, 18, 21, 35, 0, 0, time.UTC)
	if !calendar.IsOnClose(testTime) {
		t.Errorf("IsOnClose() failed, expected true, got false")
	}

	testTime = time.Date(2023, 6, 18, 21, 0, 0, 0, time.UTC)
	if calendar.IsOnClose(testTime) {
		t.Errorf("IsOnClose() failed, expected false, got true")
	}
}

func TestIsOnOpen(t *testing.T) {
	// Define a start and end time for the trading windows
	start := time.Date(2023, 6, 18, 15, 30, 0, 0, time.UTC)
	end := time.Date(2023, 6, 18, 16, 00, 0, 0, time.UTC)

	// Create a new trading calendar
	calendar, _ := NewTradingCalendar(TradingWindow{Start: start, End: end}, TradingWindow{})

	testTime := time.Date(2023, 6, 18, 15, 35, 0, 0, time.UTC)
	if !calendar.IsOnOpen(testTime) {
		t.Errorf("IsOnOpen() failed, expected true, got false")
	}

	testTime = time.Date(2023, 6, 18, 17, 0, 0, 0, time.UTC)
	if calendar.IsOnOpen(testTime) {
		t.Errorf("IsOnOpen() failed, expected false, got true")
	}
}
