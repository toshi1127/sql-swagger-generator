# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**NickName** | Pointer to **string** |  | 
**ProfileImageUri** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SocialLink** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**IdentifyStatus** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(nickName string, email string, createdAt time.Time, updatedAt time.Time, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *User) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *User) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *User) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *User) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNickName

`func (o *User) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *User) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *User) SetNickName(v string)`

SetNickName sets NickName field to given value.


### GetProfileImageUri

`func (o *User) GetProfileImageUri() string`

GetProfileImageUri returns the ProfileImageUri field if non-nil, zero value otherwise.

### GetProfileImageUriOk

`func (o *User) GetProfileImageUriOk() (*string, bool)`

GetProfileImageUriOk returns a tuple with the ProfileImageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUri

`func (o *User) SetProfileImageUri(v string)`

SetProfileImageUri sets ProfileImageUri field to given value.

### HasProfileImageUri

`func (o *User) HasProfileImageUri() bool`

HasProfileImageUri returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSocialLink

`func (o *User) GetSocialLink() string`

GetSocialLink returns the SocialLink field if non-nil, zero value otherwise.

### GetSocialLinkOk

`func (o *User) GetSocialLinkOk() (*string, bool)`

GetSocialLinkOk returns a tuple with the SocialLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLink

`func (o *User) SetSocialLink(v string)`

SetSocialLink sets SocialLink field to given value.

### HasSocialLink

`func (o *User) HasSocialLink() bool`

HasSocialLink returns a boolean if a field has been set.

### GetGender

`func (o *User) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *User) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *User) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *User) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetIdentifyStatus

`func (o *User) GetIdentifyStatus() string`

GetIdentifyStatus returns the IdentifyStatus field if non-nil, zero value otherwise.

### GetIdentifyStatusOk

`func (o *User) GetIdentifyStatusOk() (*string, bool)`

GetIdentifyStatusOk returns a tuple with the IdentifyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifyStatus

`func (o *User) SetIdentifyStatus(v string)`

SetIdentifyStatus sets IdentifyStatus field to given value.

### HasIdentifyStatus

`func (o *User) HasIdentifyStatus() bool`

HasIdentifyStatus returns a boolean if a field has been set.

### GetCustomerId

`func (o *User) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *User) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *User) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *User) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *User) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *User) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *User) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *User) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


