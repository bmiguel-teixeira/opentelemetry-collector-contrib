// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogsBuilderAppendLogRecord(t *testing.T) {
	observedZapCore, _ := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopSettings(receivertest.NopType)
	settings.Logger = zap.New(observedZapCore)
	lb := NewLogsBuilder(settings)

	rb := lb.NewResourceBuilder()
	rb.SetContainerID("container.id-val")
	rb.SetContainerImageName("container.image.name-val")
	rb.SetContainerImageTag("container.image.tag-val")
	rb.SetContainerRuntime("container.runtime-val")
	rb.SetContainerRuntimeVersion("container.runtime.version-val")
	rb.SetK8sContainerName("k8s.container.name-val")
	rb.SetK8sContainerStatusLastTerminatedReason("k8s.container.status.last_terminated_reason-val")
	rb.SetK8sCronjobName("k8s.cronjob.name-val")
	rb.SetK8sCronjobUID("k8s.cronjob.uid-val")
	rb.SetK8sDaemonsetName("k8s.daemonset.name-val")
	rb.SetK8sDaemonsetUID("k8s.daemonset.uid-val")
	rb.SetK8sDeploymentName("k8s.deployment.name-val")
	rb.SetK8sDeploymentUID("k8s.deployment.uid-val")
	rb.SetK8sHpaName("k8s.hpa.name-val")
	rb.SetK8sHpaScaletargetrefApiversion("k8s.hpa.scaletargetref.apiversion-val")
	rb.SetK8sHpaScaletargetrefKind("k8s.hpa.scaletargetref.kind-val")
	rb.SetK8sHpaScaletargetrefName("k8s.hpa.scaletargetref.name-val")
	rb.SetK8sHpaUID("k8s.hpa.uid-val")
	rb.SetK8sJobName("k8s.job.name-val")
	rb.SetK8sJobUID("k8s.job.uid-val")
	rb.SetK8sKubeletVersion("k8s.kubelet.version-val")
	rb.SetK8sNamespaceName("k8s.namespace.name-val")
	rb.SetK8sNamespaceUID("k8s.namespace.uid-val")
	rb.SetK8sNodeName("k8s.node.name-val")
	rb.SetK8sNodeUID("k8s.node.uid-val")
	rb.SetK8sPodName("k8s.pod.name-val")
	rb.SetK8sPodQosClass("k8s.pod.qos_class-val")
	rb.SetK8sPodUID("k8s.pod.uid-val")
	rb.SetK8sReplicasetName("k8s.replicaset.name-val")
	rb.SetK8sReplicasetUID("k8s.replicaset.uid-val")
	rb.SetK8sReplicationcontrollerName("k8s.replicationcontroller.name-val")
	rb.SetK8sReplicationcontrollerUID("k8s.replicationcontroller.uid-val")
	rb.SetK8sResourcequotaName("k8s.resourcequota.name-val")
	rb.SetK8sResourcequotaUID("k8s.resourcequota.uid-val")
	rb.SetK8sStatefulsetName("k8s.statefulset.name-val")
	rb.SetK8sStatefulsetUID("k8s.statefulset.uid-val")
	rb.SetOpenshiftClusterquotaName("openshift.clusterquota.name-val")
	rb.SetOpenshiftClusterquotaUID("openshift.clusterquota.uid-val")
	rb.SetOsDescription("os.description-val")
	rb.SetOsType("os.type-val")
	res := rb.Emit()

	// append the first log record
	lr := plog.NewLogRecord()
	lr.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	lr.Attributes().PutStr("type", "log")
	lr.Body().SetStr("the first log record")

	// append the second log record
	lr2 := plog.NewLogRecord()
	lr2.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	lr2.Attributes().PutStr("type", "event")
	lr2.Body().SetStr("the second log record")

	lb.AppendLogRecord(lr)
	lb.AppendLogRecord(lr2)

	logs := lb.Emit(WithLogsResource(res))
	assert.Equal(t, 1, logs.ResourceLogs().Len())

	rl := logs.ResourceLogs().At(0)
	assert.Equal(t, 1, rl.ScopeLogs().Len())

	sl := rl.ScopeLogs().At(0)
	assert.Equal(t, ScopeName, sl.Scope().Name())
	assert.Equal(t, lb.buildInfo.Version, sl.Scope().Version())

	assert.Equal(t, 2, sl.LogRecords().Len())

	attrVal, ok := sl.LogRecords().At(0).Attributes().Get("type")
	assert.True(t, ok)
	assert.Equal(t, "log", attrVal.Str())

	assert.Equal(t, pcommon.ValueTypeStr, sl.LogRecords().At(0).Body().Type())
	assert.Equal(t, "the first log record", sl.LogRecords().At(0).Body().Str())

	attrVal, ok = sl.LogRecords().At(1).Attributes().Get("type")
	assert.True(t, ok)
	assert.Equal(t, "event", attrVal.Str())

	assert.Equal(t, pcommon.ValueTypeStr, sl.LogRecords().At(1).Body().Type())
	assert.Equal(t, "the second log record", sl.LogRecords().At(1).Body().Str())
}
