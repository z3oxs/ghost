package modes

import (
    "image"

    "github.com/go-vgo/robotgo"
    "github.com/z3oxs/ghost/utils"
)

func Selection(save, format string, clipboard, output, upload bool) {
    var x1, x2, y1, y2 int
    var screenshot image.Image

    mouseFirstClick := robotgo.AddMouse("left")

    if mouseFirstClick {
        x1, y1 = robotgo.GetMousePos()

    }

    mouseLastClick := robotgo.AddMouse("left")

    if mouseLastClick {
        x2, y2 = robotgo.GetMousePos()

    } 

    if x1 < x2 {
        screenshot = robotgo.CaptureImg(x1, y1, x2 - x1, y2 - y1) 

    } else {
        screenshot = robotgo.CaptureImg(x2, y2, x1 - x2, y1 - y2)

    }

    utils.PlaySound("/opt/ghost/screenshot.wav")

    if clipboard {
        utils.SaveToClipboard(screenshot)

    }

    if save != "" {
        robotgo.Save(screenshot, save)

    }

    if output {
        utils.OutputToStdout(screenshot, format)

    }

    if upload {
        utils.UploadImage(screenshot, format)

    }
}
