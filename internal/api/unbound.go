package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type SingleHostOverrideRes struct {
	HostOverride SingleHostOverride `json:"host"`
}

type SingleHostOverride struct {
	Enabled     string  `json:"enabled"`
	HostName    string  `json:"hostname"`
	Domain      string  `json:"domain"`
	Records     Records `json:"rr"`
	MxPriority  string  `json:"mxprio"`
	MxRecord    string  `json:"mx"`
	IpAddress   string  `json:"server"`
	Description string  `json:"Description"`
}

func (s SingleHostOverride) String() string {
	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(s)
	if err != nil {
		fmt.Printf("Error encoding SingleHostOverride to JSON: %s", err.Error())
		return ""
	}
	return b.String()
}

type Records struct {
	A    Record `json:"A"`
	AAAA Record `json:"AAAA"`
	MX   Record `json:"MX"`
}

type Record struct {
	Type   string `json:"value"`
	Active int    `json:"selected"`
}

type HostOverrideListItem struct {
	ID          string `json:"uuid"`
	Enabled     string `json:"enabled"`
	HostName    string `json:"hostname"`
	Domain      string `json:"domain"`
	RecordType  string `json:"rr"`
	MxPriority  string `json:"mxprio"`
	MxHost      string `json:"mx"`
	IpAddress   string `json:"server"`
	Description string `json:"description"`
}

func (h HostOverrideListItem) String() string {
	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(h)
	if err != nil {
		fmt.Printf("Error encoding HostOverride to JSON: %s", err.Error())
		return ""
	}
	return b.String()
}

type HostOverrideList struct {
	Records []HostOverrideListItem `json:"rows"`
}

func (h HostOverrideList) String() string {
	b := &bytes.Buffer{}
	b.WriteString("[\n")
	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "\t")
	for _, v := range h.Records {
		err := encoder.Encode(v)
		if err != nil {
			fmt.Printf("Error encoding HostOverride to JSON: %s", err.Error())
		}
	}
	b.WriteString("]")
	return b.String()
}

type AddHostOverride struct {
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Enabled     string `json:"enabled"`
	HostName    string `json:"hostname"`
	Mx          string `json:"mx"`
	MxPriority  string `json:"mxprio"`
	RecordType  string `json:"rr"`
	IpAddress   string `json:"server"`
}

func (a AddHostOverride) String() string {
	temp := struct {
		Host AddHostOverride `json:"host"`
	}{Host: a}
	s, err := json.Marshal(temp)
	if err != nil {
		fmt.Printf("Error marshalling AddHostOverride to JSON: %s\n", err.Error())
		return ""
	}
	return string(s)
}

type AddItemRes struct {
	Result string `json:"result"`
	UUID   string `json:"uuid"`
}

func (a AddItemRes) String() string {
	s := &bytes.Buffer{}
	encoder := json.NewEncoder(s)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(a)
	if err != nil {
		fmt.Printf("Error when encoding AddOverrideRes to JSON: %s\n", err.Error())
		return ""
	}
	return s.String()
}

type DeleteItemRes struct {
	Result string `json:"result"`
}

func (d DeleteItemRes) String() string {
	s := &bytes.Buffer{}
	encoder := json.NewEncoder(s)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(d)
	if err != nil {
		fmt.Printf("Error when encoding DeleteItemsRes to JSON: %s\n", err.Error())
		return ""
	}
	return s.String()
}

type HostAliasListItem struct {
	UUID        string `json:"uuid"`
	Enabled     string `json:"enabled"`
	Host        string `json:"host"`
	Alias       string `json:"hostname"`
	AliasDomain string `json:"domain"`
	Description string `json:"description"`
}

func (h HostAliasListItem) String() string {
	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(h)
	if err != nil {
		fmt.Printf("Error encoding HostAliasListItem to JSON: %s\n", err.Error())
		return ""
	}
	return b.String()
}

type HostAliasList struct {
	Aliases []HostAliasListItem `json:"rows"`
}

func (h HostAliasList) String() string {
	b := &bytes.Buffer{}
	b.WriteString("[\n")
	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "\t")
	for _, v := range h.Aliases {
		err := encoder.Encode(v)
		if err != nil {
			fmt.Printf("Error encoding HostAliasListItem to JSON: %s\n", err.Error())
			return ""
		}
		b.WriteString(",")
	}
	b.WriteString("]\n")
	return b.String()
}

type SingleHostAliasRes struct {
	Alias SingleHostAlias `json:"alias"`
}

func (s SingleHostAliasRes) String() string {
	return s.Alias.String()
}

type SingleHostAlias struct {
	Enabled       string `json:"enabled"`
	HostOverrides map[string]struct {
		Hostname string `json:"value"`
		Active   int    `json:"selected"`
	} `json:"host"`
	Alias       string `json:"hostname"`
	Domain      string `json:"domain"`
	Description string `json:"description"`
}

func (s SingleHostAlias) String() string {
	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(s)
	if err != nil {
		fmt.Printf("Error encoding SingleHostAlias to JSON: %s\n", err.Error())
		return ""
	}
	return b.String()
}

type AddHostAlias struct {
	Description    string `json:"description"`
	Domain         string `json:"domain"`
	Enabled        string `json:"enabled"`
	HostOverrideID string `json:"host"`
	Alias          string `json:"hostname"`
}

func (a AddHostAlias) String() string {
	temp := struct {
		Alias AddHostAlias `json:"alias"`
	}{Alias: a}
	s, err := json.Marshal(temp)
	if err != nil {
		fmt.Printf("Error marshalling AddHostOverride to JSON: %s\n", err.Error())
		return ""
	}
	return string(s)
}
