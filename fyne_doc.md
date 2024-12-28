# https://docs.fyne.io/

Toggle navigation

- [Docs](https://docs.fyne.io)

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Fyne toolkit documentation for developers

This site is home to the documentation and examples for developers working with the Fyne toolkit. We have details for people just getting started, building their first graphical app through to detailed walkthroughs of complex topics related to building leading cross platform apps.

If you can't wait to start building your first Fyne app, you should follow our getting started guide.

[Get Started](/started/)

You can also check out the latest [API documentation](/api/), or if you came here with a question in mind then you may find an answer in the [FAQs](/faq/).

If you are new to the Go language, we recommend running through the Go tour before returning to the Fyne documentation.

There is also an [introduction to GUI development](/guiintro/) if this is your first time building an application.

For developers who prefer to learn from videos we have a collection of getting started tutorials on YouTube.

If you have questions about the videos or want to let us know that one is missing then please .

## Installing

If you have not used Fyne before then the following steps will get you up and running:

1. Install Go
2. Install Gcc
3. `go get fyne.io/fyne/v2@latest` (or, if using Go before 1.16, then `go get fyne.io/fyne/v2`)
4. You can test your installation using the app.

For more details including the exact commands for your operating system see the [getting started](/started/) page.

## Diving Right In

To explore more detailed topics you can check out the [explore](/explore/) section of this site. If you would like to request additional examples or some support for your project you can .

If you like the Fyne toolkit please consider our work.

# https://docs.fyne.io/started/

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Getting Started

Using the Fyne toolkit to build cross platform applications is very simple but does require some tools to be installed before you can begin. If your computer is set up for development with Go then the following steps may not be required, but we advise reading the tips for your operating system just in case. If later steps in this tutorial fail then you should re-visit the prerequisites below.

## Prerequisites

Fyne requires 3 basic elements to be present, the Go tools (at least version 1.12), a C compiler (to connect with system graphics drivers) and a system graphics driver. The instructions vary depending on your operating system, choose the appropriate tab below for installation instructions.

Note that these steps are just required for development - your Fyne applications will not require any setup or dependency installation for end users!

- Windows
- macOS X
- Linux
- Raspberry Pi
- BSD
- Android
- iOS
- Termux (create Android apk without PC - on Android)

The MSYS2 is the recommended approach for working on windows, proceed as follows:

1. Install MSYS2 from
2. Once installed do not use the MSYS terminal that opens
3. Open “MSYS2 MinGW 64-bit” from the start menu
4. Execute the following commands (if asked for install options be sure to choose “all”):

```
 $ pacman -Syu
 $ pacman -S git mingw-w64-x86_64-toolchain mingw-w64-x86_64-go

```

5. You will need to add ~/Go/bin to your $PATH, for MSYS2 you can paste the following command into your terminal:

```
 $ echo "export PATH=\$PATH:~/Go/bin" >> ~/.bashrc

```

6. For the compiler to work on other terminals you will need to set up the windows %PATH% variable to find these tools. Go to the “Edit the system environment variables” control panel, tap “Advanced” and add “C:\msys64\mingw64\bin” to the Path list.

The MSYS2 is the recommended approach for working on windows, proceed as follows:

1. Install MSYS2 from
2. Once installed do not use the MSYS terminal that opens
3. Open “MSYS2 MinGW 64-bit” from the start menu
4. Execute the following commands (if asked for install options be sure to choose “all”):

```
 $ pacman -Syu
 $ pacman -S git mingw-w64-x86_64-toolchain mingw-w64-x86_64-go

```

5. You will need to add ~/Go/bin to your $PATH, for MSYS2 you can paste the following command into your terminal:

```
 $ echo "export PATH=\$PATH:~/Go/bin" >> ~/.bashrc

```

6. For the compiler to work on other terminals you will need to set up the windows %PATH% variable to find these tools. Go to the “Edit the system environment variables” control panel, tap “Advanced” and add “C:\msys64\mingw64\bin” to the Path list.

1. Download Go from the and follow instructions
1. Install Xcode from the
1. Set up the Xcode command line tools by opening a Terminal window and typing the following: `xcode-select --install`
1. In macOS the graphics drivers will already be installed.

- You will need to install Go, gcc and the graphics library header files using your package manager, one of the following commands will probably work.
- **Debian / Ubuntu:** `sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev`
- **Fedora:** `sudo dnf install golang golang-misc gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel`
- **Arch Linux:** `sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi`
- **Solus:** `sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel`
- **openSUSE:** `sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel`
- **Void Linux:** `sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel libXxf86vm-devel`
- **Alpine Linux** `sudo apk add go gcc libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers mesa-dev`
- **NixOS** `nix-shell -p libGL pkg-config xorg.libX11.dev xorg.libXcursor xorg.libXi xorg.libXinerama xorg.libXrandr xorg.libXxf86vm`

- You will need to install Go, gcc and the graphics library header files using the package manager.
- `sudo apt-get install golang gcc libegl1-mesa-dev xorg-dev`

- You will need to install Go, gcc and the graphics library header files using the package manager.
- **FreeBSD:** `sudo pkg install go gcc xorg pkgconf`
- **OpenBSD:** `sudo pkg_add go`
- **NetBSD:** `sudo pkgin install go pkgconf`

- To develop apps for Android you will first need to install the tools for your current computer (Windows, macOS or Linux)
- Once complete you will need to install the Android SDK and Android NDK - the recommended approach is to install and then go to **Tools > SDK Manager** and from **SDK Tools** install the **NDK (Side by side)** package.
- Alternatively you can download the which is a more lean approach. Extract the folder and point the `ANDROID_NDK_HOME` environment variable to it.

- To develop apps for iOS you will need access to an Apple Mac computer, configured according to the **macOS** tab above.
- You will also need to create an and sign up to the developer program (costs apply) to obtain the necessary certificate to run your app on any devices.

Compiling Fyne apps on Android you will need Android 9 or above

- Install and from there
- Open Termux and install Go and Git: `pkg install golang git`
- Install NDK and SDK to termux from and set enviroment variables `ANDROID_HOME` and `ANDROID_NDK_HOME`

(you can get more help from )

## Downloading

Since Go 1.16 you will need to set up the module before you can use the package.

Run the following command and replace `MODULE_NAME` with your preferred module name (this should be called in a new folder specific for your application).

```
$ mkdir myapp
$ cd myapp
$ go mod init MODULE_NAME

```

You now need to download the Fyne module and helper tool. This will be done using the following commands:

```
$ go get fyne.io/fyne/v2@latest
$ go install fyne.io/fyne/v2/cmd/fyne@latest

```

If you are unsure of how Go modules work, consider reading .

## Check your installation

Before coding an app or running an example you can check your install using the tool. Simply download the right app for your computer from the link and run it, you should see something like the following screen:

![](/images/started/setup.png)

If there are any problems with your installation see the [troubleshooting](/faq/troubleshoot) section for hints.

## Run the demo

If you want to see the Fyne toolkit in action before you start to code your own application, you can see our running on your computer by executing:

```
$ go run fyne.io/fyne/v2/cmd/fyne_demo@latest

```

Please note that the first run has to compile some C-code and can thus take longer than usual. Subsequent builds reuse the cache and will be much faster.

### Installing

If you want to, you can also install the demo using the following command (requires Go 1.16 or later):

```
$ go install fyne.io/fyne/v2/cmd/fyne_demo@latest

```

If your `GOBIN` environment has been added to path (should be by default on macOS and Windows), you can then run the demo:

```
$ fyne_demo

```

And that’s all there is to it! Now you can write your own Fyne application in your IDE of choice. If you want to see some Fyne code in action then you can read [your first application](/started/firstapp).

# https://docs.fyne.io/started/hello

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Creating your first Fyne app

Having completed the steps in the [getting started](/started/) document you’re ready to build your first app. To illustrate the process we will build a simple hello world application.

A simple app starts by creating an app instance with app.New() and then opening a window with app.NewWindow(). Then a widget tree is defined that is set as the main content with SetContent() on a window. The app UI is then shown by calling ShowAndRun() on the window.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}

```

The code above can be built using the command `go build .` and then executed either by running the `hello` command or by double clicking the icon. You could also bypass the compiling step and just run the code directly using `go run .`.

Either approach will show a window that looks just like this:

![Hello Window](/images/started/hello-dark.png)

If you prefer a light theme then just set the environment variable `FYNE_THEME=light` and you’ll get:

![Hello Light Theme](/images/started/hello-light.png)

That’s all there is to getting started. To learn more you can read the full .

Fyne

Search

Watch later

Share

Copy link

Info

Shopping

Tap to unmute

If playback doesn't begin shortly, try restarting your device.

Share

Include playlist

An error occurred while retrieving sharing information. Please try again later.

0:00

0:00 / 3:51•Live

•

# https://docs.fyne.io/started/demo

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Run Fyne Demo

If you want to see the Fyne toolkit in action before you start to code your own application, you can see our .

### Running

If you want to, you can run the demo directly using the following command (requires Go 1.16 or later):

```
go run fyne.io/fyne/v2/cmd/fyne_demo@latest

```

For earlier versions of Go, you need to use the following command instead:

```
go run fyne.io/fyne/v2/cmd/fyne_demo

```

By browsing the different tabs of the app you can see all the features of Fyne toolkit.

### Installing

It is possible to install the app as a graphical application like all others on your computer. We have the helpful `fyne` tool to do this for you. First you will need to install the tool:

```
go install fyne.io/fyne/v2/cmd/fyne@latest

```

After that you can simply package and install the demo app:

```
fyne get fyne.io/fyne/v2/cmd/fyne_demo

```

After this step you can find “Fyne Demo” in your app launcher.

### Exploring the code

If you are interested in any of the features you should check out the or join one of the .

# https://docs.fyne.io/started/apprun

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Application and RunLoop

For a GUI application to work it needs to run an event loop (sometimes called a runloop) that processes user interactions and drawing events. In Fyne this is started using the `App.Run()` or `Window.ShowAndRun()` functions. One of these must be called from the end of your setup code in the `main()` function.

An application should only have one runloop and so you should only call `Run()` once in your code. Calling it a second time will cause errors.

```
package main
import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))
	myWindow.Show()
	myApp.Run()
	tidyUp()
}
func tidyUp() {
	fmt.Println("Exited")
}

```

For desktop runtimes an app can be quit directly by calling `App.Quit()` (mobile apps do not support this) - normally not needed in developer code. An application will also quit once all the windows are closed. See also that functions executed after `Run()` will not be called until the application exits.

# https://docs.fyne.io/started/updating

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Updating Content in your GUI

Having completed the [hello world](/started/hello) tutorial or other examples you will have created a basic user interface. In this page we see how the content of a GUI can be updated from your code.

The first step is to assign the widget you want to update to a variable. In the hello world tutorial we passed `widget.NewLabel` directly into `SetContent()`, to update it we change that to two different lines, such as:

```
	clock := widget.NewLabel("")
	w.SetContent(clock)

```

Once the content has been assigned to a variable we can call functions like `SetText("new text")`. For our example we will set the content of our label to the current time, with the help of `Time.Format`.

```
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)

```

That is all we need to do to change content of a visible item (see below for the full code). However, we can go further and update content on a regular basis.

## Running in the background

Most applications will need to have processes that run in the background, for example downloading data or responding to events. To simulate this we will extend the code above to run every second.

Like with most go code we can create a goroutine (using the `go` keyword) and run our code there. If we move the text update code to a new function it can be called on initial display as well as on a timer for regular updating. By combining a goroutine and the `time.Tick` inside a for loop we can update the label every second.

```
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

```

It is important to place this code before `ShowAndRun` or `Run` calls because they will not return until the application closes. With all of this together the code will run and update the user interface each second, creating a basic clock widget. The full code is as follows:

```
package main
import (
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}

```

Share

Include playlist

An error occurred while retrieving sharing information. Please try again later.

0:00

0:00 / 5:00•Live

•

# https://docs.fyne.io/started/windows

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Window Handling

Windows are created using `App.NewWindow()` and need to be shown using the `Show()` function. The helper method `ShowAndRun()` on `fyne.Window` allows you to show your window and run the application at the same time.

By default a window will be the right size to show its content by checking the `MinSize()` function (more on that in later examples). You can set a larger size by calling the `Window.Resize()` method. Into this is passed a `fyne.Size` which contains a width and height using device independent pixels (meaning that it will be the same across different devices), for example to make a window square by default we could:

```
	w.Resize(fyne.NewSize(100, 100))

```

Be aware that the desktop environment may have constraints that cause windows to be smaller than requested. Mobile devices will typically ignore this as they are only displayed at full-screen.

If you wish to show a second window you must only call the `Show()` function. It can also be helpful to split `Window.Show()` from `App.Run()` if you want to open multiple windows when your application starts. The example below shows how to load two windows when starting.

```
package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()
	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()
	a.Run()
}

```

The above application will exit when both windows are closed. If your app is arranged so one window is main and the others are accessory views you can set one window to be “master” so that the app exits if that window is closed. To do this use the `SetMaster()` function on `Window`.

Windows can be created at any time, we could change the code above so that the content of the second window (`w2`) is a button that opens a new window. You could also load windows from more complex workflows, but be careful because new windows will normally appear above the current active content.

```
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))

```

Share

Include playlist

An error occurred while retrieving sharing information. Please try again later.

0:00

0:00 / 4:28•Live

•

# https://docs.fyne.io/started/testing

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Testing Graphical Apps

Part of a good test suite is being able to quickly write tests and run them on a regular basis. Fyne’s API is designed to make testing applications easy. By separating component logic from its rendering definition we can load applications without actually displaying them and test the functionality completely.

### Example

We can demonstrate unit testing by extending our [Hello World](/started/hello) app to include space for users to input their name to be greeted. We start by updating the user interface to have two elements, a `Label` for the greeting and an `Entry` for the name input. We display them, one above another, using `container.NewVBox` (a vertical box container). The updated user interface code will look as follows:

```
func makeUI() (*widget.Label, *widget.Entry) {
	return widget.NewLabel("Hello world!"),
		widget.NewEntry()
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")
	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}

```

To test this input behaviour we create a new file (with a name ending `_test.go` to mark it as tests) that defines a `TestGreeter` function.

```
package main
import (
	"testing"
)
func TestGreeting(t *testing.T) {
}

```

We can add an intial test that verifies the initial state, to do this we test the `Text` field of the `Label` that is returned from `makeUI` and error the test if it is not correct. Add the following code to your test method:

```
	out, in := makeUI()
	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

```

This test will pass - next we add to the test to validate the greeter. We use the Fyne `fyne.io/fyne/v2/test` package which assists in test scenarios, calling `test.Type` to simulate user input. The following test code will check that the output updates when the user’s name is input (be sure to add the import as well):

```
	test.Type(in, "Andy")
	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}

```

You can run all of these tests using `go test .` - just like any other tests. Doing so you will now see a failure - because we did not add the greeter logic. Update the `makeUI` function to the following code:

```
func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()
	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

```

Doing so you will see that the tests now pass. You can also run the full application (using `go run .`) and see the greeting update as you enter a name in the `Entry` field. Notice also that these tests all run without displaying a window or stealing your mouse - this is another benefit of the Fyne unit testing setup.

Share

Include playlist

An error occurred while retrieving sharing information. Please try again later.

0:00

0:00 / 7:43•Live

•

# https://docs.fyne.io/started/packaging

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Packaging for Desktop

Packaging a graphical app for distribution can be complex. Graphical applications typically have icons and metadata associated with them as well as specific formats required to integrate with each environment. Windows executables need embedded icons, macOS apps are bundles and with Linux there are various metadata files that should get installed. What a hassle!

Thankfully the “fyne” app has a “package” command that can handle this automatically. Just specifying the target OS and any required metadata (such as icon) will generate the appropriate package. The icon conversion will be done automatically for .icns or .ico so just provide a .png file :). All you need is to have the application already built for the target platform…

```
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon myapp.png

```

If you’re using an older version of Go (<1.16), you should install fyne using `go get`

```
go get fyne.io/fyne/v2/cmd/fyne
fyne package -os darwin -icon myapp.png

```

Will create myapp.app, a complete bundle structure for distribution to macOS users. You could then build the linux and Windows versions too…

```
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png

```

These commands will create:

- myapp.tar.gz that contains a folder structure starting at usr/local/ that a Linux user could expand to the root of their system.
- myapp.exe (after the second build, which is required for a windows package) will have the icon and app metadata embedded.

If you just want to install the desktop app on your computer then you can make use of the helpful install subcommand. For example to install your current application system wide you could simply execute the following:

```
fyne install -icon myapp.png

```

All of these commands also support a default icon file of `Icon.png` so that you can avoid typing the parameter for each execution. Since Fyne 2.1 there is also a [metadata file](/started/metadata) where you can set default options for your project.

Of course you can still run your applications using the standard Go tools if you prefer.

Fyne

Search

Watch later

Share

Copy link

Info

Shopping

Tap to unmute

If playback doesn't begin shortly, try restarting your device.

Share

Include playlist

An error occurred while retrieving sharing information. Please try again later.

0:00

0:00 / 4:02•Live

•

# https://docs.fyne.io/started/webapp

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Run in a Browser

Fyne applications can also be run over the web using a standard web browser! A web app created with Fyne will bundle a WebAssembly binary which will run in most modern browsers

To prepare your app to be used over the web we use the “fyne” cli app again, which has a “serve” command for quick testing

```
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne serve

```

You will see, after a few moments, that a web server has been started on port :8080. Just open your web browser to https://localhost:8080 and you can use your app!

### Packaging for web distribution

The `fyne serve` command is great for local testing, but just like other platforms you’ll want to be able to distribute your app as well. To prepare the files for upload just use the `fyne package` command like with regular [Packaging](/started/packaging).

```
fyne package -os web

```

### Demo

You can see a Fyne app in action to test on any of your devices by visiting .

### Limitations

As of release v2.5.0 the web driver is fully supported but is not yet 100% complete, so your app may not be able to use the following features:

- File open/save dialog
- Storage of Documents

These issues are being worked on and will be resolved in a future release.

# https://docs.fyne.io/started/metadata

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# App Metadata

Since release v2.1.0 of the `fyne` command we support a metadata file that allows you to store information about your application in the repository. This file is optional, but can help to avoid having to remember specific build parameters for each package and release command.

## Basic configuration

The file should be named `FyneApp.toml` in the directory where you run the `fyne` command (this is normally the `main` package). The contents of the file are as follows:

```
Website = "https://example.com"
[Details]
Icon = "Icon.png"
Name = "My App"
ID = "com.example.app"
Version = "1.0.0"
Build = 1

```

The top portion of the file is metadata that will be used if you upload your app to the listing page, so it is optional.

The `Details` table contains data about your application that are used in the release process by other app stores and operating systems.

The `fyne` tool will use this file if it is found, many mandatory command parameters are not required if the metadata is present. You can still override these values by using command line parameters.

## Linux & BSD configuration

For Linux and BSD builds there is an optional table called `LinuxAndBSD`. This table contains additional parameters for a “desktop entry” configuration file of the Fyne app. All parameters are optional, but when present they will be used by the `fyne` tool (in addition to parameters from the `Details` table).

The contents of this section is as follows (with example data):

```
[LinuxAndBSD]
 GenericName = "Web Browser"
 Categories = ["Network"]
 Comment = "View sites on the Internet"
 Keywords = ["browser", "web"]
 ExecParams = "-x 42"

```

Hint: For instructions on how to define these parameters correctly, please see the from freedesktop.org.

# https://docs.fyne.io/started/cross-compiling

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Compiling for different platforms

Cross compiling with Go is designed to be simple - we just set the environment variable `GOOS` for the target Operating System (and `GOARCH` if targeting a different architecture). Unfortunately when using native graphics calls the use of CGo in Fyne makes this a little harder.

### Compiling from a development computer

To cross-compile a Fyne application you will also have to set `CGO_ENABLED=1` which tells go to enable the C compiler (this is normally turned off when the target platform is different to the current system). Doing so unfortunately means that you must have a C compiler for the target platform that you are going to compile for. After installing the appropriate compilers you will also need to set the `CC` environment variable to tell Go which compiler to use.

There are many ways to install the required tools - and different tools that can be used. The configuration recommended by the Fyne developers is:

| GOOS (target) | CC                               | provider          | download                                                                        | notes                                                                                                                      |
| ------------- | -------------------------------- | ----------------- | ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `darwin`      | `o32-clang`                      | osxcross          | You will also need to install the macOS SDK (instructions at the download link) |
| `windows`     | `x86_64-w64-mingw64-gcc`         | mingw64           | package manager                                                                 | For macOS use                                                                                                              |
| `linux`       | `gcc` or `x86_64-linux-musl-gcc` | gcc or musl-cross | or package manager                                                              | musl-cross is available from to provide the linux gcc. You will also need to install X11 and mesa headers for compilation. |

With the environment variables above set you should be able to compile in the usual manner. If further errors occur it is likely to be due to missing packages. Some target platforms require additional libraries or headers to be installed for the compilation to succeed.

### Using a virtual environment

As a Linux system is able to cross compile to macOS and Windows easily it can be simpler to use a virtualised environment when you are not developing from Linux. Docker images are a useful tool for a complex build configuration and this works for Fyne as well. There are different tools that can be used. The tool recommended by the Fyne developers is . It has been inspired by and uses a built on top of the image, that includes the MinGW compiler for windows, and a macOS SDK, along with the Fyne requirements.

fyne-cross allows to build binaries and create distribution packages for the following targets:

| GOOS    | GOARCH |
| ------- | ------ |
| darwin  | amd64  |
| darwin  | 386    |
| linux   | amd64  |
| linux   | 386    |
| linux   | arm64  |
| linux   | arm    |
| windows | amd64  |
| windows | 386    |
| android | amd64  |
| android | 386    |
| android | arm64  |
| android | arm    |

ios  
freebsd | amd64  
freebsd | arm64

> Note: iOS compilation is supported only on darwin hosts.

#### Requirements

- go >= 1.13
- docker

#### Installation

You can install `fyne-cross` using the following command (requires Go 1.16 or later):

```
go install github.com/fyne-io/fyne-cross@latest

```

For earlier versions of Go, you need to use the following command instead:

```
go get github.com/fyne-io/fyne-cross

```

#### Usage

```
fyne-cross <command> [options]
The commands are:
	darwin    Build and package a fyne application for the darwin OS
	linux     Build and package a fyne application for the linux OS
	windows    Build and package a fyne application for the windows OS
	android    Build and package a fyne application for the android OS
	ios      Build and package a fyne application for the iOS OS
	freebsd    Build and package a fyne application for the freebsd OS
	version    Print the fyne-cross version information
Use "fyne-cross <command> -help" for more information about a command.

```

#### Wildcards

The `arch` flag support wildcards in case want to compile against all supported GOARCH for a specified GOOS

Example:

```
fyne-cross windows -arch=*

```

is equivalent to

```
fyne-cross windows -arch=amd64,386

```

#### Example

The example below cross compile and package the

```
git clone https://github.com/fyne-io/examples.git
cd examples

```

##### Compile and package the main example app

```
fyne-cross linux

```

> Note: by default fyne-cross will compile the package into the current dir.
>
> The command above is equivalent to: `fyne-cross linux .`

##### Compile and package a particular example app

```
fyne-cross linux -output bugs ./cmd/bugs

```

# https://docs.fyne.io/explore/canvas

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Canvas and CanvasObject

In Fyne a `Canvas` is the area within which an application is drawn. Each window has a canvas which you can access with `Window.Canvas()` but usually you will find functions on `Window` that avoid accessing the canvas.

Everything that can be drawn in Fyne is a type of `CanvasObject`. The example here opens a new window and then shows different types of primitive graphical element by setting the content of the window canvas. There are many ways that each type of object can be customised as shown with the text and circle examples.

As well as changing the content shown using `Canvas.SetContent()` it is possible to change the content that is currently visible. If, for example, you change the `FillColour` of a rectangle you can request a refresh of this existing component using `rect.Refresh()`.

```
package main
import (
	"image/color"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)
	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rect.FillColor = green
		rect.Refresh()
	}()
	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

```

We can draw many different drawing elements in the same way, such as circle and text.

```
func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}
func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

```

## Widget

A `fyne.Widget` is a special type of canvas object that has interactive elements associated with it. In widgets the logic is separate from the way that it looks (also called the `WidgetRenderer`).

Widgets are also types of `CanvasObject` and so we can set the content of our window to a single widget. See how we create a new `widget.Entry` and set it as the content of the window in this example.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")
	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/explore/container

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Container and Layouts

In the previous example we saw how to set a `CanvasObject` to the content of a `Canvas`, but it is not very useful to only show one visual element. To show more than one item we use the `Container` type.

As the `fyne.Container` also is a `fyne.CanvasObject`, we can set it to be the content of a `fyne.Canvas`. In this example we create 3 text objects and then place them in a container using the `container.NewWithoutLayout()` function. As there is no layout set we can move the elements around like you see with `text2.Move()`.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2)
	// content := container.New(layout.NewGridLayout(2), text1, text2)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

A `fyne.Layout` implements a method for organising items within a container. By uncommenting the `container.New()` line in this example you alter the container to use a grid layout with 2 columns. Run this code and try resizing the window to see how the layout automatically configures the contents of the window. Notice also that the manual position of `text2` is ignored by the layout code.

To see more you can check out the `Layout list`[](layouts).

# https://docs.fyne.io/explore/widgets

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Widget List

## Standard Widgets (in `widget` package)

### Accordion

Accordion displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

![](https://docs.fyne.io/images/widgets/accordion-light.png)

### Activity

Display an animated activity indicator.

![](https://docs.fyne.io/images/widgets/activity-light.png)

### Button

[Button](/widget/button) widget has a text label and icon, both are optional.

![](https://docs.fyne.io/images/widgets/button-light.png)

### Card

Card widget groups elements with a header and subheader, all are optional.

![](https://docs.fyne.io/images/widgets/card-light.png)

### Check

Check widget has a text label and a checked (or unchecked) icon.

![](https://docs.fyne.io/images/widgets/check-light.png)

### Entry

[Entry](/widget/entry) widget allows simple text to be input when focused.

![](https://docs.fyne.io/images/widgets/entry-light.png)

![](https://docs.fyne.io/images/widgets/entry-valid-light.png)

![](https://docs.fyne.io/images/widgets/entry-invalid-light.png)

PasswordEntry widget hides text input and adds a button to display the text.

![](https://docs.fyne.io/images/widgets/password-light.png)

### FileIcon

FileIcon provides helpful standard icons for various types of file. It displays the type of file as an indicator icon and shows the extension of the file type.

![](https://docs.fyne.io/images/widgets/fileicon-light.png)

### Form

[Form](/widget/form) widget is two column grid where each row has a label and a widget (usually an input). The last row of the grid will contain the appropriate form control buttons if any should be shown.

![](https://docs.fyne.io/images/widgets/form-light.png)

### Hyperlink

Hyperlink widget is a text component with appropriate padding and layout. When clicked, the URL opens in your default web browser.

![](https://docs.fyne.io/images/widgets/hyperlink-light.png)

### Icon

Icon widget is a basic image component that load’s its resource to match the theme.

![](https://docs.fyne.io/images/widgets/icon-light.png)

### Label

[Label](/widget/label) widget is a label component with appropriate padding and layout.

![](https://docs.fyne.io/images/widgets/label-light.png)

### Progress bar

[ProgressBar](/widget/progressbar) widget creates a horizontal panel that indicates progress.

![](https://docs.fyne.io/images/widgets/progress-light.png)

ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called.

![](https://docs.fyne.io/images/widgets/progressinf-light.png)

### RadioGroup

RadioGroup widget has a list of text labels and radio check icons next to each.

![](https://docs.fyne.io/images/widgets/radiogroup-light.png)

### RichText

RichText widget is a text component that shows various styles and embedded objects. It supports markdown parsing to construct the widget with ease.

![](https://docs.fyne.io/images/widgets/richtext-light.png)

### Select

Select widget has a list of options, with the current one shown, and triggers an event function when clicked.

![](https://docs.fyne.io/images/widgets/select-light.png)

### SelectEntry

Select entry widget adds an editable component to the select widget. Users can select an option or enter their own value.

![](https://docs.fyne.io/images/widgets/selectentry-light.png)

### Separator

Separator widget shows a dividing line between other elements.

![](https://docs.fyne.io/images/widgets/separator-light.png)

### Slider

Slider if a widget that can slide between two fixed values.

![](https://docs.fyne.io/images/widgets/slider-light.png)

### TextGrid

TextGrid is a monospaced grid of characters. This is designed to be used by a text editor, code preview or terminal emulator.

![](https://docs.fyne.io/images/widgets/textgrid-light.png)

### Toolbar

[Toolbar](/widget/toolbar) widget creates a horizontal list of tool buttons.

![](https://docs.fyne.io/images/widgets/toolbar-light.png)

## Collection Widgets (in `widget` package)

Collection widgets provide advanced caching functionality to provide high performance rendering of massive data. This does lead to a more complex constructor, but is a good balance for the outcome it enables. Each of these widgets uses a series of callbacks, the minimum set is defined by their constructor function, which includes the data size, the creation of template items that can be re-used and finally the function that applies data to a widget as it is about to be added to the display.

### List

[List](/collection/list) provides a high performance vertical scroll of many sub-items.

![](https://docs.fyne.io/images/widgets/list-light.png)

### Table

[Table](/collection/table) provides a high performance scrolled two dimensional display of many sub-items.

![](https://docs.fyne.io/images/widgets/table-light.png)

### Tree

[Tree](/collection/tree) provides a high performance vertical scroll of items that can be expanded to reveal child elements..

![](https://docs.fyne.io/images/widgets/tree-light.png)

### GridWrap

GridWrap is a collection widget that display each child item at the same size and wraps them to new lines as required.

![](https://docs.fyne.io/images/widgets/gridwrap-light.png)

## Container Widgets (in `container` package)

Container widgets are like regular containers but they provide some additional functionality.

### AppTabs

[AppTabs](/container/apptabs) widget allows switching visible content from a list of TabItems. Each item is represented by a button at the top of the widget.

![](https://docs.fyne.io/images/widgets/apptabs-light.png)

### Scroll

ScrollContainer defines a container that is smaller than the Content.

![](https://docs.fyne.io/images/widgets/scroll-light.png)

### Split

SplitContainer defines a container whose size is split between two children.

![](https://docs.fyne.io/images/widgets/split-light.png)

# https://docs.fyne.io/explore/widgets

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Widget List

## Standard Widgets (in `widget` package)

### Accordion

Accordion displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

![](https://docs.fyne.io/images/widgets/accordion-light.png)

### Activity

Display an animated activity indicator.

![](https://docs.fyne.io/images/widgets/activity-light.png)

### Button

[Button](/widget/button) widget has a text label and icon, both are optional.

![](https://docs.fyne.io/images/widgets/button-light.png)

### Card

Card widget groups elements with a header and subheader, all are optional.

![](https://docs.fyne.io/images/widgets/card-light.png)

### Check

Check widget has a text label and a checked (or unchecked) icon.

![](https://docs.fyne.io/images/widgets/check-light.png)

### Entry

[Entry](/widget/entry) widget allows simple text to be input when focused.

![](https://docs.fyne.io/images/widgets/entry-light.png)

![](https://docs.fyne.io/images/widgets/entry-valid-light.png)

![](https://docs.fyne.io/images/widgets/entry-invalid-light.png)

PasswordEntry widget hides text input and adds a button to display the text.

![](https://docs.fyne.io/images/widgets/password-light.png)

### FileIcon

FileIcon provides helpful standard icons for various types of file. It displays the type of file as an indicator icon and shows the extension of the file type.

![](https://docs.fyne.io/images/widgets/fileicon-light.png)

### Form

[Form](/widget/form) widget is two column grid where each row has a label and a widget (usually an input). The last row of the grid will contain the appropriate form control buttons if any should be shown.

![](https://docs.fyne.io/images/widgets/form-light.png)

### Hyperlink

Hyperlink widget is a text component with appropriate padding and layout. When clicked, the URL opens in your default web browser.

![](https://docs.fyne.io/images/widgets/hyperlink-light.png)

### Icon

Icon widget is a basic image component that load’s its resource to match the theme.

![](https://docs.fyne.io/images/widgets/icon-light.png)

### Label

[Label](/widget/label) widget is a label component with appropriate padding and layout.

![](https://docs.fyne.io/images/widgets/label-light.png)

### Progress bar

[ProgressBar](/widget/progressbar) widget creates a horizontal panel that indicates progress.

![](https://docs.fyne.io/images/widgets/progress-light.png)

ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called.

![](https://docs.fyne.io/images/widgets/progressinf-light.png)

### RadioGroup

RadioGroup widget has a list of text labels and radio check icons next to each.

![](https://docs.fyne.io/images/widgets/radiogroup-light.png)

### RichText

RichText widget is a text component that shows various styles and embedded objects. It supports markdown parsing to construct the widget with ease.

![](https://docs.fyne.io/images/widgets/richtext-light.png)

### Select

Select widget has a list of options, with the current one shown, and triggers an event function when clicked.

![](https://docs.fyne.io/images/widgets/select-light.png)

### SelectEntry

Select entry widget adds an editable component to the select widget. Users can select an option or enter their own value.

![](https://docs.fyne.io/images/widgets/selectentry-light.png)

### Separator

Separator widget shows a dividing line between other elements.

![](https://docs.fyne.io/images/widgets/separator-light.png)

### Slider

Slider if a widget that can slide between two fixed values.

![](https://docs.fyne.io/images/widgets/slider-light.png)

### TextGrid

TextGrid is a monospaced grid of characters. This is designed to be used by a text editor, code preview or terminal emulator.

![](https://docs.fyne.io/images/widgets/textgrid-light.png)

### Toolbar

[Toolbar](/widget/toolbar) widget creates a horizontal list of tool buttons.

![](https://docs.fyne.io/images/widgets/toolbar-light.png)

## Collection Widgets (in `widget` package)

Collection widgets provide advanced caching functionality to provide high performance rendering of massive data. This does lead to a more complex constructor, but is a good balance for the outcome it enables. Each of these widgets uses a series of callbacks, the minimum set is defined by their constructor function, which includes the data size, the creation of template items that can be re-used and finally the function that applies data to a widget as it is about to be added to the display.

### List

[List](/collection/list) provides a high performance vertical scroll of many sub-items.

![](https://docs.fyne.io/images/widgets/list-light.png)

### Table

[Table](/collection/table) provides a high performance scrolled two dimensional display of many sub-items.

![](https://docs.fyne.io/images/widgets/table-light.png)

### Tree

[Tree](/collection/tree) provides a high performance vertical scroll of items that can be expanded to reveal child elements..

![](https://docs.fyne.io/images/widgets/tree-light.png)

### GridWrap

GridWrap is a collection widget that display each child item at the same size and wraps them to new lines as required.

![](https://docs.fyne.io/images/widgets/gridwrap-light.png)

## Container Widgets (in `container` package)

Container widgets are like regular containers but they provide some additional functionality.

### AppTabs

[AppTabs](/container/apptabs) widget allows switching visible content from a list of TabItems. Each item is represented by a button at the top of the widget.

![](https://docs.fyne.io/images/widgets/apptabs-light.png)

### Scroll

ScrollContainer defines a container that is smaller than the Content.

![](https://docs.fyne.io/images/widgets/scroll-light.png)

### Split

SplitContainer defines a container whose size is split between two children.

![](https://docs.fyne.io/images/widgets/split-light.png)

# https://docs.fyne.io/explore/dialogs

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Dialog List

## Standard Dialogs

### Color

Allow users to pick a colour from a standard set (or any color in advanced mode).

### Confirm

Ask for conformation of an action.

### FileOpen

Present this to ask user to choose a file to use inside the app. The actual dialog displayed will depend on the current operating system.

### Form

Get various input elements in a dialog, with validation.

### Information

A simple way to present some information to the app user.

### Custom

Present any content inside a dialog container.

# https://docs.fyne.io/explore/icons

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Theme Icons

Each of the following icons is available via the `theme` package as a function. For example `theme.InfoIcon()`.

The icons are also available via their source icon name by using the `ThemeIconName` with the `Icon` method on a struct implementing `fyne.Theme`. For example `theme.Icon(theme.IconNameInfo)`.

## List

- AccountIcon
- ArrowDropDownIcon
- ArrowDropUpIcon
- BrokenImageIcon
- CancelIcon
- CheckButtonIcon
- CheckButtonCheckedIcon
- ColorAchromaticIcon
- ColorChromaticIcon
- ColorPaletteIcon
- ComputerIcon
- ConfirmIcon
- ContentAddIcon
- ContentClearIcon
- ContentCopyIcon
- ContentCutIcon
- ContentPasteIcon
- ContentRedoIcon
- ContentRemoveIcon
- ContentUndoIcon
- DeleteIcon
- DocumentCreateIcon
- DocumentIcon
- DocumentPrintIcon
- DocumentSaveIcon
- DownloadIcon
- ErrorIcon
- FileApplicationIcon
- FileAudioIcon
- FileIcon
- FileImageIcon
- FileTextIcon
- FileVideoIcon
- FolderIcon
- FolderNewIcon
- FolderOpenIcon
- GridIcon
- HelpIcon
- HistoryIcon
- HomeIcon
- InfoIcon
- ListIcon
- LoginIcon
- LogoutIcon
- MailAttachmentIcon
- MailComposeIcon
- MailForwardIcon
- MailReplyAllIcon
- MailReplyIcon
- MailSendIcon
- MediaFastForwardIcon
- MediaFastRewindIcon
- MediaMusicIcon
- MediaPauseIcon
- MediaPhotoIcon
- MediaPlayIcon
- MediaRecordIcon
- MediaReplayIcon
- MediaSkipNextIcon
- MediaSkipPreviousIcon
- MediaStopIcon
- MediaVideoIcon
- MenuExpandIcon
- MenuIcon
- MoreHorizontalIcon
- MoreVerticalIcon
- MoveDownIcon
- MoveUpIcon
- NavigateBackIcon
- NavigateNextIcon
- QuestionIcon
- RadioButtonCheckedIcon
- RadioButtonIcon
- SearchReplaceIcon
- SearchIcon
- SettingsIcon
- StorageIcon
- UploadIcon
- ViewFullScreenIcon
- ViewRefreshIcon
- ViewRestoreIcon
- ViewZoomFitIcon
- ViewZoomInIcon
- ViewZoomOutIcon
- VisibilityOffIcon
- VisibilityIcon
- VolumeDownIcon
- VolumeMuteIcon
- VolumeUpIcon
- WarningIcon

## Using other color sets

Each icon can be used as a source for a particular themed color using the various public helper methods:

- `NewDisabledThemedResource`
- `NewErrorThemedResource`
- `NewInvertedThemedResource`
- `NewPrimaryThemedResource`

By default, all icons adapt to the current theme foreground using `NewThemedResource` which uses the theme foreground color. All Icons are SVG `width="24"`, `height="24"`.

# https://docs.fyne.io/explore/shortcuts

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Adding Shortcuts to an App

Shortcuts are common tasks that can be triggered by keyboard combinations or context menus. Shortcuts, much like keyboard events, can be attached to a focused element or registered on the `Canvas` to always be available in a `Window`.

## Registering with a Canvas

There are many standard shortcuts defined (such as `fyne.ShortcutCopy`) which are connected to standard keyboard shortcuts and right-click menus. The first step to adding a new `Shortcut` is to define the shortcut. For most uses this will be a keyboard triggered shortcut, which is a desktop extension. To do this we use `desktop.CustomShortcut`, for example to use the Tab key and Control modifier you might do the following:

```
	ctrlTab  := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	ctrlAltTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl | fyne.KeyModifierAlt}

```

Notice that this shortcut can be re-used so you could attach it to menus or other items as well. For this example we want it to be always available, so we register it with our window’s `Canvas` as follows:

```
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})
	w.Canvas().AddShortcut(ctrlAltTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Alt+Tab")
	})

```

As you can see there are two parts to registering a shortcut in this way - passing the shortcut definition and also a callback function. If the user types the keyboard shortcut then the function will be called and the output printed.

Shortcuts only work in combination with modifier keys. In order to react to keyboard input without modifiers use `canvas.OnTypedRune` or `canvas.OnTypedKey`.

## Adding shortcuts to an Entry

It can also be helpful to have a shortcut apply only when the current item is focused. This approach can be used for any focusable widget, and is managed by extending that widget and adding a `TypedShortcut` handler. This is much like adding key handlers, except the value passed in will be a `fyne.Shortcut`.

```
type myEntry struct {
	widget.Entry
}
func (m *myEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s)
		return
	}
	log.Println("Shortcut typed:", s)
}

```

From the excerpt above you can see how a `TypedShortcut` handler might be implemented. Inside this function you should check whether the shortcut is of the custom type used earlier. If the shortcut is a standard one it’s a good idea to call the original shortcut handler (if the widget had one). With those checks done you can compare the shortcut with the various types you are handling (if there are multiple).

# https://docs.fyne.io/explore/preferences

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Using the Preferences API

Storing user configurations and values is a common task for application developers, but implementing it across multiple platforms can be tedious and time-consuming. To make it easier, Fyne has an API for storing values on the filesystem in a clean and understandable way while the complex parts are handled for you.

Lets start with the setup of the API. It is part of the interface where storage and loading functions exist for values of `bool`, `float64`, `int`, and `string` types, and for lists of each type as well. For each type there are three different functions, one for loading, one loading with a fallback value and lastly, one for storing values. An example of the three functions and their behaviour can be seen below for the `string` type:

```
// String looks up a string value for the key
String(key string) string
// StringWithFallback looks up a string value and returns the given fallback if not found
StringWithFallback(key, fallback string) string
// SetString saves a string value for the given key
SetString(key string, value string)

```

These functions can be accessed through the created application variable and calling the `Preferences()` method on. Please note that it is necessary to create the apps with a unique ID (usually like a reversed url). This means that the application will need to be created using `app.NewWithID()` (or have ID specified in `FyneApp.toml`) to have its own place to store values. It can roughly be used like the example below:

```
a := app.NewWithID("com.example.tutorial.preferences")
[...]
a.Preferences().SetBool("Boolean", true)
a.Preferences().SetFloatList("Constants", []float64{3.142, 2.718})
number := a.Preferences().IntWithFallback("ApplicationLuckyNumber", 21)
expression := a.Preferences().String("RegularExpression")
[...]

```

To show this, we are going to build a simple little app that always closes after a set amount of time. This timeout should be user changeable and applied on the next start of the application.

Let us start by creating a variable called `timeout` that will be used to store time in the form of `time.Duration`.

```
var timeout time.Duration

```

Then we could create a select widget to let the user select the timeout from a couple pre-defined strings and then multiplying the timeout by the number of seconds that the string relates to. Lastly, the `"AppTimeout"` key is used to set the string value to the selected one.

```
timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
  switch selected {
  case "10 seconds":
    timeout = 10 * time.Second
  case "30 seconds":
    timeout = 30 * time.Second
  case "1 minute":
    timeout = time.Minute
  }
  a.Preferences().SetString("AppTimeout", selected)
})

```

Now we want to grab the set value and if none exists, we want to have a fallback that sets the timeout to the shortest one possible to save the user time when waiting. This can be done by setting the selected value of `timeoutSelector` to the loaded value or the fallback if that happens to be the case. By doing it this way, the code inside the select widget will run for that specific value.

```
timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))

```

The last part will just be to have a function that starts in a separate goroutine and tells the application to quit after the selected timeout.

```
go func() {
  time.Sleep(timeout)
  a.Quit()
}()

```

In the end, the resulting code should look something like this:

```
package main
import (
  "time"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)
func main() {
  a := app.NewWithID("com.example.tutorial.preferences")
  w := a.NewWindow("Timeout")
  var timeout time.Duration
  timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
    switch selected {
    case "10 seconds":
      timeout = 10 * time.Second
    case "30 seconds":
      timeout = 30 * time.Second
    case "1 minute":
      timeout = time.Minute
    }
    a.Preferences().SetString("AppTimeout", selected)
  })
  timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))
  go func() {
    time.Sleep(timeout)
    a.Quit()
  }()
  w.SetContent(timeoutSelector)
  w.ShowAndRun()
}

```

# https://docs.fyne.io/explore/translations

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Adding app translations

Most apps will want to add translations at some point, and since v2.5.0 Fyne helps to make this really simple. The translation features simply sit in your project as `.json` files and the content can be crowd-sourced by popular platforms like .

## Setup

The Fyne toolkit is already translated and will replace recognised strings in standard locations. To make use of this functionality for your own app strings you make use of the `fyne.io/fyne/v2/lang` package and the helper functions it provides.

The simplest way to prepare for translations is to use the `L` or `Localize` function to mark a string as translatable - if the translation is not found then the string will be used as a fallback.

```
  title := widget.NewLabel(lang.L("My App Title"))

```

In some cases it may be desirable to label a string with a unique ID instead of using the default value - for disambiguation or other reason. In this case you would use `LocalizeKey` or the `X` alias.

```
  title := widget.NewLabel(lang.X("window.title", "My App Window Title"))

```

That may be all you need to know to get started - skip to [translation files](translation-files).

## Translation files

Each translation file is a simple JSON file, if you use the `L` form all you need to do is insert the string with it’s translation for each language - 1 per file. For example, this may be your app `en.json`:

```
{"My App Title":"My App Title"}
```

Then your French translation could look like:

```
{"My App Title":"Titre de mon application"}
```

Each file can be most easily loaded using the Go `embed` feature - place each of the files in a directory called “translation” and then define them simply as:

```
//go:embed translation
var translations embed.FS

```

Finally you can tell Fyne to load these translations with a single function call between `app.New` and `Run()` in your `main` function:

```
	lang.AddTranslationsFS(translations, "translation")

```

This uses the embedded filesystem and specifies the name of the directory that the files are stored in.

When your app starts it will display using the translations for the current user’s langauge configuration.

## Plurals

In more complex cases the string will change based on the number of items it refers to. For this the `lang.LocalizePlural` function (aliased to `lang.N`) is avaialble.

```
  age := widget.NewLabel(lang.N("{{.Years}} years old", years, map[string]any{"Years": years}))

```

You can pass data in this way to any of the calls, using template syntax to insert the value. A struct with exported fields, or a map as illustrated above, can be used to insert data.

The translation file will look a little more complex for this - the key has two sub-values called “one” and “other” (some languages may use more keys).

```
{"Age":{"one":"1 year old","other":"{{.Years}} years old"}}
```

# https://docs.fyne.io/explore/systray

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# System Tray Menu

## Adding a System Tray menu

Since the v2.2.0 release Fyne has built in support for a system tray menu. This feature displays an icon on macOS, Windows and Linux computers and when tapped will pop out a menu as specified by the app.

As this is a desktop specific feature we must first do a runtime check that the app is running in desktop mode. To do this, and get a reference to the desktop features, we do a Go type assertion:

```
	if desk, ok := a.(desktop.App); ok {
...
	}

```

If the `ok` variable is true then we can set up a menu using the standard Fyne menu API that you might have used in `Window.SetMainMenu` before.

```
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				log.Println("Tapped show")
			}))
		desk.SetSystemTrayMenu(m)

```

With this code added to the setup of your application you can run the app and see that it shows a Fyne icon in the system tray. When you tap it a menu will appear containing “Show” and “Quit”.

The default icon is the Fyne logo, you can either fix this using [app metadata](/started/metadata) or by setting the app icon in `App.SetIcon` or for system tray directly using `desk.SetSystemTrayIcon`.

## Manage window lifecycle

By default a Fyne app will exit when you close all windows and this may not be what you want with a system tray app. To override the behaviour you can use the `Window.SetCloseIntercept` feature to override what happens when a window is closed. In the example below we hide the window instead of closing it by calling `Window.Hide()`. Add this before you show the window for the first time.

```
	w.SetCloseIntercept(func() {
		w.Hide()
	})

```

The benefit of hiding a window is that you can simply show it again using `Window.Show()` which is much more efficient than creating a new window if the same content is needed a second time. We update the menu created earlier to show the window that was hidden above.

```
	fyne.NewMenuItem("Show", func() {
		w.Show()
	})

```

## Complete app

That’s all there is to setting up a system tray menu with Fyne! The complete code for this tutorial is as follows.

```
package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)
func main() {
	a := app.New()
	w := a.NewWindow("SysTray")
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}

```

Fyne

Search

Watch later

Share

Copy link

Info

Shopping

Tap to unmute

If playback doesn't begin shortly, try restarting your device.

Share

Include playlist

An error occurred while retrieving sharing information. Please try again later.

0:00

0:00 / 4:48•Live

•

# https://docs.fyne.io/explore/binding

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Data Binding

Data binding was introduced in Fyne v2.0.0 and makes it easier to connect many widgets to a data source that will update over time. the `data/binding` package has many helpful bindings that can manage most standard types that will be used in an application. A data binding can be managed using the binding API (for example `NewString`) or it can be connected to an external item of data like (`BindInt(*int)`).

Widgets that support binding typically have a `...WithData` constructor to set up the binding when creating the widget. You can also call `Bind()` and `Unbind()` to manage the data of an existing widget. The following example shows how you can manage a `String` data item that is bound to a simple `Label` widget.

```
package main
import (
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()
	w.SetContent(widget.NewLabelWithData(str))
	w.ShowAndRun()
}

```

You can find out more in the [data binding](/binding/) section of this site.

# https://docs.fyne.io/explore/compiling

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Compile Options

## Build tags

Fyne will typically configure your application appropriately for the target platform by selecting the driver and configuration. The following build tags are supported and can help in your development. For example if you wish to simulate a mobile application whilst running on a desktop computer you could use the following command:

```
go run -tags mobile main.go

```

| Tag               | Description                                                                                                                                                                                                                                                                |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `debug`           | Show debug information, including visual layout to help understand your app.                                                                                                                                                                                               |
| `gles`            | Force use of embedded OpenGL (GLES) instead of full OpenGL. This is normally controlled by the target device and not normally needed.                                                                                                                                      |
| `hints`           | Display developer hints for improvements or optimisations. Running with `hints` will log when your application does not follow material design or other recommendations.                                                                                                   |
| `mobile`          | This tag runs an application in a simulated mobile window. Useful when you want to preview your app on a mobile platform without compiling and installing to the device.                                                                                                   |
| `no_animations`   | Disable the non-essential animations in standard widgets and containers.                                                                                                                                                                                                   |
| `no_emoji`        | Don’t include the embedded emoji font. This will disable emoji in your app but will make the binary smaller.                                                                                                                                                               |
| `no_native_menus` | This flag is specifically for macOS and indicates that the application should not use the macOS native menus. Instead menus will be displayed inside the application window. Most useful for testing an application on macOS to simulate the behavior on Windows or Linux. |

# https://docs.fyne.io/canvas/rectangle

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Rectangle

`canvas.Rectangle` is the simplest canvas object in Fyne. It displays a block of the specified colour. You can also set the colour using the `FillColor` field.

In this example the rectangle fills the window as it is the only content element.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Rectangle")
	rect := canvas.NewRectangle(color.White)
	w.SetContent(rect)
	w.Resize(fyne.NewSize(150, 100))
	w.ShowAndRun()
}

```

Other `fyne.CanvasObject` types have more configuration, let us look [next](text) at `canvas.Text`.

# https://docs.fyne.io/canvas/text

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Text

`canvas.Text` is used for all text rendering within Fyne. It is created by specifying the text and colour for the text. Text is rendered using the default font, specified by the current theme.

The text object allows certain configuration like the `Alignment` and `TextStyle` field. as illustrated in the example here. To use a monospaced font instead you can specify `fyne.TextStyle{Monospace: true}`.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Text")
	text := canvas.NewText("Text Object", color.White)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(text)
	w.ShowAndRun()
}

```

It is possible to use an alternative font by specifying a `FYNE_FONT` environment variable. Use this to set a `.ttf` file to use instead of the one provided in the Fyne toolkit or the current theme.

# https://docs.fyne.io/canvas/line

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Line

The `canvas.Line` object draws a line from the `Position1` (default is top, left) to `Position2` (default is bottom, right). You specify its colour and can vary the stroke width which otherwise defaults to `1`.

A line position can be manipulated using the `Position1` or `Position2` fields or by using the `Move()` and `Resize()` functions. For example a 0 width area will show a vertical line whereas 0 height would be horizontal.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Line")
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	w.SetContent(line)
	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}

```

Lines are typically used in a custom layout or controlled manually. Unlike text they have no natural (minimum) size but can be used to great effect in complex layouts.

# https://docs.fyne.io/canvas/circle

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Circle

`canvas.Circle` defines a circle shape filled by the specified colour. You may also set a `StrokeWidth` and therefore a different `StrokeColor` as shown in this example.

The circle will fill the space specified by calling `Resize()` or by the layout it is controlled by. As the example sets the circle as the window content it will resize to fill the window, within a basic padding (controlled by the theme).

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Circle")
	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{Y: 0x99}
	circle.StrokeWidth = 5
	w.SetContent(circle)
	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}

```

All these have been basic types that can be rendered by our driver with no additional information. Next we will look at more complex types starting with `Image`[](image).

# https://docs.fyne.io/canvas/image

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Image

A `canvas.Image` represents a scalable image resource in Fyne. It can be loaded from a resource (as shown in the example), from an image file, from a URI location containing an image, from an `io.Reader`, or from a Go `image.Image` in memory.

The default image fill mode is `canvas.ImageFillStretch` which will cause it to fill the space specified (through `Resize()` or layout). Alternatively you could use `canvas.ImageFillContain` to ensure that the aspect ratio is maintained and the image is within the bounds. Further to this you can use `canvas.ImageFillOriginal` (as used in the example here) which ensures that it will also have a minimum size equal to that of the original image size.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")
	image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.ShowAndRun()
}

```

Images can be bitmap based (like PNG and JPEG) or vector based (such as SVG). Where possible we recommend scalable images as they will continue to render well as sizes change. Be careful when using original image sizes as they may not behave exactly as expected with different user interface scales. As Fyne allows the entire user interface to scale, a 25px image file may not be the same height as a 25 height fyne object.

# https://docs.fyne.io/canvas/raster

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Raster

The `canvas.Raster` is like an image but draws exactly one spot for each pixel on the screen. This means that as a user interface scales or the image resizes more pixels will be requested to fill the space. To do this we use a `Generator` function as illustrated in this example - it will be used to return the colour of each pixel.

The generator functions can be pixel based (as in this example where we generate a new random colour for each pixel) or based on full images. Generating complete images (with `canvas.NewRaster()`) is more efficient but sometimes controlling pixels directly is more convenient.

```
package main
import (
	"image/color"
	"math/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")
	raster := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)), A: 0xff}
		})
	// raster := canvas.NewRasterFromImage()
	w.SetContent(raster)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}

```

If your pixel data is stored in an image you can load it through the `NewRasterFromImage()` function which will load the image to display pixel perfect on screen.

# https://docs.fyne.io/canvas/gradient

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Gradient

The last canvas primitive type is Gradient, available as `canvas.LinearGradient` and `canvas.RadialGradient` which is used to draw a gradient from one colour to another in various patterns. You can create gradients using `NewHorizontalGradient()`, `NewVerticalGradient()` or `NewRadialGradient()`.

To create a gradient you need a start and end colour - every colour in between is calculated by the canvas. In this example we use `color.Transparent` to show how a gradient (or any other type) could use an alpha value to be semi-transparent over the content behind.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")
	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	//gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)
	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}

```

# https://docs.fyne.io/canvas/animation

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Animation

Fyne includes an animation framework that allows you to smoothly transition canvas properties from one value to another over time. An animation can contain any code which means that any types of object attributes can be managed, however there are builtin animations for size, position and color.

Animations are normally created using the builtin helpers of the canvas package, such as `NewSizeAnimation`, and calling `Start()` on the created animation. You can set animations to repeat or auto reverse, as we will see below.

Let us look first at a colour animation which gradually changes the fill colour of a `Rectangle`. In the following code sample we set an rectangle to be set as the content of a window, as we have done in earlier code samples. The big difference is the animation that we start just before showing the window. The animation is created using `NewColorRGBAAnimation` which will transition the colour channels from the defined `red` state through to `blue` and it will take 2 seconds (the specified duration) to do so.

```
package main
import (
	"image/color"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))
	red := color.NRGBA{R:0xff, A:0xff}
	blue := color.NRGBA{B:0xff, A:0xff}
	canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(c color.Color) {
		obj.FillColor = c
		canvas.Refresh(obj)
	}).Start()
	w.Resize(fyne.NewSize(250, 50))
	w.SetPadded(false)
	w.ShowAndRun()
}

```

It is also possible to animate multiple properties at the same time. If you look carefully you will see that we added the rectangle into a container without layout - this means that we can manually move or resize the object. Let’s add a new position animation that will move the `Rectangle` across the window, and automatically reverse as well.

```
move := canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(200, 0), time.Second, obj.Move)
move.AutoReverse = true
move.Start()

```

Because the `Move()` function of `CanvasObject` expects a `fyne.Position` argument, and so does the position animation callback, we can simply pass the method name instead of creating a new function If you add the code above just under the first animation you will see that the object moves across the window at the same time as it changes colour!

# https://docs.fyne.io/container/box

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Box

As discussed in [Container and Layouts](/explore/container) elements within a container can be arranged using a layout. This section explores the builtin layouts and how to use them.

The most commonly used layout is `layout.BoxLayout` and it has two variants, horizontal and vertical. A box layout arranges all elements in a single row or column with optional spaces to assist alignment.

A horizontal box layout, created with `layout.NewHBoxLayout()` creates an arrangement of items in a single row. Every item in the box will have its width set to its `MinSize().Width` and the height will be equal for all items, the largest of all the `MinSize().Height` values. The layout can be used in a container or you can use the box widget `widget.NewHBox()`.

A vertical box layout is similar but it arranges items in a column. Each item will have its height set to minimum and all the widths will be equal, set to the largest of the minimum widths.

To create an expanding space between elements (so that some are left aligned and the others right aligned, for example) add a `layout.NewSpacer()` as one of the items. A spacer will expand to fill all available space. Adding a spacer at the beginning of a vertical box layout will cause all items to be bottom aligned. You can add one to the beginning and end of a horizontal arrangement to create a center alignment.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")
	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)
	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/container/grid

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Grid

The grid layout lays out the elements of a container in a grid pattern with a fixed number of columns. Items will fill a single row until the number of columns is met, after this a new row will be created. Vertical space will be split equally between each of the rows of objects.

You create a grid layout using `layout.NewGridLayout(cols)` where cols is the number of items (columns) you wish to have in each row. This layout is then passed as the first parameter to `container.New(...)`.

If you resize the container then each of the cells will resize equally to share the available space.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")
	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/container/gridwrap

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Grid Wrap

Like the previous grid layout, the grid wrap layout creates an arrangement of elements in a grid pattern. However this grid does not have a set number of columns, instead it uses a fixed size for each cell and then flows the content to as many rows as is needed to display the items.

You create a grid wrap layout using `layout.NewGridWrapLayout(size)` where size specifies the size to apply to all child elements. This layout is then passed as the first parameter to `container.New(...)`. The number of columns and rows will be calculated based on the current size of the container.

Initially a grid wrap layout will have a single column, if you resize it (as illustrated in the code comment to the right) it will rearrange the child elements to fill the space.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")
	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)
	myWindow.SetContent(grid)
	// myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/container/border

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Border

The border layout is possibly the most widely used to construct user interfaces as it allows the positioning of items around a central element which will expand to fill the space. To create a border container you need to pass the `fyne.CanvasObject`s that should be positioned in a border position to the constructor’s first four parameters. This syntax is basically just `container.NewBorder(top, bottom, left, right, center)` as illustrated in the example.

Any items passed to the container after the first four items will be positioned to the central area and will expand to fill the space available. You can also pass `nil` to border parameters that you wish to leave empty.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")
	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.NewBorder(top, nil, left, nil, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

Note that all items in the center will expand to fill the space (as if they were in a `layout.MaxLayout`[](/container/max) container). To manage the area yourself you can use any `fyne.Container` as the content instead.

# https://docs.fyne.io/container/form

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Form

The `layout.FormLayout` is like a 2 column [grid layout](/container/grid) but tweaked to lay out forms in an application. The height of each item will be the larger of the two minimum heights in each row. The width of the left item will be the largest minimum width of all items in the first column whilst the second item in each row will expand to fill the space.

This layout is more typically used within the `widget.Form` (for validation, submit and cancel buttons, etc) but it can also be used directly with `layout.NewFormLayout()` passed to the first parameter of `container.New(...)`.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")
	label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value")
	label2 := widget.NewLabel("Label 2")
	value2 := widget.NewLabel("Something")
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/container/center

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Center

`layout.CenterLayout` organises all items in its container to be centered in the available space. The objects will be drawn in the order they are passed to the container, with the last being drawn top-most.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Center Layout")
	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

The center layout causes all items to stay at their minimum size, if you wish to expand items to fill the space then see `layout.MaxLayout`[](max).

# https://docs.fyne.io/container/stack

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Max

The `layout.NewStackLayout()` is the simplest layout, it sets all items in the container to be the same size as the container. This is not often useful in general containers but can be suitable when composing widgets.

The stack layout will require the container to be at least the size of the largest item’s minimum size. The objects will be drawn in the order the are passed to the container, with the last being drawn top-most.

```
package main
import (
	"image/color"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Stack Layout")
	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewStackLayout(), img, text)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/container/apptabs

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# AppTabs

The AppTabs container is used to allow the user to switch between different content panels. Tabs are either just text or text and an icon. It is recommended not to mix some tabs having icons and some without. A tab container is created using `container.NewAppTabs(...)` and passing `container.TabItem` items (that can be created using `container.NewTabItem(...)`).

The tab container can be configured by setting the location of tabs, one of `container.TabLocationTop`, `container.TabLocationBottom`, `container.TabLocationLeading` and `container.TabLocationTrailing`. The default location is top.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))
	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

```

When loaded on a mobile device the tab location may be ignored. In a portrait orientation a leading or trailing position will be changed to bottom. When in landscape orientation, the top or bottom positions will be moved to leading.

# https://docs.fyne.io/widget/label

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Label

Widgets are the main components of a Fyne application GUI, they can be used in any place that a basic `fyne.CanvasObject` can. They manage user interactions and will always match the current theme.

The Label widget is the simplest of them - it presents text to the user. Unlike `canvas.Text` it can handle some simple formatting (such as `\n`) and wrapping (by setting the `Wrapping` field). You can create a label by calling `widget.NewLabel("some text")`, the result can be assigned to a variable or passed directly into a container.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Label Widget")
	content := widget.NewLabel("text")
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/widget/button

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Button

The button widget can contain text, an icon or both, the constructor functions are `widget.NewButton()` and `widget.NewButtonWithIcon()`. To create a text button there are just 2 parameters, the `string` content and a 0 parameter `func()` that will be called when the button is tapped. See the example for how that can be created.

The button constructor with an icon includes an additional parameter which is the `fyne.Resource` which contains the icon data. The builtin icons within the `theme` package all adapt appropriately to a change in theme. You can pass in your own image if it is loaded as a resource - helpers such as `fyne.LoadResourceFromPath()` may assist, though bundling resources is recommended where possible.

To create a button with only an icon you should pass “” as the label parameter to `widget.NewButtonWithIcon()`.

```
package main
import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")
	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})
	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/widget/entry

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Entry

The entry widget is used for user input of simple text content. An entry can be created with a simple `widget.NewEntry()` constructing function. When you create the widget keep a reference so that you can access its `Text` field later. It is also possible to use the `OnChanged` callback function to be notified every time the content changes.

Entry widgets can also have validation for verifying the text input typed into it. This can be done by setting the `Validator` field to a `fyne.StringValidator`. You can also set a `PlaceHolder` text and also set the entry to `MultiLine` to accept more than one line of text.

```
package main
import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

You can also create a password entry (where the content is obscured) using the `NewPasswordEntry()` function.

# https://docs.fyne.io/widget/choices

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Choices

There are various widgets available to present the user with a choice, these include a checkbox, radio group and select popup.

The `widget.Check` provides a simple yes/no choice and is created using a string label. Each of these widgets also takes a “changed” `func(...)` where the parameter is of the appropriate type. `widget.NewCheck(..)` therefore takes a `string` parameter for the label and a `func(bool)` param for the change handler. You can also use the `Checked` field to get the boolean value.

The radio widget is similar, but the first parameter is a slice of `string`s that represents each of the options. The change function expects a `string` parameter this time to return the currently selected value. Call `widget.NewRadioGroup(...)` to construct the radio group widget, you can use this reference later to read the `Selected` field instead of using the change callback.

The select widget is identical in the constructor signature as the radio widget. Calling `widget.NewSelect(...)` will instead show a button that displays a popup when tapped from which the user can make a selection. This is better for long lists of options.

```
package main
import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")
	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})
	myWindow.SetContent(container.NewVBox(check, radio, combo))
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/widget/form

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Form

The form widget is used to lay out many input fields with labels and optional cancel and submit buttons. In its most bare form it aligns labels to the left of each input widget. By setting OnCancel or OnSubmit the form will add a button bar with the specified handlers called when appropriate.

The widget can be created with `widget.NewForm(...)` passing a list of `widget.FormItem`s, or by using the `&widget.Form{}` syntax illustrated in the example. There is also a helpful `Form.Append(label, widget)` that can be used for an alternative syntax.

In this example we create two entries, one of which is a “multiline” (like HTML TextArea) to hold values. There is an OnSubmit handler which prints the information before closing the window (and therefore the application).

```
package main
import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")
	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			myWindow.Close()
		},
	}
	// we can also append items
	form.Append("Text", textArea)
	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/widget/progressbar

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# ProgressBar

The progress bar widget has two forms, the standard progress bar shows the user which `Value` has been reached, from `Min` to `Max`. The default min is `0.0` and the max defaults to `1.0`. To use the default values just call `widget.NewProgressBar()`. After creating you can set the `Value` field.

To set up a custom range you can set `Min` and `Max` fields manually. The label will always show the percentage completion.

The other form of progress widget is the infinite progress bar. This version simply shows that some activity is ongoing by moving a segment of the bar from left to right and back again. You create this using `widget.NewProgressBarInfinite()` and it will start animating as soon as it is shown.

```
package main
import (
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar Widget")
	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()
	myWindow.SetContent(container.NewVBox(progress, infinite))
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/widget/toolbar

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Toolbar

The toolbar widget creates a row of action buttons using icons to represent each. The `widget.NewToolbar(...)` constructor function takes a list of `widget.ToolbarItem` parameters. The builtin types of toolbar items are action, separator and spacer.

The most used item is an action that is created using the `widget.NewToolbarItemAction(..)` function. An action takes two parameters, first being the icon resource to draw and the latter is the `func()` to call when tapped. This creates a standard toolbar button.

You can use `widget.NewToolbarSeparator()` to create a small divider between items in a toolbar (usually a thin vertical line). Lastly you can use `widget.NewToolbarSpacer()` to create a flexible space between elements. This is most useful to right align the toolbar items that are listed after the spacer.

A toolbar should always be at the top of the content area so it’s normal to add it to a `fyne.Container` using the `layout.BorderLayout` to align it above other content.

```
package main
import (
	"log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar Widget")
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/collection/list

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# List

The `List` collection widget is one of the toolkit’s collection widgets. These widgets are designed to help build really performant interfaces when lots of data is being presented. You can also see the [Table](/collection/table) and [Tree](/collection/tree) widgets which have a similar API. Because of this design they are a little more complicated to use.

The `List` uses callback functions to ask for data when it is required. There are 3 main callbacks, `Length`, `CreateItem` and `UpdateItem`. The Length callback (passed first) is the simplest, it returns how many items are in the data to be presented. The others relate to templates - how graphical elements are created, cached and re-used.

The `CreateItem` callback returns a new template object. This will be re-used with real data when the widget is presented. The `MinSize` of this object will influence the `List` minimum size. Lastly `UpdateItem` is called to apply an item of data to a cached template. Use this to set the content ready for display.

```
package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
var data = []string{"a", "string", "list"}
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/collection/table

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Table

The `Table` collection widget is like the [List](/collection/list) widget (another of the toolkit’s collection widgets) with a two-dimensional index. Like `List` this is designed to help build really performant interfaces when lots of data is being presented. Because of this the widget is not created with all the data embedded, but instead calls out to the data source when needed.

The `Table` uses callback functions to ask for data when it is required. There are 3 main callbacks, `Length`, `CreateCell` and `UpdateCell`. The Length callback (passed first) is the simplest, it returns how many items are in the data to be presented, the two ints it returns represent the row and colum count. The other two relate to the content templates.

The `CreateCell` callback returns a new template object, just like list. The difference being that `MinSize` will define the standard size of each cell, and the minimum size of the table (it shows at least one cell). As previously the `UpdateCell` is called to apply data to a cell template. The index passed in is the same `(row, col)` int pair.

```
package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
var data = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}

```

For more info, for example on how to add headers to the table, see the [widget.Table API documentation](/api/v2.5/widget/table.html).

# https://docs.fyne.io/collection/tree

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Tree

The `Tree` collection widget is like the [List](/collection/list) widget (another of the toolkit’s collection widgets) with a multi-level data structure. Like `List` this is designed to help build really performant interfaces when lots of data is being presented. Because of this the widget is not created with all the data embedded, but instead calls out to the data source when needed.

The `Tree` uses callback functions to ask for data when it is required. There are 4 main callbacks, `ChildUIDs`, `IsBranch`, `CreateNode` and `UpdateNode`. The ChildUIDs callback is where we communicate the unique ID of each child node to the one requested. This will be called with the `TreeNodeID` `""` to first get a list of all the IDs that appear inside the root of the tree. The `IsBranch` callback should return `true` if the node is a branch. This is usually true if the node ID has child nodes - but you can have an empty branch.

**It is crucial that the id of each tree node is unique to the tree.** For example if you are building a file manager the ID should be the file path rather than its name.

The other two relate to the content templates.

The `CreateNode` callback returns a new template object, just like list, though there is an additional `bool` parameter that is `true` if the node can have child elements (is a branch). As previously the `UpdateNode` is called to apply data to a cell template. You should update the content according to the `TreeNodeID` and `isBranch` parameters.

```
package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")
	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"a", "b", "c"}
			case "a":
				return []widget.TreeNodeID{"a1", "a2"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "a"
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			text := id
			if branch {
				text += " (branch)"
			}
			o.(*widget.Label).SetText(text)
		})
	myWindow.SetContent(tree)
	myWindow.ShowAndRun()
}

```

# https://docs.fyne.io/binding/

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Data Binding

Data binding is a powerful new addition to the Fyne toolkit that was introduced in version `v2.0.0`. By using data binding we can avoid manually managing many standard objects like `Label`s, `Button`s and `List`s.

The builtin bindings support many primitive types (like `Int`, `String`, `Float` etc), lists (such as `StringList`, `BoolList`) as well as `Map` and `Struct` bindings. Each of these types can be created using a simple constructor function. For example to create a new string binding with a zero value you can use `binding.NewString()`. You can get or set the value of most data bindings using `Get` and `Set` methods.

It is also possible to bind to an existing value using similar functions who’s names start `Bind` and they all accept a pointer to the type bound. To bind to an existing `int` we could use `binding.BindInt(&myInt)`. By keeping a reference to a bound value instead of the original variable we can configure widgets and functions to respond to any changes automatically. If you change the external data directly, be sure to call `Reload`() to ensure that the binding system reads the new value.

```
package main
import (
	"log"
	"fyne.io/fyne/v2/data/binding"
)
func main() {
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)
	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)
}

```

We start learning about [simple values](/binding/simple) widget binding next.

# https://docs.fyne.io/binding/simple

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Binding Simple Widgets

The simplest form of binding a widget is to pass it a bound item as a value instead of a static value. Many widgets provide a `WithData` constructor that will accept a typed data binding item. To set up the binding all you need to do is pass the right type in.

Although this may not look like much of a benefit in the initial code you can then see how it ensures that the displayed content is always up to date with the source of the data. You will notice how we did not need to call `Refresh()` on the `Label` widget or even keep a reference to it and yet it updates appropriately.

```
package main
import (
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Simple")
	str := binding.NewString()
	str.Set("Initial value")
	text := widget.NewLabelWithData(str)
	w.SetContent(text)
	go func() {
		time.Sleep(time.Second * 2)
		str.Set("A new string")
	}()
	w.ShowAndRun()
}

```

In the next step we look at how to set up widgets that edit values through [two way](/binding/twoway) binding.

# https://docs.fyne.io/binding/twoway

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Two-Way Binding

So far we have looked at data binding as a way of keeping our user interface elements up to date. Far more common, however, is the need to update values from the UI widgets and keep the data up to date everywhere. Thankfully the bindings provided in Fyne are “two-way” which means that values can be pushed into them as well as read out. The change in data will be communicated to all connected code without any additional code.

To see this in action we can update the last test app to display a `Label` and an `Entry` that are bound to the same value. By setting this up you can see that editing the value through the entry will update the text in the label as well. This is all possible without calling refresh or referencing the widgets in our code.

By moving your app to use data binding you can stop saving pointers to all your widgets. By instead capturing your data as a set of bound values your user interface can be completely separate code. Cleaner to read and easier to manage.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Two Way")
	str := binding.NewString()
	str.Set("Hi!")
	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))
	w.ShowAndRun()
}

```

Next we will look at how to add [conversions](/binding/conversion) in our data.

# https://docs.fyne.io/binding/conversion

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Data Conversion

So far we have used data binding where the type of data matches the output type (for example `String` and `Label` or `Entry`). Often it will be desirable to present data that is not already in the correct format. To do this the `binding` package provides a number of helpful conversion functions.

Most commonly this will be used to convert different types of data into strings for displaying in `Label` or `Entry` widgets. See in the code how we can convert a `Float` to `String` using `binding.FloatToString`. The original value can be edited by moving the slider. Each time the data changes it will run the conversion code and update any connected widgets.

It is also possible to use format strings to add more natural output for the user interface. You can see that our `short` binding is also converting a `Float` to `String` but by using a `WithFormat` helper we can pass a format string (similar to the `fmt` package) to provide a custom output.

```
package main
import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Conversion")
	f := binding.NewFloat()
	str := binding.FloatToString(f)
	short := binding.FloatToStringWithFormat(f, "%0.0f%%")
	f.Set(25.0)
	w.SetContent(container.NewVBox(
		widget.NewSliderWithData(0, 100.0, f),
		widget.NewLabelWithData(str),
		widget.NewLabelWithData(short),
	))
	w.ShowAndRun()
}

```

Lastly in this section we will look at [list](/binding/list) data.

# https://docs.fyne.io/binding/list

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# List Data

To demonstrate how more complex types can be connected we will look at the `List` widget and how data binding can make it easier to use. Firstly we create a `StringList` data binding, which is a list of `String` data type. Once we have a data of list type we can connect this to the standard `List` widget. To do so we use the `widget.NewListWithData` constructor, much like other widgets.

Comparing this code to the [list tour](/widget/list) You will see 2 main changes, the first is that we pass the data type as the first parameter instead of a length callback function. The second change is the last parameter, our `UpdateItem` callback. The revised version takes a `binding.DataItem` value instead of `widget.ListIndexID`. When using this callback structure we should `Bind` to the template label widget instead of calling `SetText`. This means that if any of the strings change in the data source each affected row of the table will refresh.

```
package main
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")
	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}

```

In our demo code there is an “Append” `Button`, when tapped it will append a new value to the data source. Doing so will automatically trigger the data change handlers and expand the `List` widget to display the new data.

# https://docs.fyne.io/extend/custom-layout

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Building a Custom Layout

In a Fyne application each `Container` arranges its child elements using a simple layout algorithm. Fyne defines many layouts available in the `fyne.io/fyne/v2/layout` package. If you look at the code you will see that they all implement the `Layout` interface.

```
type Layout interface {
	Layout([]CanvasObject, Size)
	MinSize(objects []CanvasObject) Size
}

```

Any application can provide a custom layout to arrange widgets in a non-standard manner. To do this you need to implement the interface above in your own code. To illustrate this we will create a new layout that arranges elements in a diagonal line and is arranged to the bottom left of its container

First we will define a new type, `diagonal`, and define what its minimum size will be. To calculate this we just add the width and height of all child elements (specified as the `[]fyne.CanvasObject` parameter to `MinSize`).

```
import "fyne.io/fyne/v2"
type diagonal struct {
}
func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()
		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

```

To this type we add a `Layout()` function that should move all of the specified objects into the `fyne.Size` specified in the second parameter.

In our implementation we calculate the top-left of the widgets (this is `0` x parameter and has a y position that is the height of the container less the total of all child item heights). From the top position we simply advance each item position by the size of the previous child.

```
func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height - d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

```

That’s all there is to creating a custom layout. Now that the code is all written we can use it as the `layout` parameter to `container.New`. The code below sets up 3 `Label` widgets and place them in a container with our new layout. If you run this application you will see the diagonal widget arrangement and, upon resizing the window, they will align to the bottom left of the available space.

```
package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func main() {
	a := app.New()
	w := a.NewWindow("Diagonal")
	text1 := widget.NewLabel("topleft")
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")
	w.SetContent(container.New(&diagonal{}, text1, text2, text3))
	w.ShowAndRun()
}

```

# https://docs.fyne.io/extend/custom-widget

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Writing a Custom Widget

The standard widgets included with Fyne are designed to support standard user interactions and requirements. As a GUI often has to provide custom functionality it may be necessary to write a custom widget. This article outlines how.

A widget is split into two areas - each implementing a standard interface - the `fyne.Widget` and the `fyne.WidgetRenderer`. The widget defines behaviour and state, with the renderer being used to define how it should be drawn to screen.

### fyne.Widget

A widget in Fyne is simply a stateful canvas object that has its rendering definition separated from the main logic. As you can see from the `fyne.Widget` interface there is not much that must be implemented.

```
type Widget interface {
	CanvasObject
	CreateRenderer() WidgetRenderer
}

```

As a widget needs to be used like any other canvas item we inherit from the same interface. To save writing all the functions required we can make use of the `widget.BaseWidget` type which handles the basics.

Each widget definition will contain much more than the interface requires. It is standard in a Fyne widget to export the fields which define behaviour (just like the primitives defined in the `canvas` package).

For example, look at the `widget.Button` type:

```
type Button struct {
	BaseWidget
	Text string
	Style ButtonStyle
	Icon fyne.Resource
	OnTapped func()
}

```

You can see how each of these items store state about the widget behaviour but nothing about how it is rendered.

### fyne.WidgetRenderer

The widget renderer is responsible for managing a list of `fyne.CanvasObject` primitives that come together to create the design of our widget. It is much like a `fyne.Container` with a custom layout and some additional theme handling.

Every widget must provide a renderer, but it is entirely possible to re-use a renderer from another widget - especially if your widget is a lightweight wrapper around another standard control.

```
type WidgetRenderer interface {
	Layout(Size)
	MinSize() Size
	Refresh()
	Objects() []CanvasObject
	Destroy()
}

```

As you can see the `Layout(Size)` and `MinSize()` functions are similar to the `fyne.Layout` interface, without the `[]fyne.CanvasObject` parameter - this is because a widget does need to be laid out but it controls which objects will be included.

The `Refresh()` method is triggered when the widget this renderer draws has changed or if the theme is altered. In either situation we may need to make adjustments to how it looks. Lastly the `Destroy()` method is called when this renderer is no longer needed so it should clear any resources that would otherwise leak.

Compare again with the button widget - it’s `fyne.WidgetRenderer` implementation is based on the following type:

```
type buttonRenderer struct {
	icon  *canvas.Image
	label *canvas.Text
	shadow *fyne.CanvasObject
	objects []fyne.CanvasObject
	button *Button
}

```

As you can see it has fields to cache the actual image, text and shadow canvas objects for drawing. It keeps track of the slice of objects required by `fyne.WidgetRenderer` as a convenience.

Lastly it keeps a reference to the `widget.Button` for all state information. In the `Refresh()` method it will update the graphical state based on any changes in the underlying `widget.Button` type.

### Bring it together

A basic widget will extend the `widget.BaseWidget` type and declare any state that the widget holds. The `CreateRenderer()` function must exist and return a new `fyne.WidgetRenderer` instance. The widget and driver code in Fyne will ensure that this is cached accordingly - this method may be called many times (for example if a widget is hidden and then shown). If `CreateRenderer()` is called again you should return a new renderer instance as the old one may have been destroyed.

Take care not to keep any important state in your renderer - animation tickers are well suited to that location but user state would not be. A widget that is hidden may have it’s renderer destroyed and if it is shown again the new renderer must be able to reflect the same widget state.

#### Using SimpleRenderer for custom widgets

Custom widgets built from a single `CanvasObject`, for example a container wrapping multiple builtin widgets, can be implemented using `SimpleRenderer`. Below example is a custom widget that can be used as item in a list view, showing a title on the left hand side that will be truncated when too long, and a comment on the right hand side. The constructor would be called from the `CreateItem` function of the list, and the title and comment changed in the `UpdateItem` function:

```
type MyListItemWidget struct {
	widget.BaseWidget
	Title  *widget.Label
	Comment *widget.Label
}
func NewMyListItemWidget(title, comment string) *MyListItemWidget {
	item := &MyListItemWidget{
		Title:  widget.NewLabel(title),
		Comment: widget.NewLabel(comment),
	}
	item.Title.Truncation = fyne.TextTruncateEllipsis
	item.ExtendBaseWidget(item)
	return item
}
func (item *MyListItemWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, item.Comment, item.Title)
	return widget.NewSimpleRenderer(c)
}

```

# https://docs.fyne.io/extend/bundle

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Bundling resources

Go based applications are usually built as a single binary executable, and this is the same for Fyne applications. A single file makes it easier to distribute install our software. Unfortunately GUI applications typically require additional resources to render the user interface. To manage this challenge a Go application can bundle assets into the binary itself. The Fyne toolkit prefers the use of “fyne bundle” as it has various benefits that we will explore below.

The basic approach to bundle an asset is to execute the “fyne bundle” command. This tool has various parameters to customise the output, but in it’s basic form the file to bundle will be converted into Go source code that can be built into your application.

```
$ ls
image.png	main.go
$ fyne bundle -o bundled.go image.png
$ ls
bundled.go	image.png	main.go
$

```

The contents of `bundled.go` will be a list of resource variables that we can then access in our code. For example the code above will result in a file that includes the following:

```
var resourceImagePng = &fyne.StaticResource{
	StaticName: "image.png",
	StaticContent: []byte{
...
	}}

```

As you see the default naming is “resource<Name>.<Ext>”. The name and package used in this file can be customised in command parameters. We can then use this name to, for example, load an image on our canvas:

```
img := canvas.NewImageFromResource(resourceImagePng)

```

A fyne resource is just a collection of bytes with a unique name, so this could be a font, a sound file or any other data you wish to load. Also you can bundle many resources into a single file using the `-append` parameter. If you will be bundling many files it is recommended to save the commands in a go:generate header in one of your go files (not bundled.go):

```
//go:generate fyne bundle -o bundled.go image1.png
//go:generate fyne bundle -o bundled.go -append image2.png

```

If you then change any assets or add new ones then you can update this header and run it with “go generate” to update your `bundled.go` file. You should then add `bundled.go` to version control so others can build your app without needing to run “fyne bundle”.

# https://docs.fyne.io/extend/custom-theme

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Creating a Custom Theme

Applications are able to load custom themes that can apply small changes to the standard theme or provide a completely unique look. A theme will need to implement the functions of `fyne.Theme` interface, which is defined as follows:

```
type Theme interface {
	Color(ThemeColorName, ThemeVariant) color.Color
	Font(TextStyle) Resource
	Icon(ThemeIconName) Resource
	Size(ThemeSizeName) float32
}

```

To apply our theme changes we will first define a new type that that implements this interface.

### Define your theme

We start by defining a new type that will be our theme, a simple empty struct will do:

```
type myTheme struct {}

```

It is a good idea to assert that we implement an interface so that compile errors are closer to the defining type.

```
var _ fyne.Theme = (*myTheme)(nil)

```

At this point you could see compile errors because we still need to implement the methods, we start with color.

#### Customising colors

The `Color` function defined in the `Theme` interface asks us to define a named color and also provides a hint for the variant that the user desires (for example `theme.VariantLight` or `theme.VariantDark`). In our theme we will return a custom background color - using a different value for light and dark.

```
// The color package has to be imported from "image/color".
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}
	return theme.DefaultTheme().Color(name, variant)
}

```

You will see the last line here references the `theme.DefaultTheme()` to look up standard values. This allows us to provide custom values and yet fall back to the standard theme when we don’t want to provide our own values.

Of course colors are simpler than resources, we look at that for custom icons.

#### Overriding default icons

Icons (and Fonts) use `fyne.Resource` as values instead of simple types like `int` (for size) or `color.Color` for colors. We can build our own resource using `fyne.NewStaticResource`, or you could pass in a value that was created using .

```
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		return fyne.NewStaticResource("myHome", homeBytes)
	}

	return theme.DefaultTheme().Icon(name)
}

```

As above we return the default theme icon if we don’t want to provide a specific override.

### Load the theme

Before we can load the theme you will need to implement the `Size` and `Font` methods as well. You can just use these empty implementations if you are happy to use the defaults.

```
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}
func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

```

To set the theme for your app you will need to add the following line of code:

```
app.Settings().SetTheme(&myTheme{})

```

With these changes you can apply your own style, to make small tweaks or provide a completely custom looking application!

# https://docs.fyne.io/extend/extending-widgets

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Extending Widgets

The standard Fyne widgets provide the minimum functionality and customisation to support most use-cases. It may be required at certain times to have more advanced functionality. Rather than have developers build their own widgets it is possible to extend the existing ones.

For example we will extend the icon widget to support being tapped. To do this we declare a new struct that embeds the `widget.Icon` type. We create a constructor function that calls the important `ExtendBaseWidget` function.

```
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)
type tappableIcon struct {
	widget.Icon
}
func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	return icon
}

```

> **Note:** a widget constructor like `widget.NewIcon` may not be used for extension since it already calls `ExtendBaseWidget`.

We then add new functions to implement the `fyne.Tappable` interface, with those functions added the new `Tapped` function will be called every time the user taps our new icon. The interface required has two functions, `Tapped(*PointEvent)` and `TappedSecondary(*PointEvent)`, so we will add both.

```
import "log"
func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}
func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}

```

We can test this new widget using a simple application as follows.

```
import (
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/theme"
)
func main() {
	a := app.New()
	w := a.NewWindow("Tappable")
	w.SetContent(newTappableIcon(theme.FyneLogo()))
	w.ShowAndRun()
}

```

# https://docs.fyne.io/extend/numerical-entry

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Numerical Entry

In the traditional sense, GUI programs have used callbacks to customize actions for widgets. Fyne does not expose inserting custom callbacks to capture events on widgets, but it does not need to. The Go language is plenty extensible to make this work.

Instead we can simply use Type Embedding and extend the widget to only make it possible to enter numerical values.

First create a new type struct, we will call it `numericalEntry`.

```
type numericalEntry struct {
  widget.Entry
}

```

As mentioned in , we follow good practice and create a constructor function that extends the `BaseWidget`.

```
func newNumericalEntry() *numericalEntry {
  entry := &numericalEntry{}
  entry.ExtendBaseWidget(entry)
  return entry
}

```

Now we need to make the entry accept only numbers. This can be done by overriding the `TypedRune(rune)` method that’s part of the `fyne.Focusable` interface. This will allow us to intercept the standard handling of runes received from key presses and only pass through those that we want. Inside this method, we will use a conditional to check if the rune matches any of the numbers between zero and nine. If they do, we delegate it to the standard `TypedRune(rune)` method of the embeded entry. If they do not, we just ignore the inputs. This implementation will only allow integers to be entered, but can easily be extended to check for other keys in the future if necessary.

```
func (e *numericalEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

```

If we want to update the implementation to allow for decimal numers as well, we can simply add `.` and `,` to the list of allowed runes (some languages use commas over dots for decimal notations).

```
func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
			e.Entry.TypedRune(r)
	}
}

```

With this, the entry now only allows the user to enter numerical values when keys are pressed. However, the paste shortcut will still allow text to be entered. To fix this, we can overwrite the `TypedShortcut(fyne.Shortcut)` method that is part of the `fyne.Shortcutable` interface. First we need to do a type assertion to check if the given shortcut is of the type `*fyne.ShortcutPaste`. If it is not, we can just delegate the shortcut back to the embeded entry. If it is, we check if the clipboard content is numerical, by using `strconv.ParseFloat()` (if you want to only allow integers, `strconv.Atoi()` will be just fine), and then delegating the shortcut back to the embeded entry if the clipboard content could be parsed without errors.

```
func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}
	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

```

As a bonus, we can also make sure that mobile operating systems open the numerical keyboard instead of the default keyboard. This can be done by by first importng the `fyne.io/fyne/v2/driver/mobile` package and overwriting the `Keyboard() mobile.KeyboardType` method that is part of the `m̀obile.Keyboardable` interface. Inside the function, we then simply return the `mobile.NumberKeyboard` type.

```
func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

```

In the end, the resulting code could look something like this:

```
package main
import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)
type numericalEntry struct {
	widget.Entry
}
func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}
func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}
func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}
	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}
func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
func main() {
	a := app.New()
	w := a.NewWindow("Numerical")
	entry := newNumericalEntry()
	w.SetContent(entry)
	w.ShowAndRun()
}

```

# https://docs.fyne.io/architecture/geometry

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Geometry

Fyne apps are based on 1 canvas per window. Each canvas has a root CanvasObject which can be a single widget or a Container for many sub-objects whose size and position are controlled by a Layout.

## Position

Each canvas has its origin at the top left (0, 0) every element of the UI may be scaled depending on the output device and so the API does not describe pixels or exact measurements. The position (10, 10) may be 10 pixels right and down from the origin on, for example, a 120DPI monitor but on a HiDPI (or “Retina”) display this will probably be closer to 20 pixels.

Every position referenced by a CanvasObject is relative to it’s parent. This is important for layout algorithms but also for developers in situations such as the `Tappable.Tapped(PointEvent)` handlers. Here the X/Y coordinates will be calculated from the top left of the button not the overall canvas. This is designed to allow code to be as self-contained as possible.

## Pixel size

Like other vector based GUI libraries the Fyne coordinates need to be based on some baseline monitor resolution. All [scaling](/architecture/scaling) is relative to this value. For fyne that resolution is 120DPI. This means that the sizes referred to in `fyne.Size` will be 1=1px when your monitor is 120DPI and scale values are all set to 1. For a HiDPI screen, as mentioned above, the actual DPI may be closer to 240 and on mobile devices it could even be 360 or higher. To manage handle this complexity the toolkit manages scaling internally so your apps will always look the right size. If a user sets the scale to be smaller then their apps will always have smaller than normal fonts, buttons etc, and when they specify larger then your app will scale up to suit.

When compared to we can see that their baseline DPI is , although the maths is similar the actual numbers will be different. This means that device-independent sizes in Fyne use a smaller number to represent the same physical size. For example an icon that is `18` tall in Fyne would be sized at `24` in a standard material design (for example Android) app. This does not matter when building your application, but may be important when working with designers or experts with Material Design.

One time that pixel sizes will matter is if you start loading bitmaps images. Normally these scale appropriately, but if you specify `FillMode=fyne.FillOriginal` then the actual image size will be different on different devices, due to the pixel density. Normally this feature would be used inside a `Scroll` container. Fyne also defines a `canvas.Raster` primitive which will draw pixels exactly at the pixel density of the output device. This enables your code to draw at the highest possible output resolution without knowing details of the device you are running on. If for some reason you need “pixel perfect” positioning you need to multiply `CanvasObject.Size()` by `Canvas.Scale()`.

# https://docs.fyne.io/architecture/scaling

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Scaling

Fyne is built entirely using vector graphics, which means applications written with Fyne will scale to any size beautifully (not just whole number increments). This is a great benefit to the rising popularity of high density displays on mobile devices and high-end computers. The default scale value is calculated to match your operating system - on some systems this is user configuration and on others from your screen’s pixel density (DPI - dots per inch). If a Fyne window is moved to another screen it will re-scale and adjust the window size accordingly! We call this “auto scaling”, and it is designed to keep an app user interface the same size as you change monitor.

You can adjust the size of applications using the `fyne_settings` app or by setting a specific scale using the `FYNE_SCALE` environment variable. These values can make content larger or smaller than the system settings, using “1.5” will make things 50% larger or setting 0.8 will make it 20% smaller.

| ![Hello normal size](/images/architecture/hello-normal.png) Standard size | ![Hello small size](/images/architecture/hello-small.png) FYNE_SCALE=0.5 | ![Hello large size](/images/architecture/hello-large.png) FYNE_SCALE=2.5 |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |

# https://docs.fyne.io/architecture/widgets

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Widgets

Widgets in the Fyne toolkit are designed for a clean and pleasant user interaction, following a standard theme and supporting rapid app development, solid testing and easy maintenance. There are various design considerations that promote that ambition, we explore them in this page.

## Behaviour API

One thing that you will notice about the standard widgets is that the API is all about behaviour and state - but very little that controls the actual look of an element. This is by design. It enables our code, and that of app developers, to focus on the behaviour of a widget so that it’s rendering process can be left to other code. This makes it much easier to test, in fact full applications can be run through unit tests in memory without ever having to render the app.

You can to add new behaviours without needing to worry about how it is rendered. It is also possible to components, an application is not limited to using the provided widget set. When building your own widget you will notice that the rendering details are completely separate from the state - this is part of the design mentioned above. A `WidgetRenderer` (the code that renders a `Widget`) typically holds a reference to the widget that it will be rendering to access state or other information. When a widget state changes then `Refresh()` is called - the renderer will then be asked to refresh and it should update the display to reflect the new state. Custom widgets are recommended to use the current `Theme` but can choose to specify thier own sizes, colours and icons where that seems desirable.

## Content Padding

The standard widgets use the theme specified padding to make appropriate space around their graphical components. The `widget` package uses a standard height and baseline to ensure that provided layouts will align well by default. If you are building a custom widget it is recommended to follow these guidelines.

The value of `theme.Padding()` is used in layouts to space elements of a container, it creates a consistent space around the various parts of an application. Some widgets, however, have content that should be inset from the edges of the extents. Consider `Entry`, It has a background and a border that go out to the edges, but its content should be inset. And so we have standardised the amount of spacing used to inset so that alignment matches.

The standard inset, or content padding, of a widget is defined as `theme.InnerPadding()`. The standard value of padding is `4` and the inner padding is `8`. You can see in `Label` and `Entry` how the (text) content is inset by this much so that their content will align horizontally and vertically when placed next to each other.

![](/images/architecture/widget-inset.png)

It is recommended that custom widgets include similar dimensions so that they fit well alongside the standard widgets.

# https://docs.fyne.io/architecture/organisation

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Organisation and Packages

The Fyne project is split into many packages, each providing different types of functionality. They are as follows:

`fyne.io/fyne/v2`
This import provides the basic definitions common to all Fyne code
including data types and interfaces.
`fyne.io/fyne/v2/app`
The app package provides the APIs that start a new application.
Normally you only require `app.New()` or `app.NewWithID()`.
`fyne.io/fyne/v2/canvas`
The canvas package provides all of the drawing APIs within Fyne.
The complete Fyne toolkit is made up of these primitive graphical types.
`fyne.io/fyne/v2/container`
The container package provides containers that are used to lay out and organise applications.
`fyne.io/fyne/v2/data/binding`
The binding package contains ways of binding data sources to widgets.
`fyne.io/fyne/v2/data/validation`
The validation package provides tooling for validating data inside widgets.
`fyne.io/fyne/v2/dialog`
The dialog package contains dialogs such as confirm, error and file save/open.
`fyne.io/fyne/v2/layout`
The layout package provides various layout implementations for use
with containers (discussed in a later tutorial).
`fyne.io/fyne/v2/storage`
The storage package provides storage access and management functionality.
`fyne.io/fyne/v2/test`
Applications can be tested more easily using the tools within the test
package.
`fyne.io/fyne/v2/widget`
Most graphical applications are created using a collection of widgets.
All the widgets and interactive elements within Fyne are in this package.

# https://docs.fyne.io/faq/layout

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Layout and Widget Size

Intro

## Move and Resize

**Q: How can I move my widget to a different position or resize it?**

**A:** The position and size of elements in a Fyne app are controlled by the layout of the container that they are within. If the elements of your UI are too small consider using a different [layout](/started/layouts) or container.

A new `Window` will expand whatever element is passed to `SetContent()` to fill its size. Each time you add a container to this it will divide up the available space according to the layout. Layouts like `HBox` and `VBox` will shrink content to its `MinSize()` in one dimension or another to pack contents in. Layouts like `Max` or `Border` will expand content to fill the space. By writing a [custom layout](/extend/custom-layout) you could fully control the items within a container.

**Q: Why is my image so small?**

**A:** One of the difficulties in using a fully scalable user interface toolkit such as Fyne is that the coordinate system is device independent. This allows apps to draw at the right resolution or pixel density to get the best results based on the hardware connected. What this means for pixel based images is that their size could vary based on details not known at compile time.

Due to this complication an image loaded using `canvas.NewImageFromFile()` or similar calls will not have a size set, leading to it being very small or appearing to be hidden by default. When placed in an appropriate layout the image will stretch according to its `FillMode` property. If you desire the image to always be set to a certain size (or larger) you can call `Image.SetMinSize()` and specify a device independent size for the image.

## Containers and Layout

**Q: How can I manually control the position of elements**

**A:** In some situations you may want to have complete control over the position and size of elements in a container. To do this you create a container without a layout. The `container.NewWithoutLayout()` function will create a container for manual positioning - you should pass to that constructor a list of the graphical elements that you want to manage in this container.

Once set up then you can use `Move()` and `Resize()` on each element to position it as desired. When doing this be careful to note that it will not adjust as the available space changes - and it does not have an explicit minimum size either. To add either of those features you should replace your manual positioning with a [custom layout](/tutorial/custom-layout).

# https://docs.fyne.io/faq/theme

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Theme and Customisation

In this page we answer some common questions about the design of Fyne themes and widgets.

## Customisation

**Q: How can I change the colour of text for a`Label` widget?**

**A:** All of the standard widgets use the current `Theme` definition to set the colour, font and sizes. To make changes to your application consider using a [custom theme](/extend/custom-theme).

If your application requires text that is a different colour you can use the `canvas.Text` type instead. That allows directly setting the colour and size of the text. Be careful when doing this because a user can choose between light or dark theme variations, so you should test with both.

**Q: How can I remove the background colour from my`Entry` widget?**

**A:** The input background is set by the theme `InputBackground` color. You can change that to `color.Transparent` to remove all input background boxes. It is not possible to edit the style of a single input element - the theme API is designed to give a customisable, but consistent, design.

## Theme API

**Q: How can I use my custom theme written before v2.0.0?**

**A:** Over time you should consider updating to use the new theme API. However it is possible to use a simple adapter that was included to allow the usage of old themes during the transition time. You will find `theme.FromLegacy` function that can adapt an old theme instance to the new API.

```
myTheme := &myOldThemeType{}
updated := theme.FromLegacy(myTheme)
app.Settings().SetTheme(updated)

```

There are no performance penalties when using a theme in this mode, but in a future release this API will be removed.

# https://docs.fyne.io/faq/troubleshoot

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Troubleshooting

Some things can not go as expected during setup or in compiling your first app. We try to address these problems here. Remember you can also check your configuration using the tool.

## Compiler Issues

**Q: command not found: fyne**

**A:** If you have installed the Fyne command using `go install fyne.io/fyne/v2/cmd/fyne@latest` but you see an error when trying to run it then the most likely issue is that your Go install has not set up your PATH environment variable correctly.

Go will install tools to the `bin` directory inside the user’s `GOPATH` location (normally `~/go`) - you can check this by seeing if your `PATH` variable includes this location. If it seems to be missing then you should update your `PATH` environment variable to include `~/go/bin`, or if you have changed install location then you can use `$(go env GOPATH)/bin` instead.

**Q: build constraints exclude all Go files in …**

**A:** If you are cross compiling you may see an error about go files being excluded, followed by a build failure. When doing a standard Go cross-compile it will automatically turn off CGo. To fix this be sure to set `CGO_ENABLED=1` in your compile command.

**Q: cc1.exe: sorry, unimplemented: 64-bit mode not compiled in**

**A:** Windows compilation will sometimes complain that 64 bit mode is not available. This is normally because the wrong compiler is installed, or the configuration is incorrect. If you have followed our [installation instructions](/started/) for MSYS2 and MingW64 then make sure you are using the start menu launcher titled “MSYS2 MinGW 64-bit”.

## Distribution

**Q: Apple macOS says my app is damaged when it is downloaded**

**A:** When files are downloaded on a macOS computer they are marked with a “quarantine” flag so they are checked by the OS for problems. If your application is signed with a certificate purchased from Apple this is not a problem. However if you want to share your software without that cost this error may appear - and on M1/2 computers it is not possible to use the _System Settings_ to allow the app to run.

The fix is to remove the quarantine flag, which you can do by opening the _Terminal_ and executing the following command:

```
sudo xattr -r -d com.apple.quarantine MyApp.app

```

# https://docs.fyne.io/api/v2.5/upgrading.html

#### [ Getting Started ](#collapse-1)

[Getting Started](/started/) [Creating your first Fyne app](/started/hello) [Run Fyne Demo](/started/demo) [Application and RunLoop](/started/apprun) [Updating Content in your GUI](/started/updating) [Window Handling](/started/windows) [Testing Graphical Apps](/started/testing) [Packaging for Desktop](/started/packaging) [Mobile Packaging](/started/mobile) [Run in a Browser](/started/webapp) [App Metadata](/started/metadata) [Distributing to App Stores](/started/distribution) [Compiling for different platforms](/started/cross-compiling)

#### [ Exploring Fyne ](#collapse-2)

[Canvas and CanvasObject](/explore/canvas) [Container and Layouts](/explore/container) [Widget List](/explore/widgets) [Layout List](/explore/layouts) [Dialog List](/explore/dialogs) [Theme Icons](/explore/icons) [Adding Shortcuts to an App](/explore/shortcuts) [Using the Preferences API](/explore/preferences) [Adding app translations](/explore/translations) [System Tray Menu](/explore/systray) [Data Binding](/explore/binding) [Compile Options](/explore/compiling)

#### [ Drawing and Animation ](#collapse-3)

[Rectangle](/canvas/rectangle) [Text](/canvas/text) [Line](/canvas/line) [Circle](/canvas/circle) [Image](/canvas/image) [Raster](/canvas/raster) [Gradient](/canvas/gradient) [Animation](/canvas/animation)

#### [ Containers and Layout ](#collapse-4)

[Box](/container/box) [Grid](/container/grid) [Grid Wrap](/container/gridwrap) [Border](/container/border) [Form](/container/form) [Center](/container/center) [Max](/container/stack) [AppTabs](/container/apptabs)

#### [ Widgets ](#collapse-5)

[Label](/widget/label) [Button](/widget/button) [Entry](/widget/entry) [Choices](/widget/choices) [Form](/widget/form) [ProgressBar](/widget/progressbar) [Toolbar](/widget/toolbar)

#### [ Collections ](#collapse-6)

[List](/collection/list) [Table](/collection/table) [Tree](/collection/tree)

#### [ Data Binding ](#collapse-7)

[Data Binding](/binding/) [Binding Simple Widgets](/binding/simple) [Two-Way Binding](/binding/twoway) [Data Conversion](/binding/conversion) [List Data](/binding/list)

#### [ Extending Fyne ](#collapse-8)

[Building a Custom Layout](/extend/custom-layout) [Writing a Custom Widget](/extend/custom-widget) [Bundling resources](/extend/bundle) [Creating a Custom Theme](/extend/custom-theme) [Extending Widgets](/extend/extending-widgets) [Numerical Entry](/extend/numerical-entry)

#### [ Architecture ](#collapse-9)

[Geometry](/architecture/geometry) [Scaling](/architecture/scaling) [Widgets](/architecture/widgets) [Organisation and Packages](/architecture/organisation)

#### [ Frequently Asked Questions ](#collapse-10)

[Layout and Widget Size](/faq/layout) [Theme and Customisation](/faq/theme) [Troubleshooting](/faq/troubleshoot)

#### [ API Documentation ](#collapse-99)

[Upgrade to v2.5](/api/v2.5/upgrading.html) [Fyne API v2.5](/api/v2.5/) [app](/api/v2.5/app/) [canvas](/api/v2.5/canvas/) [container](/api/v2.5/container/) [data/binding](/api/v2.5/data/binding/) [data/validation](/api/v2.5/data/validation/) [dialog](/api/v2.5/dialog/) [driver](/api/v2.5/driver/) [driver/desktop](/api/v2.5/driver/desktop/) [driver/mobile](/api/v2.5/driver/mobile/) [driver/software](/api/v2.5/driver/software/) [lang](/api/v2.5/lang/) [layout](/api/v2.5/layout/) [storage](/api/v2.5/storage/) [storage/repository](/api/v2.5/storage/repository/) [test](/api/v2.5/test/) [theme](/api/v2.5/theme/) [widget](/api/v2.5/widget/)

#### [ List My App ](/submit/)

# Upgrading to v2.5

The 2.5 release is fully backward compatible with 2.4.5 and earlier, so upgrading is as simple as updating the version of code you compile with. From thie v2.5.0 release of Fyne it requires Go 1.19, and so all projects are assumed to be using Go modules.

## Updating

Open your `go.mod` file and alter the the `require` line to use version `v2.5.0`, or you can execute the following command inside the directory:

```
go get fyne.io/fyne/v2@v2.5.0

```

The next time you build or run your app it will be using the 2.5 API, showing the new curved corners in input widgets and selection.

## Fyne command

You should update the `fyne` tool for v2.5.0 to get some of the new features and bug fixes. You can make the upgrade by using the `go get` command similarly to above:

```
go install fyne.io/fyne/v2/cmd/fyne@v2.5.0

```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

Although this release is backwards compatible so your code will compile and run as expected, there are some changes which you may notice.

- Fyne now requires Go 1.19 - be sure to upgrade if you were on an earlier version
- Widgets are now able to have a non-standard theme if placed inside a `ThemeOverride` container, so theme code should use `theme.ForWidget()` instead of the static `theme` helper methods
- Mobile apps using a software keyboard will now be adjusted in size to fit the space - consider adding scrolling to taller containers
