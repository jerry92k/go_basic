// 대상 사이트 : 루리웹(ruliweb.com)

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// 스크래핑 대상 URL
const (
	urlRoot = "http://ruliweb.com"
)

// 첫번째 방문(메인페이지) 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		// n.DataAtom ==atom.A : 가져온 DataAtom이 atom 구조체의 a 태그와 같은지
		// n.Parent!=nil 부모노드가 있는 경우. 현재 사이트에서 찾고자하는 a태그의 부모노드는 row임
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

// 에러 체크
func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// 동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

func main() {
	// 메인 페이지 Get 방식 요청
	resp, err := http.Get(urlRoot)
	errCheck(err)

	// 요청 body 닫기
	defer resp.Body.Close()

	// 응답 데이터(Html)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	//ParseMainNodes 매소드를 크롤링(스크랩핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)

	for _, link := range urlList {
		// 대상 url 1차 출력
		//fmt.Println(" check name link : ", link, idx)
		//fmt.Println("TargetUrl : ", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)

		//fmt.Println("filename : ", fileName)
		// 작업 대기열에 추가
		wg.Add(1) // Done 개수와 일치

		// 고루틴 시작 -> 작업 대기열 개수와 같아야 함

		go scrapContents(scrape.Attr(link, "href"), fileName)

	}
	wg.Wait()

}

// URL 대상이 되는 페이지(서브 페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContents(url string, fn string) {
	// 작업 종료 알림
	defer wg.Done()

	resp, err := http.Get(url)
	errCheck(err)

	// 요청 body 닫기
	defer resp.Body.Close()

	// 응답 데이터(html)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	// Response 데이터(html)의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(n, "class") == "deco"
	}

	//fmt.Println(fn)
	file, err := os.OpenFile(fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	defer file.Close()

	w := bufio.NewWriter(file)

	// matchNode 매소드를 사용해서 원하는 노드 순회(Iterator) 하면서 출력

	for _, g := range scrape.FindAll(root, matchNode) {
		// Url 및 해당 데이터 출력
		//fmt.Println("result :", scrape.Text(g))
		// 파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}

	w.Flush()

}
