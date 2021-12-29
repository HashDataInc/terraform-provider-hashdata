package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/HashDataInc/terraform-provider-hashdata/internal/provider/cloudmgr"
	_nethttp "net/http"
	"strconv"
)

func resourceWarehouse() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Warehouse.",

		CreateContext: resourceWarehouseCreate,
		ReadContext:   resourceWarehouseRead,
		UpdateContext: resourceWarehouseUpdate,
		DeleteContext: resourceWarehouseDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		//TODO computing 扩容缩容 catalog 扩容缩容 fdb 扩容缩容 warehouse 和 volume 支持扩容
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name.",
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
			"metadata": {
				Description: "metadata.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_database": {
							Description: "default_database.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"default_user": {
							Description: "default_user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"default_password": {
							Description: "default_password.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"number_segments": {
							Description: "number_segments.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"logic_part": {
							Description: "logic_part.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
			"feature": {
				Description: "feature.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_storage": {
							Description: "is_local_storage.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"mirror_standby": {
							Description: "is_mirror_standby.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWarehouseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	body := *cloudmgr.NewCoreCreateWarehouseRequest()
	apiClient := meta.(*cloudmgr.APIClient)
	errMsg, err2 := checkWarehouseCreateSchema(d, apiClient)
	if err2 != nil {
		return diag.Errorf(errMsg)
	}

	image := d.Get("image")

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

	metadataPropertiesRaw := d.Get("metadata").(*schema.Set).List()
	var metadataProperties = metadataPropertiesRaw[0].(map[string]interface{})
	var metadata = make(map[string]interface{})
	metadata["default_database"] = metadataProperties["default_database"].(string)
	metadata["default_user"] = metadataProperties["default_user"].(string)
	metadata["default_password"] = metadataProperties["default_password"].(string)
	if numberSegments, ok := metadataProperties["number_segments"]; ok {
		if val, ok := numberSegments.(int); ok && val != 0 {
			metadata["number_segments"] = numberSegments
		}
	}
	if logicPart, ok := metadataProperties["logic_part"]; ok {
		if val, ok := logicPart.(int); ok && val != 0 {
			metadata["logic_part"] = logicPart
		}
	}
	if featureOk, ok := d.GetOk("feature"); ok {
		featurePropertiesRaw := featureOk.(*schema.Set).List()
		var featureProperties = featurePropertiesRaw[0].(map[string]interface{})
		feature := cloudmgr.CoreCreateServiceFeatureRequest{}
		if localStorage, ok := featureProperties["local_storage"]; ok {
			feature.LocalStorage = Bool(localStorage.(bool))
			if localStorage.(bool) {
				if ossOk, ok := d.GetOk("oss"); ok {
					ossProperties := ossOk.(*schema.Set).List()[0].(map[string]interface{})
					if ossName, ok := ossProperties["name"]; ok {
						oss := cloudmgr.CoreCreateServiceOssZoneRequest{
							Name: String(ossName.(string)),
						}
						body.Oss = &oss
					} else {
						return diag.Errorf("Schema oss.name field not found.")
					}
				} else {
					return diag.Errorf("Schema oss field not found.")
				}
			}
		}
		if mirrorStandby, ok := featureProperties["mirror_standby"]; ok {
			feature.LocalStorage = Bool(mirrorStandby.(bool))
			if mirrorStandby.(bool) {
				if standBy, ok := d.GetOk("standby"); ok {
					standbyProperties := standBy.(*schema.Set).List()[0].(map[string]interface{})
					if _, ok := standbyProperties["count"]; !ok {
						return diag.Errorf("Schema standby.count field not found.")
					}
					if _, ok := standbyProperties["instance_type"]; !ok {
						return diag.Errorf("Schema standby.instance_type field not found.")
					}
					if _, ok := standbyProperties["volume_type"]; !ok {
						return diag.Errorf("Schema standby.volume_type field not found.")
					}
					if _, ok := standbyProperties["volume_size"]; !ok {
						return diag.Errorf("Schema standby.volume_size field not found.")
					}
					if _, ok := standbyProperties["image"]; !ok {
						return diag.Errorf("Schema standby.image field not found.")
					}
					if _, ok := standbyProperties["zone"]; !ok {
						return diag.Errorf("Schema standby.zone field not found.")
					}

					//TODO  mirrorStabdby的image要不要和其余的一样统一
					standby := cloudmgr.CoreCreateServiceComponentRequest{
						Iaas: &cloudmgr.CloudmgrcoreIaasResource{
							Count:        Int32(standbyProperties["count"].(int)),
							InstanceType: String(standbyProperties["instance_type"].(string)),
							VolumeType:   String(standbyProperties["volume_type"].(string)),
							VolumeSize:   Int32(standbyProperties["volume_size"].(int)),
							Image:        String(standbyProperties["image"].(string)),
							Zone:         String(standbyProperties["zone"].(string)),
						},
					}
					body.Standby = &standby
				} else {
					return diag.Errorf("Schema standby field not found.")
				}
			}
		}
		body.Features = &feature
	}

	name := d.Get("name").(string)
	body.Name = &name
	body.Master = &master
	body.Segment = &segment
	body.Metadata = &metadata

	var resp cloudmgr.CommonDescribeJobResponse
	var r *_nethttp.Response
	var err error
	//Do Not retry here
	resp, r, err = apiClient.CoreWarehouseServiceApi.CreateWarehouse(ctx).Body(body).Execute()
	if err != nil {
		if errInner1, ok := err.(cloudmgr.GenericOpenAPIError); ok {
			if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
				return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse`: %s\n", *errInner2.ErrorMessage)
			}
		}
		return diag.Errorf("Error when calling `CoreWarehouseServiceApi.CreateWarehouse` (Error not format): %v\n", err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Error status code when calling `CoreWarehouseServiceApi.CreateWarehouse` : %s\n", r.Status)
	}

	d.SetId(resp.GetResourceIds()[0])
	if errRefresh := waitJobComplete(ctx, apiClient.CoreJobServiceApi, resp.GetId()); errRefresh != nil {
		return diag.Errorf(errRefresh.Error())
	}
	//return nil
	return resourceWarehouseUpdate(ctx, d, meta)
}

func resourceWarehouseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceWarehouseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	if id == "" {
		return diag.Errorf("Can not read warehouse,because id is empty")
	}
	apiClient := meta.(*cloudmgr.APIClient)
	targetSize := make(map[string]interface{})
	if d.HasChange("master") && !d.IsNewResource() {
		//volume
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
		resp, r, err := apiClient.CoreServiceApi.ListServiceInstance(ctx, id).Component(component).Execute()
		if err = checkErrAndNetResponse(err, r, "CoreServiceApi.ListServiceInstance with component segment"); err != nil {
			return diag.Errorf(err.Error())
		}
		countOld := resp.GetCount()
		segmentPropertiesRaw := d.Get("segment").(*schema.Set).List()
		var segmentProperties = segmentPropertiesRaw[0].(map[string]interface{})
		countNew := segmentProperties["count"].(int)
		if int32(countNew) != countOld {
			componentRequestMap := make(map[string]interface{})
			var respScaleOut cloudmgr.CommonDescribeJobResponse
			var rScaleOut *_nethttp.Response
			var errScaleOut error
			if int32(countNew) > countOld {
				if errCanShrink := canExpendShrinkService(ctx, id, apiClient); errCanShrink != nil {
					return diag.Errorf(errCanShrink.Error())
				}
				componentRequestMap["segment"] = cloudmgr.CoreScaleOutServiceComponentRequest{
					Iaas: &cloudmgr.CoreScaleOutIaasResource{
						Count: Int32(countNew),
					},
				}
				respScaleOut, rScaleOut, errScaleOut = apiClient.CoreServiceApi.ScaleOutService(ctx, id).Body(cloudmgr.CoreScaleOutServiceRequest{
					Component: &componentRequestMap,
				}).Execute()
				if errScaleOut != nil {
					if errInner1, ok := errScaleOut.(cloudmgr.GenericOpenAPIError); ok {
						if errInner2, ok := errInner1.Model().(cloudmgr.CommonActionResponse); ok {
							return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService`: %s\n", *errInner2.ErrorMessage)
						}
					}
					return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` (Error not format): %v\n", errScaleOut)
				}
				if rScaleOut.StatusCode != 200 {
					return diag.Errorf("Error when calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\n", rScaleOut.Status)
				}
				if errWaitJob := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respScaleOut.GetId()); errWaitJob != nil {
					return diag.Errorf("Error when wait calling `CoreServiceApi.ScaleOutService` or ScaleInService: %s\v", errWaitJob)
				}
			} else {
				return diag.Errorf("Not support shrink warehouse")
			}
		}
	}
	if len(targetSize) != 0 {
		if errResize := resizeVolume(ctx, id, apiClient, targetSize); errResize != nil {
			return diag.Errorf(errResize.Error())
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
	return resourceWarehouseRead(ctx, d, meta)
}

func resourceWarehouseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func checkWarehouseCreateSchema(d *schema.ResourceData, apiClient *cloudmgr.APIClient) (string, error) {
	res := ""
	_, ok := d.GetOk("name")
	if !ok {
		res += "schema name field is missing\n"
	}
	_, ok = d.GetOk("image")
	if !ok {
		res += "schema image field is missing\n"
	}
	masterRaw, ok := d.GetOk("master")
	if !ok {
		res += "schema master field is missing\n"
	}
	masterMap := masterRaw.(*schema.Set).List()[0].(map[string]interface{})
	if _, ok := masterMap["instance_type"]; !ok {
		res += "schema master.instance_type field is missing\n"
	}
	masterVolumeType, ok := masterMap["volume_type"]
	if !ok {
		res += "schema master.volume_type field is missing\n"
	}
	masterVolumeSize, ok := masterMap["volume_size"]
	if !ok {
		res += "schema master.volume_size field is missing\n"
	}

	resp, r, err := apiClient.CoreVolumeTypeServiceApi.DescribeVolumeType(context.Background(), masterVolumeType.(string)).Execute()
	if err = checkErrAndNetResponse(err, r, "Describe master volume type"); err != nil {
		res += err.Error()
		return "", fmt.Errorf(res)
	}

	minSize := int(*resp.MinSize)
	maxSize := int(*resp.MaxSize)
	step := int(*resp.StepSize)
	if masterVolumeSize.(int) > maxSize || masterVolumeSize.(int) < minSize {
		res += "master volume size should less than " + strconv.Itoa(maxSize) + " and more than " + strconv.Itoa(minSize)
	}
	if (masterVolumeSize.(int)-minSize)%step != 0 {
		res += "master volume size should be start with " + strconv.Itoa(minSize) + " and step of " + strconv.Itoa(step)
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
	segmentVolumeType, ok := segmentMap["volume_type"]
	if !ok {
		res += "schema segment.volume_type field is missing\n"
	}
	if segmentVolumeType != masterVolumeType {
		resp, r, err := apiClient.CoreVolumeTypeServiceApi.DescribeVolumeType(context.Background(), masterVolumeType.(string)).Execute()
		if err = checkErrAndNetResponse(err, r, "Describe segment volume type"); err != nil {
			res += err.Error()
			return "", fmt.Errorf(res)
		}
		minSize = int(*resp.MinSize)
		maxSize = int(*resp.MaxSize)
		step = int(*resp.StepSize)
	}
	segmentVolumeSize, ok := segmentMap["volume_size"]
	if !ok {
		res += "schema segment.volume_size field is missing\n"
	}

	if segmentVolumeSize.(int) > maxSize || segmentVolumeSize.(int) < minSize {
		res += "master volume size should less than " + strconv.Itoa(maxSize) + " and more than " + strconv.Itoa(minSize)
	}
	if (segmentVolumeSize.(int)-minSize)%step != 0 {
		res += "master volume size should be start with " + strconv.Itoa(minSize) + " and step of " + strconv.Itoa(step)
	}

	if _, ok := segmentMap["zone"]; !ok {
		res += "schema segment.zone field is missing\n"
	}

	metadataRaw, ok := d.GetOk("metadata")
	if !ok {
		res += "schema metadata field is missing\n"
	}
	metadataMap := metadataRaw.(*schema.Set).List()[0].(map[string]interface{})
	if _, ok := metadataMap["default_database"]; !ok {
		res += "schema metadata.default_database field is missing\n"
	}
	if _, ok := metadataMap["default_user"]; !ok {
		res += "schema metadata.default_user field is missing\n"
	}
	if _, ok := metadataMap["default_password"]; !ok {
		res += "schema metadata.default_password field is missing\n"
	}
	if res != "" {
		return res, fmt.Errorf("Input is illegal. ")
	}
	return "", nil
}
