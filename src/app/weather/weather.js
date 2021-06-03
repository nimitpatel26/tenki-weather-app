import React from "react";

import DailyForecast from "./dailyForecast"
import "./weather.sass"
import config from "../../config";

class Weather extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            location: props.location
        }

    }

    // Gets data from the lambda function and maps it to the correct format
    async componentDidMount() {
        if (this.state.location == null) {
            this.setState({
                forecast: <h3>{config.error.weather}</h3>
            });
            return;
        }

        try {
            let resp = await fetch(config.api.weather + "?location=" + this.state.location);
            let data = await resp.json();


            let forecastData = []

            if (data == null || data.periods == null) {
                forecastData.push(<h3>Unable to get weather data</h3>);
            }

            for (let forecast of data.periods) {
                forecastData.push(<DailyForecast key={forecast.name} data={forecast}/>);

            }

            this.setState({
                forecast: forecastData
            });

        }catch (e) {
            console.log(e);
            this.setState({
                forecast: <h3>{config.error.weather}</h3>
            });
        }
    }


    render() {

        return (
            <>
                <div className="forecastContainer">
                    {this.state.forecast}
                </div>
            </>
        );
    }

}


export default Weather