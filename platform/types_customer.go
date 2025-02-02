package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AnonymousCartSignInMode string

const (
	AnonymousCartSignInModeMergeWithExistingCustomerCart AnonymousCartSignInMode = "MergeWithExistingCustomerCart"
	AnonymousCartSignInModeUseAsNewActiveCustomerCart    AnonymousCartSignInMode = "UseAsNewActiveCustomerCart"
)

type AuthenticationMode string

const (
	AuthenticationModePassword     AuthenticationMode = "Password"
	AuthenticationModeExternalAuth AuthenticationMode = "ExternalAuth"
)

type Customer struct {
	// Unique identifier of the Customer.
	ID string `json:"id"`
	// The current version of the customer.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// The customer number can be used to create a more human-readable (in contrast to ID) identifier for the customer.
	// It should be unique across a project.
	// Once the field was set it cannot be changed anymore.
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// The customer's email address and the main identifier of uniqueness for a customer account.
	// Email addresses are either unique to the store they're specified for, _or_ for the entire project.
	// For more information, see Email uniquenes.
	Email string `json:"email"`
	// Only present with the default `authenticationMode`, `Password`.
	Password    *string `json:"password,omitempty"`
	FirstName   *string `json:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	MiddleName  *string `json:"middleName,omitempty"`
	Title       *string `json:"title,omitempty"`
	DateOfBirth *Date   `json:"dateOfBirth,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	VatId       *string `json:"vatId,omitempty"`
	// The addresses have unique IDs in the addresses list
	Addresses []Address `json:"addresses"`
	// The address ID in the addresses list
	DefaultShippingAddressId *string `json:"defaultShippingAddressId,omitempty"`
	// The IDs from the addresses list which are used as shipping addresses
	ShippingAddressIds []string `json:"shippingAddressIds"`
	// The address ID in the addresses list
	DefaultBillingAddressId *string `json:"defaultBillingAddressId,omitempty"`
	// The IDs from the addresses list which are used as billing addresses
	BillingAddressIds []string                `json:"billingAddressIds"`
	IsEmailVerified   bool                    `json:"isEmailVerified"`
	ExternalId        *string                 `json:"externalId,omitempty"`
	CustomerGroup     *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom            *CustomFields           `json:"custom,omitempty"`
	Locale            *string                 `json:"locale,omitempty"`
	Salutation        *string                 `json:"salutation,omitempty"`
	// User-defined unique identifier of the Customer.
	Key *string `json:"key,omitempty"`
	// References to the stores the customer account is associated with.
	// If no stores are specified, the customer is a global customer, and can log in using the Password Flow for global Customers.
	// If one or more stores are specified, the customer can only log in using the Password Flow for Customers in a Store for those specific stores.
	Stores []StoreKeyReference `json:"stores"`
	// Defines whether a Customer has a password.
	AuthenticationMode *AuthenticationMode `json:"authenticationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["shippingAddressIds"] == nil {
		delete(raw, "shippingAddressIds")
	}

	if raw["billingAddressIds"] == nil {
		delete(raw, "billingAddressIds")
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

type CustomerChangePassword struct {
	// Unique identifier of the Customer.
	ID              string `json:"id"`
	Version         int    `json:"version"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type CustomerCreateEmailToken struct {
	// Unique identifier of the email token.
	ID         string `json:"id"`
	Version    *int   `json:"version,omitempty"`
	TtlMinutes int    `json:"ttlMinutes"`
}

type CustomerCreatePasswordResetToken struct {
	Email      string `json:"email"`
	TtlMinutes *int   `json:"ttlMinutes,omitempty"`
}

type CustomerDraft struct {
	// String that uniquely identifies a customer.
	// It can be used to create more human-readable (in contrast to ID) identifier for the customer.
	// It should be **unique** across a project.
	// Once it's set it cannot be changed.
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// The customer's email address and the main identifier of uniqueness for a customer account.
	// Email addresses are either unique to the store they're specified for, _or_ for the entire project, and are case insensitive.
	// For more information, see Email uniquenes.
	Email string `json:"email"`
	// Only optional with `authenticationMode` set to `ExternalAuth`.
	Password   *string `json:"password,omitempty"`
	FirstName  *string `json:"firstName,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
	MiddleName *string `json:"middleName,omitempty"`
	Title      *string `json:"title,omitempty"`
	// Identifies a single cart that will be assigned to the new customer account.
	AnonymousCartId *string `json:"anonymousCartId,omitempty"`
	// Identifies a single cart that will be assigned to the new customer account.
	AnonymousCart *CartResourceIdentifier `json:"anonymousCart,omitempty"`
	// Identifies carts and orders belonging to an anonymous session that will be assigned to the new customer account.
	AnonymousId *string `json:"anonymousId,omitempty"`
	DateOfBirth *Date   `json:"dateOfBirth,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	VatId       *string `json:"vatId,omitempty"`
	// Sets the ID of each address to be unique in the addresses list.
	Addresses []BaseAddress `json:"addresses"`
	// The index of the address in the addresses array.
	// The `defaultShippingAddressId` of the customer will be set to the ID of that address.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// The indices of the shipping addresses in the addresses array.
	// The `shippingAddressIds` of the Customer will be set to the IDs of that addresses.
	ShippingAddresses []int `json:"shippingAddresses"`
	// The index of the address in the addresses array.
	// The `defaultBillingAddressId` of the customer will be set to the ID of that address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The indices of the billing addresses in the addresses array.
	// The `billingAddressIds` of the customer will be set to the IDs of that addresses.
	BillingAddresses []int                            `json:"billingAddresses"`
	IsEmailVerified  *bool                            `json:"isEmailVerified,omitempty"`
	ExternalId       *string                          `json:"externalId,omitempty"`
	CustomerGroup    *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Must be one of the languages supported for this project
	Locale     *string `json:"locale,omitempty"`
	Salutation *string `json:"salutation,omitempty"`
	// User-defined unique identifier for the Customer.
	Key *string `json:"key,omitempty"`
	// References to the stores the customer account is associated with.
	// If no stores are specified, the customer is a global customer, and can log in using the Password Flow for global Customers.
	// If one or more stores are specified, the customer can only log in using the Password Flow for Customers in a Store for those specific stores.
	Stores []StoreResourceIdentifier `json:"stores"`
	// Defines whether a password field is a required field for the Customer.
	AuthenticationMode *AuthenticationMode `json:"authenticationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerDraft) MarshalJSON() ([]byte, error) {
	type Alias CustomerDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addresses"] == nil {
		delete(raw, "addresses")
	}

	if raw["shippingAddresses"] == nil {
		delete(raw, "shippingAddresses")
	}

	if raw["billingAddresses"] == nil {
		delete(raw, "billingAddresses")
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

type CustomerEmailVerify struct {
	Version    *int   `json:"version,omitempty"`
	TokenValue string `json:"tokenValue"`
}

type CustomerPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int        `json:"offset"`
	Results []Customer `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Customer](ctp:api:type:Customer).
*
 */
type CustomerReference struct {
	// Unique identifier of the referenced [Customer](ctp:api:type:Customer).
	ID string `json:"id"`
	// Contains the representation of the expanded Customer. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Customers.
	Obj *Customer `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

type CustomerResetPassword struct {
	TokenValue  string `json:"tokenValue"`
	NewPassword string `json:"newPassword"`
	Version     *int   `json:"version,omitempty"`
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Customer](ctp:api:type:Customer).
*
 */
type CustomerResourceIdentifier struct {
	// Unique identifier of the referenced [Customer](ctp:api:type:Customer). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Customer](ctp:api:type:Customer). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CustomerResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

type CustomerSignInResult struct {
	Customer Customer `json:"customer"`
	// A cart that is associated to the customer.
	// Empty if the customer does not have a cart yet.
	Cart *Cart `json:"cart,omitempty"`
}

type CustomerSignin struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	AnonymousCartId *string `json:"anonymousCartId,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Cart](ctp:api:type:Cart).
	AnonymousCart           *CartResourceIdentifier  `json:"anonymousCart,omitempty"`
	AnonymousCartSignInMode *AnonymousCartSignInMode `json:"anonymousCartSignInMode,omitempty"`
	AnonymousId             *string                  `json:"anonymousId,omitempty"`
	UpdateProductData       *bool                    `json:"updateProductData,omitempty"`
}

type CustomerToken struct {
	// Unique identifier of the CustomerToken.
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	CustomerId     string     `json:"customerId"`
	ExpiresAt      time.Time  `json:"expiresAt"`
	Value          string     `json:"value"`
}

type CustomerUpdate struct {
	Version int                    `json:"version"`
	Actions []CustomerUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCustomerUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type CustomerUpdateAction interface{}

func mapDiscriminatorCustomerUpdateAction(input interface{}) (CustomerUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addAddress":
		obj := CustomerAddAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addBillingAddressId":
		obj := CustomerAddBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShippingAddressId":
		obj := CustomerAddShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addStore":
		obj := CustomerAddStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAddress":
		obj := CustomerChangeAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeEmail":
		obj := CustomerChangeEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAddress":
		obj := CustomerRemoveAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeBillingAddressId":
		obj := CustomerRemoveBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingAddressId":
		obj := CustomerRemoveShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeStore":
		obj := CustomerRemoveStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomField":
		obj := CustomerSetAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomType":
		obj := CustomerSetAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAuthenticationMode":
		obj := CustomerSetAuthenticationModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCompanyName":
		obj := CustomerSetCompanyNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := CustomerSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := CustomerSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerGroup":
		obj := CustomerSetCustomerGroupAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerNumber":
		obj := CustomerSetCustomerNumberAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDateOfBirth":
		obj := CustomerSetDateOfBirthAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultBillingAddress":
		obj := CustomerSetDefaultBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultShippingAddress":
		obj := CustomerSetDefaultShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExternalId":
		obj := CustomerSetExternalIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setFirstName":
		obj := CustomerSetFirstNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := CustomerSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLastName":
		obj := CustomerSetLastNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := CustomerSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMiddleName":
		obj := CustomerSetMiddleNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSalutation":
		obj := CustomerSetSalutationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStores":
		obj := CustomerSetStoresAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTitle":
		obj := CustomerSetTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setVatId":
		obj := CustomerSetVatIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyCustomerChangePassword struct {
	Version         int    `json:"version"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type MyCustomerResetPassword struct {
	TokenValue  string `json:"tokenValue"`
	NewPassword string `json:"newPassword"`
}

type MyCustomerSignin struct {
	Email                string                   `json:"email"`
	Password             string                   `json:"password"`
	ActiveCartSignInMode *AnonymousCartSignInMode `json:"activeCartSignInMode,omitempty"`
	UpdateProductData    *bool                    `json:"updateProductData,omitempty"`
}

type CustomerAddAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

type CustomerAddBillingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerAddShippingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerAddStoreAction struct {
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Store](ctp:api:type:Store).
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddStoreAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStore", Alias: (*Alias)(&obj)})
}

type CustomerChangeAddressAction struct {
	AddressId  *string     `json:"addressId,omitempty"`
	AddressKey *string     `json:"addressKey,omitempty"`
	Address    BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

type CustomerChangeEmailAction struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

type CustomerRemoveAddressAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

type CustomerRemoveBillingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerRemoveShippingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerRemoveStoreAction struct {
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Store](ctp:api:type:Store).
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveStoreAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeStore", Alias: (*Alias)(&obj)})
}

type CustomerSetAddressCustomFieldAction struct {
	AddressId string `json:"addressId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomField", Alias: (*Alias)(&obj)})
}

type CustomerSetAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `address` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `address`.
	Fields    *FieldContainer `json:"fields,omitempty"`
	AddressId string          `json:"addressId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomType", Alias: (*Alias)(&obj)})
}

type CustomerSetAuthenticationModeAction struct {
	AuthMode AuthenticationMode `json:"authMode"`
	// Required when `authMode` is `Password`
	Password *string `json:"password,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetAuthenticationModeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetAuthenticationModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthenticationMode", Alias: (*Alias)(&obj)})
}

type CustomerSetCompanyNameAction struct {
	// If not defined, the company name is unset.
	CompanyName *string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Customer with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Customer.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Customer.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomerGroupAction struct {
	// If not defined, the customer group is unset.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomerNumberAction struct {
	// It should be **unique** across a project.
	// Once it's set, it cannot be changed.
	CustomerNumber *string `json:"customerNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomerNumberAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerNumber", Alias: (*Alias)(&obj)})
}

type CustomerSetDateOfBirthAction struct {
	// If not defined, the date of birth is unset.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

type CustomerSetDefaultBillingAddressAction struct {
	// If not defined, the customer's `defaultBillingAddress` is unset.
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

type CustomerSetDefaultShippingAddressAction struct {
	// If not defined, the customer's `defaultShippingAddress` is unset.
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

type CustomerSetExternalIdAction struct {
	// If not defined, the external ID is unset.
	ExternalId *string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type CustomerSetFirstNameAction struct {
	FirstName *string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

type CustomerSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CustomerSetLastNameAction struct {
	LastName *string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

type CustomerSetLocaleAction struct {
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type CustomerSetMiddleNameAction struct {
	MiddleName *string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

type CustomerSetSalutationAction struct {
	Salutation *string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

type CustomerSetStoresAction struct {
	Stores []StoreResourceIdentifier `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetStoresAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetStoresAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStores", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

type CustomerSetTitleAction struct {
	Title *string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type CustomerSetVatIdAction struct {
	// If not defined, the vat Id is unset.
	VatId *string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetVatIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetVatIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}
