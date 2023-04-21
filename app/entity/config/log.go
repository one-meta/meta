package config

type Log struct {
	Web Web `json:"web"`
	App App `json:"app"`
}
type Web struct {
	Format     string     `json:"format"`
	TimeFormat string     `json:"timeFormat"`
	TimeZone   string     `json:"timezone"`
	Output     string     `json:"output"`
	Lumberjack Lumberjack `json:"lumberjack"`
}
type App struct {
	LogLevel   int        `json:"logLevel"`
	Lumberjack Lumberjack `json:"lumberjack"`
}

type Lumberjack struct {
	LogFile   string `json:"logFile"`
	MaxSize   int    `json:"maxSize"`
	MaxAge    int    `json:"maxAge"`
	MaxBackup int    `json:"maxBackup"`
	LocalTime bool   `json:"localTime"`
	Compress  bool   `json:"compress"`
}
