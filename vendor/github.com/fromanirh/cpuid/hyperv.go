// Copyright 2019 Red Hat Inc
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpuid

const (
	HypervVendorSignature    string = "Microsoft Hv"
	HypervInterfaceSignature string = "Hv#1"
)

func HypervSupportEnabled() bool {
	return HypervVendorIDSignatureString == HypervVendorSignature && HypervInterfaceSignatureString == HypervInterfaceSignature
}

var HypervSupport bool

var HypervCPUIDMax uint32

var HypervVendorIDSignatureString string

var HypervInterfaceSignatureString string

type HypervFeature uint64

var hypervFeatureFlags uint64

func HypervHasFeature(flag HypervFeature) bool {
	return (hypervFeatureFlags & uint64(flag)) != 0
}

var HypervFeatureNames = map[HypervFeature]string{
	HYPERV_GUEST_DEBUGGING:                                    "GUEST_DEBUGGING",
	HYPERV_PERFORMANCE_MONITOR:                                "PERFORMANCE_MONITOR",
	HYPERV_PCPU_DYNAMIC_PARTITIONING_EVENTS:                   "PCPU_DYNAMIC_PARTITIONING_EVENTS",
	HYPERV_HYPERCALL_INPUT_PARAMETERS_XMM:                     "HYPERCALL_INPUT_PARAMETERS_XMM",
	HYPERV_VIRTUAL_GUEST_IDLE_STATE:                           "VIRTUAL_GUEST_IDLE_STATE",
	HYPERV_HYPERVISOR_SLEEP_STATE:                             "HYPERVISOR_SLEEP_STATE",
	HYPERV_NUMA_DISTANCE_QUERY:                                "NUMA_DISTANCE_QUERY",
	HYPERV_TIMER_FREQUENCIES_QUERY:                            "TIMER_FREQUENCIES_QUERY",
	HYPERV_SYTHETIC_MCE_INJECTION:                             "SYTHETIC_MCE_INJECTION",
	HYPERV_GUEST_CRASH_MSR:                                    "GUEST_CRASH_MSR",
	HYPERV_DEBUG_MSR:                                          "DEBUG_MSR",
	HYPERV_NPIEP:                                              "NPIEP",
	HYPERV_DISABLE_HYPERVISOR:                                 "DISABLE_HYPERVISOR",
	HYPERV_EXTENDED_GVA_RANGES_FOR_FLUSH_VIRTUAL_ADDRESS_LIST: "EXTENDED_GVA_RANGES_FOR_FLUSH_VIRTUAL_ADDRESS_LIST",
	HYPERV_HYPERCALL_OUTPUT_RETURN_XMM:                        "HYPERCALL_OUTPUT_RETURN_XMM",
	HYPERV_SINT_POLLING_MODE:                                  "SINT_POLLING_MODE",
	HYPERV_HYPERCALL_MSR_LOCK:                                 "HYPERCALL_MSR_LOCK",
	HYPERV_USE_DIRECT_SYNTHETIC_TIMERS:                        "USE_DIRECT_SYNTHETIC_TIMERS",
}

const (
	_ = HypervFeature(1) << iota
	HYPERV_GUEST_DEBUGGING
	HYPERV_PERFORMANCE_MONITOR
	HYPERV_PCPU_DYNAMIC_PARTITIONING_EVENTS
	HYPERV_HYPERCALL_INPUT_PARAMETERS_XMM
	HYPERV_VIRTUAL_GUEST_IDLE_STATE
	HYPERV_HYPERVISOR_SLEEP_STATE
	HYPERV_NUMA_DISTANCE_QUERY
	HYPERV_TIMER_FREQUENCIES_QUERY
	HYPERV_SYTHETIC_MCE_INJECTION
	HYPERV_GUEST_CRASH_MSR
	HYPERV_DEBUG_MSR
	HYPERV_NPIEP
	HYPERV_DISABLE_HYPERVISOR
	HYPERV_EXTENDED_GVA_RANGES_FOR_FLUSH_VIRTUAL_ADDRESS_LIST
	HYPERV_HYPERCALL_OUTPUT_RETURN_XMM
	_
	HYPERV_SINT_POLLING_MODE
	HYPERV_HYPERCALL_MSR_LOCK
	HYPERV_USE_DIRECT_SYNTHETIC_TIMERS
)

type HypervPartitionPrivilege uint64

func (r HypervPartitionPrivilege) toFlag() uint64 {
	return uint64(r)
}

var hypervPartitionPrivilegeFlags uint64

func HypervHasPrivilege (privilege HypervPartitionPrivilege) bool {
	return (hypervPartitionPrivilegeFlags & privilege.toFlag()) != 0
}

const (
	// Access to virtual MSRs
	HYPERV_AccessVpRunTimeRegs = HypervPartitionPrivilege(1) << iota
	HYPERV_AccessPartitionReferenceCounter
	HYPERV_AccessSynicRegs
	HYPERV_AccessSyntheticTimerRegs
	HYPERV_AccessIntrCtrlRegs
	HYPERV_AccessHypercallMsrs
	HYPERV_AccessVpIndex
	HYPERV_AccessResetReg
	HYPERV_AccessStatsReg                 // 8
	HYPERV_AccessPartitionReferenceTsc
	HYPERV_AccessGuestIdleReg
	HYPERV_AccessFrequencyRegs
	HYPERV_AccessDebugRegs
	HYPERV_AccessReenlightenmentControls
	_
	_
	_                                     // 16
	_
	_
	_
	_
	_
	_
	_
	_                                     // 24
	_
	_
	_
	_
	_
	_
	_

	// Access to hypercalls
	HYPERV_CreatePartitions               // 32
	HYPERV_AccessPartitionId
	HYPERV_AccessMemoryPool
	_
	HYPERV_PostMessages
	HYPERV_SignalEvents
	HYPERV_CreatePort
	HYPERV_ConnectPort
	HYPERV_AccessStats                    // 40
	_
	_
	HYPERV_Debugging
	HYPERV_CpuManagement
	_
	_
	_
	HYPERV_AccessVSM                      // 48
	HYPERV_AccessVpRegisters
	_
	_
	HYPERV_EnableExtendedHypercalls
	HYPERV_StartVirtualProcessor
)

type HypervRecommendation uint64

func (r HypervRecommendation) toFlag() uint64 {
	return uint64(r)
}

var hypervRecommendationFlags uint64
var hypervSpinlockRetries uint64

func HypervGetSpinlockRetries() uint64 {
	return hypervSpinlockRetries
}

func HypervHasAdditionalNestedEnlightenments() bool {
	// The flag is overloaded. See spec sec. 2.4.5 for details
	return (hypervRecommendationFlags & HYPERV_NESTED_EVMCS.toFlag()) != 0
}

func HypervHasRecommendation(recommendation HypervRecommendation) bool {
	return (hypervRecommendationFlags & recommendation.toFlag()) != 0
}

var HypervRecommendationNames = map[HypervRecommendation]string{
	HYPERV_HYPERCALL_ADDRESS_SPACE_SWITCH:    "HYPERV_HYPERCALL_ADDRESS_SPACE_SWITCH",
	HYPERV_HYPERCALL_LOCAL_TLB_FLUSH:         "HYPERV_HYPERCALL_LOCAL_TLB_FLUSH",
	HYPERV_HYPERCALL_REMOTE_TLB_FLUSH:        "HYPERV_HYPERCALL_REMOTE_TLB_FLUSH",
	HYPERV_MSR_APIC_REGISTERS:                "HYPERV_MSR_APIC_REGISTERS",
	HYPERV_MSR_SYSTEM_RESET:                  "HYPERV_MSR_SYSTEM_RESET",
	HYPERV_RELAXED_TIMING:                    "HYPERV_RELAXED_TIMING",
	HYPERV_DMA_REMAPPING:                     "HYPERV_DMA_REMAPPING",
	HYPERV_INTERRUPT_REMAPPING:               "HYPERV_INTERRUPT_REMAPPING",
	HYPERV_x2APIC_MSR:                        "HYPERV_x2APIC_MSR",
	HYPERV_DEPRECATING_AUTOEOI:               "HYPERV_DEPRECATING_AUTOEOI",
	HYPERV_SYNTHETIC_CLUSTER_IPI:             "HYPERV_SYNTHETIC_CLUSTER_IPI",
	HYPERV_EXPROCESSORMASKS:                  "HYPERV_EXPROCESSORMASKS",
	HYPERV_NESTED:                            "HYPERV_NESTED",
	HYPERV_INT_MBEC_SYSCALLS:                 "HYPERV_INT_MBEC_SYSCALLS",
	HYPERV_NESTED_EVMCS:                      "HYPERV_NESTED_EVMCS",
	HYPERV_SYNCED_TIMELINE:                   "HYPERV_SYNCED_TIMELINE",
	HYPERV_DIRECT_LOCAL_FLUSH_ENTIRE:         "HYPERV_DIRECT_LOCAL_FLUSH_ENTIRE",
	HYPERV_NO_NON_ARCHITECTURAL_CORE_SHARING: "HYPERV_NO_NON_ARCHITECTURAL_CORE_SHARING",
}

const (
	HYPERV_HYPERCALL_ADDRESS_SPACE_SWITCH = HypervRecommendation(1) << iota
	HYPERV_HYPERCALL_LOCAL_TLB_FLUSH
	HYPERV_HYPERCALL_REMOTE_TLB_FLUSH
	HYPERV_MSR_APIC_REGISTERS
	HYPERV_MSR_SYSTEM_RESET
	HYPERV_RELAXED_TIMING
	HYPERV_DMA_REMAPPING
	HYPERV_INTERRUPT_REMAPPING
	HYPERV_x2APIC_MSR
	HYPERV_DEPRECATING_AUTOEOI
	HYPERV_SYNTHETIC_CLUSTER_IPI
	HYPERV_EXPROCESSORMASKS
	HYPERV_NESTED
	HYPERV_INT_MBEC_SYSCALLS
	HYPERV_NESTED_EVMCS
	HYPERV_SYNCED_TIMELINE
	_
	HYPERV_DIRECT_LOCAL_FLUSH_ENTIRE
	HYPERV_NO_NON_ARCHITECTURAL_CORE_SHARING
)

type HypervNestedFeatureMSR uint64

func (r HypervNestedFeatureMSR) toFlag() uint64 {
	return uint64(r)
}

var hypervNestedFeatureMSRFlags uint64

func HypervHasNestedFeatureMSR(nestedFeatureMSR HypervNestedFeatureMSR) bool {
	return (hypervNestedFeatureMSRFlags & nestedFeatureMSR.toFlag()) != 0
}

var HypervNestedFeatureMSRNames = map[HypervNestedFeatureMSR]string{
	HYPERV_NESTED_FEATURE_MSR_ACCESS_SYNIC_REGS:            "HYPERV_NESTED_FEATURE_MSR_ACCESS_SYNIC_REGS",
	HYPERV_NESTED_FEATURE_MSR_ACCESS_INTR_CTRL_REGS:        "HYPERV_NESTED_FEATURE_MSR_ACCESS_INTR_CTRL_REGS",
	HYPERV_NESTED_FEATURE_MSR_ACCESS_HYPERCALL_MSRS:        "HYPERV_NESTED_FEATURE_MSR_ACCESS_HYPERCALL_MSRS",
	HYPERV_NESTED_FEATURE_MSR_ACCESS_VPINDEX:               "HYPERV_NESTED_FEATURE_MSR_ACCESS_VPINDEX",
	HYPERV_NESTED_FEATURE_MSR_ACCESS_REENLIGHTENMENT_CTRLS: "HYPERV_NESTED_FEATURE_MSR_ACCESS_REENLIGHTENMENT_CTRLS",
}

const (
	_ = HypervNestedFeatureMSR(1) << iota
	_
	HYPERV_NESTED_FEATURE_MSR_ACCESS_SYNIC_REGS
	_
	HYPERV_NESTED_FEATURE_MSR_ACCESS_INTR_CTRL_REGS
	HYPERV_NESTED_FEATURE_MSR_ACCESS_HYPERCALL_MSRS
	HYPERV_NESTED_FEATURE_MSR_ACCESS_VPINDEX
	_
	_
	_
	_
	_
	HYPERV_NESTED_FEATURE_MSR_ACCESS_REENLIGHTENMENT_CTRLS
)

type HypervNestedFeatureHypercall uint64

func (r HypervNestedFeatureHypercall) toFlag() uint64 {
	return uint64(r)
}

var hypervNestedFeatureHypercallFlags uint64

func HypervHasNestedFeatureHypercall(nestedFeatureHypercall HypervNestedFeatureHypercall) bool {
	return (hypervNestedFeatureHypercallFlags & nestedFeatureHypercall.toFlag()) != 0
}

var HypervNestedFeatureHypercallNames = map[HypervNestedFeatureHypercall]string{
	HYPERV_NESTED_FEATURE_HYPERCALL_FAST_XMM_REGISTER: "HYPERV_NESTED_FEATURE_HYPERCALL_FAST_XMM_REGISTER",
	HYPERV_NESTED_FEATURE_HYPERCALL_FAST_OUTPUT:       "HYPERV_NESTED_FEATURE_HYPERCALL_FAST_OUTPUT",
	HYPERV_NESTED_FEATURE_HYPERCALL_SINT_POLLING_MODE: "HYPERV_NESTED_FEATURE_HYPERCALL_SINT_POLLING_MODE",
}

const (
	_ = HypervNestedFeatureHypercall(1) << iota
	_
	_
	_
	HYPERV_NESTED_FEATURE_HYPERCALL_FAST_XMM_REGISTER
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	HYPERV_NESTED_FEATURE_HYPERCALL_FAST_OUTPUT
	_
	HYPERV_NESTED_FEATURE_HYPERCALL_SINT_POLLING_MODE
)

type HypervNestedFeatureOptimization uint64

func (r HypervNestedFeatureOptimization) toFlag() uint64 {
	return uint64(r)
}

var hypervNestedFeatureOptimizationFlags uint64

func HypervHasNestedFeatureOptimization(nestedFeatureOptimization HypervNestedFeatureOptimization) bool {
	return (hypervNestedFeatureOptimizationFlags & nestedFeatureOptimization.toFlag()) != 0
}

var HypervNestedFeatureOptimizationNames = map[HypervNestedFeatureOptimization]string{
	HYPERV_NESTED_OPTIMIZATION_DIRECT_VIRTUAL_FLUSH_HYPERCALL: "HYPERV_NESTED_OPTIMIZATION_DIRECT_VIRTUAL_FLUSH_HYPERCALL",
	HYPERV_NESTED_OPTIMIZATION_HV_FLUSH_GUEST_PHYS_ADDR_LIST:  "HYPERV_NESTED_OPTIMIZATION_HV_FLUSH_GUEST_PHYS_ADDR_LIST",
	HYPERV_NESTED_OPTIMIZATION_ENLIGHTENED_MSR_BITMAP:         "HYPERV_NESTED_OPTIMIZATION_ENLIGHTENED_MSR_BITMAP",
}

const (
	_ = HypervNestedFeatureOptimization(1) << iota
	_
	_
	_
	_
	_
	_
	_ // 7
	_
	_
	_
	_
	_
	_
	_
	_ // 15
	_
	HYPERV_NESTED_OPTIMIZATION_DIRECT_VIRTUAL_FLUSH_HYPERCALL
	HYPERV_NESTED_OPTIMIZATION_HV_FLUSH_GUEST_PHYS_ADDR_LIST
	HYPERV_NESTED_OPTIMIZATION_ENLIGHTENED_MSR_BITMAP
)
