package gografana

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	gohttp "github.com/xm-chentl/go-http"
)

const (
	ContentType = "application/json"
)

type Option struct {
	Host     string
	Token    string
	FolderID int32
	Refresh  string
}

type grafanaInst struct {
	httpInst gohttp.IHttp

	Opt Option
}

func (g grafanaInst) DeleteDashboard(uid string) (err error) {
	res := ResponseBase{}
	resp, err := g.httpInst.SetMethod(http.MethodDelete).SetURL(
		fmt.Sprintf("%s/api/dashboards/uid/%s", g.Opt.Host, uid),
	).SetHeader(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", g.Opt.Token),
		},
	).Send(&res)
	if err != nil {
		return
	}
	if resp.Code == http.StatusOK || resp.Code == 404 {
		return
	}

	// todo: 优化重构为抽象工厂
	switch resp.Code {
	case 401:
		err = errors.New("访问grafana未经授权")
	case 403:
		err = errors.New("访问grafana被拒绝")
	default:
		err = errors.New("操作失败")
	}

	return
}

func (g grafanaInst) PreviewDashboard(req RequestPreviewDashboard) string {
	var fromTimestamp, toTimestamp int64
	currentTime := time.Now()
	if req.FromTime != nil {
		fromTimestamp = req.FromTime.Unix() * 1000
	} else {
		fromTimestamp = currentTime.Add(-time.Hour*6).UnixNano() / 1e6
	}
	if req.ToTime != nil {
		toTimestamp = req.ToTime.Unix() * 1000
	} else {
		toTimestamp = currentTime.UnixNano() / 1e6
	}

	urlVals := make([]string, 0)
	var urlVarParam string
	if len(req.Vars) > 0 {
		for k, vs := range req.Vars {
			for _, v := range vs {
				urlVals = append(urlVals, fmt.Sprintf("var-%s=%s", k, v))
			}
		}
		urlVarParam = "&" + strings.Join(urlVals, "&")
	}

	return fmt.Sprintf(
		"%s%s?orgId=1&refresh=%s%s&from=%d&to=%d&panelId=%d",
		g.Opt.Host,
		strings.ReplaceAll(req.Url, "/d/", "/d-solo/"),
		g.Opt.Refresh,
		urlVarParam,
		fromTimestamp,
		toTimestamp,
		req.PanelID,
	)
}

func (g grafanaInst) SaveDashboard(info DashboardInfo, panels ...IDashboardPanel) (res *ResponseCreateOrUpdateDashboard, err error) {
	args, err := generateDashboardArgs(info, panels...)
	if err != nil {
		return
	}

	res = &ResponseCreateOrUpdateDashboard{}
	resp, err := g.httpInst.SetURL(
		fmt.Sprintf("%s/api/dashboards/db", g.Opt.Host),
	).SetHeader(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", g.Opt.Token),
		},
	).SetBody(args).Send(res)
	if err != nil {
		return
	}
	if resp.Code == http.StatusOK {
		return
	}

	var respErrorItems []ResponseErrorItem
	if err = json.Unmarshal(resp.Body, &respErrorItems); err != nil {
		return
	}

	respError := ResponseErrorItem{}
	if len(respErrorItems) > 0 {
		respError = respErrorItems[0]
	}

	// todo: 优化重构为抽象工厂
	switch resp.Code {
	case 400:
		err = errors.New("创建的请求数据(json)不符合grafana要求")
	case 401:
		err = errors.New("访问grafana未经授权")
	case 403:
		err = errors.New("访问grafana被拒绝")
	case 412:
		err = fmt.Errorf("前提条件失败: %s", respError.Message)
	}

	return
}

func (g grafanaInst) CreateFolder(folderName string) (res *ResponseCreateOrUpdateFolder, err error) {
	res = &ResponseCreateOrUpdateFolder{}
	resp, err := g.httpInst.SetURL(
		fmt.Sprintf("%s/api/folders", g.Opt.Host),
	).SetHeader(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", g.Opt.Token),
		},
	).SetBody(map[string]interface{}{
		"title": folderName,
	}).Send(&res)
	if err != nil {
		return
	}
	if resp.Code == http.StatusOK {
		return
	}

	// todo: 优化重构为抽象工厂
	switch resp.Code {
	case 400:
		err = errors.New("创建的请求数据不符合grafana要求")
	case 401:
		err = errors.New("访问grafana未经授权")
	case 403:
		err = errors.New("访问grafana被拒绝")
	case 409:
		err = errors.New("文件夹已存在")
	}

	return
}

func (g grafanaInst) SaveFolder(uid string, folderName string) (res *ResponseCreateOrUpdateFolder, err error) {
	res = &ResponseCreateOrUpdateFolder{}
	resp, err := g.httpInst.SetURL(
		fmt.Sprintf("%s/api/folders/%s", g.Opt.Host, uid),
	).SetHeader(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", g.Opt.Token),
		},
	).SetBody(map[string]interface{}{
		"title":     folderName,
		"overwrite": true,
	}).SetMethod(http.MethodPut).Send(&res)
	if err != nil {
		return
	}
	if resp.Code == http.StatusOK {
		return
	}

	// todo: 优化重构为抽象工厂
	switch resp.Code {
	case 400:
		err = errors.New("创建的请求数据不符合grafana要求")
	case 401:
		err = errors.New("访问grafana未经授权")
	case 403:
		err = errors.New("访问grafana被拒绝")
	case 404:
		err = errors.New("未找到文件夹")
	case 412:
		err = errors.New("请求失败, grafana: " + res.Message)
	}

	return
}

func (g grafanaInst) DeleteFolder(uid string) (res *ResponseDeleteFolder, err error) {
	res = &ResponseDeleteFolder{}
	resp, err := g.httpInst.SetURL(
		fmt.Sprintf("%s/api/folders/%s", g.Opt.Host, uid),
	).SetHeader(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", g.Opt.Token),
		},
	).SetMethod(http.MethodDelete).Send(&res)
	if err != nil {
		return
	}
	if resp.Code == http.StatusOK {
		return
	}

	// todo: 优化重构为抽象工厂
	switch resp.Code {
	case 400:
		err = errors.New("数据错误，请求数据不符合grafana")
	case 401:
		err = errors.New("访问grafana未经授权")
	case 403:
		err = errors.New("访问grafana被拒绝")
	case 404:
		err = errors.New("在grafana未找到文件夹")
	}
	return
}

func (g grafanaInst) GetFolder(id int32) (res *ResponseCreateOrUpdateFolder, err error) {
	res = &ResponseCreateOrUpdateFolder{}
	resp, err := g.httpInst.SetURL(
		fmt.Sprintf("%s/api/folders/id/%d", g.Opt.Host, id),
	).SetHeader(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", g.Opt.Token),
		},
	).SetMethod(http.MethodGet).Send(&res)
	if err != nil {
		return
	}
	if resp.Code == http.StatusOK {
		return
	}

	// todo: 优化重构为抽象工厂
	switch resp.Code {
	case 400:
		err = errors.New("创建的请求数据不符合grafana要求")
	case 401:
		err = errors.New("访问grafana未经授权")
	case 403:
		err = errors.New("访问grafana被拒绝")
	case 404:
		err = errors.New("在grafana未找到文件夹")
	}
	return
}

func (g grafanaInst) PanelFactory() IDashboardPanelFactory {
	return &panelFactory{}
}

func New(opt Option, httpInst gohttp.IHttp) IGrafana {
	return &grafanaInst{
		httpInst: httpInst,
		Opt:      opt,
	}
}
