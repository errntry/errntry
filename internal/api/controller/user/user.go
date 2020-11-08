package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/errntry/errntry/internal/api/conn"
	"github.com/errntry/errntry/internal/api/controller/shared"
	"github.com/errntry/errntry/internal/api/model/user"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserCollection statically declared
const UserCollection = "user"

func Index(ctx *gin.Context) {
	db := conn.GetMongoDB()
	users := user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		shared.Abort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &users})
}

func Get(ctx *gin.Context) {
	var id = bson.ObjectIdHex(ctx.Param("id")) // Get Param
	userInfo, err := user.Info(id, UserCollection)
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			shared.AbortEntityNotFound(ctx, "User")
			return
		}
		shared.AbortBadRequest(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &userInfo})
}

// TODO: Create should be removed after creating a proper signup endpoint
func Create(ctx *gin.Context) {
	db := conn.GetMongoDB()
	userInfo := user.User{}
	err := ctx.Bind(&userInfo)
	if err != nil {
		shared.AbortBadRequest(ctx, err.Error())
		return
	}
	userInfo.ID = bson.NewObjectId()
	userInfo.CreatedAt = time.Now()
	userInfo.UpdatedAt = time.Now()
	err = db.C(UserCollection).Insert(userInfo)
	if err != nil {
		shared.AbortBadRequest(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": &userInfo})
}

func Update(ctx *gin.Context) {
	db := conn.GetMongoDB()
	var id = bson.ObjectIdHex(ctx.Param("id")) // Get Param
	userInfo, err := user.Info(id, UserCollection)
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			shared.AbortEntityNotFound(ctx, "User")
			return
		}
		shared.AbortBadRequest(ctx, err.Error())
		return
	}
	// user := user.User{}
	err = ctx.Bind(&userInfo)
	if err != nil {
		shared.AbortBadRequest(ctx, err.Error())
		return
	}
	userInfo.ID = id
	userInfo.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": &id}, userInfo)
	if err != nil {
		shared.AbortBadRequest(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &userInfo})
}

func Delete(ctx *gin.Context) {
	db := conn.GetMongoDB()

	var id = bson.ObjectIdHex(ctx.Param("id")) // Get Param
	err := db.C(UserCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			shared.AbortEntityNotFound(ctx, "User")
			return
		}
		shared.AbortBadRequest(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}
