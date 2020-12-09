# Log-Http

This is a powerful logger for services.

### Installation

Install and start the server.

```sh
$ go run main.go
$ go get github.com/gerardoCD/go-utils 
```

Import in your project

```goland

import (
	utils "github.com/gerardoCD/go-utils"
)

```


# Features!

  - Generate SoapRequest with template
  - Generate SoapCall
  - Generate Error Class
  
 # Examples!

```goland


func main() {
	url := env.URISoap + "/AppClientesBackendServicios/"
	template := templates.RequestSOAP
	httpReq, err := utils.GenerateSOAPRequest(requestSoap,url,template)
	if err != nil {
		lgs.Error("Generando SOAP Request")
		return nil, err
	}
	
	//PETICIÓN A SERVICIO SOAP
	responseSoap, err := utils.SoapCall(httpReq)
	if err != nil {
		lgs.Error("Llamada servicio SOAP")
		return nil, err
	}
}
  
```
