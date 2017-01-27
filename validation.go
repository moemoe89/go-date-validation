package validation

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func StringToInt(s string)(*int,error){

	n,err := strconv.Atoi(s)
	if err != nil {
		return nil,err
	}
	return &n,nil
}

func Validation(date string)(bool,error){

	s := strings.Split(date,"-")
	if len(s) == 3 {

		yearSplit,monthSplit,daySplit := s[0],s[1],s[2]
		yearInt,err := StringToInt(yearSplit)
		if err != nil {
			return false,err
		}

		monthInt,err := StringToInt(monthSplit)
		if err != nil {
			return false,err
		}

		dayInt,err := StringToInt(daySplit)
		if err != nil {
			return false,err
		}

		if *yearInt == 0 {
			return false,errors.New("The year shouldn't be 0")
		}

		if *monthInt > 12 {
			return false,errors.New("The month shouldn't be over 12")
		}

		if *dayInt > 31 {
			return false,errors.New("The day shouldn't be over 31")
		}

		layoutFormat := "2006-01-02"
		_,err = time.Parse(layoutFormat,date)
		if err != nil {
			return false,err
		}

	} else {
		return false,errors.New("The date format should be yyyy-mm-dd")
	}

	return true,nil
}