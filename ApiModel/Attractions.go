package ApiModel

type Attractions struct {
	Total int `json:"total"`
	Data  []struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		NameZh       interface{} `json:"name_zh"`
		OpenStatus   int         `json:"open_status"`
		Introduction string      `json:"introduction"`
		OpenTime     string      `json:"open_time"`
		Zipcode      string      `json:"zipcode"`
		Distric      string      `json:"distric"`
		Address      string      `json:"address"`
		Tel          string      `json:"tel"`
		Fax          string      `json:"fax"`
		Email        string      `json:"email"`
		Months       string      `json:"months"`
		Nlat         float64     `json:"nlat"`
		Elong        float64     `json:"elong"`
		OfficialSite string      `json:"official_site"`
		Facebook     string      `json:"facebook"`
		Ticket       string      `json:"ticket"`
		Remind       string      `json:"remind"`
		Staytime     string      `json:"staytime"`
		Modified     string      `json:"modified"`
		URL          string      `json:"url"`
		Category     []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"category"`
		Target   []interface{} `json:"target"`
		Service  []interface{} `json:"service"`
		Friendly []interface{} `json:"friendly"`
		Images   []struct {
			Src     string `json:"src"`
			Subject string `json:"subject"`
			Ext     string `json:"ext"`
		} `json:"images"`
		Files []interface{} `json:"files"`
		Links []interface{} `json:"links"`
	} `json:"data"`
}
