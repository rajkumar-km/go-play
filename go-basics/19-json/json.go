/*
json demonstrates encoding and decoding JSON with Go

JSON stands for Javascript Object Notation.
JSON is an encoding method like XML, YAML, ASN.1, and Google Protocol Buffers.
- JSON follows the Javascript values for encoding
- Go has excellent support for encodings with encoding/json, encoding/xml, and encoding/asn1, and so on.
JSON Types:
  - boolean true
  - number -273.15
  - string "She said \"Hello, BF\""
  - array ["gold", "silver", "bronze"]
  - object {
    "year": 1980,
    "event": "archery",
    "medals": ["gold", "silver", "bronze"]
    }

The package encoding/json includes:
  - json.Marshal() - encode Go data as json
  - json.MarshalIndent() - encode as json with indentation
  - json.Unmarshal() - decode json to Go data structure
  - json.NewDecoder().Decode() - streaming decoder
  - json.NewEncoder().Encode() - streaming encoder
*/
package main

import (
	"encoding/json"
	"fmt"
)

// Proxy is a sample GO struct for encoding/decoding
// Only exported fields (starting with capital letter) are considered for encoding.
//
// A field can have multiple tags separated with space. Each tags consits of key-value pairs
// without space. Example: `bson:"outbound_proxy" json:"outbound_proxy,omitempty"`
//
// By default, the encoding/decoding works by the field names.
// These tags are helpful to use a different name (or behavior like omitempty)
// for encoding.
//
// The omitempty tag ignores the zero value for encoding/decoding for the particular field.
// Note that this includes a value of "0" for int and "false" for bool and etc.,
// So, one must be careful while using omitempty. If int 0 is valid value for an application,
// then avoid using omitempty.
//
// Alternatively a field can be declared as pointer and configured omitempty. This can
// still store zero values except the it omits the nil value.
type Proxy struct {
	Ip       string `json:"ip,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username *string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	metadata string
}

func main() {
	// Sample data structure for JSON encoding/decoding
	p := &Proxy{
			Ip: "192.168.1.1",
			Port: 1234,
			Username: new(string), // Not part of the json if it is nil. Empty string is still part of json since it is a pointer field
			Password: "", // Not part of the json if it is empty string
			metadata: "secret", // This field not exported and not part of encoded json
		}

	// Encode JSON
	clusterJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("json.Marshal(%#v):\n\t%s\n", p, clusterJson)
	}

	// Encode JSON indented
	formattedClusterJson, err := json.MarshalIndent(p, "\t", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("json.MarshalIndent\n\t%s\n", formattedClusterJson)
	}

	// Decode JSON
	proxyJsonStr := `{"IP":"192.168.1.2", "port":5392, "username":"admin", "password":"admin", "metadata":"secret", "unknown":""}`
	var p2 Proxy
	err = json.Unmarshal([]byte(proxyJsonStr), &p2)
	fmt.Printf("json.Unmarshal(%s)\n\t%#v\n", proxyJsonStr, p2)
	// &main.Proxy{Ip:"192.168.1.2", Port:5392, Username:"admin", Password:"admin", metadata:""}
	// - Note the JSON string contains "IP" (all uppercase), but our json tag is "ip". 
	//   But, unmarshal worked for the field. So, the unmashal is case insensitive.
	// - Next, metadata is not populated, because it is an unexported Go field
	// - Finally, all the other fields presents in JSON string are ignored (see "unknown" field)
}