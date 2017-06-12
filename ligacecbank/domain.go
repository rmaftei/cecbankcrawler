package ligacecbank

import "time"

type Stage struct {
	stageNumber int
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

type ReverseStages []Stage

func (stages ReverseStages) Len() int {
	return len(stages)
}

func (stages ReverseStages) Swap(i, j int)      {
	stages[i], stages[j] = stages[j], stages[i]
}

func (stages ReverseStages) Less(i int, j int) bool {
	return stages[i].stageNumber > stages[j].stageNumber
}
