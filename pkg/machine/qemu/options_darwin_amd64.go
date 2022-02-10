package qemu

var (
	QemuCommand = "qemu-system-x86_64"
)

func (v *MachineVM) addArchOptions() []string {
	opts := []string{
		"-machine", "q35,accel=hvf:tcg", "-cpu", "host",
		"-fsdev",
		"local,security_model=mapped,id=fsdev-fs0,multidevs=remap,path=/",
		"-device",
		"virtio-9p,id=fs0,fsdev=fsdev-fs0,mount_tag=fs0",
	}

	return opts
}

func (v *MachineVM) prepare() error {
	return nil
}

func (v *MachineVM) archRemovalFiles() []string {
	return []string{}
}
