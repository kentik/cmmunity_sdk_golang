//nolint:forbidigo
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/examples/demos"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

func main() {
	demos.Step("Create Kentik API client")
	client, err := demos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	demos.Step("Create a device")
	id := createDevice(client)

	demos.Step("Get device")
	getDevice(client, id)

	demos.Step("Delete device")
	deleteDevice(client, id)

	demos.Step("Get all devices")
	getAllDevices(client)

	demos.Step("Get all users")
	getAllUsers(client)

	demos.Step("Finished!")
}

func createDevice(client *kentikapi.Client) models.ID {
	device := models.NewDeviceDNS(models.DeviceDNSRequiredFields{
		DeviceName:       "testapi_dns_awssubnet",
		DeviceSubType:    models.DeviceSubtypeAWSSubnet,
		DeviceSampleRate: 1,
		PlanID:           models.ID("11466"),
		CDNAttr:          models.CDNAttributeYes,
	})
	device.SiteID = models.IDPtr("8483")
	device.DeviceBGPFlowSpec = pointer.ToBool(true)

	createdDevice, err := client.Devices.Create(context.Background(), *device)
	demos.ExitOnError(err)
	fmt.Printf("Successfully created device, ID = %v\n", createdDevice.ID)

	return createdDevice.ID
}

func getDevice(client *kentikapi.Client, id models.ID) {
	fmt.Printf("Retrieving device of ID = %v\n", id)
	device, err := client.Devices.Get(context.Background(), id)
	demos.ExitOnError(err)
	demos.PrettyPrint(device)
}

func deleteDevice(client *kentikapi.Client, id models.ID) {
	fmt.Printf("Deleting device of ID = %v\n", id)
	err := client.Devices.Delete(context.Background(), id) // archive
	demos.ExitOnError(err)
	err = client.Devices.Delete(context.Background(), id) // delete
	demos.ExitOnError(err)
	fmt.Println("Successful")
}

func getAllDevices(client *kentikapi.Client) {
	devices, err := client.Devices.GetAll(context.Background())
	demos.ExitOnError(err)
	demos.PrettyPrint(devices)
	fmt.Printf("Total devices: %d\n", len(devices))
}

func getAllUsers(client *kentikapi.Client) {
	users, err := client.Users.GetAll(context.Background())
	demos.ExitOnError(err)
	demos.PrettyPrint(users)
	fmt.Printf("Total users: %d\n", len(users))
}
