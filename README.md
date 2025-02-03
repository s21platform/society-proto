# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [society.proto](#society-proto)
    - [AccessLevel](#-AccessLevel)
    - [EmptySociety](#-EmptySociety)
    - [GetAccessLevelOut](#-GetAccessLevelOut)
    - [GetPermissionsOut](#-GetPermissionsOut)
    - [GetSocietiesForUserIn](#-GetSocietiesForUserIn)
    - [GetSocietiesForUserOut](#-GetSocietiesForUserOut)
    - [GetSocietyInfoIn](#-GetSocietyInfoIn)
    - [GetSocietyInfoOut](#-GetSocietyInfoOut)
    - [GetSocietyWithOffsetIn](#-GetSocietyWithOffsetIn)
    - [GetSocietyWithOffsetOut](#-GetSocietyWithOffsetOut)
    - [GetUsersForSocietyIn](#-GetUsersForSocietyIn)
    - [GetUsersForSocietyOut](#-GetUsersForSocietyOut)
    - [Permission](#-Permission)
    - [SetSocietyIn](#-SetSocietyIn)
    - [SetSocietyOut](#-SetSocietyOut)
    - [Society](#-Society)
    - [SubscribeToSocietyIn](#-SubscribeToSocietyIn)
    - [SubscribeToSocietyOut](#-SubscribeToSocietyOut)
    - [UnsubscribeFromSocietyIn](#-UnsubscribeFromSocietyIn)
    - [UnsubscribeFromSocietyOut](#-UnsubscribeFromSocietyOut)
    - [UpdateSocietyIn](#-UpdateSocietyIn)
    - [UserSociety](#-UserSociety)
  
    - [SocietyService](#-SocietyService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="society-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## society.proto



<a name="-AccessLevel"></a>

### AccessLevel
Описание уровней доступа


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Идентификатор уровня доступа |
| access_level | [string](#string) |  | Название или описание уровня доступа |






<a name="-EmptySociety"></a>

### EmptySociety
Пустой объект






<a name="-GetAccessLevelOut"></a>

### GetAccessLevelOut
Объект возвращения уровней доступа


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| levels | [AccessLevel](#AccessLevel) | repeated | Список уровней доступа |






<a name="-GetPermissionsOut"></a>

### GetPermissionsOut
Объект возвращения списка разрешений


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissions | [Permission](#Permission) | repeated | Список разрешений. |






<a name="-GetSocietiesForUserIn"></a>

### GetSocietiesForUserIn
Данные о подписках юзера


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_uuid | [string](#string) |  | UUID юзера |






<a name="-GetSocietiesForUserOut"></a>

### GetSocietiesForUserOut
Возвращение данных о подписках юзера


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society | [Society](#Society) | repeated | Список сообществ, на которые подписан юзер |






<a name="-GetSocietyInfoIn"></a>

### GetSocietyInfoIn
Данные параметра для получения сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | id необходимого сообщества |






<a name="-GetSocietyInfoOut"></a>

### GetSocietyInfoOut
Данные, возвращаемые о информации сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Наименование сообщества |
| description | [string](#string) |  | Описание сообщества |
| ownerUUID | [string](#string) |  | UUID владельца сообщества |
| photoUrl | [string](#string) |  | URL фотографии сообщества |
| isPrivate | [bool](#bool) |  | Приватность сообщества |
| countSubscribers | [int64](#int64) |  | Кол-во подписчиков сообщества |






<a name="-GetSocietyWithOffsetIn"></a>

### GetSocietyWithOffsetIn
Данные, для получения списка сообществ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | Количество возвращаемых записей |
| offset | [int64](#int64) |  | С какой записи начинаем |
| name | [string](#string) |  | Если поле пустое - выводит все группы, иначе поиск по подстроке подходящие |






<a name="-GetSocietyWithOffsetOut"></a>

### GetSocietyWithOffsetOut
Данные, возвращаемые при получении списка сообществ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society | [Society](#Society) | repeated | Список сообществ, на которые подписан юзер |
| total | [int64](#int64) |  | Общее кол-во возвращенных сообществ |






<a name="-GetUsersForSocietyIn"></a>

### GetUsersForSocietyIn
Данные о подписчиках сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society_id | [int64](#int64) |  | ID сообщества |






<a name="-GetUsersForSocietyOut"></a>

### GetUsersForSocietyOut
Возвращение данных о подписчиках сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserSociety](#UserSociety) | repeated | Список юзеров, подписанных на сообщество |






<a name="-Permission"></a>

### Permission
Описание разрешения


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Уникальный идентификатор разрешения |
| name | [string](#string) |  | Название разрешения |
| description | [string](#string) |  | Описание разрешения |






<a name="-SetSocietyIn"></a>

### SetSocietyIn
Данные для создания сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Имя |
| description | [string](#string) |  | Описание |
| is_private | [bool](#bool) |  | Приватность |
| direction_id | [int64](#int64) |  | Направление |
| access_level_id | [int64](#int64) |  | Уровень доступа |






<a name="-SetSocietyOut"></a>

### SetSocietyOut
Объект ответа создания сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society_id | [int64](#int64) |  | id созданного сообщества |






<a name="-Society"></a>

### Society
Данные параметров о подписках юзера


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Название сообщества |
| avatar_link | [string](#string) |  | Ссылка на аватарку сообщества |
| society_id | [int64](#int64) |  | ID сообщества |
| isMember | [bool](#bool) |  | Состоит ли пользователь в группе: true - cocтоит, false - не состоит |
| isPrivate | [bool](#bool) |  | Приватное сообщество: true - да, false - нет |






<a name="-SubscribeToSocietyIn"></a>

### SubscribeToSocietyIn
Данные для подписки на сообщество


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society_id | [int64](#int64) |  | ID сообщества |






<a name="-SubscribeToSocietyOut"></a>

### SubscribeToSocietyOut
Возвращение данных о подписке на сообщество


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Получилось ли подписаться или нет |






<a name="-UnsubscribeFromSocietyIn"></a>

### UnsubscribeFromSocietyIn
Данные для отписки от сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society_id | [int64](#int64) |  | ID сообщества |






<a name="-UnsubscribeFromSocietyOut"></a>

### UnsubscribeFromSocietyOut
Возвращение данных об отписке от сообщества


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Получилось ли подписаться или нет |






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






<a name="-UserSociety"></a>

### UserSociety
Список параметров юзеров, подписанных на сообщество


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Имя пользователя |
| avatar_link | [string](#string) |  | Ссылка на аватарку пользователя |
| user_uuid | [string](#string) |  | UUID юзера |





 

 

 


<a name="-SocietyService"></a>

### SocietyService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateSociety | [.SetSocietyIn](#SetSocietyIn) | [.SetSocietyOut](#SetSocietyOut) | Метод создания сообщества |
| GetAccessLevel | [.EmptySociety](#EmptySociety) | [.GetAccessLevelOut](#GetAccessLevelOut) | Метод получения уровней доступа |
| GetPermissions | [.EmptySociety](#EmptySociety) | [.GetPermissionsOut](#GetPermissionsOut) | Метод получения разрешения |
| SubscribeToSociety | [.SubscribeToSocietyIn](#SubscribeToSocietyIn) | [.SubscribeToSocietyOut](#SubscribeToSocietyOut) | Подписка на сообщество |
| UnsubscribeFromSociety | [.UnsubscribeFromSocietyIn](#UnsubscribeFromSocietyIn) | [.UnsubscribeFromSocietyOut](#UnsubscribeFromSocietyOut) | Отписка от сообщества |
| GetUsersForSociety | [.GetUsersForSocietyIn](#GetUsersForSocietyIn) | [.GetUsersForSocietyOut](#GetUsersForSocietyOut) | Список юзеров, подписанных на сообщество |
| GetSocietiesForUser | [.GetSocietiesForUserIn](#GetSocietiesForUserIn) | [.GetSocietiesForUserOut](#GetSocietiesForUserOut) | Список сообществ, на которые подписан юзер |
| GetSocietyInfo | [.GetSocietyInfoIn](#GetSocietyInfoIn) | [.GetSocietyInfoOut](#GetSocietyInfoOut) | Метод для получения информации о сообществе |
| GetSocietyWithOffset | [.GetSocietyWithOffsetIn](#GetSocietyWithOffsetIn) | [.GetSocietyWithOffsetOut](#GetSocietyWithOffsetOut) | Метод для получения списка сообществ |
| UpdateSociety | [.UpdateSocietyIn](#UpdateSocietyIn) | [.EmptySociety](#EmptySociety) | Метод обновления сообщества |

 



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

