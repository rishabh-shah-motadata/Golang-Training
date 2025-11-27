package forrange

import "fmt"

func ForRange() {
	a := []int{1, 2, 3, 4, 5}

	for _, v := range a {
		a[1] = 10
		fmt.Println(v)
	}

	fmt.Println("Final slice:", a)

	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	for k, v := range m {
		m["four"] = 4
		fmt.Println(k, v)
	}

	fmt.Println("Final map:", m)
}
