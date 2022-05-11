package disp

import (
	"fmt"

	"github.com/miyamoto-jo/hfdp/observer"
)

// 統計情報
type StatisticsDisplay struct {
	weatherData observer.Subject
	temperature float64
	humidity    float64
	pressure    float64
}

func NewStatisticsDisplay(wd observer.Subject) *StatisticsDisplay {
	sd := &StatisticsDisplay{
		weatherData: wd,
	}
	return sd
}

func (m *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	m.temperature = temp
	m.humidity = humidity
	m.pressure = pressure
	m.Display()
}

func (m *StatisticsDisplay) Display() {
	fmt.Printf("統計情報です\n")
}