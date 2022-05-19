package config

import (
	"log"
)

type B3NodeConfig struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Category    string                 `json:"category"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Children    []string               `json:"children"`
	Child       string                 `json:"child"`
	Parameters  map[string]interface{} `json:"parameters"`
	Properties  map[string]interface{} `json:"properties"`
}

func (b *B3NodeConfig) getProperty(key string) interface{} {
	if v, ok := b.Properties[key]; ok {
		return v
	}
	log.Fatal("", "GetProperty err ,no vlaue")
	return nil
}

func (b *B3NodeConfig) GetProperty2Int(key string) int {
	if v := b.getProperty(key); v != nil {
		return v.(int)
	}
	return 0
}

func (b *B3NodeConfig) GetProperty2Int64(key string) int64 {
	if v := b.getProperty(key); v != nil {
		return v.(int64)
	}
	return 0
}

func (b *B3NodeConfig) GetProperty2float64(key string) float64 {
	if v := b.getProperty(key); v != nil {
		return v.(float64)
	}
	return 0
}

func (b *B3NodeConfig) GetProperty2String(key string) string {
	if v := b.getProperty(key); v != nil {
		return v.(string)
	}
	return ""
}

func (b *B3NodeConfig) GetProperty2bool(key string) bool {
	if v := b.getProperty(key); v != nil {
		return v.(bool)
	}
	log.Fatal("", "GetProperty err ,format not bool key:%v", key)
	return false
}
