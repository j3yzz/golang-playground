package main

import "fmt"

type inputError struct {
	message      string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func main() {
	err := validate("amir", "")
	if err != nil {
		fmt.Println(err.(*inputError))
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing field is %s\n", err.getMissingField())
		}
	}

}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{
			message:      "Name is required.",
			missingField: "name",
		}
	}

	if gender == "" {
		return &inputError{
			message:      "Gender is required.",
			missingField: "gender",
		}
	}

	return nil
}
