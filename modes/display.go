package modes

import (
    "log"

    "github.com/z3oxs/ghost/utils"
)

func Display(save string, clipboard, output, upload, file bool, display int) {
    displays := utils.GetDisplays()

    if display > len(displays) {
        log.Fatal("Invalid screen selected.")

    }

    displaySelected := displays[display]

    screenshot := utils.CaptureRect (
        displaySelected.X,
        displaySelected.Y,
        displaySelected.Width,
        displaySelected.Height,
        save,
    )

    utils.SaveHandler(clipboard, output, upload, file, screenshot)
}
