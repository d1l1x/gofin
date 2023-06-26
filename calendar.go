package gofin

import (
	"fmt"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
	"time"
)

type TradingWindow struct {
	Start time.Time
	End   time.Time
}

type TradingCalendar struct {
	Calendar    *cal.BusinessCalendar
	OpenWindow  TradingWindow
	CloseWindow TradingWindow
}

// NewTradingCalendar creates a new TradingCalendar with specified opening and closing trading windows.
// It initializes a business calendar 'c' and adds US holidays to it.
// The function takes two arguments: 'OpenWindow' and 'CloseWindow', which are TradingWindow structs representing
// the opening and closing hours of the trading day, respectively.
// It returns a TradingCalendar struct.
func NewTradingCalendar(OpenWindow TradingWindow, CloseWindow TradingWindow) (TradingCalendar, error) {
	if OpenWindow.Start.Location() != OpenWindow.End.Location() {
		return TradingCalendar{}, fmt.Errorf("start and End of OpenWindow must have the same location")
	}
	if CloseWindow.Start.Location() != CloseWindow.End.Location() {
		return TradingCalendar{}, fmt.Errorf("start and End of CloseWindow must have the same location")
	}
	if OpenWindow.Start.Location() != CloseWindow.Start.Location() {
		return TradingCalendar{}, fmt.Errorf("OpenWindow and CloseWindow must have the same location")
	}
	c := cal.NewBusinessCalendar()
	// Add US Holidays
	c.AddHoliday(
		us.NewYear,
		us.MlkDay,
		us.PresidentsDay,
		us.MemorialDay,
		us.Juneteenth,
		us.IndependenceDay,
		us.LaborDay,
		us.ThanksgivingDay,
		us.ChristmasDay,
	)
	//c.SetWorkHours(DayStartHour, DayEndHour)
	return TradingCalendar{
		Calendar:    c,
		OpenWindow:  OpenWindow,
		CloseWindow: CloseWindow,
	}, nil
}

// IsTradingDay checks if the provided time 't' falls on a trading day according to the calendar 'c'.
// It returns true if 't' is a trading day, and false otherwise.
func (c *TradingCalendar) IsTradingDay(t time.Time) bool {
	return c.Calendar.IsWorkday(t)
}

// NextDayOnOpen Returns the trading window for the opening hours of the next trading day after the provided time "t".
func (c *TradingCalendar) NextDayOnOpen(t time.Time) (TradingWindow, error) {
	if t.Location() != c.OpenWindow.Start.Location() {
		return TradingWindow{}, fmt.Errorf("input's location and calendar's location must be the same")
	}
	nextYear, nextMonth, nextDay := c.NextBusinessDay(t).Date()
	startHour, startMinute, _ := c.OpenWindow.Start.Clock()
	endHour, endMinute, _ := c.OpenWindow.End.Clock()
	return TradingWindow{
		Start: time.Date(nextYear, nextMonth, nextDay, startHour, startMinute, 0, 0, t.Location()),
		End:   time.Date(nextYear, nextMonth, nextDay, endHour, endMinute, 0, 0, t.Location()),
	}, nil
}

// NextDayOnClose Returns the trading window for the closing hours of the next trading day after the provided time "t".
func (c *TradingCalendar) NextDayOnClose(t time.Time) (TradingWindow, error) {
	if t.Location() != c.CloseWindow.Start.Location() {
		return TradingWindow{}, fmt.Errorf("input's location and calendar's location must be the same")
	}
	nextYear, nextMonth, nextDay := c.NextBusinessDay(t).Date()
	startHour, startMinute, _ := c.CloseWindow.Start.Clock()
	endHour, endMinute, _ := c.CloseWindow.End.Clock()
	return TradingWindow{
		Start: time.Date(nextYear, nextMonth, nextDay, startHour, startMinute, 0, 0, t.Location()),
		End:   time.Date(nextYear, nextMonth, nextDay, endHour, endMinute, 0, 0, t.Location()),
	}, nil
}

// IsOnOpen checks if the provided time 't' falls within the opening hours of the trading day.
// It returns true if 't' is within the opening hours, and false otherwise.
func (c *TradingCalendar) IsOnOpen(t time.Time) bool {
	return t.After(c.OpenWindow.Start) && t.Before(c.OpenWindow.End)
}

// IsOnClose checks if the provided time 't' falls within the closing hours of the trading day.
// It returns true if 't' is within the closing hours, and false otherwise.
func (c *TradingCalendar) IsOnClose(t time.Time) bool {
	return t.After(c.CloseWindow.Start) && t.Before(c.CloseWindow.End)
}

// NextBusinessDay returns the next business day after the given time t.
// It checks if the given time t is a workday according to the TradingCalendar's calendar.
// If t is not a workday (it's a holiday or a weekend), it adds one day and checks again.
// It continues this process until it finds a workday, which it then returns.
// This function does not account for business hours; it only checks the date.
// If the given time t is already a workday, it will return the same day.
//
// Parameters:
// t : A time.Time value representing the date to start from.
//
// Returns:
// The next business day (as a time.Time value) after the given time t, or the same day if t
// is already a business day.
func (c *TradingCalendar) NextBusinessDay(t time.Time) time.Time {
	t = t.AddDate(0, 0, 1)
	for {
		switch {
		case !c.Calendar.IsWorkday(t):
			// If it's a holiday or weekend, add 1 day
			t = t.AddDate(0, 0, 1)
		default:
			return t
		}
	}
}
