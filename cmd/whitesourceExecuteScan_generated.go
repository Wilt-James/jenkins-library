// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type whitesourceExecuteScanOptions struct {
	AgentDownloadURL                     string   `json:"agentDownloadUrl,omitempty"`
	AgentFileName                        string   `json:"agentFileName,omitempty"`
	AgentParameters                      []string `json:"agentParameters,omitempty"`
	AgentURL                             string   `json:"agentUrl,omitempty"`
	AggregateVersionWideReport           bool     `json:"aggregateVersionWideReport,omitempty"`
	BuildDescriptorExcludeList           []string `json:"buildDescriptorExcludeList,omitempty"`
	BuildDescriptorFile                  string   `json:"buildDescriptorFile,omitempty"`
	BuildTool                            string   `json:"buildTool,omitempty"`
	ConfigFilePath                       string   `json:"configFilePath,omitempty"`
	CreateProductFromPipeline            bool     `json:"createProductFromPipeline,omitempty"`
	CustomScanVersion                    string   `json:"customScanVersion,omitempty"`
	CvssSeverityLimit                    string   `json:"cvssSeverityLimit,omitempty"`
	EmailAddressesOfInitialProductAdmins []string `json:"emailAddressesOfInitialProductAdmins,omitempty"`
	Excludes                             []string `json:"excludes,omitempty"`
	Includes                             []string `json:"includes,omitempty"`
	InstallCommand                       string   `json:"installCommand,omitempty"`
	JreDownloadURL                       string   `json:"jreDownloadUrl,omitempty"`
	LicensingVulnerabilities             bool     `json:"licensingVulnerabilities,omitempty"`
	OrgToken                             string   `json:"orgToken,omitempty"`
	ProductName                          string   `json:"productName,omitempty"`
	ProductToken                         string   `json:"productToken,omitempty"`
	Version                              string   `json:"version,omitempty"`
	ProjectName                          string   `json:"projectName,omitempty"`
	ProjectToken                         string   `json:"projectToken,omitempty"`
	Reporting                            bool     `json:"reporting,omitempty"`
	ScanImage                            string   `json:"scanImage,omitempty"`
	ScanImageIncludeLayers               bool     `json:"scanImageIncludeLayers,omitempty"`
	ScanImageRegistryURL                 string   `json:"scanImageRegistryUrl,omitempty"`
	SecurityVulnerabilities              bool     `json:"securityVulnerabilities,omitempty"`
	ServiceURL                           string   `json:"serviceUrl,omitempty"`
	Timeout                              int      `json:"timeout,omitempty"`
	UserToken                            string   `json:"userToken,omitempty"`
	VersioningModel                      string   `json:"versioningModel,omitempty"`
	VulnerabilityReportFormat            string   `json:"vulnerabilityReportFormat,omitempty"`
	VulnerabilityReportTitle             string   `json:"vulnerabilityReportTitle,omitempty"`
	ProjectSettingsFile                  string   `json:"projectSettingsFile,omitempty"`
	GlobalSettingsFile                   string   `json:"globalSettingsFile,omitempty"`
	M2Path                               string   `json:"m2Path,omitempty"`
	InstallArtifacts                     bool     `json:"installArtifacts,omitempty"`
	DefaultNpmRegistry                   string   `json:"defaultNpmRegistry,omitempty"`
}

type whitesourceExecuteScanCommonPipelineEnvironment struct {
	custom struct {
		whitesourceProjectNames []string
	}
}

func (p *whitesourceExecuteScanCommonPipelineEnvironment) persist(path, resourceName string) {
	content := []struct {
		category string
		name     string
		value    interface{}
	}{
		{category: "custom", name: "whitesourceProjectNames", value: p.custom.whitesourceProjectNames},
	}

	errCount := 0
	for _, param := range content {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(param.category, param.name), param.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting piper environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Piper environment")
	}
}

type whitesourceExecuteScanInflux struct {
	step_data struct {
		fields struct {
			whitesource bool
		}
		tags struct {
		}
	}
	whitesource_data struct {
		fields struct {
			vulnerabilities       int
			major_vulnerabilities int
			minor_vulnerabilities int
			policy_violations     int
		}
		tags struct {
		}
	}
}

func (i *whitesourceExecuteScanInflux) persist(path, resourceName string) {
	measurementContent := []struct {
		measurement string
		valType     string
		name        string
		value       interface{}
	}{
		{valType: config.InfluxField, measurement: "step_data", name: "whitesource", value: i.step_data.fields.whitesource},
		{valType: config.InfluxField, measurement: "whitesource_data", name: "vulnerabilities", value: i.whitesource_data.fields.vulnerabilities},
		{valType: config.InfluxField, measurement: "whitesource_data", name: "major_vulnerabilities", value: i.whitesource_data.fields.major_vulnerabilities},
		{valType: config.InfluxField, measurement: "whitesource_data", name: "minor_vulnerabilities", value: i.whitesource_data.fields.minor_vulnerabilities},
		{valType: config.InfluxField, measurement: "whitesource_data", name: "policy_violations", value: i.whitesource_data.fields.policy_violations},
	}

	errCount := 0
	for _, metric := range measurementContent {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(metric.measurement, fmt.Sprintf("%vs", metric.valType), metric.name), metric.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting influx environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Influx environment")
	}
}

// WhitesourceExecuteScanCommand Execute a WhiteSource scan
func WhitesourceExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "whitesourceExecuteScan"

	metadata := whitesourceExecuteScanMetadata()
	var stepConfig whitesourceExecuteScanOptions
	var startTime time.Time
	var commonPipelineEnvironment whitesourceExecuteScanCommonPipelineEnvironment
	var influx whitesourceExecuteScanInflux

	var createWhitesourceExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Execute a WhiteSource scan",
		Long: `With this step [WhiteSource](https://www.whitesourcesoftware.com) security and license compliance scans can be executed and assessed.
WhiteSource is a Software as a Service offering based on a so called unified agent that locally determines the dependency
tree of a node.js, Java, Python, Ruby, or Scala based solution and sends it to the WhiteSource server for a policy based license compliance
check and additional Free and Open Source Software Publicly Known Vulnerabilities detection.

The step uses the so-called WhiteSource Unified Agent. For details please refer to the [WhiteSource Unified Agent Documentation](https://whitesource.atlassian.net/wiki/spaces/WD/pages/33718339/Unified+Agent).

!!! note "Docker Images"
    The underlying Docker images are public and specific to the solution's programming language(s) and therefore may have to be exchanged
    to fit to and support the relevant scenario. The default Python environment used is i.e. Python 3 based.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.OrgToken)
			log.RegisterSecret(stepConfig.UserToken)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				commonPipelineEnvironment.persist(GeneralConfig.EnvRootPath, "commonPipelineEnvironment")
				influx.persist(GeneralConfig.EnvRootPath, "influx")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			whitesourceExecuteScan(stepConfig, &telemetryData, &commonPipelineEnvironment, &influx)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addWhitesourceExecuteScanFlags(createWhitesourceExecuteScanCmd, &stepConfig)
	return createWhitesourceExecuteScanCmd
}

func addWhitesourceExecuteScanFlags(cmd *cobra.Command, stepConfig *whitesourceExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.AgentDownloadURL, "agentDownloadUrl", `https://github.com/whitesource/unified-agent-distribution/releases/latest/download/wss-unified-agent.jar`, "URL used to download the latest version of the WhiteSource Unified Agent.")
	cmd.Flags().StringVar(&stepConfig.AgentFileName, "agentFileName", `wss-unified-agent.jar`, "Locally used name for the Unified Agent jar file after download.")
	cmd.Flags().StringSliceVar(&stepConfig.AgentParameters, "agentParameters", []string{}, "[NOT IMPLEMENTED] List of additional parameters passed to the Unified Agent command line.")
	cmd.Flags().StringVar(&stepConfig.AgentURL, "agentUrl", `https://saas.whitesourcesoftware.com/agent`, "URL to the WhiteSource agent endpoint.")
	cmd.Flags().BoolVar(&stepConfig.AggregateVersionWideReport, "aggregateVersionWideReport", false, "This does not run a scan, instead just generated a report for all projects with projectVersion = config.ProductVersion")
	cmd.Flags().StringSliceVar(&stepConfig.BuildDescriptorExcludeList, "buildDescriptorExcludeList", []string{`unit-tests/pom.xml`, `integration-tests/pom.xml`}, "List of build descriptors and therefore modules to exclude from the scan and assessment activities.")
	cmd.Flags().StringVar(&stepConfig.BuildDescriptorFile, "buildDescriptorFile", os.Getenv("PIPER_buildDescriptorFile"), "Explicit path to the build descriptor file.")
	cmd.Flags().StringVar(&stepConfig.BuildTool, "buildTool", os.Getenv("PIPER_buildTool"), "Defines the tool which is used for building the artifact.")
	cmd.Flags().StringVar(&stepConfig.ConfigFilePath, "configFilePath", `./wss-unified-agent.config`, "Explicit path to the WhiteSource Unified Agent configuration file.")
	cmd.Flags().BoolVar(&stepConfig.CreateProductFromPipeline, "createProductFromPipeline", true, "Whether to create the related WhiteSource product on the fly based on the supplied pipeline configuration.")
	cmd.Flags().StringVar(&stepConfig.CustomScanVersion, "customScanVersion", os.Getenv("PIPER_customScanVersion"), "Custom version of the WhiteSource project used as source.")
	cmd.Flags().StringVar(&stepConfig.CvssSeverityLimit, "cvssSeverityLimit", `-1`, "Limit of tolerable CVSS v3 score upon assessment and in consequence fails the build, defaults to  `-1`.")
	cmd.Flags().StringSliceVar(&stepConfig.EmailAddressesOfInitialProductAdmins, "emailAddressesOfInitialProductAdmins", []string{}, "The list of email addresses to assign as product admins for newly created WhiteSource products.")
	cmd.Flags().StringSliceVar(&stepConfig.Excludes, "excludes", []string{}, "List of file path patterns to exclude in the scan.")
	cmd.Flags().StringSliceVar(&stepConfig.Includes, "includes", []string{}, "List of file path patterns to include in the scan.")
	cmd.Flags().StringVar(&stepConfig.InstallCommand, "installCommand", os.Getenv("PIPER_installCommand"), "[NOT IMPLEMENTED] Install command that can be used to populate the default docker image for some scenarios.")
	cmd.Flags().StringVar(&stepConfig.JreDownloadURL, "jreDownloadUrl", `https://github.com/SAP/SapMachine/releases/download/sapmachine-11.0.2/sapmachine-jre-11.0.2_linux-x64_bin.tar.gz`, "URL used for downloading the Java Runtime Environment (JRE) required to run the WhiteSource Unified Agent.")
	cmd.Flags().BoolVar(&stepConfig.LicensingVulnerabilities, "licensingVulnerabilities", true, "[NOT IMPLEMENTED] Whether license compliance is considered and reported as part of the assessment.")
	cmd.Flags().StringVar(&stepConfig.OrgToken, "orgToken", os.Getenv("PIPER_orgToken"), "WhiteSource token identifying your organization.")
	cmd.Flags().StringVar(&stepConfig.ProductName, "productName", os.Getenv("PIPER_productName"), "Name of the WhiteSource product used for results aggregation. This parameter is mandatory if the parameter `createProductFromPipeline` is set to `true` and the WhiteSource product does not yet exist. It is also mandatory if the parameter `productToken` is not provided.")
	cmd.Flags().StringVar(&stepConfig.ProductToken, "productToken", os.Getenv("PIPER_productToken"), "Token of the WhiteSource product to be created and used for results aggregation, usually determined automatically. Can optionally be provided as an alternative to `productName`.")
	cmd.Flags().StringVar(&stepConfig.Version, "version", os.Getenv("PIPER_version"), "Version of the WhiteSource product to be created and used for results aggregation.")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", os.Getenv("PIPER_projectName"), "The project name used for reporting results in WhiteSource. When provided, all source modules will be scanned into one aggregated WhiteSource project. For scan types `maven`, `mta`, `npm`, the default is to generate one WhiteSource project per module, whereas the project name is derived from the module's build descriptor. For NPM modules, project aggregation is not supported, the last scanned NPM module will override all previously aggregated scan results!")
	cmd.Flags().StringVar(&stepConfig.ProjectToken, "projectToken", os.Getenv("PIPER_projectToken"), "Project token to execute scan on. Ignored for scan types `maven`, `mta` and `npm`. Used for project aggregation when scanning with the Unified Agent and can be provided as an alternative to `projectName`.")
	cmd.Flags().BoolVar(&stepConfig.Reporting, "reporting", true, "Whether assessment is being done at all, defaults to `true`")
	cmd.Flags().StringVar(&stepConfig.ScanImage, "scanImage", os.Getenv("PIPER_scanImage"), "For `buildTool: docker`: Defines the docker image which should be scanned.")
	cmd.Flags().BoolVar(&stepConfig.ScanImageIncludeLayers, "scanImageIncludeLayers", true, "For `buildTool: docker`: Defines if layers should be included.")
	cmd.Flags().StringVar(&stepConfig.ScanImageRegistryURL, "scanImageRegistryUrl", os.Getenv("PIPER_scanImageRegistryUrl"), "For `buildTool: docker`: Defines the registry where the scanImage is located.")
	cmd.Flags().BoolVar(&stepConfig.SecurityVulnerabilities, "securityVulnerabilities", true, "Whether security compliance is considered and reported as part of the assessment.")
	cmd.Flags().StringVar(&stepConfig.ServiceURL, "serviceUrl", `https://saas.whitesourcesoftware.com/api`, "URL to the WhiteSource API endpoint.")
	cmd.Flags().IntVar(&stepConfig.Timeout, "timeout", 900, "Timeout in seconds until an HTTP call is forcefully terminated.")
	cmd.Flags().StringVar(&stepConfig.UserToken, "userToken", os.Getenv("PIPER_userToken"), "User token to access WhiteSource. In Jenkins use case this is automatically filled through the credentials.")
	cmd.Flags().StringVar(&stepConfig.VersioningModel, "versioningModel", `major`, "The default project versioning model used in case `projectVersion` parameter is empty for creating the version based on the build descriptor version to report results in Whitesource, can be one of `'major'`, `'major-minor'`, `'semantic'`, `'full'`")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityReportFormat, "vulnerabilityReportFormat", `xlsx`, "Format of the file the vulnerability report is written to.")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityReportTitle, "vulnerabilityReportTitle", `WhiteSource Security Vulnerability Report`, "Title of vulnerability report written during the assessment phase.")
	cmd.Flags().StringVar(&stepConfig.ProjectSettingsFile, "projectSettingsFile", os.Getenv("PIPER_projectSettingsFile"), "Path to the mvn settings file that should be used as project settings file.")
	cmd.Flags().StringVar(&stepConfig.GlobalSettingsFile, "globalSettingsFile", os.Getenv("PIPER_globalSettingsFile"), "Path to the mvn settings file that should be used as global settings file.")
	cmd.Flags().StringVar(&stepConfig.M2Path, "m2Path", os.Getenv("PIPER_m2Path"), "Path to the location of the local repository that should be used.")
	cmd.Flags().BoolVar(&stepConfig.InstallArtifacts, "installArtifacts", false, "If enabled, it will install all artifacts to the local maven repository to make them available before running whitesource. This is required if any maven module has dependencies to other modules in the repository and they were not installed before.")
	cmd.Flags().StringVar(&stepConfig.DefaultNpmRegistry, "defaultNpmRegistry", os.Getenv("PIPER_defaultNpmRegistry"), "URL of the npm registry to use. Defaults to https://registry.npmjs.org/")

	cmd.MarkFlagRequired("buildTool")
	cmd.MarkFlagRequired("orgToken")
	cmd.MarkFlagRequired("userToken")
}

// retrieve step metadata
func whitesourceExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "whitesourceExecuteScan",
			Aliases:     []config.Alias{},
			Description: "Execute a WhiteSource scan",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Resources: []config.StepResources{
					{Name: "buildDescriptor", Type: "stash"},
					{Name: "opensourceConfiguration", Type: "stash"},
					{Name: "checkmarx", Type: "stash"},
				},
				Parameters: []config.StepParameters{
					{
						Name:        "agentDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentParameters",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "aggregateVersionWideReport",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorExcludeList",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "buildTool",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "buildTool",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "configFilePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "createProductFromPipeline",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "customScanVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "cvssSeverityLimit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "emailAddressesOfInitialProductAdmins",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "excludes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "includes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "installCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "jreDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesource/jreDownloadUrl"}},
					},
					{
						Name:        "licensingVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "orgToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "orgAdminUserTokenCredentialsId",
								Type: "secret",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "whitesourceOrgToken"}, {Name: "whitesource/orgToken"}},
					},
					{
						Name:        "productName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceProductName"}, {Name: "whitesource/productName"}},
					},
					{
						Name:        "productToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceProductToken"}, {Name: "whitesource/productToken"}},
					},
					{
						Name: "version",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "artifactVersion",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "productVersion"}, {Name: "whitesourceProductVersion"}, {Name: "whitesource/productVersion"}},
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceProjectName"}},
					},
					{
						Name:        "projectToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "reporting",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "scanImage",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "scanImageIncludeLayers",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "scanImageRegistryUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "securityVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "serviceUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceServiceUrl"}, {Name: "whitesource/serviceUrl"}},
					},
					{
						Name:        "timeout",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "userToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "userTokenCredentialsId",
								Type: "secret",
							},

							{
								Name:  "",
								Paths: []string{"$(vaultPath)/whitesource", "$(vaultBasePath)/$(vaultPipelineName)/whitesource", "$(vaultBasePath)/GROUP-SECRETS/whitesource"},
								Type:  "vaultSecret",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "versioningModel",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "defaultVersioningModel"}},
					},
					{
						Name:        "vulnerabilityReportFormat",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityReportTitle",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/projectSettingsFile"}},
					},
					{
						Name:        "globalSettingsFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/globalSettingsFile"}},
					},
					{
						Name:        "m2Path",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "maven/m2Path"}},
					},
					{
						Name:        "installArtifacts",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "STEPS", "STAGES", "PARAMETERS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "defaultNpmRegistry",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "GENERAL", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "npm/defaultNpmRegistry"}},
					},
				},
			},
			Containers: []config.Container{
				{Image: "buildpack-deps:stretch-curl", WorkingDir: "/tmp", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "dub"}, {Name: "buildTool", Value: "docker"}}}}},
				{Image: "devxci/mbtci:1.0.14", WorkingDir: "/home/mta", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "mta"}}}}},
				{Image: "golang:1", WorkingDir: "/go", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "golang"}}}}},
				{Image: "gradle", WorkingDir: "/home/gradle/", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "gradle"}}}}},
				{Image: "hseeberger/scala-sbt:8u181_2.12.8_1.2.8", WorkingDir: "/tmp", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "sbt"}}}}},
				{Image: "maven:3.5-jdk-8", WorkingDir: "/tmp", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "maven"}}}}},
				{Image: "node:lts-stretch", WorkingDir: "/home/node", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "npm"}}}}},
				{Image: "python:3.6-stretch", WorkingDir: "/tmp", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "pip"}}}}},
				{Image: "ruby", WorkingDir: "/tmp", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "ruby"}}}}},
				{Image: "rust", WorkingDir: "/tmp", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "rust"}}}}},
				{Image: "node:lts-stretch", WorkingDir: "/home/node", Conditions: []config.Condition{{ConditionRef: "strings-equal", Params: []config.Param{{Name: "buildTool", Value: "yarn"}}}}},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"Name": "custom/whitesourceProjectNames"},
						},
					},
					{
						Name: "influx",
						Type: "influx",
						Parameters: []map[string]interface{}{
							{"Name": "step_data"}, {"fields": []map[string]string{{"name": "whitesource"}}},
							{"Name": "whitesource_data"}, {"fields": []map[string]string{{"name": "vulnerabilities"}, {"name": "major_vulnerabilities"}, {"name": "minor_vulnerabilities"}, {"name": "policy_violations"}}},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
