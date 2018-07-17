package service

import (
	"encoding/json"
	"github.com/SwitchCollector/g"
	"github.com/SwitchCollector/service/flow"
	"github.com/SwitchCollector/service/visit"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func FlowQuantityBytes(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse data error", http.StatusBadRequest)
		return
	}
	period := strings.TrimSpace(r.Form.Get("Period"))
	expire, err := strconv.ParseInt(period, 10, 64)
	if err != nil {
		http.Error(w, "param Period format error.", http.StatusBadRequest)
		return
	}
	flowQuantity := flow.Query(expire * 60)
	ret, err1 := json.Marshal(flowQuantity)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ret)

}

func VisitLog(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse data error", http.StatusBadRequest)
		return
	}
	period := strings.TrimSpace(r.Form.Get("Period"))
	expire, err := strconv.ParseInt(period, 10, 64)
	if err != nil {
		http.Error(w, "param Period format error.", http.StatusBadRequest)
		return
	}
	visitLog := visit.Search(expire * 60)
	ret, err1 := json.Marshal(visitLog)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ret)
}

func GetDeviceList(w http.ResponseWriter, r *http.Request) {
	deviceList := g.GetDeviceList()
	ret, err := json.Marshal(deviceList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ret)
}

func GetDeviceInfo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse data error", http.StatusBadRequest)
		return
	}
	ip := strings.TrimSpace(r.Form.Get("ip"))
	deviceInfo := g.GetDeviceInfo(ip)
	ret, err := json.Marshal(deviceInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ret)
}

func GetInterfaceInfo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse data error", http.StatusBadRequest)
		return
	}
	filter := strings.TrimSpace(r.Form.Get("filter"))
	ip := strings.TrimSpace(r.Form.Get("ip"))
	deviceInfo := g.GetInterfaceInfo(ip, filter)
	ret, err := json.Marshal(deviceInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ret)
}

func GetInterfaceMetric(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse data error", http.StatusBadRequest)
		return
	}
	filter := strings.TrimSpace(r.Form.Get("filter"))
	period := strings.TrimSpace(r.Form.Get("period"))
	p, _ := strconv.ParseInt(period, 10, 64)
	ip := strings.TrimSpace(r.Form.Get("ip"))
	deviceInfo := g.GetInterfaceMetric(ip, filter, p)
	ret, err := json.Marshal(deviceInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ret)
}

func StartHttpServer() {
	log.Println("starting http server....")
	http.HandleFunc("/GetDeviceList", GetDeviceList)
	http.HandleFunc("/GetDeviceInfo", GetDeviceInfo)
	http.HandleFunc("/GetInterfaceInfo", GetInterfaceInfo)
	http.HandleFunc("/GetInterfaceMetric", GetInterfaceMetric)
	http.ListenAndServe(g.Config().Http.Addr, nil)
}
