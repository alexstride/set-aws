package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type AssumeRoleOutput struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

func (a *AssumeRoleOutput) UnmarshalJSON(b []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	a.AccessKeyId = v["Credentials"].(map[string]interface{})["AccessKeyId"].(string)
	a.SecretAccessKey = v["Credentials"].(map[string]interface{})["SecretAccessKey"].(string)
	a.SessionToken = v["Credentials"].(map[string]interface{})["SessionToken"].(string)
	a.Expiration = v["Credentials"].(map[string]interface{})["Expiration"].(string)
	return nil
}

func main() {

	inputString, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	// Unmarshal inputString into AssumeRoleOutput struct
	var assumeRoleOutput AssumeRoleOutput
	err = json.Unmarshal(inputString, &assumeRoleOutput)
	if err != nil {
		log.Fatal(err)
	}
	if assumeRoleOutput.AccessKeyId == "" {
		log.Fatal("AccessKeyId is empty")
	}
	if assumeRoleOutput.AccessKeyId == "" {
		log.Fatal("SecretAccessKey is empty")
	}
	if assumeRoleOutput.AccessKeyId == "" {
		log.Fatal("SessionToken is empty")
	}

	fmt.Printf(`
export AWS_ACCESS_KEY_ID=%s
export AWS_SECRET_ACCESS_KEY=%s
export AWS_SESSION_TOKEN=%s
unset AWS_PROFILE
`, assumeRoleOutput.AccessKeyId, assumeRoleOutput.SecretAccessKey, assumeRoleOutput.SessionToken)

	fmt.Fprintf(os.Stderr, "Set env. Expiration: %s\n", assumeRoleOutput.Expiration)
}
