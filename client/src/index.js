// pull in desired CSS/SASS files

// import "./scss/tailwind.css";
require('./scss/main.scss');

// inject bundled Elm app into div#main
const { Elm } = require('./elm/Main.elm');
const app = Elm.Main.init({
  node: document.getElementById('main')
});