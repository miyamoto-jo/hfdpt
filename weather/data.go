package weather

import "fmt"

type WeatherData struct {
	observers   []Observer
	temperature float32
	humidity    float32
	pressure    float32
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
		ob.Update(m.temperature, m.humidity, m.pressure)
	}
}

func (m *WeatherData) MeasurementsChanged() {
	m.NotifyObservers()
}

// 気象観測所側で更新された内容をreadする処理
func (m *WeatherData) SetMeasurements(temperature, humidity, pressure float32) {
	m.temperature = temperature
	m.humidity = humidity
	m.pressure = pressure
	m.MeasurementsChanged() // Observerに上記変更をkickする
	fmt.Println("-------------------------------------------------------")
}
