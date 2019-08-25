package yapi

import (
	"net/http"
)

// InterfaceService .
type InterfaceService struct {
	client *Client
}

type ReqKVItemSimple struct {
	Name    string `json:"name" structs:"name"`
	Example string `json:"example" structs:"example"`
	Desc    string `json:"desc" structs:"desc"`
}

type ReqKVItemDetail struct {
	ReqKVItemSimple
	Type     string `json:"type" structs:"type"`
	Required string `json:"required" structs:"required"`
}

type InterfaceData struct {
	ID                  int               `json:"_id" structs:"_id"`
	ProjectID           int               `json:"project_id" structs:"project_id"`
	CatID               int               `json:"catid" structs:"catid"`
	UID                 int               `json:"uid" structs:"uid"`
	EditUID             int               `json:"edit_uid" structs:"edit_uid"`
	AddTime             int               `json:"add_time" structs:"add_time"`
	UpTime              int               `json:"up_time" structs:"up_time"`
	Status              string            `json:"status" structs:"status"`
	Title               string            `json:"title" structs:"title"`
	Path                string            `json:"path" structs:"path"`
	Method              string            `json:"method" structs:"method"`
	ResBody             string            `json:"res_body" structs:"res_body"`
	ResBodyType         string            `json:"res_body_type" structs:"res_body_type"`
	ResBodyIsJsonSchema bool              `json:"res_body_is_json_schema" structs:"res_body_is_json_schema"`
	ReqBodyType         string            `json:"req_body_type" structs:"req_body_type"`
	ReqParams           []ReqKVItemSimple `json:"req_params" structs:"req_params"`
	ReqHeaders          []ReqKVItemSimple `json:"req_headers" structs:"req_headers"`
	ReqQuery            []ReqKVItemDetail `json:"req_query" structs:"req_query"`
	ReqBodyForm         []ReqKVItemDetail `json:"req_body_form" structs:"req_body_form"`
}

type Interface struct {
	ErrCode int           `json:"errcode" structs:"errcode"`
	ErrMsg  string        `json:"errmsg" structs:"errmsg"`
	Data    InterfaceData `json:"data" structs:"data"`
}

type InterfaceParam struct {
	Token string `url:"token"`
	ID    int    `url:"id"`
}

type InterfaceListData struct {
	Count int             `json:"count" structs:"count"`
	Total int             `json:"total" structs:"total"`
	List  []InterfaceData `json:"list" structs:"list"`
}

type InterfaceList struct {
	ErrCode int               `json:"errcode" structs:"errcode"`
	ErrMsg  string            `json:"errmsg" structs:"errmsg"`
	Data    InterfaceListData `json:"data" structs:"data"`
}

type InterfaceListParam struct {
	Token string `url:"token,omitempty"`
	CatID int    `url:"catid,omitempty"`
	Page  int    `url:"Page"`
	Limit int    `url:"limit"`
}

func (s *InterfaceService) GetList(opt *InterfaceListParam) (*InterfaceList, *http.Response, error) {
	apiEndpoint := "api/interface/list_cat"
	opt.Token = s.client.Authentication.token
	url, err := addOptions(apiEndpoint, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(InterfaceList)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, NewServerError(resp, err)
	}
	return result, resp, err
}

func (s *InterfaceService) Get(id int) (*Interface, *http.Response, error) {
	apiEndpoint := "api/interface/get"
	interfaceParam := new(InterfaceParam)
	interfaceParam.ID = id
	interfaceParam.Token = s.client.Authentication.token
	url, err := addOptions(apiEndpoint, interfaceParam)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json	")
	result := new(Interface)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, NewServerError(resp, err)
	}
	return result, resp, err
}
