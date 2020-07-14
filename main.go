package main

import (
	"diablomod/front"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Settings struct {
	IncreaseStackSizes     bool          `json:"IncreaseStackSizes"`
	IncreaseMonsterDensity int           `json:"IncreaseMonsterDensity"`
	EnableTownSkills       bool          `json:"EnableTownSkills"`
	NoDropZero             bool          `json:"NoDropZero"`
	QuestDrops             bool          `json:"QuestDrops"`
	UniqueItemDropRate     int           `json:"UniqueItemDropRate"`
	StartWithCube          bool          `json:"StartWithCube"`
	RandomOptions          RandomOptions `json:"RandomOptions"`
}

type RandomOptions struct {
	Randomize    bool `json:"Randomize"`
	Seed         int  `json:"Seed"`
	IsBalanced   bool `json:"IsBalanced"`
	MinProps     int  `json:"MinProps"`
	MaxProps     int  `json:"MaxProps"`
	UseOSkills   bool `json:"UseOSkills"`
	PerfectProps bool `json:"PerfectProps"`
}

//PageVariables - GUI variables that change on webpages.
type PageVariables struct {
	Output string
}

//Global struct to save the settings.
var N = Settings{}

//Main page for the front end.
var mapage = `<title>d2modmaker config editor</title>
</head>
<body>
    <style>
        input {
            display: inline-block;
            float: left;
            margin-right: 20px;
			background-color: #ffc2c2
		}
    </style>
    <a href="https://imgbb.com/"><img src="https://i.ibb.co/kxScnHd/diablo2.jpg" alt="Diablo II Logo (small)" border="0"></a>
    <p><b>d2modmaker config editor</b></p>
    <p>View and edit the cfg.json file below...</p>
  
  
</body>
<form action="/process" method="POST">
	IncreaseStackSizes         <input name="IncreaseStackSizes" type="text" value="{{.IncreaseStackSizes}}" /> <br> </br>
	IncreaseMonsterDensity     <input name="IncreaseMonsterDensity" type="text" value="{{.IncreaseMonsterDensity}}" /> <br> </br>
	EnableTownSkills           <input name="EnableTownSkills" type="text" value="{{.EnableTownSkills}}" /> <br> </br>
	NoDropZero                 <input name="NoDropZero" type="text" value="{{.NoDropZero}}" /> <br> </br>
	QuestDrops                 <input name="QuestDrops" type="text" value="{{.QuestDrops}}" /> <br> </br>
	UniqueItemDropRate         <input name="UniqueItemDropRate" type="text" value="{{.UniqueItemDropRate}}" /> <br> </br>
	StartWithCube              <input name="StartWithCube" type="text" value="{{.StartWithCube}}" /> <br> </br>
	RandomOptions.Randomize    <input name="Randomize" type="text" value="{{.RandomOptions.Randomize}}" /> <br> </br>
	RandomOptions.Seed         <input name="Seed" type="text" value="{{.RandomOptions.Seed}}" /> <br> </br>
	RandomOptions.IsBalanced   <input name="IsBalanced" type="text" value="{{.RandomOptions.IsBalanced}}" /> <br> </br>
	RandomOptions.MinProps     <input name="MinProps" type="text" value="{{.RandomOptions.MinProps}}" /> <br> </br>
	RandomOptions.MaxProps     <input name="MaxProps" type="text" value="{{.RandomOptions.MaxProps}}" /> <br> </br>
	RandomOptions.UseOSkills   <input name="UseOSkills" type="text" value="{{.RandomOptions.UseOSkills}}" /> <br> </br>
	RandomOptions.PerfectProps <input name="PerfectProps" type="text" value="{{.RandomOptions.PerfectProps}}" /> <br> </br>
	<button type="submit" style="background-color: #ffc2c2;" value="save">Save</button> Save the CFG file.
</form>
<form action="/run" method="POST">
<button type="submit" style="background-color: #ffc2c2;" value="run">Run</button> Run (?)
</form>`

//Main page handler - the front page.
func mainpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	err2 := t.Execute(w, N)
	if err2 != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

//Main page handler - the front page.
func run(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	fmt.Println("Running the thing...")
	//
	//
	//Running app logic goes here.
	//
	//
	err2 := t.Execute(w, N)
	if err2 != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func adjustbool(input string) bool {
	switch input {
	case "true":
		return true
	case "True":
		return true
	}
	return false
}

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	r.ParseForm()
	if len(r.Form["IncreaseStackSizes"]) != 0 && r.Form["IncreaseStackSizes"][0] != "" {
		N.IncreaseStackSizes = adjustbool(r.Form["IncreaseStackSizes"][0])
	}
	if len(r.Form["IncreaseMonsterDensity"]) != 0 && r.Form["IncreaseMonsterDensity"][0] != "" {
		N.IncreaseMonsterDensity, _ = strconv.Atoi(r.Form["IncreaseMonsterDensity"][0])
	}
	if len(r.Form["EnableTownSkills"]) != 0 && r.Form["EnableTownSkills"][0] != "" {
		N.EnableTownSkills = adjustbool(r.Form["EnableTownSkills"][0])
	}
	if len(r.Form["NoDropZero"]) != 0 && r.Form["NoDropZero"][0] != "" {
		N.NoDropZero = adjustbool(r.Form["NoDropZero"][0])
	}
	if len(r.Form["QuestDrops"]) != 0 && r.Form["QuestDrops"][0] != "" {
		N.QuestDrops = adjustbool(r.Form["QuestDrops"][0])
	}
	if len(r.Form["UniqueItemDropRate"]) != 0 && r.Form["UniqueItemDropRate"][0] != "" {
		N.UniqueItemDropRate, _ = strconv.Atoi(r.Form["UniqueItemDropRate"][0])
	}
	if len(r.Form["StartWithCube"]) != 0 && r.Form["StartWithCube"][0] != "" {
		N.StartWithCube = adjustbool(r.Form["StartWithCube"][0])
	}
	if len(r.Form["Randomize"]) != 0 && r.Form["Randomize"][0] != "" {
		N.RandomOptions.Randomize = adjustbool(r.Form["Randomize"][0])
	}
	if len(r.Form["Seed"]) != 0 && r.Form["Seed"][0] != "" {
		N.RandomOptions.Seed, _ = strconv.Atoi(r.Form["Seed"][0])
	}
	if len(r.Form["IsBalanced"]) != 0 && r.Form["IsBalanced"][0] != "" {
		N.RandomOptions.IsBalanced = adjustbool(r.Form["IsBalanced"][0])
	}
	if len(r.Form["MinProps"]) != 0 && r.Form["MinProps"][0] != "" {
		N.RandomOptions.MinProps, _ = strconv.Atoi(r.Form["MinProps"][0])
	}
	if len(r.Form["MaxProps"]) != 0 && r.Form["MaxProps"][0] != "" {
		N.RandomOptions.MaxProps, _ = strconv.Atoi(r.Form["MaxProps"][0])
	}
	if len(r.Form["UseOSkills"]) != 0 && r.Form["UseOSkills"][0] != "" {
		N.RandomOptions.UseOSkills = adjustbool(r.Form["UseOSkills"][0])
	}
	if len(r.Form["PerfectProps"]) != 0 && r.Form["PerfectProps"][0] != "" {
		N.RandomOptions.PerfectProps = adjustbool(r.Form["PerfectProps"][0])
	}
	fmt.Println(N)
	output, err := json.MarshalIndent(N, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("cfg.json", output, 0755)
	fmt.Println("Saved!!")
	err2 := t.Execute(w, N)
	if err2 != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func main() {
	jsonFile, _ := ioutil.ReadFile("cfg.json")
	json.Unmarshal([]byte(jsonFile), &N)
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/process", process)
	http.HandleFunc("/run", run)
	serverPort := front.LaunchServer()
	err2 := http.ListenAndServe(serverPort, nil) // setting listening port
	if err2 != nil {
		fmt.Println(err2)
		log.Fatal("ListenAndServe: ", err2)
	}
}
