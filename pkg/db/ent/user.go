// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"salt,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// EmailAddress holds the value of the "email_address" field.
	EmailAddress string `json:"email_address,omitempty"`
	// LoginTimes holds the value of the "login_times" field.
	LoginTimes uint32 `json:"login_times,omitempty"`
	// KycVerify holds the value of the "kyc_verify" field.
	KycVerify bool `json:"kyc_verify,omitempty"`
	// GaVerify holds the value of the "ga_verify" field.
	GaVerify bool `json:"ga_verify,omitempty"`
	// SignupMethod holds the value of the "signup_method" field.
	SignupMethod string `json:"signup_method,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// Age holds the value of the "age" field.
	Age uint32 `json:"age,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday string `json:"birthday,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Career holds the value of the "career" field.
	Career string `json:"career,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldKycVerify, user.FieldGaVerify:
			values[i] = new(sql.NullBool)
		case user.FieldLoginTimes, user.FieldCreateAt, user.FieldUpdateAt, user.FieldDeleteAt, user.FieldAge:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldPassword, user.FieldSalt, user.FieldDisplayName, user.FieldPhoneNumber, user.FieldEmailAddress, user.FieldSignupMethod, user.FieldAvatar, user.FieldRegion, user.FieldGender, user.FieldBirthday, user.FieldCountry, user.FieldProvince, user.FieldCity, user.FieldCareer:
			values[i] = new(sql.NullString)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				u.PhoneNumber = value.String
			}
		case user.FieldEmailAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_address", values[i])
			} else if value.Valid {
				u.EmailAddress = value.String
			}
		case user.FieldLoginTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field login_times", values[i])
			} else if value.Valid {
				u.LoginTimes = uint32(value.Int64)
			}
		case user.FieldKycVerify:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field kyc_verify", values[i])
			} else if value.Valid {
				u.KycVerify = value.Bool
			}
		case user.FieldGaVerify:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ga_verify", values[i])
			} else if value.Valid {
				u.GaVerify = value.Bool
			}
		case user.FieldSignupMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field signup_method", values[i])
			} else if value.Valid {
				u.SignupMethod = value.String
			}
		case user.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				u.CreateAt = uint32(value.Int64)
			}
		case user.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				u.UpdateAt = uint32(value.Int64)
			}
		case user.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				u.DeleteAt = uint32(value.Int64)
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				u.Region = value.String
			}
		case user.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				u.Age = uint32(value.Int64)
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = value.String
			}
		case user.FieldBirthday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				u.Birthday = value.String
			}
		case user.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				u.Country = value.String
			}
		case user.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				u.Province = value.String
			}
		case user.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				u.City = value.String
			}
		case user.FieldCareer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field career", values[i])
			} else if value.Valid {
				u.Career = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", phone_number=")
	builder.WriteString(u.PhoneNumber)
	builder.WriteString(", email_address=")
	builder.WriteString(u.EmailAddress)
	builder.WriteString(", login_times=")
	builder.WriteString(fmt.Sprintf("%v", u.LoginTimes))
	builder.WriteString(", kyc_verify=")
	builder.WriteString(fmt.Sprintf("%v", u.KycVerify))
	builder.WriteString(", ga_verify=")
	builder.WriteString(fmt.Sprintf("%v", u.GaVerify))
	builder.WriteString(", signup_method=")
	builder.WriteString(u.SignupMethod)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", u.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", u.DeleteAt))
	builder.WriteString(", avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", region=")
	builder.WriteString(u.Region)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", gender=")
	builder.WriteString(u.Gender)
	builder.WriteString(", birthday=")
	builder.WriteString(u.Birthday)
	builder.WriteString(", country=")
	builder.WriteString(u.Country)
	builder.WriteString(", province=")
	builder.WriteString(u.Province)
	builder.WriteString(", city=")
	builder.WriteString(u.City)
	builder.WriteString(", career=")
	builder.WriteString(u.Career)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
