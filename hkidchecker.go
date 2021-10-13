// hkidchecker checks if a Hong Kong ID number is valid or not
package hkidchecker

import (
	"regexp"
	"strconv"
	"strings"
)

// Validation is implemented as described on this page:
// https://computerterminal.blogspot.com/2013/02/hong-kong-id-formula-hkid-number-check.html

var hkidRegExp *regexp.Regexp

const hkidMaxLen = 11
const hkidDataLen = 8

func init() {
	hkidRegExp = regexp.MustCompile(`^[A-Z]{1,2}[0-9]{6}\([0-9A]{1}\)$`)
}

// CheckHKID validates a Hong Kong ID
func CheckHKID(hkid string) bool {
	hkid = strings.TrimSpace(hkid)
	if !hkidRegExp.MatchString(hkid) {
		return false
	}

	if len(hkid) != hkidMaxLen {
		hkid = " " + hkid
	}

	checkDigit := hkid[hkidDataLen:][1:2]
	sum := 0
	for i := 0; i < hkidDataLen; i++ {
		var num int
		ch := hkid[i]
		if ch == ' ' {
			num = 36
		} else if ch >= '0' && ch <= '9' {
			num = int(hkid[i] - '0')
		} else {
			num = int(hkid[i]-'A') + 10
		}
		sum += num * ((hkidDataLen + 1) - i)
	}

	var calcCheck string
	const modulo = 11
	reminder := sum % modulo
	if reminder == 0 {
		calcCheck = "0"
	} else if reminder == 1 {
		calcCheck = "A"
	} else {
		calcCheck = strconv.Itoa(modulo - reminder)
	}
	return calcCheck == checkDigit
}
