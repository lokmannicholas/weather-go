package hko

import (
	"encoding/json"
	"testing"
)

func TestGetSportCenterWeather(t *testing.T) {

	result, err := FetchRealTimeWeather()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	//check is json format
	b, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(string(b))

}
