import React from "react";
import Weather from "./weather/weather";
import News from "./news/news";
import About from "./about/about";
import config from "../config";
import './app.sass';

class App extends React.Component {

    state = {
        selected: "weather",
        location: null,
        body: <h3>{config.start.weather}</h3>,
        counter: 0

    }

    constructor(props) {
        super(props);
        this.setLocation = this.setLocation.bind(this);
        this.showWeather = this.showWeather.bind(this);
        this.showNews = this.showNews.bind(this);
        this.showAbout = this.showAbout.bind(this);
        this.getHighlight = this.getHighlight.bind(this);
        this.Header = this.Header.bind(this);
        this.ToolBar = this.ToolBar.bind(this);
    }

    // After the location is set
    // Refresh the current tab with the updated location
    setLocation() {
        let val = document.getElementById("inputZip").value;
        let bodySelected = null;
        let selected = this.state.selected;


        if (selected === "weather") {
            bodySelected = <Weather key={this.state.counter} location={val}/>;
        } else if (selected === "news") {
            bodySelected = <News key={this.state.counter} location={val}/>;
        } else if (selected === "about") {
            bodySelected = <About key={this.state.counter} location={val}/>;
        }


        this.setState({
            location: val,
            body: bodySelected,
            counter: this.state.counter + 1
        });


    }

    // Switch to the weather tab
    showWeather() {
        this.setState({
            selected: "weather",
            body: <Weather location={this.state.location}/>
        })
    }

    // Switch to the news tab
    showNews() {
        this.setState({
            selected: "news",
            body: <News location={this.state.location}/>
        })
    }

    // Switch to the about tab
    showAbout() {
        this.setState({
            selected: "about",
            body: <About location={this.state.location}/>
        })
    }

    // Get a background if the current tab is selected
    getHighlight(val) {
        if (val === this.state.selected) {
            return "#7CBAB5";
        }
        return "transparent";
    }

    // Return the header
    Header() {
        return (
            <div className="header">
                <div className="title">
                    <div className="imgContainer">
                        <img className="logo" alt=""
                         src="https://d29fhpw069ctt2.cloudfront.net/clipart/96671/preview/ryanlerch_simple_sun_motif_preview_57db.png"/>
                    </div>
                    <h1>Tenki</h1>
                </div>
                <div className="search">
                    <div className="inputContainer">
                        <input id="inputZip" className="inputZip" type="text"/>
                    </div>
                    <button className="buttonZip" onClick={this.setLocation}>Search</button>
                </div>
            </div>
        );
    }

    // Return the toolbar
    ToolBar() {
        return (
            <div className="toolbar">
                <button onClick={this.showWeather} style={{background: this.getHighlight("weather")}}>Weather</button>
                <button onClick={this.showNews} style={{background: this.getHighlight("news")}}>News</button>
                <button onClick={this.showAbout} style={{background: this.getHighlight("about")}}>About</button>
            </div>
        );
    }

    render() {
        return (
            <div className="mainContainer">
                <this.Header/>
                <this.ToolBar/>
                <div className="bodyContainer">
                    {this.state.body}
                </div>
            </div>);
    }
}

export default App;
