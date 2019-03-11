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
		GuestDebugging:                cpuid.HasHypervFeature(cpuid.HYPERV_GUEST_DEBUGGING),
		PerformanceMonitor:            cpuid.HasHypervFeature(cpuid.HYPERV_PERFORMANCE_MONITOR),
		PCPUDynamicPartitioningEvents: cpuid.HasHypervFeature(cpuid.HYPERV_PCPU_DYNAMIC_PARTITIONING_EVENTS),
		HypercallInputParamsXMM:       cpuid.HasHypervFeature(cpuid.HYPERV_HYPERCALL_INPUT_PARAMETERS_XMM),
		VirtualGuestIdleState:         cpuid.HasHypervFeature(cpuid.HYPERV_VIRTUAL_GUEST_IDLE_STATE),
		HypervisorSleepState:          cpuid.HasHypervFeature(cpuid.HYPERV_HYPERVISOR_SLEEP_STATE),
		NUMADistanceQuery:             cpuid.HasHypervFeature(cpuid.HYPERV_NUMA_DISTANCE_QUERY),
		TimerFrequenciesQuery:         cpuid.HasHypervFeature(cpuid.HYPERV_TIMER_FREQUENCIES_QUERY),
		SytheticMCEInjection:          cpuid.HasHypervFeature(cpuid.HYPERV_SYTHETIC_MCE_INJECTION),
		GuestCrashMSR:                 cpuid.HasHypervFeature(cpuid.HYPERV_GUEST_CRASH_MSR),
		DebugMSR:                      cpuid.HasHypervFeature(cpuid.HYPERV_DEBUG_MSR),
		NPIEP:                         cpuid.HasHypervFeature(cpuid.HYPERV_NPIEP),
		DisableHypervisorAvailable:    cpuid.HasHypervFeature(cpuid.HYPERV_DISABLE_HYPERVISOR),
		ExtendedGvaRangesForFlushVirtualAddressList: cpuid.HasHypervFeature(cpuid.HYPERV_EXTENDED_GVA_RANGES_FOR_FLUSH_VIRTUAL_ADDRESS_LIST),
		HypercallOutputReturnXMM:                    cpuid.HasHypervFeature(cpuid.HYPERV_HYPERCALL_OUTPUT_RETURN_XMM),
		SintPollingMode:                             cpuid.HasHypervFeature(cpuid.HYPERV_SINT_POLLING_MODE),
		HypercallMsrLock:                            cpuid.HasHypervFeature(cpuid.HYPERV_HYPERCALL_MSR_LOCK),
		UseDirectSyntheticTimers:                    cpuid.HasHypervFeature(cpuid.HYPERV_USE_DIRECT_SYNTHETIC_TIMERS),
	}
}

type HypervInfo struct {
	HyperVsupport bool
	Identity      *HypervIdentity `json:",omitempty"`
	Features      *HypervFeatures `json:",omitempty"`
}

func main() {
	enabled := cpuid.HasHypervSupport()
	info := HypervInfo{
		HyperVsupport: enabled,
		Features:      GetHypervFeatures(enabled),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(&info)
	if err != nil {
		panic(err)
	}
}
