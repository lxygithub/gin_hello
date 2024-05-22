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

type KimiResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type KimiErrResp struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

type ChatgptResp struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// User 结构体定义用户数据的类型
type User struct {
	Username string
}
