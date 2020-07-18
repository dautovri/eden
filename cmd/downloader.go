package cmd

import (
	"fmt"
	"github.com/lf-edge/eden/pkg/defaults"
	"github.com/lf-edge/eden/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	eveArch   string
	eveTag    string
	outputDir string
)

var downloaderCmd = &cobra.Command{
	Use: "download",
}
var downloadEVECmd = &cobra.Command{
	Use:   "eve",
	Short: "download eve live image from docker",
	Long:  `Download eve live image from docker.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		assignCobraToViper(cmd)
		viperLoaded, err := utils.LoadConfigFile(configFile)
		if err != nil {
			return fmt.Errorf("error reading config: %s", err.Error())
		}
		if viperLoaded {
			eveTag = viper.GetString("eve.tag")
			eveArch = viper.GetString("eve.arch")
			eveHV = viper.GetString("eve.hv")
			adamDist = utils.ResolveAbsPath(viper.GetString("adam.dist"))
			eveImageFile = utils.ResolveAbsPath(viper.GetString("eve.image-file"))
			devModel = viper.GetString("eve.devmodel")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		format := "qcow2"
		if devModel == defaults.DefaultRPIModel {
			format = "raw"
		}
		if err := utils.DownloadEveLive(adamDist, eveImageFile, eveArch, eveHV, eveTag, format); err != nil {
			log.Fatal(err)
		} else {
			if devModel == defaults.DefaultRPIModel {
				log.Infof("Write file %s to sd (it is in raw format)", eveImageFile)
			}
			fmt.Println(eveImageFile)
		}
	},
}
var downloadEVERootFSCmd = &cobra.Command{
	Use:   "eve-rootfs",
	Short: "download eve rootfs image from docker",
	Long:  `Download eve rootfs image from docker.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		assignCobraToViper(cmd)
		viperLoaded, err := utils.LoadConfigFile(configFile)
		if err != nil {
			return fmt.Errorf("error reading config: %s", err.Error())
		}
		if viperLoaded {
			eveTag = viper.GetString("eve.tag")
			eveArch = viper.GetString("eve.arch")
			eveHV = viper.GetString("eve.hv")
			if outputDir == "" {
				outputDir = filepath.Dir(utils.ResolveAbsPath(viper.GetString("eve.image-file")))
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if image, err := utils.DownloadEveRootFS(outputDir, eveArch, eveHV, eveTag); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(image)
		}
	},
}

func downloaderInit() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	downloaderCmd.AddCommand(downloadEVECmd)
	downloadEVECmd.Flags().StringVarP(&eveTag, "eve-tag", "", defaults.DefaultEVETag, "tag to download")
	downloadEVECmd.Flags().StringVarP(&eveArch, "eve-arch", "", runtime.GOARCH, "arch of EVE")
	downloadEVECmd.Flags().StringVarP(&eveHV, "eve-hv", "", defaults.DefaultEVEHV, "HV of EVE (kvm or xen)")
	downloadEVECmd.Flags().StringVarP(&eveImageFile, "image-file", "i", "", "path for image drive")
	downloadEVECmd.Flags().StringVarP(&adamDist, "adam-dist", "", path.Join(currentPath, defaults.DefaultDist, defaults.DefaultAdamDist), "adam dist to start")
	downloaderCmd.AddCommand(downloadEVERootFSCmd)
	downloadEVERootFSCmd.Flags().StringVarP(&eveTag, "eve-tag", "", defaults.DefaultEVETag, "tag to download")
	downloadEVERootFSCmd.Flags().StringVarP(&eveArch, "eve-arch", "", runtime.GOARCH, "arch of EVE")
	downloadEVERootFSCmd.Flags().StringVarP(&eveHV, "eve-hv", "", defaults.DefaultEVEHV, "HV of EVE (kvm or xen)")
	downloadEVERootFSCmd.Flags().StringVarP(&outputDir, "downloader-dist", "d", "", "output directory")
}
