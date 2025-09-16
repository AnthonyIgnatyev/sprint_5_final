package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataStorage := strings.Split(datastring, ",")
	if len(dataStorage) != 3 {
		return fmt.Errorf("incorrect data count, want: 3, have: %d", len(dataStorage))
	}

	steps, err := strconv.Atoi(dataStorage[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps cannot be negative or zero")
	}

	t.Steps = steps

	t.TrainingType = dataStorage[1]

	duration, err := time.ParseDuration(dataStorage[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("the duration cannot be negative or zero")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("wrong training type")
	}

	actionInfo := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories)

	return actionInfo, nil
}
