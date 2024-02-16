package shareddb

import (
	"encoding/json"
	"strconv"
)

type ConfigType string

const (
	ConfigTypeString      ConfigType = "string"
	ConfigTypeInt         ConfigType = "int"
	ConfigTypeFloat       ConfigType = "float"
	ConfigTypeJson        ConfigType = "json"
	ConfigTypeStringArray ConfigType = "[]string"
	ConfigTypeIntArray    ConfigType = "[]int"
	ConfigTypeFloatArray  ConfigType = "[]float"
	ConfigTypeJsonArray   ConfigType = "[]json"
)

type (
	Config struct {
		Base
		RefUUID string     `gorm:"column:ref_uuid" json:"ref_uuid"`
		Key     string     `gorm:"column:key;not null" json:"key"`
		Type    ConfigType `gorm:"column:type;not null" json:"type"`
		Value   string     `gorm:"column:value;not null" json:"value"`
	}
)

func (config Config) String() string {
	return config.Value
}

func (config Config) Integer() int {
	i, _ := strconv.Atoi(config.Value)
	return i
}

func (config Config) Float() float64 {
	f, _ := strconv.ParseFloat(config.Value, 64)
	return f
}

func (config Config) Json() interface{} {
	var result interface{}
	json.Unmarshal([]byte(config.Value), &result)
	return result
}

func (config Config) Strings() []string {
	var strings []string
	json.Unmarshal([]byte(config.Value), &strings)
	return strings
}

func (config Config) Integers() []int {
	var ints []int
	json.Unmarshal([]byte(config.Value), &ints)
	return ints
}

func (config Config) Floats() []float64 {
	var floats []float64
	json.Unmarshal([]byte(config.Value), &floats)
	return floats
}
