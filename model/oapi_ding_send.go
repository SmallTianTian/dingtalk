package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingSendRequest() *OapiDingSendRequest {
	return &OapiDingSendRequest{}
}

type OapiDingSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingSendResponse
	OpenDingSendVo  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingSendRequest) SetOpenDingSendVo(openDingSendVo2 string) {
	this.OpenDingSendVo = openDingSendVo2
}
func (this *OapiDingSendRequest) GetOpenDingSendVo() string {
	return this.OpenDingSendVo
}
func (this *OapiDingSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ding.send"
}
func (this *OapiDingSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_ding_send_vo"] = this.OpenDingSendVo
	return txtParams
}
func (this *OapiDingSendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AttachmentVo struct {
	DetailType  string `json:"detail_type,omitempty"`
	FileId      string `json:"file_id,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	FileSize    int64  `json:"file_size,omitempty"`
	FileSpaceId string `json:"file_space_id,omitempty"`
	LinkPicUrl  string `json:"link_pic_url,omitempty"`
	LinkText    string `json:"link_text,omitempty"`
	LinkTitle   string `json:"link_title,omitempty"`
	LinkUrl     string `json:"link_url,omitempty"`
	Type        string `json:"type,omitempty"`
}
type OpenDingSendVo struct {
	Attachment   AttachmentVo `json:"attachment,omitempty"`
	ReceiverUids []string     `json:"receiver_uids,omitempty"`
	RemindType   int64        `json:"remind_type,omitempty"`
	TextContent  string       `json:"text_content,omitempty"`
}
type OapiDingSendResponse struct {
	taobao.TaobaoResponse
	Result CorpDingCreateResult `json:"result,omitempty"`
}
