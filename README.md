# user-balance-service

## API METHODS
[POST] - add money to users account
> INPUT-1 : user-id  
> INPUT-2 : money-amount

[POST] - remove money from account
> INPUT-1 : user-id  
> INPUT-2 : money-amount  

[POST] - transfer money from user to user
> INPUT-1 : user-id-sender  
> INPUT-2 : user-id-sender-receiver  
> INPUT-3 : money-amount  

[GET] - get money on user account in exact currency.  
> INPUT-1 : user-id  
> INPUT-2 : currency    
?userid=    
?currency=USD   
RUB - default   
    
[GET] - get transactions listing on users account with data : where when whom . sort by date\amount + pagination
> INPUT-1 : user-id  
    ?userid=


## 2 ACTORS OUTSIDE
- BILLING SERVICE
- FAVOR SERVICE


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
>go mod init github.com/truecoder34/user-balance-service<br>
>go get github.com/badoux/checkmail<br>
>go get github.com/jinzhu/gorm<br>
>go get golang.org/x/crypto/bcrypt<br>
>go get github.com/dgrijalva/jwt-go<br>
>go get github.com/gorilla/mux<br>
>go get github.com/jinzhu/gorm/dialects/mysql" //If using mysql<br>
>go get github.com/jinzhu/gorm/dialects/postgres //If using postgres<br>
>go get github.com/joho/godotenv<br>
>go get gopkg.in/go-playground/assert.v1<br>
>go get github.com/satori/go.uuid<br>
>go get github.com/badoux/checkmail<br>
>go get github.com/gin-gonic/gin<br>