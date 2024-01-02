//go:build integration

/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schematicsv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the schematicsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SchematicsV1 Integration Tests`, func() {
	const externalConfigFile = "../schematics_v1.env"

	var (
		err          error
		schematicsService *schematicsv1.SchematicsV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(schematicsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

			schematicsService, err = schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)
			Expect(err).To(BeNil())
			Expect(schematicsService).ToNot(BeNil())
			Expect(schematicsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			schematicsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSchematicsVersion - Get Schematics API information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchematicsVersion(getSchematicsVersionOptions *GetSchematicsVersionOptions)`, func() {
			getSchematicsVersionOptions := &schematicsv1.GetSchematicsVersionOptions{
			}

			versionResponse, response, err := schematicsService.GetSchematicsVersion(getSchematicsVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionResponse).ToNot(BeNil())
		})
	})

	Describe(`ListLocations - List supported locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLocations(listLocationsOptions *ListLocationsOptions)`, func() {
			listLocationsOptions := &schematicsv1.ListLocationsOptions{
			}

			schematicsLocationsList, response, err := schematicsService.ListLocations(listLocationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsLocationsList).ToNot(BeNil())
		})
	})

	Describe(`ListResourceGroup - List resource groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceGroup(listResourceGroupOptions *ListResourceGroupOptions)`, func() {
			listResourceGroupOptions := &schematicsv1.ListResourceGroupOptions{
			}

			resourceGroupResponse, response, err := schematicsService.ListResourceGroup(listResourceGroupOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupResponse).ToNot(BeNil())
		})
	})

	Describe(`ListSchematicsLocation - List supported schematics locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSchematicsLocation(listSchematicsLocationOptions *ListSchematicsLocationOptions)`, func() {
			listSchematicsLocationOptions := &schematicsv1.ListSchematicsLocationOptions{
			}

			schematicsLocations, response, err := schematicsService.ListSchematicsLocation(listSchematicsLocationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsLocations).ToNot(BeNil())
		})
	})

	Describe(`ProcessTemplateMetaData - Get variable metadata by parsing the template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ProcessTemplateMetaData(processTemplateMetaDataOptions *ProcessTemplateMetaDataOptions)`, func() {
			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			processTemplateMetaDataOptions := &schematicsv1.ProcessTemplateMetaDataOptions{
				TemplateType: core.StringPtr("testString"),
				Source: externalSourceModel,
				Region: core.StringPtr("testString"),
				SourceType: core.StringPtr("local"),
				XGithubToken: core.StringPtr("testString"),
			}

			templateMetaDataResponse, response, err := schematicsService.ProcessTemplateMetaData(processTemplateMetaDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateMetaDataResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateWorkspace - Create a workspace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)`, func() {
			serviceExtensionsModel := &schematicsv1.ServiceExtensions{
				Name: core.StringPtr("flavor"),
				Value: core.StringPtr("testString"),
				Type: core.StringPtr("string"),
			}

			catalogRefModel := &schematicsv1.CatalogRef{
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				ServiceExtensions: []schematicsv1.ServiceExtensions{*serviceExtensionsModel},
			}

			dependenciesModel := &schematicsv1.Dependencies{
				Parents: []string{"testString"},
				Children: []string{"testString"},
			}

			sharedTargetDataModel := &schematicsv1.SharedTargetData{
				ClusterCreatedOn: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				ClusterName: core.StringPtr("testString"),
				ClusterType: core.StringPtr("testString"),
				EntitlementKeys: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Namespace: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				WorkerCount: core.Int64Ptr(int64(26)),
				WorkerMachineType: core.StringPtr("testString"),
			}

			environmentValuesMetadataModel := &schematicsv1.EnvironmentValuesMetadata{
				Hidden: core.BoolPtr(true),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
			}

			injectTerraformTemplateItemTftParametersItemModel := &schematicsv1.InjectTerraformTemplateItemTftParametersItem{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			injectTerraformTemplateItemModel := &schematicsv1.InjectTerraformTemplateItem{
				TftGitURL: core.StringPtr("testString"),
				TftGitToken: core.StringPtr("testString"),
				TftPrefix: core.StringPtr("testString"),
				InjectionType: core.StringPtr("testString"),
				TftName: core.StringPtr("testString"),
				TftParameters: []schematicsv1.InjectTerraformTemplateItemTftParametersItem{*injectTerraformTemplateItemTftParametersItemModel},
			}

			workspaceVariableRequestModel := &schematicsv1.WorkspaceVariableRequest{
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
				Type: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Value: core.StringPtr("testString"),
			}

			templateSourceDataRequestModel := &schematicsv1.TemplateSourceDataRequest{
				EnvValues: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				EnvValuesMetadata: []schematicsv1.EnvironmentValuesMetadata{*environmentValuesMetadataModel},
				Folder: core.StringPtr("testString"),
				Compact: core.BoolPtr(true),
				InitStateFile: core.StringPtr("testString"),
				Injectors: []schematicsv1.InjectTerraformTemplateItem{*injectTerraformTemplateItemModel},
				Type: core.StringPtr("testString"),
				UninstallScriptName: core.StringPtr("testString"),
				Values: core.StringPtr("testString"),
				ValuesMetadata: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Variablestore: []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel},
			}

			templateRepoRequestModel := &schematicsv1.TemplateRepoRequest{
				Branch: core.StringPtr("testString"),
				Release: core.StringPtr("testString"),
				RepoShaValue: core.StringPtr("testString"),
				RepoURL: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			workspaceStatusRequestModel := &schematicsv1.WorkspaceStatusRequest{
				Frozen: core.BoolPtr(true),
				FrozenAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				FrozenBy: core.StringPtr("testString"),
				Locked: core.BoolPtr(true),
				LockedBy: core.StringPtr("testString"),
				LockedTime: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			createWorkspaceOptions := &schematicsv1.CreateWorkspaceOptions{
				AppliedShareddataIds: []string{"testString"},
				CatalogRef: catalogRefModel,
				Dependencies: dependenciesModel,
				Description: core.StringPtr("testString"),
				Location: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ResourceGroup: core.StringPtr("testString"),
				SharedData: sharedTargetDataModel,
				Tags: []string{"testString"},
				TemplateData: []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel},
				TemplateRef: core.StringPtr("testString"),
				TemplateRepo: templateRepoRequestModel,
				Type: []string{"testString"},
				WorkspaceStatus: workspaceStatusRequestModel,
				AgentID: core.StringPtr("testString"),
				XGithubToken: core.StringPtr("testString"),
			}

			workspaceResponse, response, err := schematicsService.CreateWorkspace(createWorkspaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(workspaceResponse).ToNot(BeNil())
		})
	})

	Describe(`GetAllWorkspaceInputs - Get workspace template details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllWorkspaceInputs(getAllWorkspaceInputsOptions *GetAllWorkspaceInputsOptions)`, func() {
			getAllWorkspaceInputsOptions := &schematicsv1.GetAllWorkspaceInputsOptions{
				WID: core.StringPtr("testString"),
			}

			workspaceTemplateValuesResponse, response, err := schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceTemplateValuesResponse).ToNot(BeNil())
		})
	})

	Describe(`GetTemplateActivityLog - Show logs for a workspace job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTemplateActivityLog(getTemplateActivityLogOptions *GetTemplateActivityLogOptions)`, func() {
			getTemplateActivityLogOptions := &schematicsv1.GetTemplateActivityLogOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
				ActivityID: core.StringPtr("testString"),
				LogTfCmd: core.BoolPtr(true),
				LogTfPrefix: core.BoolPtr(true),
				LogTfNullResource: core.BoolPtr(true),
				LogTfAnsible: core.BoolPtr(true),
			}

			result, response, err := schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetTemplateLogs - Show latest logs for a workspace template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTemplateLogs(getTemplateLogsOptions *GetTemplateLogsOptions)`, func() {
			getTemplateLogsOptions := &schematicsv1.GetTemplateLogsOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
				LogTfCmd: core.BoolPtr(true),
				LogTfPrefix: core.BoolPtr(true),
				LogTfNullResource: core.BoolPtr(true),
				LogTfAnsible: core.BoolPtr(true),
			}

			result, response, err := schematicsService.GetTemplateLogs(getTemplateLogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspace - Get workspace details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)`, func() {
			getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{
				WID: core.StringPtr("testString"),
			}

			workspaceResponse, response, err := schematicsService.GetWorkspace(getWorkspaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceActivityLogs - Get workspace job log URL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions *GetWorkspaceActivityLogsOptions)`, func() {
			getWorkspaceActivityLogsOptions := &schematicsv1.GetWorkspaceActivityLogsOptions{
				WID: core.StringPtr("testString"),
				ActivityID: core.StringPtr("testString"),
			}

			workspaceActivityLogs, response, err := schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivityLogs).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceInputMetadata - List workspace variable metadata`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptions *GetWorkspaceInputMetadataOptions)`, func() {
			getWorkspaceInputMetadataOptions := &schematicsv1.GetWorkspaceInputMetadataOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
			}

			result, response, err := schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceInputs - List workspace input variables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceInputs(getWorkspaceInputsOptions *GetWorkspaceInputsOptions)`, func() {
			getWorkspaceInputsOptions := &schematicsv1.GetWorkspaceInputsOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
			}

			templateValues, response, err := schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateValues).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceLogUrls - Get latest workspace job log URL for all workspace templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions *GetWorkspaceLogUrlsOptions)`, func() {
			getWorkspaceLogUrlsOptions := &schematicsv1.GetWorkspaceLogUrlsOptions{
				WID: core.StringPtr("testString"),
			}

			logStoreResponseList, response, err := schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logStoreResponseList).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceOutputs - List workspace output values`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceOutputs(getWorkspaceOutputsOptions *GetWorkspaceOutputsOptions)`, func() {
			getWorkspaceOutputsOptions := &schematicsv1.GetWorkspaceOutputsOptions{
				WID: core.StringPtr("testString"),
			}

			outputValuesItem, response, err := schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outputValuesItem).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceReadme - Show workspace template readme`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceReadme(getWorkspaceReadmeOptions *GetWorkspaceReadmeOptions)`, func() {
			getWorkspaceReadmeOptions := &schematicsv1.GetWorkspaceReadmeOptions{
				WID: core.StringPtr("testString"),
				Ref: core.StringPtr("testString"),
				Formatted: core.StringPtr("markdown"),
			}

			templateReadme, response, err := schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateReadme).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceResources - List workspace resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceResources(getWorkspaceResourcesOptions *GetWorkspaceResourcesOptions)`, func() {
			getWorkspaceResourcesOptions := &schematicsv1.GetWorkspaceResourcesOptions{
				WID: core.StringPtr("testString"),
			}

			templateResources, response, err := schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResources).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceState - Get Terraform statefile URL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceState(getWorkspaceStateOptions *GetWorkspaceStateOptions)`, func() {
			getWorkspaceStateOptions := &schematicsv1.GetWorkspaceStateOptions{
				WID: core.StringPtr("testString"),
			}

			stateStoreResponseList, response, err := schematicsService.GetWorkspaceState(getWorkspaceStateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stateStoreResponseList).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceTemplateState - Show Terraform statefile content`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions *GetWorkspaceTemplateStateOptions)`, func() {
			getWorkspaceTemplateStateOptions := &schematicsv1.GetWorkspaceTemplateStateOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
			}

			templateStateStore, response, err := schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateStateStore).ToNot(BeNil())
		})
	})

	Describe(`ListWorkspaces - List workspaces`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)`, func() {
			listWorkspacesOptions := &schematicsv1.ListWorkspacesOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Profile: core.StringPtr("ids"),
				ResourceGroup: core.StringPtr("testString"),
			}

			workspaceResponseList, response, err := schematicsService.ListWorkspaces(listWorkspacesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponseList).ToNot(BeNil())
		})
	})

	Describe(`ReplaceWorkspace - Update workspace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceWorkspace(replaceWorkspaceOptions *ReplaceWorkspaceOptions)`, func() {
			serviceExtensionsModel := &schematicsv1.ServiceExtensions{
				Name: core.StringPtr("flavor"),
				Value: core.StringPtr("testString"),
				Type: core.StringPtr("string"),
			}

			catalogRefModel := &schematicsv1.CatalogRef{
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				ServiceExtensions: []schematicsv1.ServiceExtensions{*serviceExtensionsModel},
			}

			dependenciesModel := &schematicsv1.Dependencies{
				Parents: []string{"testString"},
				Children: []string{"testString"},
			}

			sharedTargetDataModel := &schematicsv1.SharedTargetData{
				ClusterCreatedOn: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				ClusterName: core.StringPtr("testString"),
				ClusterType: core.StringPtr("testString"),
				EntitlementKeys: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Namespace: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				WorkerCount: core.Int64Ptr(int64(26)),
				WorkerMachineType: core.StringPtr("testString"),
			}

			environmentValuesMetadataModel := &schematicsv1.EnvironmentValuesMetadata{
				Hidden: core.BoolPtr(true),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
			}

			injectTerraformTemplateItemTftParametersItemModel := &schematicsv1.InjectTerraformTemplateItemTftParametersItem{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			injectTerraformTemplateItemModel := &schematicsv1.InjectTerraformTemplateItem{
				TftGitURL: core.StringPtr("testString"),
				TftGitToken: core.StringPtr("testString"),
				TftPrefix: core.StringPtr("testString"),
				InjectionType: core.StringPtr("testString"),
				TftName: core.StringPtr("testString"),
				TftParameters: []schematicsv1.InjectTerraformTemplateItemTftParametersItem{*injectTerraformTemplateItemTftParametersItemModel},
			}

			workspaceVariableRequestModel := &schematicsv1.WorkspaceVariableRequest{
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
				Type: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Value: core.StringPtr("testString"),
			}

			templateSourceDataRequestModel := &schematicsv1.TemplateSourceDataRequest{
				EnvValues: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				EnvValuesMetadata: []schematicsv1.EnvironmentValuesMetadata{*environmentValuesMetadataModel},
				Folder: core.StringPtr("testString"),
				Compact: core.BoolPtr(true),
				InitStateFile: core.StringPtr("testString"),
				Injectors: []schematicsv1.InjectTerraformTemplateItem{*injectTerraformTemplateItemModel},
				Type: core.StringPtr("testString"),
				UninstallScriptName: core.StringPtr("testString"),
				Values: core.StringPtr("testString"),
				ValuesMetadata: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Variablestore: []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel},
			}

			templateRepoUpdateRequestModel := &schematicsv1.TemplateRepoUpdateRequest{
				Branch: core.StringPtr("testString"),
				Release: core.StringPtr("testString"),
				RepoShaValue: core.StringPtr("testString"),
				RepoURL: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			workspaceStatusUpdateRequestModel := &schematicsv1.WorkspaceStatusUpdateRequest{
				Frozen: core.BoolPtr(true),
				FrozenAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				FrozenBy: core.StringPtr("testString"),
				Locked: core.BoolPtr(true),
				LockedBy: core.StringPtr("testString"),
				LockedTime: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			workspaceStatusMessageModel := &schematicsv1.WorkspaceStatusMessage{
				StatusCode: core.StringPtr("testString"),
				StatusMsg: core.StringPtr("testString"),
			}

			replaceWorkspaceOptions := &schematicsv1.ReplaceWorkspaceOptions{
				WID: core.StringPtr("testString"),
				CatalogRef: catalogRefModel,
				Description: core.StringPtr("testString"),
				Dependencies: dependenciesModel,
				Name: core.StringPtr("testString"),
				SharedData: sharedTargetDataModel,
				Tags: []string{"testString"},
				TemplateData: []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel},
				TemplateRepo: templateRepoUpdateRequestModel,
				Type: []string{"testString"},
				WorkspaceStatus: workspaceStatusUpdateRequestModel,
				WorkspaceStatusMsg: workspaceStatusMessageModel,
				AgentID: core.StringPtr("testString"),
				XGithubToken: core.StringPtr("testString"),
			}

			workspaceResponse, response, err := schematicsService.ReplaceWorkspace(replaceWorkspaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())
		})
	})

	Describe(`ReplaceWorkspaceInputs - Replace workspace input variables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions *ReplaceWorkspaceInputsOptions)`, func() {
			workspaceVariableRequestModel := &schematicsv1.WorkspaceVariableRequest{
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
				Type: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Value: core.StringPtr("testString"),
			}

			replaceWorkspaceInputsOptions := &schematicsv1.ReplaceWorkspaceInputsOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
				EnvValues: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Values: core.StringPtr("testString"),
				Variablestore: []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel},
			}

			userValues, response, err := schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userValues).ToNot(BeNil())
		})
	})

	Describe(`TemplateRepoUpload - Upload a TAR file to your workspace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TemplateRepoUpload(templateRepoUploadOptions *TemplateRepoUploadOptions)`, func() {
			templateRepoUploadOptions := &schematicsv1.TemplateRepoUploadOptions{
				WID: core.StringPtr("testString"),
				TID: core.StringPtr("testString"),
				File: CreateMockReader("This is a mock file."),
				FileContentType: core.StringPtr("testString"),
			}

			templateRepoTarUploadResponse, response, err := schematicsService.TemplateRepoUpload(templateRepoUploadOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateRepoTarUploadResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateWorkspace - Update workspace metadata`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)`, func() {
			serviceExtensionsModel := &schematicsv1.ServiceExtensions{
				Name: core.StringPtr("flavor"),
				Value: core.StringPtr("testString"),
				Type: core.StringPtr("string"),
			}

			catalogRefModel := &schematicsv1.CatalogRef{
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				ServiceExtensions: []schematicsv1.ServiceExtensions{*serviceExtensionsModel},
			}

			dependenciesModel := &schematicsv1.Dependencies{
				Parents: []string{"testString"},
				Children: []string{"testString"},
			}

			sharedTargetDataModel := &schematicsv1.SharedTargetData{
				ClusterCreatedOn: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				ClusterName: core.StringPtr("testString"),
				ClusterType: core.StringPtr("testString"),
				EntitlementKeys: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Namespace: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				WorkerCount: core.Int64Ptr(int64(26)),
				WorkerMachineType: core.StringPtr("testString"),
			}

			environmentValuesMetadataModel := &schematicsv1.EnvironmentValuesMetadata{
				Hidden: core.BoolPtr(true),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
			}

			injectTerraformTemplateItemTftParametersItemModel := &schematicsv1.InjectTerraformTemplateItemTftParametersItem{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			injectTerraformTemplateItemModel := &schematicsv1.InjectTerraformTemplateItem{
				TftGitURL: core.StringPtr("testString"),
				TftGitToken: core.StringPtr("testString"),
				TftPrefix: core.StringPtr("testString"),
				InjectionType: core.StringPtr("testString"),
				TftName: core.StringPtr("testString"),
				TftParameters: []schematicsv1.InjectTerraformTemplateItemTftParametersItem{*injectTerraformTemplateItemTftParametersItemModel},
			}

			workspaceVariableRequestModel := &schematicsv1.WorkspaceVariableRequest{
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Secure: core.BoolPtr(true),
				Type: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Value: core.StringPtr("testString"),
			}

			templateSourceDataRequestModel := &schematicsv1.TemplateSourceDataRequest{
				EnvValues: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				EnvValuesMetadata: []schematicsv1.EnvironmentValuesMetadata{*environmentValuesMetadataModel},
				Folder: core.StringPtr("testString"),
				Compact: core.BoolPtr(true),
				InitStateFile: core.StringPtr("testString"),
				Injectors: []schematicsv1.InjectTerraformTemplateItem{*injectTerraformTemplateItemModel},
				Type: core.StringPtr("testString"),
				UninstallScriptName: core.StringPtr("testString"),
				Values: core.StringPtr("testString"),
				ValuesMetadata: []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}},
				Variablestore: []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel},
			}

			templateRepoUpdateRequestModel := &schematicsv1.TemplateRepoUpdateRequest{
				Branch: core.StringPtr("testString"),
				Release: core.StringPtr("testString"),
				RepoShaValue: core.StringPtr("testString"),
				RepoURL: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			workspaceStatusUpdateRequestModel := &schematicsv1.WorkspaceStatusUpdateRequest{
				Frozen: core.BoolPtr(true),
				FrozenAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				FrozenBy: core.StringPtr("testString"),
				Locked: core.BoolPtr(true),
				LockedBy: core.StringPtr("testString"),
				LockedTime: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			workspaceStatusMessageModel := &schematicsv1.WorkspaceStatusMessage{
				StatusCode: core.StringPtr("testString"),
				StatusMsg: core.StringPtr("testString"),
			}

			updateWorkspaceOptions := &schematicsv1.UpdateWorkspaceOptions{
				WID: core.StringPtr("testString"),
				CatalogRef: catalogRefModel,
				Description: core.StringPtr("testString"),
				Dependencies: dependenciesModel,
				Name: core.StringPtr("testString"),
				SharedData: sharedTargetDataModel,
				Tags: []string{"testString"},
				TemplateData: []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel},
				TemplateRepo: templateRepoUpdateRequestModel,
				Type: []string{"testString"},
				WorkspaceStatus: workspaceStatusUpdateRequestModel,
				WorkspaceStatusMsg: workspaceStatusMessageModel,
				AgentID: core.StringPtr("testString"),
			}

			workspaceResponse, response, err := schematicsService.UpdateWorkspace(updateWorkspaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateAction - Create an action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAction(createActionOptions *CreateActionOptions)`, func() {
			userStateModel := &schematicsv1.UserState{
				State: core.StringPtr("draft"),
				SetBy: core.StringPtr("testString"),
				SetAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			credentialVariableMetadataModel := &schematicsv1.CredentialVariableMetadata{
				Type: core.StringPtr("string"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			credentialVariableDataModel := &schematicsv1.CredentialVariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: credentialVariableMetadataModel,
			}

			bastionResourceDefinitionModel := &schematicsv1.BastionResourceDefinition{
				Name: core.StringPtr("testString"),
				Host: core.StringPtr("testString"),
			}

			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			createActionOptions := &schematicsv1.CreateActionOptions{
				Name: core.StringPtr("Stop Action"),
				Description: core.StringPtr("The description of your action. The description can be up to 2048 characters long in size. **Example** you can use the description to stop the targets."),
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("testString"),
				BastionConnectionType: core.StringPtr("ssh"),
				InventoryConnectionType: core.StringPtr("ssh"),
				Tags: []string{"testString"},
				UserState: userStateModel,
				SourceReadmeURL: core.StringPtr("testString"),
				Source: externalSourceModel,
				SourceType: core.StringPtr("local"),
				CommandParameter: core.StringPtr("testString"),
				Inventory: core.StringPtr("testString"),
				Credentials: []schematicsv1.CredentialVariableData{*credentialVariableDataModel},
				Bastion: bastionResourceDefinitionModel,
				BastionCredential: credentialVariableDataModel,
				TargetsIni: core.StringPtr("testString"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				XGithubToken: core.StringPtr("testString"),
			}

			action, response, err := schematicsService.CreateAction(createActionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(action).ToNot(BeNil())
		})
	})

	Describe(`GetAction - Get action details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAction(getActionOptions *GetActionOptions)`, func() {
			getActionOptions := &schematicsv1.GetActionOptions{
				ActionID: core.StringPtr("testString"),
				Profile: core.StringPtr("summary"),
			}

			action, response, err := schematicsService.GetAction(getActionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(action).ToNot(BeNil())
		})
	})

	Describe(`ListActions - List actions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListActions(listActionsOptions *ListActionsOptions)`, func() {
			listActionsOptions := &schematicsv1.ListActionsOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Sort: core.StringPtr("testString"),
				Profile: core.StringPtr("ids"),
			}

			actionList, response, err := schematicsService.ListActions(listActionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionList).ToNot(BeNil())
		})
	})

	Describe(`UpdateAction - Update an action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAction(updateActionOptions *UpdateActionOptions)`, func() {
			userStateModel := &schematicsv1.UserState{
				State: core.StringPtr("draft"),
				SetBy: core.StringPtr("testString"),
				SetAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			credentialVariableMetadataModel := &schematicsv1.CredentialVariableMetadata{
				Type: core.StringPtr("string"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			credentialVariableDataModel := &schematicsv1.CredentialVariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: credentialVariableMetadataModel,
			}

			bastionResourceDefinitionModel := &schematicsv1.BastionResourceDefinition{
				Name: core.StringPtr("testString"),
				Host: core.StringPtr("testString"),
			}

			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			updateActionOptions := &schematicsv1.UpdateActionOptions{
				ActionID: core.StringPtr("testString"),
				Name: core.StringPtr("Stop Action"),
				Description: core.StringPtr("The description of your action. The description can be up to 2048 characters long in size. **Example** you can use the description to stop the targets."),
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("testString"),
				BastionConnectionType: core.StringPtr("ssh"),
				InventoryConnectionType: core.StringPtr("ssh"),
				Tags: []string{"testString"},
				UserState: userStateModel,
				SourceReadmeURL: core.StringPtr("testString"),
				Source: externalSourceModel,
				SourceType: core.StringPtr("local"),
				CommandParameter: core.StringPtr("testString"),
				Inventory: core.StringPtr("testString"),
				Credentials: []schematicsv1.CredentialVariableData{*credentialVariableDataModel},
				Bastion: bastionResourceDefinitionModel,
				BastionCredential: credentialVariableDataModel,
				TargetsIni: core.StringPtr("testString"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				XGithubToken: core.StringPtr("testString"),
			}

			action, response, err := schematicsService.UpdateAction(updateActionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(action).ToNot(BeNil())
		})
	})

	Describe(`UploadTemplateTarAction - Upload a TAR file to an action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UploadTemplateTarAction(uploadTemplateTarActionOptions *UploadTemplateTarActionOptions)`, func() {
			uploadTemplateTarActionOptions := &schematicsv1.UploadTemplateTarActionOptions{
				ActionID: core.StringPtr("testString"),
				File: CreateMockReader("This is a mock file."),
				FileContentType: core.StringPtr("testString"),
			}

			templateRepoTarUploadResponse, response, err := schematicsService.UploadTemplateTarAction(uploadTemplateTarActionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateRepoTarUploadResponse).ToNot(BeNil())
		})
	})

	Describe(`ApplyWorkspaceCommand - Perform a Schematics `apply` job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ApplyWorkspaceCommand(applyWorkspaceCommandOptions *ApplyWorkspaceCommandOptions)`, func() {
			workspaceActivityOptionsTemplateModel := &schematicsv1.WorkspaceActivityOptionsTemplate{
				Target: []string{"testString"},
				TfVars: []string{"testString"},
			}

			applyWorkspaceCommandOptions := &schematicsv1.ApplyWorkspaceCommandOptions{
				WID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				ActionOptions: workspaceActivityOptionsTemplateModel,
				DelegatedToken: core.StringPtr("testString"),
			}

			workspaceActivityApplyResult, response, err := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityApplyResult).ToNot(BeNil())
		})
	})

	Describe(`CreateJob - Create a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateJob(createJobOptions *CreateJobOptions)`, func() {
			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			jobStatusWorkitemModel := &schematicsv1.JobStatusWorkitem{
				WorkspaceID: core.StringPtr("testString"),
				WorkspaceName: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusFlowModel := &schematicsv1.JobStatusFlow{
				FlowID: core.StringPtr("testString"),
				FlowName: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				Workitems: []schematicsv1.JobStatusWorkitem{*jobStatusWorkitemModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusTemplateModel := &schematicsv1.JobStatusTemplate{
				TemplateID: core.StringPtr("testString"),
				TemplateName: core.StringPtr("testString"),
				FlowIndex: core.Int64Ptr(int64(38)),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusWorkspaceModel := &schematicsv1.JobStatusWorkspace{
				WorkspaceName: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				FlowStatus: jobStatusFlowModel,
				TemplateStatus: []schematicsv1.JobStatusTemplate{*jobStatusTemplateModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusActionModel := &schematicsv1.JobStatusAction{
				ActionName: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				BastionStatusCode: core.StringPtr("none"),
				BastionStatusMessage: core.StringPtr("testString"),
				TargetsStatusCode: core.StringPtr("none"),
				TargetsStatusMessage: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusSchematicsResourcesModel := &schematicsv1.JobStatusSchematicsResources{
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				SchematicsResourceID: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusSystemModel := &schematicsv1.JobStatusSystem{
				SystemStatusMessage: core.StringPtr("testString"),
				SystemStatusCode: core.StringPtr("job_pending"),
				SchematicsResourceStatus: []schematicsv1.JobStatusSchematicsResources{*jobStatusSchematicsResourcesModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusModel := &schematicsv1.JobStatus{
				PositionInQueue: core.Float64Ptr(float64(72.5)),
				TotalInQueue: core.Float64Ptr(float64(72.5)),
				WorkspaceJobStatus: jobStatusWorkspaceModel,
				ActionJobStatus: jobStatusActionModel,
				SystemJobStatus: jobStatusSystemModel,
				FlowJobStatus: jobStatusFlowModel,
			}

			cartOrderDataModel := &schematicsv1.CartOrderData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				UsageKind: []string{"servicetags"},
			}

			jobDataTemplateModel := &schematicsv1.JobDataTemplate{
				TemplateID: core.StringPtr("testString"),
				TemplateName: core.StringPtr("testString"),
				FlowIndex: core.Int64Ptr(int64(38)),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobDataWorkspaceModel := &schematicsv1.JobDataWorkspace{
				WorkspaceName: core.StringPtr("testString"),
				FlowID: core.StringPtr("testString"),
				FlowName: core.StringPtr("testString"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				TemplateData: []schematicsv1.JobDataTemplate{*jobDataTemplateModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			inventoryResourceRecordModel := &schematicsv1.InventoryResourceRecord{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("testString"),
				InventoriesIni: core.StringPtr("testString"),
				ResourceQueries: []string{"testString"},
			}

			jobDataActionModel := &schematicsv1.JobDataAction{
				ActionName: core.StringPtr("testString"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				InventoryRecord: inventoryResourceRecordModel,
				MaterializedInventory: core.StringPtr("testString"),
			}

			jobDataSystemModel := &schematicsv1.JobDataSystem{
				KeyID: core.StringPtr("testString"),
				SchematicsResourceID: []string{"testString"},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			jobDataWorkItemLastJobModel := &schematicsv1.JobDataWorkItemLastJob{
				CommandObject: core.StringPtr("workspace"),
				CommandObjectName: core.StringPtr("testString"),
				CommandObjectID: core.StringPtr("testString"),
				CommandName: core.StringPtr("workspace_plan"),
				JobID: core.StringPtr("testString"),
				JobStatus: core.StringPtr("job_pending"),
			}

			jobDataWorkItemModel := &schematicsv1.JobDataWorkItem{
				CommandObjectID: core.StringPtr("testString"),
				CommandObjectName: core.StringPtr("testString"),
				Layers: core.StringPtr("testString"),
				SourceType: core.StringPtr("local"),
				Source: externalSourceModel,
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				LastJob: jobDataWorkItemLastJobModel,
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobDataFlowModel := &schematicsv1.JobDataFlow{
				FlowID: core.StringPtr("testString"),
				FlowName: core.StringPtr("testString"),
				Workitems: []schematicsv1.JobDataWorkItem{*jobDataWorkItemModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobDataModel := &schematicsv1.JobData{
				JobType: core.StringPtr("repo_download_job"),
				WorkspaceJobData: jobDataWorkspaceModel,
				ActionJobData: jobDataActionModel,
				SystemJobData: jobDataSystemModel,
				FlowJobData: jobDataFlowModel,
			}

			bastionResourceDefinitionModel := &schematicsv1.BastionResourceDefinition{
				Name: core.StringPtr("testString"),
				Host: core.StringPtr("testString"),
			}

			jobLogSummaryRepoDownloadJobModel := &schematicsv1.JobLogSummaryRepoDownloadJob{
			}

			jobLogSummaryWorkspaceJobModel := &schematicsv1.JobLogSummaryWorkspaceJob{
			}

			jobLogSummaryWorkitemsModel := &schematicsv1.JobLogSummaryWorkitems{
				WorkspaceID: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
				LogURL: core.StringPtr("testString"),
			}

			jobLogSummaryFlowJobModel := &schematicsv1.JobLogSummaryFlowJob{
				Workitems: []schematicsv1.JobLogSummaryWorkitems{*jobLogSummaryWorkitemsModel},
			}

			jobLogSummaryActionJobRecapModel := &schematicsv1.JobLogSummaryActionJobRecap{
				Target: []string{"testString"},
				Ok: core.Float64Ptr(float64(72.5)),
				Changed: core.Float64Ptr(float64(72.5)),
				Failed: core.Float64Ptr(float64(72.5)),
				Skipped: core.Float64Ptr(float64(72.5)),
				Unreachable: core.Float64Ptr(float64(72.5)),
			}

			jobLogSummaryActionJobModel := &schematicsv1.JobLogSummaryActionJob{
				Recap: jobLogSummaryActionJobRecapModel,
			}

			jobLogSummarySystemJobModel := &schematicsv1.JobLogSummarySystemJob{
				Success: core.Float64Ptr(float64(72.5)),
				Failed: core.Float64Ptr(float64(72.5)),
			}

			jobLogSummaryModel := &schematicsv1.JobLogSummary{
				JobType: core.StringPtr("repo_download_job"),
				RepoDownloadJob: jobLogSummaryRepoDownloadJobModel,
				WorkspaceJob: jobLogSummaryWorkspaceJobModel,
				FlowJob: jobLogSummaryFlowJobModel,
				ActionJob: jobLogSummaryActionJobModel,
				SystemJob: jobLogSummarySystemJobModel,
			}

			agentInfoModel := &schematicsv1.AgentInfo{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				AssignmentPolicyID: core.StringPtr("testString"),
			}

			createJobOptions := &schematicsv1.CreateJobOptions{
				RefreshToken: core.StringPtr("testString"),
				CommandObject: core.StringPtr("workspace"),
				CommandObjectID: core.StringPtr("testString"),
				CommandName: core.StringPtr("workspace_plan"),
				CommandParameter: core.StringPtr("testString"),
				CommandOptions: []string{"testString"},
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				Tags: []string{"testString"},
				Location: core.StringPtr("us-south"),
				Status: jobStatusModel,
				CartOrderData: []schematicsv1.CartOrderData{*cartOrderDataModel},
				Data: jobDataModel,
				Bastion: bastionResourceDefinitionModel,
				LogSummary: jobLogSummaryModel,
				Agent: agentInfoModel,
			}

			job, response, err := schematicsService.CreateJob(createJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`DestroyWorkspaceCommand - Perform a Schematics `destroy` job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DestroyWorkspaceCommand(destroyWorkspaceCommandOptions *DestroyWorkspaceCommandOptions)`, func() {
			workspaceActivityOptionsTemplateModel := &schematicsv1.WorkspaceActivityOptionsTemplate{
				Target: []string{"testString"},
				TfVars: []string{"testString"},
			}

			destroyWorkspaceCommandOptions := &schematicsv1.DestroyWorkspaceCommandOptions{
				WID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				ActionOptions: workspaceActivityOptionsTemplateModel,
				DelegatedToken: core.StringPtr("testString"),
			}

			workspaceActivityDestroyResult, response, err := schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityDestroyResult).ToNot(BeNil())
		})
	})

	Describe(`GetJob - Get a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetJob(getJobOptions *GetJobOptions)`, func() {
			getJobOptions := &schematicsv1.GetJobOptions{
				JobID: core.StringPtr("testString"),
				Profile: core.StringPtr("summary"),
			}

			job, response, err := schematicsService.GetJob(getJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`GetJobFiles - Get output files from the Job record`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetJobFiles(getJobFilesOptions *GetJobFilesOptions)`, func() {
			getJobFilesOptions := &schematicsv1.GetJobFilesOptions{
				JobID: core.StringPtr("testString"),
				FileType: core.StringPtr("template_repo"),
			}

			jobFileData, response, err := schematicsService.GetJobFiles(getJobFilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobFileData).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceActivity - Get workspace job details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceActivity(getWorkspaceActivityOptions *GetWorkspaceActivityOptions)`, func() {
			getWorkspaceActivityOptions := &schematicsv1.GetWorkspaceActivityOptions{
				WID: core.StringPtr("testString"),
				ActivityID: core.StringPtr("testString"),
			}

			workspaceActivity, response, err := schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivity).ToNot(BeNil())
		})
	})

	Describe(`ListJobLogs - Get job logs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListJobLogs(listJobLogsOptions *ListJobLogsOptions)`, func() {
			listJobLogsOptions := &schematicsv1.ListJobLogsOptions{
				JobID: core.StringPtr("testString"),
			}

			jobLog, response, err := schematicsService.ListJobLogs(listJobLogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobLog).ToNot(BeNil())
		})
	})

	Describe(`ListJobs - List jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListJobs(listJobsOptions *ListJobsOptions)`, func() {
			listJobsOptions := &schematicsv1.ListJobsOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Sort: core.StringPtr("testString"),
				Profile: core.StringPtr("ids"),
				Resource: core.StringPtr("workspaces"),
				ResourceID: core.StringPtr("testString"),
				ActionID: core.StringPtr("testString"),
				WorkspaceID: core.StringPtr("testString"),
				List: core.StringPtr("all"),
			}

			jobList, response, err := schematicsService.ListJobs(listJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobList).ToNot(BeNil())
		})
	})

	Describe(`ListWorkspaceActivities - List workspace jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListWorkspaceActivities(listWorkspaceActivitiesOptions *ListWorkspaceActivitiesOptions)`, func() {
			listWorkspaceActivitiesOptions := &schematicsv1.ListWorkspaceActivitiesOptions{
				WID: core.StringPtr("testString"),
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
			}

			workspaceActivities, response, err := schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivities).ToNot(BeNil())
		})
	})

	Describe(`PlanWorkspaceCommand - Perform a Schematics `plan` job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PlanWorkspaceCommand(planWorkspaceCommandOptions *PlanWorkspaceCommandOptions)`, func() {
			workspaceActivityOptionsTemplateModel := &schematicsv1.WorkspaceActivityOptionsTemplate{
				Target: []string{"testString"},
				TfVars: []string{"testString"},
			}

			planWorkspaceCommandOptions := &schematicsv1.PlanWorkspaceCommandOptions{
				WID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				ActionOptions: workspaceActivityOptionsTemplateModel,
				DelegatedToken: core.StringPtr("testString"),
			}

			workspaceActivityPlanResult, response, err := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityPlanResult).ToNot(BeNil())
		})
	})

	Describe(`RefreshWorkspaceCommand - Perform a Schematics `refresh` job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RefreshWorkspaceCommand(refreshWorkspaceCommandOptions *RefreshWorkspaceCommandOptions)`, func() {
			refreshWorkspaceCommandOptions := &schematicsv1.RefreshWorkspaceCommandOptions{
				WID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				DelegatedToken: core.StringPtr("testString"),
			}

			workspaceActivityRefreshResult, response, err := schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityRefreshResult).ToNot(BeNil())
		})
	})

	Describe(`RunWorkspaceCommands - Run Terraform Commands`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunWorkspaceCommands(runWorkspaceCommandsOptions *RunWorkspaceCommandsOptions)`, func() {
			terraformCommandModel := &schematicsv1.TerraformCommand{
				Command: core.StringPtr("testString"),
				CommandParams: core.StringPtr("testString"),
				CommandName: core.StringPtr("testString"),
				CommandDesc: core.StringPtr("testString"),
				CommandOnError: core.StringPtr("testString"),
				CommandDependsOn: core.StringPtr("testString"),
				CommandStatus: core.StringPtr("testString"),
			}

			runWorkspaceCommandsOptions := &schematicsv1.RunWorkspaceCommandsOptions{
				WID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				Commands: []schematicsv1.TerraformCommand{*terraformCommandModel},
				OperationName: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			workspaceActivityCommandResult, response, err := schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivityCommandResult).ToNot(BeNil())
		})
	})

	Describe(`UpdateJob - Update a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateJob(updateJobOptions *UpdateJobOptions)`, func() {
			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			jobStatusWorkitemModel := &schematicsv1.JobStatusWorkitem{
				WorkspaceID: core.StringPtr("testString"),
				WorkspaceName: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusFlowModel := &schematicsv1.JobStatusFlow{
				FlowID: core.StringPtr("testString"),
				FlowName: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				Workitems: []schematicsv1.JobStatusWorkitem{*jobStatusWorkitemModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusTemplateModel := &schematicsv1.JobStatusTemplate{
				TemplateID: core.StringPtr("testString"),
				TemplateName: core.StringPtr("testString"),
				FlowIndex: core.Int64Ptr(int64(38)),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusWorkspaceModel := &schematicsv1.JobStatusWorkspace{
				WorkspaceName: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				FlowStatus: jobStatusFlowModel,
				TemplateStatus: []schematicsv1.JobStatusTemplate{*jobStatusTemplateModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusActionModel := &schematicsv1.JobStatusAction{
				ActionName: core.StringPtr("testString"),
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				BastionStatusCode: core.StringPtr("none"),
				BastionStatusMessage: core.StringPtr("testString"),
				TargetsStatusCode: core.StringPtr("none"),
				TargetsStatusMessage: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusSchematicsResourcesModel := &schematicsv1.JobStatusSchematicsResources{
				StatusCode: core.StringPtr("job_pending"),
				StatusMessage: core.StringPtr("testString"),
				SchematicsResourceID: core.StringPtr("testString"),
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusSystemModel := &schematicsv1.JobStatusSystem{
				SystemStatusMessage: core.StringPtr("testString"),
				SystemStatusCode: core.StringPtr("job_pending"),
				SchematicsResourceStatus: []schematicsv1.JobStatusSchematicsResources{*jobStatusSchematicsResourcesModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobStatusModel := &schematicsv1.JobStatus{
				PositionInQueue: core.Float64Ptr(float64(72.5)),
				TotalInQueue: core.Float64Ptr(float64(72.5)),
				WorkspaceJobStatus: jobStatusWorkspaceModel,
				ActionJobStatus: jobStatusActionModel,
				SystemJobStatus: jobStatusSystemModel,
				FlowJobStatus: jobStatusFlowModel,
			}

			cartOrderDataModel := &schematicsv1.CartOrderData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				UsageKind: []string{"servicetags"},
			}

			jobDataTemplateModel := &schematicsv1.JobDataTemplate{
				TemplateID: core.StringPtr("testString"),
				TemplateName: core.StringPtr("testString"),
				FlowIndex: core.Int64Ptr(int64(38)),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobDataWorkspaceModel := &schematicsv1.JobDataWorkspace{
				WorkspaceName: core.StringPtr("testString"),
				FlowID: core.StringPtr("testString"),
				FlowName: core.StringPtr("testString"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				TemplateData: []schematicsv1.JobDataTemplate{*jobDataTemplateModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			inventoryResourceRecordModel := &schematicsv1.InventoryResourceRecord{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("testString"),
				InventoriesIni: core.StringPtr("testString"),
				ResourceQueries: []string{"testString"},
			}

			jobDataActionModel := &schematicsv1.JobDataAction{
				ActionName: core.StringPtr("testString"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				InventoryRecord: inventoryResourceRecordModel,
				MaterializedInventory: core.StringPtr("testString"),
			}

			jobDataSystemModel := &schematicsv1.JobDataSystem{
				KeyID: core.StringPtr("testString"),
				SchematicsResourceID: []string{"testString"},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			jobDataWorkItemLastJobModel := &schematicsv1.JobDataWorkItemLastJob{
				CommandObject: core.StringPtr("workspace"),
				CommandObjectName: core.StringPtr("testString"),
				CommandObjectID: core.StringPtr("testString"),
				CommandName: core.StringPtr("workspace_plan"),
				JobID: core.StringPtr("testString"),
				JobStatus: core.StringPtr("job_pending"),
			}

			jobDataWorkItemModel := &schematicsv1.JobDataWorkItem{
				CommandObjectID: core.StringPtr("testString"),
				CommandObjectName: core.StringPtr("testString"),
				Layers: core.StringPtr("testString"),
				SourceType: core.StringPtr("local"),
				Source: externalSourceModel,
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Outputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				LastJob: jobDataWorkItemLastJobModel,
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobDataFlowModel := &schematicsv1.JobDataFlow{
				FlowID: core.StringPtr("testString"),
				FlowName: core.StringPtr("testString"),
				Workitems: []schematicsv1.JobDataWorkItem{*jobDataWorkItemModel},
				UpdatedAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			jobDataModel := &schematicsv1.JobData{
				JobType: core.StringPtr("repo_download_job"),
				WorkspaceJobData: jobDataWorkspaceModel,
				ActionJobData: jobDataActionModel,
				SystemJobData: jobDataSystemModel,
				FlowJobData: jobDataFlowModel,
			}

			bastionResourceDefinitionModel := &schematicsv1.BastionResourceDefinition{
				Name: core.StringPtr("testString"),
				Host: core.StringPtr("testString"),
			}

			jobLogSummaryRepoDownloadJobModel := &schematicsv1.JobLogSummaryRepoDownloadJob{
			}

			jobLogSummaryWorkspaceJobModel := &schematicsv1.JobLogSummaryWorkspaceJob{
			}

			jobLogSummaryWorkitemsModel := &schematicsv1.JobLogSummaryWorkitems{
				WorkspaceID: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
				LogURL: core.StringPtr("testString"),
			}

			jobLogSummaryFlowJobModel := &schematicsv1.JobLogSummaryFlowJob{
				Workitems: []schematicsv1.JobLogSummaryWorkitems{*jobLogSummaryWorkitemsModel},
			}

			jobLogSummaryActionJobRecapModel := &schematicsv1.JobLogSummaryActionJobRecap{
				Target: []string{"testString"},
				Ok: core.Float64Ptr(float64(72.5)),
				Changed: core.Float64Ptr(float64(72.5)),
				Failed: core.Float64Ptr(float64(72.5)),
				Skipped: core.Float64Ptr(float64(72.5)),
				Unreachable: core.Float64Ptr(float64(72.5)),
			}

			jobLogSummaryActionJobModel := &schematicsv1.JobLogSummaryActionJob{
				Recap: jobLogSummaryActionJobRecapModel,
			}

			jobLogSummarySystemJobModel := &schematicsv1.JobLogSummarySystemJob{
				Success: core.Float64Ptr(float64(72.5)),
				Failed: core.Float64Ptr(float64(72.5)),
			}

			jobLogSummaryModel := &schematicsv1.JobLogSummary{
				JobType: core.StringPtr("repo_download_job"),
				RepoDownloadJob: jobLogSummaryRepoDownloadJobModel,
				WorkspaceJob: jobLogSummaryWorkspaceJobModel,
				FlowJob: jobLogSummaryFlowJobModel,
				ActionJob: jobLogSummaryActionJobModel,
				SystemJob: jobLogSummarySystemJobModel,
			}

			agentInfoModel := &schematicsv1.AgentInfo{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				AssignmentPolicyID: core.StringPtr("testString"),
			}

			updateJobOptions := &schematicsv1.UpdateJobOptions{
				JobID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				CommandObject: core.StringPtr("workspace"),
				CommandObjectID: core.StringPtr("testString"),
				CommandName: core.StringPtr("workspace_plan"),
				CommandParameter: core.StringPtr("testString"),
				CommandOptions: []string{"testString"},
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				Tags: []string{"testString"},
				Location: core.StringPtr("us-south"),
				Status: jobStatusModel,
				CartOrderData: []schematicsv1.CartOrderData{*cartOrderDataModel},
				Data: jobDataModel,
				Bastion: bastionResourceDefinitionModel,
				LogSummary: jobLogSummaryModel,
				Agent: agentInfoModel,
			}

			job, response, err := schematicsService.UpdateJob(updateJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`CreateWorkspaceDeletionJob - Delete one or more workspace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptions *CreateWorkspaceDeletionJobOptions)`, func() {
			createWorkspaceDeletionJobOptions := &schematicsv1.CreateWorkspaceDeletionJobOptions{
				RefreshToken: core.StringPtr("testString"),
				Job: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Workspaces: []string{"testString"},
			}

			workspaceBulkDeleteResponse, response, err := schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceBulkDeleteResponse).ToNot(BeNil())
		})
	})

	Describe(`GetWorkspaceDeletionJobStatus - Get the workspace deletion job status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptions *GetWorkspaceDeletionJobStatusOptions)`, func() {
			getWorkspaceDeletionJobStatusOptions := &schematicsv1.GetWorkspaceDeletionJobStatusOptions{
				WjID: core.StringPtr("testString"),
			}

			workspaceJobResponse, response, err := schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceJobResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateBlueprint - Create a blueprint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBlueprint(createBlueprintOptions *CreateBlueprintOptions)`, func() {
			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			blueprintConfigItemModel := &schematicsv1.BlueprintConfigItem{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Source: externalSourceModel,
				Inputs: []schematicsv1.VariableData{*variableDataModel},
			}

			blueprintFlowModel := &schematicsv1.BlueprintFlow{
			}

			userStateModel := &schematicsv1.UserState{
				State: core.StringPtr("draft"),
				SetBy: core.StringPtr("testString"),
				SetAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			createBlueprintOptions := &schematicsv1.CreateBlueprintOptions{
				Name: core.StringPtr("Toronto Dev Environtment"),
				SchemaVersion: core.StringPtr("1.0"),
				Source: externalSourceModel,
				Config: []schematicsv1.BlueprintConfigItem{*blueprintConfigItemModel},
				Description: core.StringPtr("Deploys dev environtment instance in Toronto Region"),
				ResourceGroup: core.StringPtr("Default"),
				Tags: []string{"blueprint:Tor-Dev"},
				Location: core.StringPtr("us-south"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				Flow: blueprintFlowModel,
				UserState: userStateModel,
			}

			blueprint, response, err := schematicsService.CreateBlueprint(createBlueprintOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())
		})
	})

	Describe(`GetBlueprint - Get a blueprint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBlueprint(getBlueprintOptions *GetBlueprintOptions)`, func() {
			getBlueprintOptions := &schematicsv1.GetBlueprintOptions{
				BlueprintID: core.StringPtr("testString"),
				Profile: core.StringPtr("ids"),
			}

			blueprint, response, err := schematicsService.GetBlueprint(getBlueprintOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())
		})
	})

	Describe(`ListBlueprint - List blueprint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBlueprint(listBlueprintOptions *ListBlueprintOptions)`, func() {
			listBlueprintOptions := &schematicsv1.ListBlueprintOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
			}

			blueprintList, response, err := schematicsService.ListBlueprint(listBlueprintOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprintList).ToNot(BeNil())
		})
	})

	Describe(`ReplaceBlueprint - Update a blueprint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceBlueprint(replaceBlueprintOptions *ReplaceBlueprintOptions)`, func() {
			gitSourceModel := &schematicsv1.GitSource{
				ComputedGitRepoURL: core.StringPtr("testString"),
				GitRepoURL: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
				GitRepoFolder: core.StringPtr("testString"),
				GitRelease: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
			}

			catalogSourceModel := &schematicsv1.CatalogSource{
				CatalogName: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				OfferingName: core.StringPtr("testString"),
				OfferingVersion: core.StringPtr("testString"),
				OfferingKind: core.StringPtr("testString"),
				OfferingTargetKind: core.StringPtr("testString"),
				OfferingID: core.StringPtr("testString"),
				OfferingVersionID: core.StringPtr("testString"),
				OfferingVersionFlavourName: core.StringPtr("testString"),
				OfferingRepoURL: core.StringPtr("testString"),
				OfferingProvisionerWorkingDirectory: core.StringPtr("testString"),
				DryRun: core.BoolPtr(true),
				OwningAccount: core.StringPtr("testString"),
				ItemIconURL: core.StringPtr("testString"),
				ItemID: core.StringPtr("testString"),
				ItemName: core.StringPtr("testString"),
				ItemReadmeURL: core.StringPtr("testString"),
				ItemURL: core.StringPtr("testString"),
				LaunchURL: core.StringPtr("testString"),
			}

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
				Git: gitSourceModel,
				Catalog: catalogSourceModel,
			}

			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			blueprintConfigItemModel := &schematicsv1.BlueprintConfigItem{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Source: externalSourceModel,
				Inputs: []schematicsv1.VariableData{*variableDataModel},
			}

			blueprintFlowModel := &schematicsv1.BlueprintFlow{
			}

			userStateModel := &schematicsv1.UserState{
				State: core.StringPtr("draft"),
				SetBy: core.StringPtr("testString"),
				SetAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			replaceBlueprintOptions := &schematicsv1.ReplaceBlueprintOptions{
				BlueprintID: core.StringPtr("testString"),
				Name: core.StringPtr("Toronto Dev Environtment"),
				SchemaVersion: core.StringPtr("1.0"),
				Source: externalSourceModel,
				Config: []schematicsv1.BlueprintConfigItem{*blueprintConfigItemModel},
				Description: core.StringPtr("Deploys dev environtment instance in Toronto Region"),
				ResourceGroup: core.StringPtr("Default"),
				Tags: []string{"blueprint:Tor-Dev"},
				Location: core.StringPtr("us-south"),
				Inputs: []schematicsv1.VariableData{*variableDataModel},
				Settings: []schematicsv1.VariableData{*variableDataModel},
				Flow: blueprintFlowModel,
				UserState: userStateModel,
				Profile: core.StringPtr("ids"),
			}

			blueprint, response, err := schematicsService.ReplaceBlueprint(replaceBlueprintOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())
		})
	})

	Describe(`UploadTemplateTarBlueprint - Upload a TAR file to a blueprint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UploadTemplateTarBlueprint(uploadTemplateTarBlueprintOptions *UploadTemplateTarBlueprintOptions)`, func() {
			uploadTemplateTarBlueprintOptions := &schematicsv1.UploadTemplateTarBlueprintOptions{
				BlueprintID: core.StringPtr("testString"),
				File: CreateMockReader("This is a mock file."),
				FileContentType: core.StringPtr("testString"),
			}

			blueprintTemplateRepoTarUploadResponse, response, err := schematicsService.UploadTemplateTarBlueprint(uploadTemplateTarBlueprintOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprintTemplateRepoTarUploadResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateInventory - Create an inventory definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInventory(createInventoryOptions *CreateInventoryOptions)`, func() {
			createInventoryOptions := &schematicsv1.CreateInventoryOptions{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("testString"),
				InventoriesIni: core.StringPtr("testString"),
				ResourceQueries: []string{"testString"},
			}

			inventoryResourceRecord, response, err := schematicsService.CreateInventory(createInventoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())
		})
	})

	Describe(`CreateResourceQuery - Create resource query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceQuery(createResourceQueryOptions *CreateResourceQueryOptions)`, func() {
			resourceQueryParamModel := &schematicsv1.ResourceQueryParam{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			resourceQueryModel := &schematicsv1.ResourceQuery{
				QueryType: core.StringPtr("workspaces"),
				QueryCondition: []schematicsv1.ResourceQueryParam{*resourceQueryParamModel},
				QuerySelect: []string{"testString"},
			}

			createResourceQueryOptions := &schematicsv1.CreateResourceQueryOptions{
				Type: core.StringPtr("vsi"),
				Name: core.StringPtr("testString"),
				Queries: []schematicsv1.ResourceQuery{*resourceQueryModel},
			}

			resourceQueryRecord, response, err := schematicsService.CreateResourceQuery(createResourceQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())
		})
	})

	Describe(`ExecuteResourceQuery - Run the resource query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ExecuteResourceQuery(executeResourceQueryOptions *ExecuteResourceQueryOptions)`, func() {
			executeResourceQueryOptions := &schematicsv1.ExecuteResourceQueryOptions{
				QueryID: core.StringPtr("testString"),
			}

			resourceQueryResponseRecord, response, err := schematicsService.ExecuteResourceQuery(executeResourceQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryResponseRecord).ToNot(BeNil())
		})
	})

	Describe(`GetInventory - Get an inventory definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInventory(getInventoryOptions *GetInventoryOptions)`, func() {
			getInventoryOptions := &schematicsv1.GetInventoryOptions{
				InventoryID: core.StringPtr("testString"),
				Profile: core.StringPtr("summary"),
			}

			inventoryResourceRecord, response, err := schematicsService.GetInventory(getInventoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())
		})
	})

	Describe(`GetResourcesQuery - Get resources query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourcesQuery(getResourcesQueryOptions *GetResourcesQueryOptions)`, func() {
			getResourcesQueryOptions := &schematicsv1.GetResourcesQueryOptions{
				QueryID: core.StringPtr("testString"),
			}

			resourceQueryRecord, response, err := schematicsService.GetResourcesQuery(getResourcesQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())
		})
	})

	Describe(`ListInventories - List inventory definitions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInventories(listInventoriesOptions *ListInventoriesOptions)`, func() {
			listInventoriesOptions := &schematicsv1.ListInventoriesOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Sort: core.StringPtr("testString"),
				Profile: core.StringPtr("ids"),
			}

			inventoryResourceRecordList, response, err := schematicsService.ListInventories(listInventoriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecordList).ToNot(BeNil())
		})
	})

	Describe(`ListResourceQuery - List resource queries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceQuery(listResourceQueryOptions *ListResourceQueryOptions)`, func() {
			listResourceQueryOptions := &schematicsv1.ListResourceQueryOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Sort: core.StringPtr("testString"),
				Profile: core.StringPtr("ids"),
			}

			resourceQueryRecordList, response, err := schematicsService.ListResourceQuery(listResourceQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecordList).ToNot(BeNil())
		})
	})

	Describe(`ReplaceInventory - Update an inventory definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceInventory(replaceInventoryOptions *ReplaceInventoryOptions)`, func() {
			replaceInventoryOptions := &schematicsv1.ReplaceInventoryOptions{
				InventoryID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("testString"),
				InventoriesIni: core.StringPtr("testString"),
				ResourceQueries: []string{"testString"},
			}

			inventoryResourceRecord, response, err := schematicsService.ReplaceInventory(replaceInventoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())
		})
	})

	Describe(`ReplaceResourcesQuery - Update resources query definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceResourcesQuery(replaceResourcesQueryOptions *ReplaceResourcesQueryOptions)`, func() {
			resourceQueryParamModel := &schematicsv1.ResourceQueryParam{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			resourceQueryModel := &schematicsv1.ResourceQuery{
				QueryType: core.StringPtr("workspaces"),
				QueryCondition: []schematicsv1.ResourceQueryParam{*resourceQueryParamModel},
				QuerySelect: []string{"testString"},
			}

			replaceResourcesQueryOptions := &schematicsv1.ReplaceResourcesQueryOptions{
				QueryID: core.StringPtr("testString"),
				Type: core.StringPtr("vsi"),
				Name: core.StringPtr("testString"),
				Queries: []schematicsv1.ResourceQuery{*resourceQueryModel},
			}

			resourceQueryRecord, response, err := schematicsService.ReplaceResourcesQuery(replaceResourcesQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())
		})
	})

	Describe(`CreateAgentData - Create an agent`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAgentData(createAgentDataOptions *CreateAgentDataOptions)`, func() {
			agentInfrastructureModel := &schematicsv1.AgentInfrastructure{
				InfraType: core.StringPtr("ibm_kubernetes"),
				ClusterID: core.StringPtr("testString"),
				ClusterResourceGroup: core.StringPtr("testString"),
				CosInstanceName: core.StringPtr("testString"),
				CosBucketName: core.StringPtr("testString"),
				CosBucketRegion: core.StringPtr("testString"),
			}

			agentMetadataInfoModel := &schematicsv1.AgentMetadataInfo{
				Name: core.StringPtr("purpose"),
				Value: []string{"git", "terraform", "ansible"},
			}

			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			agentUserStateModel := &schematicsv1.AgentUserState{
				State: core.StringPtr("enable"),
			}

			agentKpiDataModel := &schematicsv1.AgentKPIData{
				AvailabilityIndicator: core.StringPtr("available"),
				LifecycleIndicator: core.StringPtr("consistent"),
				PercentUsageIndicator: core.StringPtr("testString"),
				ApplicationIndicators: []interface{}{"testString"},
				InfraIndicators: []interface{}{"testString"},
			}

			createAgentDataOptions := &schematicsv1.CreateAgentDataOptions{
				Name: core.StringPtr("MyDevAgent"),
				ResourceGroup: core.StringPtr("Default"),
				Version: core.StringPtr("v1.0.0"),
				SchematicsLocation: core.StringPtr("us-south"),
				AgentLocation: core.StringPtr("us-south"),
				AgentInfrastructure: agentInfrastructureModel,
				Description: core.StringPtr("Create Agent"),
				Tags: []string{"testString"},
				AgentMetadata: []schematicsv1.AgentMetadataInfo{*agentMetadataInfoModel},
				AgentInputs: []schematicsv1.VariableData{*variableDataModel},
				UserState: agentUserStateModel,
				AgentKpi: agentKpiDataModel,
			}

			agentData, response, err := schematicsService.CreateAgentData(createAgentDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agentData).ToNot(BeNil())
		})
	})

	Describe(`DeployAgentJob - Run the agent deployment job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeployAgentJob(deployAgentJobOptions *DeployAgentJobOptions)`, func() {
			deployAgentJobOptions := &schematicsv1.DeployAgentJobOptions{
				AgentID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
			}

			agentDeployJob, response, err := schematicsService.DeployAgentJob(deployAgentJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentDeployJob).ToNot(BeNil())
		})
	})

	Describe(`GetAgent - Get the registered agent details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAgent(getAgentOptions *GetAgentOptions)`, func() {
			getAgentOptions := &schematicsv1.GetAgentOptions{
				AgentID: core.StringPtr("testString"),
				Profile: core.StringPtr("summary"),
			}

			agent, response, err := schematicsService.GetAgent(getAgentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agent).ToNot(BeNil())
		})
	})

	Describe(`GetAgentData - Get agent details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAgentData(getAgentDataOptions *GetAgentDataOptions)`, func() {
			getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
				AgentID: core.StringPtr("testString"),
				Profile: core.StringPtr("summary"),
			}

			agentData, response, err := schematicsService.GetAgentData(getAgentDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentData).ToNot(BeNil())
		})
	})

	Describe(`GetDeployAgentJob - Get agent deployment job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDeployAgentJob(getDeployAgentJobOptions *GetDeployAgentJobOptions)`, func() {
			getDeployAgentJobOptions := &schematicsv1.GetDeployAgentJobOptions{
				AgentID: core.StringPtr("testString"),
			}

			agentDeployJob, response, err := schematicsService.GetDeployAgentJob(getDeployAgentJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentDeployJob).ToNot(BeNil())
		})
	})

	Describe(`GetHealthCheckAgentJob - Get agent health check job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetHealthCheckAgentJob(getHealthCheckAgentJobOptions *GetHealthCheckAgentJobOptions)`, func() {
			getHealthCheckAgentJobOptions := &schematicsv1.GetHealthCheckAgentJobOptions{
				AgentID: core.StringPtr("testString"),
			}

			agentHealthJob, response, err := schematicsService.GetHealthCheckAgentJob(getHealthCheckAgentJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentHealthJob).ToNot(BeNil())
		})
	})

	Describe(`GetPrsAgentJob - Get pre-requisite scanner job status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrsAgentJob(getPrsAgentJobOptions *GetPrsAgentJobOptions)`, func() {
			getPrsAgentJobOptions := &schematicsv1.GetPrsAgentJobOptions{
				AgentID: core.StringPtr("testString"),
			}

			agentPrsJob, response, err := schematicsService.GetPrsAgentJob(getPrsAgentJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentPrsJob).ToNot(BeNil())
		})
	})

	Describe(`HealthCheckAgentJob - Run agent health check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HealthCheckAgentJob(healthCheckAgentJobOptions *HealthCheckAgentJobOptions)`, func() {
			healthCheckAgentJobOptions := &schematicsv1.HealthCheckAgentJobOptions{
				AgentID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
			}

			agentHealthJob, response, err := schematicsService.HealthCheckAgentJob(healthCheckAgentJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentHealthJob).ToNot(BeNil())
		})
	})

	Describe(`ListAgent - Get all registered/unregistered agents in the Account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAgent(listAgentOptions *ListAgentOptions)`, func() {
			listAgentOptions := &schematicsv1.ListAgentOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Profile: core.StringPtr("summary"),
				Filter: core.StringPtr("all"),
			}

			agentList, response, err := schematicsService.ListAgent(listAgentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentList).ToNot(BeNil())
		})
	})

	Describe(`ListAgentData - List agents`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAgentData(listAgentDataOptions *ListAgentDataOptions)`, func() {
			listAgentDataOptions := &schematicsv1.ListAgentDataOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Profile: core.StringPtr("summary"),
				Filter: core.StringPtr("all"),
			}

			agentDataList, response, err := schematicsService.ListAgentData(listAgentDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentDataList).ToNot(BeNil())
		})
	})

	Describe(`PrsAgentJob - Run pre-requisite scanner job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PrsAgentJob(prsAgentJobOptions *PrsAgentJobOptions)`, func() {
			prsAgentJobOptions := &schematicsv1.PrsAgentJobOptions{
				AgentID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
			}

			agentPrsJob, response, err := schematicsService.PrsAgentJob(prsAgentJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentPrsJob).ToNot(BeNil())
		})
	})

	Describe(`RegisterAgent - Register the agent with schematics`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RegisterAgent(registerAgentOptions *RegisterAgentOptions)`, func() {
			agentUserStateModel := &schematicsv1.AgentUserState{
				State: core.StringPtr("enable"),
			}

			registerAgentOptions := &schematicsv1.RegisterAgentOptions{
				Name: core.StringPtr("MyDevAgent"),
				AgentLocation: core.StringPtr("us-south"),
				Location: core.StringPtr("us-south"),
				ProfileID: core.StringPtr("testString"),
				Description: core.StringPtr("Register agent"),
				ResourceGroup: core.StringPtr("testString"),
				Tags: []string{"testString"},
				UserState: agentUserStateModel,
			}

			agent, response, err := schematicsService.RegisterAgent(registerAgentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agent).ToNot(BeNil())
		})
	})

	Describe(`UpdateAgentData - Update agent`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAgentData(updateAgentDataOptions *UpdateAgentDataOptions)`, func() {
			agentInfrastructureModel := &schematicsv1.AgentInfrastructure{
				InfraType: core.StringPtr("ibm_kubernetes"),
				ClusterID: core.StringPtr("testString"),
				ClusterResourceGroup: core.StringPtr("testString"),
				CosInstanceName: core.StringPtr("testString"),
				CosBucketName: core.StringPtr("testString"),
				CosBucketRegion: core.StringPtr("testString"),
			}

			agentMetadataInfoModel := &schematicsv1.AgentMetadataInfo{
				Name: core.StringPtr("purpose"),
				Value: []string{"git", "terraform", "ansible"},
			}

			variableMetadataModel := &schematicsv1.VariableMetadata{
				Type: core.StringPtr("boolean"),
				Aliases: []string{"testString"},
				Description: core.StringPtr("testString"),
				CloudDataType: core.StringPtr("testString"),
				DefaultValue: core.StringPtr("testString"),
				LinkStatus: core.StringPtr("normal"),
				Secure: core.BoolPtr(true),
				Immutable: core.BoolPtr(true),
				Hidden: core.BoolPtr(true),
				Required: core.BoolPtr(true),
				Options: []string{"testString"},
				MinValue: core.Int64Ptr(int64(38)),
				MaxValue: core.Int64Ptr(int64(38)),
				MinLength: core.Int64Ptr(int64(38)),
				MaxLength: core.Int64Ptr(int64(38)),
				Matches: core.StringPtr("testString"),
				Position: core.Int64Ptr(int64(38)),
				GroupBy: core.StringPtr("testString"),
				Source: core.StringPtr("testString"),
			}

			variableDataModel := &schematicsv1.VariableData{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
				UseDefault: core.BoolPtr(true),
				Metadata: variableMetadataModel,
			}

			agentUserStateModel := &schematicsv1.AgentUserState{
				State: core.StringPtr("enable"),
			}

			agentKpiDataModel := &schematicsv1.AgentKPIData{
				AvailabilityIndicator: core.StringPtr("available"),
				LifecycleIndicator: core.StringPtr("consistent"),
				PercentUsageIndicator: core.StringPtr("testString"),
				ApplicationIndicators: []interface{}{"testString"},
				InfraIndicators: []interface{}{"testString"},
			}

			updateAgentDataOptions := &schematicsv1.UpdateAgentDataOptions{
				AgentID: core.StringPtr("testString"),
				Name: core.StringPtr("MyDevAgent"),
				ResourceGroup: core.StringPtr("Default"),
				Version: core.StringPtr("v1.0.0"),
				SchematicsLocation: core.StringPtr("us-south"),
				AgentLocation: core.StringPtr("us-south"),
				AgentInfrastructure: agentInfrastructureModel,
				Description: core.StringPtr("Create Agent"),
				Tags: []string{"testString"},
				AgentMetadata: []schematicsv1.AgentMetadataInfo{*agentMetadataInfoModel},
				AgentInputs: []schematicsv1.VariableData{*variableDataModel},
				UserState: agentUserStateModel,
				AgentKpi: agentKpiDataModel,
			}

			agentData, response, err := schematicsService.UpdateAgentData(updateAgentDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agentData).ToNot(BeNil())
		})
	})

	Describe(`UpdateAgentRegistration - Update the agent registration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAgentRegistration(updateAgentRegistrationOptions *UpdateAgentRegistrationOptions)`, func() {
			agentUserStateModel := &schematicsv1.AgentUserState{
				State: core.StringPtr("enable"),
			}

			updateAgentRegistrationOptions := &schematicsv1.UpdateAgentRegistrationOptions{
				AgentID: core.StringPtr("testString"),
				Name: core.StringPtr("MyDevAgent"),
				AgentLocation: core.StringPtr("us-south"),
				Location: core.StringPtr("us-south"),
				ProfileID: core.StringPtr("testString"),
				Description: core.StringPtr("Register agent"),
				ResourceGroup: core.StringPtr("testString"),
				Tags: []string{"testString"},
				UserState: agentUserStateModel,
			}

			agent, response, err := schematicsService.UpdateAgentRegistration(updateAgentRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agent).ToNot(BeNil())
		})
	})

	Describe(`GetKmsSettings - Get a KMS settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetKmsSettings(getKmsSettingsOptions *GetKmsSettingsOptions)`, func() {
			getKmsSettingsOptions := &schematicsv1.GetKmsSettingsOptions{
				Location: core.StringPtr("testString"),
			}

			kmsSettings, response, err := schematicsService.GetKmsSettings(getKmsSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsSettings).ToNot(BeNil())
		})
	})

	Describe(`ListKms - List KMS instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListKms(listKmsOptions *ListKmsOptions)`, func() {
			listKmsOptions := &schematicsv1.ListKmsOptions{
				EncryptionScheme: core.StringPtr("testString"),
				Location: core.StringPtr("testString"),
				ResourceGroup: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(100)),
				Sort: core.StringPtr("testString"),
			}

			kmsDiscovery, response, err := schematicsService.ListKms(listKmsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsDiscovery).ToNot(BeNil())
		})
	})

	Describe(`UpdateKmsSettings - Update a KMS settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateKmsSettings(updateKmsSettingsOptions *UpdateKmsSettingsOptions)`, func() {
			kmsSettingsPrimaryCrkModel := &schematicsv1.KMSSettingsPrimaryCrk{
				KmsName: core.StringPtr("testString"),
				KmsPrivateEndpoint: core.StringPtr("testString"),
				KeyCrn: core.StringPtr("testString"),
			}

			kmsSettingsSecondaryCrkModel := &schematicsv1.KMSSettingsSecondaryCrk{
				KmsName: core.StringPtr("testString"),
				KmsPrivateEndpoint: core.StringPtr("testString"),
				KeyCrn: core.StringPtr("testString"),
			}

			updateKmsSettingsOptions := &schematicsv1.UpdateKmsSettingsOptions{
				Location: core.StringPtr("testString"),
				EncryptionScheme: core.StringPtr("testString"),
				ResourceGroup: core.StringPtr("testString"),
				PrimaryCrk: kmsSettingsPrimaryCrkModel,
				SecondaryCrk: kmsSettingsSecondaryCrkModel,
			}

			kmsSettings, response, err := schematicsService.UpdateKmsSettings(updateKmsSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsSettings).ToNot(BeNil())
		})
	})

	Describe(`CreatePolicy - Create a policy account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicy(createPolicyOptions *CreatePolicyOptions)`, func() {
			userStateModel := &schematicsv1.UserState{
				State: core.StringPtr("draft"),
				SetBy: core.StringPtr("testString"),
				SetAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			policyObjectSelectorModel := &schematicsv1.PolicyObjectSelector{
				Kind: core.StringPtr("workspace"),
				Tags: []string{"testString"},
				ResourceGroups: []string{"testString"},
				Locations: []string{"us-south"},
			}

			policyObjectsModel := &schematicsv1.PolicyObjects{
				SelectorKind: core.StringPtr("ids"),
				SelectorIds: []string{"testString"},
				SelectorScope: []schematicsv1.PolicyObjectSelector{*policyObjectSelectorModel},
			}

			agentAssignmentPolicyParameterModel := &schematicsv1.AgentAssignmentPolicyParameter{
				SelectorKind: core.StringPtr("ids"),
				SelectorIds: []string{"testString"},
				SelectorScope: []schematicsv1.PolicyObjectSelector{*policyObjectSelectorModel},
			}

			policyParameterModel := &schematicsv1.PolicyParameter{
				AgentAssignmentPolicyParameter: agentAssignmentPolicyParameterModel,
			}

			scopedResourceModel := &schematicsv1.ScopedResource{
				Kind: core.StringPtr("workspace"),
				ID: core.StringPtr("testString"),
			}

			createPolicyOptions := &schematicsv1.CreatePolicyOptions{
				Name: core.StringPtr("Agent1-DevWS"),
				Description: core.StringPtr("Policy for job execution of secured workspaces on agent1"),
				ResourceGroup: core.StringPtr("Default"),
				Tags: []string{"policy:secured-job"},
				Location: core.StringPtr("us-south"),
				State: userStateModel,
				Kind: core.StringPtr("agent_assignment_policy"),
				Target: policyObjectsModel,
				Parameter: policyParameterModel,
				ScopedResources: []schematicsv1.ScopedResource{*scopedResourceModel},
			}

			policy, response, err := schematicsService.CreatePolicy(createPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
		})
	})

	Describe(`GetPolicy - Get policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicy(getPolicyOptions *GetPolicyOptions)`, func() {
			getPolicyOptions := &schematicsv1.GetPolicyOptions{
				PolicyID: core.StringPtr("testString"),
				Profile: core.StringPtr("summary"),
			}

			policy, response, err := schematicsService.GetPolicy(getPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
	})

	Describe(`ListPolicy - List policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicy(listPolicyOptions *ListPolicyOptions)`, func() {
			listPolicyOptions := &schematicsv1.ListPolicyOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit: core.Int64Ptr(int64(100)),
				Profile: core.StringPtr("summary"),
			}

			policyList, response, err := schematicsService.ListPolicy(listPolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyList).ToNot(BeNil())
		})
	})

	Describe(`UpdatePolicy - Update policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions)`, func() {
			userStateModel := &schematicsv1.UserState{
				State: core.StringPtr("draft"),
				SetBy: core.StringPtr("testString"),
				SetAt: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			policyObjectSelectorModel := &schematicsv1.PolicyObjectSelector{
				Kind: core.StringPtr("workspace"),
				Tags: []string{"testString"},
				ResourceGroups: []string{"testString"},
				Locations: []string{"us-south"},
			}

			policyObjectsModel := &schematicsv1.PolicyObjects{
				SelectorKind: core.StringPtr("ids"),
				SelectorIds: []string{"testString"},
				SelectorScope: []schematicsv1.PolicyObjectSelector{*policyObjectSelectorModel},
			}

			agentAssignmentPolicyParameterModel := &schematicsv1.AgentAssignmentPolicyParameter{
				SelectorKind: core.StringPtr("ids"),
				SelectorIds: []string{"testString"},
				SelectorScope: []schematicsv1.PolicyObjectSelector{*policyObjectSelectorModel},
			}

			policyParameterModel := &schematicsv1.PolicyParameter{
				AgentAssignmentPolicyParameter: agentAssignmentPolicyParameterModel,
			}

			scopedResourceModel := &schematicsv1.ScopedResource{
				Kind: core.StringPtr("workspace"),
				ID: core.StringPtr("testString"),
			}

			updatePolicyOptions := &schematicsv1.UpdatePolicyOptions{
				PolicyID: core.StringPtr("testString"),
				Name: core.StringPtr("Agent1-DevWS"),
				Description: core.StringPtr("Policy for job execution of secured workspaces on agent1"),
				ResourceGroup: core.StringPtr("Default"),
				Tags: []string{"policy:secured-job"},
				Location: core.StringPtr("us-south"),
				State: userStateModel,
				Kind: core.StringPtr("agent_assignment_policy"),
				Target: policyObjectsModel,
				Parameter: policyParameterModel,
				ScopedResources: []schematicsv1.ScopedResource{*scopedResourceModel},
			}

			policy, response, err := schematicsService.UpdatePolicy(updatePolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
	})

	Describe(`DeleteWorkspace - Delete a workspace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions)`, func() {
			deleteWorkspaceOptions := &schematicsv1.DeleteWorkspaceOptions{
				RefreshToken: core.StringPtr("testString"),
				WID: core.StringPtr("testString"),
				DestroyResources: core.StringPtr("testString"),
			}

			result, response, err := schematicsService.DeleteWorkspace(deleteWorkspaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`DeleteAction - Delete an action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAction(deleteActionOptions *DeleteActionOptions)`, func() {
			deleteActionOptions := &schematicsv1.DeleteActionOptions{
				ActionID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
				Propagate: core.BoolPtr(true),
			}

			response, err := schematicsService.DeleteAction(deleteActionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteJob - Stop the running Job, and delete the Job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteJob(deleteJobOptions *DeleteJobOptions)`, func() {
			deleteJobOptions := &schematicsv1.DeleteJobOptions{
				JobID: core.StringPtr("testString"),
				RefreshToken: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
				Propagate: core.BoolPtr(true),
			}

			response, err := schematicsService.DeleteJob(deleteJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteWorkspaceActivity - Stop the workspace job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteWorkspaceActivity(deleteWorkspaceActivityOptions *DeleteWorkspaceActivityOptions)`, func() {
			deleteWorkspaceActivityOptions := &schematicsv1.DeleteWorkspaceActivityOptions{
				WID: core.StringPtr("testString"),
				ActivityID: core.StringPtr("testString"),
			}

			workspaceActivityApplyResult, response, err := schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityApplyResult).ToNot(BeNil())
		})
	})

	Describe(`DeleteBlueprint - Delete a blueprint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBlueprint(deleteBlueprintOptions *DeleteBlueprintOptions)`, func() {
			deleteBlueprintOptions := &schematicsv1.DeleteBlueprintOptions{
				BlueprintID: core.StringPtr("testString"),
				Profile: core.StringPtr("ids"),
				Destroy: core.BoolPtr(true),
			}

			response, err := schematicsService.DeleteBlueprint(deleteBlueprintOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteInventory - Delete an inventory definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInventory(deleteInventoryOptions *DeleteInventoryOptions)`, func() {
			deleteInventoryOptions := &schematicsv1.DeleteInventoryOptions{
				InventoryID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
				Propagate: core.BoolPtr(true),
			}

			response, err := schematicsService.DeleteInventory(deleteInventoryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteResourcesQuery - Delete resources query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteResourcesQuery(deleteResourcesQueryOptions *DeleteResourcesQueryOptions)`, func() {
			deleteResourcesQueryOptions := &schematicsv1.DeleteResourcesQueryOptions{
				QueryID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
				Propagate: core.BoolPtr(true),
			}

			response, err := schematicsService.DeleteResourcesQuery(deleteResourcesQueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteAgent - Deregister the agent`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAgent(deleteAgentOptions *DeleteAgentOptions)`, func() {
			deleteAgentOptions := &schematicsv1.DeleteAgentOptions{
				AgentID: core.StringPtr("testString"),
			}

			response, err := schematicsService.DeleteAgent(deleteAgentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteAgentData - Delete agent`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAgentData(deleteAgentDataOptions *DeleteAgentDataOptions)`, func() {
			deleteAgentDataOptions := &schematicsv1.DeleteAgentDataOptions{
				AgentID: core.StringPtr("testString"),
			}

			response, err := schematicsService.DeleteAgentData(deleteAgentDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicy - Delete policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicy(deletePolicyOptions *DeletePolicyOptions)`, func() {
			deletePolicyOptions := &schematicsv1.DeletePolicyOptions{
				PolicyID: core.StringPtr("testString"),
			}

			response, err := schematicsService.DeletePolicy(deletePolicyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
