# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/grpc/protobuf/clients_v1/clients.proto](#api_grpc_protobuf_clients_v1_clients-proto)
    - [LoginRequest](#clients_v1-service-LoginRequest)
    - [LoginResponse](#clients_v1-service-LoginResponse)
    - [RegisterRequest](#clients_v1-service-RegisterRequest)
    - [RegisterResponse](#clients_v1-service-RegisterResponse)
    - [UserID](#clients_v1-service-UserID)
    - [Username](#clients_v1-service-Username)
  
    - [ClientsService](#clients_v1-service-ClientsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_grpc_protobuf_clients_v1_clients-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/grpc/protobuf/clients_v1/clients.proto
Сервис CRUD операций с клиентами


<a name="clients_v1-service-LoginRequest"></a>

### LoginRequest
Запрос на аутентификацию


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | Логин |
| password | [string](#string) |  | Пароль |






<a name="clients_v1-service-LoginResponse"></a>

### LoginResponse
Ответ на аутентификацию


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| username | [string](#string) |  | Логин |
| password | [string](#string) |  | Пароль |
| fullname | [string](#string) |  | ФИО |






<a name="clients_v1-service-RegisterRequest"></a>

### RegisterRequest
Запрос на регистрацию


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | Логин |
| password | [string](#string) |  | Пароль |
| name | [string](#string) |  | Имя |
| surname | [string](#string) |  | Фамилия |
| patronymic | [string](#string) |  | Отчество |






<a name="clients_v1-service-RegisterResponse"></a>

### RegisterResponse
Ответ на регистрацию


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="clients_v1-service-UserID"></a>

### UserID
Запрос на получение пользователя по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Логин |






<a name="clients_v1-service-Username"></a>

### Username
Ответ на получение пользователя по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | Логин |





 

 

 


<a name="clients_v1-service-ClientsService"></a>

### ClientsService
Сервис CRUD операций с клиентами

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterRequest](#clients_v1-service-RegisterRequest) | [RegisterResponse](#clients_v1-service-RegisterResponse) |  |
| Login | [LoginRequest](#clients_v1-service-LoginRequest) | [LoginResponse](#clients_v1-service-LoginResponse) |  |
| GetUserByID | [UserID](#clients_v1-service-UserID) | [Username](#clients_v1-service-Username) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

