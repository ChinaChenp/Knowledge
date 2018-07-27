package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type HttpRequest struct {
	Addr   string
	Client *http.Client
}

func NewHttpClient(addr string, rwTimeout, connTimeout time.Duration) *HttpRequest {
	c := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, connTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(rwTimeout))
				return c, nil
			},
		},
	}

	httpclient := &HttpRequest{
		Addr:   addr,
		Client: c,
	}

	return httpclient
}

func (this *HttpRequest) Get(header, params map[string]string) (response *Response, err error) {
	query := buildQuery(this.Addr, params)
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, err
	}

	this.setDefaultHeader(req.Header)
	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := this.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{HttpResp: resp, HttpCode: resp.StatusCode}, nil
}

func (this *HttpRequest) Post(header, params map[string]string) (response *Response, err error) {
	body, err := mapToJson(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", this.Addr, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	this.setDefaultHeader(req.Header)
	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := this.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{HttpResp: resp, HttpCode: resp.StatusCode}, nil
}

func (this *Response) ReadHeader(keys []string) (ret map[string]string) {
	ret = make(map[string]string)
	for _, key := range keys {
		ret[key] = this.HttpResp.Header.Get(key)
	}
	return
}

func (this *Response) ToString() (string, error) {
	bytes, err := this.readAll()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (this *Response) ToMap() (ret map[string]interface{}, err error) {
	bytes, err := this.readAll()
	if err != nil {
		return nil, err
	}
	ret = make(map[string]interface{})
	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		return nil, err
	}
	return
}

func buildQuery(url string, params map[string]string) string {
	ret := url + "?"
	for key, value := range params {
		ret += key + "=" + value + "&"
	}
	return ret
}

func mapToJson(params map[string]string) (string, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (this *HttpRequest) setDefaultHeader(header http.Header) {
	header.Set("Content-Type", "application/json")
	header.Set("Accept-Charset", "utf-8")
}

type Response struct {
	HttpCode int
	HttpResp *http.Response
}

func (this *Response) readAll() (bytes []byte, err error) {
	var reader io.ReadCloser
	defer func() {
		if reader != nil {
			err = reader.Close()
		}
	}()
	switch this.HttpResp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(this.HttpResp.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = this.HttpResp.Body
	}
	bytes, err = ioutil.ReadAll(reader)
	return
}

func FaceRecognition(params map[string]string) (bool, error) {
	rep, err := FaceService.Post(nil, params)
	if err != nil {
		fmt.Printf("FaceRecognition Post error, %v", err)
		return false, err
	}
	defer rep.HttpResp.Body.Close()

	retstr, err := rep.ToString()
	if err != nil {
		fmt.Printf("reponse to string failed %v", err)
		return false, err
	}

	type Result struct {
		ApiCode int32  `json:"apiCode"`
		ApiMsg  string `json:"apiMsg"`
		Data    struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	}

	var re Result
	if err = json.Unmarshal([]byte(retstr), &re); err != nil {
		fmt.Printf("unmarshal result failed, err:%v", err)
		return false, err
	}

	if re.ApiCode != 200 || re.Data.Code != 100000 {
		fmt.Printf("return code error apicode %v, data code %v\n", re.ApiCode, re.Data.Code)
		return false, errors.New(re.Data.Message)
	}

	return true, nil
}

var FaceService *HttpRequest

func main() {
	connTimeout, err := time.ParseDuration("100ms")
	if err != nil {
		panic(fmt.Sprintf("parse http connect timeout %v", err))
	}

	rwTimeout, err := time.ParseDuration("500ms")
	if err != nil {
		panic(fmt.Sprintf("parse http read & write timeout %v", err))
	}

	FaceService = NewHttpClient("http://10.94.121.45:8000/gateway?api=risk.face.check&apiVersion=1.0.0", rwTimeout, connTimeout)
	b, err := FaceRecognition(map[string]string{
		"idcard":    "xxxxxxx",
		"sessionId": "c8ec54ab-b2bf-4c33-a12a-c764f67a5f92",
	})

	fmt.Printf("result = %v, err %v", b, err)
}
