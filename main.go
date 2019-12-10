package main

import (
	"GoTicket/general"
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"strings"
	"time"
)

var config map[string]string
var webDriver selenium.WebDriver

func init() {
	config = general.TicketConfig()
	webDriver = general.Driver()
}

func main() {
	get_cookie()
	order_ticket()
}

func login() {

}

func get_cookie() {
	var web_title string
	if err := webDriver.Get(config["damai_url"]); err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	fmt.Println("### 请点击登录 ###")
	for {
		web_title, _ = webDriver.Title()
		if strings.Contains(web_title, "大麦网-全球演出赛事官方购票平台") == true {
			time.Sleep(1)
		} else {
			break
		}
	}

	fmt.Println("### 请选择扫码登陆 ###")
	for {
		web_title, _ = webDriver.Title()
		if web_title == "大麦登录" || strings.Contains(web_title, "大麦网-全球演出赛事官方购票平台") != true {
			time.Sleep(1)
		} else {
			break
		}

	}

	cookie, err := webDriver.GetCookies()
	if err != nil {
		log.Printf("Cookie获取失败\n%v", err)
	}

	general.Writefile(cookie)
	fmt.Println("### 登录成功 ###")
}

//TODO 自动读取cookie
func set_cookie() {
	var cookies []selenium.Cookie
	var cookie_dict map[string]interface{}

	filePtr, err := os.Open("cookies.pkl")
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&cookies)
	if err != nil {
		fmt.Println(err)
	}
	for _, cookie := range cookies {
		cookie_dict["domain"] = ".damai.cn"
		cookie_dict["name"] = cookie.Name
		cookie_dict["value"] = cookie.Value
		cookie_dict["expires"] = ""
		cookie_dict["path"] = "/"
		cookie_dict["httpOnly"] = false
		cookie_dict["HostOnly"] = false
		cookie_dict["Secure"] = false
	}
	//webDriver.AddCookie(cookie_dict)
}

func order_ticket()  {
	webDriver.Get(config["target_url"])
	webDriver.Refresh()

}