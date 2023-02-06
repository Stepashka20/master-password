package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"github.com/AllenDang/giu"
	"golang.org/x/crypto/scrypt"
	"image"
	"log"
	"strconv"
)

var (
	username          string
	masterPassword    string
	site              string
	login                   = false
	encVariants             = []string{"Maximum", "Long", "Medium", "Basic", "Short", "PIN", "Name", "Phrase"}
	encSelected       int32 = 0
	passNum           int32 = 1
	generatedPassword       = ""
	auto                    = false
)

var characters = map[byte]string{
	'V': "AEIOU",
	'C': "BCDFGHJKLMNPQRSTVWXYZ",
	'v': "aeiou",
	'c': "bcdfghjklmnpqrstvwxyz",
	'A': "AEIOUBCDFGHJKLMNPQRSTVWXYZ",
	'a': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz",
	'n': "0123456789",
	'o': "@&%?,=[]_:-+*$#!'^~;()/.",
	'x': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()",
	' ': " ",
}

var templates = map[string][]string{
	"templateMaximum": {
		"anoxxxxxxxxxxxxxxxxx",
		"axxxxxxxxxxxxxxxxxno",
	},
	"templateLong": {
		"CvcvnoCvcvCvcv",
		"CvcvCvcvnoCvcv",
		"CvcvCvcvCvcvno",
		"CvccnoCvcvCvcv",
		"CvccCvcvnoCvcv",
		"CvccCvcvCvcvno",
		"CvcvnoCvccCvcv",
		"CvcvCvccnoCvcv",
		"CvcvCvccCvcvno",
		"CvcvnoCvcvCvcc",
		"CvcvCvcvnoCvcc",
		"CvcvCvcvCvccno",
		"CvccnoCvccCvcv",
		"CvccCvccnoCvcv",
		"CvccCvccCvcvno",
		"CvcvnoCvccCvcc",
		"CvcvCvccnoCvcc",
		"CvcvCvccCvccno",
		"CvccnoCvcvCvcc",
		"CvccCvcvnoCvcc",
		"CvccCvcvCvccno",
	},
	"templateMedium": {
		"CvcnoCvc",
		"CvcCvcno",
	},
	"templateShort": {
		"Cvcn",
	},
	"templateBasic": {
		"aaanaaan",
		"aannaaan",
		"aaannaaa",
	},
	"templatePIN": {
		"nnnn",
	},
	"templateName": {
		"cvccvcvcv",
	},
	"templatePhrase": {
		"cvcc cvc cvccvcv cvc",
		"cvc cvccvcvcv cvcv",
		"cv cvccv cvc cvcvccv",
	},
}

//go:embed key.png
var iconData []byte
var icons []image.Image

func generateUniqueKey() string {
	// Generate a unique string from the username, master password, site and password number
	seed := hmac.New(sha256.New, []byte(masterPassword+strconv.Itoa(len(site))+site+strconv.Itoa(int(passNum))))
	salt := []byte(site + string(rune(len(username))) + username + hex.EncodeToString(seed.Sum(nil)))
	dk, _ := scrypt.Key([]byte(masterPassword), salt, 32768, 8, 2, 64)
	final := base64.StdEncoding.EncodeToString(dk)
	return final
}

func generatePassword() {
	originalString := generateUniqueKey()
	result := ""
	var lettersCode []int
	// Convert the string to a slice of integers
	for _, r := range originalString {
		lettersCode = append(lettersCode, int(r))
	}
	templateType := "template" + encVariants[encSelected]
	resultTemplate := templates[templateType][lettersCode[0]%len(templates[templateType])]
	i := 0
	// Generate the password using the template
	for _, characterClass := range resultTemplate {
		characters := characters[byte(characterClass)]
		result += string(characters[lettersCode[i+1]%len(characters)])
		i++
	}
	generatedPassword = result

}

// Auto-generate password when the user changes the settings
func autoGeneratePassword() {
	if auto {
		generatePassword()
	}
}
func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			if !login {
				giu.PopupModal("Login:##" + strconv.Itoa(1)).Flags(6).Layout(
					giu.Custom(func() {
						giu.Layout{
							giu.Dummy(180, 0),
							giu.Label("Username:"),
							giu.InputText(&username).Size(-1),
							giu.Label("Master password:"),
							giu.InputText(&masterPassword).Size(-1).Flags(giu.InputTextFlagsPassword),
							giu.Button("Login").Size(-1, 30).OnClick(func() {
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
		giu.Layout{
			giu.Label("Generate password for " + username),
			giu.Separator(),
			giu.Row(
				giu.Label("Site name:"),
				giu.InputText(&site).Size(-1).OnChange(autoGeneratePassword),
			),
			giu.Style().SetDisabled(site == "").To(
				giu.Row(
					giu.Combo("##splitter", encVariants[encSelected], encVariants, &encSelected).Size(-1).OnChange(autoGeneratePassword),
					giu.Tooltip("Type of password to generate"),

					giu.InputInt(&passNum).Size(-1).Flags(giu.InputTextFlagsCharsDecimal).OnChange(autoGeneratePassword),
					giu.Tooltip("Number of password"),
				),

				giu.Row(
					giu.Checkbox("", &auto),
					giu.Tooltip("Auto generate password"),
					giu.Button("Generate").Size(-1, 0).OnClick(generatePassword),
				),
				giu.Style().SetDisabled(generatedPassword == "").To(
					giu.Dummy(0, 10),
					giu.Align(giu.AlignCenter).To(
						giu.Label(generatedPassword),
					),
					giu.Dummy(0, 10),
					giu.Button("copy").Size(-1, 0).OnClick(func() {
						giu.Context.GetPlatform().SetClipboard(generatedPassword)

					}),
				),
			),
		},
	)
}

func main() {
	wnd := giu.NewMasterWindow("Master Password", 318, 180, giu.MasterWindowFlagsNotResizable)
	// Set the icon of the window
	icons = make([]image.Image, 0)
	icon1, _, err := image.Decode(bytes.NewReader(iconData))
	if err != nil {
		log.Fatal(err)
	}
	icons = append(icons, icon1)
	wnd.SetIcon(icons)

	wnd.Run(loop)

}
