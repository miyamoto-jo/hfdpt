package weather_test

import (
	"testing"

	"github.com/miyamoto-jo/hfdp/weather"
)

func TestWeatherStation(t *testing.T) {
	wd := weather.NewWeatherData()
	ccd := weather.NewCurrentConditionDisplay(wd) // 現在の天気
	sd := weather.NewStatisticsDisplay(wd)        // 統計情報
	fd := weather.NewForecastDisplay(wd)          // 気象予報

	// 登録
	ccd.RegisterSubject()
	sd.RegisterSubject()
	fd.RegisterSubject()

	// SubjectにObserver通知イベントをkick
	wd.SetMeasurements(27, 65, 30.4)
	wd.SetMeasurements(28, 70, 29.2)
	wd.SetMeasurements(26, 90, 29.2)

	// 気象予報だけ削除
	fd.RemoveSubject()

	wd.SetMeasurements(15, 30, 20.6)
}
