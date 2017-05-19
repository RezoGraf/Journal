package models

type Config struct {
	Database struct {
			 Host string `json:"host"`
			 Password string `json:"password"`
		 } `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}