# Aura-Server

##  [Link to Postman](https://www.postman.com/wiaderek/workspace/aurahub)

## Packages to install:
#### 1. Gin Web Framework 
```bash
go get -u github.com/gin-gonic/gin
```
#### 2. MongoDB 
```bash
go get go.mongodb.org/mongo-driver/mongo
```
#### 3. Compile Daemon
```bash
go get github.com/githubnemo/CompileDaemon
```
#### 4. GoDotEnv
```bash
go get github.com/joho/godotenv
```
#### 5. Google UUID
```bash
go get github.com/google/uuid
```
#### 6. JWT v5
```bash
go get -u github.com/golang-jwt/jwt/v5
```
#### 7. Crypto - Bcrypt
```bash
go get -u golang.org/x/crypto/bcrypt
```
#### 8. Eclipse Paho MQTT Go client
```bash
go get github.com/eclipse/paho.mqtt.golang

go get github.com/gorilla/websocket
go get golang.org/x/net/proxy
```

---

# Available attributes
## Devices
* OnOff - (0; 255) - "bool"
* Brightness - (true/false) - "value"

## Triggers
* Button
  * Click
  * DoubleClick
  * Hold
* Movement sensor
* Switch

### Actions
* Set value - "set"
* Switch - "switch"

## Sensors
