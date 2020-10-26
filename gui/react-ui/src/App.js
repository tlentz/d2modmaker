import React from 'react';
import "./App.scss"
import D2ModMaker from "./components/D2ModMaker/Main"

export default function App() {
    return (
        <React.Fragment>
            <div className="AppContainer">
                <div className="AppContainerOuter">
                    <div className="AppContainerInner">
                        <D2ModMaker />
                    </div>
                </div>
            </div>
        </React.Fragment>
    );
}