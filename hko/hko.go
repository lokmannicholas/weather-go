package hko

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/lokmannicholas/weather-go/database/mysql"
	"github.com/lokmannicholas/weather-go/hko/client"
	"github.com/lokmannicholas/weather-go/hko/models"
)

func InitDB() {
	CreateHKOWeatherTables()
}
func CreateHKOWeatherTables() {

	sql :=
		`CREATE TABLE IF NOT EXISTS hko_weather (
		  location varchar(32) NOT NULL,
		  temp float DEFAULT NULL,
		  humidity float DEFAULT NULL,
		  speed float DEFAULT NULL,
		  gust float DEFAULT NULL,
		  direction varchar(32) DEFAULT NULL,
		  time timestamp NOT NULL,
		  PRIMARY KEY (time,location)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	r, err := mysql.Get().DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
}

func FetchRealTimeWeather() ([]models.HKOWeather, error) {
	hkos := []models.HKOWeather{}
	db := mysql.Get().DB
	sql := `Insert into %s (temp,humidity,speed,gust,direction,time,location
			) VALUES (%f,%f,%f,%f,"%s","%s","%s")`

	sql2 := ` ON DUPLICATE KEY UPDATE
				 	temp = %f,
				 	humidity = %f,
				 	speed = %f,
				 	gust = %f,
					direction = "%s",
					time = "%s",
					location = "%s"
			;`
	table := "hko_weather"
	for _, loc := range client.Location {

		hkoWeather, errs := client.GetSportCenterWeather(loc)
		if len(errs) > 0 {
			for _, err := range errs {
				log.Println(err.Error())
			}
			return nil, errors.New("multiple errors")
		}
		s := fmt.Sprintf(sql, table, hkoWeather.Temperature, hkoWeather.Humidity, hkoWeather.Speed, hkoWeather.Gust, hkoWeather.Direction, hkoWeather.Time, hkoWeather.Location)
		s2 := fmt.Sprintf(sql2, hkoWeather.Temperature, hkoWeather.Humidity, hkoWeather.Speed, hkoWeather.Gust, hkoWeather.Direction, hkoWeather.Time, hkoWeather.Location)
		stm, err := db.Prepare(s + s2)
		if err != nil {
			log.Println(err.Error())
		} else {
			_, err = stm.ExecContext(context.Background())
			if err != nil {
				log.Println(err.Error())
			}
		}
		hkos = append(hkos, *hkoWeather)

	}
	return hkos, nil
}
