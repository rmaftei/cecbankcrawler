package ligacecbank

import "time"

type Stage struct {
	Games []Game
}

type Game struct {
	StartTime time.Time
	Location string
	Team1 string
	Team2 string
	PointsTeam1 int
	PointsTeam2 int
	LiveTransmission string

}

type Fixtures interface {
	GetFixtures() []Stage
}

