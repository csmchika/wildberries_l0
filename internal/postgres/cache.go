package postgres

import (
	"log"
	"wb-l0/internal/postgres/models"
)

type Cache struct {
	DBInst     *DB
	CacheModel map[int64]*models.Model
}

func NewCache(db *DB) *Cache {
	csh := Cache{}
	csh.Init(db)
	return &csh
}

func (c *Cache) Init(db *DB) {
	c.DBInst = db
	c.GetCacheFromDatabase()
}

func (c *Cache) GetCacheFromDatabase() {
	log.Printf("check & download cache from database\n")
	c.CacheModel = c.DBInst.SetCache(c)

}

func (c *Cache) AddModel(uid int64, m *models.Model) {
	c.CacheModel[uid] = m
}

func (c *Cache) CountElems() int {
	return len(c.CacheModel)
}
