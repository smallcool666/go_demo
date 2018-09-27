// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github
/*
	通过Github的issue查询服务来演示通过HTTP接口发送JSON格式请求并返回JSON格式的信息
	本文件定义了基础的数据结构
*/
import "time"
const IssuesURL = "https://api.github.com/search/issues"
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}
type Issue struct {
	Number int
	//因为有些JSON成员名字和Go结构体成员名字并不相同，因此需要Go语言结构体成员Tag来指定对应的JSON名字。同样，在解码的时候也需要做同样的处理
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string // in Markdown format
}
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}