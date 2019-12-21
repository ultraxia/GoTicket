package main

import (
	"GoTicket/general"
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var num = 0
var cookie = "cookies.pkl"
var config map[string]string
var webDriver selenium.WebDriver

func init() {
	config = general.TicketConfig()
	webDriver = general.Driver()
}

func main() {
	login()
	order_ticket()

}

func login() {
	if exist(cookie) == false {
		fmt.Println("### 未找到cookies，正在调用登录组件 ###")
		get_cookie()
	} else {
		set_cookie()
		fmt.Println("### 发现cookies，自动登录成功 ###")
	}
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func get_cookie() {
	var web_title string

	if err := webDriver.Get(config["login_url"]); err != nil {
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
		fmt.Printf("Cookie获取失败\n%v", err)
	}

	general.Writefile(cookie)
	fmt.Printf("### 登录成功 ###\n")
}

func set_cookie() {
	webDriver.Get(config["target_url"])
	var cookies []selenium.Cookie

	filePtr, err := os.Open(cookie)
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&cookies)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range cookies {
		session := &selenium.Cookie{
			Name:   v.Name,
			Value:  v.Value,
			Expiry: math.MaxUint32,
		}
		session.Domain = v.Domain
		session.Path = "/"
		session.Secure = false

		err = webDriver.AddCookie(session)

	}

}

func order_ticket() {
	webDriver.Get(config["target_url"])
	showName, _ := webDriver.Title()
	fmt.Printf("### 正在加载：%v ###\n",showName)
	//startTime := time.Now()
	//fmt.Println(startTime)

	//for {
	num += 1
	web_title, _ := webDriver.Title()
	if strings.Contains(web_title, "确认订单") == true {
		fmt.Printf("### 经过%v轮的努力，恭喜您成功抢到票 ###\n", num)
	} else {
		select_city()
		select_date()
		select_price()
		confirm_order()

	}

}

//选择城市
func select_city() {
	if config["date"] != "0" || len(config["date"]) == 0 {
		orderElements := general.GetOrderPicker(webDriver)
		cityElements := orderElements[0]
		cityList, err := cityElements.FindElements(selenium.ByTagName, "span")
		if err != nil {
			fmt.Println(nil)
		}
		cityConfig, _ := strconv.Atoi(config["city"])
		cityButton := cityList[cityConfig-1]
		cityName, err := cityButton.Text()
		if err != nil {
			log.Fatalln(err)
		}
		err = cityButton.Click()
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Printf("### 已选择： %v ###", cityName)
		}
	}

}

//选择场次
func select_date() {
	if config["date"] != "0" || len(config["date"]) == 0 {
		orderElements := general.GetOrderPicker(webDriver)
		dateElements := orderElements[1]
		dateList, err := dateElements.FindElements(selenium.ByTagName, "span")
		if err != nil {
			log.Fatalln(err)
		}
		dateConfig, _ := strconv.Atoi(config["date"])
		dateButton := dateList[dateConfig-1]
		date, err := dateButton.Text()
		if err != nil {
			log.Fatalln(err)
		}
		err = dateButton.Click()
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Printf("### 已选择： %v ###", date)
		}
	}
}

//选择票档
func select_price() {
	orderElements := general.GetOrderPicker(webDriver)
	priceElements := orderElements[2]
	orderList, err := priceElements.FindElements(selenium.ByClassName, "skuname")
	if err != nil {
		log.Fatalln(err)
	}
	priceConfig, _ := strconv.Atoi(config["price"])
	priceButton := orderList[priceConfig-1]
	price, err := priceButton.Text()
	if err != nil {
		log.Fatalln(err)
	}

	err = orderList[priceConfig-1].Click()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("### 票档选择成功,已选择 %v ###\n", price)
	}

}

func confirm_order() {
	pickers, err := webDriver.FindElement(selenium.ByClassName, "buybtn")
	if err != nil {
		log.Fatalln(err)
	}
	button_text, err := pickers.Text()

	if strings.Contains(button_text, "选座购买") == true {
		pickers.Click()
		fmt.Printf("### 请手工选座并提交 ###\n")
	}

	if strings.Contains(button_text, "提交缺货登记") == true {
		pickers.Click()
		fmt.Printf("### 票已售罄，已为您提交缺货登记 ###\n")
	}

	if strings.Contains(button_text, "立即预定") == true {
		pickers.Click()
		fmt.Println("### 预售已开启，请手动确认订单 ###")
	}
}
