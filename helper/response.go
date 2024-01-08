package helper

type Meta struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

func ResponseAPI(success bool, code int, message string, data any) Response {
	return Response{
		Meta: Meta{
			Code:    code,
			Success: success,
			Message: message,
		},
		Data: data,
	}
}
