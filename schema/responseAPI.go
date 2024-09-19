package schema

type ResponseEmpty struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ResponseData struct {
	Message    	string 		`json:"message"`
	StatusCode 	int    		`json:"statusCode"`
	Data    	interface{} `json:"data"`
}