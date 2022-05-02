package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
)

/**
 * This function is used to create unique IDs for trees and nodes.
 *
 * (consult http://www.ietf.org/rfc/rfc4122.txt).
 *
 * @class createUUID
 * @construCtor
 * @return {String} A unique ID.
**/
func CreateUUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	s := base64.URLEncoding.EncodeToString(b)
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//定义注册结构map
type RegisterStructMaps struct {
	maps map[string]reflect.Type
}

func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{make(map[string]reflect.Type)}
}

//根据name初始化结构
//在这里根据结构的成员注解进行DI注入，这里没有实现，只是简单都初始化
func (rsm *RegisterStructMaps) New(name string) (interface{}, error) {
	fmt.Println("New ", name)
	var c interface{}
	var err error
	if v, ok := rsm.maps[name]; ok {
		c = reflect.New(v).Interface()
		fmt.Println("found ", name, "  ", reflect.TypeOf(c))
		return c, nil
	} else {
		err = fmt.Errorf("not found %s struct", name)
		fmt.Println("New no found", name, "  ", len(rsm.maps))
	}
	return nil, err
}

//查询是否存在
func (rsm *RegisterStructMaps) CheckElem(name string) bool {
	if _, ok := rsm.maps[name]; ok {
		return true
	}
	return false
}

//根据名字注册实例
func (rsm *RegisterStructMaps) Register(name string, c interface{}) {
	rsm.maps[name] = reflect.TypeOf(c).Elem()
}
