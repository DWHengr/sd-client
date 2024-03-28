package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sd-client/httpclient/header"
	"sd-client/jwt"
	"sd-client/logger"
)

// MarshalNotHtml not encode html
func MarshalNotHtml(data interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return nil, err
	}
	return bf.Bytes(), nil
}

// POST http post
func POST(ctx context.Context, uri string, params interface{}, entity interface{}) error {
	client = GetHttpClient()
	paramByte, err := MarshalNotHtml(params)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(paramByte)
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	if ctx != nil {
		fmt.Println("Request-Id:" + header.GetRequestId(ctx) + " http request uri:" + uri + "http request params:" + string(paramByte))
		req.Header.Add(header.GetRequestIDKV(ctx).Wreck())
	}
	logger.Logger.Info("http post request uri:" + uri + " http request params:" + string(paramByte))
	//添加token
	token, _ := jwt.GenerateJWTToken("sd-client")
	req.Header.Set("token", token)
	response, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	defer response.Body.Close()
	err = DecomposeResp(response, entity)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	return err
}

func GET(uri string, entity interface{}) error {
	logger.Logger.Info("http get request uri:" + uri)
	client = GetHttpClient()
	req, err := http.NewRequest("GET", uri, nil)
	//添加token
	token, _ := jwt.GenerateJWTToken("sd-client")
	req.Header.Set("token", token)
	response, err := client.Do(req)
	response.Header.Set("", "")
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	if entity != nil {
		err = json.Unmarshal(body, entity)
	}
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	return nil
}

// Request http
func Request(uri string, method string, params interface{}, entity interface{}, headers map[string]string) error {
	client = GetHttpClient()
	paramByte, err := json.Marshal(params)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(paramByte)
	req, err := http.NewRequest(method, uri, reader)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	if entity != nil {
		err = json.Unmarshal(body, entity)
	}
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	return nil
}
