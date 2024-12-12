// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/fluidstackio/fluidstack-go-sdk/internal"
)

type CreateSshKeyRequest struct {
	// The name of the SSH key.
	Name string `json:"name" url:"-"`
	// The public key of the SSH key.
	PublicKey string `json:"public_key" url:"-"`
}

type SshKeysListRequest struct {
	// Show all SSH keys
	ShowAll *bool `json:"-" url:"show_all,omitempty"`
}

type SshKeyResponse struct {
	// The name of the SSH key.
	Name string `json:"name" url:"name"`
	// The public key.
	PublicKey *string `json:"public_key,omitempty" url:"public_key,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SshKeyResponse) GetName() string {
	if s == nil {
		return ""
	}
	return s.Name
}

func (s *SshKeyResponse) GetPublicKey() *string {
	if s == nil {
		return nil
	}
	return s.PublicKey
}

func (s *SshKeyResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SshKeyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SshKeyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SshKeyResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SshKeyResponse) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
