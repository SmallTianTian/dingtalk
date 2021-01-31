package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMiniappAppinfoQueryRequest() *OapiMiniappAppinfoQueryRequest {
	return &OapiMiniappAppinfoQueryRequest{}
}

type OapiMiniappAppinfoQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMiniappAppinfoQueryResponse
	ModelKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMiniappAppinfoQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMiniappAppinfoQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMiniappAppinfoQueryRequest) SetModelKey(modelKey2 string) {
	this.ModelKey = modelKey2
}
func (this *OapiMiniappAppinfoQueryRequest) GetModelKey() string {
	return this.ModelKey
}
func (this *OapiMiniappAppinfoQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.miniapp.appinfo.query"
}
func (this *OapiMiniappAppinfoQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMiniappAppinfoQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMiniappAppinfoQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMiniappAppinfoQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMiniappAppinfoQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_key"] = this.ModelKey
	return txtParams
}
func (this *OapiMiniappAppinfoQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelKey, "modelKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMiniappAppinfoQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMiniappAppinfoQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMiniappAppinfoQueryResponse struct {
	taobao.TaobaoResponse
	Data    AppInfoDoModel `json:"data,omitempty"`
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
}
type AppInfoDoModel struct {
	Alias         string `json:"alias,omitempty"`
	AppChannel    int64  `json:"app_channel,omitempty"`
	AppId         string `json:"app_id,omitempty"`
	AppKey        string `json:taobao.APP_KE,omitemptyY`
	AppType       int64  `json:"app_type,omitempty"`
	AutoInstall   int64  `json:"auto_install,omitempty"`
	ClientId      int64  `json:"client_id,omitempty"`
	Desc          string `json:"desc,omitempty"`
	DevId         string `json:"dev_id,omitempty"`
	EnglishName   string `json:"english_name,omitempty"`
	GmtCreate     int64  `json:"gmt_create,omitempty"`
	GmtModified   int64  `json:"gmt_modified,omitempty"`
	IconUrl       string `json:"icon_url,omitempty"`
	Id            int64  `json:"id,omitempty"`
	InheritConfig string `json:"inherit_config,omitempty"`
	InstId        int64  `json:"inst_id,omitempty"`
	IsDeleted     int64  `json:"is_deleted,omitempty"`
	LastDiscards  string `json:"last_discards,omitempty"`
	MainUrl       string `json:"main_url,omitempty"`
	Name          string `json:"name,omitempty"`
	Origin        int64  `json:"origin,omitempty"`
	Preset        int64  `json:"preset,omitempty"`
	Size          int64  `json:"size,omitempty"`
	Slogan        string `json:"slogan,omitempty"`
	Status        int64  `json:"status,omitempty"`
	SubType       int64  `json:"sub_type,omitempty"`
	VHost         string `json:"v_host,omitempty"`
}
