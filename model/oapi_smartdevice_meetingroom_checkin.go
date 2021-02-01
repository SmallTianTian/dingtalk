package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceMeetingroomCheckinRequest() *OapiSmartdeviceMeetingroomCheckinRequest {
	return &OapiSmartdeviceMeetingroomCheckinRequest{}
}

type OapiSmartdeviceMeetingroomCheckinRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceMeetingroomCheckinResponse
	Bookid          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) SetBookid(bookid2 string) {
	this.Bookid = bookid2
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetBookid() string {
	return this.Bookid
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.meetingroom.checkin"
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bookid"] = this.Bookid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Bookid, "bookid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceMeetingroomCheckinRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceMeetingroomCheckinResponse struct {
	taobao.TaobaoResponse
}
