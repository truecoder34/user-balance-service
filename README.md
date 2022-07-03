# user-balance-service

## API METHODS
[POST] - add money to users account
    INPUT: user-id, money-amount
[POST] - remove money from account
    INPUT: user-id, money-amount
[POST] - transfer money from user to user
    INPUT: user-id-sender, user-id-sender-receiver, money-amount
[GET] - get money on user account in exact currency.  
    INPUT-PARAM:  user-id, currency
    ?userid=
    ?currency=USD
    RUB - default
[GET] - get transactions listing on users account with data : where when whom . sort by date\amount + pagination
    ?userid=


## 2 ACTORS OUTSIDE
-BILLING SERVICE
-SERVICIES SERVICE


## ENTITIES
### USER
id
date-create
date-update
name
surname

### ACCOUNT [1-*USER]
user-id
money-amount


### GO MODULES 
go mod init github.com/truecoder34/user-balance-service

go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/mysql" //If using mysql 
go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1

go get github.com/satori/go.uuid
go get github.com/badoux/checkmail