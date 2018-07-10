# Protocol Documentation
<a name="top"/>

## Table of Contents

- [message.proto](#message.proto)
    - [Message](#messaging.Message)
    - [Thread](#messaging.Thread)
  
  
  
  

- [message_service.proto](#message_service.proto)
    - [AuthedMessage](#messaging.AuthedMessage)
    - [ThreadReply](#messaging.ThreadReply)
  
  
  
    - [MessageService](#messaging.MessageService)
  

- [user.proto](#user.proto)
    - [Ident](#messaging.Ident)
    - [Profile](#messaging.Profile)
  
  
  
  

- [user_service.proto](#user_service.proto)
    - [AuthedProfile](#messaging.AuthedProfile)
  
  
  
    - [UserService](#messaging.UserService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="message.proto"/>
<p align="right"><a href="#top">Top</a></p>

## message.proto



<a name="messaging.Message"/>

### Message
Message is a message from one user to another


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the message&#39;s unique ID. |
| from | [Profile](#messaging.Profile) |  | from is the user the message was sent from. |
| to | [Profile](#messaging.Profile) |  | to is the user the message was sent to. |
| title | [string](#string) |  | title is the message&#39;s title. |
| body | [string](#string) |  | body is the message&#39;s body. |






<a name="messaging.Thread"/>

### Thread
Thread is a sequence of messages from one user to another


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the thread&#39;s unique ID. |
| messages | [Message](#messaging.Message) | repeated | messages is the list of messages in the thread. |





 

 

 

 



<a name="message_service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## message_service.proto



<a name="messaging.AuthedMessage"/>

### AuthedMessage
AuthedMessage is a combined authorization and message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [Ident](#messaging.Ident) |  |  |
| message | [Message](#messaging.Message) |  |  |






<a name="messaging.ThreadReply"/>

### ThreadReply
ThreadReply is a message to be added to a thread.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [Ident](#messaging.Ident) |  |  |
| thread | [Thread](#messaging.Thread) |  |  |
| message | [Message](#messaging.Message) |  |  |





 

 

 


<a name="messaging.MessageService"/>

### MessageService
MessageService is our message microservice.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Send | [AuthedMessage](#messaging.AuthedMessage) | [Thread](#messaging.AuthedMessage) | Send sends a new message to a user. |
| Reply | [ThreadReply](#messaging.ThreadReply) | [Thread](#messaging.ThreadReply) | Reply replies to a thread. |
| Delete | [AuthedMessage](#messaging.AuthedMessage) | [Thread](#messaging.AuthedMessage) | Delete deletes a message. |

 



<a name="user.proto"/>
<p align="right"><a href="#top">Top</a></p>

## user.proto
This is important.  proto3 has some major advantages you don&#39;t want to miss out on.


<a name="messaging.Ident"/>

### Ident
Ident is a message used to identify a user.  It should be used when a user needs
to log in or create an account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | username is the unique name of the user. |
| password | [string](#string) |  | password is the user&#39;s password. |






<a name="messaging.Profile"/>

### Profile
Profile is the set of public details about a user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the user&#39;s unique ID. |
| email | [string](#string) |  | email is the user&#39;s email address. |
| display_name | [string](#string) |  | display_name is the user&#39;s name to show to other users. |





 

 

 

 



<a name="user_service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## user_service.proto



<a name="messaging.AuthedProfile"/>

### AuthedProfile
AuthedProfile is a public user profile with authorization attached.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [Ident](#messaging.Ident) |  |  |
| profile | [Profile](#messaging.Profile) |  |  |





 

 

 


<a name="messaging.UserService"/>

### UserService
UserService is our user microservice.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [Ident](#messaging.Ident) | [Profile](#messaging.Ident) | CreateUser creates a user. |
| SetProfile | [AuthedProfile](#messaging.AuthedProfile) | [Profile](#messaging.AuthedProfile) | SetProfile updates a user&#39;s profile. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

