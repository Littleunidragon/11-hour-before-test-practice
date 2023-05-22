package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/TwiN/go-color"
	"log"
)

type sastavdala struct {
	veids, modelis string
	cena string
	id int
}
type list struct {
	all []sastavdala
}

func Add(Scanner *bufio.Scanner, l list) list {
	var new sastavdala
	for new.veids == "" {
	fmt.Println(color.Green + "Veids: " + color.Reset)
	Scanner.Scan()
	new.veids = Scanner.Text()		
	}
	for new.modelis == "" {
	fmt.Println(color.Green +"Modelis: "+ color.Reset)
	Scanner.Scan()
	new.modelis = Scanner.Text()
	}
	for new.cena == "" {
	fmt.Println(color.Green +"Cena: "+ color.Reset)
	Scanner.Scan()
	new.cena = Scanner.Text()	
	}
	new.id = GiveID(l)
	l.all = append(l.all, new)
	return l
}
func GiveID(l list) int {
	return len(l.all)+1
}
func List(l list){
		fmt.Println("-Personālā datora sastāvdaļa-")	
	for i := range l.all {
		fmt.Printf("%v)\n", i+1)
		fmt.Printf("ID: %v\n",l.all[i].id)
		fmt.Printf("Veids: %s\n",l.all[i].veids)
		fmt.Printf("Modelis: %s\n", l.all[i].modelis)
		fmt.Printf("Cena: %s\n", l.all[i].cena)
	}
}
func Change(l list, a int, scanner *bufio.Scanner) list {
for i := range l.all {
	if l.all[i].id == a {
		fmt.Println(color.Green + "Veids: " + color.Reset)
		scanner.Scan()
		scanner.Scan()
		l.all[i].veids = scanner.Text()
		fmt.Println(color.Green +"Modelis: "+ color.Reset)
		scanner.Scan()
		l.all[i].modelis = scanner.Text()
		fmt.Println(color.Green +"Price: "+ color.Reset)
		scanner.Scan()
		l.all[i].cena = scanner.Text()	
	}
}
return l
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("starting...")
	var action int
	var l list
	for {
		fmt.Println(color.Green +`Type:
		1 - to add
		2 - to read
		3 - to change
		4 - to exit` + color.Reset)
		fmt.Scan(&action)
		if action == 1 {
			l = Add(scanner, l)
		}
		if action == 2 {
			List(l)
		}
		if action == 3 {
			fmt.Println(color.Green +"Enter ID of part you want change: " + color.Reset)
			var a int
			fmt.Scan(&a)
			l = Change(l, a, scanner)
		}
		if action == 4 {
			os.Exit(1)
		}
		f, err := os.Create("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		for i := range l.all {
			_, err := f.WriteString("-Personāla datora sastāvdaļa-\n")
			_, err1 := f.WriteString("Veids: " +l.all[i].veids + "\n")
			_, err2 := f.WriteString("Modelis: " +l.all[i].modelis + "\n")
			_, err3 := f.WriteString("Modelis: " + l.all[i].cena + "EUR\n")		
			if err != nil {
				log.Fatal(err)
			}
			if err1 != nil {
				log.Fatal(err1)
			}
			if err2 != nil {
				log.Fatal(err2)
			}
			if err3 != nil {
				log.Fatal(err3)
			}
		}
	}
}
