package github

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

// Edit is a function
func Edit(value *NewIssue) error {
	editor := "code"

	tempFile, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return err
	}

	tempFileName := tempFile.Name()
	defer os.Remove(tempFileName)

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		return err
	}

	err = tempFile.Close()
	if err != nil {
		return err
	}

	cmd := exec.Command(editor, tempFileName, "-w")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, value)
	if err != nil {
		return err
	}
	return nil
}
