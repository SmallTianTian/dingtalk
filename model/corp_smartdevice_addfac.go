package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpSmartdeviceAddfaceRequest() *CorpSmartdeviceAddfaceRequest {
	return &CorpSmartdeviceAddfaceRequest{}
}

type CorpSmartdeviceAddfaceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpSmartdeviceAddfaceResponse
	FaceVo          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpSmartdeviceAddfaceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpSmartdeviceAddfaceRequest) SetFaceVo(faceVo2 string) {
	this.FaceVo = faceVo2
}
func (this *CorpSmartdeviceAddfaceRequest) GetFaceVo() string {
	return this.FaceVo
}
func (this *CorpSmartdeviceAddfaceRequest) GetApiMethodName() string {
	return "dingtalk.corp.smartdevice.addface"
}
func (this *CorpSmartdeviceAddfaceRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpSmartdeviceAddfaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpSmartdeviceAddfaceRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpSmartdeviceAddfaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpSmartdeviceAddfaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpSmartdeviceAddfaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["face_vo"] = this.FaceVo
	return txtParams
}
func (this *CorpSmartdeviceAddfaceRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpSmartdeviceAddfaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpSmartdeviceAddfaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DidoFaceVO struct {
	EndDate   int64  `json:"end_date,omitempty"`
	MediaId   string `json:"media_id,omitempty"`
	StartDate int64  `json:"start_date,omitempty"`
	UserType  string `json:"user_type,omitempty"`
	Userid    string `json:"userid,omitempty"`
}
type CorpSmartdeviceAddfaceResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
