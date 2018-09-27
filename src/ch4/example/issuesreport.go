package main
/*
	输出的文本模板
	执行go run issurereport.go  repo:golang/go is:open json decoder得到预期结果
*/
import (
	"ch4/example/github"
	"log"
	"os"
	"time"
	"text/template"
)

//模板中{{range .Items}}和{{end}}对应一个循环action
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//模板通常在编译时就测试好了，如果模板解析失败将是一个致命的错误。template.Must辅助函数可以简化这个致命错误的处理
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

