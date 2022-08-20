package utils

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const HttpHeaderKeyAuthorization = "Authorization"
const HttpHeaderKeyContentType = "Content-Type"
const HttpHeaderValueContentTypeApplicationJson = "application/json"
const HttpHeaderValueContentTypeApplicationXml = "application/xml"

type ResponseHelper struct {
	Writer http.ResponseWriter
}

//#region Json
func (h *ResponseHelper) setContentTypeAsApplicationJson() {
	h.Writer.
		Header().
		Set(HttpHeaderKeyContentType, HttpHeaderValueContentTypeApplicationJson)
}

func (h *ResponseHelper) setJsonBody(body interface{}) {
	jsn, _ := json.Marshal(body)
	h.Writer.
		Write(jsn)
}
//#endregion

//#region Xml
func (h *ResponseHelper) setContentTypeAsApplicationXml() {
	h.Writer.
		Header().
		Set(HttpHeaderKeyContentType, HttpHeaderValueContentTypeApplicationXml)
}

func (h *ResponseHelper) setXmlBody(body interface{}) {
	x, _ := xml.Marshal(body)
	h.Writer.
		Write(x)
}
//#endregion

func (h *ResponseHelper) setStatusCode(statusCode int) {
	h.Writer.
		WriteHeader(statusCode)
}

func (h *ResponseHelper) SetJsonResponse(statusCode int, body interface{}) {
	h.setContentTypeAsApplicationJson()
	h.setStatusCode(statusCode)
	h.setJsonBody(body)
}

func (h *ResponseHelper) SetXmlResponse(statusCode int, body interface{}) {
	h.setContentTypeAsApplicationXml()
	h.setStatusCode(statusCode)
	h.setXmlBody(body)
}
