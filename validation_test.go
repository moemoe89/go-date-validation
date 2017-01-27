package validation

import (
	"testing"
)

func TestStringToInt(t *testing.T) {

	s := "1"
	_,err := StringToInt(s)
	if err != nil {
		t.Errorf("Should return no error, got %s", err)
	}
}

func TestWrongStringToInt(t *testing.T) {

	s := "s"
	_,err := StringToInt(s)
	if err == nil {
		t.Errorf("Should return wrong string, got %s", err)
	}
}

func TestDateFormat(t *testing.T) {

	date := "2017-02-28"
	_,err := Validation(date)
	if err != nil {
		t.Errorf("Should return no error, got %s", err)
	}
}

func TestWrongDateFormat(t *testing.T) {

	date := "lorem-epsum"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong date format, got %s", err)
	}
}

func TestWrongNumberYearFormat(t *testing.T) {

	date := "s-12-12"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong year format, got %s", err)
	}
}


func TestWrongNumberMonthFormat(t *testing.T) {

	date := "2017-s-12"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong month format, got %s", err)
	}
}


func TestWrongNumberDayFormat(t *testing.T) {

	date := "2017-12-s"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong day format, got %s", err)
	}
}

func TestWrongYearFormat(t *testing.T) {

	date := "0-12-12"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong year format, got %s", err)
	}
}


func TestWrongMonthFormat(t *testing.T) {

	date := "2017-13-12"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong month format, got %s", err)
	}
}


func TestWrongDayFormat(t *testing.T) {

	date := "2017-12-32"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error wrong day format, got %s", err)
	}
}

func TestOutOfRangeDateFormat(t *testing.T) {

	date := "2017-02-29"
	_,err := Validation(date)
	if err == nil {
		t.Errorf("Should return error out of range date format, got %s", err)
	}
}