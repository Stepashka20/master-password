package main

import (
	"bytes"
	_ "embed"
	"github.com/AllenDang/giu"
	"image"
	"log"
	"strconv"
)

var (
	username          string
	masterPassword    string
	site              string
	login             bool   = false
	encVariants              = []string{"Maximum Security", "Long", "Medium", "Short", "Basic", "PIN"}
	encSelected       int32  = 0
	passNum           int32  = 1
	generatedPassword string = "TestPassword#234"
)

//go:embed key.png
var iconData []byte
var icons []image.Image

func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			if !login {
				giu.PopupModal("Login:##" + strconv.Itoa(1)).Flags(6).Layout(
					giu.Custom(func() {
						const buttonW, buttonH = -1, 30
						//_, availableH := giu.GetAvailableRegion()
						//_, itemSpacingY := giu.GetItemSpacing()
						giu.Layout{
							giu.Dummy(180, 0),
							giu.Label("Username:"),
							giu.InputText(&username).Size(-1),
							giu.Label("Master password:"),
							giu.InputText(&masterPassword).Size(-1).Flags(giu.InputTextFlagsPassword),
							giu.Button("Login").Size(buttonW, buttonH).OnClick(func() {
								if (username == "") || (masterPassword == "") {
									return
								}
								login = true
							}),
						}.Build()
					}),
				).Build()
				giu.OpenPopup("Login:##" + strconv.Itoa(1))
				giu.Update()
			}
		}),
		giu.Custom(func() { // TODO delete Custom
			giu.Layout{
				giu.Label("Generate password for " + username),
				giu.Separator(),
				giu.Row(
					giu.Label("Site name:"),
					giu.InputText(&site).Size(-1),
				),
				giu.Style().SetDisabled(site == "").To(
					giu.Row(
						giu.Combo("##splitter", encVariants[encSelected], encVariants, &encSelected).Size(-1),
						giu.Tooltip("Type of password to generate"),

						giu.InputInt(&passNum).Size(-1).Flags(giu.InputTextFlagsCharsDecimal),
						giu.Tooltip("Number of password"),
					),

					giu.Row(
						giu.Button("Generate").Size(-1, 30).OnClick(func() {
							// TODO

						}),
					),
					giu.Dummy(0, 10),
					giu.Align(giu.AlignCenter).To(
						giu.Label(generatedPassword),
					),
					giu.Dummy(0, 10),
					giu.Button("copy").Size(-1, 0).OnClick(func() {
						// TODO

					}),
				),
			}.Build()
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Master Password", 318, 180, giu.MasterWindowFlagsNotResizable)
	icons = make([]image.Image, 0)

	icon1, _, err := image.Decode(bytes.NewReader(iconData))
	if err != nil {
		log.Fatal(err)
	}

	icons = append(icons, icon1)

	wnd.SetIcon(icons)
	wnd.Run(loop)

}
