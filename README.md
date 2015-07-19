# go-net-upnp

go-net-upnp is a open source framework for Go and UPnP™ developers.

UPnP™ is a standard protocol for IoT, the protocols consist of other standard protocols, such as GENA, SSDP, SOAP, HTTPU and HTTP. Therefore you have to understand and implement these protocols to create UPnP™ applications.

go-net-upnp hansles these protocols automatically to support to create UPnP devices and control points quickly.

## Installation

To use go-net-upnp, run `go get` as the following:

```
go get -u github.com/cybergarage/go-net-upnp/net/upnp
```

## Examples

```
example/
├── ctrlpoint
│   ├── upnpctrl
│   │   ├── control_point.go
│   │   ├── control_point_action.go
│   │   ├── control_point_action_mgr.go
│   │   ├── control_point_actions.go
│   │   └── upnpctrl.go
│   ├── upnpdump
│   │   ├── control_point.go
│   │   └── upnpdump.go
│   ├── upnpgwlist
│   │   ├── upnpgwdev.go
│   │   └── upnpgwlist.go
│   └── upnpsearch
│       └── upnpsearch.go
└── dev
    └── lightdev
        ├── lightdev.go
        ├── lightdev_desc.go
        └── main.go
```
## Repositories

The project is hosted on the following sites. Please check the following sites to know about go-net-upnp in more detail.

- [GitHub](https://github.com/cybergarage/go-net-upnp)
