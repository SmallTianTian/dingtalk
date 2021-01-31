package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAlitripBtripProjectDeleteRequest() *OapiAlitripBtripProjectDeleteRequest {
	return &OapiAlitripBtripProjectDeleteRequest{}
}

type OapiAlitripBtripProjectDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripProjectDeleteResponse
	Corpid          string
	ThirdPartId     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripProjectDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripProjectDeleteRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiAlitripBtripProjectDeleteRequest) SetThirdPartId(thirdPartId2 string) {
	this.ThirdPartId = thirdPartId2
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetThirdPartId() string {
	return this.ThirdPartId
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.project.delete"
}
func (this *OapiAlitripBtripProjectDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripProjectDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripProjectDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corpid"] = this.Corpid
	txtParams["third_part_id"] = this.ThirdPartId
	return txtParams
}
func (this *OapiAlitripBtripProjectDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Corpid, "corpid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripProjectDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripProjectDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Module  bool   `json:"module,omitempty"`
	Success bool   `json:"success,omitempty"`
}
