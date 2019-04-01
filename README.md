# hvinfo

hvinfo dumps the `HyperV` feature support bits as documented in the [HyperV Top Level Functional Specifications](https://github.com/MicrosoftDocs/Virtualization-Documentation/blob/live/tlfs/Hypervisor%20Top%20Level%20Functional%20Specification%20v5.0C.pdf).

If you don't know what any of this does mean, you don't need `hvinfo` :)

hvinfo is (C) 2019 Red Hat Inc and licensed under the MIT license.

## compile

Cross compile on linux for windows:
```bash
GOOS=windows GOARCH=amd64 go build -o binaries/hvinfo.exe hvinfo.go
```

Should you need to compile for linux (negative test?):
```bash
go build -o binaries/hvinfo hvinfo.go
```

## usage

`hvinfo` takes no arguments and just emits JSON outputs.
`hvinfo` does NOT require admin privileges to run.
Being a commandline tool, is hard to give a "screenshot", so we will just show a sample output.

```
Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

PS C:\Users\admin> .\hvinfo.exe
{
  "HyperVsupport": true,
  "Features": {
    "GuestDebugging": false,
    "PerformanceMonitor": false,
    "PCPUDynamicPartitioningEvents": true,
    "HypercallInputParamsXMM": false,
    "VirtualGuestIdleState": false,
    "HypervisorSleepState": false,
    "NUMADistanceQuery": false,
    "TimerFrequenciesQuery": false,
    "SytheticMCEInjection": false,
    "GuestCrashMSR": false,
    "DebugMSR": false,
    "NPIEP": false,
    "DisableHypervisorAvailable": false,
    "ExtendedGvaRangesForFlushVirtualAddressList": false,
    "HypercallOutputReturnXMM": false,
    "SintPollingMode": false,
    "HypercallMsrLock": false,
    "UseDirectSyntheticTimers": false
  },
  "Recommendations": {
    "HypercallAddressSpaceSwitch": false,
    "HypercallLocalTLBFlush": false,
    "HypercallRemoteTLBFlush": false,
    "MSRAPICRegisters": true,
    "MSRSystemReset": false,
    "RelaxedTiming": true,
    "DMARemapping": false,
    "InterruptRemapping": false,
    "X2APICMSR": false,
    "DeprecatingAutoEOI": false,
    "SyntheticClusterIPI": false,
    "ExProcessorMasks": false,
    "Nested": false,
    "INTForMBECSyscalls": false,
    "NestedEVMCS": false,
    "SyncedTimeline": false,
    "DirectLocalFlushEntire": false,
    "NoNonArchitecturalCoreSharing": false,
    "SpinlockRetries": 8191
  }
}
```

## bugs

Trivial bugs and typos in the otherwise simple but repetitive implementation of `hvinfo` may end up in incorrect reporting.
Please send bugs to fromani at redhat dot com.
Always include:
- the version (git commit hash from which the tool is compiled)
- the output of the tool
- the errors: what you should see and you don't, or the other way around
