package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/cartmanis/call_forwarding/app/models"
)

//ReadConfig чтение конфигурационного файла ini, если файла не существует он создается со значениями по умолчанию и сообщением выхода из программы
func ReadConfig() ([]*models.Settings,error) {
	data, err := ioutil.ReadFile("config.conf")
	if err != nil {
		return nil, err
	}
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("not exist configuration file config.conf: %v", err)
	}
	arrFile := strings.Split(string(data), "\n")
	if len(arrFile) == 0 {
		return nil, fmt.Errorf("empty config.conf")
	}
	listSettings := make([]*models.Settings, 0)

	for _, a := range arrFile {
		if strings.Trim(a, " ") == "" {
			continue
		}
		arr := strings.Split(a, " ")
		if len(arr) < 4 {
			return nil, fmt.Errorf("not valid config.conf. Example one rows: '192.168.41.26 9039 11.0.0.35 3050 #Гурьевская городская больница'")
		}
		portListner, err := strconv.Atoi(arr[1])
		if err != nil {
			return nil, fmt.Errorf("port for listner not integer: %v",err)
		}
		portForward, err := strconv.Atoi(arr[3])
		if err != nil {
			return nil, fmt.Errorf("port for forward not integer: %v",err)
		}
		s := &models.Settings{
			ListnerIP:   arr[0],
			ListnerPort: portListner,
			ForwardIP:   arr[2],
			ForwardPort: portForward,
			Comment:     getComment(arr),
		}
		listSettings = append(listSettings, s)
	}
	return listSettings, nil
}

func getComment(arr []string) string {
	comment := ""
	for i := 4; i < len(arr); i++ {
		row := strings.Trim(arr[i], " ")
		if len(row) > 1 && string(row[0]) == "#" {
			comment = row[1:]
		}
	}
	return comment
}
