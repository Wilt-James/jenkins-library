// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type mavenStaticCodeChecksOptions struct {
	SpotBugs                  bool     `json:"spotBugs,omitempty"`
	Pmd                       bool     `json:"pmd,omitempty"`
	MavenModulesExcludes      []string `json:"mavenModulesExcludes,omitempty"`
	SpotBugsExcludeFilterFile string   `json:"spotBugsExcludeFilterFile,omitempty"`
	SpotBugsIncludeFilterFile string   `json:"spotBugsIncludeFilterFile,omitempty"`
	PmdExcludes               []string `json:"pmdExcludes,omitempty"`
	PmdRuleSets               []string `json:"pmdRuleSets,omitempty"`
}

// MavenStaticCodeChecksCommand Execute static code checks for Maven based projects. The plugins SpotBugs and PMD are used.
func MavenStaticCodeChecksCommand() *cobra.Command {
	metadata := mavenStaticCodeChecksMetadata()
	var stepConfig mavenStaticCodeChecksOptions
	var startTime time.Time

	var createMavenStaticCodeChecksCmd = &cobra.Command{
		Use:   "mavenStaticCodeChecks",
		Short: "Execute static code checks for Maven based projects. The plugins SpotBugs and PMD are used.",
		Long: `Executes Spotbugs Maven plugin as well as Pmd Maven plugin for static code checks.
SpotBugs is a program to find bugs in Java programs. It looks for instances of “bug patterns” — code instances that are likely to be errors.
For more information please visit https://spotbugs.readthedocs.io/en/latest/maven.html
PMD is a source code analyzer. It finds common programming flaws like unused variables, empty catch blocks, unnecessary object creation, and so forth. It supports Java, JavaScript, Salesforce.com Apex and Visualforce, PLSQL, Apache Velocity, XML, XSL.
For more information please visit https://pmd.github.io/`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("mavenStaticCodeChecks")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "mavenStaticCodeChecks", &stepConfig, config.OpenPiperFile)
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, "mavenStaticCodeChecks")
			mavenStaticCodeChecks(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addMavenStaticCodeChecksFlags(createMavenStaticCodeChecksCmd, &stepConfig)
	return createMavenStaticCodeChecksCmd
}

func addMavenStaticCodeChecksFlags(cmd *cobra.Command, stepConfig *mavenStaticCodeChecksOptions) {
	cmd.Flags().BoolVar(&stepConfig.SpotBugs, "spotBugs", true, "Parameter to turn off SpotBugs.")
	cmd.Flags().BoolVar(&stepConfig.Pmd, "pmd", true, "Parameter to turn off PMD.")
	cmd.Flags().StringSliceVar(&stepConfig.MavenModulesExcludes, "mavenModulesExcludes", []string{}, "Maven modules which should be excluded by the static code checks. By default the modules 'unit-tests' and 'integration-tests' will be excluded.")
	cmd.Flags().StringVar(&stepConfig.SpotBugsExcludeFilterFile, "spotBugsExcludeFilterFile", os.Getenv("PIPER_spotBugsExcludeFilterFile"), "Path to a filter file with bug definitions which should be excluded.")
	cmd.Flags().StringVar(&stepConfig.SpotBugsIncludeFilterFile, "spotBugsIncludeFilterFile", os.Getenv("PIPER_spotBugsIncludeFilterFile"), "Path to a filter file with bug definitions which should be included.")
	cmd.Flags().StringSliceVar(&stepConfig.PmdExcludes, "pmdExcludes", []string{}, "A comma-separated list of exclusions (.java source files) expressed as an Ant-style pattern relative to the sources root folder, i.e. application/src/main/java for maven projects.")
	cmd.Flags().StringSliceVar(&stepConfig.PmdRuleSets, "pmdRuleSets", []string{}, "The PMD rulesets to use. See the Stock Java Rulesets for a list of available rules. Defaults to a custom ruleset provided by this maven plugin.")

}

// retrieve step metadata
func mavenStaticCodeChecksMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "spotBugs",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pmd",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "mavenModulesExcludes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "spotBugsExcludeFilterFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "spotBugs/excludeFilterFile"}},
					},
					{
						Name:        "spotBugsIncludeFilterFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "spotBugs/includeFilterFile"}},
					},
					{
						Name:        "pmdExcludes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "pmd/excludes"}},
					},
					{
						Name:        "pmdRuleSets",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "pmd/ruleSets"}},
					},
				},
			},
		},
	}
	return theMetaData
}
