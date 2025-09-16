package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Print("error:", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Print("error:", data, err)
			continue
		}
		log.Println(info)
	}

}
