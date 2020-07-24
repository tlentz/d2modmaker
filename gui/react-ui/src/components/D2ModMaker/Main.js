import React from 'react';
import "./Main.scss"
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Button from '@material-ui/core/Button';


const defaultCfg = {
    "MeleeSplash": true,
    "IncreaseStackSizes": true,
    "IncreaseMonsterDensity": 1,
    "EnableTownSkills": true,
    "NoDropZero": true,
    "QuestDrops": true,
    "UniqueItemDropRate": -1,
    "RuneDropRate": -1,
    "StartWithCube": true,
    "Cowzzz": true,
    "EnterToExit": true,
    "RandomOptions": {
        "Randomize": true,
        "Seed": -1,
        "IsBalanced": true,
        "BalancedPropCount": true,
        "AllowDuplicateProps": false,
        "MinProps": -1,
        "MaxProps": -1,
        "UseOSkills": true,
        "PerfectProps": false
    }
}


function getStepContent(step) {
    return step
}

export default function D2ModMaker() {
    const [state, setState] = React.useState(defaultCfg);

    const createCheckbox = (key) => (
        <FormControlLabel
            control={<Checkbox color="primary" name={key} value="true" />}
            label={key}
            checked={state[key]}
            onChange={() => setState({ [key] : !state[key] } ) }
        />
    );

    return (
        <div className="D2ModMakerContainer">
            <Button
                variant="contained"
                color="primary"
            >
                Run
            </Button>
            <React.Fragment>
                <Grid container spacing={3}>
                    <Grid item xs={12}>
                        {createCheckbox("MeleeSplash")}
                        {createCheckbox("IncreaseStackSizes")}
                        {createCheckbox("EnableTownSkills")}
                        {createCheckbox("NoDropZero")}
                        {createCheckbox("QuestDrops")}
                        {createCheckbox("UniqueItemDropRate")}
                        {createCheckbox("RuneDropRate")}
                        {createCheckbox("StartWithCube")}
                        {createCheckbox("Cowzzz")}
                    </Grid>
                </Grid>
            </React.Fragment>
        </div>
    );
}


