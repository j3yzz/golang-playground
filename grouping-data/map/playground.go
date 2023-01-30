package main

import "fmt"

func main() {
	m := map[string]int {
		"Abbas Boa'azar": 32,
		"Asqar SagDast": 27,
	}

	fmt.Println(m["Abbas Boa'azar"])

	v, ok := m["Sakine"]
	fmt.Println(v, ok)

	m["mamali"] = 18
	if v, ok := m["Sakine"]; ok {
		fmt.Println(v)
	}

	if v, ok := m["mamali"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "mamali")
	if _, ok := m["mamali"]; !ok {
		fmt.Println("mamali is deleted.")
	}

	delete(m, "zahra")
}
