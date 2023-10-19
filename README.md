# weather-cli
Weather CLI delivers daily or hourly forecasts to your terminal.

### Installation
1. clone the repo
2. if not using linux, you can run `build.sh` shell script to compile the Go code (requires Golang installed)
3. add `path/to/installation/bin` to your $PATH variable
4. register for free on weatherapi.com and export your API key to `WEATHER_API_KEY` environment variable (e.g. add `export WEATHER_API_KEY="<your_key>"` in `.bashrc` file)

### Usage
Typing in your terminal
```
weather
```
will display all the available commands.<br />
+ You can search for cities using `weather search <location>` (e.g. `weather search London`) to find the available cities.<br />
+ `weather daily <location>` will display a 5 day weather forecast.<br />
+ `weather hourly <location>` will display a 3 day weather forecast hour-by-hour.<br />
<br />
Have fun!
