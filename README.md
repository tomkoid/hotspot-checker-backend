# Hotspot Checker Backend 

This backend is used to check hotspot availability. [The firmware](https://codeberg.org/tomkoid/hotspot-checker-firmware) connects to a hotspot Wi-Fi and sends ping requests to a remote server each 5 seconds. If the remote backend server has not received any message in an amount of time, it will send a notification to the phone sharing the hotspot through [Ntfy](https://ntfy.sh/).

## How to use

To get started with the backend, clone the repository and make a .env file (example file in .env.example). Then you can just run `go run .` like in any other Go project.

Make sure to set your values such as backend password in the `.env` file.

## Roadmap

- [x] Add base functionality (e.g. receive ping requests)
- [x] Send notifications through [Ntfy](https://ntfy.sh/) 
- [x] Custom backend password 
- [x] .env configuration 

# License

This project is licensed under the MIT License. See LICENSE for more information.

Copyright 2024 Tomkoid
