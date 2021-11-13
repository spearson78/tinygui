module github.com/spearson78/tinygui

go 1.17

require (
	github.com/sago35/tinydisplay v0.0.0-20211221140237-b980dfd11c01
	github.com/shopspring/decimal v1.3.1
	github.com/spearson78/tinyamifont v0.0.0-20220201082506-452b1c82d8bd
	tinygo.org/x/drivers v0.19.0
	tinygo.org/x/tinydraw v0.0.0-20220125063109-43cae6615eb5
)

require (
	fyne.io/fyne/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fredbi/uri v0.0.0-20181227131451-3dcfdacbaaf3 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211213063430-748e38ca8aec // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/goki/freetype v0.0.0-20220119013949-7a161fd3728c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/srwiley/oksvg v0.0.0-20220128195007-1f435e4c2b44 // indirect
	github.com/srwiley/rasterx v0.0.0-20220128185129-2efea2b9ea41 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/yuin/goldmark v1.4.4 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace tinygo.org/x/drivers => github.com/spearson78/drivers v0.18.1-0.20220201074301-10484ed1d809

replace tinygo.org/x/tinydraw => github.com/spearson78/tinydraw v0.0.0-20220201143000-b08d64c952ed
