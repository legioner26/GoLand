package post

import (
	"encoding/csv"
	"os"
	cities "skillbox/attestat/models"
	"strconv"
)

// NewSQLPostRepo retunrs implement of post repository interface
const (
	nameIdx = iota
	regionIdx
	districtIdx
	populationIdx
	foundationIdx
)
const (
	errNotFoundId = "ERROR: CITY WITH THIS ID WAS NOT FOUND"
)

type DataBase struct {
	records map[int][]string
	lastID  int
}

type CityListDB struct {
	db *DataBase
}

func NewCityListDB(db *DataBase) *CityListDB {
	return &CityListDB{db: db}
}
func NewDataBase(filePath string) (*DataBase, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	cities, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	cityList := new(DataBase)
	cityList.records = make(map[int][]string)

	if len(cities) == 0 {
		return cityList, nil
	}

	for _, city := range cities {
		id, _ := strconv.Atoi(city[0])
		if err != nil {
			return nil, err
		}
		cityList.records[id] = make([]string, 5)
		copy(cityList.records[id], city[1:])
	}

	cityList.lastID = 0
	for cityID := range cityList.records {
		if cityList.lastID < cityID {
			cityList.lastID = cityID
		}
	}
	return cityList, nil
}

// SaveCSV saves the database to a csv file.
// Creates a csv file and copies data from the DataBase structure into it.
// If successful, overwrites the csv file into the filePath file.
func (db *DataBase) SaveCSV(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var data [][]string
	for id, description := range db.records {
		var cityLine []string
		cityLine = append(cityLine, strconv.Itoa(id))
		cityLine = append(cityLine, description...)
		data = append(data, cityLine)
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}
	return nil
}

// Create places a new city in the database and assigns it an id.
// Returns the id of the city
func (r *CityListDB) Create(city cities.CityRequest) (string, error) {
	newID := r.newId()
	citySlice := []string{
		city.Name,
		city.Region,
		city.District,
		strconv.Itoa(city.Population),
		strconv.Itoa(city.Foundation),
	}

	r.db.records[newID] = citySlice
	return strconv.Itoa(newID), nil
}

// Delete deletes the city with the id from the database
func (r *CityListDB) Delete(id int) error {
	_, ok := r.db.records[id]
	if !ok {
		return errors.New(errNotFoundId)
	}

	delete(r.db.records, id)
	return nil
}

// SetPopulation updates the city population by id in the database
func (r *CityListDB) SetPopulation(id, population int) error {
	_, ok := r.db.records[id]
	if !ok {
		return errors.New(errNotFoundId)
	}

	r.db.records[id][populationIdx] = strconv.Itoa(population)
	return nil
}

// GetFromRegion returns the list of cities by region
func (r *CityListDB) GetFromRegion(region string) ([]string, error) {
	return r.findCities(regionIdx, region), nil
}

// GetFromDistrict returns the list of cities by district
func (r *CityListDB) GetFromDistrict(district string) ([]string, error) {
	return r.findCities(districtIdx, district), nil
}

// GetFromPopulation returns the list of cities by population
func (r *CityListDB) GetFromPopulation(start, end int) ([]string, error) {
	return r.findRangeCities(populationIdx, start, end)
}

// GetFromFoundation returns the list of cities by foundation
func (r *CityListDB) GetFromFoundation(start, end int) ([]string, error) {
	return r.findRangeCities(foundationIdx, start, end)
}

// GetFull searches for a city by id.
// If successful, it returns full information about the city
func (r *CityListDB) GetFull(id int) (*cities.City, error) {
	_, ok := r.db.records[id]
	if !ok {
		return nil, errors.New(errNotFoundId)
	}

	city := new(cities.City)
	city.Id = id
	city.Name = r.db.records[id][nameIdx]
	city.Region = r.db.records[id][regionIdx]
	city.District = r.db.records[id][districtIdx]

	population, err := strconv.Atoi(r.db.records[id][populationIdx])
	if err != nil {
		return nil, err
	}
	city.Population = population

	foundation, err := strconv.Atoi(r.db.records[id][foundationIdx])
	if err != nil {
		return nil, err
	}
	city.Foundation = foundation

	return city, nil
}

// newId returns a free identifier in the database
func (r *CityListDB) newId() int {
	r.db.lastID++
	return r.db.lastID
}

// findCities searches for cities based on idxType.
// Returns a list of found cities
func (r *CityListDB) findCities(idxType int, searchText string) []string {
	cityNames := make([]string, 0)
	for _, cityLine := range r.db.records {
		if cityLine[idxType] == searchText {
			cityNames = append(cityNames, cityLine[nameIdx])
		}
	}
	return cityNames
}

// findRangeCities searches for cities in the range [start; end] based on idxType.
// Returns a list of found cities
func (r *CityListDB) findRangeCities(idxType int, start, end int) ([]string, error) {
	if start > end {
		start, end = end, start
	}

	cityNames := make([]string, 0)
	for _, cityLine := range r.db.records {

		year, err := strconv.Atoi(cityLine[idxType])
		if err != nil {
			return nil, err
		}

		if year >= start && year <= end {
			cityNames = append(cityNames, cityLine[nameIdx])
		}
	}
	return cityNames, nil
}
