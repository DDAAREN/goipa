// Use of this source code is governed by a BSD style
// license that can be found in the LICENSE file.

package ipa

import (
	"encoding/json"
	"fmt"
)

type HBACRuleRecord struct {
	Dn                 string    `json:"dn"`
	Cn                 []string  `json:"cn"`
	AccessRuleType     IpaString `json:"accessruletype"`
	UserCategory       IpaString `json:"usercategory"`
	HostCategory       IpaString `json:"hostcategory"`
	SourceHostCategory IpaString `json:"sourcehostcategory"`
	ServiceCategory    []string  `json:"servicecategory"`
	Description        string    `json:"description"`
	MemberUser         []string  `json:"memberuser_user"`
	MemberUserGroup    []string  `json:"memberuser_group"`
	MemberHost         []string  `json:"memberhost_host"`
	MemberHostGroup    []string  `json:"memberhost_hostgroup"`
	SourceHost         []string  `json:"sourcehost_host"`
	SourceHostGroup    []string  `json:"sourcehost_hostgroup"`
	MemberService      []string  `json:"memberservice_hbacsvc"`
	MemberServiceGroup []string  `json:"memberservice_hbacsvcgroup"`
	ExternalHost       IpaString `json:"externalhost"`
}

func (c *Client) HBACAdd(cn string) (*HBACRuleRecord, error) {
	options := map[string]interface{}{
		"accessruletype":  "allow",
		"servicecategory": "all",
		"no_members":      false,
	}

	res, err := c.rpc("hbacrule_add", []string{cn}, options)
	if err != nil {
		return nil, err
	}

	var hbacRec HBACRuleRecord
	err = json.Unmarshal(res.Result.Data, &hbacRec)
	if err != nil {
		return nil, err
	}

	return &hbacRec, nil
}

func (c *Client) HBACShow(cn string) (*HBACRuleRecord, error) {
	options := map[string]interface{}{
		"no_members": false,
	}

	res, err := c.rpc("hbacrule_show", []string{cn}, options)

	if err != nil {
		return nil, err
	}

	fmt.Println(string(res.Result.Data))
	var hbacRec HBACRuleRecord
	err = json.Unmarshal(res.Result.Data, &hbacRec)
	if err != nil {
		return nil, err
	}

	return &hbacRec, nil
}

func (c *Client) HBACAddHost(cn string, host []string) (*HBACRuleRecord, error) {
	options := map[string]interface{}{
		"host":       host,
		"no_members": false,
	}

	res, err := c.rpc("hbacrule_add_host", []string{cn}, options)
	if err != nil {
		return nil, err
	}

	var hbacRec HBACRuleRecord
	err = json.Unmarshal(res.Result.Data, &hbacRec)
	if err != nil {
		return nil, err
	}

	return &hbacRec, nil
}

func (c *Client) HBACAddUser(cn string, user []string) (*HBACRuleRecord, error) {
	options := map[string]interface{}{
		"user":       user,
		"no_members": false,
	}

	res, err := c.rpc("hbacrule_add_user", []string{cn}, options)
	if err != nil {
		return nil, err
	}

	var hbacRec HBACRuleRecord
	err = json.Unmarshal(res.Result.Data, &hbacRec)
	if err != nil {
		return nil, err
	}

	return &hbacRec, nil
}
