package weather

import "fmt"

// 現在の気温情報
type CurrentConditionDisplay struct {
	weatherData Subject
	temperature float32
	humidity    float32
}

func NewCurrentConditionDisplay(wd Subject) *CurrentConditionDisplay {
	ccd := &CurrentConditionDisplay{
		weatherData: wd,
	}
	return ccd
}

// サブスク登録
func (m *CurrentConditionDisplay) RegisterSubject() {
	m.weatherData.RegisterObserver(m)
}

// サブスク削除
func (m *CurrentConditionDisplay) RemoveSubject() {
	m.weatherData.RemoveObserver(m)
}

func (m *CurrentConditionDisplay) Update(temp, humidity, _ float32) {
	m.temperature = temp
	m.humidity = humidity
	m.Display()
}

func (m *CurrentConditionDisplay) Display() {
	fmt.Printf("現在の気象状況：温度%f度　湿度%f％です\n", m.temperature, m.humidity)
}

// 統計情報
type StatisticsDisplay struct {
	weatherData Subject
	temperature float32
	humidity    float32
	pressure    float32
}

func NewStatisticsDisplay(wd Subject) *StatisticsDisplay {
	sd := &StatisticsDisplay{
		weatherData: wd,
	}
	return sd
}

// サブスク登録
func (m *StatisticsDisplay) RegisterSubject() {
	m.weatherData.RegisterObserver(m)
}

// サブスク削除
func (m *StatisticsDisplay) RemoveSubject() {
	m.weatherData.RemoveObserver(m)
}

func (m *StatisticsDisplay) Update(temp, humidity, pressure float32) {
	m.temperature = temp
	m.humidity = humidity
	m.pressure = pressure
	m.Display()
}

func (m *StatisticsDisplay) Display() {
	fmt.Printf("統計情報です\n")
}

// 気象予報
type ForecastDisplay struct {
	weatherData Subject
	temperature float32
	humidity    float32
	pressure    float32
}

func NewForecastDisplay(wd Subject) *ForecastDisplay {
	fd := &ForecastDisplay{
		weatherData: wd,
	}
	return fd
}

// サブスク登録
func (m *ForecastDisplay) RegisterSubject() {
	m.weatherData.RegisterObserver(m)
}

// サブスク削除
func (m *ForecastDisplay) RemoveSubject() {
	m.weatherData.RemoveObserver(m)
}

func (m *ForecastDisplay) Update(temp, humidity, pressure float32) {
	m.temperature = temp
	m.humidity = humidity
	m.pressure = pressure
	m.Display()
}

func (m *ForecastDisplay) Display() {
	fmt.Printf("気象予報です\n")
}
