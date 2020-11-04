package config

import "os"

// IsLocal はローカル環境かどうか返却します。
func IsLocal() bool {
	return os.Getenv("ENV") == "local"
}

// IsDev はDev環境かどうか返却します。
func IsDev() bool {
	return os.Getenv("ENV") == "dev"
}

// Port はサーバのポート番号を取得します。
func Port() string {
	return os.Getenv("PORT")
}

// CORSAllowOrigin はCORSのAllow Originを取得します。
func CORSAllowOrigin() string {
	return os.Getenv("CORS_ALLOW_ORIGIN")
}
