package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpSmartdeviceHasfaceRequest() *CorpSmartdeviceHasfaceRequest {
	return &CorpSmartdeviceHasfaceRequest{}
}

type CorpSmartdeviceHasfaceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpSmartdeviceHasfaceResponse
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *CorpSmartdeviceHasfaceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpSmartdeviceHasfaceRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *CorpSmartdeviceHasfaceRequest) GetUseridList() string {
	return this.UseridList
}
func (this *CorpSmartdeviceHasfaceRequest) GetApiMethodName() string {
	return "dingtalk.corp.smartdevice.hasface"
}
func (this *CorpSmartdeviceHasfaceRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpSmartdeviceHasfaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpSmartdeviceHasfaceRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpSmartdeviceHasfaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpSmartdeviceHasfaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpSmartdeviceHasfaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *CorpSmartdeviceHasfaceRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UseridList, "useridList"); err != nil {
		return err
	}
	return nil
}
func (this *CorpSmartdeviceHasfaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpSmartdeviceHasfaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpSmartdeviceHasfaceResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
