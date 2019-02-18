// Issues prints a table of GitHub issues matching the search terms
package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"gopl.io/ch4/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// Simple print
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
