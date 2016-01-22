package config

import (
	"encoding/json"
	"io/ioutil"
)

type RedShiftConfig struct {
	ClusterIdentifier         string  `json:"clusterIdentifier"`
	MasterUserPassword        string  `json:"masterUserPassword"`
	MasterUsername            string  `json:"masterUsername"`
	NodeType                  string  `json:"nodeType"`
	AvailabilityZone          string  `json:"availabilityZone"`
	ClusterParameterGroupName string  `json:"clusterParameterGroupName"`
	ClusterSubnetGroupName    string  `json:"clusterSubnetGroupName"`
	DbName                    string  `json:"dbName"`
	Encrypted                 bool    `json:"encrypted"`
	NumberOfNodes             float64 `json:"numberOfNodes"`
	Port                      float64 `json:"port"`
	PubliclyAccessible        bool    `json:"publiclyAccessible"`
	VpcSecurityGroupIDs       string  `json:"vpcSecurityGroupIDs"`
}

func FromFile(path string) RedShiftConfig {

	var settings RedShiftConfig
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic("opening config file")
	}

	if err = json.Unmarshal(configFile, &settings); err != nil {
		panic("parsing config file")
	}

	return settings
}
