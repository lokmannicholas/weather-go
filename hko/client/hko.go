package client

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"strconv"

	"time"

	"github.com/lokmannicholas/weather-go/hko/models"
)

var Location = []string{"tmt", "tap", "sha", "skg", "jkb", "xtm", "tun", "sty", "cch"}

func GetSportCenterWeather(loc string) (*models.HKOWeather, []error) {
	errs := []error{}
	hkoWeather := &models.HKOWeather{}
	resp, _ := http.Get(fmt.Sprintf("http://www.hko.gov.hk/sports/%s.js", loc))
	defer resp.Body.Close()
	// Load the HTML document
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		data := scanner.Text()
		if strings.Contains(data, "dir") {
			re := regexp.MustCompile(`"(\w+)"`)
			rps := re.FindAllString(scanner.Text(), -1)
			if len(rps) > 0 {
				rp := strings.Replace(rps[0], "\"", "", -1)
				hkoWeather.Direction = rp
			}
		} else if strings.Contains(data, "rh") {
			re := regexp.MustCompile(`"(\d+)"`)
			rps := re.FindAllString(scanner.Text(), -1)
			if len(rps) > 0 {
				rp := strings.Replace(rps[0], "\"", "", -1)
				f64, err := strconv.ParseFloat(rp, 64)
				if err != nil {
					errs = append(errs, err)
					return nil, errs
				}
				hkoWeather.Humidity = f64
			}
		} else if strings.Contains(data, "temp") {
			re := regexp.MustCompile(`"(\d+)"`)
			rps := re.FindAllString(scanner.Text(), -1)
			if len(rps) > 0 {
				rp := strings.Replace(rps[0], "\"", "", -1)
				f64, err := strconv.ParseFloat(rp, 64)
				if err != nil {
					errs = append(errs, err)
					return nil, errs
				}
				hkoWeather.Temperature = f64
			}
		} else if strings.Contains(data, "spd") {
			re := regexp.MustCompile(`"(\d+)"`)
			rps := re.FindAllString(scanner.Text(), -1)
			if len(rps) > 0 {
				rp := strings.Replace(rps[0], "\"", "", -1)
				f64, err := strconv.ParseFloat(rp, 64)
				if err != nil {
					errs = append(errs, err)
					return nil, errs
				}
				hkoWeather.Speed = f64
			}
		} else if strings.Contains(data, "gust") {
			re := regexp.MustCompile(`"(\d+)"`)
			rps := re.FindAllString(scanner.Text(), -1)
			if len(rps) > 0 {
				rp := strings.Replace(rps[0], "\"", "", -1)
				f64, err := strconv.ParseFloat(rp, 64)
				if err != nil {
					errs = append(errs, err)
					return nil, errs
				}
				hkoWeather.Gust = f64
			}
		} else if strings.Contains(data, "time") {
			re := regexp.MustCompile(`(\d+)`)
			rps := re.FindAllString(scanner.Text(), -1)
			if len(rps) == 2 {
				hour, err := strconv.Atoi(rps[0])
				if err != nil {
					errs = append(errs, err)
					return nil, errs
				}
				min, err := strconv.Atoi(rps[1])
				if err != nil {
					errs = append(errs, err)
					return nil, errs
				}
				now := time.Now()
				if hour == 23 && now.Hour() == 0 {
					now = now.AddDate(0, 0, -1)
				}
				dataTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.Local)
				hkoWeather.Time = dataTime
			}
		}
		hkoWeather.Location = loc
	}

	return hkoWeather, errs
}
