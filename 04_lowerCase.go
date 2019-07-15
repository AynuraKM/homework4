package main

import "unicode"

const xiaomiId = "9hf0P6MBHM7OY6UOWm"
const expected = "9hf0_p6_m_b_h_m7_o_y6_u_o_wm"

func LowerCase(id string) string {
	newId := ""
	for _, s := range id {
		if unicode.IsUpper(s) {
			newId += "_"
			newId += string(unicode.ToLower(s))
		} else {
			newId += string(s)
		}
	}
	return newId
}
