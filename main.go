

package main

import (
	"github.com/AllenDang/giu"
	"strconv"
)

var (
	username string
	masterPassword string
	site string
	login    bool = false
	encVariants = []string{"Maximum Security", "Long", "Medium", "Short", "Basic", "PIN"}
	encSelected  int32 = 0
	passNum int32 = 1
	generatedPassword string = "TestPassword#234"
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			if !login {
				giu.PopupModal("Login:##"+strconv.Itoa(1)).Flags(6).Layout(
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
								if (username == "") || (masterPassword == "") { return }
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
				giu.Row(
					//giu.Label("Type:"),

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

				giu.Separator(),
				giu.Row(
					giu.InputText(&generatedPassword).Size(318-50).Flags(giu.InputTextFlagsReadOnly),
					giu.Button("Copy").Size(-1, 0).OnClick(func() {
						// TODO
					}),
				),
			}.Build()
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Master Password", 318, 167, giu.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}