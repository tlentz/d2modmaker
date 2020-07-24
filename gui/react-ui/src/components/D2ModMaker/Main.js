import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Paper from '@material-ui/core/Paper';
import Stepper from '@material-ui/core/Stepper';
import Step from '@material-ui/core/Step';
import StepLabel from '@material-ui/core/StepLabel';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import "./Main.scss"

const steps = ['Shipping address', 'Payment details', 'Review your order'];

function getStepContent(step) {
    return step
}

export default function D2ModMaker() {
    const [activeStep, setActiveStep] = React.useState(0);

    const handleNext = () => {
        setActiveStep(activeStep + 1);
    };

    const handleBack = () => {
        setActiveStep(activeStep - 1);
    };

    const handleRun = () => {
        console.log("run")
    }

    return (
        <div className="D2ModMakerContainer">
            <Button
                variant="contained"
                color="primary"
                onClick={handleRun}
            >
                Run
            </Button>
        </div>
    );
}