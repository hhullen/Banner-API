package controllers

import (
	sw "main/go-server-generated/go"
	"net/http"

	"github.com/gorilla/mux"

	"main/core/application"
	"main/infrastructure/storage"
)

const (
	RoleHeaderKey     = "Role"
	InvalidDataErr    = "{\"error\": \"Invalid data\"}"
	InternalServerErr = "{\"error\": \"Internal error occured\"}"
)

type ControllerREST struct {
	router *mux.Router
	app    *application.Application
}

func NewControllerREST() *ControllerREST {
	router := sw.NewRouter()

	database := &storage.MOCKDB{}
	ctrl := &ControllerREST{
		router: router,
		app:    application.NewApplication(database),
	}

	router.GetRoute("UserBannerGet").HandlerFunc(ctrl.handleUserBannerGet)
	router.GetRoute("BannerGet").HandlerFunc(ctrl.handleBannerGet)
	router.GetRoute("BannerIdDelete").HandlerFunc(ctrl.handleBannerIdDelete)
	router.GetRoute("BannerIdPatch").HandlerFunc(ctrl.handleBannerIdPatch)
	router.GetRoute("BannerPost").HandlerFunc(ctrl.handleBannerPost)
	router.GetRoute("Index").HandlerFunc(ctrl.handleIndex)
	router.Use(ctrl.handleMiddleware)

	return ctrl
}

func (me *ControllerREST) Serve(address string) error {
	return http.ListenAndServe(":8080", me.router)
}
