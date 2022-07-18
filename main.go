package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/micmonay/keybd_event"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/mainthread"
)

const (
	F1 = hotkey.Key(112)
	F2 = hotkey.Key(113)
	F3 = hotkey.Key(114)
	F4 = hotkey.Key(115)
)

func main() { mainthread.Init(fn) }
func fn() {
	go func() {
		err := OnKeyPress([]hotkey.Modifier{hotkey.ModCtrl}, F1, func() {
			kb, _ := keybd_event.NewKeyBonding()
			kb.SetKeys(keybd_event.VK_MEDIA_PLAY_PAUSE)
			kb.Launching()
		})
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := OnKeyPress([]hotkey.Modifier{hotkey.ModCtrl}, F2, func() {
			kb, _ := keybd_event.NewKeyBonding()
			kb.SetKeys(keybd_event.VK_MEDIA_PREV_TRACK)
			kb.Launching()
		})
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := OnKeyPress([]hotkey.Modifier{hotkey.ModCtrl}, F3, func() {
			kb, _ := keybd_event.NewKeyBonding()
			kb.SetKeys(keybd_event.VK_MEDIA_NEXT_TRACK)
			kb.Launching()
		})
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := OnKeyPress([]hotkey.Modifier{hotkey.ModCtrl}, F4, TakeScreenShot)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := OnKeyPress([]hotkey.Modifier{hotkey.ModWin, hotkey.ModAlt}, F4, TakeScreenShot)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := OnKeyPress([]hotkey.Modifier{hotkey.ModWin, hotkey.ModAlt}, hotkey.KeyO, TakeScreenShot)
		if err != nil {
			panic(err)
		}
	}()

	select {}
}

func OnKeyPress(mods []hotkey.Modifier, key hotkey.Key, handler func()) error {
	hk := hotkey.New(mods, key)
	if err := hk.Register(); err != nil {
		return err
	}

	fmt.Printf("Listening for %v\n", hk)
	for range hk.Keydown() {
		handler()
	}

	return nil
}

func TakeScreenShot() {
	picPath := "C:\\Users\\Ali\\Pictures\\Screenshots"

	sc, _ := screenshot.CaptureDisplay(0)

	rawData := bytes.NewBuffer([]byte{})
	png.Encode(rawData, sc)
	data, _ := ioutil.ReadAll(rawData)

	fileName := filepath.Join(picPath, fmt.Sprintf(`Zephyrus_%d.png`, time.Now().Unix()))
	ioutil.WriteFile(fileName, data, os.ModePerm)

	clipboard.Write(clipboard.FmtImage, data)
}
