package defaults

// DefaultEdenTemplate is configuration template for Eden
const DefaultEdenTemplate = `#config is generated by eden
adam:
    #tag on adam container to pull
    tag: '{{parse "adam.tag"}}'

    #location of adam
    dist: '{{parse "adam.dist"}}'

    #port of adam
    port: {{parse "adam.port"}}

    #domain of adam
    domain: '{{parse "adam.domain"}}'

    #ip of adam for EVE access
    eve-ip: '{{parse "adam.eve-ip"}}'

    #ip of adam for EDEN access
    ip: '{{parse "adam.ip"}}'

    redis:
      #host of adam's redis for EDEN access
      eden: '{{parse "adam.redis.eden"}}'
      #host of adam's redis for ADAM access
      adam: '{{parse "adam.redis.adam"}}'

    #force adam rebuild
    force: {{parse "adam.force"}}

    #certificate for communication with adam
    ca: '{{parse "adam.ca"}}'

    #use remote adam
    remote:
        enabled: {{parse "adam.remote.enabled"}}

        #load logs and info from redis instead of http stream
        redis: {{parse "adam.remote.redis"}}

    #use v1 api
    v1: {{parse "adam.v1"}}

    caching:
        enabled: {{parse "adam.caching.enabled"}}

        #caching logs and info to redis instead of local
        redis: {{parse "adam.caching.redis"}}

        #prefix for directory/redis stream
        prefix: '{{parse "adam.caching.prefix"}}'

eve:
    #name
    name: '{{parse "eve.name"}}'

    #devmodel
    devmodel: '{{parse "eve.devmodel"}}'

    #devmodel file overwrite
    devmodelfile: '{{parse "eve.devmodelfile"}}'

    #Path to a file with JSON-formatted device config (EdgeDevConfig, but mostly just networking),
    #used to bootstrap device (i.e. establish connectivity with the controller and onboard).
    #The config will be reformatted to binary proto, signed and embedded into the EVE image.
    #Note: for legacy override.json use "eve-config-dir" arg of "eden setup".
    bootstrap-file: '{{parse "eve.bootstrap-file"}}'

    #Path to a file with JSON-formatted network config (modeled by DevicePortConfig
    #struct from EVE repo), applied in runtime using a specially formatted USB stick.
    #Also known as "usb.json".
    #Typically used for device bootstrapping as a second step after EVE installation.
    #This is a legacy method soon to be replaced by EdgeDevConfig-based bootstrapping.
    usbnetconf-file: '{{parse "eve.usbnetconf-file"}}'

    #EVE arch (amd64/arm64)
    arch: '{{parse "eve.arch"}}'

    #EVE platform (imx8mk_evk, opi3_lts and others)
    platform: '{{parse "eve.platform"}}'

    #EVE os (linux/darwin)
    os: '{{parse "eve.os"}}'

    #EVE acceleration (set to false if you have problems with qemu)
    accel: {{parse "eve.accel"}}

    #variant of hypervisor of EVE (kvm/xen)
    hv: '{{parse "eve.hv"}}'

    #serial number in SMBIOS
    serial: '{{parse "eve.serial"}}'

    #onboarding certificate of EVE to put into adam
    cert: '{{parse "eve.cert"}}'

    #device certificate of EVE to put into adam
    device-cert: '{{parse "eve.device-cert"}}'

    #EVE pid file
    pid: '{{parse "eve.pid"}}'

    #EVE log file
    log: '{{parse "eve.log"}}'

    #EVE firmware
    firmware: {{parse "eve.firmware"}}

    #eve repo used in clone mode (eden.download = false)
    repo: '{{parse "eve.repo"}}'

    #eve registry to use
    registry: '{{parse "eve.registry"}}'

    #eve tag
    tag: '{{parse "eve.tag"}}'

    #port forwarding for EVE VM [(HOST:EVE)] when running without Eden-SDN
    hostfwd: {{  parsemap "eve.hostfwd" }}

    #location of eve directory
    dist: '{{parse "eve.dist"}}'

    #file to save qemu config
    qemu-config: '{{parse "eve.qemu-config"}}'

    #uuid of EVE to use in cert
    uuid: '{{parse "eve.uuid"}}'

    #live image of EVE
    image-file: '{{parse "eve.image-file"}}'

    #custom pre-built EVE installer
    #It is assumed that the installer is configured for the zedcloud controller
    #(or another 3rd party controller, but not Adam)
    custom-installer:
        #path to the installer image file
        path: '{{parse "eve.custom-installer.path"}}'

        #format of the installer image (should be "raw" or "iso")
        format: '{{parse "eve.custom-installer.format"}}'

    #dtb directory of EVE
    dtb-part: '{{parse "eve.dtb-part"}}'

    #config part of EVE
    config-part: '{{parse "eve.config-part"}}'

    #is EVE remote or local
    remote: {{parse "eve.remote"}}

    #EVE address for access from Eden
    remote-addr: '{{parse "eve.remote-addr"}}'

    #min level of logs saved in files on device
    log-level: '{{parse "eve.log-level"}}'

    #min level of logs sent to controller
    adam-log-level: '{{parse "eve.adam-log-level"}}'

    #port for telnet (console access)
    telnet-port: {{parse "eve.telnet-port"}}

    #ssid for wifi
    ssid: '{{parse "eve.ssid"}}'

    #cpu count
    cpu: {{parse "eve.cpu"}}

    #memory (MB)
    ram: {{parse "eve.ram"}}

    #disk (MB)
    disk: {{parse "eve.disk"}}

    #tpm
    tpm: {{parse "eve.tpm"}}

    #additional disks count
    disks: {{parse "eve.disks"}}

    #configuration specific to QEMU-emulated device
    qemu:
        #port for QEMU Monitor
        monitor-port: {{parse "eve.qemu.monitor-port"}}

        #base port for socket-based ethernet interfaces used in QEMU
        netdev-socket-port: {{parse "eve.qemu.netdev-socket-port"}}

eden:
    #root directory of eden
    root: '{{parse "eden.root"}}'
    #directory with tests
    tests: '{{parse "eden.tests"}}'
    images:
        #directory to save images
        dist: '{{parse "eden.images.dist"}}'

    #download eve instead of build
    download: {{parse "eden.download"}}

    #eserver is tool for serve images
    eserver:
        #ip (domain name) of eserver for EVE access
        eve-ip: '{{parse "eden.eserver.eve-ip"}}'

        #ip of eserver for EDEN access
        ip: '{{parse "eden.eserver.ip"}}'

        #port for eserver
        port: {{parse "eden.eserver.port"}}

        #tag of eserver container
        tag: '{{parse "eden.eserver.tag"}}'

        #force eserver rebuild
        force: {{parse "eden.eserver.force"}}

    #eclient is tool we use in tests
    eclient:
        #tag of eclient container
        tag: '{{parse "eden.eclient.tag"}}'
        #image of eclient container
        image: '{{parse "eden.eclient.image"}}'

    #directory to save certs
    certs-dist: '{{parse "eden.certs-dist"}}'

    #directory to save binaries
    bin-dist: '{{parse "eden.bin-dist"}}'

    #ssh-key to put into EVE
    ssh-key: '{{parse "eden.ssh-key"}}'

    #eden binary
    eden-bin: '{{parse "eden.eden-bin"}}'

    #test binary
    test-bin: '{{parse "eden.test-bin"}}'

    #test scenario
    test-scenario: '{{parse "eden.test-scenario"}}'

gcp:
    #path to the key to interact with gcp
    key: '{{parse "gcp.key"}}'

packet:
    #path to the key to interact with packet
    key: '{{parse "packet.key"}}'

redis:
    #port for access redis
    port: {{parse "redis.port"}}

    #tag for redis image
    tag: '{{parse "redis.tag"}}'

    #directory to use for redis persistence
    dist: '{{parse "redis.dist"}}'

registry:
    #port for registry access
    port: {{parse "registry.port"}}

    #tag for registry image
    tag: '{{parse "registry.tag"}}'

    #ip of registry for EDEN access
    ip: '{{parse "registry.ip"}}'

    # dist path to store registry data
    dist: '{{parse "registry.dist"}}'

sdn:
    #disable SDN
    disable: '{{parse "sdn.disable"}}'

    #directory with SDN source code
    source-dir: '{{parse "sdn.source-dir"}}'

    #directory where to put generated SDN-related config files
    config-dir: '{{parse "sdn.config-dir"}}'

    #live image of SDN
    image-file: '{{parse "sdn.image-file"}}'

    #path to linuxkit binary used to build SDN VM
    linuxkit-bin: '{{parse "sdn.linuxkit-bin"}}'

    #CPU count for SDN VM
    cpu: {{parse "sdn.cpu"}}

    #memory (MB) for SDN VM
    ram: {{parse "sdn.ram"}}

    #SDN pid file
    pid: '{{parse "sdn.pid"}}'

    #SDN file where console output is logged
    #Not as useful as logs from the SDN mgmt agent (get with: eden sdn logs)
    console-log: '{{parse "sdn.console-log"}}'

    #port for telnet (console access) to SDN VM
    telnet-port: {{parse "sdn.telnet-port"}}

    #port for SSH access to SDN VM
    ssh-port: {{parse "sdn.ssh-port"}}

    #port for access to the management agent running inside SDN VM
    mgmt-port: {{parse "sdn.mgmt-port"}}

    #path to JSON file with network model to apply into SDN
    #leave empty for default network model
    network-model: '{{parse "sdn.network-model"}}'
`

// DefaultQemuTemplate is configuration template for qemu
const DefaultQemuTemplate = `#qemu config file generated by eden
{{- if .Firmware }}
{{ $firmwareLength := len .Firmware }}{{ if eq $firmwareLength 1 }}
[machine]
  firmware = "{{ index .Firmware 0 }}"
{{- else if eq $firmwareLength 2 }}
[drive]
  if = "pflash"
  format = "raw"
  unit = "0"
  readonly = "on"
  file = "{{ index .Firmware 0 }}"

[drive]
  if = "pflash"
  format = "raw"
  unit = "1"
  file = "{{ index .Firmware 1 }}"
{{end}}
{{end}}
{{if .DTBDrive }}
[drive]
  file = "fat:rw:{{ .DTBDrive }}"
  format = "vvfat"
  label = "QEMU_DTB""
{{end}}
[rtc]
  base = "utc"
  clock = "rt"

[global]
  driver = "ICH9-LPC"
  property = "noreboot"
  value = "false"

[memory]
  size = "{{ .MemoryMB }}"

[smp-opts]
  cpus = "{{ .CPUs }}"

[device "usb"]
  driver = "qemu-xhci"

{{- if .USBTablets -}}
{{ range $i := .USBTablets }}
[device]
  driver = "usb-tablet"
{{ end }}
{{- end -}}

{{- if .USBSerials -}}
{{ range $i := .USBSerials }}
[chardev "charserial{{ $i }}"]
  backend = "pty"

[device "serial{{ $i }}"]
  driver = "usb-serial"
  chardev = "charserial{{ $i }}"
{{ end }}
{{- end -}}

{{ range .Disks }}
[drive]
  format = "qcow2"
  file = "{{.}}"
{{ end }}
`

// ParallelsDiskTemplate is template for disk annotation of parallels
const ParallelsDiskTemplate = `<?xml version='1.0' encoding='UTF-8'?>
<Parallels_disk_image Version="1.0">
    <Disk_Parameters>
        <Disk_size>{{ .DiskSize }}</Disk_size>
        <Cylinders>{{ .Cylinders }}</Cylinders>
        <PhysicalSectorSize>512</PhysicalSectorSize>
        <Heads>16</Heads>
        <Sectors>32</Sectors>
        <Padding>0</Padding>
        <Encryption>
            <Engine>{00000000-0000-0000-0000-000000000000}</Engine>
            <Data></Data>
        </Encryption>
        <UID>{{ .UID }}</UID>
        <Name>eve</Name>
        <Miscellaneous>
            <CompatLevel>level2</CompatLevel>
            <Bootable>1</Bootable>
            <SuspendState>0</SuspendState>
        </Miscellaneous>
    </Disk_Parameters>
    <StorageData>
        <Storage>
            <Start>0</Start>
            <End>{{ .DiskSize }}</End>
            <Blocksize>2048</Blocksize>
            <Image>
                <GUID>{{ .SnapshotUID }}</GUID>
                <Type>Compressed</Type>
                <File>live.0.{{ .SnapshotUID }}.hds</File>
            </Image>
        </Storage>
    </StorageData>
    <Snapshots>
        <Shot>
            <GUID>{{ .SnapshotUID }}</GUID>
            <ParentGUID>{00000000-0000-0000-0000-000000000000}</ParentGUID>
        </Shot>
    </Snapshots>
</Parallels_disk_image>`

// DefaultEdenTemplate is configuration template for Eden
const DefaultActivateShTemplate = `# This file must be used with "source bin/activate" *from bash*
# you cannot run it directly

if [ "${BASH_SOURCE-}" = "$0" ]; then
    echo "You must source this script: \$ source $0" >&2
    exit 33
fi

eden_deactivate () {
    # reset old environment variables
    # ! [ -z ${VAR+_} ] returns true if VAR is declared at all
    if ! [ -z "${_OLD_EDEN_PATH:+_}" ] ; then
        PATH="$_OLD_EDEN_PATH"
        export PATH
        unset _OLD_EDEN_PATH
    fi

    # This should detect bash and zsh, which have a hash command that must
    # be called to get it to forget past commands.  Without forgetting
    # past commands the $PATH changes we made may not be respected
    if [ -n "${BASH-}" ] || [ -n "${ZSH_VERSION-}" ] ; then
        hash -r 2>/dev/null
    fi

    if ! [ -z "${_OLD_EDEN_PS1+_}" ] ; then
        PS1="$_OLD_EDEN_PS1"
        export PS1
        unset _OLD_EDEN_PS1
    fi

    unset EDEN_HOME
    if [ ! "${1-}" = "nondestructive" ] ; then
    # Self destruct!
        unset -f eden_deactivate
        unset -f eden_config
        unset -f eden-config
        unset -f eden+config
    fi
}

eden_config () {
    if [ $# -eq 0 ]
    then
        echo Usage: eden_config config
        return
    fi

    eden config set $1
    PS1="EDEN-$(eden config get)_${_OLD_EDEN_PS1-}"
}

eden+config () {
    if [ $# -eq 0 ]
    then
        echo Usage: eden+config config
        return
    fi

    cd $(eden config get --key eden.root)/..
    eden config add $1
    cd -
}

eden-config () {
    if [ $# -eq 0 ]
    then
        echo Usage: eden-config config
        return
    fi

    eden config delete $1
    eden_config default
}

# unset irrelevant variables
eden_deactivate nondestructive

EDEN_HOME={{.Eden.Root}}
EDEN_BIN={{.Eden.BinDir}}
export EDEN_HOME

_OLD_EDEN_PATH="$PATH"
PATH="$EDEN_BIN:$PATH"
export PATH

if [ -z "${EDEN_HOME_DISABLE_PROMPT-}" ] ; then
    _OLD_EDEN_PS1="${PS1-}"
    PS1="EDEN-$(eden config get)_${PS1-}"
    export PS1
fi

# This should detect bash and zsh, which have a hash command that must
# be called to get it to forget past commands.  Without forgetting
# past commands the $PATH changes we made may not be respected
if [ -n "${BASH-}" ] || [ -n "${ZSH_VERSION-}" ] ; then
    hash -r 2>/dev/null
fi
`

const DefaultActivateCshTemplate = `# This file must be used with "source bin/activate.csh" *from csh*.
# You cannot run it directly.

set newline='\
'

alias eden_deactivate 'test $?_OLD_EDEN_PATH != 0 && setenv PATH "$_OLD_EDEN_PATH:q" && unset _OLD_EDEN_PATH; rehash; test $?_OLD_EDEN_PROMPT != 0 && set prompt="$_OLD_EDEN_PROMPT:q" && unset _OLD_EDEN_PROMPT; unsetenv EDEN_HOME; test "\!:*" != "nondestructive" && unalias eden_deactivate && unalias eden_config && unalias eden+config && unalias eden-config'

alias eden_config 'eden config set \!:1 && set prompt="EDEN-$(eden config get)_$_OLD_EDEN_PROMPT:q"'

alias eden+config 'cd $(eden config get --key eden.root)/..; eden config add \!:1; cd -'
alias eden-config 'eden config delete \!:1; eden_config default'

# Unset irrelevant variables.
eden_deactivate nondestructive

setenv EDEN_HOME "{{.Eden.Root}}"

set _OLD_EDEN_PATH="$PATH:q"
setenv PATH "{{.Eden.BinDir}}:$PATH:q"

if ( $?EDEN_DISABLE_PROMPT ) then
    if ( $EDEN_DISABLE_PROMPT == "" ) then
        set do_prompt = "1"
    else
        set do_prompt = "0"
    endif
else
    set do_prompt = "1"
endif

if ( $do_prompt == "1" ) then
    # Could be in a non-interactive environment,
    # in which case, $prompt is undefined and we wouldn't
    # care about the prompt anyway.
    if ( $?prompt ) then
        set _OLD_EDEN_PROMPT="$prompt:q"
        if ( "$prompt:q" =~ *"$newline:q"* ) then
            :
        else
            set prompt = "eden-$(eden config get)_$prompt:q"
        endif
    endif
endif

unset do_prompt

rehash`
