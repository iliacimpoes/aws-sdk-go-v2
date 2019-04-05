// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sagemakeriface provides an interface to enable mocking the Amazon SageMaker Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sagemakeriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

// SageMakerAPI provides an interface to enable mocking the
// sagemaker.SageMaker service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon SageMaker Service.
//    func myFunc(svc sagemakeriface.SageMakerAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sagemaker.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSageMakerClient struct {
//        sagemakeriface.SageMakerAPI
//    }
//    func (m *mockSageMakerClient) AddTags(input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSageMakerClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SageMakerAPI interface {
	AddTagsRequest(*sagemaker.AddTagsInput) sagemaker.AddTagsRequest

	CreateAlgorithmRequest(*sagemaker.CreateAlgorithmInput) sagemaker.CreateAlgorithmRequest

	CreateCodeRepositoryRequest(*sagemaker.CreateCodeRepositoryInput) sagemaker.CreateCodeRepositoryRequest

	CreateCompilationJobRequest(*sagemaker.CreateCompilationJobInput) sagemaker.CreateCompilationJobRequest

	CreateEndpointRequest(*sagemaker.CreateEndpointInput) sagemaker.CreateEndpointRequest

	CreateEndpointConfigRequest(*sagemaker.CreateEndpointConfigInput) sagemaker.CreateEndpointConfigRequest

	CreateHyperParameterTuningJobRequest(*sagemaker.CreateHyperParameterTuningJobInput) sagemaker.CreateHyperParameterTuningJobRequest

	CreateLabelingJobRequest(*sagemaker.CreateLabelingJobInput) sagemaker.CreateLabelingJobRequest

	CreateModelRequest(*sagemaker.CreateModelInput) sagemaker.CreateModelRequest

	CreateModelPackageRequest(*sagemaker.CreateModelPackageInput) sagemaker.CreateModelPackageRequest

	CreateNotebookInstanceRequest(*sagemaker.CreateNotebookInstanceInput) sagemaker.CreateNotebookInstanceRequest

	CreateNotebookInstanceLifecycleConfigRequest(*sagemaker.CreateNotebookInstanceLifecycleConfigInput) sagemaker.CreateNotebookInstanceLifecycleConfigRequest

	CreatePresignedNotebookInstanceUrlRequest(*sagemaker.CreatePresignedNotebookInstanceUrlInput) sagemaker.CreatePresignedNotebookInstanceUrlRequest

	CreateTrainingJobRequest(*sagemaker.CreateTrainingJobInput) sagemaker.CreateTrainingJobRequest

	CreateTransformJobRequest(*sagemaker.CreateTransformJobInput) sagemaker.CreateTransformJobRequest

	CreateWorkteamRequest(*sagemaker.CreateWorkteamInput) sagemaker.CreateWorkteamRequest

	DeleteAlgorithmRequest(*sagemaker.DeleteAlgorithmInput) sagemaker.DeleteAlgorithmRequest

	DeleteCodeRepositoryRequest(*sagemaker.DeleteCodeRepositoryInput) sagemaker.DeleteCodeRepositoryRequest

	DeleteEndpointRequest(*sagemaker.DeleteEndpointInput) sagemaker.DeleteEndpointRequest

	DeleteEndpointConfigRequest(*sagemaker.DeleteEndpointConfigInput) sagemaker.DeleteEndpointConfigRequest

	DeleteModelRequest(*sagemaker.DeleteModelInput) sagemaker.DeleteModelRequest

	DeleteModelPackageRequest(*sagemaker.DeleteModelPackageInput) sagemaker.DeleteModelPackageRequest

	DeleteNotebookInstanceRequest(*sagemaker.DeleteNotebookInstanceInput) sagemaker.DeleteNotebookInstanceRequest

	DeleteNotebookInstanceLifecycleConfigRequest(*sagemaker.DeleteNotebookInstanceLifecycleConfigInput) sagemaker.DeleteNotebookInstanceLifecycleConfigRequest

	DeleteTagsRequest(*sagemaker.DeleteTagsInput) sagemaker.DeleteTagsRequest

	DeleteWorkteamRequest(*sagemaker.DeleteWorkteamInput) sagemaker.DeleteWorkteamRequest

	DescribeAlgorithmRequest(*sagemaker.DescribeAlgorithmInput) sagemaker.DescribeAlgorithmRequest

	DescribeCodeRepositoryRequest(*sagemaker.DescribeCodeRepositoryInput) sagemaker.DescribeCodeRepositoryRequest

	DescribeCompilationJobRequest(*sagemaker.DescribeCompilationJobInput) sagemaker.DescribeCompilationJobRequest

	DescribeEndpointRequest(*sagemaker.DescribeEndpointInput) sagemaker.DescribeEndpointRequest

	DescribeEndpointConfigRequest(*sagemaker.DescribeEndpointConfigInput) sagemaker.DescribeEndpointConfigRequest

	DescribeHyperParameterTuningJobRequest(*sagemaker.DescribeHyperParameterTuningJobInput) sagemaker.DescribeHyperParameterTuningJobRequest

	DescribeLabelingJobRequest(*sagemaker.DescribeLabelingJobInput) sagemaker.DescribeLabelingJobRequest

	DescribeModelRequest(*sagemaker.DescribeModelInput) sagemaker.DescribeModelRequest

	DescribeModelPackageRequest(*sagemaker.DescribeModelPackageInput) sagemaker.DescribeModelPackageRequest

	DescribeNotebookInstanceRequest(*sagemaker.DescribeNotebookInstanceInput) sagemaker.DescribeNotebookInstanceRequest

	DescribeNotebookInstanceLifecycleConfigRequest(*sagemaker.DescribeNotebookInstanceLifecycleConfigInput) sagemaker.DescribeNotebookInstanceLifecycleConfigRequest

	DescribeSubscribedWorkteamRequest(*sagemaker.DescribeSubscribedWorkteamInput) sagemaker.DescribeSubscribedWorkteamRequest

	DescribeTrainingJobRequest(*sagemaker.DescribeTrainingJobInput) sagemaker.DescribeTrainingJobRequest

	DescribeTransformJobRequest(*sagemaker.DescribeTransformJobInput) sagemaker.DescribeTransformJobRequest

	DescribeWorkteamRequest(*sagemaker.DescribeWorkteamInput) sagemaker.DescribeWorkteamRequest

	GetSearchSuggestionsRequest(*sagemaker.GetSearchSuggestionsInput) sagemaker.GetSearchSuggestionsRequest

	ListAlgorithmsRequest(*sagemaker.ListAlgorithmsInput) sagemaker.ListAlgorithmsRequest

	ListCodeRepositoriesRequest(*sagemaker.ListCodeRepositoriesInput) sagemaker.ListCodeRepositoriesRequest

	ListCompilationJobsRequest(*sagemaker.ListCompilationJobsInput) sagemaker.ListCompilationJobsRequest

	ListEndpointConfigsRequest(*sagemaker.ListEndpointConfigsInput) sagemaker.ListEndpointConfigsRequest

	ListEndpointsRequest(*sagemaker.ListEndpointsInput) sagemaker.ListEndpointsRequest

	ListHyperParameterTuningJobsRequest(*sagemaker.ListHyperParameterTuningJobsInput) sagemaker.ListHyperParameterTuningJobsRequest

	ListLabelingJobsRequest(*sagemaker.ListLabelingJobsInput) sagemaker.ListLabelingJobsRequest

	ListLabelingJobsForWorkteamRequest(*sagemaker.ListLabelingJobsForWorkteamInput) sagemaker.ListLabelingJobsForWorkteamRequest

	ListModelPackagesRequest(*sagemaker.ListModelPackagesInput) sagemaker.ListModelPackagesRequest

	ListModelsRequest(*sagemaker.ListModelsInput) sagemaker.ListModelsRequest

	ListNotebookInstanceLifecycleConfigsRequest(*sagemaker.ListNotebookInstanceLifecycleConfigsInput) sagemaker.ListNotebookInstanceLifecycleConfigsRequest

	ListNotebookInstancesRequest(*sagemaker.ListNotebookInstancesInput) sagemaker.ListNotebookInstancesRequest

	ListSubscribedWorkteamsRequest(*sagemaker.ListSubscribedWorkteamsInput) sagemaker.ListSubscribedWorkteamsRequest

	ListTagsRequest(*sagemaker.ListTagsInput) sagemaker.ListTagsRequest

	ListTrainingJobsRequest(*sagemaker.ListTrainingJobsInput) sagemaker.ListTrainingJobsRequest

	ListTrainingJobsForHyperParameterTuningJobRequest(*sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) sagemaker.ListTrainingJobsForHyperParameterTuningJobRequest

	ListTransformJobsRequest(*sagemaker.ListTransformJobsInput) sagemaker.ListTransformJobsRequest

	ListWorkteamsRequest(*sagemaker.ListWorkteamsInput) sagemaker.ListWorkteamsRequest

	RenderUiTemplateRequest(*sagemaker.RenderUiTemplateInput) sagemaker.RenderUiTemplateRequest

	SearchRequest(*sagemaker.SearchInput) sagemaker.SearchRequest

	StartNotebookInstanceRequest(*sagemaker.StartNotebookInstanceInput) sagemaker.StartNotebookInstanceRequest

	StopCompilationJobRequest(*sagemaker.StopCompilationJobInput) sagemaker.StopCompilationJobRequest

	StopHyperParameterTuningJobRequest(*sagemaker.StopHyperParameterTuningJobInput) sagemaker.StopHyperParameterTuningJobRequest

	StopLabelingJobRequest(*sagemaker.StopLabelingJobInput) sagemaker.StopLabelingJobRequest

	StopNotebookInstanceRequest(*sagemaker.StopNotebookInstanceInput) sagemaker.StopNotebookInstanceRequest

	StopTrainingJobRequest(*sagemaker.StopTrainingJobInput) sagemaker.StopTrainingJobRequest

	StopTransformJobRequest(*sagemaker.StopTransformJobInput) sagemaker.StopTransformJobRequest

	UpdateCodeRepositoryRequest(*sagemaker.UpdateCodeRepositoryInput) sagemaker.UpdateCodeRepositoryRequest

	UpdateEndpointRequest(*sagemaker.UpdateEndpointInput) sagemaker.UpdateEndpointRequest

	UpdateEndpointWeightsAndCapacitiesRequest(*sagemaker.UpdateEndpointWeightsAndCapacitiesInput) sagemaker.UpdateEndpointWeightsAndCapacitiesRequest

	UpdateNotebookInstanceRequest(*sagemaker.UpdateNotebookInstanceInput) sagemaker.UpdateNotebookInstanceRequest

	UpdateNotebookInstanceLifecycleConfigRequest(*sagemaker.UpdateNotebookInstanceLifecycleConfigInput) sagemaker.UpdateNotebookInstanceLifecycleConfigRequest

	UpdateWorkteamRequest(*sagemaker.UpdateWorkteamInput) sagemaker.UpdateWorkteamRequest

	WaitUntilEndpointDeleted(context.Context, *sagemaker.DescribeEndpointInput, ...aws.WaiterOption) error

	WaitUntilEndpointInService(context.Context, *sagemaker.DescribeEndpointInput, ...aws.WaiterOption) error

	WaitUntilNotebookInstanceDeleted(context.Context, *sagemaker.DescribeNotebookInstanceInput, ...aws.WaiterOption) error

	WaitUntilNotebookInstanceInService(context.Context, *sagemaker.DescribeNotebookInstanceInput, ...aws.WaiterOption) error

	WaitUntilNotebookInstanceStopped(context.Context, *sagemaker.DescribeNotebookInstanceInput, ...aws.WaiterOption) error

	WaitUntilTrainingJobCompletedOrStopped(context.Context, *sagemaker.DescribeTrainingJobInput, ...aws.WaiterOption) error

	WaitUntilTransformJobCompletedOrStopped(context.Context, *sagemaker.DescribeTransformJobInput, ...aws.WaiterOption) error
}

var _ SageMakerAPI = (*sagemaker.SageMaker)(nil)
