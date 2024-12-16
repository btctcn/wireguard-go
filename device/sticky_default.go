//go:build !linux

package device

import (
	"github.com/btctcn/wireguard-go/conn"
	"github.com/btctcn/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
