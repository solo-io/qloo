package version

import (
	"github.com/pkg/errors"
	version "github.com/solo-io/go-utils/versionutils"
	"github.com/solo-io/go-utils/log"
)

func CheckVersions() error {
	log.Printf("Checking expected solo kit and gloo versions...")
	tomlTree, err := version.ParseToml()
	if err != nil {
		return err
	}

	expectedGlooVersion, err := version.GetVersion(version.GlooPkg, tomlTree)
	if err != nil {
		return err
	}

	expectedSoloKitVersion, err := version.GetVersion(version.SoloKitPkg, tomlTree)
	if err != nil {
		return err
	}

	log.Printf("Checking repo versions...")
	actualGlooVersion, err := version.GetGitVersion("../gloo")
	if err != nil {
		return err
	}
	expectedTaggedGlooVersion := version.GetTag(expectedGlooVersion)
	if expectedTaggedGlooVersion != actualGlooVersion {
		return errors.Errorf("Expected gloo version %s, found gloo version %s in repo. Run 'make pin-repos' or fix manually.", expectedTaggedGlooVersion, actualGlooVersion)
	}

	actualSoloKitVersion, err := version.GetGitVersion("../solo-kit")
	if err != nil {
		return err
	}
	expectedTaggedSoloKitVersion := version.GetTag(expectedSoloKitVersion)
	if expectedTaggedSoloKitVersion != actualSoloKitVersion {
		return errors.Errorf("Expected solo kit version %s, found solo kit version %s in repo. Run 'make pin-repos' or fix manually.", expectedTaggedSoloKitVersion, actualSoloKitVersion)
	}
	log.Printf("Versions are pinned correctly.")
	return nil
}
