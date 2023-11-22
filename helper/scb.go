package helper

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"strconv"
	"time"
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
	req.Header.Add("requestUId", "PICME_PAYMENT_SYSTEM")

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

func ScbCreateQrPayment(amount uint, transactionId string) payload.ScbCreateQrDataResponse {
	accessToken := scbGetAccessToken()

	url := mod.Conf.ScbUrl + "/v1/payment/qrcode/create"
	createQrBody := payload.ScbCreateQrPaymentRequest{
		QrType: "PP",
		Amount: strconv.Itoa(int(amount)) + ".00",
		PpType: "BILLERID",
		PpId:   mod.Conf.ScbBillerId,
		Ref1:   transactionId,
		Ref2:   "PICME01",
		Ref3:   "PICME01",
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

	return qrResponse.Data
}

func ScbInquiryPayment(transactionId string) (*payload.ScbInquiryQrDataResponse, error) {
	accessToken := scbGetAccessToken()

	// get current date in format 2019-10-28
	currentDate := time.Now().Format("2006-01-02")
	url := mod.Conf.ScbUrl + "/v1/payment/billpayment/inquiry?eventCode=00300100&transactionDate=" + currentDate + "&billerId=" + mod.Conf.ScbBillerId + "&reference1=" + transactionId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error("Unable to construct request ", err)
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept-language", "EN")
	req.Header.Add("resourceOwnerId", mod.Conf.ScbAppKey)
	req.Header.Add("requestUId", "1234567890")
	req.Header.Add("authorization", "Bearer "+accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error("Unable to send request ", err)
		return nil, err
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		logrus.Error("Unable to read response body ", err)
		return nil, err
	}

	var qrResponse *payload.ScbInquiryQrResponse
	if err := json.Unmarshal(body, &qrResponse); err != nil {
		logrus.Error("Unable to parse response body ", err)
		return nil, err
	}

	if qrResponse.Data == nil || len(qrResponse.Data) == 0 {
		return nil, nil
	}

	return &qrResponse.Data[0], nil
}
