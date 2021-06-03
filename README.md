# Tenki Weather App

## Introduction
This repository contains code to a weather application. It shows weekly weather forecast and when clicking on a weekly forecast, it will show detailed forecast. It also shows news for that particular location. The front-end is in React.js and the back-end is lambda functions written in Go. The application is available live on [weather.nimitpatel.me](https://weather.nimitpatel.me).


## Folder Info
### **go**
* Contains code related to the Go functions
* api folder contains all the functions that are created
* Each folder (except api) in go directory is a package that can be used individually

### **public**
* Contains the main HTML file

### **src**
* Contains code related to React application

## Environment Variables
* **FAUNA_COLLECTION_NAME**: collection (table) name in Fauna database to be used
* **FAUNA_ITEM_NUM**: item number in the database collection to be retrieved
* **FAUNA_SECRET**: secret to access the database
* **MAPBOX_API**: MapBox geocode API to do forward geocoding; used to convert address to longitude and latitude
* **MAPBOX_API_KEY**: MapBox API key
* **NEWS_API**: NewsAPI used to get news data
* **NEWS_API_KEY**: NewsAPI key
* **NOAA_API**: NOAA API used to get weather data (no API key needed)