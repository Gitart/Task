package main

import (
	    "fmt"
        "net/http"
        "time"
        "html/template"
        "runtime"
        "path"
        "log"
	  r "github.com/dancannon/gorethink"
)


// Declaration inetrfaces & structure & other type
type (
	Mii interface{}                                          // Interface
	Mif []interface{}                                        // Cрез Interface
	Msr []string                                             // Срез String
	Mst map[string]interface{}                               // Map - string - interface
	Mss map[string]string                                    // Map - string - string
	Msi map[string]int64                                     // Map - string - int64
	Mis map[int64]string                                     // Map - int64 - string
	Msl map[int]string                                       // Map - int   - string
	Mil []int64
)


// System Constants
const (
	AdressService          = "127.0.0.1"                      // Local адресс
	PortService            = ":5556"                          // Local Port 5555
	DatabaseName           = "Work"                           // System - Test Databse
	AdressProductionServer = "10.0.50.16:28015"               // IP Server Morion
	AdressMorionServer     = "10.0.50.16"                     // IP Local Server
	AdressExternal         = "195.128.18.66:28015"            // IP World External for Connect
	AddressMainserver      = "localhost:28015"                // Local Host Mashine
	ServerVersion          = "2.0.5"                          // Current Version
	CodeMirrorVer          = "4.8.001"                        // Editor version
	ServerDescript         = "Draft"                          // Draft Version
	DSYS                   = "System"                         // System Database
	Organization           = "Starlight"                      // Organization Name
    ServiceClientKey       = "S0864AA791CE7E7B00R1T$$"        // Secret Client Key By Default  
    ServiceSecretKey       = "A0AEC09A647A688A64A28"          // Secret Service Key By Default 
    DatabaseKey            = ""                               // Secret Key for Database
    SecKey                 = "KeySecret$"                     // AccessKey 
)    


var (
   sessionArray []*r.Session
   IpPort string
   Startdisplay string
)




/**********************************************************
  Подключение к базе данных 
**********************************************************/
func Dbini() {

	// Инициализация подключения к базе
	// на той машине где расположен и стартует сервис
	// Для переключения на тестовую машину

	if runtime.GOOS == "windows" {
	   IpPort = "172.125.164.16:28015"                        // Тестовый сервис

        // Connect to Production Server
        production := false
        if production == true { 
           IpPort = "193.111.9.202:28015"
        }

	}  else {
	     IpPort = "localhost:28015"                       // Локальный ресурс 
	}


	// Сессия подключения
	// MaxActive: 100, IdleTimeout: time.Second * 10, MaxIdle: 10})

	// Для случая когда будет установлен пароль  базы данных
	// ServiceClientKey       = "S0864AA791CE7E7B00RT"
    // AuthKey: DatabaseKey,
    // 

	// session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: DatabaseName, Password: ServiceClientKey})
	// session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: DatabaseName, AuthKey: DatabaseKey, Username:"admin", Password: ServiceClientKey})
	session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: DatabaseName, AuthKey: DatabaseKey })


    // Закрыть сеесию 
	// defer session.Close()

	// Для случая когда будет установлен пароль для админа + ключ базы данных
	// session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: DatabaseName, AuthKey: DatabaseKey, Username: "admin", Password: ServiceClientKey})

	// Для случая когда будет установлен пароль для админа
	// session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: DatabaseName, Username: "admin", Password: ""})

	// Обработка ошибок
	if err != nil {
	   fmt.Println("NO CONNECTION.")
	   return
	}

	// Максимальное количество подключений
	// По умолчанию 200
	session.SetMaxOpenConns(200)
	session.SetMaxIdleConns(200)
	sessionArray = append(sessionArray, session)
}




/**********************************************************
 
 Стартовая процедура
 
**********************************************************/
func main(){

Startdisplay = `
   *************************************************
   Name    : Report Center
   Port    : 5555
   Author  : Savchenko Arthur
   Version : 1.0 Draft 
   Date    : 20.03.2017 17:30
   Company : (c) Star
   Start   : %s
   *************************************************
`
Starttime:=time.Now().Format("02-01-2006 15:04:05")

     Dbini()

     // Hendlers 
     http.HandleFunc("/",                StaticPage)    
     http.HandleFunc("/tasks/",          Tasks)    

     fmt.Printf(Startdisplay,Starttime)
	 http.ListenAndServe(":5555", nil)   // Listing Port PortService
}


/****************************************************************
 *   Статические странички
 *   c установкой разрешений и доступов на операции
 *
 *	 http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, r.URL.Path[1:])})
 *	 http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.FileServer(http.Dir("/static/"))})
 *	 /static/....
 *
 ****************************************************************/
func StaticPage(w http.ResponseWriter, r *http.Request) {
 	 w.Header().Set("Access-Control-Allow-Origin", "*")     // Allows
	 http.ServeFile(w, r, r.URL.Path[1:])                  //  File static page
}


/******************************************************************
 *
 * Показ одной задачи в окне
 * /api/task/one/ - путь
 * http://10.0.3.24:5555/api/task/one/8625E011B99E4AECAC25
 *
*******************************************************************/
func Tasks(w http.ResponseWriter, rq *http.Request) {

	// Param := rq.URL.Path[len("/tasks/"):]
	var response []Mst

	// Получение одной записи по ид
	Res, err := r.DB("slbc").Table("tasks").Run(sessionArray[0])
	
	if err !=nil{
       log.Fatalf("Error connect to database", err)
	}
	defer Res.Close()
	Res.All(&response)

	// p:=`Документ # 3546783`
	// profile:= Mst{"SS":"ddd","INFO":"Inf","WARNING":"News","DANGER":Param}

	fp        := path.Join("tmpl", "task.html")                                         // подключение страницы
	tmpl, err := template.ParseFiles(fp, "temp/main.gtpl", "temp/foot.gtpl")     // подключение (шаблон) шаблонов   

	// Error
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    // Формирование датнных для отчета
    tags := []string{"(c) STAR","2017", "Department developing Software"}
    Dat  := Mst{"DATA": response, "HANDLE": tags, "Title": "Отчет о текущих задачах департамента разработки программного обеспечения"}
    Dat["Description"] = "Отчет по загруженности отдела задачами"

    // Load data to template
	tmpl.Execute(w, Dat)
    
    // if  errsd != nil {
    //     http.Error(w, errsd.Error(), http.StatusInternalServerError)
    // }
}
