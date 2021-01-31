package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFaceFeatureRequest() *OapiSmartdeviceFaceFeatureRequest {
	return &OapiSmartdeviceFaceFeatureRequest{}
}

type OapiSmartdeviceFaceFeatureRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFaceFeatureResponse
	ModelType       int64
	ModelVersion    string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartdeviceFaceFeatureRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFaceFeatureRequest) SetModelType(modelType2 int64) {
	this.ModelType = modelType2
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetModelType() int64 {
	return this.ModelType
}
func (this *OapiSmartdeviceFaceFeatureRequest) SetModelVersion(modelVersion2 string) {
	this.ModelVersion = modelVersion2
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetModelVersion() string {
	return this.ModelVersion
}
func (this *OapiSmartdeviceFaceFeatureRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.face.feature"
}
func (this *OapiSmartdeviceFaceFeatureRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFaceFeatureRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFaceFeatureRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_type"] = this.ModelType
	txtParams["model_version"] = this.ModelVersion
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartdeviceFaceFeatureRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelType, "modelType"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFaceFeatureRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFaceFeatureResponse struct {
	taobao.TaobaoResponse
	Errcode int64           `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	Result  []DidoFeatureVo `json:"result,omitempty"`
}
type DidoFeatureVo struct {
	FeatureInfo string `json:"feature_info,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
