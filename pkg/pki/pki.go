package pki

import (
	"context"
	"crypto"
	"crypto/x509"
	"time"
)

// PublicKeyInfraestructure defines certificate lifecycle operations.
type PublicKeyInfraestructure interface {
	IssuePersonalCertificate(ctx context.Context, req PersonalCertRequest) (*Certificate, error)
	IssueCodeSignCertificate(ctx context.Context, req CodeSignCertRequest) (*Certificate, error)
	IssueHostCertificate(ctx context.Context, req HostCertRequest) (*Certificate, error)
	IssueTLSCertificate(ctx context.Context, req TLSCertRequest) (*Certificate, error)
	RevokeCertificate(ctx context.Context, serial string) error
	ValidateCertificate(ctx context.Context, certPEM []byte) (*ValidationResult, error)
	RenewCertificate(ctx context.Context, serial string) (*Certificate, error)
}

// Certificate represents an issued certificate and some metadata.
type Certificate struct {
	PEM       []byte
	Serial    string
	NotBefore time.Time
	NotAfter  time.Time
	Parsed    *x509.Certificate
}

// Requests for different certificate issuance types.
type PersonalCertRequest struct {
	CommonName string
	SANs       []string
	PublicKey  crypto.PublicKey
	TTL        time.Duration
	Template   *x509.Certificate // optional advanced template
}

type CodeSignCertRequest struct {
	Subject   string
	PublicKey crypto.PublicKey
	TTL       time.Duration
}

type HostCertRequest struct {
	Hostname  string
	PublicKey crypto.PublicKey
	TTL       time.Duration
}

type TLSCertRequest struct {
	Hostname  string
	SANs      []string
	PublicKey crypto.PublicKey
	TTL       time.Duration
}

// ValidationResult describes certificate validation output.
type ValidationResult struct {
	Valid      bool
	Expired    bool
	Revoked    bool
	ChainError error
}