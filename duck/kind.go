package duck

var (
	MallardDuck = New(`マガモ`, &Quack{}, &FlyWithWings{})
	RedHeadDuck = New(`アメリカホシハジロ`, &Quack{}, &FlyWithWings{})
	RubberDuck  = New(`ゴム製のカモ`, &Squeak{}, &FlyNoWay{})
	DecoyDuck   = New(`おとりのカモ`, &MuteQuack{}, &FlyNoWay{})
)
