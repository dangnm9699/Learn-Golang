package main

import (
	"errors"
	"regexp"
)

func stringToModify(a, b string) string {
	if b == "" {
		return a
	}
	return b
}

func validate(req *Request) error {
	msisdnReg := regexp.MustCompile("84(\\d{9})")
	imsiReg := regexp.MustCompile("45204(\\d{0,10})")
	if len(req.Data.MSISDN) != 11 || len(req.Data.IMSI) > 15 || len(req.Data.IMSI) < 5 {
		return errors.New("wrong")
	}
	result := msisdnReg.FindStringSubmatch(req.Data.MSISDN)
	if len(result) != 2 {
		return errors.New("wrong")
	}
	result = imsiReg.FindStringSubmatch(req.Data.IMSI)
	if len(result) != 2 {
		return errors.New("wrong")
	}
	return nil
}
