package main

import (
	"diablomod/front"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
type Checkboxes struct {
	IncreaseStackSizesBOX string
}

//Global struct to save the settings.
var N = Settings{}
var P = Checkboxes{}

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
	IncreaseStackSizesBOX (true/false)         <input name="IncreaseStackSizesBOX" type="checkbox" value="{{.IncreaseStackSizes}}"> <br> </br>
	IncreaseStackSizes (true/false)         <input name="IncreaseStackSizes" type="text" value="{{.IncreaseStackSizes}}" /> <br> </br>
	IncreaseMonsterDensity (min:0 max:30 or -1 to omit)     <input name="IncreaseMonsterDensity" type="text" value="{{.IncreaseMonsterDensity}}" /> <br> </br>
	EnableTownSkills (true/false)            <input name="EnableTownSkills" type="text" value="{{.EnableTownSkills}}" /> <br> </br>
	NoDropZero (true/false)                 <input name="NoDropZero" type="text" value="{{.NoDropZero}}" /> <br> </br>
	QuestDrops (true/false)                  <input name="QuestDrops" type="text" value="{{.QuestDrops}}" /> <br> </br>
	UniqueItemDropRate (min:0 max: 450 or -1 to omit)        <input name="UniqueItemDropRate" type="text" value="{{.UniqueItemDropRate}}" /> <br> </br>
	StartWithCube (true/false)               <input name="StartWithCube" type="text" value="{{.StartWithCube}}" /> <br> </br>
	RandomOptions.Randomize (true/false)     <input name="Randomize" type="text" value="{{.RandomOptions.Randomize}}" /> <br> </br>
	RandomOptions.Seed (set to -1 to generate random seed)        <input name="Seed" type="text" value="{{.RandomOptions.Seed}}" /> <br> </br>
	RandomOptions.IsBalanced (true/false)    <input name="IsBalanced" type="text" value="{{.RandomOptions.IsBalanced}}" /> <br> </br>
	RandomOptions.MinProps (set to -1 to omit)    <input name="MinProps" type="text" value="{{.RandomOptions.MinProps}}" /> <br> </br>
	RandomOptions.MaxProps (set to -1 to omit)    <input name="MaxProps" type="text" value="{{.RandomOptions.MaxProps}}" /> <br> </br>
	RandomOptions.UseOSkills (true/false)    <input name="UseOSkills" type="text" value="{{.RandomOptions.UseOSkills}}" /> <br> </br>
	RandomOptions.PerfectProps (true/false)  <input name="PerfectProps" type="text" value="{{.RandomOptions.PerfectProps}}" /> <br> </br>
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

func adjustboolx(input []string) bool {
	if len(input) != 0 && input[0] != "" {
		switch strings.ToLower(input[0]) {
		case "true":
			return true
		}
	}
	return false
}

func adjustinput(change *int, input []string) {
	ch := *change
	if len(input) != 0 && input[0] != "" {
		converted, _ := strconv.Atoi(input[0])
		ch = converted
	}
	*change = ch
}

func checkboxGrabinput(input []string) {
	if len(input) != 0 && input[0] != "" {
		fmt.Println("checkbox is set to true")
		return
	}
	fmt.Println("checkbox is set to false")
}

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	r.ParseForm()
	//Test demonstrating grabbing of input from 'IncreaseStackSizesBOX' checkbox
	checkboxGrabinput(r.Form["IncreaseStackSizesBOX"])
	//endtest
	//refactored below into two functions
	adjustinput(&N.IncreaseMonsterDensity, r.Form["IncreaseMonsterDensity"])
	adjustinput(&N.UniqueItemDropRate, r.Form["UniqueItemDropRate"])
	adjustinput(&N.RandomOptions.Seed, r.Form["Seed"])
	adjustinput(&N.RandomOptions.MinProps, r.Form["MinProps"])
	adjustinput(&N.RandomOptions.MaxProps, r.Form["MaxProps"])
	N.EnableTownSkills = adjustboolx(r.Form["EnableTownSkills"])
	N.NoDropZero = adjustboolx(r.Form["NoDropZero"])
	N.QuestDrops = adjustboolx(r.Form["QuestDrops"])
	N.StartWithCube = adjustboolx(r.Form["StartWithCube"])
	N.RandomOptions.Randomize = adjustboolx(r.Form["Randomize"])
	N.RandomOptions.IsBalanced = adjustboolx(r.Form["IsBalanced"])
	N.RandomOptions.UseOSkills = adjustboolx(r.Form["UseOSkills"])
	N.RandomOptions.PerfectProps = adjustboolx(r.Form["PerfectProps"])
	N.IncreaseStackSizes = adjustboolx(r.Form["IncreaseStackSizes"])
	fmt.Println(N)
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
