package config

import "testing"

func TestLoadRedshiftConfig(T *testing.T) {
	path := "./config.json"

	cfg := FromFile(path)

	if cfg.Encrypted != false {
		T.Log("Expected cfg.Encrypted to be false, got", cfg.Encrypted)
		T.Fail()
	}
}
