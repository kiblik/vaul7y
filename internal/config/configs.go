package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"gopkg.in/yaml.v2"
)

type Config struct {
	VaultAddr         string `yaml:"vault_addr"`
	VaultNamespace    string `yaml:"vault_namespace"`
	VaultToken        string `yaml:"vault_token"`
	VaultCaCert       string `yaml:"vault_cacert"`
	VaultClientCert   string `yaml:"vault_client_cert"`
	VaultClientKey    string `yaml:"vault_client_key"`
	VaultyLogFile     string `yaml:"vaulty_log_file"`
	VaultyLogLevel    string `yaml:"vaulty_log_level"`
	VaultyRefreshRate int    `yaml:"vaulty_refresh_rate"`
}

func checkForVaultAddress() {
	if os.Getenv("VAULT_ADDR") == "" {
		fmt.Println("VAULT_ADDR is not set. Please set it and try again.")
		os.Exit(1)
	}

	if os.Getenv("VAULT_TOKEN") == "" {
		fmt.Println("VAULT_TOKEN is not set. Please set it and try again.")
		os.Exit(1)
	}

}

func LoadConfig(cfgFile string) Config {
	// Load the config from the YAML file
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory")
		os.Exit(1)
	}

	var data []byte
	if cfgFile == "" {
		fmt.Println("No config file specified, using default")
		yamlFilePath := filepath.Join(home, ".vaul7y.yaml")
		data, err = ioutil.ReadFile(yamlFilePath)
		if err != nil {
			fmt.Printf("Error reading YAML file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Using config file: ", cfgFile)
		data, err = ioutil.ReadFile(cfgFile)
		if err != nil {
			fmt.Printf("Error reading YAML file: %v\n", err)
			os.Exit(1)
		}
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %v\n", err)
		os.Exit(1)
	}

	// Overwrite with environment variables if they are set
	if vaultAddr := os.Getenv("VAULT_ADDR"); vaultAddr != "" {
		config.VaultAddr = vaultAddr
	}
	if vaultNamespace := os.Getenv("VAULT_NAMESPACE"); vaultNamespace != "" {
		config.VaultNamespace = vaultNamespace
	}
	if vaultToken := os.Getenv("VAULT_TOKEN"); vaultToken != "" {
		config.VaultToken = vaultToken
	}
	if vaultCaCert := os.Getenv("VAULT_CACERT"); vaultCaCert != "" {
		config.VaultCaCert = vaultCaCert
	}
	if vaultClientCert := os.Getenv("VAULT_CLIENT_CERT"); vaultClientCert != "" {
		config.VaultClientCert = vaultClientCert
	}
	if vaultClientKey := os.Getenv("VAULT_CLIENT_KEY"); vaultClientKey != "" {
		config.VaultClientKey = vaultClientKey
	}
	if vaultyLogFile := os.Getenv("VAULTY_LOG_FILE"); vaultyLogFile != "" {
		config.VaultyLogFile = vaultyLogFile
	}
	if vaultyLogLevel := os.Getenv("VAULTY_LOG_LEVEL"); vaultyLogLevel != "" {
		config.VaultyLogLevel = vaultyLogLevel
	}
	if vaultyRefreshRate := os.Getenv("VAULTY_REFRESH_RATE"); vaultyRefreshRate != "" {
		vaultyRefreshRateInt, err := strconv.Atoi(vaultyRefreshRate)
		if err != nil {
			fmt.Printf("Error converting VAULTY_REFRESH_RATE to int: %v", err)
		} else {
			config.VaultyRefreshRate = vaultyRefreshRateInt
		}
	}

	if config.VaultAddr == "" {
		fmt.Println("VAULT_ADDR is not set. Please set it and try again.")
		os.Exit(1)
	}

	if config.VaultToken == "" {
		fmt.Println("VAULT_TOKEN is not set. Please set it and try again.")
		os.Exit(1)
	}

	if config.VaultyRefreshRate == 0 {
		config.VaultyRefreshRate = 30
	}

	if strings.EqualFold(config.VaultyLogLevel, "debug") {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGTERM)

			<-ch
			fmt.Println("Dumping goroutines")
			bufsize := int(10 * 1024 * 1024) // 10 MiB
			buf := make([]byte, bufsize)
			n := runtime.Stack(buf, true)
			filename := fmt.Sprintf("%s.dump", config.VaultyLogFile)

			ioutil.WriteFile(filename, buf[:n], 0644)
			os.Exit(1)
		}()

	}

	return config
}
