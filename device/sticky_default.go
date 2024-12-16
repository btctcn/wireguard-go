//go:build !linux

package device

import (
	"github.com/btctcn/wireguard-go/rwcancel"
	"github.com/btctcn/wireguard-god-go/conn"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
