package acme

const (
	ErrAccountDoesNotExist Error = "The request specified an account that does not exist"
	ErrAlreadyRevoked Error = "The request specified a certificate to revoke that has already been revoked"
	ErrBadCSR Error = "The CSR is unacceptable (e.g., due to a short key)"
	ErrBadNonce Error = "The client sent an unacceptable antireply nonce"
	ErrBadPublicKey Error = "The JWS was signed by a public key the server does not support"
	ErrBadRevocationReason Error = "The revocation reason provided is not allowed by the server"
	ErrBadSignatureAlgorithm Error = "The JWS was signed with an algorithm the server does not support"
	ErrCaa Error = "Certification Authority Authorization (CAA) records forbid the CA from issuing a certificate for the identifier"
	ErrCompound Error = "The server detected multiple problems with the request; the server will attempt to provide more than one error code and more than one error description in the response"
	ErrDns Error = "The identifier is invalid"
	ErrIncorrectResponse Error = "The server could not parse a response from the client"
	ErrInvalidContact Error = "One or more of the contact URIs provided by the client were invalid"
	ErrOrderNotReady Error = "The request specified an order that is not ready to be finalized"
	ErrRateLimited Error = "The request exceeds a rate limit"
	ErrRejectedIdentifier Error = "The server will not issue certificates for the identifier"
	ErrServerInternal Error = "The server experienced an internal error"
	ErrTls Error = "The server received a TLS error during validation"
	ErrUnauthorized Error = "The client lacks sufficient authorization"
	ErrUnsupportedIdentifier Error = "An identifier is of an unsupported type"
	ErrUserActionRequired Error = "Visit the 'instance' URL and take actions specified there"
)

/* Samples: https://www.rfc-editor.org/rfc/rfc8555#page-19

HTTP/1.1 403 Forbidden
Content-Type: application/problem+json
Link: <https://example.com/acme/directory>;rel="index"

{
    "type": "urn:ietf:params:acme:error:malformed",
    "detail": "Some of the identifiers requested were rejected",
    "subproblems": [
        {
            "type": "urn:ietf:params:acme:error:malformed",
            "detail": "Invalid underscore in DNS name \"_example.org\"",
            "identifier": {
                "type": "dns",
                "value": "_example.org"
            }
        },
        {
            "type": "urn:ietf:params:acme:error:rejectedIdentifier",
            "detail": "This CA will not issue for \"example.net\"",
            "identifier": {
                "type": "dns",
                "value": "example.net"
            }
        }
    ]
}

*/

/* https://www.rfc-editor.org/rfc/rfc8555#page-21

   The following table illustrates a typical sequence of requests
   required to establish a new account with the server, prove control of
   an identifier, issue a certificate, and fetch an updated certificate
   some time after issuance.  The "->" is a mnemonic for a Location
   header field pointing to a created resource.

   +-------------------+--------------------------------+--------------+
   | Action            | Request                        | Response     |
   +-------------------+--------------------------------+--------------+
   | Get directory     | GET  directory                 | 200          |
   |                   |                                |              |
   | Get nonce         | HEAD newNonce                  | 200          |
   |                   |                                |              |
   | Create account    | POST newAccount                | 201 ->       |
   |                   |                                | account      |
   |                   |                                |              |
   | Submit order      | POST newOrder                  | 201 -> order |
   |                   |                                |              |
   | Fetch challenges  | POST-as-GET order's            | 200          |
   |                   | authorization urls             |              |
   |                   |                                |              |
   | Respond to        | POST authorization challenge   | 200          |
   | challenges        | urls                           |              |
   |                   |                                |              |
   | Poll for status   | POST-as-GET order              | 200          |
   |                   |                                |              |
   | Finalize order    | POST order's finalize url      | 200          |
   |                   |                                |              |
   | Poll for status   | POST-as-GET order              | 200          |
   |                   |                                |              |
   | Download          | POST-as-GET order's            | 200          |
   | certificate       | certificate url                |              |
   +-------------------+--------------------------------+--------------+
*/

/* Endpoints
HEAD /v1/adme/new-nonce
POST /v1/acme/new-nonce
POST /v1/acme/new-account

POST /v1/acme/new-order
POST /v1/acme/new-authz
POST /v1/acme/revoke-cert
POST /v1/acme/key-change
*/


/*
   HTTP/1.1 200 OK
   Content-Type: application/json

   {
     "newNonce": "https://example.com/acme/new-nonce",
     "newAccount": "https://example.com/acme/new-account",
     "newOrder": "https://example.com/acme/new-order",
     "newAuthz": "https://example.com/acme/new-authz",
     "revokeCert": "https://example.com/acme/revoke-cert",
     "keyChange": "https://example.com/acme/key-change",
     "meta": {
       "termsOfService": "https://example.com/acme/terms/2017-5-30",
       "website": "https://www.example.com/",
       "caaIdentities": ["example.com"],
       "externalAccountRequired": false
     }
   }
*/

/* Account Objects
   {
     "status": "valid",
     "contact": [
       "mailto:cert-admin@example.org",
       "mailto:admin@example.org"
     ],
     "termsOfServiceAgreed": true,
     "orders": "https://example.com/acme/orders/rzGoeA"
   }
*/

/* Order Lists
   HTTP/1.1 200 OK
   Content-Type: application/json
   Link: <https://example.com/acme/directory>;rel="index"
   Link: <https://example.com/acme/orders/rzGoeA?cursor=2>;rel="next"

   {
     "orders": [
       "https://example.com/acme/order/TOlocE8rfgo",
       "https://example.com/acme/order/4E16bbL5iSw",
       // more URLs not shown for example brevity
       "https://example.com/acme/order/neBHYLfw0mg"
     ]
   }
*/

/* Order Objects

   {
     "status": "valid",
     "expires": "2016-01-20T14:09:07.99Z",

     "identifiers": [
       { "type": "dns", "value": "www.example.org" },
       { "type": "dns", "value": "example.org" }
     ],

     "notBefore": "2016-01-01T00:00:00Z",
     "notAfter": "2016-01-08T00:00:00Z",

     "authorizations": [
       "https://example.com/acme/authz/PAniVnsZcis",
       "https://example.com/acme/authz/r4HqLzrSrpI"
     ],

     "finalize": "https://example.com/acme/order/TOlocE8rfgo/finalize",

     "certificate": "https://example.com/acme/cert/mAt3xBGaobw"
   }
*/

/* Challenge

   {
     "status": "valid",
     "expires": "2015-03-01T14:09:07.99Z",

     "identifier": {
       "type": "dns",
       "value": "www.example.org"
     },

     "challenges": [
       {
         "url": "https://example.com/acme/chall/prV_B7yEyA4",
         "type": "http-01",
         "status": "valid",
         "token": "DGyRejmCefe7v4NfDGDKfA",
         "validated": "2014-12-01T12:05:58.16Z"
       }
     ],

     "wildcard": false
   }
*/

/* Getting a nonce

   HEAD /acme/new-nonce HTTP/1.1
   Host: example.com

   HTTP/1.1 200 OK
   Replay-Nonce: oFvnlFP1wIhRlYS2jTaXbA
   Cache-Control: no-store
   Link: <https://example.com/acme/directory>;rel="index"
*/

/* Account Management

   POST /acme/new-account HTTP/1.1
   Host: example.com
   Content-Type: application/jose+json

   {
     "protected": base64url({
       "alg": "ES256",
       "jwk": {...},
       "nonce": "6S8IqOGY7eL2lsGoTZYifg",
       "url": "https://example.com/acme/new-account"
     }),
     "payload": base64url({
       "termsOfServiceAgreed": true,
       "contact": [
         "mailto:cert-admin@example.org",
         "mailto:admin@example.org"
       ]
     }),
     "signature": "RZPOnYoPs1PhjszF...-nh6X1qtOFPB519I"
   }

   HTTP/1.1 201 Created
   Content-Type: application/json
   Replay-Nonce: D8s4D2mLs8Vn-goWuPQeKA
   Link: <https://example.com/acme/directory>;rel="index"
   Location: https://example.com/acme/acct/evOfKhNU60wg

   {
     "status": "valid",

     "contact": [
       "mailto:cert-admin@example.org",
       "mailto:admin@example.org"
     ],

     "orders": "https://example.com/acme/acct/evOfKhNU60wg/orders"
   }
*/

/* Account Update

   POST /acme/acct/evOfKhNU60wg HTTP/1.1
   Host: example.com
   Content-Type: application/jose+json

   {
     "protected": base64url({
       "alg": "ES256",
       "kid": "https://example.com/acme/acct/evOfKhNU60wg",
       "nonce": "ax5RnthDqp_Yf4_HZnFLmA",
       "url": "https://example.com/acme/acct/evOfKhNU60wg"
     }),
     "payload": base64url({
       "contact": [
         "mailto:certificates@example.org",
         "mailto:admin@example.org"
       ]
     }),
     "signature": "hDXzvcj8T6fbFbmn...rDzXzzvzpRy64N0o"
   }
*/

/* External Account Binding

   POST /acme/new-account HTTP/1.1
   Host: example.com
   Content-Type: application/jose+json

   {
     "protected": base64url({
       "alg": "ES256",
       "jwk": // account key,
       "nonce": "K60BWPrMQG9SDxBDS_xtSw",
       "url": "https://example.com/acme/new-account"
     }),
     "payload": base64url({
       "contact": [
         "mailto:cert-admin@example.org",
         "mailto:admin@example.org"
       ],
       "termsOfServiceAgreed": true,

       "externalAccountBinding": {
         "protected": base64url({
           "alg": "HS256",
           "kid": // key identifier from CA,
           "url": "https://example.com/acme/new-account"
         }),
         "payload": base64url(same as in "jwk" above),
         "signature": // MAC using MAC key from CA
       }
     }),
     "signature": "5TWiqIYQfIDfALQv...x9C2mg8JGPxl5bI4"
   }
*/

/* Account Key rollover

   POST /acme/key-change HTTP/1.1
   Host: example.com
   Content-Type: application/jose+json

   {
     "protected": base64url({
       "alg": "ES256",
       "kid": "https://example.com/acme/acct/evOfKhNU60wg",
       "nonce": "S9XaOcxP5McpnTcWPIhYuB",
       "url": "https://example.com/acme/key-change"
     }),
     "payload": base64url({
       "protected": base64url({
         "alg": "ES256",
         "jwk": // new key ,
         "url": "https://example.com/acme/key-change"
       }),
       "payload": base64url({
         "account": "https://example.com/acme/acct/evOfKhNU60wg",
         "oldKey": // old key
       }),
       "signature": "Xe8B94RD30Azj2ea...8BmZIRtcSKPSd8gU"
     }),
     "signature": "5TWiqIYQfIDfALQv...x9C2mg8JGPxl5bI4"
   }
*/

/* Account Deactivation

   POST /acme/acct/evOfKhNU60wg HTTP/1.1
   Host: example.com
   Content-Type: application/jose+json

   {
     "protected": base64url({
       "alg": "ES256",
       "kid": "https://example.com/acme/acct/evOfKhNU60wg",
       "nonce": "ntuJWWSic4WVNSqeUmshgg",
       "url": "https://example.com/acme/acct/evOfKhNU60wg"
     }),
     "payload": base64url({
       "status": "deactivated"
     }),
     "signature": "earzVLd3m5M4xJzR...bVTqn7R08AKOVf3Y"
   }
*/

/* Certificate Issuance

   POST /acme/new-order HTTP/1.1
   Host: example.com
   Content-Type: application/jose+json

   {
     "protected": base64url({
       "alg": "ES256",
       "kid": "https://example.com/acme/acct/evOfKhNU60wg",
       "nonce": "5XJ1L3lEkMG7tR6pA00clA",
       "url": "https://example.com/acme/new-order"
     }),
     "payload": base64url({
       "identifiers": [
         { "type": "dns", "value": "www.example.org" },
         { "type": "dns", "value": "example.org" }
       ],
       "notBefore": "2016-01-01T00:04:00+04:00",
       "notAfter": "2016-01-08T00:04:00+04:00"
     }),
     "signature": "H6ZXtGjTZyUnPeKn...wEA4TklBdh3e454g"
   }

   HTTP/1.1 201 Created
   Replay-Nonce: MYAuvOpaoIiywTezizk5vw
   Link: <https://example.com/acme/directory>;rel="index"
   Location: https://example.com/acme/order/TOlocE8rfgo

   {
     "status": "pending",
     "expires": "2016-01-05T14:09:07.99Z",

     "notBefore": "2016-01-01T00:00:00Z",
     "notAfter": "2016-01-08T00:00:00Z",

     "identifiers": [
       { "type": "dns", "value": "www.example.org" },
       { "type": "dns", "value": "example.org" }
     ],

     "authorizations": [
       "https://example.com/acme/authz/PAniVnsZcis",
       "https://example.com/acme/authz/r4HqLzrSrpI"
     ],

     "finalize": "https://example.com/acme/order/TOlocE8rfgo/finalize"
   }
*/