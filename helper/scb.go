package helper

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
)

func scbGetAccessToken() string {
	url := mod.Conf.ScbUrl + "/v1/oauth/token"

	getTokenBody := payload.ScbGetTokenRequest{
		ApplicationKey:    mod.Conf.ScbAppKey,
		ApplicationSecret: mod.Conf.ScbAppSecret,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(getTokenBody)
	if err != nil {
		logrus.Error("Unable to marshal JSON ", err)
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error("Unable to construct request ", err)
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept-language", "EN")
	req.Header.Add("resourceOwnerId", mod.Conf.ScbAppKey)
	req.Header.Add("requestUId", "1234567890")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error("Unable to send request ", err)
		panic(err)
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		logrus.Error("Unable to read response body ", err)
		panic(err)
	}

	var tokenResponse *payload.ScbGetTokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		logrus.Error("Unable to parse response body ", err)
		panic(err)
	}

	return tokenResponse.Data.AccessToken
}

func ScbCreateQrPayment() {
	accessToken := scbGetAccessToken()

	url := mod.Conf.ScbUrl + "/payment/qrcode/create"

	createQrBody := payload.ScbCreateQrPaymentRequest{
		QrType: "PP",
		Amount: "100.00",
		PpType: "BILLERID",
		PpId:   "1234567890",
		Ref1:   "1234567890",
		Ref2:   "1234567890",
		Ref3:   "1234567890",
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(createQrBody)
	if err != nil {
		logrus.Error("Unable to marshal JSON ", err)
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Error("Unable to construct request ", err)
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept-language", "EN")
	req.Header.Add("resourceOwnerId", mod.Conf.ScbAppKey)
	req.Header.Add("requestUId", "1234567890")
	req.Header.Add("authorization", "Bearer "+accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error("Unable to send request ", err)
		panic(err)
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		logrus.Error("Unable to read response body ", err)
		panic(err)
	}

	var qrResponse *payload.ScbCreateQrResponse
	if err := json.Unmarshal(body, &qrResponse); err != nil {
		logrus.Error("Unable to parse response body ", err)
		panic(err)
	}
}
