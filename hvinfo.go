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

type HypervIdentity struct{} // we don't care atm

type HypervFeatures struct {
	GuestDebugging                              bool `json:",omitempty"`
	PerformanceMonitor                          bool `json:",omitempty"`
	PCPUDynamicPartitioningEvents               bool `json:",omitempty"`
	HypercallInputParamsXMM                     bool `json:",omitempty"`
	VirtualGuestIdleState                       bool `json:",omitempty"`
	HypervisorSleepState                        bool `json:",omitempty"`
	NUMADistanceQuery                           bool `json:",omitempty"`
	TimerFrequenciesQuery                       bool `json:",omitempty"`
	SytheticMCEInjection                        bool `json:",omitempty"`
	GuestCrashMSR                               bool `json:",omitempty"`
	DebugMSR                                    bool `json:",omitempty"`
	NPIEP                                       bool `json:",omitempty"`
	DisableHypervisorAvailable                  bool `json:",omitempty"`
	ExtendedGvaRangesForFlushVirtualAddressList bool `json:",omitempty"`
	HypercallOutputReturnXMM                    bool `json:",omitempty"`
	SintPollingMode                             bool `json:",omitempty"`
	HypercallMsrLock                            bool `json:",omitempty"`
	UseDirectSyntheticTimers                    bool `json:",omitempty"`
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

type HypervRecommendations struct {
	HypercallAddressSpaceSwitch   bool   `json:",omitempty"`
	HypercallLocalTLBFlush        bool   `json:",omitempty"`
	HypercallRemoteTLBFlush       bool   `json:",omitempty"`
	MSRAPICRegisters              bool   `json:",omitempty"`
	MSRSystemReset                bool   `json:",omitempty"`
	RelaxedTiming                 bool   `json:",omitempty"`
	DMARemapping                  bool   `json:",omitempty"`
	InterruptRemapping            bool   `json:",omitempty"`
	X2APICMSR                     bool   `json:",omitempty"`
	DeprecatingAutoEOI            bool   `json:",omitempty"`
	SyntheticClusterIPI           bool   `json:",omitempty"`
	ExProcessorMasks              bool   `json:",omitempty"`
	Nested                        bool   `json:",omitempty"`
	INTForMBECSyscalls            bool   `json:",omitempty"`
	NestedEVMCS                   bool   `json:",omitempty"`
	SyncedTimeline                bool   `json:",omitempty"`
	DirectLocalFlushEntire        bool   `json:",omitempty"`
	NoNonArchitecturalCoreSharing bool   `json:",omitempty"`
	SpinlockRetries               uint64 `json:",omitempty"`
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

type HypervInfo struct {
	HyperVsupport   bool
	Identity        *HypervIdentity        `json:",omitempty"`
	Features        *HypervFeatures        `json:",omitempty"`
	Recommendations *HypervRecommendations `json:",omitempty"`
}

func main() {
	enabled := cpuid.HypervSupportEnabled()
	info := HypervInfo{
		HyperVsupport:   enabled,
		Features:        GetHypervFeatures(enabled),
		Recommendations: GetHypervRecommendations(enabled),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(&info)
	if err != nil {
		panic(err)
	}
}
