package repository

import (
	"github.com/go-pg/pg/v10"
)

type CounterDao struct {
    tableName struct{} `pg:"counter,alias:counter"`
    ID        uint64   `pg:"id,pk"`
    CurrentValue     uint64   `pg:"current_value,use_zero"`
}

type CounterRepositoryImpl struct {
	db *pg.DB
}

func NewCounterRepositoryImpl(db *pg.DB) *CounterRepositoryImpl {
	return &CounterRepositoryImpl{
		db: db,
	}
}

func (r *CounterRepositoryImpl) GetNextCounter() (uint64, error) {
	var counter CounterDao

	_, err := r.db.Model(&counter).
		Where("id = ?", 1).
		Set("current_value = current_value + 1").
		Returning("*").
		Update()

	if err != nil {
		if err == pg.ErrNoRows {
			counter = CounterDao{
				ID:    1,
				CurrentValue: 1,
			}
			_, err = r.db.Model(&counter).Insert()
			if err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	return counter.CurrentValue, nil
}
