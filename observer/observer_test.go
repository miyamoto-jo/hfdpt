package observer_test

import (
	"testing"

	"github.com/miyamoto-jo/hfdp/observer"
	"github.com/miyamoto-jo/hfdp/observer/disp"
)

var (
	// Subject
	wd = observer.NewWeatherData()

	// Observers
	ccd = disp.NewCurrentConditionDisplay(wd) // 現在の天気
	sd  = disp.NewStatisticsDisplay(wd)       // 統計情報
	fd  = disp.NewForecastDisplay(wd)         // 気象予報
	hi  = disp.NewHeatIndexDisplay(wd)        // 熱指数

	obList = []observer.Observer{
		ccd, sd, fd, hi,
	}
)

func init() {
	// 予め決めておいたObserverをSubjectに登録
	regist(wd)
}

func TestWeatherStation(t *testing.T) {
	// SubjectにObserver通知イベントをkick
	wd.SetMeasurements(27, 65, 30.4)
	wd.SetMeasurements(28, 70, 29.2)
	wd.SetMeasurements(26, 90, 29.2)

	// 気象予報だけ削除
	wd.RemoveObserver(fd)

	wd.SetMeasurements(15, 30, 20.6)
}

// オブザーバーを登録
func regist(wd observer.Subject) {
	for _, ob := range obList {
		wd.RegisterObserver(ob)
	}
}
