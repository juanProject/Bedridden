package BedManager

import (
	"bytes"
	"encoding/json"
	"github.com/hidez8891/shm"
	"log"
)

func (_manager *bedManager) sync(){
	segment, err := shm.Open("bedManagment", 256)
	if err != nil {
		panic(err.Error())
	}
	defer segment.Close()

	content := make([]byte, 256)
	_, err = segment.ReadAt(content, 0)
	if err != nil{
		panic(err.Error())
	}

	trimContent := bytes.Trim(content, string(byte(0)))
	if len(trimContent) == 0 {
		_, err = segment.Write(_manager.toJson())
		if err != nil {
			panic(err.Error())
		}
	} else {
		err = _manager.updateFromJson(trimContent)
		if err != nil {
			panic(err.Error())
		}

	}

}

func (_manager *bedManager) update(){
	segment, err := shm.Open("bedManagment", 256)
	if err != nil {
		segment, err = shm.Create("bedManagment", 256)
		if err != nil{
			panic(err.Error())
		}
	}
	defer segment.Close()

	segment.WriteAt(_manager.toJson(), 0)
}

func (_manager *bedManager) toJson() []byte{
	_json, err := json.Marshal(_manager)
	if err != nil {
		log.Fatal(err)
	}
	toPad := 256 - len(_json)
	padding := make([]byte, toPad)
	return bytes.Join([][]byte{_json, padding}, []byte(""))
}

func (_manager *bedManager) updateFromJson(jsonSource []byte) error{
	newManager := bedManager{0,0}
	err := json.Unmarshal(jsonSource, &newManager)
	if err != nil {
		return err
	}
	_manager.TotalBeds = newManager.TotalBeds
	_manager.NaBeds = newManager.NaBeds
	return nil
}

func (_manager *bedManager) AddPatient(){
	_manager.sync()
	_manager.addPatient()
	_manager.update()
}

func (_manager *bedManager) SubPatient(){
	_manager.sync()
	_manager.subPatient()
	_manager.update()
}

func (_manager *bedManager) DisplayState(){
	_manager.sync()
	_manager.displayState()
}
