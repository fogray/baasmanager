module github.com/fogray/baasmanager/baas-gateway

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fogray/baasmanager/baas-core v0.0.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.4.0
	github.com/go-xorm/xorm v0.7.1
	github.com/spf13/viper v1.4.0
)

replace github.com/fogray/baasmanager/baas-core v0.0.0 => ../baas-core
