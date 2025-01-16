package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type NginxConfig struct {
	Host       string
	Port       string
	ServerName string
}

func main() {
	fmt.Println("Parsing Environment Variables...")

	tmpl, err := template.ParseFiles("nginx.conf.template")
	if err != nil {
		fmt.Printf("Error reading template file: %v\n", err)
		os.Exit(1)
	}

	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key, value := parts[0], parts[1]

		if strings.HasPrefix(key, "APP_") {

			envParts := strings.SplitN(value, ":", 3)
			if len(envParts) != 3 {
				fmt.Printf("Invalid format for %s environment variable\n", key)
				continue
			}
	
			domain, host, port := envParts[0], envParts[1], envParts[2]
			config := NginxConfig{
				Host:       host,
				Port:       port,
				ServerName: domain,
			}
	
			outputPath := fmt.Sprintf("/etc/nginx/conf.d/%s.conf", domain)
			if err := createConfigFile(outputPath, tmpl, config); err != nil {
				fmt.Printf("Error creating config for %s: %v\n", domain, err)
				continue
			}
	
			fmt.Printf("Created config for Domain: %s, Host: %s, Port: %s\n", domain, host, port)
		}
	}
}

func createConfigFile(path string, tmpl *template.Template, config NginxConfig) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, config); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
