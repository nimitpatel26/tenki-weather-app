import React from 'react';
import "./hourlyForecast.sass";

class HourlyForecast extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            children: props.children
        };
    }

    // Creates table rows for detailed forecast
    getTableRows() {
        let tableRows = [];

        let index = 0;
        for (let child of this.state.children) {
            let startTime = new Date(Date.parse(child.startTime));
            let startHours = startTime.getHours();
            let startAmPm = startHours >= 12 ? 'PM' : 'AM';
            startHours = startHours % 12;
            startHours = startHours ? startHours : 12;

            tableRows.push(
                <tbody key={index}>
                <tr>
                    <td><img src={child.icon} alt="weather-icon"/></td>
                    <td>{startHours} {startAmPm}</td>
                    <td>{child.shortForecast}</td>
                    <td>{child.temperature} {child.temperatureUnit}</td>
                    <td>{child.windSpeed}</td>
                </tr>
                </tbody>
            );
            index++;
        }
        return tableRows;
    }

    // Renders a table that will be a popup for detailed forecast
    render() {

        if (this.state.children == null || this.state.children.length === 0) {
            return <h3 key="failed">Unable to get detailed weather data</h3>;
        }


        let tableHeader = (
            <thead>
            <tr>
                <th/>
                <th>Time</th>
                <th>Forecast</th>
                <th>Temp</th>
                <th>Wind</th>
            </tr>
            </thead>
        );


        return (
            <table>
                {tableHeader}
                {this.getTableRows()}
            </table>
        );
    }


}

export default HourlyForecast;