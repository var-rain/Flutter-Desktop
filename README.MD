## Flutter Desktop
### Special thanks
[Go Flutter desktop embedder](https://github.com/Drakirus/go-flutter-desktop-embedder)
### Screenshot
![hot_reload](screenshot/Screenshot_0.gif)
![hot_reload](screenshot/Screenshot_1.gif)
### Feature
- Support Hot Reload. (Test Only Visual Studio Code)
- Support Windows and Linux and Mac OS. ([Go Flutter desktop embedder](https://github.com/Drakirus/go-flutter-desktop-embedder))
- Support window title.
- Support Text input.
- Support Clipboard. (through shortcuts and UI)
- Support Configuration file.
- More...
### Deploy
- Clone this project.
- Edit main method.
```dart
import 'package:flutter/foundation.dart'
    show debugDefaultTargetPlatformOverride;

void main() {
    // must
    debugDefaultTargetPlatformOverride = TargetPlatform.android;
    runApp(MyApp());
}
```
- Copy **assets**, **flutter_engine.dll**, **config.json**, **icudtl.dat**, **launcher.exe**, **run.bat** to flutter project dirctory. if you will be run as Linux or Mac OS, you can to [build target platform](https://github.com/Drakirus/go-flutter-desktop-embedder/blob/master/README.md).
- Run **run.bat**.
- Configuration Visual Studio Code Run configuration (launch.json).
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Flutter",
            "request": "attach",
            "deviceId": "flutter-tester",
            "observatoryUri": "http://127.0.0.1:50300/",
            "type": "dart"
        }
    ]
}
```
- Click Visual Studio Code Run button to enable hot reload.
### Configuration File
```json
{
    // Window default width.
    "WIDTH": 480,
    // Window default height.
    "HEIGHT": 800,
    // Pixel ratio.
    "RATIO": 1.0,
    // ICON image path.
    "ICON": "assets/icon.png",
    // Flutter assets path.
    "FLUTTER": "build/flutter_assets",
    // icudtl.dat file path.
    "ICU": "icudtl.dat",
    // VM args.
    "ARGS": [
        "--dart-non-checked-mode",
        "--observatory-port=50300"
    ]
}
```