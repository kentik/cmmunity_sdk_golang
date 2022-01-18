//go:build examples
// +build examples

//nolint:forbidigo,goconst,testpackage
package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestDemonstrateCloudExportAPIWithAWS(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudExportAPIWithAWS()
	assert.NoError(t, err)
}

func TestDemonstrateCloudExportAPIWithAzure(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudExportAPIWithAzure()
	assert.NoError(t, err)
}

func TestDemonstrateCloudExportAPIWithGCE(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudExportAPIWithGCE()
	assert.NoError(t, err)
}

func TestDemonstrateCloudExportAPIWithIBM(t *testing.T) {
	t.Parallel()
	err := demonstrateCloudExportAPIWithIBM()
	assert.NoError(t, err)
}

func demonstrateCloudExportAPIWithAWS() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Getting all cloud exports")
	getAllResp, err := client.CloudExports.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("client.CloudExports.GetAll: %w", err)
	}

	fmt.Println("Invalid cloud exports count:", getAllResp.InvalidCloudExportsCount)
	fmt.Printf("Listed cloud exports: %+v\n", getAllResp.CloudExports)

	fmt.Println("Creating AWS cloud export")
	ce := models.NewAWSCloudExport(models.CloudExportAWSRequiredFields{
		Name:          "example-aws-export",
		PlanID:        "11467",
		CloudProvider: models.CloudProviderAWS,
		AWSProperties: models.AWSPropertiesRequiredFields{
			Bucket: "dummy-bucket",
		},
	})
	ce.Type = models.CloudExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy AWS description"
	ce.APIRoot = "https://dummy.api.kentik.com"
	ce.FlowDestination = "https://dummy.flow.kentik.com"
	ce.AWSProperties.IAMRoleARN = "dummy-iam-role-arn"
	ce.AWSProperties.Region = "dummy-region"
	ce.AWSProperties.DeleteAfterRead = pointer.ToBool(true)
	ce.AWSProperties.MultipleBuckets = pointer.ToBool(true)
	ce.BGP = &models.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}

	ce, err = client.CloudExports.Create(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Create: %w", err)
	}

	fmt.Printf("Created AWS cloud export: %+v\n", ce)

	fmt.Println("Getting AWS cloud export")
	ce, err = client.CloudExports.Get(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Get: %w", err)
	}

	fmt.Printf("Got AWS cloud export: %+v\n", ce)

	fmt.Println("Updating AWS cloud export")
	ce.Name = "updated-example-aws-export"
	ce.Description = "Updated description"
	ce.AWSProperties.Bucket = "updated-bucket"
	ce.BGP.UseBGPDeviceID = "updated-bgp-device-id"
	ce, err = client.CloudExports.Update(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Update: %w", err)
	}

	fmt.Printf("Updated cloud export: %+v\n", ce)

	fmt.Println("Deleting AWS cloud export")
	err = client.CloudExports.Delete(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Delete: %w", err)
	}

	fmt.Println("Deleted AWS cloud export")
	return nil
}

func demonstrateCloudExportAPIWithAzure() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Creating Azure cloud export")
	ce := models.NewAzureCloudExport(models.CloudExportAzureRequiredFields{
		Name:          "example-azure-export",
		PlanID:        "11467",
		CloudProvider: models.CloudProviderAzure,
		AzureProperties: models.AzurePropertiesRequiredFields{
			Location:       "dummy-location",
			ResourceGroup:  "dummy-rg",
			StorageAccount: "dummy-sa",
			SubscriptionID: "dummy-sid",
		},
	})
	ce.Type = models.CloudExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy Azure description"
	ce.APIRoot = "https://dummy.api.kentik.com"
	ce.FlowDestination = "https://dummy.flow.kentik.com"
	ce.AzureProperties.SecurityPrincipalEnabled = pointer.ToBool(true)
	ce.BGP = &models.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}
	ce, err = client.CloudExports.Create(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Create: %w", err)
	}

	fmt.Printf("Created Azure cloud export: %+v\n", ce)

	fmt.Println("Deleting Azure cloud export")
	err = client.CloudExports.Delete(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Delete: %w", err)
	}

	fmt.Println("Deleted Azure cloud export")
	return nil
}

func demonstrateCloudExportAPIWithGCE() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Creating GCE cloud export")
	fmt.Println("Creating Azure cloud export")
	ce := models.NewGCECloudExport(models.CloudExportGCERequiredFields{
		Name:          "example-gce-export",
		PlanID:        "11467",
		CloudProvider: models.CloudProviderGCE,
		GCEProperties: models.GCEPropertiesRequiredFields{
			Project:      "dummy-project",
			Subscription: "dummy-subscription",
		},
	})
	ce.Type = models.CloudExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy GCE description"
	ce.APIRoot = "https://dummy.api.kentik.com"
	ce.FlowDestination = "https://dummy.flow.kentik.com"
	ce.BGP = &models.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}

	ce, err = client.CloudExports.Create(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Create: %w", err)
	}

	fmt.Printf("Created GCE cloud export: %+v\n", ce)

	fmt.Println("Deleting GCE cloud export")
	err = client.CloudExports.Delete(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Delete: %w", err)
	}

	fmt.Println("Deleted GCE cloud export")
	return nil
}

func demonstrateCloudExportAPIWithIBM() error {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Creating IBM cloud export")
	ce := models.NewIBMCloudExport(models.CloudExportIBMRequiredFields{
		Name:          "example-ibm-export",
		PlanID:        "11467",
		CloudProvider: models.CloudProviderIBM,
		IBMProperties: models.IBMPropertiesRequiredFields{
			Bucket: "dummy-bucket",
		},
	})
	ce.Type = models.CloudExportTypeKentikManaged
	ce.Enabled = pointer.ToBool(true)
	ce.Description = "Dummy IBM description"
	ce.APIRoot = "https://dummy.api.kentik.com"
	ce.FlowDestination = "https://dummy.flow.kentik.com"
	ce.BGP = &models.BGPProperties{
		ApplyBGP:       pointer.ToBool(true),
		UseBGPDeviceID: "dummy-device-id",
		DeviceBGPType:  "dummy-device-bgp-type",
	}

	ce, err = client.CloudExports.Create(ctx, ce)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Create: %w", err)
	}

	fmt.Printf("Created IBM cloud export: %+v\n", ce)

	fmt.Println("Deleting IBM cloud export")
	err = client.CloudExports.Delete(ctx, ce.ID)
	if err != nil {
		return fmt.Errorf("client.CloudExports.Delete: %w", err)
	}

	fmt.Println("Deleted IBM cloud export")
	return nil
}
