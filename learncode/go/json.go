package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	type Road struct {
		Name    string `json:"n"`
		Number  int    `json:",omitempty"`
		Address string `json:"-"`
		Phone   uint64 `json:"-,"`
	}
	roads := []Road{
		{"Diamond Fork", 29, "beijing", 21321321},
		{Name: "Sheep Creek", Address: "chongqing", Phone: 12345},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "  ")
	out.WriteTo(os.Stdout)

	type Result struct {
		ApiCode int32  `json:"apiCode"`
		ApiMsg  string `json:"apiMsg"`
		Data    struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	}

	type Result1 struct {
		ApiCode int32  `json:"apiCode"`
		ApiMsg  string `json:"apiMsg"`
	}

	re := Result1{
		ApiCode: 14,
		ApiMsg:  "shit",
	}

	m, err := json.Marshal(re)
	if err != nil {
		fmt.Printf("marshal result failed, err:%v", err)
		return
	}
	fmt.Printf("marshal %v\n", string(m))

	fuck := &Result{}
	if err = json.Unmarshal([]byte(m), &fuck); err != nil {
		fmt.Printf("unmarshal result failed, err:%v", err)
		return
	}
	fmt.Printf("unmarshal %v\n", fuck)

	shit1 := map[string]interface{}{
		"shit1":  "fuck",
		"number": 32,
		"float":  34.5678,
	}
	m, err = json.Marshal(shit1)
	if err != nil {
		fmt.Printf("marshal result failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(m))

	fmt.Printf("tostring %s\n", toString(shit1))

	str := `{"cell":"50630066003","code":"888888","cpf":"506.300.660-03","email":"sg@%%","scene":"14","smstype":0,"appversion":"6.2.0","channel":"0","city_id":"1","client_tag":"didi","country_id":"%2B55","datatype":1,"imei":"354117073210795CF4970A24B3272E5FD928B2868E3CEC6","lang":"pt-BR","lat":"0.0","lng":"0.0","loc_country":-1,"maptype":"soso","model":"Moto G (4)","networkType":"WIFI","origin_id":"5","os":"7.0","role":2,"source":0,"suuid":"37EC38521A8D2FDEEF66DF4FB81A49D7","vcode":620}`
	//var out1 map[string]interface{}
	out1 := DecodeJsonSafe(0, []byte(str))
	fmt.Println(out1)
	fmt.Println(out1["scene"])
}

func DecodeJsonSafe(logid int64, data []byte) (parameters map[string]interface{}) {
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)
	decoder.UseNumber()
	var v interface{}
	if err := decoder.Decode(&v); err == nil {
		var ok bool
		if parameters, ok = v.(map[string]interface{}); ok {
			return
		}
	} else {
		fmt.Printf("%d, decode err %v %s", logid, err, string(data))
	}
	return nil
}

func toString(params map[string]interface{}) string {
	ret := ""
	l := len(params)
	for key, value := range params {
		if l > 1 {
			ret += fmt.Sprintf("%v=%v&", key, value)
		} else {
			ret += fmt.Sprintf("%v=%v", key, value)
		}
		l--
	}
	return ret
}

func ParseStringWithNoPanic(item interface{}) string {
	if item == nil {
		return ""
	}
	if s, ok := item.(string); ok {
		return s
	}
	return ""
}
