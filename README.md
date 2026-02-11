Cross-platform project for remote computer control via a telegram bot

## Build
Specify NAME yourself (for Windows, add .exe extension to name)
```
git clone https://github.com/Konare1ka/TG-commander
cd TG-commander
go build -o NAME ./src
```

## Settings
In same directory where executable file is located there should be config.json and plugins directory with plugins

config.yaml must contain a field with:
- telegram bot token, which is issued by [BotFather](https://telegram.me/BotFather)
- your username (without @) to receive admin privileges
- plugins available to everyone (optional)
- path to the directory with plugins (optional)

## Usage 
To simple use, run binary without additional arguments

The help command is also available `./bin -h`\

Messages may arrive with a delay, before you decide that nothing is working for you, wait a minute

## Plugins
Plugins can be any executable files and must be called via the telegram bot /plugin_name
>Ex. file `help.py` - command in bot `/help`

Plugins can only be accessed by the user specified in config. However, publicly accessible plugins can be specified in third config field,
> like ["help", "getFile", "music"]
> 

### Return value of plugins
Any output in the program will be transferred to the telegram bot. To send any media or document (including .jar, .zip, .exe, etc.), you need to send the message type (image-img, video-vid, audio-aud, document-doc) followed by a space and full path to file. 

>Ex. img /home/user/Pictures/image.png
>

### Usage arguments
You can pass arguments to the plugin; to do this, when calling, you need to specify the arguments separated by a space
`/plugin arg1 arg2`
And to process arguments in the plugin you need to use %1 - arg1, %2 - arg2, etc.
**I recommend using a check for the presence of these arguments, and also checking whether everything was done as it should. Since the program does not monitor how the script was executed, it only monitors whether it was executed**
>An example of using arguments is in the examples, the name of the script is shareFiles
>

## Service

-s, --service to open service interface


