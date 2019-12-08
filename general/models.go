package general

type Config struct {
	Sess              int     `json:"sess"`
	Price             int     `json:"price"`
	Date              int     `json:"date"`
	RealName          string  `json:"real_name"`
	NickName          string  `json:"nick_name"`
	TicketNum         string  `json:"ticket_num"`
	DamaiUrl          string  `json:"damai_url"`
	TargetUrl         string  `json:"target_url"`
	TotalWaitTime     int     `json:"total_wait_time"`
	RefreshWaitTime   float64 `json:"refresh_wait_time"`
	IntersectWaitTime float64 `json:"intersect_wait_time"`
}
