import React from "react";
import Popup from 'reactjs-popup';

import 'reactjs-popup/dist/index.css';
import "./dailyForecast.sass";
import HourlyForecast from "./hourlyForecast";

class DailyForecast extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            ...props.data,
            children: props.data.children
        };


    }

    // Renders code for the card and the detailed view table
    // The card is displayed on screen
    // After clicking the card a table is shown with detailed forecast
    render() {
        return (
            <>
                <Popup
                    trigger={

                        <div className="card">
                            <p><b>{this.state.name}</b></p>
                            <img src={this.state.icon} alt={this.state.name + " image"}/>
                            <p><b>{this.state.shortForecast}</b></p>
                            <p>{this.state.detailedForecast}</p>
                        </div>


                    }
                    modal
                    closeOnDocumentClick
                    lockScroll={true}
                    contentStyle={{
                        border: ".3rem #7CBAB5 solid",
                        maxWidth: "90%", maxHeight: "80%", minWidth: "300px", overflow: "auto",
                        position: "fixed", top: "50%", left: "50%", transform: "translate(-50%, -50%)"
                    }}

                >
                    <div className="popupContent">
                        <HourlyForecast children={this.state.children}/>
                    </div>
                </Popup>
            </>
        );
    }

}

export default DailyForecast