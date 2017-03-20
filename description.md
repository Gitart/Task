## Task Description
Описание задач и структуры приложения


|Fields|Description
|---|---
|Title| Наименование задачи
|Description| Краткое описание задачи
|Project| Проект
|Codep|Код проекта
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

