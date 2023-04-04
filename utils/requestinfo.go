package utils

import (
	"net/http"
	"strings"

	"github.com/luozi-csu/lzblogs/model"
)

type RequestInfo struct {
	Path       string          // 请求路径
	Action     model.Operation // 资源操作
	APIPrefix  string          // API前缀
	APIVersion string          // API版本
	Resource   *model.Resource // 资源
}

func GetRequestInfo(req *http.Request) *RequestInfo {
	requestInfo := &RequestInfo{
		Path: req.URL.Path,
	}

	currentParts := splitPath(requestInfo.Path)
	// API前缀和版本占用两个part，因此长度小于3时为非资源请求
	if len(currentParts) < 3 {
		return requestInfo
	}

	requestInfo.APIPrefix = currentParts[0]
	currentParts = currentParts[1:]

	requestInfo.APIVersion = currentParts[0]
	currentParts = currentParts[1:]

	switch strings.ToLower(req.Method) {
	case "get", "head":
		requestInfo.Action = model.GetOperation
	case "post":
		requestInfo.Action = model.CreateOperation
	case "put":
		requestInfo.Action = model.UpdateOperation
	case "delete":
		requestInfo.Action = model.RemoveOperation
	default:
		requestInfo.Action = ""
	}

	resource := new(model.Resource)
	if len(currentParts) >= 2 {
		resource.Kind = currentParts[0]
		resource.Name = currentParts[1]
	} else if len(currentParts) >= 1 {
		resource.Kind = currentParts[0]
		if requestInfo.Action == model.GetOperation {
			requestInfo.Action = model.ListOperation
		}
	}

	return requestInfo
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return []string{}
	}
	return strings.Split(path, "/")
}
