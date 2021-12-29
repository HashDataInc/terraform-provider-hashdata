package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/HashDataInc/terraform-provider-hashdata/internal/provider/cloudmgr"
	_nethttp "net/http"
)

func resourceComputing() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Computing.",

		CreateContext: resourceComputingCreate,
		ReadContext:   resourceComputingRead,
		UpdateContext: resourceComputingUpdate,
		DeleteContext: resourceComputingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"catalog": {
				Description: "catalog UUID.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"image": {
				Description: "image.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"operation": {
				Description: "operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"master": {
				Description: "master.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type": {
							Description: "master instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "master volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "master volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"zone": {
							Description: "master zone.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"segment": {
				Description: "segment.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Description: "segment count.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"instance_type": {
							Description: "segment instance_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_type": {
							Description: "segment volume_type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_size": {
							Description: "segment volume_size.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"zone": {
							Description: "segment zone.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"extra": {
				Description: "extra.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc": {
							Description: "vpc.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"subnet": {
							Description: "subnet.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"keypair": {
							Description: "keypair.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceComputingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := *cloudmgr.NewCoreCreateWarehouseRequest()
	apiClient := meta.(*cloudmgr.APIClient)
	errMsg, err2 := checkComputingCreateSchema(d)
	if err2 != nil {
		return diag.Errorf(errMsg)
	}
	image := d.Get("image")
	catalog := d.Get("catalog").(string)

	masterPropertiesRaw := d.Get("master").(*schema.Set).List()
	var masterProperties = masterPropertiesRaw[0].(map[string]interface{})
	var masterCount int32 = 1

	master := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        &masterCount,
			InstanceType: String(masterProperties["instance_type"].(string)),
			VolumeType:   String(masterProperties["volume_type"].(string)),
			VolumeSize:   Int32(masterProperties["volume_size"].(int)),
			Image:        String(image.(string)),
			Zone:         String(masterProperties["zone"].(string)),
		},
	}

	segmentPropertiesRaw := d.Get("segment").(*schema.Set).List()
	var segmentProperties = segmentPropertiesRaw[0].(map[string]interface{})
	segment := cloudmgr.CoreCreateServiceComponentRequest{
		Iaas: &cloudmgr.CloudmgrcoreIaasResource{
			Count:        Int32(segmentProperties["count"].(int)),
			InstanceType: String(segmentProperties["instance_type"].(string)),
			VolumeType:   String(segmentProperties["volume_type"].(string)),
			VolumeSize:   Int32(segmentProperties["volume_size"].(int)),
			Image:        String(image.(string)),
			Zone:         String(segmentProperties["zone"].(string)),
		},
	}

	if extraRaw, ok := d.GetOk("extra"); ok {
		extraMap := extraRaw.(*schema.Set).List()[0].(map[string]interface{})
		extra := cloudmgr.CoreCreateServiceIaasExtraRequest{}
		if vpc, ok := extraMap["vpc"]; ok {
			extra.Vpc = String(vpc.(string))
		}
		if subnet, ok := extraMap["subnet"]; ok {
			extra.Subnet = String(subnet.(string))
		}
		if keypair, ok := extraMap["keypair"]; ok {
			extra.Keypair = String(keypair.(string))
		}
		body.Extras = &extra
	}

	name := d.Get("name").(string)
	body.Name = &name
	body.Master = &master
	body.Segment = &segment
	body.Catalog = &catalog
	var resp cloudmgr.CommonDescribeJobResponse
	var r *_nethttp.Response
	var err error
	resp, r, err = apiClient.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
	if err != nil {
		if errInner1, ok := err.(cloudmgr.GenericOpenAPIError); ok {
			if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
				return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse_Computing`: %s\n", *errInner2.ErrorMessage)
			}
		}
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse_Computing` (Error not format): %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse_Computing``: %s\n", r.Status)
	}
	// response from `CreateWarehouse`: CommonDescribeJobResponse
	d.SetId(resp.GetResourceIds()[0])
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}
	return resourceComputingUpdate(ctx, d, meta)
}

func resourceComputingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	if id == "" {
		return diag.Errorf("Can not read warehouse,because id is empty")
	}

	apiClient := meta.(*cloudmgr.APIClient)

	var resp cloudmgr.CoreListInstanceResponse
	var r *_nethttp.Response
	var err error

	resp, r, err = apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component([]string{"master"}).Execute() //.DescribeInstance(ctx, id).Execute()
	if err != nil {
		if errInner1, ok := err.(cloudmgr.GenericOpenAPIError); ok {
			if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
				return diag.Errorf("Error when calling `CoreWarehouseServiceApi.DescribeInstance`: %s\n", *errInner2.ErrorMessage)
			}
		}
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.DescribeInstance` (Error not format): %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error status code when calling `CoreWarehouseServiceApi.CreateWarehouse``: %d \n", r.StatusCode)
	}
	if *resp.Count == 0 || *resp.Count > 1 {
		return diag.Errorf("Error when ListServiceInstance")
	}
	master := (*resp.Content)[0]
	if param, ok := master.GetArchOk(); !ok {
		d.Set("arch", param)
	}
	if param, ok := master.GetComponentsOk(); !ok {
		d.Set("components", param)
	}
	if param, ok := master.GetCreatedAtOk(); !ok {
		d.Set("created_at", param)
	}
	if param, ok := master.GetDestroyedAtOk(); !ok {
		d.Set("destroyed_at", param)
	}
	if nic, ok := master.GetElasticNicOk(); !ok {
		var nic_map = make(map[string]interface{})
		if param, ok2 := nic.GetElasticNicIdOk(); !ok2 {
			nic_map["elastic_nic_id"] = param
		}
		if param, ok2 := nic.GetIpaddressesOk(); !ok2 {
			nic_map["ipaddresses"] = param
		}
		d.Set("nic", nic_map)
	}
	if param, ok := master.GetHealthStatusOk(); !ok {
		d.Set("health_status", param)
	}
	if param, ok := master.GetHostnameOk(); !ok {
		d.Set("hostname", param)
	}
	if param, ok := master.GetImageOk(); !ok {
		d.Set("image", param)
	}
	if internet, ok := master.GetInternetOk(); !ok {
		var internet_map = make(map[string]interface{})
		if param, ok2 := internet.GetBandwidthOk(); !ok2 {
			internet_map["band_width"] = param
		}
		if param, ok2 := internet.GetElasticIpOk(); !ok2 {
			internet_map["elastic_ip"] = param
		}
		if param, ok2 := internet.GetElasticIpIdOk(); !ok2 {
			internet_map["elastic_ip_id"] = param
		}
		if param, ok2 := internet.GetPublicIpOk(); !ok2 {
			internet_map["public_ip"] = param
		}
		d.Set("internet", internet_map)
	}
	if param, ok := master.GetIpaddressesOk(); !ok {
		d.Set("ipaddresses", param)
	}
	if param, ok := master.GetMessageOk(); !ok {
		d.Set("message", param)
	}
	if param, ok := master.GetNameOk(); !ok {
		d.Set("name", param)
	}
	if param, ok := master.GetResourcePoolOk(); !ok {
		d.Set("resource_pool", param)
	}
	if param, ok := master.GetScaleTypeOk(); !ok {
		d.Set("scale_type", param)
	}
	if param, ok := master.GetServiceOk(); !ok {
		d.Set("service", param)
	}
	if param, ok := master.GetStatusOk(); !ok {
		d.Set("status", param)
	}
	if param, ok := master.GetTenantOk(); !ok {
		d.Set("tenant", param)
	}
	if param, ok := master.GetTypeOk(); !ok {
		d.Set("type", param)
	}
	if param, ok := master.GetVendorOk(); !ok {
		d.Set("vendor", param)
	}
	if param, ok := master.GetZoneOk(); !ok {
		d.Set("zone", param)
	}

	//TODO 判断是否被删除
	return nil
}

func resourceComputingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//You should make sure that when expending cluster the service is in stop status
	id := d.Id()
	if id == "" {
		return diag.Errorf("Can not read warehouse,because id is empty")
	}
	apiClient := meta.(*cloudmgr.APIClient)
	targetSize := make(map[string]interface{})
	if d.HasChange("master") && !d.IsNewResource() {
		if err := getVolumeResizeMap(ctx, "master", id, apiClient, d, &targetSize); err != nil {
			return diag.Errorf(err.Error())
		}
	}
	if d.HasChange("segment") && !d.IsNewResource() {
		//step 1: describe instance
		//step 2: getCount
		//step 3: judgment expend or shrink
		//step 4: stop service
		//step 4: async job
		if err := getVolumeResizeMap(ctx, "segment", id, apiClient, d, &targetSize); err != nil {
			return diag.Errorf(err.Error())
		}
		component := []string{"segment"}
		respListInstance, rListInstance, errListInstance := apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component(component).Execute()
		if errListInstance != nil {
			if errInner1, ok := errListInstance.(cloudmgr.GenericOpenAPIError); ok {
				if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
					return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance`: %s\n", *errInner2.ErrorMessage)
				}
			}
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance` (Error not format): %v\n", errListInstance)
		}
		if rListInstance.StatusCode != 200 {
			return diag.Errorf("Error when calling `CoreServiceApi.ListServiceInstance``: %s\n", rListInstance.Status)
		}
		countOld := respListInstance.GetCount()

		segmentPropertiesRaw := d.Get("segment").(*schema.Set).List()
		var segmentProperties = segmentPropertiesRaw[0].(map[string]interface{})
		countNew := segmentProperties["count"].(int)

		if int32(countNew) != countOld {
			if errCanShrink := canExpendShrinkService(ctx, id, apiClient); errCanShrink != nil {
				return diag.Errorf(errCanShrink.Error())
			}
			componentRequestMap := make(map[string]interface{})
			var respScale cloudmgr.CommonDescribeJobResponse
			var rScale *_nethttp.Response
			var errScale error
			if int32(countNew) > countOld {
				componentRequestMap["segment"] = cloudmgr.CoreScaleOutServiceComponentRequest{
					Iaas: &cloudmgr.CoreScaleOutIaasResource{
						Count: Int32(countNew),
					},
				}
				respScale, rScale, errScale = apiClient.CoreServiceApi.ScaleOutService(ctx, id).Body(cloudmgr.CoreScaleOutServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			} else {
				var remainInstances = make([]string, int(countOld)-countNew)
				for i := 0; i < int(countOld)-countNew; i++ {
					remainInstances[i] = (*respListInstance.Content)[i].GetId()
				}
				componentRequestMap["segment"] = cloudmgr.CoreScaleInServiceComponentRequest{
					Instances: &remainInstances,
				}
				respScale, rScale, errScale = apiClient.CoreServiceApi.ScaleInService(ctx, id).Body(cloudmgr.CoreScaleInServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
			}
			if errScale != nil {
				if errInner1, ok := errScale.(cloudmgr.GenericOpenAPIError); ok {
					if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
						return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService`: %s\n", *errInner2.ErrorMessage)
					}
				}
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` (Error not format): %v\n", errScale)
			}
			if rScale.StatusCode != 200 {
				return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\n", rListInstance.Status)
			}
			if errWaitJob := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respScale.GetId()); errWaitJob != nil {
				return diag.Errorf("Error when wait calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\v", errWaitJob)
			}

		}
	}
	if len(targetSize) != 0 {
		if len(targetSize) != 0 {
			if errResize := resizeVolume(ctx, id, apiClient, targetSize); errResize != nil {
				return diag.Errorf(errResize.Error())
			}
		}
	}

	if op, ok := d.GetOk("operation"); ok {
		if op == OPERATE_START {
			if errStart := startService(ctx, id, apiClient); errStart != nil {
				return diag.Errorf(errStart.Error())
			}
		} else if op == OPERATE_STOP {
			if errStop := stopService(ctx, id, apiClient); errStop != nil {
				return diag.Errorf(errStop.Error())
			}
		} else if op == OPERATE_RESTART {
			if errReStart := restartService(ctx, id, apiClient); errReStart != nil {
				return diag.Errorf(errReStart.Error())
			}
		} else {
			return diag.Errorf("Wrong operation,operation must in one of 'start','stop' ,'restart'.")
		}
	}
	return resourceComputingRead(ctx, d, meta)
}

func resourceComputingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*cloudmgr.APIClient)
	resourceId := d.Id()
	if resourceId == "" {
		return diag.Errorf(WAREHOUSE_ID + " not found! ")
	}
	resp, r, err := apiClient.CoreServiceApi.DeleteService(ctx, resourceId).Force(true).Execute()
	if err != nil {
		if errInner1, ok := err.(cloudmgr.GenericOpenAPIError); ok {
			if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
				return diag.Errorf("Error when calling `CoreServiceApi.DeleteService`: %s\n", *errInner2.ErrorMessage)
			}
		}
		return diag.Errorf("Error when calling `CoreServiceApi.DeleteService` (Error not format): %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Delete resource fail with %d . ", r.StatusCode)
	}
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}
	d.SetId("")
	return nil
}

func checkComputingCreateSchema(d *schema.ResourceData) (string, error) {
	res := ""
	_, ok := d.GetOk("name")
	if !ok {
		res += "schema name field is missing\n"
	}
	_, ok = d.GetOk("image")
	if !ok {
		res += "schema image field is missing\n"
	}
	_, ok = d.GetOk("catalog")
	if !ok {
		res += "schema name field is missing\n"
	}

	masterRaw, ok := d.GetOk("master")
	if !ok {
		res += "schema master field is missing\n"
	}
	masterMap := masterRaw.(*schema.Set).List()[0].(map[string]interface{})
	if _, ok := masterMap["instance_type"]; !ok {
		res += "schema master.instance_type field is missing\n"
	}
	if _, ok := masterMap["volume_type"]; !ok {
		res += "schema master.volume_type field is missing\n"
	}
	if _, ok := masterMap["volume_size"]; !ok {
		res += "schema master.volume_size field is missing\n"
	}
	if _, ok := masterMap["zone"]; !ok {
		res += "schema master.zone field is missing\n"
	}

	segmentRaw, ok := d.GetOk("segment")
	if !ok {
		res += "schema segment field is missing\n"
	}
	segmentMap := segmentRaw.(*schema.Set).List()[0].(map[string]interface{})
	if _, ok := segmentMap["count"]; !ok {
		res += "schema segment.count field is missing\n"
	}
	if _, ok := segmentMap["instance_type"]; !ok {
		res += "schema segment.instance_type field is missing\n"
	}
	if _, ok := segmentMap["volume_type"]; !ok {
		res += "schema segment.volume_type field is missing\n"
	}
	if _, ok := segmentMap["volume_size"]; !ok {
		res += "schema segment.volume_size field is missing\n"
	}
	if _, ok := segmentMap["zone"]; !ok {
		res += "schema segment.zone field is missing\n"
	}
	//extraRaw, ok := d.GetOk("extra")
	//if !ok {
	//	res += "schema extra field is missing\n"
	//}
	//extraMap := extraRaw.(map[string]interface{})
	//if _, ok := extraMap["vpc"]; !ok {
	//	res += "schema extra.vpc field is missing\n"
	//}
	//if _, ok := extraMap["subnet"]; !ok {
	//	res += "schema extra.subnet field is missing\n"
	//}
	//if _, ok := extraMap["keypair"]; !ok {
	//	res += "schema extra.keypair field is missing\n"
	//}
	if res != "" {
		return res, fmt.Errorf("Input is illegal. ")
	}
	return "", nil
}
