package test

import "testing"
import (
	"fmt"
	"github.com/zhenyangjiang/qcloud-sms/QcloudSms"
	"io/ioutil"
	"net/http"
)

func TestSend(t *testing.T) {
	qcloudsms, _ := QcloudSms.NewQcloudSms(12312, "xxxx")
	resp, err := qcloudsms.SmsSingleSender.SendWithParam(86, "xxxx", 375634, []string{"12312"}, "xxxx", "", "", callback)
	if err == nil {
		var data []byte
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(string(data))
		}
	}
}
func callback(err error, resp *http.Response, resData string) {
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("response data: ", resData)
	}
}
