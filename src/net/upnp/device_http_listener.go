// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"io/ioutil"

	"net/upnp/control"
	"net/upnp/http"
)

func writeStatusCode(httpRes http.ResponseWriter, code int) error {
	httpRes.WriteHeader(code)
	return nil
}

func responseInternalServerError(httpRes http.ResponseWriter) error {
	return writeStatusCode(httpRes, http.StatusInternalServerError)
}

func responseBadRequest(httpRes http.ResponseWriter) error {
	return writeStatusCode(httpRes, http.StatusBadRequest)
}

func writeXMLHeader(httpRes http.ResponseWriter) error {
	httpRes.Header().Set(http.ContentType, http.ContentTypeXML)
	return nil
}

func writeContent(httpRes http.ResponseWriter, content []byte) error {
	httpRes.Write(content)
	return nil
}

func responseXMLContent(httpRes http.ResponseWriter, content string) error {
	writeStatusCode(httpRes, http.StatusOK)
	writeXMLHeader(httpRes)
	writeContent(httpRes, []byte(content))

	return nil
}

func (self *Device) isDescriptionUri(path string) bool {
	if path == self.DescriptionURL {
		return true
	}
	return false
}

func (self *Device) responseDeviceDescription(httpRes http.ResponseWriter) error {
	devDesc, err := self.DescriptionString()
	if err != nil {
		return err
	}
	return responseXMLContent(httpRes, devDesc)
}

func (self *Device) responseServiceDescription(httpRes http.ResponseWriter, service *Service) error {
	srvDesc, err := service.DescriptionString()
	if err != nil {
		return err
	}
	return responseXMLContent(httpRes, srvDesc)
}

func (self *Device) httpGetRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) bool {
	path := httpReq.URL.Path

	// Device Description ?
	if self.isDescriptionUri(path) {
		err := self.responseDeviceDescription(httpRes)
		if err != nil {
			responseInternalServerError(httpRes)
		}
		return true
	}

	// Service Description ?
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.isDescriptionURL(path) {
			err := self.responseServiceDescription(httpRes, service)
			if err != nil {
				responseInternalServerError(httpRes)
			}
			return true
		}
	}

	return false
}

func (self *Device) httpActionRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter, action *Action) error {
	defer httpReq.Body.Close()
	soapReqBytes, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		return err
	}

	actionReq, err := control.NewActionRequestFromSOAPBytes(soapReqBytes)
	if err != nil {
		return err
	}

	err = action.SetArgumentsByActionRequest(actionReq)
	if err != nil {
		return err
	}

	return nil
}

func (self *Device) httpSoapRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) bool {
	ctrlURL := httpReq.URL.Path
	service, err := self.GetServiceByControlURL(ctrlURL)
	if err != nil {
		return false
	}

	actionName, ok := httpReq.GetSOAPServiceActionName()
	if !ok {
		return false
	}

	action, err := service.GetActionByName(actionName)
	if err != nil {
		return false
	}

	err = self.httpActionRequestReceived(httpReq, httpRes, action)
	if err != nil {
		return false
	}

	return true
}

func (self *Device) httpPostRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) bool {
	if httpReq.IsSOAPRequest() {
		return self.httpSoapRequestReceived(httpReq, httpRes)
	}

	return false
}

func (self *Device) HTTPRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) {
	switch httpReq.Method {
	case http.GET:
		if self.httpGetRequestReceived(httpReq, httpRes) {
			return
		}

	case http.POST:
		if self.httpPostRequestReceived(httpReq, httpRes) {
			return
		}
	}

	if self.Listener != nil {
		self.Listener.HTTPRequestReceived(httpReq, httpRes)
		return
	}

	responseBadRequest(httpRes)
}
