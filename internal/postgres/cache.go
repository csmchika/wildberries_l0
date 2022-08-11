package postgres

import (
	"bytes"
	"encoding/json"
	"log"
	"wb-l0/internal/postgres/models"
)

type Cache struct {
	DBInst     *DB
	CacheModel map[int]*models.Model
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
	log.Printf("Проверяю и загружаю кеш из БД\n")
	c.CacheModel = c.DBInst.SetCache(c)

}

func (c *Cache) AddModelCache(uid int, m *models.Model) {
	c.CacheModel[uid-1] = m
}

func (c *Cache) CountElems() int {
	return len(c.CacheModel)
}

func (c *Cache) ToString(id int) string {
	buf := bytes.Buffer{}
	marshal, _ := json.Marshal(c.CacheModel[id])
	_ = json.Indent(&buf, marshal, "", "\t")
	return buf.String()
}
