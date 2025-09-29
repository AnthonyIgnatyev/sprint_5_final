package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataStorage := strings.Split(datastring, ",")
	if len(dataStorage) != 2 {
		return fmt.Errorf("incorrect data count")
	}

	steps, err := strconv.Atoi(dataStorage[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps cannot be negative or zero")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(dataStorage[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("the duration cannot be negative or zero")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err

	}

	actionInfo := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories)

	return actionInfo, nil

}
