// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/ans"
	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/SAP/jenkins-library/pkg/validation"
	"github.com/spf13/cobra"
)

type gctsExecuteABAPQualityChecksOptions struct {
	Username             string `json:"username,omitempty"`
	Password             string `json:"password,omitempty"`
	Host                 string `json:"host,omitempty"`
	Repository           string `json:"repository,omitempty"`
	Client               string `json:"client,omitempty"`
	AUnitTest            bool   `json:"aUnitTest,omitempty"`
	AtcCheck             bool   `json:"atcCheck,omitempty"`
	AtcVariant           string `json:"atcVariant,omitempty"`
	Scope                string `json:"scope,omitempty"`
	Commit               string `json:"commit,omitempty"`
	Workspace            string `json:"workspace,omitempty"`
	AtcResultsFileName   string `json:"atcResultsFileName,omitempty"`
	AUnitResultsFileName string `json:"aUnitResultsFileName,omitempty"`
}

// GctsExecuteABAPQualityChecksCommand Runs ABAP unit tests and ATC (ABAP Test Cockpit) checks for a specified object scope.
func GctsExecuteABAPQualityChecksCommand() *cobra.Command {
	const STEP_NAME = "gctsExecuteABAPQualityChecks"

	metadata := gctsExecuteABAPQualityChecksMetadata()
	var stepConfig gctsExecuteABAPQualityChecksOptions
	var startTime time.Time
	var logCollector *log.CollectorHook
	var splunkClient *splunk.Splunk
	telemetryClient := &telemetry.Telemetry{}

	var createGctsExecuteABAPQualityChecksCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Runs ABAP unit tests and ATC (ABAP Test Cockpit) checks for a specified object scope.",
		Long: `This step executes ABAP unit test and ATC checks for a specified scope of objects that exist in a local Git repository on an ABAP system.
Depending on your use case, you can specify a scope of objects for which you want to execute the checks. In addition, you can choose whether you want to execute only ABAP units tests, or only ATC checks, or both.
By default, both checks are executed.
The results of the checks are stored in a [Checkstyle](https://checkstyle.sourceforge.io/) format. With the help of the Jenkins [Warnings-Next-Generation](https://plugins.jenkins.io/warnings-ng/) Plugin), you can view the issues found, and navigate to the exact line of the source code where the issue occurred.
To make the findings visible in Jenkins interface, you will need to use step recordIssues. An example will be shown in the Example section.
<br />
You can use this step as of SAP S/4HANA 2020.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			GeneralConfig.GitHubAccessTokens = ResolveAccessTokens(GeneralConfig.GitHubTokens)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.Username)
			log.RegisterSecret(stepConfig.Password)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient = &splunk.Splunk{}
				logCollector = &log.CollectorHook{CorrelationID: GeneralConfig.CorrelationID}
				log.RegisterHook(logCollector)
			}

			validation, err := validation.New(validation.WithJSONNamesForStructFields(), validation.WithPredefinedErrorMessages())
			if err != nil {
				return err
			}
			if err = validation.ValidateStruct(stepConfig); err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			stepTelemetryData := telemetry.CustomData{}
			stepTelemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				stepTelemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				stepTelemetryData.ErrorCategory = log.GetErrorCategory().String()
				stepTelemetryData.PiperCommitHash = GitCommit
				telemetryClient.SetData(&stepTelemetryData)
				telemetryClient.Send()
				if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
					splunkClient.Send(telemetryClient.GetData(), logCollector)
				}
				if len(GeneralConfig.ANSServiceKey) > 0 {
					ans.Send(GeneralConfig.ANSServiceKey)
				}
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetryClient.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient.Initialize(GeneralConfig.CorrelationID,
					GeneralConfig.HookConfig.SplunkConfig.Dsn,
					GeneralConfig.HookConfig.SplunkConfig.Token,
					GeneralConfig.HookConfig.SplunkConfig.Index,
					GeneralConfig.HookConfig.SplunkConfig.SendLogs)
			}
			gctsExecuteABAPQualityChecks(stepConfig, &stepTelemetryData)
			stepTelemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addGctsExecuteABAPQualityChecksFlags(createGctsExecuteABAPQualityChecksCmd, &stepConfig)
	return createGctsExecuteABAPQualityChecksCmd
}

func addGctsExecuteABAPQualityChecksFlags(cmd *cobra.Command, stepConfig *gctsExecuteABAPQualityChecksOptions) {
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User that authenticates to the ABAP system. Note – Don´t provide this parameter directly. Either set it in the environment, or in the Jenkins credentials store, and provide the ID as value of the abapCredentialsId parameter.")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "Password of the ABAP  user that authenticates to the ABAP system. Note – Don´t provide this parameter directly. Either set it in the environment, or in the Jenkins credentials store, and provide the ID as value of the abapCredentialsId parameter.")
	cmd.Flags().StringVar(&stepConfig.Host, "host", os.Getenv("PIPER_host"), "Protocol and host of the ABAP system, including the port. Please provide in the format <protocol>://<host>:<port>. Supported protocols are http and https.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Name (ID) of the local repository on the ABAP system")
	cmd.Flags().StringVar(&stepConfig.Client, "client", os.Getenv("PIPER_client"), "Client of the ABAP system in which you want to execute the checks")
	cmd.Flags().BoolVar(&stepConfig.AUnitTest, "aUnitTest", true, "Indication whether you want to execute the unit test checks.")
	cmd.Flags().BoolVar(&stepConfig.AtcCheck, "atcCheck", true, "Indication whether you want to execute the ATC checks.")
	cmd.Flags().StringVar(&stepConfig.AtcVariant, "atcVariant", `DEFAULT_REMOTE_REF`, "Variant for ATC checks")
	cmd.Flags().StringVar(&stepConfig.Scope, "scope", `repository`, "Scope of objects for which you want to execute the checks:\n\n * localChangedObjects - object delta derived from last activity in the local repository. The checks are executed for the individual objects.\n * remoteChangedObjects - object delta between the commit that triggered the pipeline and the current commit in the remote repository. The checks are executed for the individual objects.\n * localChangedPackages - object delta derived from last activity in the local repository. All objects are resolved into packages. The checks are executed for the packages.\n * remoteChangedPackages - object delta between the commit that triggered the pipeline and the current commit in the remote repository. All objects are resolved into packages. The checks are executed for the packages.\n * repository - all objects that are part of the local repository. The checks are executed for the individual objects. Packages (DEVC) are excluded. This is the default scope.\n * packages - all packages that are part of the local repository . The checks are executed for the packages.\n")
	cmd.Flags().StringVar(&stepConfig.Commit, "commit", os.Getenv("PIPER_commit"), "ID of the commit that triggered the pipeline or any other commit to compare objects. For scopes remoteChangedObjects and remoteChangedPackages secifying a commit it's mandatory.")
	cmd.Flags().StringVar(&stepConfig.Workspace, "workspace", os.Getenv("PIPER_workspace"), "Absolute path to directory which contains the source code that your CI/CD tool checks out. For example in Jenkins, the workspace parameter is /var/jenkins_home/workspace/<jobName>/")
	cmd.Flags().StringVar(&stepConfig.AtcResultsFileName, "atcResultsFileName", `ATCResults.xml`, "Specifies output file name for the results from the ATC checks")
	cmd.Flags().StringVar(&stepConfig.AUnitResultsFileName, "aUnitResultsFileName", `AUnitResults.xml`, "Specifies output file name for the results from the AUnit tests")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("repository")
	cmd.MarkFlagRequired("client")
	cmd.MarkFlagRequired("workspace")
}

// retrieve step metadata
func gctsExecuteABAPQualityChecksMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "gctsExecuteABAPQualityChecks",
			Aliases:     []config.Alias{},
			Description: "Runs ABAP unit tests and ATC (ABAP Test Cockpit) checks for a specified object scope.",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "abapCredentialsId", Description: "ID taken from the Jenkins credentials store containing user name and password of the user that authenticates to the ABAP system on which you want to execute the checks.", Type: "jenkins"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "username",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "abapCredentialsId",
								Param: "username",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_username"),
					},
					{
						Name: "password",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "abapCredentialsId",
								Param: "password",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_password"),
					},
					{
						Name:        "host",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_host"),
					},
					{
						Name:        "repository",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_repository"),
					},
					{
						Name:        "client",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_client"),
					},
					{
						Name:        "aUnitTest",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "atcCheck",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "atcVariant",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `DEFAULT_REMOTE_REF`,
					},
					{
						Name:        "scope",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `repository`,
					},
					{
						Name:        "commit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_commit"),
					},
					{
						Name:        "workspace",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_workspace"),
					},
					{
						Name:        "atcResultsFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `ATCResults.xml`,
					},
					{
						Name:        "aUnitResultsFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `AUnitResults.xml`,
					},
				},
			},
		},
	}
	return theMetaData
}
