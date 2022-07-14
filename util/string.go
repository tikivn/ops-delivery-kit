package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func ParseAcronym(str string, sep string) string {
	if strings.TrimSpace(str) == "" || sep == "" {
		return str
	}

	r := "[^" + sep + "]+"
	strArray := regexp.MustCompile(r).FindAllString(str, -1)

	var fisrtCharString string
	for _, item := range strArray {
		fisrtCharString = fisrtCharString + string(item[0])
	}

	return strings.ToUpper(fisrtCharString)
}

// Remove null Unicode from string because type jsonb in postgres does not accept (\u0000, ....)
func RemoveUnicodeNull(str string) string {
	return strings.Replace(str, "\x00", "", -1)
}
func InterfaceToFloat64(data interface{}) (float64, error) {
	var dataF float64

	if data == nil {
		return dataF, nil
	}

	if f, ok := data.(float64); ok { // yeah, JSON numbers are floats, gotcha!
		dataF = f
	} else if s, ok := data.(string); ok {
		if s == "" {
			return 0, nil
		}

		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return dataF, err
		}
		dataF = n
	} else {
		return dataF, errors.New(fmt.Sprintf("undefined type of value %v", data))
	}
	return dataF, nil
}

func ParseValidForHereMap(addr string) string {
	subdistrictLowerCase := strings.ToLower(strings.TrimSpace(addr))

	r := regexp.MustCompile(`\A(ph.*ng|qu.*n)\s+(\d+)\z`)

	res := r.FindStringSubmatch(subdistrictLowerCase)

	if len(res) <= 1 {
		return addr
	}

	name, err := strconv.Atoi(res[2])
	if err != nil {
		return addr
	}

	res[2] = strconv.Itoa(name)

	return strings.Join([]string{res[1], res[2]}, " ")
}

// return district_tiki_code, region_tiki_code
func ExtractTikiCode(wardTikiCode string) (string, string) {
	if len(wardTikiCode) < 5 {
		return "", ""
	}

	if len(wardTikiCode) < 8 {
		return "", wardTikiCode[0:5]
	}

	return wardTikiCode[0:8], wardTikiCode[0:5]
}
