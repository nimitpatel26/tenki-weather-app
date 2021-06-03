import React from "react"
import config from "../../config";

import "./about.sass";

class About extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            title:null,
            details: null
        }

    }

    // Gets the data to display in the about section from the lambda function
    async componentDidMount() {

        try {
            let resp = await fetch(config.api.about);
            let data = await resp.json();

            let aboutDetails = data["Details"]
            aboutDetails = aboutDetails.map(detail => <p>{detail}</p>);

            this.setState({
                title: data["Title"],
                details: aboutDetails
            });
        }catch (e){
            console.log(e);
            this.setState({
                details: <h3>{config.error.about}</h3>
            });
        }


    }

    render() {
        return (
            <div className="aboutContainer">
                {this.state.details}
            </div>
        );
    }
}

export default About;