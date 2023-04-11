package flyweight

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleton().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *Game) AddTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) AddCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	g.terrorists = append(g.terrorists, player)
}
