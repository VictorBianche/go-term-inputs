package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

func InputPromp(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func SensitivePrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		pw, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(pw)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}

var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "enter your name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "pet",
		Prompt: &survey.MultiSelect{
			Message: "Choose your pets:",
			Options: []string{"dogs", "reptiles", "cats", "birds", "fish", "rabbits", "pigs", "rats", "mices"},
			Default: "dogs",
		},
	},
	{
		Name:   "rating",
		Prompt: &survey.Input{Message: "Rate our website (integer number):"},
	},
}

func main() {
	answer := struct {
		Name   string   // survey match the question and field names
		Pet    []string `survey:"pet"` //tag fields to match a specific name
		Rating int      // if the types don't match, survey will convert it
	}{}

	err := survey.Ask(qs, &answer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	petString := strings.Join(answer.Pet, ", ")

	prompt := promptui.Select{
		Label: "Select your characters:",
		Items: []string{"Harry Potter", "Ron Weasley", "Hermione Granger", "Ginny Weasley", "Neville Longbottom",
			"Luna Lovegood", "Draco Malfoy", "Albus Dumbledore"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	sensitiveString := SensitivePrompt("Enter your password:")
	message := InputPromp("Enter your message:")

	fmt.Printf("You choose %q\n", result)
	fmt.Printf("Your message: %s!\n", message)
	fmt.Printf("%s likes %s.\n", answer.Name, petString)
	fmt.Printf("Your password is %q\n", sensitiveString)
}
