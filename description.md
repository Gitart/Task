## Task Description
Описание задач и структуры приложения


|Fields|Description
|---|---
|Title| Наименование задачи
|Description| Краткое описание задачи
|Project| Проект
|ProjecCode|Код проекта
|Start|Старт задачи
|Finish|Окончание задачи
|Owner|Собственник задачи
|Respons|Исполнитель
|Priority|Приоритет (важно,среднее, неважно)
|Status|Выполненно, План, в работе, ожидание
|Percent|Процент выполнения



```golang

/******************************************************************************************************************************************************
 *
 * Показ одной задачи в окне
 * /api/task/one/ - путь
 * http://10.0.3.24:5555/api/task/one/8625E011B99E4AECAC25
 *
******************************************************************************************************************************************************/
func Tasks(w http.ResponseWriter, rq *http.Request) {

	// Param := rq.URL.Path[len("/tasks/"):]
	var response []Mst

	// Получение одной записи по ид
	Res, _ := r.DB("slbc").Table("tasks").Run(sessionArray[0])
	defer Res.Close()
	Res.All(&response)

	// p:=`Документ # 3546783`
	// profile:= Mst{"SS":"ddd","INFO":"dddddsss","WARNING":"ddddcccc","DANGER":Param}

	fp := path.Join("tmpl", "task.html")                                         // подключение страницы
	tmpl, err := template.ParseFiles(fp, "temp/main.gtpl", "temp/foot.gtpl")     // подключение (шаблон) шаблонов   

	// Error
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    tags := []string{"(c) SILVERSTAR","2017", "Department developin Software"}
    Dat  := Mst{"DATA": response, "HANDLE": tags, "Title": "Отчет о текущих задачах департамента разработки программного обеспечения"}
    Dat["Description"] = "Отчет  по загруженности отдела задачами"

    // Load data to template
	tmpl.Execute(w, Dat)
    
    // if  errsd != nil {
    //     http.Error(w, errsd.Error(), http.StatusInternalServerError)
    // }
}
```


### Формирование шаблона динамически

```golang
/********************************************************************************************************************************
 *
 * Использование шаблонов
 * http.HandleFunc("/api/system/report/",               GitInn)                        // Preview Report
 *
 *********************************************************************************************************************************/
func GitInn(rp http.ResponseWriter, rq *http.Request) {

	//  Structure
	type Inventory struct {
		Country string
		Index   float64
	}

	sweaters := []Inventory{{"Head", 1}, {"Move", 2}, {"System", 3}}

	sss := `<!DOCTYPE html>
	            <htmL>
	                <head>
				          <title>Head Office</title>
					      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

                          <style type="text/css">
                                 html {color: #3B3B3B; font-size:15px; font-family:Calibri;}
                                 td   {border:1px gray solid; padding:5px;}
                                .hhd  {border:1px gray solid; padding:5px; background-color: #E7E7E7; font-weight: bold;}
                                .hdsf {color: #3366FF;	font-weight: bold;	font-size:17px;}
                          </style>
  				    </head>

	                <body>
	                      <p>Результат фильтрации<p>
	                      <br>
	                      <table style='width: 300px; border:1px gray solid; border-collapse:collapse;'>

                                              <tr>
		                                          <td><b>Страна</b></td>
		                                          <td><b>Коєф</b></td>

		                                     </tr>
	                             {{range .}}

		                                     <tr>
		                                          <td><b>{{printf "%s" .Country}}</b></td>
		                                          <td><b>{{printf "%v" .Index}}</b></td>

		                                     </tr>

		                         {{end}}
		                  </table>
                    </body>
                </html>`

	tmpl, err := template.New("test").Parse(sss)

	if err != nil {
		panic(err)
	}

	// yyy := tmpl.Execute(os.Stdout, sweaters)
	err = tmpl.Execute(rp, sweaters)

	if err != nil {
		panic(err)
	}
}
```


