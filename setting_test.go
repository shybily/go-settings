package settings

import (
	"reflect"
	"testing"
)

func TestSetting(t *testing.T) {
	if !Config.HasKey("host") {
		t.Logf("parse ini faile failed")
		t.Fail()
	}
}

func TestVal(t *testing.T) {
	host := Val("db.mongo.default")
	if len(host) == 0 || host != "mongodb://127.0.0.1:27017" {
		t.Logf("get host val[%s]", host)
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	port := Int("port")
	if port == 0 {
		t.Logf("pars port failed")
		t.Fail()
	}

	if v, ok := interface{}(port).(int); !ok {
		t.Logf("value [%d] type [%v] not type of int", v, reflect.TypeOf(port))
		t.Failed()
	}
}

func TestInt64(t *testing.T) {
	uniq := Int64("uniq")
	if uniq == 0 {
		t.Logf("pars port failed")
		t.Fail()
	}

	if v, ok := interface{}(uniq).(int64); !ok {
		t.Logf("value [%d] type [%v] not type of int64", v, reflect.TypeOf(uniq))
		t.Failed()
	}
}
