//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the schematics service.
//
// The following configuration properties are assumed to be defined:
// SCHEMATICS_URL=<service base url>
// SCHEMATICS_AUTH_TYPE=iam
// SCHEMATICS_APIKEY=<IAM apikey>
// SCHEMATICS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`SchematicsV1 Examples Tests`, func() {

	const externalConfigFile = "../schematics_v1.env"

	var (
		schematicsService *schematicsv1.SchematicsV1
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(schematicsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

			schematicsService, err = schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(schematicsService).ToNot(BeNil())
		})
	})

	Describe(`SchematicsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSchematicsLocation request example`, func() {
			fmt.Println("\nListSchematicsLocation() result:")
			// begin-list_schematics_location

			listSchematicsLocationOptions := schematicsService.NewListSchematicsLocationOptions()

			schematicsLocations, response, err := schematicsService.ListSchematicsLocation(listSchematicsLocationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(schematicsLocations, "", "  ")
			fmt.Println(string(b))

			// end-list_schematics_location

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsLocations).ToNot(BeNil())
		})
		It(`ListLocations request example`, func() {
			fmt.Println("\nListLocations() result:")
			// begin-list_locations

			listLocationsOptions := schematicsService.NewListLocationsOptions()

			schematicsLocationsList, response, err := schematicsService.ListLocations(listLocationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(schematicsLocationsList, "", "  ")
			fmt.Println(string(b))

			// end-list_locations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsLocationsList).ToNot(BeNil())
		})
		It(`ListResourceGroup request example`, func() {
			fmt.Println("\nListResourceGroup() result:")
			// begin-list_resource_group

			listResourceGroupOptions := schematicsService.NewListResourceGroupOptions()

			resourceGroupResponse, response, err := schematicsService.ListResourceGroup(listResourceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceGroupResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupResponse).ToNot(BeNil())
		})
		It(`GetSchematicsVersion request example`, func() {
			fmt.Println("\nGetSchematicsVersion() result:")
			// begin-get_schematics_version

			getSchematicsVersionOptions := schematicsService.NewGetSchematicsVersionOptions()

			versionResponse, response, err := schematicsService.GetSchematicsVersion(getSchematicsVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(versionResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_schematics_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionResponse).ToNot(BeNil())
		})
		It(`ProcessTemplateMetaData request example`, func() {
			fmt.Println("\nProcessTemplateMetaData() result:")
			// begin-ProcessTemplateMetaData

			externalSourceModel := &schematicsv1.ExternalSource{
				SourceType: core.StringPtr("local"),
			}

			processTemplateMetaDataOptions := schematicsService.NewProcessTemplateMetaDataOptions(
				"testString",
				externalSourceModel,
			)

			templateMetaDataResponse, response, err := schematicsService.ProcessTemplateMetaData(processTemplateMetaDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateMetaDataResponse, "", "  ")
			fmt.Println(string(b))

			// end-ProcessTemplateMetaData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateMetaDataResponse).ToNot(BeNil())
		})
		It(`ListWorkspaces request example`, func() {
			fmt.Println("\nListWorkspaces() result:")
			// begin-list_workspaces

			listWorkspacesOptions := schematicsService.NewListWorkspacesOptions()

			workspaceResponseList, response, err := schematicsService.ListWorkspaces(listWorkspacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceResponseList, "", "  ")
			fmt.Println(string(b))

			// end-list_workspaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponseList).ToNot(BeNil())
		})
		It(`CreateWorkspace request example`, func() {
			fmt.Println("\nCreateWorkspace() result:")
			// begin-create_workspace

			createWorkspaceOptions := schematicsService.NewCreateWorkspaceOptions()

			workspaceResponse, response, err := schematicsService.CreateWorkspace(createWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(workspaceResponse).ToNot(BeNil())
		})
		It(`GetWorkspace request example`, func() {
			fmt.Println("\nGetWorkspace() result:")
			// begin-get_workspace

			getWorkspaceOptions := schematicsService.NewGetWorkspaceOptions(
				"testString",
			)

			workspaceResponse, response, err := schematicsService.GetWorkspace(getWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())
		})
		It(`ReplaceWorkspace request example`, func() {
			fmt.Println("\nReplaceWorkspace() result:")
			// begin-replace_workspace

			replaceWorkspaceOptions := schematicsService.NewReplaceWorkspaceOptions(
				"testString",
			)

			workspaceResponse, response, err := schematicsService.ReplaceWorkspace(replaceWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())
		})
		It(`UpdateWorkspace request example`, func() {
			fmt.Println("\nUpdateWorkspace() result:")
			// begin-update_workspace

			updateWorkspaceOptions := schematicsService.NewUpdateWorkspaceOptions(
				"testString",
			)

			workspaceResponse, response, err := schematicsService.UpdateWorkspace(updateWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())
		})
		It(`GetWorkspaceReadme request example`, func() {
			fmt.Println("\nGetWorkspaceReadme() result:")
			// begin-get_workspace_readme

			getWorkspaceReadmeOptions := schematicsService.NewGetWorkspaceReadmeOptions(
				"testString",
			)

			templateReadme, response, err := schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateReadme, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_readme

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateReadme).ToNot(BeNil())
		})
		It(`TemplateRepoUpload request example`, func() {
			fmt.Println("\nTemplateRepoUpload() result:")
			// begin-template_repo_upload

			templateRepoUploadOptions := schematicsService.NewTemplateRepoUploadOptions(
				"testString",
				"testString",
			)

			templateRepoTarUploadResponse, response, err := schematicsService.TemplateRepoUpload(templateRepoUploadOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateRepoTarUploadResponse, "", "  ")
			fmt.Println(string(b))

			// end-template_repo_upload

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateRepoTarUploadResponse).ToNot(BeNil())
		})
		It(`GetWorkspaceInputs request example`, func() {
			fmt.Println("\nGetWorkspaceInputs() result:")
			// begin-get_workspace_inputs

			getWorkspaceInputsOptions := schematicsService.NewGetWorkspaceInputsOptions(
				"testString",
				"testString",
			)

			templateValues, response, err := schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateValues, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_inputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateValues).ToNot(BeNil())
		})
		It(`ReplaceWorkspaceInputs request example`, func() {
			fmt.Println("\nReplaceWorkspaceInputs() result:")
			// begin-replace_workspace_inputs

			replaceWorkspaceInputsOptions := schematicsService.NewReplaceWorkspaceInputsOptions(
				"testString",
				"testString",
			)

			userValues, response, err := schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userValues, "", "  ")
			fmt.Println(string(b))

			// end-replace_workspace_inputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userValues).ToNot(BeNil())
		})
		It(`GetAllWorkspaceInputs request example`, func() {
			fmt.Println("\nGetAllWorkspaceInputs() result:")
			// begin-get_all_workspace_inputs

			getAllWorkspaceInputsOptions := schematicsService.NewGetAllWorkspaceInputsOptions(
				"testString",
			)

			workspaceTemplateValuesResponse, response, err := schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceTemplateValuesResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_all_workspace_inputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceTemplateValuesResponse).ToNot(BeNil())
		})
		It(`GetWorkspaceInputMetadata request example`, func() {
			fmt.Println("\nGetWorkspaceInputMetadata() result:")
			// begin-get_workspace_input_metadata

			getWorkspaceInputMetadataOptions := schematicsService.NewGetWorkspaceInputMetadataOptions(
				"testString",
				"testString",
			)

			templateMetadata, response, err := schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateMetadata, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_input_metadata

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateMetadata).ToNot(BeNil())
		})
		It(`GetWorkspaceOutputs request example`, func() {
			fmt.Println("\nGetWorkspaceOutputs() result:")
			// begin-get_workspace_outputs

			getWorkspaceOutputsOptions := schematicsService.NewGetWorkspaceOutputsOptions(
				"testString",
			)

			outputValuesInner, response, err := schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outputValuesInner, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_outputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outputValuesInner).ToNot(BeNil())
		})
		It(`GetWorkspaceResources request example`, func() {
			fmt.Println("\nGetWorkspaceResources() result:")
			// begin-get_workspace_resources

			getWorkspaceResourcesOptions := schematicsService.NewGetWorkspaceResourcesOptions(
				"testString",
			)

			templateResources, response, err := schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateResources, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResources).ToNot(BeNil())
		})
		It(`GetWorkspaceState request example`, func() {
			fmt.Println("\nGetWorkspaceState() result:")
			// begin-get_workspace_state

			getWorkspaceStateOptions := schematicsService.NewGetWorkspaceStateOptions(
				"testString",
			)

			stateStoreResponseList, response, err := schematicsService.GetWorkspaceState(getWorkspaceStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(stateStoreResponseList, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stateStoreResponseList).ToNot(BeNil())
		})
		It(`GetWorkspaceTemplateState request example`, func() {
			fmt.Println("\nGetWorkspaceTemplateState() result:")
			// begin-get_workspace_template_state

			getWorkspaceTemplateStateOptions := schematicsService.NewGetWorkspaceTemplateStateOptions(
				"testString",
				"testString",
			)

			templateStateStore, response, err := schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateStateStore, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_template_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateStateStore).ToNot(BeNil())
		})
		It(`GetWorkspaceActivityLogs request example`, func() {
			fmt.Println("\nGetWorkspaceActivityLogs() result:")
			// begin-get_workspace_activity_logs

			getWorkspaceActivityLogsOptions := schematicsService.NewGetWorkspaceActivityLogsOptions(
				"testString",
				"testString",
			)

			workspaceActivityLogs, response, err := schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityLogs, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_activity_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivityLogs).ToNot(BeNil())
		})
		It(`GetWorkspaceLogUrls request example`, func() {
			fmt.Println("\nGetWorkspaceLogUrls() result:")
			// begin-get_workspace_log_urls

			getWorkspaceLogUrlsOptions := schematicsService.NewGetWorkspaceLogUrlsOptions(
				"testString",
			)

			logStoreResponseList, response, err := schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logStoreResponseList, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_log_urls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logStoreResponseList).ToNot(BeNil())
		})
		It(`GetTemplateLogs request example`, func() {
			fmt.Println("\nGetTemplateLogs() result:")
			// begin-get_template_logs

			getTemplateLogsOptions := schematicsService.NewGetTemplateLogsOptions(
				"testString",
				"testString",
			)

			templateLogStoreString, response, err := schematicsService.GetTemplateLogs(getTemplateLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateLogStoreString, "", "  ")
			fmt.Println(string(b))

			// end-get_template_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateLogStoreString).ToNot(BeNil())
		})
		It(`GetTemplateActivityLog request example`, func() {
			fmt.Println("\nGetTemplateActivityLog() result:")
			// begin-get_template_activity_log

			getTemplateActivityLogOptions := schematicsService.NewGetTemplateActivityLogOptions(
				"testString",
				"testString",
				"testString",
			)

			workspaceActivityTemplateLogString, response, err := schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityTemplateLogString, "", "  ")
			fmt.Println(string(b))

			// end-get_template_activity_log

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivityTemplateLogString).ToNot(BeNil())
		})
		It(`ListActions request example`, func() {
			fmt.Println("\nListActions() result:")
			// begin-list_actions

			listActionsOptions := schematicsService.NewListActionsOptions()

			actionList, response, err := schematicsService.ListActions(listActionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionList, "", "  ")
			fmt.Println(string(b))

			// end-list_actions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionList).ToNot(BeNil())
		})
		It(`CreateAction request example`, func() {
			fmt.Println("\nCreateAction() result:")
			// begin-create_action

			createActionOptions := schematicsService.NewCreateActionOptions()

			action, response, err := schematicsService.CreateAction(createActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(action, "", "  ")
			fmt.Println(string(b))

			// end-create_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(action).ToNot(BeNil())
		})
		It(`GetAction request example`, func() {
			fmt.Println("\nGetAction() result:")
			// begin-get_action

			getActionOptions := schematicsService.NewGetActionOptions(
				"testString",
			)

			action, response, err := schematicsService.GetAction(getActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(action, "", "  ")
			fmt.Println(string(b))

			// end-get_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(action).ToNot(BeNil())
		})
		It(`UpdateAction request example`, func() {
			fmt.Println("\nUpdateAction() result:")
			// begin-update_action

			updateActionOptions := schematicsService.NewUpdateActionOptions(
				"testString",
			)

			action, response, err := schematicsService.UpdateAction(updateActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(action, "", "  ")
			fmt.Println(string(b))

			// end-update_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(action).ToNot(BeNil())
		})
		It(`UploadTemplateTarAction request example`, func() {
			fmt.Println("\nUploadTemplateTarAction() result:")
			// begin-upload_template_tar_action

			uploadTemplateTarActionOptions := schematicsService.NewUploadTemplateTarActionOptions(
				"testString",
			)

			templateRepoTarUploadResponse, response, err := schematicsService.UploadTemplateTarAction(uploadTemplateTarActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateRepoTarUploadResponse, "", "  ")
			fmt.Println(string(b))

			// end-upload_template_tar_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateRepoTarUploadResponse).ToNot(BeNil())
		})
		It(`ListWorkspaceActivities request example`, func() {
			fmt.Println("\nListWorkspaceActivities() result:")
			// begin-list_workspace_activities

			listWorkspaceActivitiesOptions := schematicsService.NewListWorkspaceActivitiesOptions(
				"testString",
			)

			workspaceActivities, response, err := schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivities, "", "  ")
			fmt.Println(string(b))

			// end-list_workspace_activities

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivities).ToNot(BeNil())
		})
		It(`GetWorkspaceActivity request example`, func() {
			fmt.Println("\nGetWorkspaceActivity() result:")
			// begin-get_workspace_activity

			getWorkspaceActivityOptions := schematicsService.NewGetWorkspaceActivityOptions(
				"testString",
				"testString",
			)

			workspaceActivity, response, err := schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivity, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_activity

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivity).ToNot(BeNil())
		})
		It(`RunWorkspaceCommands request example`, func() {
			fmt.Println("\nRunWorkspaceCommands() result:")
			// begin-run_workspace_commands

			runWorkspaceCommandsOptions := schematicsService.NewRunWorkspaceCommandsOptions(
				"testString",
				"testString",
			)

			workspaceActivityCommandResult, response, err := schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityCommandResult, "", "  ")
			fmt.Println(string(b))

			// end-run_workspace_commands

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivityCommandResult).ToNot(BeNil())
		})
		It(`ApplyWorkspaceCommand request example`, func() {
			fmt.Println("\nApplyWorkspaceCommand() result:")
			// begin-apply_workspace_command

			applyWorkspaceCommandOptions := schematicsService.NewApplyWorkspaceCommandOptions(
				"testString",
				"testString",
			)

			workspaceActivityApplyResult, response, err := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityApplyResult, "", "  ")
			fmt.Println(string(b))

			// end-apply_workspace_command

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityApplyResult).ToNot(BeNil())
		})
		It(`DestroyWorkspaceCommand request example`, func() {
			fmt.Println("\nDestroyWorkspaceCommand() result:")
			// begin-destroy_workspace_command

			destroyWorkspaceCommandOptions := schematicsService.NewDestroyWorkspaceCommandOptions(
				"testString",
				"testString",
			)

			workspaceActivityDestroyResult, response, err := schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityDestroyResult, "", "  ")
			fmt.Println(string(b))

			// end-destroy_workspace_command

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityDestroyResult).ToNot(BeNil())
		})
		It(`PlanWorkspaceCommand request example`, func() {
			fmt.Println("\nPlanWorkspaceCommand() result:")
			// begin-plan_workspace_command

			planWorkspaceCommandOptions := schematicsService.NewPlanWorkspaceCommandOptions(
				"testString",
				"testString",
			)

			workspaceActivityPlanResult, response, err := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityPlanResult, "", "  ")
			fmt.Println(string(b))

			// end-plan_workspace_command

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityPlanResult).ToNot(BeNil())
		})
		It(`RefreshWorkspaceCommand request example`, func() {
			fmt.Println("\nRefreshWorkspaceCommand() result:")
			// begin-refresh_workspace_command

			refreshWorkspaceCommandOptions := schematicsService.NewRefreshWorkspaceCommandOptions(
				"testString",
				"testString",
			)

			workspaceActivityRefreshResult, response, err := schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityRefreshResult, "", "  ")
			fmt.Println(string(b))

			// end-refresh_workspace_command

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityRefreshResult).ToNot(BeNil())
		})
		It(`ListJobs request example`, func() {
			fmt.Println("\nListJobs() result:")
			// begin-list_jobs

			listJobsOptions := schematicsService.NewListJobsOptions()

			jobList, response, err := schematicsService.ListJobs(listJobsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobList, "", "  ")
			fmt.Println(string(b))

			// end-list_jobs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobList).ToNot(BeNil())
		})
		It(`CreateJob request example`, func() {
			fmt.Println("\nCreateJob() result:")
			// begin-create_job

			createJobOptions := schematicsService.NewCreateJobOptions(
				"testString",
			)

			job, response, err := schematicsService.CreateJob(createJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-create_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(job).ToNot(BeNil())
		})
		It(`GetJob request example`, func() {
			fmt.Println("\nGetJob() result:")
			// begin-get_job

			getJobOptions := schematicsService.NewGetJobOptions(
				"testString",
			)

			job, response, err := schematicsService.GetJob(getJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-get_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
		It(`UpdateJob request example`, func() {
			fmt.Println("\nUpdateJob() result:")
			// begin-update_job

			updateJobOptions := schematicsService.NewUpdateJobOptions(
				"testString",
				"testString",
			)

			job, response, err := schematicsService.UpdateJob(updateJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-update_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
		It(`ListJobLogs request example`, func() {
			fmt.Println("\nListJobLogs() result:")
			// begin-list_job_logs

			listJobLogsOptions := schematicsService.NewListJobLogsOptions(
				"testString",
			)

			jobLog, response, err := schematicsService.ListJobLogs(listJobLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobLog, "", "  ")
			fmt.Println(string(b))

			// end-list_job_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobLog).ToNot(BeNil())
		})
		It(`GetJobFiles request example`, func() {
			fmt.Println("\nGetJobFiles() result:")
			// begin-get_job_files

			getJobFilesOptions := schematicsService.NewGetJobFilesOptions(
				"testString",
				"template_repo",
			)

			jobFileData, response, err := schematicsService.GetJobFiles(getJobFilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobFileData, "", "  ")
			fmt.Println(string(b))

			// end-get_job_files

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobFileData).ToNot(BeNil())
		})
		It(`CreateWorkspaceDeletionJob request example`, func() {
			fmt.Println("\nCreateWorkspaceDeletionJob() result:")
			// begin-create_workspace_deletion_job

			createWorkspaceDeletionJobOptions := schematicsService.NewCreateWorkspaceDeletionJobOptions(
				"testString",
			)

			workspaceBulkDeleteResponse, response, err := schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceBulkDeleteResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_workspace_deletion_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceBulkDeleteResponse).ToNot(BeNil())
		})
		It(`GetWorkspaceDeletionJobStatus request example`, func() {
			fmt.Println("\nGetWorkspaceDeletionJobStatus() result:")
			// begin-get_workspace_deletion_job_status

			getWorkspaceDeletionJobStatusOptions := schematicsService.NewGetWorkspaceDeletionJobStatusOptions(
				"testString",
			)

			workspaceJobResponse, response, err := schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceJobResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_deletion_job_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceJobResponse).ToNot(BeNil())
		})
		It(`ListInventories request example`, func() {
			fmt.Println("\nListInventories() result:")
			// begin-list_inventories

			listInventoriesOptions := schematicsService.NewListInventoriesOptions()

			inventoryResourceRecordList, response, err := schematicsService.ListInventories(listInventoriesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecordList, "", "  ")
			fmt.Println(string(b))

			// end-list_inventories

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecordList).ToNot(BeNil())
		})
		It(`CreateInventory request example`, func() {
			fmt.Println("\nCreateInventory() result:")
			// begin-create_inventory

			createInventoryOptions := schematicsService.NewCreateInventoryOptions()

			inventoryResourceRecord, response, err := schematicsService.CreateInventory(createInventoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecord, "", "  ")
			fmt.Println(string(b))

			// end-create_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())
		})
		It(`GetInventory request example`, func() {
			fmt.Println("\nGetInventory() result:")
			// begin-get_inventory

			getInventoryOptions := schematicsService.NewGetInventoryOptions(
				"testString",
			)

			inventoryResourceRecord, response, err := schematicsService.GetInventory(getInventoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecord, "", "  ")
			fmt.Println(string(b))

			// end-get_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())
		})
		It(`ReplaceInventory request example`, func() {
			fmt.Println("\nReplaceInventory() result:")
			// begin-replace_inventory

			replaceInventoryOptions := schematicsService.NewReplaceInventoryOptions(
				"testString",
			)

			inventoryResourceRecord, response, err := schematicsService.ReplaceInventory(replaceInventoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecord, "", "  ")
			fmt.Println(string(b))

			// end-replace_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())
		})
		It(`ListResourceQuery request example`, func() {
			fmt.Println("\nListResourceQuery() result:")
			// begin-list_resource_query

			listResourceQueryOptions := schematicsService.NewListResourceQueryOptions()

			resourceQueryRecordList, response, err := schematicsService.ListResourceQuery(listResourceQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecordList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecordList).ToNot(BeNil())
		})
		It(`CreateResourceQuery request example`, func() {
			fmt.Println("\nCreateResourceQuery() result:")
			// begin-create_resource_query

			createResourceQueryOptions := schematicsService.NewCreateResourceQueryOptions()

			resourceQueryRecord, response, err := schematicsService.CreateResourceQuery(createResourceQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecord, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())
		})
		It(`GetResourcesQuery request example`, func() {
			fmt.Println("\nGetResourcesQuery() result:")
			// begin-get_resources_query

			getResourcesQueryOptions := schematicsService.NewGetResourcesQueryOptions(
				"testString",
			)

			resourceQueryRecord, response, err := schematicsService.GetResourcesQuery(getResourcesQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecord, "", "  ")
			fmt.Println(string(b))

			// end-get_resources_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())
		})
		It(`ReplaceResourcesQuery request example`, func() {
			fmt.Println("\nReplaceResourcesQuery() result:")
			// begin-replace_resources_query

			replaceResourcesQueryOptions := schematicsService.NewReplaceResourcesQueryOptions(
				"testString",
			)

			resourceQueryRecord, response, err := schematicsService.ReplaceResourcesQuery(replaceResourcesQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecord, "", "  ")
			fmt.Println(string(b))

			// end-replace_resources_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())
		})
		It(`ExecuteResourceQuery request example`, func() {
			fmt.Println("\nExecuteResourceQuery() result:")
			// begin-execute_resource_query

			executeResourceQueryOptions := schematicsService.NewExecuteResourceQueryOptions(
				"testString",
			)

			resourceQueryResponseRecord, response, err := schematicsService.ExecuteResourceQuery(executeResourceQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryResponseRecord, "", "  ")
			fmt.Println(string(b))

			// end-execute_resource_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryResponseRecord).ToNot(BeNil())
		})
		It(`ListAgent request example`, func() {
			fmt.Println("\nListAgent() result:")
			// begin-list_agent

			listAgentOptions := schematicsService.NewListAgentOptions()

			agentList, response, err := schematicsService.ListAgent(listAgentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentList, "", "  ")
			fmt.Println(string(b))

			// end-list_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentList).ToNot(BeNil())
		})
		It(`RegisterAgent request example`, func() {
			fmt.Println("\nRegisterAgent() result:")
			// begin-register_agent

			registerAgentOptions := schematicsService.NewRegisterAgentOptions(
				"MyDevAgent",
				"us-south",
				"us-south",
				"testString",
			)

			agent, response, err := schematicsService.RegisterAgent(registerAgentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agent, "", "  ")
			fmt.Println(string(b))

			// end-register_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agent).ToNot(BeNil())
		})
		It(`GetAgent request example`, func() {
			fmt.Println("\nGetAgent() result:")
			// begin-get_agent

			getAgentOptions := schematicsService.NewGetAgentOptions(
				"testString",
			)

			agent, response, err := schematicsService.GetAgent(getAgentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agent, "", "  ")
			fmt.Println(string(b))

			// end-get_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agent).ToNot(BeNil())
		})
		It(`UpdateAgentRegistration request example`, func() {
			fmt.Println("\nUpdateAgentRegistration() result:")
			// begin-update_agent_registration

			updateAgentRegistrationOptions := schematicsService.NewUpdateAgentRegistrationOptions(
				"testString",
				"MyDevAgent",
				"us-south",
				"us-south",
				"testString",
			)

			agent, response, err := schematicsService.UpdateAgentRegistration(updateAgentRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agent, "", "  ")
			fmt.Println(string(b))

			// end-update_agent_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agent).ToNot(BeNil())
		})
		It(`ListAgentData request example`, func() {
			fmt.Println("\nListAgentData() result:")
			// begin-list_agent_data

			listAgentDataOptions := schematicsService.NewListAgentDataOptions()

			agentDataList, response, err := schematicsService.ListAgentData(listAgentDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentDataList, "", "  ")
			fmt.Println(string(b))

			// end-list_agent_data

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentDataList).ToNot(BeNil())
		})
		It(`CreateAgentData request example`, func() {
			fmt.Println("\nCreateAgentData() result:")
			// begin-create_agent_data

			agentInfrastructureModel := &schematicsv1.AgentInfrastructure{
			}

			createAgentDataOptions := schematicsService.NewCreateAgentDataOptions(
				"MyDevAgent",
				"Default",
				"v1.0.0",
				"us-south",
				"us-south",
				agentInfrastructureModel,
			)

			agentData, response, err := schematicsService.CreateAgentData(createAgentDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentData, "", "  ")
			fmt.Println(string(b))

			// end-create_agent_data

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agentData).ToNot(BeNil())
		})
		It(`GetAgentData request example`, func() {
			fmt.Println("\nGetAgentData() result:")
			// begin-get_agent_data

			getAgentDataOptions := schematicsService.NewGetAgentDataOptions(
				"testString",
			)

			agentData, response, err := schematicsService.GetAgentData(getAgentDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentData, "", "  ")
			fmt.Println(string(b))

			// end-get_agent_data

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentData).ToNot(BeNil())
		})
		It(`UpdateAgentData request example`, func() {
			fmt.Println("\nUpdateAgentData() result:")
			// begin-update_agent_data

			agentInfrastructureModel := &schematicsv1.AgentInfrastructure{
			}

			updateAgentDataOptions := schematicsService.NewUpdateAgentDataOptions(
				"testString",
				"MyDevAgent",
				"Default",
				"v1.0.0",
				"us-south",
				"us-south",
				agentInfrastructureModel,
			)

			agentData, response, err := schematicsService.UpdateAgentData(updateAgentDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentData, "", "  ")
			fmt.Println(string(b))

			// end-update_agent_data

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agentData).ToNot(BeNil())
		})
		It(`GetAgentVersions request example`, func() {
			fmt.Println("\nGetAgentVersions() result:")
			// begin-get_agent_versions

			getAgentVersionsOptions := schematicsService.NewGetAgentVersionsOptions()

			agentVersions, response, err := schematicsService.GetAgentVersions(getAgentVersionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentVersions, "", "  ")
			fmt.Println(string(b))

			// end-get_agent_versions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentVersions).ToNot(BeNil())
		})
		It(`GetPrsAgentJob request example`, func() {
			fmt.Println("\nGetPrsAgentJob() result:")
			// begin-get_prs_agent_job

			getPrsAgentJobOptions := schematicsService.NewGetPrsAgentJobOptions(
				"testString",
			)

			agentPrsJob, response, err := schematicsService.GetPrsAgentJob(getPrsAgentJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentPrsJob, "", "  ")
			fmt.Println(string(b))

			// end-get_prs_agent_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentPrsJob).ToNot(BeNil())
		})
		It(`PrsAgentJob request example`, func() {
			fmt.Println("\nPrsAgentJob() result:")
			// begin-prs_agent_job

			prsAgentJobOptions := schematicsService.NewPrsAgentJobOptions(
				"testString",
			)

			agentPrsJob, response, err := schematicsService.PrsAgentJob(prsAgentJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentPrsJob, "", "  ")
			fmt.Println(string(b))

			// end-prs_agent_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentPrsJob).ToNot(BeNil())
		})
		It(`GetHealthCheckAgentJob request example`, func() {
			fmt.Println("\nGetHealthCheckAgentJob() result:")
			// begin-get_health_check_agent_job

			getHealthCheckAgentJobOptions := schematicsService.NewGetHealthCheckAgentJobOptions(
				"testString",
			)

			agentHealthJob, response, err := schematicsService.GetHealthCheckAgentJob(getHealthCheckAgentJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentHealthJob, "", "  ")
			fmt.Println(string(b))

			// end-get_health_check_agent_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentHealthJob).ToNot(BeNil())
		})
		It(`HealthCheckAgentJob request example`, func() {
			fmt.Println("\nHealthCheckAgentJob() result:")
			// begin-health_check_agent_job

			healthCheckAgentJobOptions := schematicsService.NewHealthCheckAgentJobOptions(
				"testString",
			)

			agentHealthJob, response, err := schematicsService.HealthCheckAgentJob(healthCheckAgentJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentHealthJob, "", "  ")
			fmt.Println(string(b))

			// end-health_check_agent_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentHealthJob).ToNot(BeNil())
		})
		It(`GetDeployAgentJob request example`, func() {
			fmt.Println("\nGetDeployAgentJob() result:")
			// begin-get_deploy_agent_job

			getDeployAgentJobOptions := schematicsService.NewGetDeployAgentJobOptions(
				"testString",
			)

			agentDeployJob, response, err := schematicsService.GetDeployAgentJob(getDeployAgentJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentDeployJob, "", "  ")
			fmt.Println(string(b))

			// end-get_deploy_agent_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentDeployJob).ToNot(BeNil())
		})
		It(`DeployAgentJob request example`, func() {
			fmt.Println("\nDeployAgentJob() result:")
			// begin-deploy_agent_job

			deployAgentJobOptions := schematicsService.NewDeployAgentJobOptions(
				"testString",
			)

			agentDeployJob, response, err := schematicsService.DeployAgentJob(deployAgentJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentDeployJob, "", "  ")
			fmt.Println(string(b))

			// end-deploy_agent_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(agentDeployJob).ToNot(BeNil())
		})
		It(`GetKmsSettings request example`, func() {
			fmt.Println("\nGetKmsSettings() result:")
			// begin-get_kms_settings

			getKmsSettingsOptions := schematicsService.NewGetKmsSettingsOptions(
				"testString",
			)

			kmsSettings, response, err := schematicsService.GetKmsSettings(getKmsSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(kmsSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_kms_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsSettings).ToNot(BeNil())
		})
		It(`UpdateKmsSettings request example`, func() {
			fmt.Println("\nUpdateKmsSettings() result:")
			// begin-update_kms_settings

			updateKmsSettingsOptions := schematicsService.NewUpdateKmsSettingsOptions()

			kmsSettings, response, err := schematicsService.UpdateKmsSettings(updateKmsSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(kmsSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_kms_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsSettings).ToNot(BeNil())
		})
		It(`ListKms request example`, func() {
			fmt.Println("\nListKms() result:")
			// begin-list_kms

			listKmsOptions := schematicsService.NewListKmsOptions(
				"testString",
				"testString",
			)

			kmsDiscovery, response, err := schematicsService.ListKms(listKmsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(kmsDiscovery, "", "  ")
			fmt.Println(string(b))

			// end-list_kms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsDiscovery).ToNot(BeNil())
		})
		It(`ListPolicy request example`, func() {
			fmt.Println("\nListPolicy() result:")
			// begin-list_policy

			listPolicyOptions := schematicsService.NewListPolicyOptions()

			policyList, response, err := schematicsService.ListPolicy(listPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyList, "", "  ")
			fmt.Println(string(b))

			// end-list_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyList).ToNot(BeNil())
		})
		It(`CreatePolicy request example`, func() {
			fmt.Println("\nCreatePolicy() result:")
			// begin-create_policy

			createPolicyOptions := schematicsService.NewCreatePolicyOptions()

			policy, response, err := schematicsService.CreatePolicy(createPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-create_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
		})
		It(`GetPolicy request example`, func() {
			fmt.Println("\nGetPolicy() result:")
			// begin-get_policy

			getPolicyOptions := schematicsService.NewGetPolicyOptions(
				"testString",
			)

			policy, response, err := schematicsService.GetPolicy(getPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-get_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`UpdatePolicy request example`, func() {
			fmt.Println("\nUpdatePolicy() result:")
			// begin-update_policy

			updatePolicyOptions := schematicsService.NewUpdatePolicyOptions(
				"testString",
			)

			policy, response, err := schematicsService.UpdatePolicy(updatePolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-update_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`DeleteWorkspace request example`, func() {
			fmt.Println("\nDeleteWorkspace() result:")
			// begin-delete_workspace

			deleteWorkspaceOptions := schematicsService.NewDeleteWorkspaceOptions(
				"testString",
				"testString",
			)

			workspaceDeleteResponse, response, err := schematicsService.DeleteWorkspace(deleteWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceDeleteResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceDeleteResponse).ToNot(BeNil())
		})
		It(`DeleteAction request example`, func() {
			// begin-delete_action

			deleteActionOptions := schematicsService.NewDeleteActionOptions(
				"testString",
			)

			response, err := schematicsService.DeleteAction(deleteActionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAction(): %d\n", response.StatusCode)
			}

			// end-delete_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteWorkspaceActivity request example`, func() {
			fmt.Println("\nDeleteWorkspaceActivity() result:")
			// begin-delete_workspace_activity

			deleteWorkspaceActivityOptions := schematicsService.NewDeleteWorkspaceActivityOptions(
				"testString",
				"testString",
			)

			workspaceActivityApplyResult, response, err := schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceActivityApplyResult, "", "  ")
			fmt.Println(string(b))

			// end-delete_workspace_activity

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityApplyResult).ToNot(BeNil())
		})
		It(`DeleteJob request example`, func() {
			// begin-delete_job

			deleteJobOptions := schematicsService.NewDeleteJobOptions(
				"testString",
				"testString",
			)

			response, err := schematicsService.DeleteJob(deleteJobOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteJob(): %d\n", response.StatusCode)
			}

			// end-delete_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteInventory request example`, func() {
			// begin-delete_inventory

			deleteInventoryOptions := schematicsService.NewDeleteInventoryOptions(
				"testString",
			)

			response, err := schematicsService.DeleteInventory(deleteInventoryOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteInventory(): %d\n", response.StatusCode)
			}

			// end-delete_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteResourcesQuery request example`, func() {
			// begin-delete_resources_query

			deleteResourcesQueryOptions := schematicsService.NewDeleteResourcesQueryOptions(
				"testString",
			)

			response, err := schematicsService.DeleteResourcesQuery(deleteResourcesQueryOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourcesQuery(): %d\n", response.StatusCode)
			}

			// end-delete_resources_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAgent request example`, func() {
			// begin-delete_agent

			deleteAgentOptions := schematicsService.NewDeleteAgentOptions(
				"testString",
			)

			response, err := schematicsService.DeleteAgent(deleteAgentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAgent(): %d\n", response.StatusCode)
			}

			// end-delete_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAgentData request example`, func() {
			// begin-delete_agent_data

			deleteAgentDataOptions := schematicsService.NewDeleteAgentDataOptions(
				"testString",
			)

			response, err := schematicsService.DeleteAgentData(deleteAgentDataOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAgentData(): %d\n", response.StatusCode)
			}

			// end-delete_agent_data

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAgentResources request example`, func() {
			// begin-delete_agent_resources

			deleteAgentResourcesOptions := schematicsService.NewDeleteAgentResourcesOptions(
				"testString",
				"testString",
			)

			response, err := schematicsService.DeleteAgentResources(deleteAgentResourcesOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteAgentResources(): %d\n", response.StatusCode)
			}

			// end-delete_agent_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeletePolicy request example`, func() {
			// begin-delete_policy

			deletePolicyOptions := schematicsService.NewDeletePolicyOptions(
				"testString",
			)

			response, err := schematicsService.DeletePolicy(deletePolicyOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeletePolicy(): %d\n", response.StatusCode)
			}

			// end-delete_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
