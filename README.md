# Lethal Mod Manager

The sexiest mod-manager for Lethal Company (Only for windows).

![image](https://github.com/KonstantinBelenko/lethal-mod-manager/assets/90444271/d584e916-4edc-41fa-82d0-c22d6a90095c)

A mod manager built with ***Love, Care, and Golang***

## Planned Features

- Automatic mod updates
- Syncronize mods with your friends
- Malware protection

## How to install

Just go to the [releases page](https://github.com/KonstantinBelenko/lethal-mod-manager/releases) and install & run the latest one. (Windows defender might not like it).

## Dev Installation

Clone the Repository & install all the required packages.

```sh
git clone https://github.com/KonstantinBelenko/lethal-mod-manager.git
cd lethalmodmanager

npm install
go mod tidy
```

### Run the project:

This will launch the go server using `serve_debug.go` and the react app separately, without packaging them together. This allows you to develop the frontend part without needing to constantly re-start the app itself.

```sh
npm run start
```

### Build the Project:

This will generate a new assets.go and package the react app together with go backend. It will output a fresh new binary to `bin/app.exe`.

```sh
npm run build
```

## Contributing

Any contributions to Lethal Mod Manager are welcome.

## Shout out to

- [Thunderstore](https://thunderstore.io/), for providing a platform for everyone to enjoy mods.
- [v0.dev](https://v0.dev/), for helping me to start with the project's design.
- [OpenAI chatgpt](https://chat.openai.com/), for letting me focus on the important details.
- [Shadcn/ui](https://ui.shadcn.com/docs/cli), for creatign such a beautiful library of reusable components that we're using now.

## License

Lethal Mod Manager is licensed under the MIT License.
