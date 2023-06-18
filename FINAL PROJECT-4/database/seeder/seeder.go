package seeder

import (
	"finalproject4/database/faker"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterAdmin(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.Admin(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	seeder := RegisterAdmin(db)
	for _, s := range seeder {
		err := db.Create(s.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
