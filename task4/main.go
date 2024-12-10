package main

import "fmt"

func sliceDifference(sl1, sl2 []string) []string {
	out := []string{}
	m := map[string]struct{}{}
	for _, str := range sl2 {
		m[str] = struct{}{}
	}
	for _, str := range sl1 {

		if _, ok := m[str]; !ok {
			out = append(out, str)
		}
	}
	return out
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := sliceDifference(slice1, slice2)
	fmt.Println(res)
}
