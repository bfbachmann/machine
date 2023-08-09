package errdriver

import (
	"encoding/json"
	"fmt"

	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/mcnflag"
	"github.com/rancher/machine/libmachine/state"
)

type Driver struct {
	Name string
}

type NotLoadable struct {
	Name string
}

func (e NotLoadable) Error() string {
	return fmt.Sprintf("Driver %q not found. Do you have the plugin binary accessible in your PATH?", e.Name)
}

func NewDriver(Name string) drivers.Driver {
	return &Driver{
		Name: Name,
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "not-found"
}

func (d *Driver) PreCreateCheck() error {
	return NotLoadable{d.Name}
}

func (d *Driver) GetFlags() []mcnflag.Flag {
	return nil
}

// LoadConfigFromJSON loads driver config from JSON.
func (d *Driver) LoadConfigFromJSON(data []byte) error {
	if err := json.Unmarshal(data, &d); err != nil {
		return fmt.Errorf("error unmarshalling driver config from JSON: %w", err)
	}

	return nil
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	return NotLoadable{d.Name}
}

func (d *Driver) GetURL() (string, error) {
	return "", NotLoadable{d.Name}
}

func (d *Driver) GetMachineName() string {
	return d.Name
}

func (d *Driver) GetIP() (string, error) {
	return "1.2.3.4", NotLoadable{d.Name}
}

func (d *Driver) GetSSHHostname() (string, error) {
	return "", NotLoadable{d.Name}
}

func (d *Driver) GetSSHKeyPath() string {
	return ""
}

func (d *Driver) GetSSHPort() (int, error) {
	return 0, NotLoadable{d.Name}
}

func (d *Driver) GetSSHUsername() string {
	return ""
}

func (d *Driver) GetState() (state.State, error) {
	return state.Error, NotLoadable{d.Name}
}

func (d *Driver) Create() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Remove() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Start() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Stop() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Restart() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Kill() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Upgrade() error {
	return NotLoadable{d.Name}
}
