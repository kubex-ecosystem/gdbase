package types

import (
	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type JWT struct {
	Reference             *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath              string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	RefreshSecret         string       `gorm:"omitempty" json:"refresh_secret" yaml:"refresh_secret" xml:"refresh_secret" toml:"refresh_secret" mapstructure:"refresh_secret"`
	PrivateKey            string       `gorm:"omitempty" json:"private_key" yaml:"private_key" xml:"private_key" toml:"private_key" mapstructure:"private_key"`
	PublicKey             string       `gorm:"omitempty" json:"public_key" yaml:"public_key" xml:"public_key" toml:"public_key" mapstructure:"public_key"`
	ExpiresIn             int          `gorm:"omitempty" json:"expires_in" yaml:"expires_in" xml:"expires_in" toml:"expires_in" mapstructure:"expires_in"`
	IDExpirationSecs      int          `gorm:"omitempty" json:"id_expiration_secs" yaml:"id_expiration_secs" xml:"id_expiration_secs" toml:"id_expiration_secs" mapstructure:"id_expiration_secs"`
	RefreshExpirationSecs int          `gorm:"omitempty" json:"refresh_expiration_secs" yaml:"refresh_expiration_secs" xml:"refresh_expiration_secs" toml:"refresh_expiration_secs" mapstructure:"refresh_expiration_secs"`
	Mapper                ci.IMapper[JWT]
}
