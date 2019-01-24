package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsValidateSlackEndpoint(t *testing.T) {
	var err error

	err = ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		EndpointType: endpointTypeSlack,
	})
	assert.NoError(t, err)

	err = ValidateEndpointRequest(Endpoint{
		Title: "title",
	})
	assert.Error(t, err)
}

func TestEndpointsValidateCustomEndpoint(t *testing.T) {
	var err error

	err = ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		Method:       "method",
		Headers:      map[string]string{"key": "value"},
		BodyTemplate: "{\"hello\":\"there\"}",
		EndpointType: endpointTypeCustom,
	})
	assert.NoError(t, err)

	err = ValidateEndpointRequest(Endpoint{
		Title: "title",
	})
	assert.Error(t, err)
}

func TestEndpointsValidatePagerDuty(t *testing.T) {
	var err error

	err = ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ServiceKey:   "serviceKey",
		EndpointType: endpointTypePagerDuty,
	})
	assert.NoError(t, err)

	err = ValidateEndpointRequest(Endpoint{
		Title: "title",
	})
	assert.Error(t, err)
}

func TestEndpointsValidateBigPanda(t *testing.T) {
	var err error

	err = ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ApiToken:     "ApiToken",
		AppKey:       "AppKey",
		EndpointType: endpointTypeBigPanda,
	})
	assert.NoError(t, err)

	err = ValidateEndpointRequest(Endpoint{
		Title: "title",
	})
	assert.Error(t, err)
}

func TestEndpointsValidateDataDog(t *testing.T) {
	var err error

	err = ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ApiKey:       "ApiKey",
		EndpointType: endpointTypeDataDog,
	})
	assert.NoError(t, err)

	err = ValidateEndpointRequest(Endpoint{
		Title: "title",
	})
	assert.Error(t, err)
}

func TestEndpointsValidateVictorOps(t *testing.T) {
	var err error

	err = ValidateEndpointRequest(Endpoint{
		Title:         "title",
		Description:   "description",
		RoutingKey:    "routingKey",
		MessageType:   "messageType",
		ServiceApiKey: "serviceApiKey",
		EndpointType:  endpointTypeVictorOps,
	})
	assert.NoError(t, err)

	err = ValidateEndpointRequest(Endpoint{
		Title: "title",
	})
	assert.Error(t, err)
}
