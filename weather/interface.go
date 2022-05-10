package weather

type Subject interface {
	RegisterObserver(o Observer) // オブザーバーの登録
	RemoveObserver(o Observer)   // オブザーバーの削除
	NotifyObservers()            // オブザーバーらに通知
}

type Observer interface {
	Update(temp, humidity, pressure float32)
	RegisterSubject() // Subjectに登録
	RemoveSubject()   // Subjectから削除
}

type DisplayElement interface {
	Display()
}
