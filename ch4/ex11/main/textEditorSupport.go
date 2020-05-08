/*
create.go contains functions supporting the creation of new GitHub issues.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"the_go_programming_language/ch4/ex11/github"
)

const (
	defaultEditor = "vim"
	instructions  = "\n\n# Enter the issue description. Lines starting with # will be ignored."
)

// getIssueDetails gets the title and body of the new issue from the user
func getNewIssueDetails() (*github.Issue, error) {
	// Collect the issue title via the terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an issue title: ")
	scanner.Scan()
	title := scanner.Text()

	// Collect the issue body via the user's preferred text editor
	fmt.Print("Enter the issue body: ")
	body, err := getInputWithTextEditor("")
	if err != nil {
		return nil, err
	}
	fmt.Println(body)

	issue := github.Issue{
		Title: title,
		Body:  body,
	}
	return &issue, nil
}

// getUpdatedIssueDetails prompts the user to edit the title and body of an exiting issue
func getUpdatedIssueDetails(current *github.Issue) (*github.Issue, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Title: %s\n", current.Title)
	fmt.Print("Enter new title (leave blank for unchanged): ")
	var title string
	if scanner.Scan() {
		title = scanner.Text()
	} else {
		title = current.Title
	}

	fmt.Printf("Edit the issue body: ")
	body, err := getInputWithTextEditor(current.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(body)

	issue := github.Issue{
		Title: title,
		Body:  body,
	}
	return &issue, nil
}

func getInputWithTextEditor(currentBody string) (string, error) {
	// Create a temporary file to store text editor input
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return "", fmt.Errorf("Couldn't create tempfile: %s", err)
	}
	filename := file.Name()
	defer os.Remove(filename)

	// Write any existing issue body text for editing
	if currentBody != "" {
		_, err = file.WriteString(currentBody)
	}
	// Write instructions for the user to the tempfile, then close it
	if _, err = file.WriteString(instructions); err != nil {
		return "", fmt.Errorf("Couldn't write instructions to tempfile: %s", err)
	}
	if err = file.Close(); err != nil {
		return "", fmt.Errorf("Couldn't close tempfile: %s", err)
	}

	// Get user input into the newly created file
	if err = openTempFileInEditor(filename); err != nil {
		return "", err
	}

	// Read user input from the file and return it
	body, err := parseTempFile(filename)
	if err != nil {
		return "", fmt.Errorf("Couldn't read from tempfile: %s", err)
	}

	return body, nil
}

func openTempFileInEditor(filename string) error {
	// Get the user's preferred editor, or set the default if none is found
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}
	executable, err := exec.LookPath(editor)
	if err != nil {
		return fmt.Errorf("Couldn't find text editor %s: %s", editor, err)
	}
	// Open the temp file in the editor, setting its input and output to Stdin and Stdout
	cmd := exec.Command(executable, filename)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

// parseTempFile removes all comments from tempFile before returning the remaining text
func parseTempFile(filename string) (string, error) {
	tempFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	var lines []string
	scanner := bufio.NewScanner(tempFile)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			// Include blank lines
			lines = append(lines, "\n")
		} else if text[0] != '#' {
			lines = append(lines, text)
		}
	}
	// Preserve all newlines except the last one before the instructional text
	return strings.Join(lines[:len(lines)-1], "\n"), scanner.Err()
}
