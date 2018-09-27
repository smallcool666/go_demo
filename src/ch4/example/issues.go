package main
/*
	通过Github的issue查询服务来演示通过HTTP接口发送JSON格式请求并返回JSON格式的信息
	本文件调用封装好的包格式化输出结果
	执行go run issues.go repo:golang/go is:open json decoder命令运行
*/
import (
	"fmt"
	"log"
	"os"
	"ch4/example/github"
)
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
		item.Number, item.User.Login, item.Title)
	}
}