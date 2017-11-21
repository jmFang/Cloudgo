package service

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", nil)
	}
}

func submit(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.URL)
		formatter.HTML(w, http.StatusOK, "result", struct {
			Huawei string `json:"huawei"`
			Oppo   string `json:"oppo"`
			Xiaomi string `json:"xiaomi"`
			Ov     string `json:"ov"`
			Apple  string `json:"apple"`
		}{
			Huawei: r.FormValue("huawei"),
			Oppo:   r.FormValue("oppo"),
			Xiaomi: r.FormValue("xiaomi"),
			Ov:     r.FormValue("ov"),
			Apple:  r.FormValue("apple")})
	}
}

func apiFetchDataHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.Text(w, http.StatusOK, "更多精彩请积极认真学习golang")
	}
}
