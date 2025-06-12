/********************************************************************************
 *******************************************************************************/

package mqtt

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"iiot-backend/pkg/go-mod-messaging/internal/pkg"
	"iiot-backend/pkg/go-mod-messaging/pkg/types"
)

// MQTTClientConfig contains all the configurations for the MQTT client.
type MQTTClientConfig struct {
	BrokerURL string
	MQTTClientOptions
}

// ConnectionOptions contains the connection configurations for the MQTT client.
//
// NOTE: The connection properties resides in its own struct in order to avoid the property being loaded in via
//
//	reflection during the load process.
type ConnectionOptions struct {
	BrokerURL string
}

// MQTTClientOptions contains the client options which are loaded via reflection
type MQTTClientOptions struct {
	// Client Identifiers
	Username string
	Password string
	ClientId string
	// Connection information
	Qos            int
	KeepAlive      int // Seconds
	Retained       bool
	AutoReconnect  bool
	CleanSession   bool // MQTT Default is true if never set
	ConnectTimeout int  // Seconds
	pkg.TlsConfigurationOptions
}

// CreateMQTTClientConfiguration constructs a MQTTClientConfig based on the provided MessageBusConfig.
func CreateMQTTClientConfiguration(messageBusConfig types.MessageBusConfig) (MQTTClientConfig, error) {
	var brokerUrl string
	if !messageBusConfig.Broker.IsHostInfoEmpty() {
		brokerUrl = messageBusConfig.Broker.GetHostURL()
	} else {
		return MQTTClientConfig{}, errors.New("broker info not specified")
	}

	_, err := url.Parse(brokerUrl)
	if err != nil {
		return MQTTClientConfig{}, pkg.NewBrokerURLErr(fmt.Sprintf("Failed to parse broker: %v", err))
	}

	mqttClientOptions := CreateMQTTClientOptionsWithDefaults()
	err = pkg.Load(messageBusConfig.Optional, &mqttClientOptions)
	if err != nil {
		return MQTTClientConfig{}, err
	}

	tlsConfig := pkg.TlsConfigurationOptions{}
	err = pkg.Load(messageBusConfig.Optional, &tlsConfig)
	if err != nil {
		return MQTTClientConfig{}, err
	}

	mqttClientOptions.TlsConfigurationOptions = tlsConfig

	return MQTTClientConfig{
		BrokerURL:         brokerUrl,
		MQTTClientOptions: mqttClientOptions,
	}, nil
}

// CreateMQTTClientOptionsWithDefaults constructs MQTTClientOptions instance with defaults.
func CreateMQTTClientOptionsWithDefaults() MQTTClientOptions {
	// Does not need to be cryptographically random client id
	// nolint: gosec
	randomClientId := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000))
	return MQTTClientOptions{
		Username: "",
		Password: "",
		// Client ID is required or else can cause unexpected errors. This was observed with Eclipse's Mosquito MQTT server.
		ClientId:                randomClientId,
		Qos:                     0,
		KeepAlive:               0,
		Retained:                false,
		ConnectTimeout:          5, // 5 seconds
		AutoReconnect:           false,
		CleanSession:            true, // This is the MQTT default
		TlsConfigurationOptions: pkg.CreateDefaultTlsConfigurationOptions(),
	}
}
