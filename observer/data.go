package observer

import "fmt"

type WeatherData struct {
	observers   []Observer
	temperature float64 // 温度
	humidity    float64 // 湿度
	pressure    float64 // 気圧
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

// オブザーバーを登録
func (m *WeatherData) RegisterObserver(o Observer) {
	m.observers = append(m.observers, o)
}

// オブザーバーを削除
func (m *WeatherData) RemoveObserver(o Observer) {
	newObservers := []Observer{}
	for _, ob := range m.observers {
		if ob == o {
			continue
		}
		newObservers = append(newObservers, ob)
	}
	m.observers = newObservers
}

// オブザーバーらに通知
func (m *WeatherData) NotifyObservers() {
	for _, ob := range m.observers {
		ob.Update(m)
	}
}

func (m *WeatherData) MeasurementsChanged() {
	m.NotifyObservers()
}

// 気象観測所側で更新された内容をreadする処理
func (m *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	m.temperature = temperature
	m.humidity = humidity
	m.pressure = pressure
	m.MeasurementsChanged() // Observerに上記変更をkickする
	fmt.Println("-------------------------------------------------------")
}

func (m *WeatherData) GetTemperature() float64 {
	return m.temperature
}
func (m *WeatherData) GetHumidity() float64 {
	return m.humidity
}
func (m *WeatherData) GetPressure() float64 {
	return m.pressure
}
