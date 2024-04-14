package controllers

import (
	sw "main/go-server-generated/go"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"

	"main/core/application"
	"main/core/repository"
)

const (
	RoleHeaderKey        = "Role"
	InvalidBodyDataErr   = "{\"error\": \"Invalid body data\"}"
	InvalidURIDataErr    = "{\"error\": \"Invalid URI data\"}"
	NotExistsResourceErr = "{\"error\": \"Resource does not exists\"}"
	InternalServerErr    = "{\"error\": \"Internal error occured\"}"
	AbsenceOfQueryParam  = "{\"error\": \"Absence of required query param\"}"
	CacheExpiringMinutes = time.Minute * 5
)

type ControllerREST struct {
	router *mux.Router
	app    *application.Application
	cache  *redis.Client
}

func NewControllerREST(storage repository.IRepository) *ControllerREST {
	cache := redis.NewClient(&redis.Options{
		Addr:     "app_cache:6379",
		Password: "",
		DB:       0,
	})

	router := sw.NewRouter()
	ctrl := &ControllerREST{
		router: router,
		app:    application.NewApplication(storage),
		cache:  cache,
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
