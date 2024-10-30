# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/grpc/protobuf/books_v1/books.proto](#api_grpc_protobuf_books_v1_books-proto)
    - [AddAuthorRequest](#books_v1-service-AddAuthorRequest)
    - [AddBookRequest](#books_v1-service-AddBookRequest)
    - [Author](#books_v1-service-Author)
    - [AuthorID](#books_v1-service-AuthorID)
    - [AuthorResponse](#books_v1-service-AuthorResponse)
    - [AuthorWithBooks](#books_v1-service-AuthorWithBooks)
    - [Book](#books_v1-service-Book)
    - [BookWithAuthor](#books_v1-service-BookWithAuthor)
    - [BooksResponse](#books_v1-service-BooksResponse)
    - [Empty](#books_v1-service-Empty)
    - [Response](#books_v1-service-Response)
  
    - [BooksService](#books_v1-service-BooksService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_grpc_protobuf_books_v1_books-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/grpc/protobuf/books_v1/books.proto
Сервис CRUD операций с клиентами


<a name="books_v1-service-AddAuthorRequest"></a>

### AddAuthorRequest
Запрос на добавление автора


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Имя |
| surname | [string](#string) |  | Фамилия |
| patronymic | [string](#string) |  | Отчество |






<a name="books_v1-service-AddBookRequest"></a>

### AddBookRequest
Запрос на добавление книги


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | Название |
| authorId | [int64](#int64) |  | ID Автора |






<a name="books_v1-service-Author"></a>

### Author
Структура автора


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID Автора |
| fullName | [string](#string) |  | Полное имя автора |






<a name="books_v1-service-AuthorID"></a>

### AuthorID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID Автора |






<a name="books_v1-service-AuthorResponse"></a>

### AuthorResponse
Список авторов с их книгами


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authors | [AuthorWithBooks](#books_v1-service-AuthorWithBooks) | repeated | Список авторов |






<a name="books_v1-service-AuthorWithBooks"></a>

### AuthorWithBooks
Структура автора с книгами


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID автора |
| fullName | [string](#string) |  | ФИО автора |
| books | [Book](#books_v1-service-Book) | repeated | Список книг автора |






<a name="books_v1-service-Book"></a>

### Book
Структура книг


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID Книги |
| title | [string](#string) |  | Название книги |
| authorId | [int64](#int64) |  | Автор книги |






<a name="books_v1-service-BookWithAuthor"></a>

### BookWithAuthor
Структура книг с автором


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID Книги |
| title | [string](#string) |  | Название книги |
| author | [Author](#books_v1-service-Author) |  | Автор книги |






<a name="books_v1-service-BooksResponse"></a>

### BooksResponse
Ответ с книгами


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| books | [BookWithAuthor](#books_v1-service-BookWithAuthor) | repeated | Список книг |






<a name="books_v1-service-Empty"></a>

### Empty







<a name="books_v1-service-Response"></a>

### Response
Ответ на добавление книги или автора


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |





 

 

 


<a name="books_v1-service-BooksService"></a>

### BooksService
Сервис CRUD операций с книгами

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddBook | [AddBookRequest](#books_v1-service-AddBookRequest) | [Response](#books_v1-service-Response) |  |
| AddAuthor | [AddAuthorRequest](#books_v1-service-AddAuthorRequest) | [Response](#books_v1-service-Response) |  |
| GetAllBookWithAuthor | [Empty](#books_v1-service-Empty) | [BooksResponse](#books_v1-service-BooksResponse) |  |
| GetBookWithAuthor | [AuthorID](#books_v1-service-AuthorID) | [BookWithAuthor](#books_v1-service-BookWithAuthor) |  |
| GetAllAuthorWithBooks | [Empty](#books_v1-service-Empty) | [AuthorResponse](#books_v1-service-AuthorResponse) |  |

 



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

