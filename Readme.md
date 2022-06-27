# GYM STREAK
[**WIP**] - An application for tracking workout!

## Tech Stack

| Category | List        | github-link                                 | links | go command                                    |
|----------|-------------|---------------------------------------------|-------|-----------------------------------------------|
| router   | gorilla-mux | https://github.com/gorilla/mux              | -     | go get -u github.com/gorilla/mux              |
| router   | http-router | https://github.com/julienschmidt/httprouter | -     | go get -u github.com/julienschmidt/httprouter |
| logging  | logrus      | https://github.com/Sirupsen/logrus          | -     | go get -u github.com/sirupsen/logrus          |
|          |             |                                             |       |                                               |


## Tech Debts
- [ ] Use Viper to pickup value from env variable 
- [ ] Pickup logger-settings from Viper Config
- [ ] Add middleware to log http request
- [ ] Inject Correlation ID to logs

