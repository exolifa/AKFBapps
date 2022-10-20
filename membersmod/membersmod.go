package membersmod

import (
	Fishmod "AKFBapps/fishmod"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Members struct {
	Members []Member `json:"members"`
}
type Member struct {
	Memberid    string            `json:"memberid"`
	Name        string            `json:"name"`
	Firstname   string            `json:"firstname"`
	Email       string            `json:"email"`
	Maintenance []Fishmod.Poisson `json:"maintenance"`
}

var Memlist = Members{}

// creer une fonction Display_member qui va faire la liste de tous les members dans Memlist
// la fonction affiche le Firstname,name, email
func Display_member() {
	for x, _ := range Memlist.Members {
		fmt.Printf("%s, %s, %s\n", Memlist.Members[x].Name, Memlist.Members[x].Firstname, Memlist.Members[x].Email)
	}
}

// Creer une fonction Get_member_by_name qui accepte le name comme string en parametre et renvoie la meme member correspondante
// elle retourne un json de type Member
func Get_member_by_name(name string) (member *Member, err error) {
	for x, y := range Memlist.Members {
		if name == y.Name {
			member = &Memlist.Members[x]
			return
		}
	}
	err = fmt.Errorf("%s member not found", name)
	return
}

// create a function to
// get all files from current directory where file name contains fishes_v and file name contains .json
func GetFilesList() []string {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info != nil {
			if info.IsDir() {
				// do nothing
			} else {
				if strings.Contains(path, ".json") && strings.Contains(path, "members_v") {
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
	return "members_v" + strconv.Itoa(lastversion+1) + ".json"
}

func init() {
	filename := getLastFile()
	if Load_memlist(filename) {
		fmt.Println("Loaded member list from file:", filename)
	} else {
		panic(errors.New("Loaded member list from file: " + filename + " failed."))
	}
}

// générer une fonction Load_memlist qui va charger Memlist à partir d'un fichier json appelé "members.json"
// la fonction retourne un booleen suivant que le chargement s'est bien passé ou pas
func Load_memlist(filename string) (success bool) {
	var members Members
	f, err := os.Open(filename)
	if err == nil {
		defer f.Close()
		err = json.NewDecoder(f).Decode(&members)
		if err == nil {
			Memlist = members
			success = true
		} else {
			success = false
		}
	} else {
		success = false
	}
	return
}

// EnregistrerMemlist va enregistrer Memlist dans un fichier json appelé "members.json"
// la fonction retourne un booleen suivant que le chargement s'est bien passé ou pas
func EnregistrerMemlist() (success bool) {
	var filename = getNextVersion()
	var members = Memlist
	f, err := os.Create(filename)
	if err == nil {
		defer f.Close()
		err = json.NewEncoder(f).Encode(members)
		if err == nil {
			success = true
		} else {
			success = false
		}
	} else {
		success = false
	}
	return
}

// ecrire une fonction Get_Memlist_Maintenance qui accepte en paramètre un email string
// l' email est la cle pour identifier un element de memlist
// La fonction va retourner la table de Fishmod/Poisson contenue dans Maintenance
func Get_Memlist_Maintenance(email string) []Fishmod.Poisson {
	var table = []Fishmod.Poisson{}
	// rechercher le meme element
	for _, el := range Memlist.Members {
		if el.Email == email {
			return el.Maintenance
		}
	}
	return table
}

// generate a function to add members to memlist
func Add_Member(id string, email string, nom string, prenom string, maintenance []Fishmod.Poisson) bool {
	// on parcour la table de memlist
	for _, el := range Memlist.Members {
		if el.Email == email {
			return false
		}
	}
	Memlist.Members = append(Memlist.Members, Member{id, nom, prenom, email, maintenance})
	return true

}

// générer une fonction qui ajoute une fishmod.Poisson a Maintenance pour une Member identifié par
// email.
func Add_Maintenance(email string, poisson Fishmod.Poisson) bool {
	// on parcour la table de memlist
	for i, el := range Memlist.Members {
		if el.Email == email {
			fmt.Printf("found email %s on ajoute %s", email, poisson)
			Memlist.Members[i].Maintenance = append(Memlist.Members[i].Maintenance, poisson)
			fmt.Printf("from the struct %s", el.Maintenance)
			return true
		}
	}
	return false
}

// générer une fonction qui delete une fishmod.Poisson a Maintenance pour une Member identifié par
// email.
func Delete_Maintenance(email string, poisson Fishmod.Poisson) bool {
	// on parcour la table de memlist
	for i, el := range Memlist.Members {
		if el.Email == email {
			fmt.Printf("found email %s on ajoute %s", email, poisson)
			Memlist.Members[i].Maintenance = append(Memlist.Members[i].Maintenance[:], Memlist.Members[i].Maintenance[:len(Memlist.Members[i].Maintenance)-1]...)
			fmt.Printf("from the struct %s", el.Maintenance)
			return true
		}
	}
	return false
}

// retourne une table avec la liste des emails à traiter
func Email_list() string {
	var liste string
	for _, el := range Memlist.Members {
		liste = liste + el.Email + "\n"
	}
	return liste
}

// retourne la liste des membres
func Members_list() string {
	var liste string
	for _, el := range Memlist.Members {
		liste = liste + el.Name + " " + el.Firstname + "\n"
	}
	return liste
}

// chercher les members qui ont un poisson dans leur maintenance
// le poisson est identifié par genre,espece,population la fonction retourne une table des emails
func Membership_maintenance(genre, espece, population string) string {
	var liste string
	liste = ""
	for _, el := range Memlist.Members {
		// on regarde si la personne n'a plus de maintenance parmi les membres de la meme espace et genre
		// si c'est le cas on ajoute dans la liste des emails
		for i, e := range el.Maintenance {
			if e.Genre == genre && e.Espece == espece && e.Population == population {
				liste = liste + el.Email + "\n"
			}
		}
	}
	if liste == "" {
		liste = "Nous n'avons pas encore de maintenance pour cette espece\n"
	}
	return liste
}
