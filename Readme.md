# GYM STREAK
[**WIP**] - An application for tracking workout!

## Tech Stack

| Category            | List        | github-link                                 | links | go command                                    |
|---------------------|-------------|---------------------------------------------|-------|-----------------------------------------------|
| router              | gorilla-mux | https://github.com/gorilla/mux              | -     | go get -u github.com/gorilla/mux              |
| router              | http-router | https://github.com/julienschmidt/httprouter | -     | go get -u github.com/julienschmidt/httprouter |
| logging             | logrus      | https://github.com/Sirupsen/logrus          | -     | go get -u github.com/sirupsen/logrus          |
| config              | spf13-viper | https://github.com/spf13/viper              | -     | go get github.com/spf13/viper                 |
| log-rotation        | lumberjack  | https://github.com/natefinch/lumberjack     | -     |                                               |
| database-driver     | mysql       | https://github.com/go-sql-driver/mysql      | -     | go get github.com/go-sql-driver/mysql         |
| database-migrations | migrate     | https://github.com/golang-migrate/migrate   | -     | go get github.com/go-sql-driver/mysql         |



## Tech Debts
- [x] Use Viper to pickup value from env variable 
- [x] Pickup logger-settings from Viper Config
- [ ] environment based configuration file
- [ ] Add middleware to log http request
- [ ] Inject Correlation ID to logs
- [ ] Improve Exception Handling logic
- [ ] Replace gorilla-mux with GIN
- [ ] Replace go-sql-driver with GORM
- [ ] Standardize logging
- [ ] Add context support
- [ ] Use Pointer Receiver at required places

