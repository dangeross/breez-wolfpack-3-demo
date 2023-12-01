package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/breez/breez-sdk-go/breez_sdk"
	"github.com/dangeross/breez-wolfpack-3-demo/internal/util"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type Cli struct {
	*grumble.App

	log *logrus.Logger
	sdk *breez_sdk.BlockingBreezServices

	dataDir string
	config  *Config
}

func Run() {
	c := &Cli{
		App: grumble.New(&grumble.Config{
			Name:                  "sdk",
			Description:           "A demo Go Breez SDK CLI for Wolfpack 3",
			HistoryFile:           "history.txt",
			Prompt:                "sdk> ",
			PromptColor:           color.New(color.FgHiCyan, color.Bold),
			HelpHeadlineColor:     color.New(color.FgHiCyan),
			HelpHeadlineUnderline: true,
			HelpSubCommands:       true,

			Flags: func(f *grumble.Flags) {
				f.String("d", "data_dir", "", "optional data dir, default to current directory")
			},
		}),
	}

	c.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("__________                                                                 ")
		a.Println("\\______   \\_______   ____   ____ ________                                  ")
		a.Println(" |    |  _/\\_  __ \\_/ __ \\_/ __ \\\\___   /                                  ")
		a.Println(" |    |   \\ |  | \\/\\  ___/\\  ___/ /    /                                   ")
		a.Println(" |______  / |__|    \\___  >\\___  >_____ \\                                  ")
		a.Println("        \\/              \\/     \\/      \\/                                  ")
		a.Println("           __      __      .__   _____                     __     ________  ")
		a.Println("          /  \\    /  \\____ |  |_/ ____\\__________    ____ |  | __ \\_____  \\ ")
		a.Println("          \\   \\/\\/   /  _ \\|  |\\   __\\\\____ \\__  \\ _/ ___\\|  |/ /   _(__  < ")
		a.Println("           \\        (  <_> )  |_|  |  |  |_> > __ \\\\  \\___|    <   /       \\")
		a.Println("            \\__/\\  / \\____/|____/__|  |   __(____  /\\___  >__|_ \\ /______  /")
		a.Println("                 \\/                   |__|       \\/     \\/     \\/        \\/ ")
		a.Println()
	})

	c.OnInit(func(app *grumble.App, flags grumble.FlagMap) (err error) {
		c.dataDir = flags.String("data_dir")

		if len(c.dataDir) == 0 {
			if c.dataDir, err = os.Getwd(); err != nil {
				return fmt.Errorf("failed to get the current working directory: %v", err)
			}
		}

		if c.dataDir, err = filepath.Abs(c.dataDir); err != nil {
			return err
		}

		if err = os.MkdirAll(c.dataDir, os.ModePerm); err != nil {
			return err
		}

		app.Config().HistoryFile = filepath.Join(c.dataDir, "history.txt")

		// Init config
		if err := c.readConfig(); err != nil {
			panic(err)
		}

		// Init log
		c.initLog()
		breez_sdk.SetLogStream(c)

		return c.load()
	})

	c.OnClose(func() (err error) {
		if c.sdk != nil {
			c.sdk.Disconnect()
			c.sdk.Destroy()
			c.sdk = nil
		}

		return
	})

	grumble.Main(c.App)
}

func (c *Cli) PrettyPrint(i interface{}) {
	if b, err := json.MarshalIndent(i, "", "  "); err == nil {
		c.Println(string(b))
	}
}

func (c *Cli) PrintSuccess(str string) {
	color.New(color.FgGreen, color.Bold).Fprintln(c.App, str)
}

func (c *Cli) Prompt(prompt string, a ...any) (string, error) {
	var response string

	c.Printf(prompt, a...)
	_, err := fmt.Scanln(&response)

	return response, err
}

func (c *Cli) PromptNil(prompt string, a ...any) (*string, error) {
	response, err := c.Prompt(prompt, a...)
	if err != nil {
		return nil, err
	}

	return util.NilString(response), err
}

func (c *Cli) PromptUint64(prompt string, a ...any) (uint64, error) {
	var response string

	c.Printf(prompt, a...)
	fmt.Scanln(&response)

	return strconv.ParseUint(response, 10, 64)
}
