package general

import (
	"github.com/Unknwon/goconfig"
	"log"
	"path/filepath"
)

func TicketConfig() map[string]string {
	configPath, err := filepath.Abs("config.ini")
	if err != nil {
		log.Println(err)
	}

	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil {
		log.Println(err)
	}

	sess, _ := cfg.GetValue("ticket", "sess")
	price, _ := cfg.GetValue("ticket", "price")
	date, _ := cfg.GetValue("ticket", "date")
	realName, _ := cfg.GetValue("ticket", "real_name")
	nickName, _ := cfg.GetValue("ticket", "nick_name")
	ticketNum, _ := cfg.GetValue("ticket", "ticket_num")
	damaiUrl, _ := cfg.GetValue("ticket", "damai_url")
	targetUrl, _ := cfg.GetValue("ticket", "target_url")

	config := make(map[string]string, 0)
	config["secc"] = sess
	config["price"] = price
	config["date"] = date
	config["realName"] = realName
	config["nickName"] = nickName
	config["ticketNum"] = ticketNum
	config["damaiUrl"] = damaiUrl
	config["targetUrl"] = targetUrl

	return config
}
