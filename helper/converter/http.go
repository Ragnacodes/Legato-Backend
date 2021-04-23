package converter

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)

func HttpDbToHttpInfo(s legatoDb.Http) api.HttpInfo {
	hi := api.HttpInfo{}
	hi.Id = s.ID
	hi.Url = s.Url
	hi.Method = s.Method

	return hi
}

func DataToHttp(data interface{}) legatoDb.Http {
	var w legatoDb.Http
	if data != nil {
		d := data.(map[string]interface{})
		w.Url = d["url"].(string)
		w.Method = d["method"].(string)
	}

	return w
}

func HttpDbToServiceNode(h legatoDb.Http) api.ServiceNode {
	var sn api.ServiceNode
	sn = ServiceDbToServiceNode(h.Service)
	// Http data
	sn.Data = HttpDbToHttpInfo(h)

	return sn
}
