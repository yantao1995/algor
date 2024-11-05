package got

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestSocks5h(t *testing.T) {
	//动态端口转发
	//ssh -D 8989 servername -N
	proxyUrl, _ := url.Parse("socks5h://localhost:8989")
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}

	resp, err := client.Get("https://api.binance.com/api/v3/ticker/price")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
