package main



import (
	"net/http"
)


type Dinosaur struct {
	Nom              string   `json:"name"`
	Nom_Scientifique string   `json:"scientific_name"`
	Classification   string   `json:"classification"`
	Periode          []string `json:"period"`
	Alimentation     string   `json:"diet"`
	Hauteur          string   `json:"height_m"`
	Largueur         string   `json:"length_m"`
	Poids            string   `json:"weight_kg"`
	Id               string   `json:"slug"`
}

type Periode struct {
	Nom            string `json:"name"`
	Ere            string `json:"era"`
	Annee_debut    string `json:"start_mya"`
	Annee_fin      string `json:"end_mya"`
	Description    string `json:"description"`
	Nombre_especes string `json:"species_count"`
	Id             string `json:"slug"`
}

type Site struct {
	Nom                string `json:"formation_name"`
	Pays               string `json:"country_name"`
	Nombre_decouvertes int    `json:"discovery_count"`
	Id                 string `json:"slug"`
}



func GetJson(Url,  ) {
	req, err:= 
}