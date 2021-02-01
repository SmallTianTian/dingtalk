package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceHasfaceRequest() *OapiSmartdeviceHasfaceRequest {
	return &OapiSmartdeviceHasfaceRequest{}
}

type OapiSmartdeviceHasfaceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceHasfaceResponse
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartdeviceHasfaceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceHasfaceRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceHasfaceRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartdeviceHasfaceRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartdeviceHasfaceRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.hasface"
}
func (this *OapiSmartdeviceHasfaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceHasfaceRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceHasfaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceHasfaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceHasfaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartdeviceHasfaceRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UseridList, "useridList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceHasfaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceHasfaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceHasfaceResponse struct {
	taobao.TaobaoResponse

	UseridList []string `json:"userid_list,omitempty"`
}
