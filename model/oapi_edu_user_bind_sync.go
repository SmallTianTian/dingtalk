package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUserBindSyncRequest() *OapiEduUserBindSyncRequest {
	return &OapiEduUserBindSyncRequest{}
}

type OapiEduUserBindSyncRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUserBindSyncResponse
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiEduUserBindSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUserBindSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUserBindSyncRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiEduUserBindSyncRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiEduUserBindSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.user.bind.sync"
}
func (this *OapiEduUserBindSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUserBindSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUserBindSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUserBindSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUserBindSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *OapiEduUserBindSyncRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UserId, "userId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUserBindSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUserBindSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUserBindSyncResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
