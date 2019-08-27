package playerstatus

//status 0：未在房间中
//1：未就绪
//2：已准备
//3：选好英雄
//4：游戏进行中
const (
	NotInRoom byte = iota
	NotReady
	IsReady
	Locked
	Loaded
	Gaming
	Deleted
)
