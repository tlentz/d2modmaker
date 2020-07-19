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

//Added additional fields to help with check/uncheck boxes. Shouldn't affect mods in any way. Once this GUI is prod, people won't be editing the cfg.json anymore anyway.
type Settings struct {
	IncreaseStackSizes     bool `json:"IncreaseStackSizes"`
	IncreaseStackSizess    string
	IncreaseMonsterDensity int  `json:"IncreaseMonsterDensity"`
	EnableTownSkills       bool `json:"EnableTownSkills"`
	EnableTownSkillss      string
	NoDropZero             bool `json:"NoDropZero"`
	NoDropZeros            string
	QuestDrops             bool `json:"QuestDrops"`
	QuestDropss            string
	UniqueItemDropRate     int  `json:"UniqueItemDropRate"`
	StartWithCube          bool `json:"StartWithCube"`
	StartWithCubes         string
	RandomOptions          RandomOptions `json:"RandomOptions"`
}

type RandomOptions struct {
	Randomize     bool `json:"Randomize"`
	Randomizes    string
	Seed          int  `json:"Seed"`
	IsBalanced    bool `json:"IsBalanced"`
	IsBalanceds   string
	MinProps      int  `json:"MinProps"`
	MaxProps      int  `json:"MaxProps"`
	UseOSkills    bool `json:"UseOSkills"`
	UseOSkillss   string
	PerfectProps  bool `json:"PerfectProps"`
	PerfectPropss string
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
	IncreaseStackSizes        <input name="IncreaseStackSizes" type="checkbox" {{.IncreaseStackSizess}}> <br> </br>

	IncreaseMonsterDensity (min:0 max:30 or -1 to omit)     <input name="IncreaseMonsterDensity" type="text" value="{{.IncreaseMonsterDensity}}" /> <br> </br>

	EnableTownSkills            <input name="EnableTownSkills" type="checkbox" {{.EnableTownSkillss}}><br> </br>
	                                       
	NoDropZero (true/false)                 <input name="NoDropZero" type="checkbox" {{.NoDropZeros}}> <br> </br>
                                            
	QuestDrops (true/false)                  <input name="QuestDrops" type="checkbox" {{.QuestDropss}}> <br> </br>

	UniqueItemDropRate (min:0 max: 450 or -1 to omit)        <input name="UniqueItemDropRate" type="text" value="{{.UniqueItemDropRate}}" /> <br> </br>

	StartWithCube (true/false)               <input name="StartWithCube" type="checkbox" {{.StartWithCubes}}> <br> </br>

	RandomOptions.Randomize (true/false)     <input name="Randomize" type="checkbox" {{.RandomOptions.Randomizes}}> <br> </br>

	RandomOptions.Seed (set to -1 to generate random seed)        <input name="Seed" type="text" value="{{.RandomOptions.Seed}}" /> <br> </br>

	RandomOptions.IsBalanced (true/false)    <input name="IsBalanced"  type="checkbox" {{.RandomOptions.IsBalanceds}}> <br> </br>

	RandomOptions.MinProps (set to -1 to omit)    <input name="MinProps" type="text" value="{{.RandomOptions.MinProps}}" /> <br> </br>

	RandomOptions.MaxProps (set to -1 to omit)    <input name="MaxProps" type="text" value="{{.RandomOptions.MaxProps}}" /> <br> </br>

	RandomOptions.UseOSkills (true/false)    <input name="UseOSkills"  type="checkbox" {{.RandomOptions.UseOSkillss}}> <br> </br>

	RandomOptions.PerfectProps (true/false)  <input name="PerfectProps"  type="checkbox" {{.RandomOptions.PerfectPropss}}> <br> </br>

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

func adjustinput(change *int, input []string) {
	ch := *change
	if len(input) != 0 && input[0] != "" {
		converted, _ := strconv.Atoi(input[0])
		ch = converted
	}
	*change = ch
}

func checkboxGrabinput(input []string) (string, bool) {
	if len(input) != 0 && input[0] != "" {
		return "checked", true
	}
	return "", false
}

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	r.ParseForm()
	//Test demonstrating grabbing of input from 'IncreaseStackSizesBOX' checkbox
	N.IncreaseStackSizess, N.IncreaseStackSizes = checkboxGrabinput(r.Form["IncreaseStackSizes"])
	//endtest
	//refactored below into two functions
	adjustinput(&N.IncreaseMonsterDensity, r.Form["IncreaseMonsterDensity"])
	adjustinput(&N.UniqueItemDropRate, r.Form["UniqueItemDropRate"])
	adjustinput(&N.RandomOptions.Seed, r.Form["Seed"])
	adjustinput(&N.RandomOptions.MinProps, r.Form["MinProps"])
	adjustinput(&N.RandomOptions.MaxProps, r.Form["MaxProps"])
	N.EnableTownSkillss, N.EnableTownSkills = checkboxGrabinput(r.Form["EnableTownSkills"])
	N.NoDropZeros, N.NoDropZero = checkboxGrabinput(r.Form["NoDropZero"])
	N.QuestDropss, N.QuestDrops = checkboxGrabinput(r.Form["QuestDrops"])
	N.StartWithCubes, N.StartWithCube = checkboxGrabinput(r.Form["StartWithCube"])
	N.RandomOptions.Randomizes, N.RandomOptions.Randomize = checkboxGrabinput(r.Form["Randomize"])
	N.RandomOptions.IsBalanceds, N.RandomOptions.IsBalanced = checkboxGrabinput(r.Form["IsBalanced"])
	N.RandomOptions.UseOSkillss, N.RandomOptions.UseOSkills = checkboxGrabinput(r.Form["UseOSkills"])
	N.RandomOptions.PerfectPropss, N.RandomOptions.PerfectProps = checkboxGrabinput(r.Form["PerfectProps"])
	output, err := json.MarshalIndent(N, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("cfg.json", output, 0755)
	fmt.Println("Saved!!")
	fmt.Println("---")
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
