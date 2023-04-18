package restapi

import "fmt"

func ExecExample() error {
	client := NewAPIClient("yoonjeong", "yj-password")
	code, err := client.GetGoogle()
	if err != nil {
		return err
	}

	fmt.Printf("Result of GetGoogle: %d\n", code)
	return nil
}
