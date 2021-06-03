/**
 * This file will contain different configuration information
 * that will be used throughout the application
 * */
const config = {
    api: {
        weather: "https://tenki.nimitpatel.me/.netlify/functions/weather",
        news: "https://tenki.nimitpatel.me/.netlify/functions/news",
        about: "https://tenki.nimitpatel.me/.netlify/functions/about",
    },
    start:{
        weather: "Please enter a valid location to get started."
    },
    error: {
        weather: "Please enter a valid location to see weather data.",
        news: "Please enter a valid location to see news data.",
        about: "Could not fetch data."
    }
};

module.exports = config;