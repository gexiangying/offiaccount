package poi

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAddpoi        = "/cgi-bin/poi/addpoi"
	apiGetpoi        = "/cgi-bin/poi/getpoi"
	apiGetPoiList    = "/cgi-bin/poi/getpoilist"
	apiUpdatepoi     = "/cgi-bin/poi/updatepoi"
	apiDelpoi        = "/cgi-bin/poi/delpoi"
	apiGetWXCategory = "/cgi-bin/poi/getwxcategory"
)

/*
创建门店

创建门店接口是为商户提供创建自己门店数据的接口，门店数据字段越完整，商户页面展示越丰富，越能够吸引更多用户，并提高曝光度

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST http://api.weixin.qq.com/cgi-bin/poi/addpoi?access_token=TOKEN
*/
func Addpoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddpoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询门店信息

创建门店后获取poi_id 后，商户可以利用poi_id，查询具体某条门店的信息。 若在查询时，update_status 字段为1，表明在5 个工作日内曾用update 接口修改过门店扩展字段，该扩展字段为最新的修改字段，尚未经过审核采纳，因此不是最终结果。最终结果会在5 个工作日内，最终确认是否采纳，并前端生效（但该扩展字段的采纳过程不影响门店的可用性，即available_state仍为审核通过状态）

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST http://api.weixin.qq.com/cgi-bin/poi/getpoi?access_token=TOKEN
*/
func Getpoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetpoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询门店列表

商户可以通过该接口，批量查询自己名下的门店list，并获取已审核通过的poiid、商户自身sid 用于对应、商户名、分店名、地址字段

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token=TOKEN
*/
func GetPoiList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetPoiList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改门店服务信息

商户可以通过该接口，修改门店的服务信息，包括：sid、图片列表、营业时间、推荐、特色服务、简介、人均价格、电话8个字段（名称、坐标、地址等不可修改）修改后需要人工审核

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token=TOKEN
*/
func Updatepoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdatepoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除门店

商户可以通过该接口，删除已经成功创建的门店。请商户慎重调用该接口

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token=TOKEN
*/
func Delpoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelpoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
门店类目表

类目名称接口是为商户提供自己门店类型信息的接口。门店类目定位的越规范，能够精准的吸引更多用户，提高曝光率

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST http://api.weixin.qq.com/cgi-bin/poi/getwxcategory?access_token=TOKEN
*/
func GetWXCategory(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetWXCategory, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
