package main
/*
	输出的html模板
	执行go run issueshtml.go  repo:golang/go is:open json decoder获取输出
*/
import (
	"ch4/example/github"
	"html/template"
	"log"
	"os"
)
var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th></tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
