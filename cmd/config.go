package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/ziwon/crawli/crawli"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create default config file",
	Run:   config,
}

type CrawliConfig struct {
	Default struct {
		Home string `toml:"home"`
	} `toml:"default"`

	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
	} `toml:"database"`

	Workers struct {
		Min     int    `toml:"min"`
		Max     int    `toml:"max"`
		Crontab string `toml:"crontab"`
	} `toml:"workers"`
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func config(cmd *cobra.Command, args []string) {
	cfg := &CrawliConfig{
		Default: struct {
			Home string `toml:"home"`
		}{
			Home: crawli.DefaultAppHome,
		},
		Database: struct {
			Host     string `toml:"host"`
			Port     int    `toml:"port"`
			User     string `toml:"user"`
			Password string `toml:"password"`
		}{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "",
		},
		Workers: struct {
			Min     int    `toml:"min"`
			Max     int    `toml:"max"`
			Crontab string `toml:"crontab"`
		}{
			Min:     1,
			Max:     10,
			Crontab: "0 3 * * *",
		},
	}

	if err := createConfig(cfg); err != nil {
		os.Exit(1)
	}
}

func createConfig(cfg *CrawliConfig) error {
	cfgPath := path.Join(cfg.Default.Home, "config")
	err := os.MkdirAll(cfgPath, os.ModePerm)
	if err != nil {
		return err
	}

	data, err := toml.Marshal(*cfg)
	if err != nil {
		return err
	}

	cfgFile := path.Join(cfgPath, "config.toml")
	if info, err := os.Stat(cfgFile); err == nil && !info.IsDir() {
		err = os.Rename(cfgFile, path.Join(cfgPath, "config.toml.bak"))
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(cfgFile, data, os.ModePerm)
}
