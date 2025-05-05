package cli

import (
	"github.com/kelsonic-networks/kelca/internal/ca"
	"github.com/kelsonic-networks/kelca/internal/storage"
	"github.com/spf13/cobra"
)

func RegisterCommands(rootCmd *cobra.Command) {
	/* Root CA commands */
	rootCmd.AddCommand(createRootCACmd())
	// rootCmd.AddCommand(listRootCACmd())

	/* Certificate commands */
	// rootCmd.AddCommand(createCertCmd())
	// rootCmd.AddCommand(revokeCertCmd())
	// rootCmd.AddCommand(listCertsCmd())

	/* Other command groups */
}

func createRootCACmd() *cobra.Command {
	var commonName, organization string
	var keyType string
	var keySize int
	var validity int

	cmd := &cobra.Command{
		Use:   "create-root-ca",
		Short: "Create a new Root CA",
		RunE: func(cmd *cobra.Command, args []string) error {
			password, err := promptMasterPassword(true)
			if err != nil {
				return err
			}

			store, err := storage.NewSecureStorage(password)
			if err != nil {
				return err
			}

			rootCA := &ca.RootCA{
				CommonName:   commonName,
				Organization: organization,
				KeyType:      keyType,
				KeySize:      keySize,
				Validity:     validity,
			}

			return rootCA.Create(store)
		},
	}

	cmd.Flags().StringVar(&commonName, "common-name", "", "Common Name for the CA (required)")
	cmd.Flags().StringVar(&organization, "org", "Kelsonic Networks", "Organization name")
	cmd.Flags().StringVar(&keyType, "key-type", "RSA", "Key type (RSA, ECDSA)")
	cmd.Flags().IntVar(&keySize, "key-size", 4096, "Key size in bits")
	cmd.Flags().IntVar(&validity, "validity", 3650, "Validity period in days")
	cmd.MarkFlagRequired("common-name")

	return cmd
}
