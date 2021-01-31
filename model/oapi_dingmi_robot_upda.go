package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingmiRobotUpdateRequest() *OapiDingmiRobotUpdateRequest {
	return &OapiDingmiRobotUpdateRequest{}
}

type OapiDingmiRobotUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiRobotUpdateResponse
	TopHttpMethod   string
	TopResponseType string
	Type            string
	UpdateBotModel  string
}

func (this *OapiDingmiRobotUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiRobotUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiRobotUpdateRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiDingmiRobotUpdateRequest) GetType() string {
	return this.Type
}
func (this *OapiDingmiRobotUpdateRequest) SetUpdateBotModel(updateBotModel2 string) {
	this.UpdateBotModel = updateBotModel2
}
func (this *OapiDingmiRobotUpdateRequest) GetUpdateBotModel() string {
	return this.UpdateBotModel
}
func (this *OapiDingmiRobotUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.robot.update"
}
func (this *OapiDingmiRobotUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiRobotUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiRobotUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiRobotUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiRobotUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["type"] = this.Type
	txtParams["update_bot_model"] = this.UpdateBotModel
	return txtParams
}
func (this *OapiDingmiRobotUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Type, "type"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingmiRobotUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiRobotUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UpdateBotRequestDTO struct {
	Avatar          string `json:"avatar,omitempty"`
	Brief           string `json:"brief,omitempty"`
	Description     string `json:"description,omitempty"`
	Name            string `json:"name,omitempty"`
	PreviewMediaUrl string `json:"preview_media_url,omitempty"`
}
type OapiDingmiRobotUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
