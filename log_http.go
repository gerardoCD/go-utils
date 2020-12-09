package utils

import (
	// "app/model"
	// "strconv"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	"text/template" 

	"github.com/google/logger"
	lgs "github.com/gerardoCD/log-http-go"
	"encoding/json"
);


type BadResponse struct {
	ErrorMessage  string      `json:"errorMessage"`
	ErrorDev      string      `json:"errorDev"`
	ErrorResponse interface{} `json:"errorResponse"`
}


// Helper response Error and logger
func ErrorLog(response *BadResponse, mensaje string, code int, logs string, err error) {
	response.ErrorMessage=mensaje
	response.ErrorDev=err.Error()
	response.ErrorResponse=err

	reponseBodyString,_ := json.Marshal(response)
	lgs.Reponse(code,string(reponseBodyString))
}


//Construlle Request SOAP
func GenerateSOAPRequest(Req interface{}, url string, templates string) (*http.Request, error) {
	//Utilizar la cadena del template para construir el request para el servicio SOAP
	template, err := template.New("InputRequest").Parse(templates)
	if err != nil {
		logger.Error("Error al formar el objeto de template: ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	err = template.Execute(doc, Req)
	if err != nil {
		logger.Error("Error al reemplazar los valores del request en el template: ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		logger.Error("Encode error", err.Error())
		return nil, err
	}


	 
	 lgs.RequestServiceStringXML(doc.String(),url,"Get Parametros")

	r, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(doc.String())))
	r.Header.Set("Content-type", "text/xml;charset=UTF-8")
	//r.SetBasicAuth(env.UserService, env.PasswordService)
	if err != nil {
		logger.Error("Error al construir el request: ", err.Error())
		return nil, err
	}



	return r, nil
}

//Realiza llamda SOAP - Despues de construccion del request
func SoapCall(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)


	if err != nil {
		logger.Error("Error al realizar el request al servicio SOAP: ", err.Error())
		return nil, err
	}
	//Obtener el body del POST request hacia el servicio SOAP
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Error("Error al obtener el body que devuelve el Servicio SOAP: ", err.Error())
		return nil, err
	}


	lgs.ReponseServiceStringXML(resp.StatusCode,string(body),req.URL.String(),"Get Parametros")
	
	defer resp.Body.Close()
	
	return body, nil
}

