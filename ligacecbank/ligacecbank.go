package ligacecbank

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
	"strconv"
)

const RESOURCE_URL = "https://www.super-liga.ro/superliga/program-rezultate/"
const SELECTOR = "[style^=\"padding:2px;border-bottom:1px solid #ccc;vertical-align:middle;\"]"
const STAGE_SIZE = 3

type Stage struct {
	games []Game
}

type Game struct {
	startTime time.Time
	location string
	team1 string
	team2 string
	pointsTeam1 int
	pointsTeam2 int
	liveTransmission string

}

func GetFixtures() []Stage {
	var stages = make([]Stage, 0)

	result := getDataFromWeb(RESOURCE_URL, SELECTOR)

	currentStage := make([]Game, 0)

	for i := 0; i < len(result); i+=6 {
		data := result[i: i + 6]

		currentStage = append(currentStage, dataToGame(data))

		if len(currentStage) % STAGE_SIZE == 0 {
			stages = append(stages, Stage{currentStage})
		}

	}

	return stages
}

func getDataFromWeb(url string, selector string ) []string {

	var result = make([]string, 0)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		text := strings.Replace(s.Text(), "\n", "", -1)
		text = strings.Replace(text, "\t", "", -1)

		if len(text) > 0 {
			result = append(result, text)
		}

	})

	return result
}

func dataToGame(data []string) Game {
	dateTime := strings.Split(data[0], "ora")

	date := strings.Split(dateTime[0], ".")
	clock := strings.Split(strings.TrimSpace(dateTime[1]), ":")

	year, _:= strconv.Atoi(date[2])
	month, _:= strconv.Atoi(date[1])
	day, _:= strconv.Atoi(date[0])
	hour, _:= strconv.Atoi(clock[0])
	minute, _:= strconv.Atoi(clock[1])

	dateObject := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	teams := strings.Split(data[3], " - ")
	score := strings.Split(data[4], "-")
	scoreTeam1, _ := strconv.Atoi(score[0])
	scoreTeam2, _ := strconv.Atoi(score[1])

	return Game{
		dateObject,
		data[2],
		teams[0],
		teams[1],
		scoreTeam1,
		scoreTeam2,
		data[5]}

}
