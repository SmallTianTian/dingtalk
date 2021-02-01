package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSceneservicegroupGroupsetCreateRequest() *OapiSceneservicegroupGroupsetCreateRequest {
	return &OapiSceneservicegroupGroupsetCreateRequest{}
}

type OapiSceneservicegroupGroupsetCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSceneservicegroupGroupsetCreateResponse
	GroupTemplateid string
	GroupsetName    string
	OpenTeamid      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSceneservicegroupGroupsetCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) SetGroupTemplateid(groupTemplateid2 string) {
	this.GroupTemplateid = groupTemplateid2
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetGroupTemplateid() string {
	return this.GroupTemplateid
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) SetGroupsetName(groupsetName2 string) {
	this.GroupsetName = groupsetName2
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetGroupsetName() string {
	return this.GroupsetName
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) SetOpenTeamid(openTeamid2 string) {
	this.OpenTeamid = openTeamid2
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetOpenTeamid() string {
	return this.OpenTeamid
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sceneservicegroup.groupset.create"
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_templateid"] = this.GroupTemplateid
	txtParams["groupset_name"] = this.GroupsetName
	txtParams["open_teamid"] = this.OpenTeamid
	return txtParams
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenTeamid, "openTeamid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSceneservicegroupGroupsetCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSceneservicegroupGroupsetCreateResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
