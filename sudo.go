package ipa

import (
	"encoding/json"
)

type SudoRuleRecord struct {
	Dn                        string    `json:"dn"`
	Cn                        []string  `json:"cn"`
	UserCactegory             IpaString `json:"usercategory"`
	HostCategory              IpaString `json:"hostcategory"`
	CmdCategory               IpaString `json:"cmdcategory"`
	IpaSudoRunAsUserCategory  IpaString `json:"ipasudorunasusercategory"`
	IpaSudoRunAsGroupCategory IpaString `json:"ipasudorunasgroupcategory"`
	SudoOrder                 int       `json:"sudoorder"`
	MemberUser                []string  `json:"memberuser_user"`
	MemberUserGroup           []string  `json:"memberuser_group"`
	MemberHost                []string  `json:"memberhost_host"`
	MemeberHostGroup          []string  `json:"memberhost_hostgroup"`
	MemberAllCmd              []string  `json:"memberallowcmd_sudocmd"`
	MemberDenyCmd             []string  `json:"memberdenycmd_sudocmd"`
	MemberAllowCmdGroup       []string  `json:"memberallowcmd_sudocmdgroup"`
	MemberDenyCmdGroup        []string  `json:"memberdenycmd_sudocmdgroup"`
	IpaSudoRunAsUser          []string  `json:"ipasudorunas_user"`
	IpaSudoeRunAsGroup        []string  `json:"ipasudorunas_group"`
}

func (c *Client) SudoShow(cn string) (*SudoRuleRecord, error) {
	options := map[string]interface{}{
		"no_members": false,
	}

	return c.RpcSudoRuleRecord("sudorule_show", []string{cn}, options)
}

func (c *Client) SudoAdd(cn string) (*SudoRuleRecord, error) {
	options := map[string]interface{}{
		"cmdcategory": "all",
		"no_members":  false,
	}

	return c.RpcSudoRuleRecord("sudorule_add", []string{cn}, options)
}

func (c *Client) SudoAddHost(cn string, host []string) (*SudoRuleRecord, error) {
	options := map[string]interface{}{
		"host":       host,
		"no_members": false,
	}

	return c.RpcSudoRuleRecord("sudorule_add_host", []string{cn}, options)
}

func (c *Client) SudoAddUser(cn string, user []string) (*SudoRuleRecord, error) {
	options := map[string]interface{}{
		"user":       user,
		"no_members": false,
	}

	return c.RpcSudoRuleRecord("sudorule_add_user", []string{cn}, options)
}

func (c *Client) SudoAddRunAsUser(cn string, sudoUser []string) (*SudoRuleRecord, error) {
	options := map[string]interface{}{
		"user":       sudoUser,
		"no_members": false,
	}

	return c.RpcSudoRuleRecord("sudorule_add_runasuser", []string{cn}, options)
}

func (c *Client) RpcSudoRuleRecord(action string, cn []string, options map[string]interface{}) (*SudoRuleRecord, error) {
	res, err := c.rpc(action, cn, options)
	if err != nil {
		return nil, err
	}

	var sudoRec SudoRuleRecord
	err = json.Unmarshal(res.Result.Data, &sudoRec)
	if err != nil {
		return nil, err
	}

	return &sudoRec, nil

}
