# Go-Listfont

Cross platform list font of system base on [go-findfont](https://github.com/flopp/go-findfont).

## Supported platform

- [x] Linux
- [x] Darwin
- [x] Window

## Installation

```
go get -u github.com/Nguyen-Hoang-Nam/go-listfont
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/Nguyen-Hoang-Nam/go-listfont"
)

func main() {
  fonts := listfont.List()

  for _, font := range fonts {
    fmt.Printf("%s: %s:style=%s\n", font.Uri, font.Family, font.SubFamily)
  }
}
```

### Linux

```
$ go run main.go | grep -i noto
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans CJK KR:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans CJK SC:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans CJK TC:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans CJK HK:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans Mono CJK JP:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans Mono CJK KR:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans Mono CJK SC:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans Mono CJK TC:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Bold.ttc: Noto Sans Mono CJK HK:style=Bold
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans CJK KR:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans CJK SC:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans CJK TC:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans CJK HK:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans Mono CJK JP:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans Mono CJK KR:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans Mono CJK SC:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans Mono CJK TC:style=Regular
/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc: Noto Sans Mono CJK HK:style=Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Bold.ttc: Noto Serif CJK KR:style=Bold
/usr/share/fonts/opentype/noto/NotoSerifCJK-Bold.ttc: Noto Serif CJK SC:style=Bold
/usr/share/fonts/opentype/noto/NotoSerifCJK-Bold.ttc: Noto Serif CJK TC:style=Bold
/usr/share/fonts/opentype/noto/NotoSerifCJK-Bold.ttc: Noto Serif CJK HK:style=Bold
/usr/share/fonts/opentype/noto/NotoSerifCJK-Regular.ttc: Noto Serif CJK KR:style=Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Regular.ttc: Noto Serif CJK SC:style=Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Regular.ttc: Noto Serif CJK TC:style=Regular
/usr/share/fonts/opentype/noto/NotoSerifCJK-Regular.ttc: Noto Serif CJK HK:style=Regular
/usr/share/fonts/truetype/noto/NotoColorEmoji.ttf: Noto Color Emoji:style=Regular
/usr/share/fonts/truetype/noto/NotoMono-Regular.ttf: Noto Mono:style=Regular
/usr/share/fonts/truetype/noto/NotoSansMono-Bold.ttf: Noto Sans Mono:style=Bold
/usr/share/fonts/truetype/noto/NotoSansMono-Regular.ttf: Noto Sans Mono:style=Regular
```

### Window

```
D:\Dev\listfont>go run main.go
C:\Windows\Fonts\Candara.ttf: Candara:style=Regular
C:\Windows\Fonts\Candarab.ttf: Candara:style=Bold
C:\Windows\Fonts\Candarai.ttf: Candara:style=Italic
C:\Windows\Fonts\Candaral.ttf: Candara:style=Light
C:\Windows\Fonts\Candarali.ttf: Candara:style=Light Italic
C:\Windows\Fonts\Candaraz.ttf: Candara:style=Bold Italic
C:\Windows\Fonts\Gabriola.ttf: Gabriola:style=Regular
C:\Windows\Fonts\Inkfree.ttf: Ink Free:style=Regular
C:\Windows\Fonts\LeelUIsl.ttf: Leelawadee UI Semilight:style=Normal
C:\Windows\Fonts\LeelaUIb.ttf: Leelawadee UI:style=Negreta
C:\Windows\Fonts\LeelawUI.ttf: Leelawadee UI:style=Normal
C:\Windows\Fonts\Nirmala.ttf: Nirmala UI:style=Regular
C:\Windows\Fonts\NirmalaB.ttf: Nirmala UI:style=Bold
C:\Windows\Fonts\NirmalaS.ttf: Nirmala UI Semilight:style=Regular
C:\Windows\Fonts\SansSerifCollection.ttf: Sans Serif Collection:style=Regular
C:\Windows\Fonts\SegUIVar.ttf: Segoe UI Variable:style=Regular
C:\Windows\Fonts\SegoeIcons.ttf: Segoe Fluent Icons:style=Regular
C:\Windows\Fonts\SitkaVF-Italic.ttf: Sitka Text:style=Italic
C:\Windows\Fonts\SitkaVF.ttf: Sitka Text:style=Regular
C:\Windows\Fonts\VNI 06 Springtime.ttf: VNI 06 Springtime:style=Bold
C:\Windows\Fonts\VNI 08 Springtime2.ttf: VNI 08 Springtime2:style=BoldOblique
C:\Windows\Fonts\VNI 09 Baroque.ttf: VNI 09 Baroque :style=Regular
```

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit)