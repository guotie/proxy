package proxy

import (
	"fmt"
	"net/http"
	//"strings"
	"time"
	//"io/ioutil"
	//"net/http"

	"github.com/PuerkitoBio/goquery"
)

var _ = fmt.Printf

const (
	ProxyHttp = iota
	ProxyHttps
	ProxySocks4
	ProxySocks5
)

// proxy 的信息
type ProxyInfo struct {
	Addr        string    // 地址
	SecretLevel string    // 匿名等级
	Location    string    // proxy地理位置
	Typ         int       // 类型 http, https, sock4, sock5
	Speed       int       // 响应时间, 毫秒
	VerifyTime  time.Time // 最后验证时间
}

// NewProxyInfo 建立一个新的proxy info
func NewProxyInfo(addr, secretLevel, location string, typ, speed int, verified time.Time) *ProxyInfo {
	return &ProxyInfo{addr, secretLevel, location, typ, speed, verified}
}

// CrawlProxyPage 爬取proxy页面
func CrawlProxyPage(resp *http.Response,
	fn func(*goquery.Document) []*ProxyInfo) (proxies []*ProxyInfo, err error) {
	// new document
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return
	}

	proxies = fn(doc)

	return
}

// GetProxyFn 生成并返回一个函数，该函数用来根据document提取ProxyInfo
func GetProxyFn() func(*goquery.Document) []*ProxyInfo {
	return nil
}

/*
// 从http://www.kuaidaili.com/proxylist下载最新的代理服务器列表

// 获得一个页面下的代理地址
func getPageProxyList(addr, id string, ipSel, portSel string) ([]string, error) {
	_, rd, err := getWithProxy(addr, "")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, err
	}

	var proxyList []string
	// 免费高速HTTP代理IP列表
	doc.Find("#" + id).Find("tr").Each(func(i int, node *goquery.Selection) {
		//fmt.Println("found tr node:", node.Text())
		ip := node.Find("td[data-title=IP]").Text()
		//fmt.Println("ip:", ip)
		port := node.Find("td[data-title=PORT]").First().Text()
		ip = strings.ToLower(strings.TrimSpace(ip))
		port = strings.TrimSpace(port)
		if ip != "" && port != "" {
			if strings.HasPrefix(ip, "http://") == false {
				ip = "http://" + ip
			}
			proxyList = append(proxyList, ip+":"+port)
		}
	})

	//fmt.Println(proxyList)

	return proxyList, nil
}

func getProxyLists(prefix, id string, pages int) []string {
	var kuaidaili = "%s%d/"
	var res []string

	if strings.HasSuffix(prefix, "/") == false {
		prefix += "/"
	}

	for i := 1; i <= pages; i++ {
		addr := fmt.Sprintf(kuaidaili, prefix, i)
		l, err := getPageProxyList(addr, id)
		if err == nil {
			res = append(res, l...)
		}
		time.Sleep(time.Second)
	}

	fmt.Printf("Got %d proxy from %s\n", len(res), prefix)

	return res
}

func getAllProxyLists() []string {
	var lists []string

	res := getProxyLists("http://www.kuaidaili.com/proxylist/", "index_free_list", 10)
	lists = append(lists, res...)

	res = getProxyLists("http://www.kuaidaili.com/free/inha/", "list", 10)
	lists = append(lists, res...)

	res = getProxyLists("http://www.kuaidaili.com/free/intr/", "list", 10)
	lists = append(lists, res...)

	time.Sleep(time.Second)
	res = getProxyLists("http://www.kuaidaili.com/free/outha/", "list", 10)
	lists = append(lists, res...)

	time.Sleep(time.Second)
	res = getProxyLists("http://www.kuaidaili.com/free/outtr/", "list", 10)
	lists = append(lists, res...)

	lists = append([]string{}, lists...)
	fmt.Printf("Total proxy: %d\n", len(lists))
	return lists
}

//
// http://www.xicidaili.com/nn/
*/
