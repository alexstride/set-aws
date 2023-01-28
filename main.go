package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type AssumeRoleOutput struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

func main() {

	inputString, err := io.ReadAll(os.Stdin)

	fmt.Println("Input string: ", string(inputString))

	if err != nil {
		panic(err)
	}

	// Unmarshal inputString into AssumeRoleOutput struct
	var assumeRoleOutput AssumeRoleOutput
	err = json.Unmarshal(inputString, &assumeRoleOutput)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`
export AWS_ACCESS_KEY_ID=%s
export AWS_SECRET_ACCESS_KEY=%s
export AWS_SESSION_TOKEN=%s
unset AWS_PROFILE
`, assumeRoleOutput.AccessKeyId, assumeRoleOutput.SecretAccessKey, assumeRoleOutput.SessionToken)

	fmt.Fprintf(os.Stderr, "Set env. Expiration: %s\n", assumeRoleOutput.Expiration)
}
