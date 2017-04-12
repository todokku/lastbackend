//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package vendor

import (
	"encoding/json"
	"github.com/lastbackend/lastbackend/pkg/apis/types"
)

func New(obj *types.Vendor) *Vendor {
	v := new(Vendor)
	v.Vendor = obj.Vendor
	v.Username = obj.Username
	v.Host = obj.Host
	return v
}

func (obj *Vendor) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func NewList(obj map[string]*types.Vendor) *VendorList {
	v := make(VendorList)
	for index, item := range obj {
		v[index] = New(item)
	}
	return &v
}

func (obj *VendorList) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}
