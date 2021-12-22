package helper

type ResponseStruct struct {
	Data        interface{} `json:"data"`
	Description string      `json:"description"`
}

func Response(data interface{}, description string) ResponseStruct {
	return ResponseStruct{Data: data, Description: description}
}
