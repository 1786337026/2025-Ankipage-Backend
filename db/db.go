package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"time"
)

type Note struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Title     string    `gorm:"size:255;not null"`
	Url       string    `gorm:"size:255;not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
type User struct {
	ID        int       `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"` // 加密后的密码
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

var DB *gorm.DB

func InitDB() {
	config, err := LoadConfig("F:\\project\\2025-Ankipage-Backend\\config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to the database using the configuration
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		panic(1)
	}

	// 自动迁移
	DB.AutoMigrate(&Note{}, &User{})

}
