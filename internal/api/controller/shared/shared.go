package shared

import (
	"fmt"
	"net/http"

	"github.com/errntry/errntry/internal/i18n"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log = &logrus.Logger{}

func logError(prefix string, err error) {
	if err != nil {
		Log.Errorf("%s: %s", prefix, err.Error())
	}
}

func Abort(c *gin.Context, code int, msg string) {
	resp := i18n.NewResponse(code, msg)

	Log.Debugf("api: abort %s with code %d (%s)", c.FullPath(), code, resp.String())

	c.AbortWithStatusJSON(code, resp)
}

func Error(c *gin.Context, code int, err error, msg string) {
	resp := i18n.NewResponse(code, msg)

	if err != nil {
		resp.Details = err.Error()
		Log.Errorf("api: error %q with code %d in %s (%s)", err.Error(), code, c.FullPath(), resp.String())
	}

	c.AbortWithStatusJSON(code, resp)
}

func AbortUnauthorized(c *gin.Context, s string) {
	Abort(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized request: %s", s))
}

func AbortEntityNotFound(c *gin.Context, s string) {
	Abort(c, http.StatusNotFound, fmt.Sprintf("%s Not Found", s))
}

func AbortSaveFailed(c *gin.Context) {
	Abort(c, http.StatusInternalServerError, "Changes could not be saved")
}

func AbortDeleteFailed(c *gin.Context) {
	Abort(c, http.StatusInternalServerError, "Could not be deleted")
}

func AbortUnexpected(c *gin.Context) {
	Abort(c, http.StatusInternalServerError, "Unexpected error, please try again")
}

func AbortBadRequest(c *gin.Context, s string) {
	Abort(c, http.StatusBadRequest, fmt.Sprintf("Invalid Request: %s", s))
}

func AbortAlreadyExists(c *gin.Context, s string) {
	Abort(c, http.StatusConflict, fmt.Sprintf("%s Already Exists", s))
}

func AbortFeatureDisabled(c *gin.Context) {
	Abort(c, http.StatusForbidden, "Feature Disabled")
}
