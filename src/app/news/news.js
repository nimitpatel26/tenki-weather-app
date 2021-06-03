import React from "react"
import config from "../../config";

import "./news.sass"

class News extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            location: props.location,
            articles: null

        }

    }

    // Gets the news data from the lambda function
    async getNewsData(){
        let resp = await fetch(config.api.news + "?location=" + this.state.location);
        let data = await resp.json();


        if (data.length === 0){
            throw new Error('No articles found for the given location!');
        }
        return data;
    }

    // Gets and formats the news data into a card format to display
    async componentDidMount() {
        if (this.state.location == null) {
            this.setState({
                articles: <h3>{config.error.news}</h3>
            });
            return;
        }

        try {
            let data = await this.getNewsData();
            let index = -1;
            data = data.map(article => {
                index++;
                return (
                    <div key={index} className="articleCard" onClick={() => {
                        window.open(article.url)
                    }}>
                        <img src={article.urlToImage} alt="unavailable" style={{width: "100%"}}/>
                        <h4>{article.title}</h4>
                        <h5>{article.author}</h5>
                        <p>{article.description}</p>
                    </div>
                );
            });


            this.setState({
                articles: data
            });
        }catch (e) {
            console.log(e);
            this.setState({
                articles: <h3>{config.error.news}</h3>
            });
        }

    }


    render() {
        return (
            <div className="articlesContainer" style={{display: "flex", flexFlow: "row wrap"}}>
                {this.state.articles}
            </div>
        );
    }
}

export default News;