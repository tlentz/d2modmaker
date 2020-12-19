import React, { useEffect } from "react";
import "./Main.scss";
import { withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import {
  Button,
  ButtonGroup,
  Checkbox,
  FormControlLabel,
  Grid,
  Tooltip,
} from "@material-ui/core";
import Slider from "@material-ui/core/Slider";
import InputNumber from "rc-input-number";
import HelpOutlineOutlinedIcon from "@material-ui/icons/HelpOutlineOutlined";
import Badge from "@material-ui/core/Badge";
import Divider from "@material-ui/core/Divider";
import TextField from "@material-ui/core/TextField";

const _ = require('lodash');
const axios = require("axios");

const defaultCfg = {
  Version: "v0.5.2-alpha-10",
  SourceDir: "",
  OutputDir: "",
  MeleeSplash: true,
  IncreaseStackSizes: true,
  IncreaseMonsterDensity: 1,
  EnableTownSkills: true,
  NoDropZero: true,
  QuestDrops: true,
  UniqueItemDropRate: 1,
  RuneDropRate: 1,
  StartWithCube: true,
  Cowzzz: true,
  RemoveLevelRequirements: false,
  RemoveAttRequirements: false,
  RemoveUniqCharmLimit: false,
  PerfectProps: false,
  UseOSkills: true,
  EnterToExit: false,
  RandomOptions: {
    Randomize: true,
    UseSeed: false,
    Seed: -1,
    SetsSeed: -1,
    IsBalanced: true,
    BalancedPropCount: true,
    AllowDupProps: false,
    MinProps: 0,
    MaxProps: 20,
    NumClones: 9,
  },
  GeneratorOptions: {
    Generate: true,
    UseSeed: false,
    Seed: -1,
    SetsSeed: -1,
    BalancedPropCount: true,
    MinProps: 2,
    MaxProps: 20,
    NumClones: 9,
    PropScoreMultiplier: 1,
  }
};

export default function D2ModMaker() {
  const [state, setState] = React.useState(defaultCfg);

  async function loadConfig() {
    const result = await axios("http://localhost:8148/api/cfg")
    var data = _.merge(defaultCfg, result.data);
    data = result.data;
    data.Version = defaultCfg.Version;
    return data;
  }

  useEffect(() => {
    async function fetchData() {
      let data = await loadConfig();
      setState(data);
    }
    fetchData();
  }, [])

  const updateRandomOptions = (oldState, key, val) => {
    let randomOptions = oldState.RandomOptions;
    randomOptions[key] = val;
    return { ...oldState, RandomOptions: randomOptions };
  };

  const mkRandoCheckbox = ({ key, tooltip }) => {
    return (
      <React.Fragment>
        <FormControlLabel
          control={
            <Checkbox
              color="primary"
              name={key}
              value={state.RandomOptions[key]}
            />
          }
          label={key}
          checked={state.RandomOptions[key]}
          onChange={(e, checked) => {
            return setState(updateRandomOptions(state, key, checked));
          }}
        />
        <StyledTooltip title={tooltip} placement="bottom" enterDelay={250}>
          <span className={"help-icon"}>
            <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
          </span>
        </StyledTooltip>
      </React.Fragment>
    );
  };

  const mkCheckbox = ({ key, tooltip }) => {
    return (
      <React.Fragment>
        <FormControlLabel
          control={<Checkbox color="primary" name={key} value={state[key]} />}
          label={key}
          checked={state[key]}
          onChange={(e, checked) => {
            return setState({ ...state, [key]: checked });
          }}
        />
        <StyledTooltip title={tooltip} placement="bottom" enterDelay={250}>
          <span className={"help-icon"}>
            <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
          </span>
        </StyledTooltip>
      </React.Fragment>
    );
  };

  const seed = () => {
    if (state.RandomOptions.Seed >= 1) {
      return state.RandomOptions.Seed;
    } else {
      return newSeed();
    }
  };

  const newSeed = () => {
    return Math.round(Math.random() * Number.MAX_SAFE_INTEGER);
  };

  const seedInput = () => {
    if (state.RandomOptions.UseSeed) {
      return (
        <React.Fragment>
          <InputNumber
            aria-label="Seed number input"
            min={1}
            max={Number.MAX_SAFE_INTEGER}
            style={{ width: 100 }}
            value={state.RandomOptions.Seed}
            onChange={(value) => {
              return updateRandomOptions(state, "Seed", value);
            }}
          />
        </React.Fragment>
      );
    }
  };

  const qolOptions = () => {
    return (
      <Grid container>
        <Typography
          variant="h4"
          align={"center"}
          className={"HeaderText2"}
          gutterBottom
        >
          Quality of Life
        </Typography>

        <Grid container>
          <Grid item xs={6}>
            {mkCheckbox({
              key: "EnableTownSkills",
              tooltip: "Enable the ability to use all skills in town.",
            })}
          </Grid>
          <Grid item xs={6}>
            {mkCheckbox({
              key: "StartWithCube",
              tooltip: "Newly created characters will start with a cube.",
            })}
          </Grid>
        </Grid>

        <Grid container>
          <Grid item xs={6}>
            {mkCheckbox({
              key: "Cowzzz",
              tooltip:
                "Enables the ability to recreate a cow portal after killing the cow king.  Adds cube recipe to cube a single tp scroll to create the cow portal4.",
            })}
          </Grid>
          <Grid item xs={6}>
            {mkCheckbox({
              key: "IncreaseStackSizes",
              tooltip:
                "Increases tome sizes to 100.  Increases arrows/bolts stack sizes to 511.  Increases key stack sizes to 100.",
            })}
          </Grid>

          <Grid item xs={6}>
            {mkCheckbox({
              key: "RemoveLevelRequirements",
              tooltip:
                "Removes level requirements from items.",
            })}
          </Grid>

          <Grid item xs={6}>
            {mkCheckbox({
              key: "RemoveAttRequirements",
              tooltip:
                "Removes stat requirements from items.",
            })}
          </Grid>

          <Grid item xs={6}>
            {mkCheckbox({
              key: "RemoveUniqCharmLimit",
              tooltip:
                "Removes unique charm limit in inventory.",
            })}
          </Grid>



        </Grid>
      </Grid>
    );
  };

  const dirOptions = () => {
    return (
      <Grid container>
        <Grid container spacing={5}>
          <Grid item xs={6}>
            <StyledTooltip
              title={
                "The path to the source directory containing the diablo 2 source files. Leave this blank to use 113c source files. example: C:/d2/data/global/excel/"
              }
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
            <TextField
              id="source-dir"
              label="Source Directory"
              value={state.SourceDir}
              onChange={(e) =>
                setState({ ...state, SourceDir: e.target.value })
              }
              fullWidth
            />
          </Grid>
          <Grid item xs={6}>
            <StyledTooltip
              title={
                "The directory that the data folder will be placed. Leave blank to use current directory. This requires a trailing slash. example: /Users/{username}/{folder}/"
              }
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
            <TextField
              id="output-dir"
              label="Output Directory"
              value={state.OutputDir}
              onChange={(e) =>
                setState({ ...state, OutputDir: e.target.value })
              }
              fullWidth
            />
          </Grid>
        </Grid>
      </Grid>
    );
  };

  const otherOptions = () => {
    return (
      <React.Fragment>
        <Typography
          variant="h4"
          align={"center"}
          className={"HeaderText2"}
          gutterBottom
          xs={12}
        >
          Other Awesome Options
        </Typography>
        <Grid item xs={12}>
          {mkCheckbox({
            key: "MeleeSplash",
            tooltip:
              "Enables Splash Damage.  Can spawn as an affix on magic and rare jewels.",
          })}
        </Grid>
        <Grid item xs={12} className={"SliderWrapper"}>
          <Typography
            id="IncreaseMonsterDensity"
            align={"center"}
            gutterBottom
            className={"primary"}
          >
            Increase Monster Density
            <StyledTooltip
              title={
                "Increases monster density throughout the map by the given factor."
              }
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
          </Typography>
          <Slider
            defaultValue={1}
            getAriaValueText={valuetext}
            aria-labelledby="IncreaseMonsterDensity"
            step={0.5}
            min={1}
            max={30}
            marks={[
              {
                value: 1,
                label: "Vanilla",
              },
              {
                value: 30,
                label: "Insanity",
              },
            ]}
            valueLabelDisplay="on"
            onChange={(e, n) => setState({ ...state, IncreaseMonsterDensity: n })}
          />
        </Grid>
      </React.Fragment>
    );
  };

  const dropRateOptions = () => {
    return (
      <React.Fragment>
        <Typography
          variant="h4"
          align={"center"}
          className={"HeaderText2"}
          gutterBottom
        >
          Drop Rates
        </Typography>
        <Grid container>
          <Grid item xs={4}>
            {mkCheckbox({
              key: "NoDropZero",
              tooltip: "Guarantees that a monster drops something upon death.",
            })}
          </Grid>
          <Grid item xs={4}>
            {mkCheckbox({
              key: "QuestDrops",
              tooltip: "Act bosses will always drop quest drops.",
            })}
          </Grid>
        </Grid>
        <Grid item xs={12} className={"SliderWrapper"}>
          <Typography
            id="UniqueItemDropRate"
            align={"center"}
            gutterBottom
            className={"primary slider-label"}
          >
            Unique Item Drop Rate
            <StyledTooltip
              title={
                "Increases the drop rate of unique and set items.  When using this setting, high values prevent some monsters from dropping set items."
              }
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
          </Typography>
          <Slider
            defaultValue={1}
            getAriaValueText={valuetext}
            aria-labelledby="UniqueItemDropRate"
            step={0.5}
            min={1}
            max={100}
            valueLabelDisplay="on"
            onChange={(e, n) => setState({ ...state, UniqueItemDropRate: n })}
          />
        </Grid>
        <Grid item xs={12} className={"SliderWrapper"}>
          <Typography
            id="RuneDropRate"
            align={"center"}
            gutterBottom
            className={"primary slider-label"}
          >
            Rune Drop Rate
            <StyledTooltip
              title={
                "Increases rune drop rates. Each increase of 1 raises the drop rate of the highest runes by ~5% cumulatively. E.g. Zod is 12.5x more common at 50 (1/418), and 156x (1/33) more common at 100."
              }
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
          </Typography>
          <Slider
            defaultValue={1}
            getAriaValueText={valuetext}
            aria-labelledby="RuneDropRate"
            step={0.5}
            min={1}
            max={100}
            marks={[
              {
                value: 1,
                label: "Vanilla",
              },
              {
                value: 100,
                label: "Zod 1/33",
              },
            ]}
            valueLabelDisplay="on"
            onChange={(e, n) => setState({ ...state, RuneDropRate: n })}
          />
        </Grid>
      </React.Fragment>
    );
  };

  const randomOptions = () => {
    return (
      <React.Fragment>
        <Typography
          variant="h4"
          align={"center"}
          className={"HeaderText2"}
          gutterBottom
          xs={12}
        >
          Randomization
        </Typography>

        <Grid container>
          <Grid item xs={4}>
            {mkRandoCheckbox({
              key: "Randomize",
              tooltip: "Randomize all all uniques, sets, and runewords.",
            })}
          </Grid>
          <Grid item xs={8}>
            <React.Fragment>
              <FormControlLabel
                control={
                  <Checkbox
                    color="primary"
                    name={"UseSeed"}
                    value={state.RandomOptions["UseSeed"]}
                  />
                }
                label={"UseSeed"}
                checked={state.RandomOptions["UseSeed"]}
                onChange={(e, checked) => {
                  return setState(
                    updateRandomOptions(
                      updateRandomOptions(state, "UseSeed", checked),
                      "Seed",
                      checked ? seed() : -1
                    )
                  );
                }}
              />
              <StyledTooltip
                title={
                  "Provide a specific seed to use.  Toggling on/off will generate a new seed."
                }
                placement="bottom"
                enterDelay={250}
              >
                <span className={"help-icon"}>
                  <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
                </span>
              </StyledTooltip>
            </React.Fragment>
            {seedInput()}
          </Grid>
        </Grid>

        <Grid container>
          <Grid item xs={4}>
            {mkRandoCheckbox({
              key: "UseOSkills",
              tooltip: "Change class only skill props to spawn as oskills.",
            })}
          </Grid>
          <Grid item xs={4}>
            {mkRandoCheckbox({
              key: "PerfectProps",
              tooltip:
                "All props will have a perfect max value when spawning on an item.",
            })}
          </Grid>
        </Grid>

        <Grid container>
          <Grid item xs={4}>
            {mkRandoCheckbox({
              key: "AllowDupProps",
              tooltip:
                "If turned off, prevents the same prop from being placed on an item more than once. e.g. two instances of all resist will not get stacked on the same randomized item.",
            })}
          </Grid>
          <Grid item xs={4}>
            {mkRandoCheckbox({
              key: "IsBalanced",
              tooltip:
                "Allows props only from items within 10 levels of the base item so that you don't get crazy hell stats on normal items, but still get a wide range of randomization.",
            })}
          </Grid>
        </Grid>

        <Grid container>
          <Grid item xs={12}>
            {mkRandoCheckbox({
              key: "BalancedPropCount",
              tooltip:
                "Pick prop count on items based on counts from vanilla items. Picks from items up to 10 levels higher when randomizing.",
            })}
          </Grid>
        </Grid>
        <Grid item xs={12} className={"SliderWrapper"}>
          <Typography
            id="MinProps"
            align={"center"}
            gutterBottom
            className={"primary"}
          >
            MinProps
            <StyledTooltip
              title={"Minimum number of props an item can have."}
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
          </Typography>
          <Slider
            defaultValue={0}
            getAriaValueText={valuetext}
            aria-labelledby="MinProps"
            step={1}
            max={20}
            marks={propMarks}
            disabled={state.RandomOptions.BalancedPropCount}
            valueLabelDisplay="on"
            onChange={(e, n) =>
              setState(updateRandomOptions(state, "MinProps", n))
            }
          />
        </Grid>
        <Grid item xs={12} className={"SliderWrapper"}>
          <Typography
            id="MaxProps"
            gutterBottom
            align={"center"}
            className={"primary"}
          >
            MaxProps
            <StyledTooltip
              title={"Maximum number of props an item can have."}
              placement="bottom"
              enterDelay={250}
            >
              <span className={"help-icon"}>
                <HelpOutlineOutlinedIcon></HelpOutlineOutlinedIcon>
              </span>
            </StyledTooltip>
          </Typography>
          <Slider
            defaultValue={20}
            getAriaValueText={valuetext}
            aria-labelledby="min-num-props"
            step={1}
            max={20}
            marks={propMarks}
            disabled={state.RandomOptions.BalancedPropCount}
            valueLabelDisplay="on"
            onChange={(e, n) =>
              setState(updateRandomOptions(state, "MaxProps", n))
            }
          />
        </Grid>
      </React.Fragment>
    );
  };

  return (
    <div className="D2ModMakerContainer">
      <Grid container alignItems={"center"} className={"HeaderText MainHeader"}>
        <Badge badgeContent={state.Version} color="primary">
          <Typography variant={"h2"}>D2 Mod Maker</Typography>
        </Badge>
      </Grid>
      <ButtonGroup
        size="large" color="primary" aria-label="large outlined primary button group"
        fullWidth={true}
        className="btns"
      >
        <Button
          onClick={async () => {
            let cfg = await loadConfig();
            console.log(cfg);
            setState({ ...cfg });
          }}
        >
          Load Config
        </Button>
        <Button
          onClick={() => saveConfig(state)}
        >
          Save Config
        </Button>
        <Button
          onClick={() => makeRunRequest(state)}
        >
          Run
        </Button>
      </ButtonGroup>
      {/* <Grid container spacing={10}>
        <Grid item xs={4}>
          <Button
            variant="contained"
            color="primary"
            className={"run-btn"}
            onClick={() => {
              makeRunRequest(state);
            }}
          >
            Run
      </Button>
        </Grid> */}

      <div className={"D2ModMakerContainerInner"}>
        {divider()}
        <Grid container>{randomOptions()}</Grid>
        {divider()}
        <Grid container>{otherOptions()}</Grid>
        {divider()}
        <Grid container> {qolOptions()}</Grid>
        {divider()}
        <Grid container>{dropRateOptions()}</Grid>
        {divider()}
        <Grid container>{dirOptions()}</Grid>
        {divider()}
        {/*<pre id={"state"}>{JSON.stringify(state, null, 2)}</pre>*/}
      </div>
    </div >
  );
}

function valuetext(value) {
  return `${value}`;
}

const propMarks = [
  {
    value: 0,
    label: "0",
  },
  {
    value: 7,
    label: "Runewords",
  },
  {
    value: 12,
    label: "Uniques",
  },
  {
    value: 19,
    label: "Sets",
  },
];

const StyledTooltip = withStyles((theme) => ({
  tooltip: {
    backgroundColor: "#f5f5f9",
    borderColor: "#f5f5f9",
    color: "rgba(0, 0, 0, 0.87)",
    maxWidth: 300,
    fontSize: theme.typography.pxToRem(12),
    fontWeight: 800,
  },
}))(Tooltip);

const divider = () => {
  return (
    <div className={"divider"}>
      <Divider></Divider>
    </div>
  );
};

async function makeRunRequest(data) {
  const headers = {
    "Content-Type": "text/plain",
  };

  await axios
    .post("http://localhost:8148/api/run", data, { headers })
    .then((response) => {
      console.log("Success ========>", response);
    })
    .catch((error) => {
      console.log("Error ========>", error);
    });
}

function saveConfig(data) {
  console.log(data);
  const headers = {
    "Content-Type": "text/plain",
  };

  axios
    .post("http://localhost:8148/api/cfg", data, { headers })
    .then((response) => {
      console.log("Success ========>", response);
    })
    .catch((error) => {
      console.log("Error ========>", error);
    });
}
