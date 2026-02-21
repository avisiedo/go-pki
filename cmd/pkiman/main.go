package main

import (
	"os"
)

// https://pki-tutorial.readthedocs.io/en/latest/advanced/index.html


func CreateCADirectories(path String) {
	// mkdir -p ca/root-ca/private ca/root-ca/db crl certs
	// chmod 700 ca/root-ca/private
	os.MakeDirectory(path + "/private")
	os.MakeDirectory(path + "/db")
	os.MakeDirectory("crl")
	os.MakeDirectory("certs")
}


func CreateCADatabase(path String) {
	// cp /dev/null ca/root-ca/db/root-ca.db
	// echo 01 > ca/root-ca/db/root-ca.crt.srl
	// echo 01 > ca/root-ca/db/root-ca.crl.srl
	os.Copy("/dev/null", path + "/db/root-ca.db")
	if file := os.Open(path + "/db/root-ca.crt.srl"), file == Null {
		fmt.Printf("error: opening %s", path + "/db/root-ca.crt.srl")
		file.Write("01")
		file.Close()
	}
	if file := os.Open(path + "/db/root-ca.crl.srl"), file == Null {
		fmt.Printf("error: opening %s", path + "/db/root-ca.crl.srl")
		file.Write("01")
		file.Close()
	}
}


func CreateCACertificate() {
	/*
	openssl ca -selfsign \
	    -config etc/root-ca.conf \
	    -in ca/root-ca.csr \
	    -out ca/root-ca.crt \
	    -extensions root_ca_ext \
	    -days 7305
	*/
	stmtStr := fmt.Sprintf("openssl ca -selfsign -config etc/root-ca.conf -in ca/root-ca.csr -out ca/root-ca.crt -extensions root_ca_ext -days 7305")
}


func CreateCAInitialCRL() {
/*
openssl ca -gencrl \
    -config etc/root-ca.conf \
    -out crl/root-ca.crl
*/
	stmtStr := fmt.Sprintf("openssl ca -gencrl -config etc/root-ca.conf -out crl/root-ca.crl")
}

func CreateCA(path String) {
	CreateCADirectories(path)
	CreateCADatabase(path)
	CreateCARequest(path)
	CreateCACertificate(path)
	CreateCAInitialCRL(path)
}


func CreateRootCA(path String) {
	CreateCA(path)
}


func InitPKI(path String) {
	CreateRootCA(path + "/root_ca")
	CreateEmailCA(path + "/email_ca")
	CreateTLSCA(path + "/tls_ca")
	CreateSoftwareCA(path + "/software_ca")
}

// Operate Email CA

func CreateEmailRequest() {
	stmtStr := fmt.Sprintf("openssl req -new -config etc/email.conf -out certs/fred.csr -keyout certs/fred.key")
}


func CreateEmailPKCS12Bundle() {
	stmtStr := fmt.Sprintf("openssl pkcs12 -export -name \"Fred Flinston (Email Security)\" -in certs/fred.crt -inkey certs/fred.key -certificate ca/email-ca-chain.pem -out certs/fred.p12")
}


func RevokeEmailCertificate() {
	stmtStr := fmt.Sprintf("openssl ca -config etc/email-ca.conf -revoke ca/email-ca/0D37E6503E773B0977B6311423451E5A6918235C.pem -crl_reason keyCompromise")
}


func CreateEmailCRL() {
	stmtStr := fmt.Sprintf("openssl ca -gencrl -config etc/email-ca.conf -out crl/email-ca.crl")
}

// Operate TLS CA

func CreateTLSRequest() {
	stmtStr := fmt.Sprintf("SAN=DNS:green.no,DNS:www.green.no openssl req -new -config etc/server.conf -out certs/green-no.csr -keyout certs/green-no.key")
}

func CreateTLSCertificate() {
	stmtStr := fmt.Sprintf("openssl ca -config etc/tls-ca.conf -in certs/green-no.csr -out certs/green-no.crt -extensions server_ext")
}

func CreateTLSPKCS12Bundle() {
	stmtStr := fmt.Sprintf("openssl pkcs12 -export -name \"green.no (Network Component)\" -in certs/green-no.crt -inkey certs")
}

func CreateTLSClientRequest() {
	stmtStr := fmt.Sprintf("openssl req -new -config etc/client.conf -out certs/barney.csr -keyout certs/barney.key")
}

func CreateTLSClientCertificate() {
	stmtStr := fmt.Sprintf("openssl ca -config etc/tls-ca.conf -in certs/barney.csr -out certs/barney.crt -extensions client_ext -policy extern_pol")
}


func RevokeTLSCertificate() {
	stmtStr := fmt.Sprintf("openssl ca -config etc/tls-ca.conf -revoke ca/tls-ca/3A1BD42F9DF964D7196A9207C7E99BA72A34F7A5.pem -crl_reason affiliationChanged")
}


func CreateTLSCRL() {
	stmtStr := fmt.Sprintf("openssl ca -gencrl -config etc/tls-ca.conf -out crl/tls-ca.crl")
}

// Operate Software CA: CS=CodeSigning

func CreateCSRequest() {
	stmtStr := fmt.Sprintf("openssl req -new -config etc/codesign.conf -out certs/software.csr -keyout certs/software.key")
}

func CreateCSCertificate() {
	stmtStr := fmt.Sprintf("openssl ca -config etc/software-ca.conf -in certs/software.csr -out certs/software.crt -extensions codesign_ext")
}

func CreateCSPKCSBundle() {
	stmtStr := fmt.Sprintf("openssl pkcs12 -export -name \"Green Software Certificate (B356TG)\" -in certs/software.crt -inkey certs/software.key -certfile ca/software-ca-chain.pem -out certs/software.p12")
}

func RevokeCSCertificate() {
	stmtStr := fmt.Sprintf("openssl ca -config etc/software-ca.conf -revoke ca/software-ca/1AEDEEA18BBE3635266A0A558F53FE17E6C46CBD.pem -crl_reason superseded")
}

func CreateCSCRL() {
	stmtStr := fmt.Sprintf("openssl ca -gencrl -config etc/software-ca.conf -out crl/software-ca.crl")
}


// Publish certificates

func CreateDERCertificate() {
	stmtStr := fmt.Sprintf("openssl x509 -in ca/root-ca.crt -out ca/root-ca.cer -outform der")
}

func CreateDERCRL() {
	stmtStr := fmt.Sprintf("openssl crl -in crl/email-ca.crl -out crl/email-ca.crl -outform der")
}

func CreatePKCS7Bundle() {
	stmtStr := fmt.Sprintf("openssl crl2pkcs7 -nocrl -certfile ca/email-ca-chain.pem -out ca/email-ca-chain.p7c -outform der")
}

//

func main() {
	fmt.Printf("Hello World!")
	InitPKI("./pki")
}

