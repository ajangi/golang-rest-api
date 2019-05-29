package utils

// ResponseApi : this struct is just for error responses in APIs
type ResponseApi struct {
	Result     string           `json:"result"`
	StatusCode int              `json:"status_code"`
	Messages   ResponseMessages `json:"messages"`
	Data       ResponseData     `json:"data"`
}

// ResponseData : this struct is for the data we want to have in API responses
type ResponseData []interface{}

// ResponseMessages : this struct is for the data we want to have in API responses
type ResponseMessages []MessageItem

// MessageItem : Single message item should be in this format
type MessageItem struct {
	Fa string `json:"fa"`
	En string `json:"en"`
}
