package util

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	gh_ship_code_min_len = 14
)

// Dự đoán code là ordercode hay shipcode
func FuzzyDecision(code string, clientPrefix []string) (string, string) {
	for _, prefix := range clientPrefix {
		if prefix != "" && strings.HasPrefix(code, prefix) {
			return code, ""
		}
	}

	// Đoạn này thử parse code ra dạng số xem được không, nếu có thì có từ 14 ký tự trở lên hay không.
	// Mục đích là để response fe về khi code không được tìm thấy ở cả cột orderCode hay shipCode
	_, err := strconv.Atoi(code)
	if err != nil {
		match, _ := regexp.MatchString(`\d+-\d+`, code)
		if match {
			return code, ""
		}

		return "", code
	}

	if len(code) >= gh_ship_code_min_len {
		return "", code
	}

	return code, ""
}

func HasMultiPrefix(code string, prefixs []string) bool {
	for _, prefix := range prefixs {
		if strings.HasPrefix(code, prefix) && prefix != "" {
			return true
		}
	}

	return false
}
