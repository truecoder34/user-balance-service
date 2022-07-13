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


[GET] get all users
```
    curl --location --request GET 'http://localhost:8080/users'
```
<br>

[GET] get user data by uuid
```
    curl --location --request GET 'http://localhost:8080/users/<USER-UUID>'
```
<br>

```
    curl --location --request GET 'http://localhost:8080/users/236a9d24-ed23-4533-82a1-5a92d92efa97'
```
<br>

[POST] create user
```
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Timothée1",
    "surname": "Chalame1",
    "nickname": "SweetBoy1",
    "email": "tc-sweet@gmail.com1",
    "phone_number": "+144712312341"
}'
```

[GET] get all accounts data
```
    curl --location --request GET 'http://localhost:8080/accounts'
```

[GET] get account data by USER ID 
```
curl --location --request GET 'http://localhost:8080/accounts'
```

[POST] create account
```
    curl --location --request POST 'http://localhost:8080/accounts/2c294273-e10e-41f1-947a-ff4de0baad77'
```

[POST] - transfer data between accounts 
```
curl --location --request POST 'http://localhost:8080/money-transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id_sender": "992e82c7-aefc-4761-bb4d-f75cf5f1a1a7",
    "user_id_receiver": "2c294273-e10e-41f1-947a-ff4de0baad77",
    "money_amount": 4770
}'
```

[POST] add/remove data from account by user id
to Remove money
```
curl --location --request POST 'http://localhost:8080/money' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "abcf371f-e9c0-4ce7-94a2-8b74eb6432f4",
    "money_amount": 1000,
    "action": true
}'
```
or to Add money
```
curl --location --request POST 'http://localhost:8080/money' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "4f09a456-92af-41d8-97a1-efbbd8a578fc",
    "money_amount": 200000,
    "action": true
}'
```

[GET] get account balance by user id
```
curl --location --request GET 'http://localhost:8080/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "4f09a456-92af-41d8-97a1-efbbd8a578fc",
    "currency": "USD"
}'
```


## 2 ACTORS OUTSIDE
- BILLING SERVICE
- FAVOR SERVICE


## ENTITIES
### USER
```
id
date-create
date-update
date-delete
name
surname
nickname
email
phone_number
```

### ACCOUNT [1-1USER]
```
id
date-create
date-update
date-delete
user-id
money-amount
comment
```
![alt text](https://www.linkpicture.com/q/DB_1.png)



### GO MODULES 
>go mod init github.com/truecoder34/user-balance-service<br>
>go get github.com/badoux/checkmail<br>
>go get gorm.io/gorm<br>
>go get golang.org/x/crypto/bcrypt<br>
>go get github.com/dgrijalva/jwt-go<br>
>go get github.com/gorilla/mux<br>
>go get gorm.io/driver/postgres/mysql<br>
>go get gorm.io/driver/postgres<br>
>go get github.com/joho/godotenv<br>
>go get gopkg.in/go-playground/assert.v1<br>
>go get github.com/satori/go.uuid<br>
>go get github.com/badoux/checkmail<br>
>go get github.com/gin-gonic/gin<br>