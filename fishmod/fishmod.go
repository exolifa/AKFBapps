package Fishmod

// fishmod package initialiases from reading a json file into a structure Poisson of Type Poisson
// type Poisson is a struct witch contains 3 strings: Genre, Espece, Population
//module also provide published functions :
//func Find_fish_by_Genre returns all Poisson with the given Genre
//func Find_fish_by_Espece returns all Poisson with the given Espece
//func Find_fish_by_Population returns all Poisson with the given Population

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Catalog_fish is a structure of array of Poisson
type Catalog_fish struct {
	Poissons []Poisson `json:"Fishes"`
}

// Poisson is a struct witch contains 3 strings: Genre, Espece, Population
type Poisson struct {
	Genre      string `json:"Genre"`
	Espece     string `json:"Espece"`
	Population string `json:"Population"`
}

var Catalogue_poissons = Catalog_fish{}

// create a function to
// get all files from current directory where file name contains fishes_v and file name contains .json
func GetFilesList() []string {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info != nil {
			if info.IsDir() {
				// do nothing
			} else {
				if strings.Contains(path, ".json") && strings.Contains(path, "fishes_v") {
					// check if file name contains fishes_v and does not start with fishes_v
					//check if file name contains fishes_v
					files = append(files, path)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	return files
}
func getLastFile() string {
	filetable := GetFilesList()
	re := regexp.MustCompile(`[-]?\d[\d,]*`)
	lastversion := 0
	mostcurrfile := ""
	for i := 0; i < len(filetable); i++ {
		version, _ := strconv.Atoi(re.FindAllString(filetable[i], -1)[0])
		fmt.Println(version)
		if version > lastversion {
			lastversion = version
			mostcurrfile = filetable[i]
		}
	}
	fmt.Println(mostcurrfile)
	return mostcurrfile
}

func getNextVersion() string {
	filetable := GetFilesList()
	re := regexp.MustCompile(`[-]?\d[\d,]*`)
	lastversion := 0
	for i := 0; i < len(filetable); i++ {
		version, _ := strconv.Atoi(re.FindAllString(filetable[i], 1)[0])
		if version > lastversion {
			lastversion = version
		}
	}
	return "fishes_v" + strconv.Itoa(lastversion+1) + ".json"
}

func init() {
	filename := getLastFile()
	fmt.Println(filename) // import content of json file named fishes.json into Fish
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal([]byte(file), &Catalogue_poissons)
}

func MergeJSON(filename string) {
	var temp_poissons = Catalog_fish{}
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal([]byte(file), &temp_poissons)
	for _, poisson := range temp_poissons.Poissons {
		Add_fish(poisson.Genre, poisson.Espece, poisson.Population)
	}
}

// Display_fishes will print all fishes in Catalogue_poissons
func Display_fishes() {
	for _, Fish := range Catalogue_poissons.Poissons {
		fmt.Println(Fish.Genre, Fish.Espece, Fish.Population)
	}
}

// Find_fish_by_Genre will return an array of fishes which have this Genre
func Find_fish_by_Genre(Genre string) []Poisson {
	var fishes []Poisson
	for _, Fish := range Catalogue_poissons.Poissons {
		if strings.Contains(Fish.Genre, Genre) {
			fishes = append(fishes, Fish)
		}
	}
	return fishes
}

// Find_fish_by_Espece will return an array of fishes which have this Espece
func Find_fish_by_Espece(Espece string) []Poisson {
	var fishes []Poisson
	for _, Fish := range Catalogue_poissons.Poissons {
		if strings.Contains(Fish.Espece, Espece) {
			fishes = append(fishes, Fish)
		}
	}
	return fishes
}

// Find_fish_by_genre_and_espece will return an array of fishes which have this genre and this espece
func Find_fish_by_genre_and_espece(Genre string, Espece string) []Poisson {
	var fishes []Poisson
	for _, Fish := range Catalogue_poissons.Poissons {
		if strings.Contains(Fish.Genre, Genre) && strings.Contains(Fish.Espece, Espece) {
			fishes = append(fishes, Fish)
		}
	}
	return fishes
}

// Find_fish_by_Population will return an array of fishes which have this Population
func Find_fish_by_Population(Population string) []Poisson {
	var fishes []Poisson
	for _, Fish := range Catalogue_poissons.Poissons {
		if strings.Contains(Fish.Population, Population) {
			fishes = append(fishes, Fish)
		}
	}
	return fishes
}

// ADD_fish is a function to aadd new fish to the Catalogue_poissons
// from genre, espece, population giving as parameters
// if the new fish already exists in the Catalogue_poissons, it will be added again
func Add_fish(Genre string, Espece string, Population string) {
	for _, Fish := range Catalogue_poissons.Poissons {
		if Fish.Genre == Genre && Fish.Espece == Espece && Fish.Population == Population {
			return
		}
	}
	Catalogue_poissons.Poissons = append(Catalogue_poissons.Poissons, Poisson{Genre, Espece, Population})
}

// Del_fish is a fonction to delete a fish based on genre, espece,population received in parameters
func Del_fish(genre string, espece string, population string) {
	for index, fish := range Catalogue_poissons.Poissons {
		if fish.Genre == genre &&
			fish.Espece == espece &&
			fish.Population == population {
			//Delete the fish at that position
			Catalogue_poissons.Poissons = append(Catalogue_poissons.Poissons[:index], Catalogue_poissons.Poissons[index+1:]...)
		}
	}
}

// Save_catalogue is a function to save the Catalogue_poissons to a json file with title passed in parameter
// the catalogue primary item is named Fish that is an array of all the items in Catalogue_poissons
func Save_catalogue() {
	title := getNextVersion()
	file, _ := os.Create(title)
	defer file.Close()
	//save to the json file
	err := json.NewEncoder(file).Encode(&Catalogue_poissons)
	if err != nil {
		panic(err)
	}
}

// Check_Genre est une fonction qui vérifie que le genre existe dans Catalogue_poissons
// la fonction retourne un boolean pour savoir si le genre existe
func Check_Genre(genre string) bool {
	for _, poisson := range Catalogue_poissons.Poissons {
		if poisson.Genre == genre {
			return true
		}
	}
	return false
}

// Check_Espece est une fonction qui vérifie que le espece existe dans Catalogue_poissons
// la fonction retourne un boolean pour savoir si le espece existe
func Check_Espece(espece string) bool {
	for _, poisson := range Catalogue_poissons.Poissons {
		if poisson.Espece == espece {
			return true
		}
	}
	return false
}

// Check_Population est une fonction qui vérifie que la population existe dans Catalogue_poissons
// la fonction retourne un boolean pour savoir si la population existe
func Check_Population(population string) bool {
	for _, poisson := range Catalogue_poissons.Poissons {
		if poisson.Population == population {
			return true
		}
	}
	return false
}
