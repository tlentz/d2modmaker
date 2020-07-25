import React from "react";
import "./Main.scss";
import { makeStyles } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import {
  FormGroup,
  FormControlLabel,
  Checkbox,
  Button,
  Grid,
} from "@material-ui/core";
import Slider from "@material-ui/core/Slider";

const defaultCfg = {
  MeleeSplash: true,
  IncreaseStackSizes: true,
  IncreaseMonsterDensity: 1,
  EnableTownSkills: true,
  NoDropZero: true,
  QuestDrops: true,
  UniqueItemDropRate: -1,
  RuneDropRate: -1,
  StartWithCube: true,
  Cowzzz: true,
  EnterToExit: true,
  RandomOptions: {
    Randomize: true,
    Seed: -1,
    IsBalanced: true,
    BalancedPropCount: true,
    AllowDuplicateProps: false,
    MinProps: -1,
    MaxProps: -1,
    UseOSkills: true,
    PerfectProps: false,
  },
};

function getStepContent(step) {
  return step;
}

export default function D2ModMaker() {
  const [state, setState] = React.useState(defaultCfg);

  const createCheckbox = (key) => (
    <FormControlLabel
      control={<Checkbox color="primary" name={key} value={state[key]} />}
      label={key}
      checked={state[key]}
      onChange={(e, checked) => {
        setState({ ...state, [key]: checked });
      }}
    />
  );

  const updateRandomOptions = (oldState, key, val) => {
    let randomOptions = oldState.RandomOptions;
    randomOptions[key] = val;
    return { ...oldState, RandomOptions: randomOptions };
  };

  const mkRandoCheckbox = (key) => {
    return (
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
    );
  };

  const randomOptions = () => {
    return (
      <Grid item direction="row">
        <Typography variant="h6" gutterBottom>
          Random Options
        </Typography>
        {mkRandoCheckbox("Randomize")}
        {mkRandoCheckbox("IsBalanced")}
        {mkRandoCheckbox("BalancedPropCount")}
        {mkRandoCheckbox("AllowDuplicateProps")}
        <Grid item>
          <Typography id="min-num-props" gutterBottom className={"primary"}>
            MinProps
          </Typography>
          <Slider
            defaultValue={0}
            getAriaValueText={valuetext}
            aria-labelledby="min-num-props"
            step={1}
            max={20}
            marks={propMarks}
            valueLabelDisplay="auto"
            onChange={(e, n) =>
              setState(updateRandomOptions(state, "MinProps", n))
            }
          />
        </Grid>
        <Grid item>
          <Typography id="MaxProps" gutterBottom className={"primary"}>
            MaxProps
          </Typography>
          <Slider
            defaultValue={20}
            getAriaValueText={valuetext}
            aria-labelledby="min-num-props"
            step={1}
            max={20}
            marks={propMarks}
            valueLabelDisplay="on"
            onChange={(e, n) =>
              setState(updateRandomOptions(state, "MaxProps", n))
            }
          />
        </Grid>
        {mkRandoCheckbox("UseOSkills")}
        {mkRandoCheckbox("PerfectProps")}
      </Grid>
    );
  };

  return (
    <div className="D2ModMakerContainer">
      <Button variant="contained" color="primary">
        Run
      </Button>
      {randomOptions()}
      <React.Fragment>
        {/*<Grid container spacing={3}>*/}
        {/*  <Grid item xs={12}>*/}
        {/*    {createCheckbox("MeleeSplash")}*/}
        {/*    {createCheckbox("IncreaseStackSizes")}*/}
        {/*    {createCheckbox("EnableTownSkills")}*/}
        {/*    {createCheckbox("NoDropZero")}*/}
        {/*    {createCheckbox("QuestDrops")}*/}
        {/*    {createCheckbox("UniqueItemDropRate")}*/}
        {/*    {createCheckbox("RuneDropRate")}*/}
        {/*    {createCheckbox("StartWithCube")}*/}
        {/*    {createCheckbox("Cowzzz")}*/}
        {/*  </Grid>*/}
        {/*</Grid>*/}
      </React.Fragment>
    </div>
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
