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
func ReadConfig() ([]*models.Settings, error) {
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
		main, comment := mainSplitComment(a)
		if main == "" {
			return nil, fmt.Errorf("empty settings for main")
		}
		arr := strings.Split(a, "#")
		if len(arr) == 0 {
			return nil, fmt.Errorf("not valid config.conf. Example one rows: '192.168.41.26 9039 11.0.0.35 3050 #Гурьевская городская больница'")
		}
		portListner, err := strconv.Atoi(arr[1])
		if err != nil {
			return nil, fmt.Errorf("port for listner not integer: %v", err)
		}
		portForward, err := strconv.Atoi(arr[3])
		if err != nil {
			return nil, fmt.Errorf("port for forward not integer: %v", err)
		}
		s := &models.Settings{
			ListnerIP:   arr[0],
			ListnerPort: portListner,
			ForwardIP:   arr[2],
			ForwardPort: portForward,
			Comment:     comment,
		}
		listSettings = append(listSettings, s)
	}
	return listSettings, nil
}

func mainSplitComment(row string) (string, string) {
	arr := strings.Split(row, "#")
	if len(arr) <= 0 {
		return "", ""
	}
	if len(arr) == 1 {
		return arr[0], ""
	}
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	if len(arr) > 2 {
		for i := 1; i < len(arr); i++ {
			if strings.Trim(arr[i], " ") != "" && strings.Trim(arr[i]) != "#" {
				return arr[0], arr[i]
			}
		}
	}
	return "", ""
}
