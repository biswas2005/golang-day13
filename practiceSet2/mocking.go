package practiceset2

import "fmt"

type Database interface {
	Save(data string)
}
type RealDB struct{}

func (r RealDB) Save(data string) {
	fmt.Println("RealDB: Saving Data to Database", data)
}

type MockDB struct {
	Records []string
}

func (m *MockDB) Save(data string) {
	m.Records = append(m.Records, data)
	fmt.Println("MockDB: Pretending to save data", data)
}

func Process(data string, db Database) {
	result := "Processed:" + data
	db.Save(result)
}

func Mocking() {
	realDB := RealDB{}
	Process("Production Data", realDB)

	mockDB := &MockDB{}
	Process("Test Data", mockDB)

	fmt.Println("MockDB stored records:", mockDB.Records)
}
