package messages_go

// ResponseData Mensaje de respuesta
type ResponseData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Href    string `json:"href"`
}

// ResponseMessage contiene los mensajes y errores
type ResponseMessage struct {
	Data     interface{}    `json:"data"`
	Messages []ResponseData `json:"messages"`
	Errors   []ResponseData `json:"errors"`
}

// AddMessage permite agregar un mensaje a ResponseMessage
func (rm *ResponseMessage) AddMessage(c, m, h string) {
	rd := ResponseData{
		Code:    c,
		Message: m,
		Href:    h,
	}
	rm.Messages = append(rm.Messages, rd)
}

// AddError permite agregar un error a ResponseMessage
func (rm *ResponseMessage) AddError(c, m, h string) {
	rd := ResponseData{
		Code:    c,
		Message: m,
		Href:    h,
	}
	rm.Errors = append(rm.Errors, rd)
}
