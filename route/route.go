package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zzjbattlefield/delay_queue/queue"
)

type IDRequest struct {
	ID string `json:"id"`
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func PushJob(w http.ResponseWriter, r *http.Request) {
	jobInfo := &queue.JobItem{}
	err := readBody(w, r, &jobInfo)
	if err != nil {
		return
	}
	jobInfo.ID = strings.TrimSpace(jobInfo.ID)
	jobInfo.Body = strings.TrimSpace(jobInfo.Body)
	jobInfo.Topic = strings.TrimSpace(jobInfo.Topic)
	jobInfo.Delay = time.Now().Unix() + jobInfo.Delay
	err = queue.Push(*jobInfo)
	if err != nil {
		w.Write(createResponseBody(500, "添加job失败", nil))
		return
	} else {
		w.Write(createResponseBody(200, "添加job成功", jobInfo))
	}
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	idRequest := new(IDRequest)
	err := readBody(w, r, &idRequest)
	if err != nil {
		return
	}
	if err = queue.Remove(idRequest.ID); err != nil {
		w.Write(createResponseBody(500, "删除job失败", nil))
		return
	}
	w.Write(createResponseBody(200, "删除job成功", nil))
}

func readBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("读取http body失败 # %s", err.Error())
		w.Write(createResponseBody(400, "读取http body失败", nil))
		return err
	}
	return json.Unmarshal(body, v)
}

func createResponseBody(code int, msg string, data interface{}) []byte {
	responseBody := ResponseBody{Code: code, Message: msg, Data: data}
	bodyByte, err := json.Marshal(responseBody)
	if err != nil {
		log.Printf("生成response body,转换json失败#%s", err.Error())
		return []byte(`{"code":"1", "message": "生成响应body异常", "data":[]}`)
	}
	return bodyByte
}
