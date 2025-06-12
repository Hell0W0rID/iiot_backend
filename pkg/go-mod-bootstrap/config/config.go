//
// Copyright (C) 2021-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package config

// DatabaseInfo defines database configuration
type DatabaseInfo struct {
        Type     string `toml:"Type" yaml:"Type"`
        Host     string `toml:"Host" yaml:"Host"`
        Port     int    `toml:"Port" yaml:"Port"`
        Timeout  int    `toml:"Timeout" yaml:"Timeout"`
        Name     string `toml:"Name" yaml:"Name"`
        Username string `toml:"Username" yaml:"Username"`
        Password string `toml:"Password" yaml:"Password"`
}