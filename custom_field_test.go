package moneybird

import (
	"testing"
)

func TestCustomFieldGatewayList(t *testing.T) {
	_, err := testClient.CustomField().List()

	if err != nil {
		t.Error(err)
	}
}
