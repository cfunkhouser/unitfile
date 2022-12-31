package unitfile

import "time"

// Unit describes generic details about a SystemD Unit which are not dependent
// on the type of the Unit.
//
// See https://www.freedesktop.org/software/systemd/man/systemd.unit.html#id-1.10
type Unit struct {
	Description                     string
	Documentation                   []string      `unit:"Documentation"`
	Wants                           []string      `unit:"Wants"`
	Requires                        []string      `unit:"Requires"`
	Requisite                       []string      `unit:"Requisite"`
	BindsTo                         []string      `unit:"BindsTo"`
	PartOf                          []string      `unit:"PartOf"`
	Upholds                         []string      `unit:"Upholds"`
	Conflicts                       []string      `unit:"Conflicts"`
	Before                          []string      `unit:"Before"`
	After                           []string      `unit:"After"`
	OnFailure                       []string      `unit:"OnFailure"`
	OnSuccess                       []string      `unit:"OnSuccess"`
	PropagatesReloadTo              []string      `unit:"PropagatesReloadTo"`
	ReloadPropagatedFrom            []string      `unit:"ReloadPropagatedFrom"`
	PropagatesStopTo                []string      `unit:"PropagatesStopTo"`
	StopPropagatedFrom              []string      `unit:"StopPropagatedFrom"`
	JoinsNamespaceOf                []string      `unit:"JoinsNamespaceOf"`
	RequiresMountsFor               []string      `unit:"RequiresMountsFor"`
	OnFailureJobMode                []string      `unit:"OnFailureJobMode"`
	IgnoreOnIsolate                 []string      `unit:"IgnoreOnIsolate"`
	StopWhenUnneeded                []string      `unit:"StopWhenUnneeded"`
	RefuseManualStart               []string      `unit:"RefuseManualStart"`
	RefuseManualStop                []string      `unit:"RefuseManualStop"`
	AllowIsolate                    []string      `unit:"AllowIsolate"`
	DefaultDependencies             []string      `unit:"DefaultDependencies"`
	CollectMode                     []string      `unit:"CollectMode"`
	FailureAction                   []string      `unit:"FailureAction"`
	SuccessAction                   []string      `unit:"SuccessAction"`
	FailureActionExitStatus         []string      `unit:"FailureActionExitStatus"`
	SuccessActionExitStatus         []string      `unit:"SuccessActionExitStatus"`
	JobTimeoutSec                   []string      `unit:"JobTimeoutSec"`
	JobRunningTimeoutSec            []string      `unit:"JobRunningTimeoutSec"`
	JobTimeoutAction                []string      `unit:"JobTimeoutAction"`
	JobTimeoutRebootArgument        []string      `unit:"JobTimeoutRebootArgument"`
	StartLimitInterval              time.Duration `unit:"StartLimitIntervalSec"`
	StartLimitBurst                 int           `unit:"StartLimitBurst"`
	StartLimitAction                []string      `unit:"StartLimitAction"`
	RebootArgument                  []string      `unit:"RebootArgument"`
	SourcePath                      []string      `unit:"SourcePath"`
	ConditionArchitecture           []string      `unit:"ConditionArchitecture"`
	ConditionFirmware               []string      `unit:"ConditionFirmware"`
	ConditionVirtualization         []string      `unit:"ConditionVirtualization"`
	ConditionHost                   []string      `unit:"ConditionHost"`
	ConditionKernelCommandLine      []string      `unit:"ConditionKernelCommandLine"`
	ConditionKernelVersion          []string      `unit:"ConditionKernelVersion"`
	ConditionCredential             []string      `unit:"ConditionCredential"`
	ConditionEnvironment            []string      `unit:"ConditionEnvironment"`
	ConditionSecurity               []string      `unit:"ConditionSecurity"`
	ConditionCapability             []string      `unit:"ConditionCapability"`
	ConditionACPower                []string      `unit:"ConditionACPower"`
	ConditionNeedsUpdate            []string      `unit:"ConditionNeedsUpdate"`
	ConditionFirstBoot              []string      `unit:"ConditionFirstBoot"`
	ConditionPathExists             []string      `unit:"ConditionPathExists"`
	ConditionPathExistsGlob         []string      `unit:"ConditionPathExistsGlob"`
	ConditionPathIsDirectory        []string      `unit:"ConditionPathIsDirectory"`
	ConditionPathIsSymbolicLink     []string      `unit:"ConditionPathIsSymbolicLink"`
	ConditionPathIsMountPoint       []string      `unit:"ConditionPathIsMountPoint"`
	ConditionPathIsReadWrite        []string      `unit:"ConditionPathIsReadWrite"`
	ConditionPathIsEncrypted        []string      `unit:"ConditionPathIsEncrypted"`
	ConditionDirectoryNotEmpty      []string      `unit:"ConditionDirectoryNotEmpty"`
	ConditionFileNotEmpty           []string      `unit:"ConditionFileNotEmpty"`
	ConditionFileIsExecutable       []string      `unit:"ConditionFileIsExecutable"`
	ConditionUser                   []string      `unit:"ConditionUser"`
	ConditionGroup                  []string      `unit:"ConditionGroup"`
	ConditionControlGroupController []string      `unit:"ConditionControlGroupController"`
	ConditionMemory                 []string      `unit:"ConditionMemory"`
	ConditionCPUs                   []string      `unit:"ConditionCPUs"`
	ConditionCPUFeature             []string      `unit:"ConditionCPUFeature"`
	ConditionOSRelease              []string      `unit:"ConditionOSRelease"`
	ConditionMemoryPressure         []string      `unit:"ConditionMemoryPressure"`
	ConditionCPUPressure            []string      `unit:"ConditionCPUPressure"`
	ConditionIOPressure             []string      `unit:"ConditionIOPressure"`
	AssertArchitecture              []string      `unit:"AssertArchitecture"`
	AssertVirtualization            []string      `unit:"AssertVirtualization"`
	AssertHost                      []string      `unit:"AssertHost"`
	AssertKernelCommandLine         []string      `unit:"AssertKernelCommandLine"`
	AssertKernelVersion             []string      `unit:"AssertKernelVersion"`
	AssertCredential                []string      `unit:"AssertCredential"`
	AssertEnvironment               []string      `unit:"AssertEnvironment"`
	AssertSecurity                  []string      `unit:"AssertSecurity"`
	AssertCapability                []string      `unit:"AssertCapability"`
	AssertACPower                   []string      `unit:"AssertACPower"`
	AssertNeedsUpdate               []string      `unit:"AssertNeedsUpdate"`
	AssertFirstBoot                 []string      `unit:"AssertFirstBoot"`
	AssertPathExists                []string      `unit:"AssertPathExists"`
	AssertPathExistsGlob            []string      `unit:"AssertPathExistsGlob"`
	AssertPathIsDirectory           []string      `unit:"AssertPathIsDirectory"`
	AssertPathIsSymbolicLink        []string      `unit:"AssertPathIsSymbolicLink"`
	AssertPathIsMountPoint          []string      `unit:"AssertPathIsMountPoint"`
	AssertPathIsReadWrite           []string      `unit:"AssertPathIsReadWrite"`
	AssertPathIsEncrypted           []string      `unit:"AssertPathIsEncrypted"`
	AssertDirectoryNotEmpty         []string      `unit:"AssertDirectoryNotEmpty"`
	AssertFileNotEmpty              []string      `unit:"AssertFileNotEmpty"`
	AssertFileIsExecutable          []string      `unit:"AssertFileIsExecutable"`
	AssertUser                      []string      `unit:"AssertUser"`
	AssertGroup                     []string      `unit:"AssertGroup"`
	AssertControlGroupController    []string      `unit:"AssertControlGroupController"`
	AssertMemory                    []string      `unit:"AssertMemory"`
	AssertCPUs                      []string      `unit:"AssertCPUs"`
	AssertCPUFeature                []string      `unit:"AssertCPUFeature"`
	AssertOSRelease                 []string      `unit:"AssertOSRelease"`
	AssertMemoryPressure            []string      `unit:"AssertMemoryPressure"`
	AssertCPUPressure               []string      `unit:"AssertCPUPressure"`
	AssertIOPressure                []string      `unit:"AssertIOPressure"`
}

// Service is a SystemD Unit which describes a process controlled and supervised
// by SystemD.
//
// See: https://www.freedesktop.org/software/systemd/man/systemd.service.html
type Service struct {
	Unit    `unit:"Unit"`
	Service ServiceSection `unit:"Service"`
	Install InstallSection `unit:"Install"`
}

// ServiceSection contains the options for the Service section of SystemD Units.
//
// See https://www.freedesktop.org/software/systemd/man/systemd.service.html#id-1.8
type ServiceSection struct {
	Type                     string        `unit:"Type"`
	ExitType                 string        `unit:"ExitType"`
	RemainAfterExit          string        `unit:"RemainAfterExit"`
	GuessMainPID             string        `unit:"GuessMainPID"`
	PIDFile                  string        `unit:"PIDFile"`
	BusName                  string        `unit:"BusName"`
	ExecStart                string        `unit:"ExecStart"`
	ExecStartPre             string        `unit:"ExecStartPre"`
	ExecStartPost            string        `unit:"ExecStartPost"`
	ExecCondition            string        `unit:"ExecCondition"`
	ExecReload               string        `unit:"ExecReload"`
	ExecStop                 string        `unit:"ExecStop"`
	ExecStopPost             string        `unit:"ExecStopPost"`
	RestartSec               string        `unit:"RestartSec"`
	TimeoutStart             time.Duration `unit:"TimeoutStartSec"`
	TimeoutStop              time.Duration `unit:"TimeoutStopSec"`
	TimeoutAbort             time.Duration `unit:"TimeoutAbortSec"`
	Timeout                  time.Duration `unit:"TimeoutSec"`
	TimeoutStartFailureMode  string        `unit:"TimeoutStartFailureMode"`
	TimeoutStopFailureMode   string        `unit:"TimeoutStopFailureMode"`
	RuntimeMax               time.Duration `unit:"RuntimeMaxSec"`
	RuntimeRandomizedExtra   time.Duration `unit:"RuntimeRandomizedExtraSec"`
	WatchdogDuration         time.Duration `unit:"WatchdogSec"`
	Restart                  string        `unit:"Restart"`
	SuccessExitStatus        string        `unit:"SuccessExitStatus"`
	RestartPreventExitStatus int           `unit:"RestartPreventExitStatus"`
	RestartForceExitStatus   string        `unit:"RestartForceExitStatus"`
	RootDirectoryStartOnly   string        `unit:"RootDirectoryStartOnly"`
	NonBlocking              string        `unit:"NonBlocking"`
	NotifyAccess             string        `unit:"NotifyAccess"`
	Sockets                  string        `unit:"Sockets"`
	FileDescriptorStoreMax   string        `unit:"FileDescriptorStoreMax"`
	USBFunctionDescriptors   string        `unit:"USBFunctionDescriptors"`
	USBFunctionStrings       string        `unit:"USBFunctionStrings"`
	OOMPolicy                string        `unit:"OOMPolicy"`
}

// InstallSection contains the options for the Install section of SystemD Units.
//
// See https://www.freedesktop.org/software/systemd/man/systemd.unit.html#id-1.12
type InstallSection struct {
	Alias           []string `unit:"Alias"`
	WantedBy        []string `unit:"WantedBy"`
	RequiredBy      []string `unit:"RequiredBy"`
	Also            []string `unit:"Also"`
	DefaultInstance []string `unit:"DefaultInstance"`
}
