package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

var client = &http.Client{Timeout: 3 * time.Second}

func GetJson(Url string, destination any) error {
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return fmt.Errorf("erreur de requete: %d", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(destination)
	if err != nil {
		return fmt.Errorf("erreur de traduction: %d", err)
	}
	return nil
}
