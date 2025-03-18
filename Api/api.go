package Api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Structure représentant un artiste
type Artist struct {
	ID              int                 `json:"id"`
	Image           string              `json:"image"`
	Name            string              `json:"name"`
	Members         []string            `json:"members"`
	CreationDate    int                 `json:"creationDate"`
	FirstAlbum      string              `json:"firstAlbum"`
	LocationsURL    string              `json:"locations"`
	ConcertDatesURL string              `json:"concertDates"`
	RelationsURL    string              `json:"relations"`
	Relations       map[string][]string `json:"-"` // Champ utilisé pour stocker les relations (exclu du JSON)
}

// Structure représentant les lieux des concerts
type Location struct {
	Locations []string `json:"locations"` // Liste des lieux où l'artiste a joué
}

// Structure représentant les dates de concert
type ConcertDate struct {
	Dates []string `json:"dates"` // Liste des dates de concerts
}

// Structure représentant les relations entre dates et lieux de concert
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"` // Mapping entre les dates et les lieux des concerts
}

// Fonctions pour récupérer les données des APIs

// Récupère la liste des artistes depuis l'API
func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	return fetchData[[]Artist](url) // Utilisation de la fonction générique fetchData
}

// Récupère les lieux des concerts d'un artiste donné
func GetLocations(url string) ([]string, error) {
	var locData Location
	err := fetchDataFromURL(url, &locData) // Récupération des données JSON depuis l'URL
	if err != nil {
		return nil, err // Retourne une erreur en cas d'échec
	}
	return locData.Locations, nil // Retourne la liste des lieux
}

// Récupère les dates de concert d'un artiste donné
func GetConcertDates(url string) ([]string, error) {
	var dateData ConcertDate
	err := fetchDataFromURL(url, &dateData) // Récupération des données JSON depuis l'URL
	if err != nil {
		return nil, err // Retourne une erreur en cas d'échec
	}
	return dateData.Dates, nil // Retourne la liste des dates
}

// Récupère les relations entre dates et lieux de concert d'un artiste donné
func GetRelations(url string) (map[string][]string, error) {
	var relData Relation
	err := fetchDataFromURL(url, &relData) // Récupération des données JSON depuis l'URL
	if err != nil {
		return nil, err // Retourne une erreur en cas d'échec
	}
	return relData.DatesLocations, nil // Retourne le mapping entre dates et lieux
}

// Fonction générique pour récupérer des données JSON depuis une URL
func fetchData[T any](url string) (T, error) {
	resp, err := http.Get(url) // Effectue une requête HTTP GET sur l'URL
	if err != nil {
		return *new(T), fmt.Errorf("erreur lors de la requête : %w", err) // Retourne une erreur en cas d'échec
	}
	defer resp.Body.Close() // Ferme la réponse HTTP à la fin de la fonction

	body, err := ioutil.ReadAll(resp.Body) // Lit le corps de la réponse
	if err != nil {
		return *new(T), fmt.Errorf("erreur lors de la lecture de la réponse : %w", err) // Retourne une erreur si la lecture échoue
	}

	var data T
	err = json.Unmarshal(body, &data) // Décode le JSON en structure Go
	if err != nil {
		return *new(T), fmt.Errorf("erreur lors du décodage JSON : %w", err) // Retourne une erreur en cas d'échec
	}

	return data, nil // Retourne les données récupérées
}

// Fonction pour récupérer et décoder des données JSON à partir d'une URL spécifique
func fetchDataFromURL(url string, target any) error {
	resp, err := http.Get(url) // Effectue une requête HTTP GET sur l'URL
	if err != nil {
		return fmt.Errorf("erreur lors de la requête : %w", err) // Retourne une erreur en cas d'échec
	}
	defer resp.Body.Close() // Ferme la réponse HTTP à la fin de la fonction

	body, err := ioutil.ReadAll(resp.Body) // Lit le corps de la réponse
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture de la réponse : %w", err) // Retourne une erreur si la lecture échoue
	}

	err = json.Unmarshal(body, target) // Décode le JSON dans la structure cible
	if err != nil {
		return fmt.Errorf("erreur lors du décodage JSON : %w", err) // Retourne une erreur en cas d'échec
	}

	return nil // Retourne nil si tout s'est bien passé
}
