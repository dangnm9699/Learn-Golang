package main

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randMSISDN() string {
	a := strconv.Itoa(int(r.Int31n(1000000000)))
	return "84" + strings.Repeat("0", 9-len(a)) + a
}

func randIMSI() string {
	a := int(r.Int63n(10000000000))
	e := len(strconv.Itoa(a))
	return strconv.Itoa(45204*int(math.Pow10(e)) + a)
}

func randID() string {
	a := strconv.Itoa(int(r.Int31n(1000000000)))
	return strings.Repeat("0", 9-len(a)) + a
}

func randDOB() string {
	month := int(r.Int31n(12) + 1)
	year := int(r.Int31n(55) + 1950)
	day := 0
	if month == 2 {
		if isLeapYear(year) {
			day = int(r.Int31n(29) + 1)
		} else {
			day = int(r.Int31n(28) + 1)
		}
	} else if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		day = int(r.Int31n(31) + 1)
	} else {
		day = int(r.Int31n(30) + 1)
	}
	dStr := strconv.Itoa(day)
	if len(dStr) == 1 {
		dStr = "0" + dStr
	}
	mStr := strconv.Itoa(month)
	if len(mStr) == 1 {
		mStr = "0" + mStr
	}
	yStr := strconv.Itoa(year)
	return dStr + mStr + yStr
}

func isLeapYear(year int) bool {
	return (((year%4 == 0) && (year%100 != 0)) ||
		(year%400 == 0))
}
