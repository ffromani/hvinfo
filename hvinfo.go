/*
 * Copyright (c) 2019, Red Hat Inc
 * License: MIT
 */

package main

import (
	"encoding/json"
	"os"

	"github.com/fromanirh/cpuid"
)

type HypervFeatures struct {
	GuestDebugging                              bool
	PerformanceMonitor                          bool
	PCPUDynamicPartitioningEvents               bool
	HypercallInputParamsXMM                     bool
	VirtualGuestIdleState                       bool
	HypervisorSleepState                        bool
	NUMADistanceQuery                           bool
	TimerFrequenciesQuery                       bool
	SytheticMCEInjection                        bool
	GuestCrashMSR                               bool
	DebugMSR                                    bool
	NPIEP                                       bool
	DisableHypervisorAvailable                  bool
	ExtendedGvaRangesForFlushVirtualAddressList bool
	HypercallOutputReturnXMM                    bool
	SintPollingMode                             bool
	HypercallMsrLock                            bool
	UseDirectSyntheticTimers                    bool
}

func GetHypervFeatures(enabled bool) *HypervFeatures {
	if !enabled {
		return nil
	}
	return &HypervFeatures{
		GuestDebugging:                cpuid.HypervHasFeature(cpuid.HYPERV_GUEST_DEBUGGING),
		PerformanceMonitor:            cpuid.HypervHasFeature(cpuid.HYPERV_PERFORMANCE_MONITOR),
		PCPUDynamicPartitioningEvents: cpuid.HypervHasFeature(cpuid.HYPERV_PCPU_DYNAMIC_PARTITIONING_EVENTS),
		HypercallInputParamsXMM:       cpuid.HypervHasFeature(cpuid.HYPERV_HYPERCALL_INPUT_PARAMETERS_XMM),
		VirtualGuestIdleState:         cpuid.HypervHasFeature(cpuid.HYPERV_VIRTUAL_GUEST_IDLE_STATE),
		HypervisorSleepState:          cpuid.HypervHasFeature(cpuid.HYPERV_HYPERVISOR_SLEEP_STATE),
		NUMADistanceQuery:             cpuid.HypervHasFeature(cpuid.HYPERV_NUMA_DISTANCE_QUERY),
		TimerFrequenciesQuery:         cpuid.HypervHasFeature(cpuid.HYPERV_TIMER_FREQUENCIES_QUERY),
		SytheticMCEInjection:          cpuid.HypervHasFeature(cpuid.HYPERV_SYTHETIC_MCE_INJECTION),
		GuestCrashMSR:                 cpuid.HypervHasFeature(cpuid.HYPERV_GUEST_CRASH_MSR),
		DebugMSR:                      cpuid.HypervHasFeature(cpuid.HYPERV_DEBUG_MSR),
		NPIEP:                         cpuid.HypervHasFeature(cpuid.HYPERV_NPIEP),
		DisableHypervisorAvailable:    cpuid.HypervHasFeature(cpuid.HYPERV_DISABLE_HYPERVISOR),
		ExtendedGvaRangesForFlushVirtualAddressList: cpuid.HypervHasFeature(cpuid.HYPERV_EXTENDED_GVA_RANGES_FOR_FLUSH_VIRTUAL_ADDRESS_LIST),
		HypercallOutputReturnXMM:                    cpuid.HypervHasFeature(cpuid.HYPERV_HYPERCALL_OUTPUT_RETURN_XMM),
		SintPollingMode:                             cpuid.HypervHasFeature(cpuid.HYPERV_SINT_POLLING_MODE),
		HypercallMsrLock:                            cpuid.HypervHasFeature(cpuid.HYPERV_HYPERCALL_MSR_LOCK),
		UseDirectSyntheticTimers:                    cpuid.HypervHasFeature(cpuid.HYPERV_USE_DIRECT_SYNTHETIC_TIMERS),
	}
}

type HypervPartitionPrivileges struct {
	AccessVpRunTimeReg              bool
	AccessPartitionReferenceCounter bool
	AccessSynicRegs                 bool
	AccessSyntheticTimerRegs        bool
	AccessIntrCtrlRegs              bool
	AccessHypercallMsrs             bool
	AccessVpIndex                   bool
	AccessResetReg                  bool
	AccessStatsReg                  bool
	AccessPartitionReferenceTsc     bool
	AccessGuestIdleReg              bool
	AccessFrequencyRegs             bool
	AccessDebugRegs                 bool
	AccessReenlightenmentControls   bool

	CreatePartitions         bool
	AccessPartitionId        bool
	AccessMemoryPool         bool
	PostMessages             bool
	SignalEvents             bool
	CreatePort               bool
	ConnectPort              bool
	AccessStats              bool
	Debugging                bool
	CpuManagement            bool
	AccessVSM                bool
	AccessVpRegisters        bool
	EnableExtendedHypercalls bool
	StartVirtualProcessor    bool
}

func GetHypervPrivileges(enabled bool) *HypervPartitionPrivileges {
	if !enabled {
		return nil
	}
	return &HypervPartitionPrivileges{
		AccessVpRunTimeReg:              cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessVpRunTimeRegs),
		AccessPartitionReferenceCounter: cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessPartitionReferenceCounter),
		AccessSynicRegs:                 cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessSynicRegs),
		AccessSyntheticTimerRegs:        cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessSyntheticTimerRegs),
		AccessIntrCtrlRegs:              cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessIntrCtrlRegs),
		AccessHypercallMsrs:             cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessHypercallMsrs),
		AccessVpIndex:                   cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessVpIndex),
		AccessResetReg:                  cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessResetReg),
		AccessStatsReg:                  cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessStatsReg),
		AccessPartitionReferenceTsc:     cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessPartitionReferenceTsc),
		AccessGuestIdleReg:              cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessGuestIdleReg),
		AccessFrequencyRegs:             cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessFrequencyRegs),
		AccessDebugRegs:                 cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessDebugRegs),
		AccessReenlightenmentControls:   cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessReenlightenmentControls),

		CreatePartitions:         cpuid.HypervHasPrivilege(cpuid.HYPERV_CreatePartitions),
		AccessPartitionId:        cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessPartitionId),
		AccessMemoryPool:         cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessMemoryPool),
		PostMessages:             cpuid.HypervHasPrivilege(cpuid.HYPERV_PostMessages),
		SignalEvents:             cpuid.HypervHasPrivilege(cpuid.HYPERV_SignalEvents),
		CreatePort:               cpuid.HypervHasPrivilege(cpuid.HYPERV_CreatePort),
		ConnectPort:              cpuid.HypervHasPrivilege(cpuid.HYPERV_ConnectPort),
		AccessStats:              cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessStats),
		Debugging:                cpuid.HypervHasPrivilege(cpuid.HYPERV_Debugging),
		CpuManagement:            cpuid.HypervHasPrivilege(cpuid.HYPERV_CpuManagement),
		AccessVSM:                cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessVSM),
		AccessVpRegisters:        cpuid.HypervHasPrivilege(cpuid.HYPERV_AccessVpRegisters),
		EnableExtendedHypercalls: cpuid.HypervHasPrivilege(cpuid.HYPERV_EnableExtendedHypercalls),
		StartVirtualProcessor:    cpuid.HypervHasPrivilege(cpuid.HYPERV_StartVirtualProcessor),
	}
}

type HypervRecommendations struct {
	HypercallAddressSpaceSwitch   bool
	HypercallLocalTLBFlush        bool
	HypercallRemoteTLBFlush       bool
	MSRAPICRegisters              bool
	MSRSystemReset                bool
	RelaxedTiming                 bool
	DMARemapping                  bool
	InterruptRemapping            bool
	X2APICMSR                     bool
	DeprecatingAutoEOI            bool
	SyntheticClusterIPI           bool
	ExProcessorMasks              bool
	Nested                        bool
	INTForMBECSyscalls            bool
	NestedEVMCS                   bool
	SyncedTimeline                bool
	DirectLocalFlushEntire        bool
	NoNonArchitecturalCoreSharing bool
	SpinlockRetries               uint64
}

func GetHypervRecommendations(enabled bool) *HypervRecommendations {
	if !enabled {
		return nil
	}
	return &HypervRecommendations{
		HypercallAddressSpaceSwitch:   cpuid.HypervHasRecommendation(cpuid.HYPERV_HYPERCALL_ADDRESS_SPACE_SWITCH),
		HypercallLocalTLBFlush:        cpuid.HypervHasRecommendation(cpuid.HYPERV_HYPERCALL_LOCAL_TLB_FLUSH),
		HypercallRemoteTLBFlush:       cpuid.HypervHasRecommendation(cpuid.HYPERV_HYPERCALL_REMOTE_TLB_FLUSH),
		MSRAPICRegisters:              cpuid.HypervHasRecommendation(cpuid.HYPERV_MSR_APIC_REGISTERS),
		MSRSystemReset:                cpuid.HypervHasRecommendation(cpuid.HYPERV_MSR_SYSTEM_RESET),
		RelaxedTiming:                 cpuid.HypervHasRecommendation(cpuid.HYPERV_RELAXED_TIMING),
		DMARemapping:                  cpuid.HypervHasRecommendation(cpuid.HYPERV_DMA_REMAPPING),
		InterruptRemapping:            cpuid.HypervHasRecommendation(cpuid.HYPERV_INTERRUPT_REMAPPING),
		X2APICMSR:                     cpuid.HypervHasRecommendation(cpuid.HYPERV_x2APIC_MSR),
		DeprecatingAutoEOI:            cpuid.HypervHasRecommendation(cpuid.HYPERV_DEPRECATING_AUTOEOI),
		SyntheticClusterIPI:           cpuid.HypervHasRecommendation(cpuid.HYPERV_SYNTHETIC_CLUSTER_IPI),
		ExProcessorMasks:              cpuid.HypervHasRecommendation(cpuid.HYPERV_EXPROCESSORMASKS),
		Nested:                        cpuid.HypervHasRecommendation(cpuid.HYPERV_NESTED),
		INTForMBECSyscalls:            cpuid.HypervHasRecommendation(cpuid.HYPERV_INT_MBEC_SYSCALLS),
		NestedEVMCS:                   cpuid.HypervHasRecommendation(cpuid.HYPERV_NESTED_EVMCS),
		SyncedTimeline:                cpuid.HypervHasRecommendation(cpuid.HYPERV_SYNCED_TIMELINE),
		DirectLocalFlushEntire:        cpuid.HypervHasRecommendation(cpuid.HYPERV_DIRECT_LOCAL_FLUSH_ENTIRE),
		NoNonArchitecturalCoreSharing: cpuid.HypervHasRecommendation(cpuid.HYPERV_NO_NON_ARCHITECTURAL_CORE_SHARING),
		SpinlockRetries:               cpuid.HypervGetSpinlockRetries(),
	}
}

type HypervNestedFeatures struct {
	MSRAccessSyNICRegs           bool
	MSRAccessIntrCtrlRegs        bool
	MSRAccessHypercallRegs       bool
	MSRAccessVPIndex             bool
	MSRAccessReenlighenmentCtrls bool
	HypercallFastXMMRegister     bool
	HypercallFastOutput          bool
	HypercallSintPollingMode     bool
}

func GetHypervNestedFeatures(enabled bool) *HypervNestedFeatures {
	if !enabled {
		return nil
	}
	return &HypervNestedFeatures{
		MSRAccessSyNICRegs:           cpuid.HypervHasNestedFeatureMSR(cpuid.HYPERV_NESTED_FEATURE_MSR_ACCESS_SYNIC_REGS),
		MSRAccessIntrCtrlRegs:        cpuid.HypervHasNestedFeatureMSR(cpuid.HYPERV_NESTED_FEATURE_MSR_ACCESS_INTR_CTRL_REGS),
		MSRAccessHypercallRegs:       cpuid.HypervHasNestedFeatureMSR(cpuid.HYPERV_NESTED_FEATURE_MSR_ACCESS_HYPERCALL_MSRS),
		MSRAccessVPIndex:             cpuid.HypervHasNestedFeatureMSR(cpuid.HYPERV_NESTED_FEATURE_MSR_ACCESS_VPINDEX),
		MSRAccessReenlighenmentCtrls: cpuid.HypervHasNestedFeatureMSR(cpuid.HYPERV_NESTED_FEATURE_MSR_ACCESS_REENLIGHTENMENT_CTRLS),
		HypercallFastXMMRegister:     cpuid.HypervHasNestedFeatureHypercall(cpuid.HYPERV_NESTED_FEATURE_HYPERCALL_FAST_XMM_REGISTER),
		HypercallFastOutput:          cpuid.HypervHasNestedFeatureHypercall(cpuid.HYPERV_NESTED_FEATURE_HYPERCALL_FAST_OUTPUT),
		HypercallSintPollingMode:     cpuid.HypervHasNestedFeatureHypercall(cpuid.HYPERV_NESTED_FEATURE_HYPERCALL_SINT_POLLING_MODE),
	}
}

type HypervNestedOptimizations struct {
	DirectVirtualFlushHypercall bool
	HvFlushGuestPhysAddrList    bool
	EnlightenedMSRBitmap        bool
}

func GetHypervNestedOptimizations(enabled bool) *HypervNestedOptimizations {
	if !enabled {
		return nil
	}
	return &HypervNestedOptimizations{
		DirectVirtualFlushHypercall: cpuid.HypervHasNestedFeatureOptimization(cpuid.HYPERV_NESTED_OPTIMIZATION_DIRECT_VIRTUAL_FLUSH_HYPERCALL),
		HvFlushGuestPhysAddrList:    cpuid.HypervHasNestedFeatureOptimization(cpuid.HYPERV_NESTED_OPTIMIZATION_HV_FLUSH_GUEST_PHYS_ADDR_LIST),
		EnlightenedMSRBitmap:        cpuid.HypervHasNestedFeatureOptimization(cpuid.HYPERV_NESTED_OPTIMIZATION_ENLIGHTENED_MSR_BITMAP),
	}
}

type HypervNestedInfo struct {
	Features      *HypervNestedFeatures      `json:",omitempty"`
	Optimizations *HypervNestedOptimizations `json:",omitempty"`
}

func GetHypervNestedInfo(enabled bool) *HypervNestedInfo {
	if !enabled {
		return nil
	}
	return &HypervNestedInfo{
		Features:      GetHypervNestedFeatures(enabled),
		Optimizations: GetHypervNestedOptimizations(enabled),
	}
}

type HypervInfo struct {
	HyperVsupport   bool
	VendorId        string                     `json:",omitempty"`
	Features        *HypervFeatures            `json:",omitempty"`
	Privileges      *HypervPartitionPrivileges `json:",omitempty"`
	Recommendations *HypervRecommendations     `json:",omitempty"`
	NestedInfo      *HypervNestedInfo          `json:",omitempty"`
}

func hypervEnabled() bool {
	return cpuid.HypervInterfaceSignatureString == cpuid.HypervInterfaceSignature
}

func main() {
	enabled := hypervEnabled()
	info := HypervInfo{
		HyperVsupport:   enabled,
		VendorId:        cpuid.HypervVendorIDSignatureString,
		Features:        GetHypervFeatures(enabled),
		Privileges:      GetHypervPrivileges(enabled),
		Recommendations: GetHypervRecommendations(enabled),
		NestedInfo:      GetHypervNestedInfo(enabled && cpuid.HypervHasAdditionalNestedEnlightenments()),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(&info)
	if err != nil {
		panic(err)
	}
}
