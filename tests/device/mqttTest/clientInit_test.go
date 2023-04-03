package mqttTest

import (
	"fmt"
	"github.com/i-Things/things/shared/domain/deviceAuth"
	"testing"
)

func TestGetMqttClient(t *testing.T) {
	_, err := GetMqttClient(&ClientInfo{
		MqttBrokers:  []string{"tcp://106.15.225.172:1883"},
		ProductID:    "254pwnKQsvK",
		DeviceName:   "test5",
		DeviceSecret: "6skuocmEga94+OhVYRGWUphWlX0=",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestAssert(t *testing.T) {
	clientID, userName, pwd := deviceAuth.GenSecretDeviceInfo(deviceAuth.HmacSha256, "2625ope58nS", "testDevA", "b6nSU9U3P83G4tLLfsZ9hIUi+vw=")
	fmt.Printf("get id:%s,uname:%s,pwd:%s\n", clientID, userName, pwd)
	pwdInfo, _ := deviceAuth.NewPwdInfo(pwd)
	e := pwdInfo.CmpPwd(userName, "b6nSU9U3P83G4tLLfsZ9hIUi+vw=")
	if e == nil {
		//88a4124084ea4b289282cba19fce7f73d4665ab40899b747dd73431b228f5bab
	} else {
		fmt.Printf("get err:%s\n", e.Error())
	}
}
