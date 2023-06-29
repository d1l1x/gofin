package gofin

import (
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
	"time"
)

type TimeWindow struct {
	Start    time.Time
	End      time.Time
	Location *time.Location
}

func NewTimeWindow(startHour, startMinute, endHour, endMinute int, location *time.Location) TimeWindow {
	return TimeWindow{
		Start:    time.Date(0, 0, 0, startHour, startMinute, 0, 0, location),
		End:      time.Date(0, 0, 0, endHour, endMinute, 0, 0, location),
		Location: location,
	}
}

type TradingCalendar struct {
	Calendar     *cal.BusinessCalendar
	OnOpen       TimeWindow
	OnClose      TimeWindow
	TradingHours TimeWindow
}

var US = "America/New_York"

// TradingWindowUS creates a new TradingWindow struct representing the standard trading hours of
// AMEX, ARCA, BATS, NYSE, NASDAQ, NYSEARCA.
// Here are the standard trading hours for these exchanges (in Eastern Time):
//   - Pre-Market Trading Hours: 4:00 a.m. to 9:30 a.m.
//   - Regular Trading Hours: 9:30 a.m. to 4:00 p.m.
//   - After-Market Hours: 4:00 p.m. to 8:00 p.m.
func TradingWindowUS() (TimeWindow, error) {
	location, err := time.LoadLocation(US)
	if err != nil {
		return TimeWindow{}, err
	}
	tradingHours := NewTimeWindow(9, 30, 16, 0, location)
	return tradingHours, nil
}

func TradingWindowUSOnOpen() (TimeWindow, error) {
	location, err := time.LoadLocation(US)
	if err != nil {
		return TimeWindow{}, err
	}
	tradingHours := NewTimeWindow(9, 45, 10, 15, location)
	return tradingHours, nil
}

func TradingWindowUSOnClose() (TimeWindow, error) {
	location, err := time.LoadLocation(US)
	if err != nil {
		return TimeWindow{}, err
	}
	tradingHours := NewTimeWindow(15, 15, 15, 45, location)
	return tradingHours, nil
}

// NewTradingCalendarUS creates a new TradingCalendar with specified opening and closing trading windows.
// It initializes a business calendar 'c' and adds US holidays to it.
// The function takes two arguments: 'OnOpen' and 'OnClose', which are TradingWindow structs representing
// the opening and closing hours of the trading day, respectively.
// It returns a TradingCalendar struct.
func NewTradingCalendarUS() (TradingCalendar, error) {
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
	tradingWindow, err := TradingWindowUS()
	if err != nil {
		return TradingCalendar{}, err
	}
	onOpen, err := TradingWindowUSOnOpen()
	if err != nil {
		return TradingCalendar{}, err
	}
	onClose, err := TradingWindowUSOnClose()
	if err != nil {
		return TradingCalendar{}, err
	}
	return TradingCalendar{
		Calendar:     c,
		OnOpen:       onOpen,
		OnClose:      onClose,
		TradingHours: tradingWindow,
	}, nil
}

// IsTradingDay checks if the provided time 't' falls on a trading day according to the calendar 'c'.
// It returns true if 't' is a trading day, and false otherwise.
func (c *TradingCalendar) IsTradingDay(t time.Time) bool {
	return c.Calendar.IsWorkday(t)
}

// NextDayOnOpen Returns the trading window for the opening hours of the next trading day after the provided time "t".
func (c *TradingCalendar) NextDayOnOpen(t time.Time) TimeWindow {
	nextYear, nextMonth, nextDay := c.NextBusinessDay(t).Date()
	startHour, startMinute, _ := c.OnOpen.Start.Clock()
	endHour, endMinute, _ := c.OnOpen.End.Clock()

	start := time.Date(nextYear, nextMonth, nextDay, startHour, startMinute, 0, 0, c.OnOpen.Location).In(t.Location())
	end := time.Date(nextYear, nextMonth, nextDay, endHour, endMinute, 0, 0, c.OnOpen.Location).In(t.Location())

	return TimeWindow{Start: start, End: end, Location: t.Location()}
}

// NextDayOnClose Returns the trading window for the closing hours of the next trading day after the provided time "t".
func (c *TradingCalendar) NextDayOnClose(t time.Time) TimeWindow {

	nextYear, nextMonth, nextDay := c.NextBusinessDay(t).Date()
	startHour, startMinute, _ := c.OnClose.Start.Clock()
	endHour, endMinute, _ := c.OnClose.End.Clock()

	start := time.Date(nextYear, nextMonth, nextDay, startHour, startMinute, 0, 0, c.OnClose.Location).In(t.Location())
	end := time.Date(nextYear, nextMonth, nextDay, endHour, endMinute, 0, 0, c.OnClose.Location).In(t.Location())

	return TimeWindow{Start: start, End: end, Location: t.Location()}
}

// IsOnOpen checks if the provided time 't' falls within the opening hours of the trading day.
// It does so by converting the given time to the same location as the calendar.
// It returns true if 't' is within the opening hours, and false otherwise.
func (c *TradingCalendar) IsOnOpen(t time.Time) bool {
	// convert given time to the same location as the calendar
	targetLocation := c.OnOpen.Location
	givenTime := t.In(targetLocation)

	// setup time objects for comparison
	startHour, startMinute, _ := c.OnOpen.Start.Clock()
	endHour, endMinute, _ := c.OnOpen.End.Clock()

	start := time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), startHour, startMinute, 0, 0, targetLocation)
	end := time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), endHour, endMinute, 0, 0, targetLocation)

	return givenTime.After(start) && givenTime.Before(end)

}

// IsOnClose checks if the provided time 't' falls within the closing hours of the trading day.
// It does so by converting the given time to the same location as the calendar.
// It returns true if 't' is within the closing hours, and false otherwise.
func (c *TradingCalendar) IsOnClose(t time.Time) bool {
	// convert given time to the same location as the calendar
	targetLocation := c.OnClose.Location
	givenTime := t.In(targetLocation)

	// setup time objects for comparison
	startHour, startMinute, _ := c.OnClose.Start.Clock()
	endHour, endMinute, _ := c.OnClose.End.Clock()

	start := time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), startHour, startMinute, 0, 0, targetLocation)
	end := time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), endHour, endMinute, 0, 0, targetLocation)

	return givenTime.After(start) && givenTime.Before(end)
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
