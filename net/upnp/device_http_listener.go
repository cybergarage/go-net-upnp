// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"io/ioutil"

	"github.com/cybergarage/go-net-upnp/net/upnp/control"
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
	"github.com/cybergarage/go-net-upnp/net/upnp/log"
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

func writeStatusCode(httpRes http.ResponseWriter, code int) error {
	httpRes.WriteHeader(code)
	return nil
}

func writeServerHeader(httpRes http.ResponseWriter) error {
	httpRes.Header().Set(http.ServerHeader, util.GetServer())
	return nil
}

func writeXMLHeader(httpRes http.ResponseWriter) error {
	httpRes.Header().Set(http.ContentType, http.ContentTypeXML)
	return nil
}

func writeContent(httpRes http.ResponseWriter, content []byte) error {
	httpRes.Write(content)
	return nil
}

func responseInternalServerError(httpRes http.ResponseWriter) error {
	writeStatusCode(httpRes, http.StatusInternalServerError)
	writeServerHeader(httpRes)
	return nil
}

func responseBadRequest(httpRes http.ResponseWriter) error {
	writeStatusCode(httpRes, http.StatusBadRequest)
	writeServerHeader(httpRes)
	return nil
}

func responseXMLContent(httpRes http.ResponseWriter, status int, content string) error {
	writeStatusCode(httpRes, status)
	writeServerHeader(httpRes)
	writeXMLHeader(httpRes)
	writeContent(httpRes, []byte(content))
	return nil
}

func responseSuccessXMLContent(httpRes http.ResponseWriter, content string) error {
	return responseXMLContent(httpRes, http.StatusOK, content)
}

func responseUPnPError(httpRes http.ResponseWriter, upnpErr Error) error {
	errRes := NewErrorResponseFromError(upnpErr)
	errStr, _ := errRes.SOAPContentString()
	return responseXMLContent(httpRes, http.StatusInternalServerError, errStr)
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
	return responseSuccessXMLContent(httpRes, devDesc)
}

func (self *Device) responseServiceDescription(httpRes http.ResponseWriter, service *Service) error {
	srvDesc, err := service.DescriptionString()
	if err != nil {
		return err
	}
	return responseSuccessXMLContent(httpRes, srvDesc)
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
	// has listener ?

	if self.ActionListener == nil {
		upnpErr := control.NewUPnPErrorFromCode(control.ErrorOptionalActionNotImplemented)
		return responseUPnPError(httpRes, upnpErr)
	}

	// read request

	defer httpReq.Body.Close()
	soapReqBytes, err := ioutil.ReadAll(httpReq.Body)

	if err != nil {
		upnpErr := control.NewUPnPErrorFromCode(control.ErrorInvalidAction)
		return responseUPnPError(httpRes, upnpErr)
	}

	log.Trace(fmt.Sprintf("action req = \n%s", string(soapReqBytes)))

	// parse request

	actionReq, err := control.NewActionRequestFromSOAPBytes(soapReqBytes)
	if err != nil {
		upnpErr := control.NewUPnPErrorFromCode(control.ErrorInvalidAction)
		return responseUPnPError(httpRes, upnpErr)
	}

	err = action.SetArgumentsByActionRequest(actionReq)
	if err != nil {
		upnpErr := control.NewUPnPErrorFromCode(control.ErrorInvalidArgs)
		return responseUPnPError(httpRes, upnpErr)
	}

	// run listener

	upnpErr := self.ActionListener.ActionRequestReceived(action)
	if upnpErr != nil {
		return responseUPnPError(httpRes, upnpErr)
	}

	// return listener response

	actionRes, err := NewActionResponseFromAction(action)
	errStr, _ := actionRes.SOAPContentString()
	return responseSuccessXMLContent(httpRes, errStr)
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

	return self.httpSoapRequestReceived(httpReq, httpRes)
}

func (self *Device) HTTPRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) {
	log.Info(fmt.Sprintf("%s %s", httpReq.Method, httpReq.URL.Path))

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

	if self.HTTPListener != nil {
		self.HTTPListener.HTTPRequestReceived(httpReq, httpRes)
		return
	}

	responseBadRequest(httpRes)
}
