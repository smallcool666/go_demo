package github
/*
	通过Github的issue查询服务来演示通过HTTP接口发送JSON格式请求并返回JSON格式的信息
*/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
// SearchIssues queries the GitHub issue tracker.
//SearchIssues函数发出一个HTTP请求，然后解码返回的JSON格式的结果
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	//因为用户提供的查询条件可能包含类似?和&之类的特殊字符，为了避免对URL造成冲突，我们用url.QueryEscape来对查询中的特殊字符进行转义操作
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	//使用了基于流式的解码器json.Decoder，它可以从一个输入流解码JSON数据，尽管这不是必须的。
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}