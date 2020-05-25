package library

import (
	"encoding/json"
	"fmt"
	"github.com/guoruibiao/gorequests"
	"github.com/guoruibiao/goworktools/models"
	"sort"
	"strconv"
)

func GetSortedMapValues(m map[string]float64) []float64 {
	values := make([]float64, 0)
	for _, value := range m {
		values = append(values, value)
	}
	sort.Float64s(values)
	return values
}


func GetInstanceByBNS(bns string)(bnsList []models.BNSItem, err error) {

	url := "http://bns.noah.xxx.com/webfoot/index.php?r=webfoot/ApiInstanceInfo&serviceName=" + bns
	resp, err := gorequests.NewRequest("GET", url).DoRequest()
	if err != nil {
		return nil, err
	}
	content, err := resp.Content()
	if err != nil {
		return nil, err
	}
	fmt.Println(content)

	bytes := []byte(content)
	var instanceContainer models.InstanceContainer
	json.Unmarshal(bytes, &instanceContainer)
	for _, instance := range instanceContainer.Data {
		port, err := strconv.Atoi(instance.Port)
		if err != nil {
			continue
		}
		bnsList = append(bnsList, models.BNSItem{
			Host: instance.HostName,
			Port: port,
		})
	}
	fmt.Println("-------*******************************-------")
	fmt.Println(bnsList)
	bnsList = append(bnsList, models.BNSItem{
		Host: "localhost",
		Port: 6379,
	})

	/*commander := commands.New()

	params := []string{" -p ", bns}
	succ, output := commander.GetStatusOutput("get_instance_by_service", params...)
	if !succ {
		return nil, fmt.Errorf("no response for `get_instance_by_service` %s", output)
	}


	// 解析 output，封装成 BNSList
	fmt.Println(commander, output)
	template := `bjyz-oxp-osp-14091.bjyz 7060
	bjyz-oxp-osp-14091.bjyz 7064
	bjyz-feed-backup-pool-13330.bjyz 7142
	bjyz-tm-psin-p011.bjyz 2003
	bjyz-tm-psin-p011.bjyz 2007
	bjhw-bdrp-share346.bjhw 2007
	bjhw-bdrp-share346.bjhw 2003
	bjhw-bdrp-share323.bjhw 7078
	bjhw-bdrp-share323.bjhw 7082
	bjhw-bdrp-share323.bjhw 7086
	bjhw-bdrp-share383.bjhw 7072`
	instances := strings.Split(template, "\n")
	for _, item := range instances {
		row := strings.Split(item, " ")
		host := strings.Trim(row[0], " \t")
		port, _ := strconv.Atoi(row[1])
		bnsList = append(bnsList, models.BNSItem{
			Host: host,
			Port: port,
		})

	}
	bnsList = append(bnsList, models.BNSItem{
		Host: "localhost",
		Port: 6379,
	})*/
	return bnsList, nil
}
