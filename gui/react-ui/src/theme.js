
import { red } from '@material-ui/core/colors';
import { createMuiTheme } from '@material-ui/core/styles';

// A custom theme for this app
const theme = createMuiTheme({
    palette: {
        type: "dark",
        primary: {
            main: 'rgb(103,15,18)',
        },
        secondary: {
            main: 'rgb(55,55,55)',
        },
        error: {
            main: red.A400,
        },
        background: {
            default: 'black',
        },
    },
});

export default theme;