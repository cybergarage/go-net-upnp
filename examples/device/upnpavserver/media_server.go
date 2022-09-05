// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

type MediaServer struct {
	*upnp.Device
}

func NewMediaServer() (*MediaServer, error) {
	dev, err := upnp.NewDeviceFromDescription(mediaServerVerOneDeviceDescription)
	if err != nil {
		return nil, err
	}

	conDirService, err := dev.GetServiceByType("urn:schemas-upnp-org:service:ContentDirectory:1")
	if err != nil {
		return nil, err
	}

	err = conDirService.LoadDescriptionBytes([]byte(contentDirectoryOneServiceDescription))
	if err != nil {
		return nil, err
	}

	conMgrService, err := dev.GetServiceByType("urn:schemas-upnp-org:service:ConnectionManager:1")
	if err != nil {
		return nil, err
	}

	err = conMgrService.LoadDescriptionBytes([]byte(connectionManagerOneServiceDescription))
	if err != nil {
		return nil, err
	}

	mediaServer := &MediaServer{
		Device: dev,
	}

	mediaServer.ActionListener = mediaServer

	return mediaServer, nil
}

func (self *MediaServer) ActionRequestReceived(action *upnp.Action) upnp.Error {
	switch action.Name {
	}

	return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
}
