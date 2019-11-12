package common

import (
	logging "github.com/sirupsen/logrus"
	"github.com/torusresearch/bijson"
)

func Stringify(i interface{}) string {
	bytArr, ok := i.([]byte)
	if ok {
		return string(bytArr)
	}
	str, ok := i.(string)
	if ok {
		return str
	}
	byt, err := bijson.Marshal(i)
	if err != nil {
		logging.WithError(err).Error("Could not bijsonmarshal")
	}
	return string(byt)
}
