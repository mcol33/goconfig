package config

import (
    "os"
	"log"
    "encoding/json"
    "io/ioutil"
)

type ApplicationsConfigurations struct {
	ApplicationsConfigurations []ApplicationConfigurations `json:"applicationsConfigurations"`
}

type ApplicationConfigurations struct {
	ApplicationName string `json:"applicationName"`
	Configurations []Configuration `json:"configurations"`
}

type Configuration struct {
	EnvironmentName string `json:"environmentName"`
	AuthDbName string `json:"authDbName"`
	Databases []Database `json:"databases"`
	EmailGateway EmailGateway `json:"emailGateway"`
	OAuth OAuth `json:"oauth"`
	LogOutput `json:"logOutput"`
	Files []File `json:"files"`
	Temporal Temporal `json:"temporal"`
}

type Database struct { 
	DatabaseName string `json:"databaseName"`
	DatabaseProduct string `json:"databaseProduct"`
	DatabaseUserName string `json:"databaseUserName"`
	DatabasePassword string `json:"databasePassword"`
	DatabaseHost string `json:"databaseHost"`
	DatabasePort string `json:"databasePort"`
}

type EmailGateway struct {
	From string `json:"from"`
    Password string `json:"password"`
	Host string `json:"host"`
    Port string `json:"port"`

}

type OAuth struct {
	RedirectURL string `json:"redirectURL"`
	ClientID string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Scopes string `json:"scopes"`
	Endpoint string `json:"endpoint"`
}

type LogOutput struct {
	LogOutput string `json:"logOutput"`//"stdstream","file","database"
	LogOutputReference string `json:"logOutputReference"` //"stdout", "stderr", "file_name", "database_name"
}

type File struct {
	FileName string `json:"fileName"`
	FileDirectoryPath string `json:"fileDirectoryPath"`
}

type Temporal struct {
	HostPort string `json:"hostPort"`
	TaskQueue string `json:"taskQueue"`
	WorkflowID string `json:"workflowID"`
}

func GetEnvironment () (string){
	envVariableValue, ok := os.LookupEnv("ENVIRONMENT")
    if ok == false {
        return "Development"
    } else {
		 return envVariableValue
	}
}

func getConfigFilePathName () (string){
	const configFilePath = "config"
	const configName = "config"
	return configFilePath+"/"+configName
}

func GetConfig (configuration Configuration, configurationItemName string)(string){

	switch configurationItemName {
	case "from" :
		return configuration.EmailGateway.From
	case "password" :
		return configuration.EmailGateway.Password
	case "host" :
		return configuration.EmailGateway.Host
	case "port" :
		return configuration.EmailGateway.Port
	case "redirectURL" :
		return configuration.OAuth.RedirectURL
	case "clientID" :
		return configuration.OAuth.ClientID
	case "clientSecret" :
		return configuration.OAuth.ClientSecret
	case "scopes" :
		return configuration.OAuth.Scopes
	case "endpoint" :
		return configuration.OAuth.Endpoint	
	default :
		log.Print("In default")
		return ""
	}
}

func GetDatabaseConfig (configuration Configuration, databaseName string)(Database){
	var indexDB int
	for i:=0; i<len(configuration.Databases); i++ {
		if configuration.Databases[i].DatabaseName==databaseName {
			indexDB = i
		}
	}
	return configuration.Databases[indexDB]
}

func LoadConfigInMemory (applicationName string)(Configuration,error){
	//for _, env := range os.Environ() {
	//	log.Print(env)
	//}
    
    // Open our jsonFile
    //jsonFile, err := os.Open("config/config")


	//var configFilePathName string
	//configFilePathName= configFilePath+"/"+configName

	jsonFile, err := os.Open(getConfigFilePathName())
    if err != nil {
        log.Print(err)
    }

    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

   
    var applicationsConfigurations ApplicationsConfigurations
	
    json.Unmarshal(byteValue, &applicationsConfigurations)
   
	var indexApp int
	for i:=0; i<len(applicationsConfigurations.ApplicationsConfigurations); i++ {
		log.Print("Application name (2) = "+applicationsConfigurations.ApplicationsConfigurations[i].ApplicationName)
		if applicationsConfigurations.ApplicationsConfigurations[i].ApplicationName==applicationName{
			indexApp = i
			//log.Print("indexEnv :")
			//log.Print(indexEnv)
		}
	}
	log.Print("environement = "+GetEnvironment())
	var indexEnv int 
	for j:=0; j<len(applicationsConfigurations.ApplicationsConfigurations[indexApp].Configurations); j++ {
		log.Print("Environment name (2) = "+ applicationsConfigurations.ApplicationsConfigurations[indexApp].Configurations[j].EnvironmentName)
		if applicationsConfigurations.ApplicationsConfigurations[indexApp].Configurations[j].EnvironmentName==GetEnvironment() {
			indexEnv = j
			//log.Print("indexEnv :")
			//log.Print(indexEnv)
		}
	}
	return applicationsConfigurations.ApplicationsConfigurations[indexApp].Configurations[indexEnv], nil
}
