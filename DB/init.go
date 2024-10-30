package db

import (
	"fmt"
	"os"
	"strconv"

	colors "YBOOST/Lib/Func/Colors"
	json "YBOOST/Lib/Func/JSON"
	terminal "YBOOST/Lib/Func/Terminal"
	types "YBOOST/Lib/Types"
)

func InitDatabase() {
	terminal.ClearTerminal()
	fmt.Println("Initialisation de la base de données...")
	if dbExists() {
		fmt.Println("La base de données existe déjà.")
	} else {
		askDatabaseInfo()
	}
}

// ////////////////////////////////////////////////////////////////////
// Vérifier si la base de données existe déjà ////////////////////////
// ////////////////////////////////////////////////////////////////////
func dbExists() bool {
	// Vérifier si il y a un fichier config.json dans le dossier YBOOST/DB
	path := "DB/config.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

//////////////////////////////////////////////////////////////////////
// Demander à l'utilisateur de créer la base de données //////////////
//////////////////////////////////////////////////////////////////////

func askDatabaseInfo() {
	terminal.ClearTerminal()
	g, b, _, d := colors.GetTColors()
	fmt.Println(g, "Veuillez entrer les informations pour la base de données.")
	fmt.Println(g, "Host (exemple : ", b, "localhost - 10.6.0.161):", d)
	var host string
	fmt.Scanln(&host)
	fmt.Println(g, "Port (exemple : ", b, "3306):", d)
	var port string
	fmt.Scanln(&port)
	fmt.Println(g, "Nom de la base de données (exemple : ", b, "cocktails):", d)
	var dbname string
	fmt.Scanln(&dbname)
	fmt.Println(g, "Nom d'utilisateur (exemple : ", b, "root):", d)
	var user string
	fmt.Scanln(&user)
	fmt.Println(g, "Mot de passe (exemple : ", b, "password):", d)
	var password string
	fmt.Scanln(&password)
	terminal.ClearTerminal()

	//transformer port en int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Erreur lors de la conversion du port en int:", err)
		return
	}

	config := types.MySQLConfig{
		User:     user,
		Password: password,
		Database: dbname,
		Host:     host,
		Port:     portInt,
	}

	// Créer le fichier config.json
	json.WriteJsonConfig(config, "DB/config.json")
}
