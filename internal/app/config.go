package app

import "github.com/spf13/viper"

func InitConfig() {

	// viper.SetDefault("game_url", "https://uat-web-game-fe-op.bpweg.com")
	// viper.SetDefault("operator_id", "hwidradminjci70")
	// viper.SetDefault("app_secret", "CuA3JE7v2FytMNHPF9lJOet2votjy03AAHlHbEMcJEM=")

	viper.SetDefault("game_url", "https://release-web-game-fe.wehosts247.com")
	viper.SetDefault("operator_id", "ta1admin1edag")
	viper.SetDefault("app_secret", "Ac_l8-l76hX127nBxa-8fZnLPKpNHzQm-z1McYxXGdo=")
}
