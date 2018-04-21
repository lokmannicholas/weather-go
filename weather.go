package main

import (
	"os"

	"github.com/lokmannicholas/weather-go/hko"
	"github.com/urfave/cli"
)

var App = cli.NewApp()

func main() {
	App.Name = "Weather"
	App.Usage = "Weather"
	SetCommand(CreateHKOTables())
	SetCommand(GetHKOWeatherInfo())
	App.Run(os.Args)
}

func SetCommand(cmd cli.Command) {
	App.Commands = append(App.Commands, cmd)
}
func CreateHKOTables() cli.Command {
	return cli.Command{
		Name:  "CreateHKOTables",
		Usage: "Create All Tables",
		Action: func(c *cli.Context) error {
			hko.InitDB()
			return nil
		},
	}
}
func GetHKOWeatherInfo() cli.Command {
	return cli.Command{
		Name:  "GetHKOWeatherInfo",
		Usage: "Fetch HKO Weather Info ",
		Action: func(c *cli.Context) error {
			hko.FetchRealTimeWeather()
			return nil
		},
	}
}
