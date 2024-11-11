# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [society.proto](#society-proto)
    - [AccessLevel](#-AccessLevel)
    - [Empty](#-Empty)
    - [GetAccessLevelOut](#-GetAccessLevelOut)
    - [GetPermissionsOut](#-GetPermissionsOut)
    - [Permission](#-Permission)
    - [SetSocietyIn](#-SetSocietyIn)
    - [SetSocietyOut](#-SetSocietyOut)
  
    - [SocietyService](#-SocietyService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="society-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## society.proto



<a name="-AccessLevel"></a>

### AccessLevel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| access_level | [string](#string) |  |  |






<a name="-Empty"></a>

### Empty







<a name="-GetAccessLevelOut"></a>

### GetAccessLevelOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| levels | [AccessLevel](#AccessLevel) | repeated |  |






<a name="-GetPermissionsOut"></a>

### GetPermissionsOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissions | [Permission](#Permission) | repeated |  |






<a name="-Permission"></a>

### Permission



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="-SetSocietyIn"></a>

### SetSocietyIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| is_private | [bool](#bool) |  |  |
| direction_id | [int64](#int64) |  |  |
| access_level_id | [int64](#int64) |  |  |






<a name="-SetSocietyOut"></a>

### SetSocietyOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| society_id | [int64](#int64) |  |  |





 

 

 


<a name="-SocietyService"></a>

### SocietyService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateSociety | [.SetSocietyIn](#SetSocietyIn) | [.SetSocietyOut](#SetSocietyOut) |  |
| GetAccessLevel | [.Empty](#Empty) | [.GetAccessLevelOut](#GetAccessLevelOut) |  |
| GetPermissions | [.Empty](#Empty) | [.GetPermissionsOut](#GetPermissionsOut) |  |

 



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

