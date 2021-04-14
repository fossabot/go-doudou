package router

import (
	"context"
	"encoding/json"
	"example/doudou/service"
	"example/doudou/vo"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/kit/ddl/query"
	"net/http"
	"strconv"
)

var decoder = schema.NewDecoder()

func handleError(w http.ResponseWriter, err error, status ...int) {
	logrus.Errorln(fmt.Sprintf("%+v", err))
	if len(status) > 0 {
		w.WriteHeader(status[0])
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err = json.NewEncoder(w).Encode(vo.Ret{
		Code: 1,
		Data: nil,
		Msg:  err.Error(),
	}); err != nil {
		panic(err)
	}
}

func postSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = r.ParseForm(); err != nil {
		handleError(w, err)
		return
	}

	var form vo.SignUpForm
	if err = decoder.Decode(&form, r.PostForm); err != nil {
		handleError(w, err)
		return
	}

	us := service.NewUserService()

	var uid int
	if uid, err = us.PostSignUp(context.Background(), form); err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(vo.Ret{
		Code: 0,
		Data: uid,
		Msg:  "",
	}); err != nil {
		handleError(w, err)
	}
}

func postLogInHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = r.ParseForm(); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	var form vo.LogInForm
	if err = decoder.Decode(&form, r.PostForm); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	us := service.NewUserService()

	var auth vo.Auth
	if auth, err = us.PostLogIn(context.Background(), form); err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(vo.Ret{
		Code: 0,
		Data: auth,
		Msg:  "",
	}); err != nil {
		handleError(w, err)
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var uid int
	if uid, err = strconv.Atoi(r.FormValue("id")); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	us := service.NewUserService()

	var userVo vo.UserVo
	if userVo, err = us.GetUser(context.Background(), uid); err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(vo.Ret{
		Code: 0,
		Data: userVo,
		Msg:  "",
	}); err != nil {
		handleError(w, err)
	}
}

func postPageUsersHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p vo.PageQuery
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	us := service.NewUserService()

	var pr query.PageRet

	if pr, err = us.PostPageUsers(context.Background(), p); err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(vo.Ret{
		Code: 0,
		Data: pr,
		Msg:  "",
	}); err != nil {
		handleError(w, err)
	}
}
