package repository

import "github.com/go-pg/pg/v10"

type CounterDao struct {
	tableName struct{} `pg:"counter"`
	ID        uint64   `pg:"id,pk"`
	Value     uint64   `pg:"value"`
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
		Where("id = 1").
		Returning("value").
		Update("value = value + 1")

	if err != nil {
		if err == pg.ErrNoRows {
			counter = CounterDao{
				ID:    1,
				Value: 1,
			}
			_, err = r.db.Model(&counter).Insert()
			if err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	return counter.Value, nil
}
