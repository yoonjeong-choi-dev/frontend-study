package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

var (
	host         = flag.String("host", "localhost", "Certificate's comma-separated host name and IPs")
	certFileName = flag.String("cert", "cert.pem", "certificate file name to generate")
	keyFileName  = flag.String("key", "key.pem", "private key file name to generate")
)

func main() {
	flag.Parse()

	// 직접 서명한 인증서 생성을 위한 랜덤 번호 생성
	// => 이후 랜덤한 인증서 생성에 사용
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		panic(err)
	}

	// 인증서 유효 기간 설정을 위한 변수
	// => 생성 시간부터 10년간 유효
	notBefore := time.Now()
	certDuration := 10 * 356 * 24 * time.Hour

	// 인증서 생성 템플릿 정의
	template := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			Organization: []string{"Yoonjeong Choi"},
		},
		NotBefore: notBefore,
		NotAfter:  notBefore.Add(certDuration),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			// 클라이언트 인증서를 인증하기 위해서 필요한 값
			x509.ExtKeyUsageClientAuth,
		},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// 인증서에 신뢰할 호스트 이름 및 IP 설정
	for _, h := range strings.Split(*host, ",") {
		// 호스트와 IP 구분하여 설정
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	// P-256 타원 곡선을 이용하여 개인키 생성
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	// 암호화를 위한 난수, 인증서 템플릿, 상위 인증서, 공개키 및 공개키와 페어인 개인키를 이용하여 DER 인증서 바이트 슬라이스 생성
	// 스스로 서명한 인증서를 생성하기 때문에, 상위 인증서 또한 인증서 템플릿 자신으로 설정
	der, err := x509.CreateCertificate(rand.Reader, &template, &template, &private.PublicKey, private)
	if err != nil {
		panic(err)
	}

	// Save generated der to cert.pem
	cert, err := os.Create(*certFileName)
	if err != nil {
		panic(err)
	}

	// 인증서를 pem 포멧으로 인코딩하여 저장
	err = pem.Encode(cert, &pem.Block{Type: "CERTIFICATE", Bytes: der})
	if err != nil {
		panic(err)
	}

	if err := cert.Close(); err != nil {
		panic(err)
	}
	log.Printf("wrote public key: %s\n", *certFileName)

	// Save private key
	// 개인키는 공유하지 않음 => 권한 설정 필수
	key, err := os.OpenFile(*keyFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	privateKey, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		panic(err)
	}

	err = pem.Encode(key, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKey})
	if err != nil {
		log.Fatal(err)
	}

	if err := key.Close(); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote private key: %s\n", *keyFileName)
}
