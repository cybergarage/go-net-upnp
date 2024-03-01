// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"io"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-net-upnp/net/upnp/control"
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
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

func (dev *Device) isDescriptionURI(path string) bool {
	return path == dev.DescriptionURL
}

func (dev *Device) responseDeviceDescription(httpRes http.ResponseWriter) error {
	devDesc, err := dev.DescriptionString()
	if err != nil {
		return err
	}
	return responseSuccessXMLContent(httpRes, devDesc)
}

func (dev *Device) responseServiceDescription(httpRes http.ResponseWriter, service *Service) error {
	srvDesc, err := service.DescriptionString()
	if err != nil {
		return err
	}
	return responseSuccessXMLContent(httpRes, srvDesc)
}

func (dev *Device) httpGetRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) bool {
	path := httpReq.URL.Path

	// Device Description ?
	if dev.isDescriptionURI(path) {
		err := dev.responseDeviceDescription(httpRes)
		if err != nil {
			responseInternalServerError(httpRes)
		}
		return true
	}

	// Service Description ?
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		if service.isDescriptionURL(path) {
			err := dev.responseServiceDescription(httpRes, service)
			if err != nil {
				responseInternalServerError(httpRes)
			}
			return true
		}
	}

	return false
}

func (dev *Device) httpActionRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter, action *Action) error {
	// has listener ?

	if dev.ActionListener == nil {
		upnpErr := control.NewUPnPErrorFromCode(control.ErrorOptionalActionNotImplemented)
		return responseUPnPError(httpRes, upnpErr)
	}

	// read request

	defer httpReq.Body.Close()
	soapReqBytes, err := io.ReadAll(httpReq.Body)

	if err != nil {
		upnpErr := control.NewUPnPErrorFromCode(control.ErrorInvalidAction)
		return responseUPnPError(httpRes, upnpErr)
	}

	log.Tracef(fmt.Sprintf("action req = \n%s", string(soapReqBytes)))

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

	upnpErr := dev.ActionListener.ActionRequestReceived(action)
	if upnpErr != nil {
		return responseUPnPError(httpRes, upnpErr)
	}

	// return listener response

	actionRes, err := NewActionResponseFromAction(action)
	errStr, _ := actionRes.SOAPContentString()
	return responseSuccessXMLContent(httpRes, errStr)
}

func (dev *Device) httpSoapRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) bool {
	ctrlURL := httpReq.URL.Path
	service, err := dev.GetServiceByControlURL(ctrlURL)
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

	err = dev.httpActionRequestReceived(httpReq, httpRes, action)

	return err == nil
}

func (dev *Device) httpPostRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) bool {
	if httpReq.IsSOAPRequest() {
		return dev.httpSoapRequestReceived(httpReq, httpRes)
	}

	return dev.httpSoapRequestReceived(httpReq, httpRes)
}

func (dev *Device) HTTPRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) {
	log.Infof(fmt.Sprintf("%s %s", httpReq.Method, httpReq.URL.Path))

	switch httpReq.Method {
	case http.GET:
		if dev.httpGetRequestReceived(httpReq, httpRes) {
			return
		}

	case http.POST:
		if dev.httpPostRequestReceived(httpReq, httpRes) {
			return
		}
	}

	if dev.HTTPListener != nil {
		dev.HTTPListener.HTTPRequestReceived(httpReq, httpRes)
		return
	}

	responseBadRequest(httpRes)
}
