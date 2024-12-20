package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/cpuinfo"
	"time"
)

type CPUInfoHandler struct {
	svc cpuinfo.CPUInfoService
}

func NewCPUInfoHandler(svc cpuinfo.CPUInfoService) *CPUInfoHandler {
	return &CPUInfoHandler{svc: svc}
}

type cpuInfoResp struct {
	Id   string  `json:"id"`
	Load float32 `json:"load"`
}

func (cih *CPUInfoHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting CPU info")

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	cich := make(chan []byte)
	prevData, _ := cih.svc.GetCPUInfo()
	time.Sleep(2 * time.Second)
	go func() {
		for {
			cpuInfos, _ := cih.svc.GetCPUInfo()

			resp := []cpuInfoResp{}
			for i, cpuInfo := range cpuInfos {
				// https://stackoverflow.com/a/23376195
				prevIdle := prevData[i].Idle + prevData[i].Iowait
				idle := cpuInfo.Idle + cpuInfo.Iowait

				prevNonIdle := prevData[i].User + prevData[i].Nice + prevData[i].System + prevData[i].Irq + prevData[i].Softirq + prevData[i].Steal
				nonIdle := cpuInfo.User + cpuInfo.Nice + cpuInfo.System + cpuInfo.Irq + cpuInfo.Softirq + cpuInfo.Steal

				prevTotal := prevIdle + prevNonIdle
				total := idle + nonIdle

				totalDiff := total - prevTotal
				idleDiff := idle - prevIdle

				resp = append(resp, cpuInfoResp{Id: cpuInfo.Id, Load: float32(totalDiff-idleDiff) / float32(totalDiff)})
			}

			respJson, _ := json.Marshal(resp)

			cich <- respJson

			prevData = cpuInfos
			time.Sleep(2 * time.Second)
		}
	}()

	for ci := range cich {
		conn.WriteMessage(1, ci)
	}
}
