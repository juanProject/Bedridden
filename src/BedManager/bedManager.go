package BedManager

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"math"
)

type bedManager struct {
	TotalBeds      int
	NaBeds         int
}

func NewBedManagerFromFile(path string) *bedManager{
	manager := bedManager{0, 0}
	getData(path, &manager)
	return &manager
}

func getData(path string, manager *bedManager){
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, manager)
	if err != nil {
		log.Fatal(err)
	}
}

func (_manager *bedManager) addPatient(){
	if _manager.TotalBeds > _manager.NaBeds {
		_manager.NaBeds += 1
	} else {
		fmt.Println("No beds available")
	}
}

func (_manager *bedManager) subPatient(){
	if _manager.NaBeds > 0 {
		_manager.NaBeds -= 1
	} else {
		fmt.Println("No patients")
	}
}

func (_manager *bedManager) displayState(){
	bed := html.UnescapeString("&#128719;")
	naBedsPercent := float64(_manager.NaBeds) / float64(_manager.TotalBeds) * 10
	availableBedsPercent := math.Abs(naBedsPercent - 10)

	fmt.Println("Beds Available:", availableBedsPercent * 10, "%" )
	for i := 0.0; i < naBedsPercent; i++ {
		fmt.Print("\033[1;31m", bed, "\033[0m")
	}
	for i := 0.0; i < availableBedsPercent; i++ {
		fmt.Print("\033[1;30m", bed, "\033[0m")
	}
	fmt.Println()
}