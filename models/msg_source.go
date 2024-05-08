package models
type MsgSource struct {
	Room struct {
		Events struct {
		} `json:"_events"`
		EventsCount int    `json:"_eventsCount"`
		ID          string `json:"id"`
		Payload     struct {
			AdminIDList []interface{} `json:"adminIdList"`
			Avatar      string        `json:"avatar"`
			ID          string        `json:"id"`
			Topic       string        `json:"topic"`
			MemberList  []struct {
				Avatar string `json:"avatar"`
				ID     string `json:"id"`
				Name   string `json:"name"`
				Alias  string `json:"alias"`
			} `json:"memberList"`
		} `json:"payload"`
	} `json:"room"`
	To struct {
		Events struct {
		} `json:"_events"`
		EventsCount int    `json:"_eventsCount"`
		ID          string `json:"id"`
		Payload     struct {
			Alias     string        `json:"alias"`
			Avatar    string        `json:"avatar"`
			Friend    bool          `json:"friend"`
			Gender    int           `json:"gender"`
			ID        string        `json:"id"`
			Name      string        `json:"name"`
			Phone     []interface{} `json:"phone"`
			Signature string        `json:"signature"`
			Star      bool          `json:"star"`
			Type      int           `json:"type"`
		} `json:"payload"`
	} `json:"to"`
	From struct {
		Events struct {
		} `json:"_events"`
		EventsCount int    `json:"_eventsCount"`
		ID          string `json:"id"`
		Payload     struct {
			Address   string        `json:"address"`
			Alias     string        `json:"alias"`
			Avatar    string        `json:"avatar"`
			City      string        `json:"city"`
			Friend    bool          `json:"friend"`
			Gender    int           `json:"gender"`
			ID        string        `json:"id"`
			Name      string        `json:"name"`
			Phone     []interface{} `json:"phone"`
			Province  string        `json:"province"`
			Signature string        `json:"signature"`
			Star      bool          `json:"star"`
			Weixin    string        `json:"weixin"`
			Type      int           `json:"type"`
		} `json:"payload"`
	} `json:"from"`
}