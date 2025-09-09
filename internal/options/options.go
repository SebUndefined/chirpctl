package options

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/spf13/pflag"
)

type Options struct {
	Profile    string        `json:"profile"`
	Addr       string        `json:"addr"`
	Insecure   bool          `json:"insecure"`
	CACert     string        `json:"caCert"`
	ServerName string        `json:"serverName"`
	Token      string        `json:"token"`
	Timeout    time.Duration `json:"timeout"`
	Output     string        `json:"output"`
	Verbose    bool          `json:"verbose"`
	ApiVersion string        `json:"apiVersion"`
}

func (o Options) MarshalJSON() ([]byte, error) {
	type Alias Options
	masked := Alias(o)

	// Table des champs sensibles et leur logique de masquage
	if o.Token != "" {
		masked.Token = fmt.Sprintf("%s############################", lo.Substring(o.Token, 0, 4))
	}
	return json.Marshal(masked)
}

// Resolve : construit les Options avec ordre de priorité
// 1. Profil YAML
// 2. Variables d'environnement CHIPCTL_*
// 3. Flags Cobra (--xxx)
func Resolve(cfg *config.File, flags *pflag.FlagSet, env map[string]string) Options {
	// helpers pour ENV
	getEnv := func(k string) (string, bool) {
		if env != nil {
			v, ok := env[k]
			return v, ok
		}
		return os.LookupEnv(k)
	}
	getEnvBool := func(k string) (bool, bool) {
		if s, ok := getEnv(k); ok {
			if s == "" {
				return true, true
			}
			if b, err := strconv.ParseBool(s); err == nil {
				return b, true
			}
		}
		return false, false
	}
	getEnvDur := func(k string) (time.Duration, bool) {
		if s, ok := getEnv(k); ok {
			if d, err := time.ParseDuration(s); err == nil {
				return d, true
			}
		}
		return 0, false
	}

	var o Options

	// --- 1) Profil YAML ---
	// choisir le profil (flag > env > default)
	profileName := ""
	if flags != nil && flags.Changed("profile") {
		profileName, _ = flags.GetString("profile")
	}
	if profileName == "" {
		if v, ok := getEnv("CHIPCTL_PROFILE"); ok {
			profileName = v
		}
	}
	if profileName == "" && cfg.DefaultContext != "" {
		profileName = cfg.DefaultContext
	}
	if p, ok := cfg.Contexts[profileName]; ok {
		o.Profile = profileName
		o.Addr = p.Addr
		o.Insecure = p.Insecure
		o.CACert = p.CACert
		o.ServerName = p.ServerName
		o.Token = p.Token
		if p.Timeout != "" {
			if d, err := time.ParseDuration(p.Timeout); err == nil {
				o.Timeout = d
			}
		}
		if p.ApiVersion != "" {
			o.ApiVersion = strings.ToLower(p.ApiVersion)
		}
	}

	// --- 2) Overrides ENV ---
	if v, ok := getEnv("CHIPCTL_ADDR"); ok {
		o.Addr = v
	}
	if v, ok := getEnv("CHIPCTL_CACERT"); ok {
		o.CACert = v
	}
	if v, ok := getEnv("CHIPCTL_SERVER_NAME"); ok {
		o.ServerName = v
	}
	if v, ok := getEnv("CHIPCTL_TOKEN"); ok {
		o.Token = v
	}
	if b, ok := getEnvBool("CHIPCTL_INSECURE"); ok {
		o.Insecure = b
	}
	if d, ok := getEnvDur("CHIPCTL_TIMEOUT"); ok && d > 0 {
		o.Timeout = d
	}
	if v, ok := getEnv("CHIPCTL_OUTPUT"); ok {
		o.Output = strings.ToLower(v)
	}
	if b, ok := getEnvBool("CHIPCTL_VERBOSE"); ok {
		o.Verbose = b
	}
	if v, ok := getEnv("CHIPCTL_API_VERSION"); ok {
		o.ApiVersion = strings.ToLower(v)
	}

	// --- 3) Overrides FLAGS ---
	if flags != nil {
		if flags.Changed("addr") {
			o.Addr, _ = flags.GetString("addr")
		}
		if flags.Changed("insecure") {
			o.Insecure, _ = flags.GetBool("insecure")
		}
		if flags.Changed("cacert") {
			o.CACert, _ = flags.GetString("cacert")
		}
		if flags.Changed("server-name") {
			o.ServerName, _ = flags.GetString("server-name")
		}
		if flags.Changed("token") {
			o.Token, _ = flags.GetString("token")
		}
		if flags.Changed("timeout") {
			o.Timeout, _ = flags.GetDuration("timeout")
		}
		if flags.Changed("output") {
			o.Output, _ = flags.GetString("output")
		}
		if flags.Changed("verbose") {
			o.Verbose, _ = flags.GetBool("verbose")
		}
		if flags.Changed("api-version") {
			o.ApiVersion, _ = flags.GetString("api-version")
		}
		// profile déjà géré au-dessus
	}

	// --- 4) Defaults & normalisation ---
	if o.Output == "" {
		o.Output = "table"
	}
	if o.Timeout <= 0 {
		o.Timeout = 10 * time.Second
	}
	if o.ApiVersion != "" {
		o.ApiVersion = strings.TrimPrefix(strings.ToLower(o.ApiVersion), "v")
	}

	return o
}
