package observer

type Subject interface {
	RegisterObserver(o Observer) // オブザーバーの登録
	RemoveObserver(o Observer)   // オブザーバーの削除
	NotifyObservers()            // オブザーバーらに通知
}

type Observer interface {
	Update(s Subject)
}
