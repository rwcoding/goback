package dblog

import (
	"encoding/json"
	"github.com/rwcoding/goback/models"
)

var db = models.GetDb

func Add(log *models.Log) {
	db().Model(log).Create(log)
}

func NewONData() *oldNewData {
	return &oldNewData{data: map[string]interface{}{}}
}

type oldNewData struct {
	data map[string]interface{}
}

func (s *oldNewData) Old(k string, v string) *oldNewData {
	s.data["old_"+k] = v
	return s
}

func (s *oldNewData) New(k string, v string) *oldNewData {
	s.data["new_"+k] = v
	return s
}

func (s *oldNewData) OldStruct(o interface{}) *oldNewData {
	s.data["old"] = o
	return s
}

func (s *oldNewData) NewStruct(o interface{}) *oldNewData {
	s.data["new"] = o
	return s
}

func (s *oldNewData) Json() string {
	if len(s.data) == 0 {
		return ""
	}
	b, _ := json.Marshal(s.data)
	return string(b)
}
