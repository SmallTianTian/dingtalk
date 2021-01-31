package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceMeetingroomParticipantListRequest() *OapiSmartdeviceMeetingroomParticipantListRequest {
	return &OapiSmartdeviceMeetingroomParticipantListRequest{}
}

type OapiSmartdeviceMeetingroomParticipantListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceMeetingroomParticipantListResponse
	Bookid          string
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) SetBookid(bookid2 string) {
	this.Bookid = bookid2
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetBookid() string {
	return this.Bookid
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.meetingroom.participant.list"
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bookid"] = this.Bookid
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Bookid, "bookid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceMeetingroomParticipantListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceMeetingroomParticipantListResponse struct {
	taobao.TaobaoResponse
	Result OpenPageResult `json:"result,omitempty"`
}
type MemeberOpenInfo struct {
	Userid string `json:"userid,omitempty"`
}
