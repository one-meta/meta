package config

import "fmt"

type DB struct {
	MaxIdleConns    int      `json:"maxIdleConns"`
	MaxOpenConns    int      `json:"maxOpenConns"`
	ConnMaxLifetime int      `json:"connMaxLifetime"`
	MySQL           MySQL    `json:"mySQL"`
	PostGres        PostGres `json:"postGres"`
	MariaDB         MariaDB  `json:"mariaDB"`
	Sqlite3         Sqlite3  `json:"sqlite3"`
}
type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DataBase string `json:"dataBase"`
	Param    string `json:"param"`
}
type PostGres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DataBase string `json:"dataBase"`
	Param    string `json:"param"`
}
type MariaDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DataBase string `json:"dataBase"`
	Param    string `json:"param"`
}
type Sqlite3 struct {
	Mode     string `json:"mode"`
	DataBase string `json:"dataBase"`
}

func (p *PostGres) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?%s", p.Username, p.Password, p.Host, p.DataBase, p.Param)
}
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DataBase, m.Param)
}

func (m *MariaDB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DataBase, m.Param)
}

func (s *Sqlite3) DSN() string {
	return fmt.Sprintf("file:%s?mode=%s&cache=shared&_fk=1", s.DataBase, s.Mode)
}
