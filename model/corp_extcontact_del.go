package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpExtcontactDeleteRequest() *CorpExtcontactDeleteRequest {
	return &CorpExtcontactDeleteRequest{}
}

type CorpExtcontactDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtcontactDeleteResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpExtcontactDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtcontactDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpExtcontactDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpExtcontactDeleteRequest) GetApiMethodName() string {
	return "dingtalk.corp.extcontact.delete"
}
func (this *CorpExtcontactDeleteRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtcontactDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtcontactDeleteRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtcontactDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtcontactDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtcontactDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpExtcontactDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *CorpExtcontactDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtcontactDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtcontactDeleteResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
