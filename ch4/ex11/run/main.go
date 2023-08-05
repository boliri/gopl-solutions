// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"issues"
)

var editor string = "nano"

func main() {
	subcommand, args := os.Args[1], os.Args[2:]

	for i := 0; i < len(args); i++ {
		if args[i] == "-e" || args[i] == "--editor" {
			abspath, err := exec.LookPath(args[i+1])
			if err != nil {
				fmt.Printf("Could not find binary path for text editor %s: %s\n", args[i+1], err)
				os.Exit(1)
			}
			editor = abspath

			copy(args[i:], args[i+2:])
			args = args[:len(args)-2]
			break
		}
	}

	pArgs := parseArgs(args)
	switch subcommand {
		case "help":
			help()
		case "search":
			reqAll := issues.SearchAllRequest{
				Owner: pArgs["owner"],
				Repo: pArgs["repo"],
				State: pArgs["state"],
				Auth: issues.Auth{
					Username: pArgs["username"],
					Token: pArgs["token"],
				},
			}
			if pArgs["issuenum"] != "" {
				issuenum, _ := strconv.ParseUint(pArgs["issuenum"], 10, 0)  // ignoring parse errors
				req := &issues.SearchRequest{
					SearchAllRequest: reqAll,
					IssueNumber: uint(issuenum),
				}

				resp, err := issues.Search(req)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				printSearch(resp.Issue)
			} else {
				resp, err := issues.SearchAll(&reqAll)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				printSearchAll(resp.Items)
			}
		case "create":
			reqParams := issues.CreateRequestParameters{}
			promptIssueEditor(&reqParams.ModifyIssueParameters)

			req := issues.CreateRequest{
				Owner: pArgs["owner"],
				Repo: pArgs["repo"],
				CreateRequestParameters: reqParams,
				Auth: issues.Auth{
					Username: pArgs["username"],
					Token: pArgs["token"],
				},
			}
			resp, err := issues.Create(&req)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			printCreateResults(resp.Issue)
		case "edit":
			issuenum, _ := strconv.ParseUint(pArgs["issuenum"], 10, 0)  // ignoring parse errors

			reqSearch := &issues.SearchRequest{
				SearchAllRequest: issues.SearchAllRequest{
					Owner: pArgs["owner"],
					Repo: pArgs["repo"],
					State: pArgs["state"],
					Auth: issues.Auth{
						Username: pArgs["username"],
						Token: pArgs["token"],
					},
				},
				IssueNumber: uint(issuenum),
			}

			respSearch, err := issues.Search(reqSearch)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			reqParams := issues.EditRequestParameters{}
			reqParams.Title = respSearch.Issue.Title
			reqParams.Body = respSearch.Issue.Body
			promptIssueEditor(&reqParams.ModifyIssueParameters)

			reqEdit := issues.EditRequest{
				Owner: pArgs["owner"],
				Repo: pArgs["repo"],
				IssueNumber: uint(issuenum),
				EditRequestParameters: reqParams,
				Auth: issues.Auth{
					Username: pArgs["username"],
					Token: pArgs["token"],
				},
			}
			resp, err := issues.Edit(&reqEdit)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			printEditResults(resp.Issue)
		case "close":
			issuenum, _ := strconv.ParseUint(pArgs["issuenum"], 10, 0)  // ignoring parse errors

			req := issues.CloseRequest{
				Owner: pArgs["owner"],
				Repo: pArgs["repo"],
				IssueNumber: uint(issuenum),
				CloseRequestParameters: issues.CloseRequestParameters{
					State: "closed",
				},
				Auth: issues.Auth{
					Username: pArgs["username"],
					Token: pArgs["token"],
				},
			}
			resp, err := issues.Close(&req)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			printCloseResults(resp.Issue)
	}
}

func help() {
	fmt.Println("This tool lets you manage GitHub issues. Available subcommands are:")
	fmt.Println("\thelp: prints this help message")
	fmt.Println("\tsearch: looks for either all issues in a repo, or the details of a specific issue")
	fmt.Println("\t\targs:")
	fmt.Println("\t\t\t-r, --repo\tcan be either a repo name or a string with format <owner>/<repo>")
	fmt.Println("\t\t\t-o, --owner\toptional, but only if -r already specifies an owner. overwrites owner defined by -r if defined")
	fmt.Println("\t\t\t-s, --state\toptional. can take one of these values: open, closed")
	fmt.Println("\t\t\t-i, --issue\toptional. if missing, the search will be global")
	fmt.Println("\t\t\t-u, --username\tonly needed if you need to display issues from private repos you own. if defined, -t must be specified as well")
	fmt.Println("\t\t\t-t, --token\tonly needed if you need to display issues from private repos you own. if defined, -u must be specified as well")
	fmt.Println("\tcreate: creates a new issue in the specified repo")
	fmt.Println("\t\targs:")
	fmt.Println("\t\t\t-r, --repo\tcan be either a repo name or a string with format <owner>/<repo>")
	fmt.Println("\t\t\t-o, --owner\toptional, but only if -r already specifies an owner. overwrites owner defined by -r if defined")
	fmt.Println("\t\t\t-u, --username\tmandatory")
	fmt.Println("\t\t\t-t, --token\tmandatory")
	fmt.Println("\t\tnotes:")
	fmt.Println("\t\t\tthis subcommand opens the editor specified in the -e (or --editor) option. if no editor is specified, nano will be used instead")
	fmt.Println("\t\t\tyou should write an issue title and the issue body when the editor opens")
	fmt.Println("\tedit: updates an existing issue in the specified repo")
	fmt.Println("\t\targs:")
	fmt.Println("\t\t\t-r, --repo\tcan be either a repo name or a string with format <owner>/<repo>")
	fmt.Println("\t\t\t-o, --owner\toptional, but only if -r already specifies an owner. overwrites owner defined by -r if defined")
	fmt.Println("\t\t\t-i, --issue\tmandatory, it's the issue number")
	fmt.Println("\t\t\t-u, --username\tmandatory")
	fmt.Println("\t\t\t-t, --token\tmandatory")
	fmt.Println("\t\tnotes:")
	fmt.Println("\t\t\tthis subcommand opens the editor specified in the -e (or --editor) option. if no editor is specified, nano will be used instead")
	fmt.Println("\t\t\tthe editor will show the current title and body of the issue, so you can edit them as you wish")
	fmt.Println("\tclose: closes an existing issue in the specified repo")
	fmt.Println("\t\targs:")
	fmt.Println("\t\t\t-r, --repo\tcan be either a repo name or a string with format <owner>/<repo>")
	fmt.Println("\t\t\t-o, --owner\toptional, but only if -r already specifies an owner. overwrites owner defined by -r if defined")
	fmt.Println("\t\t\t-i, --issue\tmandatory, it's the issue number")
	fmt.Println("\t\t\t-u, --username\tmandatory")
	fmt.Println("\t\t\t-t, --token\tmandatory")
	fmt.Println("Miscelaneous arguments:")
	fmt.Println("\t\t\t-e, --editor\teditor you want to use to edit issues' titles and bodies, for example vim or nano")
}

func promptIssueEditor(issueParams *issues.ModifyIssueParameters) {
	file, err := ioutil.TempFile("/tmp", "*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	if issueParams.Title != "" && issueParams.Title != "" {
		file.WriteString(issueParams.Title)
		file.WriteString("\n\n")
		file.WriteString(issueParams.Body)
		file.Seek(0, 0)
	}

	cmd := exec.Command(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var title, body string
	var firstNewlineFound bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !firstNewlineFound && scanner.Text() == "" {
			firstNewlineFound = true
			continue
		}
		
		if firstNewlineFound {
			body += scanner.Text()
		} else {
			title += scanner.Text()
		}
	}

	issueParams.Body = body
	issueParams.Title = title
}

func parseArgs(args []string) map[string]string {
	result := make(map[string]string)

	for i := 0; i < len(args); {
		if args[i] == "-s" || args[i] == "--state" {
			result["state"] = args[i+1]
		} else if args[i] == "-r" || args[i] == "--repo" {
			pcs := strings.Split(args[i+1], "/")
			if len(pcs) > 1 {
				result["owner"], result["repo"] = pcs[0], pcs[1]
			} else {
				result["repo"] = pcs[0]
			}
		} else if args[i] == "-o" || args[i] == "--owner" {
			result["owner"] = args[i+1]
		} else if args[i] == "-i" || args[i] == "--issue" {
			result["issuenum"] = args[i+1]
		} else if args[i] == "-u" || args[i] == "--username" {
			result["username"] = args[i+1]
		} else if args[i] == "-t" || args[i] == "--token" {
			result["token"] = args[i+1]
		}
		i += 2
	}

	return result
}

func printSearchAll(s []*issues.Issue) {
	fmt.Printf("%d issues:\n\n", len(s))
	for _, item := range s {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func printSearch(issue *issues.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n\n%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func printCreateResults(issue *issues.Issue) {
	fmt.Printf("Issue successfully created!\n\n")
	fmt.Printf("#%-5d %9.9s %.55s\n\n%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func printEditResults(issue *issues.Issue) {
	fmt.Printf("Issue successfully updated!\n\n")
	fmt.Printf("#%-5d %9.9s %.55s\n\n%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func printCloseResults(issue *issues.Issue) {
	fmt.Printf("Issue #%d successfully closed!\n\n", issue.Number)
	fmt.Printf("#%-5d %9.9s %.55s\n\n%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
}

//!-
