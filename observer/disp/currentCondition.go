package disp

import (
	"fmt"

	"github.com/miyamoto-jo/hfdp/observer"
)

// 現在の気温情報
type CurrentConditionDisplay struct {
	weatherData observer.Subject
	temperature float64
	humidity    float64
}

func NewCurrentConditionDisplay(wd observer.Subject) *CurrentConditionDisplay {
	ccd := &CurrentConditionDisplay{
		weatherData: wd,
	}
	return ccd
}

func (m *CurrentConditionDisplay) Update(s observer.Subject) {
	if w, ok := s.(*observer.WeatherData); ok {
		m.temperature = w.GetTemperature()
		m.humidity = w.GetHumidity()
	}
	m.Display()
}

func (m *CurrentConditionDisplay) Display() {
	fmt.Printf("現在の気象状況：温度%f度　湿度%f％です\n", m.temperature, m.humidity)
}
