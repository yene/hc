package container

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"testing"
)

var info = model.Info{
	Name:         "Accessory1",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Accessory",
}

func TestContainer(t *testing.T) {
	acc1 := accessory.New(info, accessory.TypeOther)
	info.Name = "Accessory2"
	acc2 := accessory.New(info, accessory.TypeOther)

	if is, want := acc1.GetID(), model.InvalidID; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := acc2.GetID(), model.InvalidID; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c := NewContainer()
	c.AddAccessory(acc1)
	c.AddAccessory(acc2)

	if is, want := len(c.Accessories), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := acc1.GetID(); x == model.InvalidID {
		t.Fatal(x)
	}
	if x := acc2.GetID(); x == model.InvalidID {
		t.Fatal(x)
	}
	if acc1.GetID() == acc2.GetID() {
		t.Fatal("equal ids not allowed")
	}

	c.RemoveAccessory(acc2)

	if is, want := len(c.Accessories), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestAccessoryCount(t *testing.T) {
	accessory := accessory.New(info, accessory.TypeOther)
	c := NewContainer()
	c.AddAccessory(accessory)

	if is, want := len(c.Accessories), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.RemoveAccessory(accessory)

	if is, want := len(c.Accessories), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestAccessoryType(t *testing.T) {
	a1 := accessory.New(info, accessory.TypeLightBulb)
	a2 := accessory.New(info, accessory.TypeSwitch)

	c := NewContainer()
	c.AddAccessory(a1)

	if is, want := c.AccessoryType(), accessory.TypeLightBulb; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.AddAccessory(a2)

	if is, want := c.AccessoryType(), accessory.TypeBridge; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
