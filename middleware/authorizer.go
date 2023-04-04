package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/common"
	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/repository"
	"github.com/luozi-csu/lzblogs/utils"
	"github.com/pkg/errors"
)

var (
	showableResource = []string{"articles"}
)

func Authorizer(rbac repository.RBACRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestInfo := utils.GetRequestInfo(c.Request)
		// 非资源请求
		if requestInfo.Resource == nil {
			c.Next()
		}

		var isShowableResource bool
		for _, kind := range showableResource {
			if requestInfo.Resource.Kind == kind {
				isShowableResource = true
				break
			}
		}

		action := requestInfo.Action
		if isShowableResource && (action == model.GetOperation || action == model.ListOperation) {
			c.Next()
		}

		user := common.GetUser(c)
		if user == nil {
			common.ResponseFailed(c, http.StatusUnauthorized, errors.New("empty user info"))
			c.Abort()
			return
		}

		resource, err := json.Marshal(*requestInfo.Resource)
		if err != nil {
			common.ResponseFailed(c, http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		ok, err := rbac.HasPermission(user, string(action), string(resource))
		if !ok || err != nil {
			common.ResponseFailed(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
