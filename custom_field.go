package moneybird

import (
	"encoding/json"
)

// CustomField contains info about a custom field in Moneybird
type CustomField struct {
	ID               string `json:"id"`
	AdministrationID string `json:"administration_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Source           string `json:"source,omitempty"`
}

// CustomFieldGateway encapsulates all /custom_fields related endpoints
type CustomFieldGateway struct {
	*Client
}

// CustomField returns a CustomFieldGateway instance
func (c *Client) CustomField() *CustomFieldGateway {
	return &CustomFieldGateway{c}
}

// List returns all custom fields stored in Moneybird
func (c *CustomFieldGateway) List() ([]*CustomField, error) {
	var err error
	var customFields []*CustomField

	res, err := c.execute("GET", "custom_fields", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(&customFields)
		return customFields, err
	}
	return nil, res.error()
}
