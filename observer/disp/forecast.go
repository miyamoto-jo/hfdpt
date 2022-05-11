package disp

import (
	"fmt"

	"github.com/miyamoto-jo/hfdp/observer"
)

// 気象予報
type ForecastDisplay struct {
	weatherData observer.Subject
	temperature float64
	humidity    float64
	pressure    float64
}

func NewForecastDisplay(wd observer.Subject) *ForecastDisplay {
	fd := &ForecastDisplay{
		weatherData: wd,
	}
	return fd
}

func (m *ForecastDisplay) Update(temp, humidity, pressure float64) {
	m.temperature = temp
	m.humidity = humidity
	m.pressure = pressure
	m.Display()
}

func (m *ForecastDisplay) Display() {
	fmt.Printf("気象予報です\n")
}
