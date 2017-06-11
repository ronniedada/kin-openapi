package openapi3

import (
	"github.com/jban332/kinapi/jsoninfo"
)

// Info is specified by OpenAPI/Swagger standard version 3.0.
type Info struct {
	jsoninfo.ExtensionProps
	Title          string   `json:"title,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version,omitempty"`
}

func (value *Info) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStructFields(value)
}

func (value *Info) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStructFields(data, value)
}

// Contact is specified by OpenAPI/Swagger standard version 3.0.
type Contact struct {
	jsoninfo.ExtensionProps
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

func (value *Contact) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStructFields(value)
}

func (value *Contact) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStructFields(data, value)
}

// License is specified by OpenAPI/Swagger standard version 3.0.
type License struct {
	jsoninfo.ExtensionProps
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

func (value *License) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStructFields(value)
}

func (value *License) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStructFields(data, value)
}
