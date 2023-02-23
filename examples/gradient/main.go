package main

import "image"
import "image/png"

import "github.com/tinne26/badcolor"
import "github.com/tinne26/badcli"

const ProgramDescription = "Generates a color gradient and exports it as a PNG file."

func main() {
	// ---- program setup ----
	// prepare relevant program variables
	from   := badcli.NewColorString(255, 0, 0)
	to     := badcli.NewColorString(0, 255, 255)
	steps  := badcli.NewBoundedInt(6, 1, 8192)
	width  := badcli.NewBoundedInt(512, 1, 8192)
	height := badcli.NewBoundedInt(256, 1, 8192)
	bwidth := badcli.NewBoundedInt(0, 1, 1024)
	output := badcli.NewFilePath("gradient.png", "png", "PNG")

	// setup cli and flags using the previous program variables
	cli := badcli.NewCLI("badcolor-gradient", ProgramDescription)
	cli.AddUsageSection(
		"Color flags accept the following formats:\n" + badcli.ColorStringFormatsInfo,
	)
	cli.RegisterFlag("from", "gradient start color", from)
	cli.RegisterFlag("to"  , "gradient end color", to)
	cli.RegisterFlag("steps" , "number of colors in the gradient", steps)
	cli.RegisterFlag("width" , "gradient image width, in pixels", width)
	cli.RegisterFlag("height", "gradient image height, in pixels", height)
	cli.RegisterFlag("bwidth", "width of a color band in the gradient, in pixels", bwidth)
	cli.RegisterFlag("output", "output image file path", output)
	cli.DisallowExtraArgs()
	cli.ParseArguments()

	// warnings
	if cli.AllFlagsSetByUser("bwidth", "width") {
		cli.Warn("both --width and --bwidth given, prioritizing --bwidth")
	}
	
	// generate random colors if missing
	if !cli.FlagSetByUser("from") { randomizeColorFlag(cli, "from") }
	if !cli.FlagSetByUser("to"  ) { randomizeColorFlag(cli, "to"  ) }

	// ---- gradient operations ----
	// create gradient
	gradient := make([]badcolor.RGBA, steps.Value())
	cli.Printf("...generating gradient with %d colors\n", steps.Value())
	badcolor.OklabGradient(gradient, from.RGBA8(), to.RGBA8())

	// draw gradient to image
	var rect image.Rectangle
	if cli.FlagSetByUser("bwidth") {
		rect = image.Rect(0, 0, bwidth.Value()*steps.Value(), height.Value())
	} else {
		rect = image.Rect(0, 0, width.Value(), height.Value())
	}

	cli.Printf("...creating %dx%d gradient image\n", rect.Dx(), rect.Dy())
	img := image.NewRGBA(rect)
	for x := 0; x < rect.Dx(); x++ {
		// find appropriate color for the column
		progress := float64(x)/float64(rect.Dx() - 1) // between [0, 1]
		index := int(float64(len(gradient))*progress)
		if index == len(gradient) { index = len(gradient) - 1 }
		clr := gradient[index]

		// apply color to the whole column
		for y := 0; y < rect.Dy(); y++ {
			img.SetRGBA(x, y, clr)
		}
	}

	// ---- save results ----
	cli.Print("...saving gradient image to file\n")
	cli.ExportImage(output.Value(), img, png.Encode)
	cli.Printf("Gradient successfully saved at '%s'\n", output.Reference())
}

// Helper function to randomize a color flag.
func randomizeColorFlag(cli *badcli.CLI, flagName string) {
	target := cli.GetFlagValue(flagName).(*badcli.ColorString)
	clr, err := badcolor.RngRGB()
	if err != nil {
		cli.Fatal("failed to generate random --%s color: %s", flagName, err)
	}
	target.SetRGBA8(clr)
	cli.Printf("...randomizing --%s color as %s\n", flagName, target.String())
}
