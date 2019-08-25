package yapi

import (
	"net/http"
)

type CatMenuData []CatData

type CatData struct {
	ID   int    `json:"_id" structs:"_id"`
	UID  int    `json:"uid" structs:"uid"`
	Name string `json:"name" structs:"name"`
	Desc string `json:"desc" structs:"desc"`
}

type CatMenu struct {
	ErrCode int         `json:"errcode" structs:"errcode"`
	ErrMsg  string      `json:"errmsg" structs:"errmsg"`
	Data    CatMenuData `json:"data" structs:"data"`
}

// CatMenuService .
type CatMenuService struct {
	client *Client
}

type CatMenuParam struct {
	Token     string `url:"token"`
	ProjectID int    `url:"project_id"`
}

func (s *CatMenuService) Get(projectId int) (*CatMenu, *http.Response, error) {
	apiEndpoint := "api/interface/getCatMenu"
	catMenuParam := new(CatMenuParam)
	catMenuParam.ProjectID = projectId
	catMenuParam.Token = s.client.Authentication.token
	url, err := addOptions(apiEndpoint, catMenuParam)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(CatMenu)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, NewServerError(resp, err)
	}
	return result, resp, err
}
