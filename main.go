package main

import (
	"./src/BedManager"
	"fmt"
	"github.com/hidez8891/shm"
)

func main(){
	manager := BedManager.NewBedManagerFromFile("data.json")
	segment, err := shm.Create("bedManagment", 256)
	if err != nil{
		panic(err.Error())
	}
	defer segment.Close()
	for {
		choice := '0'
		exit := 0
		printMenuUpperPart()
		fmt.Println("/   \\                     Bedridden                             /   \\")
		fmt.Println("\\___/                                                           \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 1.     Add Patient                        \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 2.     Sub Patient                        \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 3.     Status                             \\___/")
		fmt.Println("/   \\                                                           /   \\")
		fmt.Println("\\___/                 q.     Exit                               \\___/")
		printMenuBottomPart()
		_, _ = fmt.Scanf("%c\n", &choice)
		switch choice {
		case '1':
			manager.AddPatient()
			break
		case '2':
			manager.SubPatient()
			break
		case '3':
			manager.DisplayState()
			break
		case 'q':
			exit = 1
			break
		default:
			fmt.Println("Incorrect input, try again...")
		}
		if exit == 1 {
			fmt.Println("Bye Bye! Come back soon!")
			//time.Sleep(2 * time.Second)
			break
		}
		//time.Sleep(3 * time.Second)
	}
}

func printMenuUpperPart() {
	fmt.Println("     ___     ___     ___     ___     ___     ___     ___     ___")
	fmt.Println(" ___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___")
	fmt.Println("/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\")
	fmt.Println("\\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
	fmt.Println("/   \\___/                                                   \\___/   \\")
	fmt.Println("\\___/                                                           \\___/")
}
func printMenuBottomPart() {
	fmt.Println("/   \\___                                                     ___/   \\")
	fmt.Println("\\___/   \\___     ___     ___     ___     ___     ___     ___/   \\___/")
	fmt.Println("/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\")
	fmt.Println("\\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
	fmt.Println("    \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/   \\___/")
}