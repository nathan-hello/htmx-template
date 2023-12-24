// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type AuthAalLevel string

const (
	AuthAalLevelAal1 AuthAalLevel = "aal1"
	AuthAalLevelAal2 AuthAalLevel = "aal2"
	AuthAalLevelAal3 AuthAalLevel = "aal3"
)

func (e *AuthAalLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthAalLevel(s)
	case string:
		*e = AuthAalLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthAalLevel: %T", src)
	}
	return nil
}

type NullAuthAalLevel struct {
	AuthAalLevel AuthAalLevel
	Valid        bool // Valid is true if AuthAalLevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthAalLevel) Scan(value interface{}) error {
	if value == nil {
		ns.AuthAalLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthAalLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthAalLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthAalLevel), nil
}

type AuthCodeChallengeMethod string

const (
	AuthCodeChallengeMethodS256  AuthCodeChallengeMethod = "s256"
	AuthCodeChallengeMethodPlain AuthCodeChallengeMethod = "plain"
)

func (e *AuthCodeChallengeMethod) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthCodeChallengeMethod(s)
	case string:
		*e = AuthCodeChallengeMethod(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthCodeChallengeMethod: %T", src)
	}
	return nil
}

type NullAuthCodeChallengeMethod struct {
	AuthCodeChallengeMethod AuthCodeChallengeMethod
	Valid                   bool // Valid is true if AuthCodeChallengeMethod is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthCodeChallengeMethod) Scan(value interface{}) error {
	if value == nil {
		ns.AuthCodeChallengeMethod, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthCodeChallengeMethod.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthCodeChallengeMethod) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthCodeChallengeMethod), nil
}

type AuthFactorStatus string

const (
	AuthFactorStatusUnverified AuthFactorStatus = "unverified"
	AuthFactorStatusVerified   AuthFactorStatus = "verified"
)

func (e *AuthFactorStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthFactorStatus(s)
	case string:
		*e = AuthFactorStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthFactorStatus: %T", src)
	}
	return nil
}

type NullAuthFactorStatus struct {
	AuthFactorStatus AuthFactorStatus
	Valid            bool // Valid is true if AuthFactorStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthFactorStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AuthFactorStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthFactorStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthFactorStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthFactorStatus), nil
}

type AuthFactorType string

const (
	AuthFactorTypeTotp     AuthFactorType = "totp"
	AuthFactorTypeWebauthn AuthFactorType = "webauthn"
)

func (e *AuthFactorType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthFactorType(s)
	case string:
		*e = AuthFactorType(s)
	default:
		return fmt.Errorf("unsupported scan type for AuthFactorType: %T", src)
	}
	return nil
}

type NullAuthFactorType struct {
	AuthFactorType AuthFactorType
	Valid          bool // Valid is true if AuthFactorType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuthFactorType) Scan(value interface{}) error {
	if value == nil {
		ns.AuthFactorType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuthFactorType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuthFactorType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AuthFactorType), nil
}

// Auth: Audit trail for user actions.
type AuthAuditLogEntry struct {
	InstanceID uuid.NullUUID
	ID         uuid.UUID
	Payload    pqtype.NullRawMessage
	CreatedAt  sql.NullTime
	IpAddress  string
}

// stores metadata for pkce logins
type AuthFlowState struct {
	ID                   uuid.UUID
	UserID               uuid.NullUUID
	AuthCode             string
	CodeChallengeMethod  AuthCodeChallengeMethod
	CodeChallenge        string
	ProviderType         string
	ProviderAccessToken  sql.NullString
	ProviderRefreshToken sql.NullString
	CreatedAt            sql.NullTime
	UpdatedAt            sql.NullTime
	AuthenticationMethod string
}

// Auth: Stores identities associated to a user.
type AuthIdentity struct {
	ProviderID   string
	UserID       uuid.UUID
	IdentityData json.RawMessage
	Provider     string
	LastSignInAt sql.NullTime
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	// Auth: Email is a generated column that references the optional email property in the identity_data
	Email sql.NullString
	ID    uuid.UUID
}

// Auth: Manages users across multiple sites.
type AuthInstance struct {
	ID            uuid.UUID
	Uuid          uuid.NullUUID
	RawBaseConfig sql.NullString
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
}

// auth: stores authenticator method reference claims for multi factor authentication
type AuthMfaAmrClaim struct {
	SessionID            uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	AuthenticationMethod string
	ID                   uuid.UUID
}

// auth: stores metadata about challenge requests made
type AuthMfaChallenge struct {
	ID         uuid.UUID
	FactorID   uuid.UUID
	CreatedAt  time.Time
	VerifiedAt sql.NullTime
	IpAddress  pqtype.Inet
}

// auth: stores metadata about factors
type AuthMfaFactor struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	FriendlyName sql.NullString
	FactorType   AuthFactorType
	Status       AuthFactorStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Secret       sql.NullString
}

// Auth: Store of tokens used to refresh JWT tokens once they expire.
type AuthRefreshToken struct {
	InstanceID uuid.NullUUID
	ID         int64
	Token      sql.NullString
	UserID     sql.NullString
	Revoked    sql.NullBool
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	Parent     sql.NullString
	SessionID  uuid.NullUUID
}

// Auth: Manages SAML Identity Provider connections.
type AuthSamlProvider struct {
	ID               uuid.UUID
	SsoProviderID    uuid.UUID
	EntityID         string
	MetadataXml      string
	MetadataUrl      sql.NullString
	AttributeMapping pqtype.NullRawMessage
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
}

// Auth: Contains SAML Relay State information for each Service Provider initiated login.
type AuthSamlRelayState struct {
	ID            uuid.UUID
	SsoProviderID uuid.UUID
	RequestID     string
	ForEmail      sql.NullString
	RedirectTo    sql.NullString
	FromIpAddress pqtype.Inet
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
	FlowStateID   uuid.NullUUID
}

// Auth: Manages updates to the auth system.
type AuthSchemaMigration struct {
	Version string
}

// Auth: Stores session data associated to a user.
type AuthSession struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	FactorID  uuid.NullUUID
	Aal       NullAuthAalLevel
	// Auth: Not after is a nullable column that contains a timestamp after which the session should be regarded as expired.
	NotAfter    sql.NullTime
	RefreshedAt sql.NullTime
	UserAgent   sql.NullString
	Ip          pqtype.Inet
	Tag         sql.NullString
}

// Auth: Manages SSO email address domain mapping to an SSO Identity Provider.
type AuthSsoDomain struct {
	ID            uuid.UUID
	SsoProviderID uuid.UUID
	Domain        string
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
}

// Auth: Manages SSO identity provider information; see saml_providers for SAML.
type AuthSsoProvider struct {
	ID uuid.UUID
	// Auth: Uniquely identifies a SSO provider according to a user-chosen resource ID (case insensitive), useful in infrastructure as code.
	ResourceID sql.NullString
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

// Auth: Stores user login data within a secure schema.
type AuthUser struct {
	InstanceID               uuid.NullUUID
	ID                       uuid.UUID
	Aud                      sql.NullString
	Role                     sql.NullString
	Email                    sql.NullString
	EncryptedPassword        sql.NullString
	EmailConfirmedAt         sql.NullTime
	InvitedAt                sql.NullTime
	ConfirmationToken        sql.NullString
	ConfirmationSentAt       sql.NullTime
	RecoveryToken            sql.NullString
	RecoverySentAt           sql.NullTime
	EmailChangeTokenNew      sql.NullString
	EmailChange              sql.NullString
	EmailChangeSentAt        sql.NullTime
	LastSignInAt             sql.NullTime
	RawAppMetaData           pqtype.NullRawMessage
	RawUserMetaData          pqtype.NullRawMessage
	IsSuperAdmin             sql.NullBool
	CreatedAt                sql.NullTime
	UpdatedAt                sql.NullTime
	Phone                    sql.NullString
	PhoneConfirmedAt         sql.NullTime
	PhoneChange              sql.NullString
	PhoneChangeToken         sql.NullString
	PhoneChangeSentAt        sql.NullTime
	ConfirmedAt              sql.NullTime
	EmailChangeTokenCurrent  sql.NullString
	EmailChangeConfirmStatus sql.NullInt16
	BannedUntil              sql.NullTime
	ReauthenticationToken    sql.NullString
	ReauthenticationSentAt   sql.NullTime
	// Auth: Set this column to true when the account comes from SSO. These accounts can have duplicate emails.
	IsSsoUser bool
	DeletedAt sql.NullTime
}

type Todo struct {
	ID        int64
	CreatedAt time.Time
	Body      string
}