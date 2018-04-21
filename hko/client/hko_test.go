package client

import (
	"encoding/json"
	"testing"
)

func TestGetSportCenterWeather(t *testing.T) {

	result, errs := GetSportCenterWeather("tmt")
	if len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
		t.Fail()
	} else {
		//check is json format
		b, err := json.Marshal(result)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		t.Log(string(b))
	}

}
