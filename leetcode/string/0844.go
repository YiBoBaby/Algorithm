package main

import (
	"strings"
)

func backspaceCompare(S string, T string) bool {
	funR := func(str string) []string {
		rs := make([]string, 0, len(str))
		for i := 0; i < len(str); i++ {
			tmp := str[i : i+1]
			if tmp == "#" {
				if len(rs) != 0 {
					rs = rs[0 : len(rs)-1]
				}
			} else {
				rs = append(rs, tmp)
			}
		}
		return rs
	}
	rs := strings.Join(funR(S), "")
	rt := strings.Join(funR(T), "")
	return rs == rt
}
