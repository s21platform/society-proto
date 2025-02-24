# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [society.proto](#society-proto)
    - [EmptySociety](#-EmptySociety)
    - [GetSocietyForUserWithOffsetIn](#-GetSocietyForUserWithOffsetIn)
    - [GetSocietyForUserWithOffsetOut](#-GetSocietyForUserWithOffsetOut)
    - [GetSocietyInfoIn](#-GetSocietyInfoIn)
    - [GetSocietyInfoOut](#-GetSocietyInfoOut)
    - [GetSocietyWithOffsetIn](#-GetSocietyWithOffsetIn)
    - [GetSocietyWithOffsetOut](#-GetSocietyWithOffsetOut)
    - [RemoveSocietyIn](#-RemoveSocietyIn)
    - [SetSocietyIn](#-SetSocietyIn)
    - [SetSocietyOut](#-SetSocietyOut)
    - [Society](#-Society)
    - [SubscribeToSocietyIn](#-SubscribeToSocietyIn)
    - [TagsID](#-TagsID)
    - [UnSubscribeToSocietyIn](#-UnSubscribeToSocietyIn)
    - [UpdateSocietyIn](#-UpdateSocietyIn)
  
    - [SocietyService](#-SocietyService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="society-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## society.proto



<a name="-EmptySociety"></a>

### EmptySociety







<a name="-GetSocietyForUserWithOffsetIn"></a>

### GetSocietyForUserWithOffsetIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | Лимит возвращаемых записей |
| total | [int64](#int64) |  | Тотал записи с которой возвращаем |
| userUUID | [string](#string) |  | UUID пользователя |






<a name="-GetSocietyForUserWithOffsetOut"></a>

### GetSocietyForUserWithOffsetOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societies | [Society](#Society) | repeated | Список сообществ |
| total | [int64](#int64) |  | Количество найденых записей |






<a name="-GetSocietyInfoIn"></a>

### GetSocietyInfoIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID сообщества |






<a name="-GetSocietyInfoOut"></a>

### GetSocietyInfoOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Название сообщества |
| description | [string](#string) |  | Описание сообщества |
| ownerUUID | [string](#string) |  | Владелец сообщества |
| photoURL | [string](#string) |  | Аватар сообщества |
| formatID | [int64](#int64) |  | Формат сообщества |
| postPermission | [int64](#int64) |  | Правила сообщений в сообществе |
| isSearch | [bool](#bool) |  | Настройка поиска сообщества |
| countSubscribe | [int64](#int64) |  | Количество подписчиков сообщества |
| tagsID | [TagsID](#TagsID) | repeated | Теги сообщества |






<a name="-GetSocietyWithOffsetIn"></a>

### GetSocietyWithOffsetIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | Лимит возвращаемых записей |
| total | [int64](#int64) |  | Тотал записи с которой возвращаем |
| name | [string](#string) |  | Название сообщества по которому происходит поиск, если пусто то все сообщества |






<a name="-GetSocietyWithOffsetOut"></a>

### GetSocietyWithOffsetOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societies | [Society](#Society) | repeated | Список сообществ |
| total | [int64](#int64) |  | Количество найденых записей |






<a name="-RemoveSocietyIn"></a>

### RemoveSocietyIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID сообщества |






<a name="-SetSocietyIn"></a>

### SetSocietyIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Название сообщества |
| formatID | [int64](#int64) |  | Формат сообщества |
| postPermissionID | [int64](#int64) |  | Правила сообщений в сообществе |
| isSearch | [bool](#bool) |  | Настройка поиска сообщества |






<a name="-SetSocietyOut"></a>

### SetSocietyOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID созданного сообщества |






<a name="-Society"></a>

### Society



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID сообщества |
| name | [string](#string) |  | Название сообщества |
| photoURL | [string](#string) |  | Аватар сообщества |
| isMember | [bool](#bool) |  | Состоит пользователь в сообществе: true - cocтоит, false - не состоит |
| isPrivate | [bool](#bool) |  | Приватное сообщество: true - да, false - нет |






<a name="-SubscribeToSocietyIn"></a>

### SubscribeToSocietyIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID сообщества |






<a name="-TagsID"></a>

### TagsID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tagID | [int64](#int64) |  | ID тега |






<a name="-UnSubscribeToSocietyIn"></a>

### UnSubscribeToSocietyIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID сообщества |






<a name="-UpdateSocietyIn"></a>

### UpdateSocietyIn

Данные, для обновления сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Id-society |
| name | [string](#string) |  | Название сообщества |
| description | [string](#string) |  | Описание сообщества |
| is_private | [bool](#bool) |  | Приватность сообщества |
| direction_id | [int64](#int64) |  | Направление |
| access_level_id | [int64](#int64) |  | Уровень доступа |




| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| societyUUID | [string](#string) |  | UUID сообщества |
| name | [string](#string) |  | Название сообщества |
| description | [string](#string) |  | Описание сообщества |
| photoURL | [string](#string) |  | Аватар сообщества |
| formatID | [int64](#int64) |  | Формат сообщества |
| postPermission | [int64](#int64) |  | Правила сообщений в сообществе |
| isSearch | [bool](#bool) |  | Настройка поиска сообщества |
| tagsID | [TagsID](#TagsID) | repeated | Список тегов сообщества |





 

 

 


<a name="-SocietyService"></a>

### SocietyService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateSociety | [.SetSocietyIn](#SetSocietyIn) | [.SetSocietyOut](#SetSocietyOut) | Создание сообщества |
| UpdateSociety | [.UpdateSocietyIn](#UpdateSocietyIn) | [.EmptySociety](#EmptySociety) | Настройка сообщества |
| GetSocietyInfo | [.GetSocietyInfoIn](#GetSocietyInfoIn) | [.GetSocietyInfoOut](#GetSocietyInfoOut) | Информация по сообществу |
| RemoveSociety | [.RemoveSocietyIn](#RemoveSocietyIn) | [.EmptySociety](#EmptySociety) | Удаление сообщества |
| SubscribeToSociety | [.SubscribeToSocietyIn](#SubscribeToSocietyIn) | [.EmptySociety](#EmptySociety) | Подписка на сообщество |
| UnSubscribeToSociety | [.UnSubscribeToSocietyIn](#UnSubscribeToSocietyIn) | [.EmptySociety](#EmptySociety) | Отписка от сообщества |
| GetSocietyWithOffset | [.GetSocietyWithOffsetIn](#GetSocietyWithOffsetIn) | [.GetSocietyWithOffsetOut](#GetSocietyWithOffsetOut) | Список сообществ для поиска по названию сообщества |
| GetSocietyForUserWithOffset | [.GetSocietyForUserWithOffsetIn](#GetSocietyForUserWithOffsetIn) | [.GetSocietyForUserWithOffsetOut](#GetSocietyForUserWithOffsetOut) | Список сообществ юзера |

 



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

