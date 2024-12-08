// NOT TESTED YET

package commands

import (
	"bebra/helpers"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/pavlo-v-chernykh/keystore-go/v4"
	"github.com/spf13/cobra"
)

var keystoreCmd = &cobra.Command{
	Use:   "keystore",
	Short: "Keystore Tool is a CLI tool to create and manage Java KeyStores.",
	Run:   createKeyStore,
}

func init() {
	keystoreCmd.Flags().StringP("password", "p", "", "Password for the keystore (required)")
	keystoreCmd.Flags().StringP("output", "o", "keystore.jks", "Output keystore file name")
	keystoreCmd.MarkFlagRequired("password")
}

func generateKeyPair() (*rsa.PrivateKey, []byte, error) {
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create a self-signed certificate template
	certTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Bebra Company"},
			CommonName:   "my-key-alias",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10000 * 24 * time.Hour), // Validity: 10000 days
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Self-sign the certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, certBytes, nil
}

func writeKeyStore(ks keystore.KeyStore, filename string, password []byte) {
	f := helpers.CreateFile(filename)
	err := ks.Store(f, password)
	if err != nil {
		log.Fatal(err)
	}
}

func zeroing(buf []byte) {
	for i := range buf {
		buf[i] = 0
	}
}

func createKeyStore(cmd *cobra.Command, args []string) {
	password, _ := cmd.Flags().GetString("password")
	outputFile, _ := cmd.Flags().GetString("output")

	// Convert password to byte slice and securely zero it out after use
	passBytes := []byte(password)
	defer zeroing(passBytes)

	// Generate RSA key pair and self-signed certificate
	privateKey, certBytes, err := generateKeyPair()
	if err != nil {
		log.Fatalf("Failed to generate key pair: %v", err)
		os.Exit(1)
	}

	// Marshal private key to PKCS#8
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Failed to marshal private key: %v", err)
		os.Exit(1)
	}

	// Initialize keystore
	ks := keystore.New()

	// Create PrivateKeyEntry
	pkeIn := keystore.PrivateKeyEntry{
		CreationTime:     time.Now(),
		PrivateKey:       privateKeyBytes,
		CertificateChain: []keystore.Certificate{{Type: "X509", Content: certBytes}},
	}

	// Add to keystore
	if err := ks.SetPrivateKeyEntry("my-key-alias", pkeIn, passBytes); err != nil {
		helpers.ErrorLog(fmt.Sprintf("Set private key entry failed: %v\n", err))
		os.Exit(1)
	}

	// Write keystore to file
	writeKeyStore(ks, outputFile, passBytes)

	fmt.Printf("Keystore saved to %s\n", outputFile)
}
