package disp

import (
	"fmt"

	"github.com/miyamoto-jo/hfdp/observer"
)

// 熱指数
type HeatIndexDisplay struct {
	weatherData observer.Subject
	heatindex   float64
}

func NewHeatIndexDisplay(wd observer.Subject) *HeatIndexDisplay {
	hi := &HeatIndexDisplay{
		weatherData: wd,
	}
	return hi
}

func (m *HeatIndexDisplay) Update(temp, humidity, pressure float64) {
	temp = 1.8*temp + 32 // 摂氏から華氏に変換
	m.heatindex = m.computeHeatIndex(temp, humidity)
	m.Display()
}

func (m *HeatIndexDisplay) Display() {
	fmt.Printf("熱指数 %f\n", m.heatindex)
}

// 熱指数を計算する
func (m *HeatIndexDisplay) computeHeatIndex(t, rh float64) float64 {
	return (float64)(
		(16.923 + (0.185212 * t)) +
			(5.37941 * rh) -
			(0.100254 * t * rh) +
			(0.00941695 * (t * t)) +
			(0.00728898 * (rh * rh)) +
			(0.000345372 * (t * t * rh)) -
			(0.000814971 * (t * rh * rh)) +
			(0.0000102102 * (t * t * rh * rh)) -
			(0.000038646 * (t * t * t)) +
			(0.0000291583 * (rh * rh * rh)) +
			(0.00000142721 * (t * t * t * rh)) +
			(0.000000197483 * (t * rh * rh * rh)) -
			(0.0000000218429 * (t * t * t * rh * rh)) +
			(0.000000000843296 * (t * t * rh * rh * rh)) -
			(0.0000000000481975 * (t * t * t * rh * rh * rh)))
}
