package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/HashDataInc/terraform-provider-hashdata/internal/provider/cloudmgr"
	"log"
	"math/rand"
	_nethttp "net/http"
	"strconv"
	"strings"
	"time"
)

func Int32(v int) *int32 {
	t := int32(v)
	return &t
}
func String(v string) *string {
	return &v
}
func Bool(v bool) *bool {
	return &v
}

func simpleRetry(fn func() error) error {
	return retry(MAX_RETRY_TIMES, 10*time.Second, fn)
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func waitJobComplete(ctx context.Context, clt *cloudmgr.CoreJobServiceApiService, id string) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}
	for i := 1; i <= 500; i++ {
		var resp cloudmgr.CommonDescribeJobResponse
		var r *_nethttp.Response
		var err error
		resp, r, err = clt.DescribeJob(ctx, id).Execute()
		if err != nil {
			return err
		}
		if r.StatusCode != 200 {
			return fmt.Errorf("Bad response code with %s ", r.Status)
		}
		//TODO 判断是否被删除
		status := strings.ToLower(resp.GetStatus())
		if status == JOB_WAIT_PENDING || status == JOB_WAIT_RUNNING {
			time.Sleep(10 * time.Second)
			log.Println("Still waiting current state is " + status + ". [" + strconv.Itoa(i*10) + "s elapsed]")
			continue
		} else if status == JOB_FAILED_ABANDONED || status == JOB_FAILED_FAILURE {
			return fmt.Errorf("Error instance create failed, request id %s, status %s ", id, status)
		} else if status == JOB_SUCCESS {
			return nil
		} else {
			return fmt.Errorf("Error unknow status code %s ", status)
		}
	}
	return fmt.Errorf("Timeout while waiting for state become 'success'. ")
}

func startService(ctx context.Context, id string, apiClient *cloudmgr.APIClient) error {

	respDes,rDes,errDes:= apiClient.CoreServiceApi.DescribeService(ctx,id).Execute()
	if errDes = checkErrAndNetResponse(errDes, rDes, "CoreServiceApi.DescribeService"); errDes != nil {
		return errDes
	}
	if strings.ToLower(*respDes.Status) == "started"{
		return nil
	}
	if strings.ToLower(*respDes.Status) != "stopped"{
		return fmt.Errorf("Service should be stopped when calling startService. ")
	}
	respDep, rDep, errDep := apiClient.CoreServiceApi.ListServiceDependent(ctx, id).Execute()
	if errDep = checkErrAndNetResponse(errDep, rDep, "CoreServiceApi.ListServiceDependent"); errDep != nil {
		return errDep
	}
	respStart, rStart, errStart := apiClient.CoreServiceApi.StartService(ctx, id).Body(cloudmgr.CoreStartServiceRequest{}).Execute()
	if errStart = checkErrAndNetResponse(errStart, rStart, "CoreServiceApi.StartService"); errStart != nil {
		return errStart
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStart.GetId()); errRefresh != nil {
		return errRefresh
	}
	for i := 0; i < int(*respDep.Count); i++ {
		respStart, rStart, errStart := apiClient.CoreServiceApi.StartService(ctx, *(*respDep.Content)[i].Dependent).Body(cloudmgr.CoreStartServiceRequest{}).Execute()
		if errStart = checkErrAndNetResponse(errStart, rStart, "CoreServiceApi.StartService"); errStart != nil {
			return errStart
		}
		if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStart.GetId()); errRefresh != nil {
			return errRefresh
		}
	}
	return nil
}

func stopService(ctx context.Context, id string, apiClient *cloudmgr.APIClient) error {
	respDes,rDes,errDes:= apiClient.CoreServiceApi.DescribeService(ctx,id).Execute()
	if errDes = checkErrAndNetResponse(errDes, rDes, "CoreServiceApi.DescribeService"); errDes != nil {
		return errDes
	}
	if strings.ToLower(*respDes.Status) == "stopped"{
		return nil
	}
	if strings.ToLower(*respDes.Status) != "started"{
		return fmt.Errorf("Service should be stopped when calling startService. ")
	}
	force := true
	respDep, rDep, errDep := apiClient.CoreServiceApi.ListServiceDependent(ctx, id).Execute()
	if errDep = checkErrAndNetResponse(errDep, rDep, "CoreServiceApi.ListServiceDependent"); errDep != nil {
		return errDep
	}
	for i := 0; i < int(*respDep.Count); i++ {
		respStop, rStop, errStop := apiClient.CoreServiceApi.StopService(ctx, *(*respDep.Content)[i].Dependent).Body(cloudmgr.CoreStopServiceRequest{
			Force: &force,
		}).Execute()
		if errStop = checkErrAndNetResponse(errStop, rStop, "CoreServiceApi.StopService"); errStop != nil {
			return errStop
		}
		if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStop.GetId()); errRefresh != nil {
			return errRefresh
		}
	}

	respStop, rStop, errStop := apiClient.CoreServiceApi.StopService(ctx, id).Body(cloudmgr.CoreStopServiceRequest{
		Force: &force,
	}).Execute()
	if errStop = checkErrAndNetResponse(errStop, rStop, "CoreServiceApi.StopService"); errStop != nil {
		return errStop
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStop.GetId()); errRefresh != nil {
		return errRefresh
	}
	return nil
}

func restartService(ctx context.Context, id string, apiClient *cloudmgr.APIClient) error {
	if errStop := stopService(ctx, id, apiClient); errStop != nil {
		return errStop
	}
	if errStart := startService(ctx, id, apiClient); errStart != nil {
		return errStart
	}
	return nil
}

func canExpendShrinkService(ctx context.Context, id string, apiClient *cloudmgr.APIClient) error {
	resp, r, err := apiClient.CoreServiceApi.DescribeService(ctx, id).Execute()
	if err != nil {
		if errInner1, ok := err.(cloudmgr.GenericOpenAPIError); ok {
			if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
				return fmt.Errorf("Error when calling `CoreServiceApi.DescribeService`: %s\n", *errInner2.ErrorMessage)
			}
		}
		return fmt.Errorf("Error when calling `CoreServiceApi.DescribeService` (Error not format): %v\n", err)
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("Error when calling `CoreServiceApi.DescribeService` StatusCode != 200: %s\n", r.Status)
	}
	currStatus := strings.ToLower(*resp.Status)
	if currStatus == SERVICE_STARTED {
		return nil
	}
	return fmt.Errorf("Service status should be STARTED ")
}

func checkErrAndNetResponse(err error, r *_nethttp.Response, content string) error {
	if err != nil {
		if errInner1, ok := err.(cloudmgr.GenericOpenAPIError); ok {
			if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
				return fmt.Errorf("Error when calling `%s`: %s\n", content, *errInner2.ErrorMessage)
			}
		}
		return fmt.Errorf("Error when calling `%s` (Error not format): %v\n", content, err)
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("Wrong response with %s and code %d ", content, r.StatusCode)
	}
	return nil
}

func getVolumeResizeMap(ctx context.Context, componentInput string, id string, apiClient *cloudmgr.APIClient, d *schema.ResourceData, targetSize *map[string]interface{}) error {
	//volume
	component := []string{componentInput}
	respListInstance, rListInstance, errListInstance := apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component(component).Execute()
	if errListInstance = checkErrAndNetResponse(errListInstance, rListInstance, "CoreServiceApi.ListServiceInstance with component "+componentInput); errListInstance != nil {
		return fmt.Errorf(errListInstance.Error())
	}
	respListVolume, rListVolume, errListVolume := apiClient.CoreInstanceServiceApi.ListVolume(ctx, respListInstance.GetContent()[0].GetId()).Execute()
	if errListVolume = checkErrAndNetResponse(errListVolume, rListVolume, "CoreServiceApi.ListVolume with id "+respListInstance.GetContent()[0].GetId()); errListVolume != nil {
		return fmt.Errorf(errListVolume.Error())
	}
	componentPropertiesRaw := d.Get(componentInput).(*schema.Set).List()
	var componentProperties = componentPropertiesRaw[0].(map[string]interface{})
	currServiceSize := int(*(*(respListVolume.Content))[0].Size)
	currLocalSize := componentProperties["volume_size"].(int)
	if currLocalSize < currServiceSize {
		return fmt.Errorf("Not support volume size shrink. ")
	}
	if currLocalSize > currServiceSize {
		respDescribeVolumeType, rDescribeVolumeType, errDescribeVolumeType := apiClient.CoreVolumeTypeServiceApi.DescribeVolumeType(context.Background(), componentProperties["volume_type"].(string)).Execute()
		if errDescribeVolumeType = checkErrAndNetResponse(errDescribeVolumeType, rDescribeVolumeType, "Describe master volume type"); errDescribeVolumeType != nil {
			return fmt.Errorf(errDescribeVolumeType.Error())
		}
		maxSize := int(*respDescribeVolumeType.MaxSize)
		step := int(*respDescribeVolumeType.StepSize)
		if currLocalSize > maxSize {
			return fmt.Errorf(componentInput + " volume size should less than " + strconv.Itoa(maxSize))
		}
		if (currLocalSize-currServiceSize)%step != 0 {
			return fmt.Errorf(componentInput + " volume size should be step of " + strconv.Itoa(step))
		}
		(*targetSize)[componentInput] = currLocalSize
	}
	return nil
}

func resizeVolume(ctx context.Context, id string, apiClient *cloudmgr.APIClient, targetSize map[string]interface{}) error {
	if errStop := stopService(ctx, id, apiClient); errStop != nil {
		return errStop
	}
	respResize, rResize, errResize := apiClient.CoreServiceApi.ResizeVolumes(ctx, id).Body(cloudmgr.CoreResizeServiceVolumesRequest{
		TargetVolumeSize: &targetSize,
	}).Execute()
	if errResize = checkErrAndNetResponse(errResize, rResize, "CoreServiceApi.ResizeVolumes"); errResize != nil {
		return errResize
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respResize.GetId()); errRefresh != nil {
		return errRefresh
	}
	if errStart := startService(ctx, id, apiClient); errStart != nil {
		return errStart
	}
	return nil
}
