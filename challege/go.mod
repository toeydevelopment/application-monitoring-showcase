module challenge

go 1.15

replace gitlab.com/fillgoods-library/microservice/datasource => gitlab.com/fillgoods-library/microservice/datasource.git v0.5.1

replace gitlab.com/fillgoods-library/microservice/fg-notify => gitlab.com/fillgoods-library/microservice/fg-notify.git v0.0.1

replace gitlab.com/fillgoods-library/microservice/vault => gitlab.com/fillgoods-library/microservice/vault.git v1.0.5

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/spf13/viper v1.7.1
	gitlab.com/fillgoods-library/microservice/datasource v0.0.0-00010101000000-000000000000
	gitlab.com/fillgoods-library/microservice/fg-notify v0.0.0-00010101000000-000000000000
	gitlab.com/fillgoods-library/microservice/vault v0.0.0-00010101000000-000000000000
	go.elastic.co/apm/module/apmgin v1.9.0 // indirect
	go.elastic.co/apm/module/apmsql v1.9.0 // indirect
)
