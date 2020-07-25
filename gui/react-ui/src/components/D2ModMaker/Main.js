import React from "react";
import "./Main.scss";
import { withStyles, makeStyles } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import { ToggleButtonGroup, ToggleButton } from "@material-ui/lab";
import {
  FormGroup,
  FormControlLabel,
  Checkbox,
  Button,
  Grid,
  Tooltip,
  ButtonGroup,
  InputAdornment,
  IconButton,
} from "@material-ui/core";
import Slider from "@material-ui/core/Slider";
import InputNumber from "rc-input-number";
import RefreshIcon from "@material-ui/icons/Refresh";
import HelpOutlineOutlinedIcon from "@material-ui/icons/HelpOutlineOutlined";
import HelpIcon from "@material-ui/icons/Help";
import Badge from "@material-ui/core/Badge";
import Divider from "@material-ui/core/Divider";
const defaultCfg = {
  Version: "v0.5.0",
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
    UseSeed: false,
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

  const seed = () => {
    if (state.RandomOptions.Seed >= 1) {
      return state.RandomOptions.Seed;
    } else {
      return newSeed();
    }
  };

  const newSeed = () => {
    console.log(state.RandomOptions.Seed);
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
            value={seed()}
            onChange={(value) => {
              return updateRandomOptions(state, "Seed", value);
            }}
          />
        </React.Fragment>
      );
    }
  };

  const propCounts = () => {};

  const randomOptions = () => {
    return (
      <Grid item>
        <Typography
          variant="h6"
          align={"center"}
          className={"HeaderText2"}
          gutterBottom
        >
          Random Options
        </Typography>

        <Grid container>
          <Grid item xs={4}>
            {mkRandoCheckbox({
              key: "Randomize",
              tooltip: "Randomize all all uniques, sets, and runewords.",
            })}
          </Grid>
          <Grid item xs={8}>
            {mkRandoCheckbox({
              key: "UseSeed",
              tooltip:
                "Provide a specific seed to use.  Toggling on/off will generate a new seed.",
            })}
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
          <Grid item xs={6}>
            {mkRandoCheckbox({
              key: "AllowDuplicateProps",
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
          <Grid item xs={6}>
            {mkRandoCheckbox({
              key: "BalancedPropCount",
              tooltip:
                "Pick prop count on items based on counts from vanilla items. Picks from items up to 10 levels higher when randomizing.",
            })}
          </Grid>
        </Grid>

        <Grid item>
          <Typography
            id="min-num-props"
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
            aria-labelledby="min-num-props"
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
        <Grid item>
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
      </Grid>
    );
  };

  return (
    <div className="D2ModMakerContainer">
      <Grid container alignItems={"center"} className={"HeaderText"}>
        <Badge badgeContent={state.Version} color="primary">
          <Typography variant={"h2"}>D2 Mod Maker</Typography>
        </Badge>
      </Grid>
      <Button variant="contained" color="primary" className={"run-btn"}>
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
