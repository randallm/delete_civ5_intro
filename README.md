# delete_civ5_intro

## Distribution

Use https://github.com/machinebox/appify to package:

```sh
go build
~/go/bin/appify -name "Remove CivV Intro" -icon Smile-thumbs-up.png civ5
codesign --force --deep -s - Remove\ CivV\ Intro.app
```