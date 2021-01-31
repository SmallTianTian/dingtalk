package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUseridGetRequest() *OapiEduUseridGetRequest {
	return &OapiEduUseridGetRequest{}
}

type OapiEduUseridGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUseridGetResponse
	Mobiles         string
	Operator        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduUseridGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUseridGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUseridGetRequest) SetMobiles(mobiles2 string) {
	this.Mobiles = mobiles2
}
func (this *OapiEduUseridGetRequest) GetMobiles() string {
	return this.Mobiles
}
func (this *OapiEduUseridGetRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduUseridGetRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduUseridGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.userid.get"
}
func (this *OapiEduUseridGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUseridGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUseridGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUseridGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUseridGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mobiles"] = this.Mobiles
	txtParams["operator"] = this.Operator
	return txtParams
}
func (this *OapiEduUseridGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Mobiles, "mobiles"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUseridGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUseridGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUseridGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenEduUserIdListResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type UserIdHold struct {
	Mobile string `json:"mobile,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OpenEduUserIdListResponse struct {
	Users []UserIdHold `json:"users,omitempty"`
}
