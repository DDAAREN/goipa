// Use of this source code is governed by a BSD style
// license that can be found in the LICENSE file.

package ipa

import (
	"encoding/json"
)

type HostRecord struct {
	Dn               string   `json:"dn"`
	Fqdn             []string `json:"fqdn"`
	MemberOfHBACRule []string `json:"memberof_hbacrule"`
}

func (c *Client) HostShow(fqdn string) (*HostRecord, error) {
	options := map[string]interface{}{
		"no_members": false,
	}

	res, err := c.rpc("host_show", []string{fqdn}, options)
	if err != nil {
		return nil, err
	}

	var hostRec HostRecord
	err = json.Unmarshal(res.Result.Data, &hostRec)
	if err != nil {
		return nil, err
	}

	return &hostRec, nil
}


type HostPubKey struct {
    Dn          string      `json:"dn"`
    Fqdn        []string    `json:"fqdn"`
    //SshPubKey   []map[string]string    `json:"ipasshpubkey"`
    SshPubKey   []interface{}    `json:"ipasshpubkey"`
}

func (c *Client) HostPubKey(fqdn string, sshpubkeys []string) (*HostPubKey ,error) {
    options := map[string]interface{}{
        "ipasshpubkey": sshpubkeys }

    res, err := c.rpc("host_mod", []string{fqdn}, options)
    if err != nil {
        return nil, err
    }
        
    var hostPubKey HostPubKey
    err = json.Unmarshal(res.Result.Data, &hostPubKey)
    if err != nil {
        return nil, err
    }

    return &hostPubKey, nil
}
