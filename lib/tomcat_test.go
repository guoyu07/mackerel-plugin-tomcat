package mptomcat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMetrics(t *testing.T) {

	xml := `<?xml version="1.0" encoding="utf-8"?><?xml-stylesheet type="text/xsl" href="/manager/xform.xsl" ?>
<status><jvm><memory free='120588912' total='167247872' max='3817865216'/><memorypool name='PS Eden Space' type='Heap memory' usageInit='67108864' usageCommitted='67108864' usageMax='1409286144' usageUsed='30620784'/><memorypool name='PS Old Gen' type='Heap memory' usageInit='179306496' usageCommitted='89128960' usageMax='2863661056' usageUsed='16038176'/><memorypool name='PS Survivor Space' type='Heap memory' usageInit='11010048' usageCommitted='11010048' usageMax='11010048' usageUsed='0'/><memorypool name='Code Cache' type='Non-heap memory' usageInit='2555904' usageCommitted='8978432' usageMax='251658240' usageUsed='8866688'/><memorypool name='Compressed Class Space' type='Non-heap memory' usageInit='0' usageCommitted='2883584' usageMax='1073741824' usageUsed='2706696'/><memorypool name='Metaspace' type='Non-heap memory' usageInit='0' usageCommitted='26738688' usageMax='-1' usageUsed='25961400'/></jvm><connector name='"ajp-nio-8009"'><threadInfo  maxThreads="200" currentThreadCount="10" currentThreadsBusy="0" /><requestInfo  maxTime="0" processingTime="0" requestCount="0" errorCount="0" bytesReceived="0" bytesSent="0" /><workers></workers></connector><connector name='"http-nio-8080"'><threadInfo  maxThreads="200" currentThreadCount="10" currentThreadsBusy="1" /><requestInfo  maxTime="642" processingTime="1041" requestCount="110" errorCount="6" bytesReceived="0" bytesSent="1096491" /><workers><worker  stage="R" requestProcessingTime="0" requestBytesSent="0" requestBytesReceived="0" remoteAddr="&#63;" virtualHost="&#63;" method="&#63;" currentUri="&#63;" currentQueryString="&#63;" protocol="&#63;" /><worker  stage="R" requestProcessingTime="0" requestBytesSent="0" requestBytesReceived="0" remoteAddr="&#63;" virtualHost="&#63;" method="&#63;" currentUri="&#63;" currentQueryString="&#63;" protocol="&#63;" /><worker  stage="S" requestProcessingTime="2" requestBytesSent="0" requestBytesReceived="0" remoteAddr="0:0:0:0:0:0:0:1" virtualHost="localhost" method="GET" currentUri="/manager/status/all" currentQueryString="XML=true" protocol="HTTP/1.1" /><worker  stage="R" requestProcessingTime="0" requestBytesSent="0" requestBytesReceived="0" remoteAddr="&#63;" virtualHost="&#63;" method="&#63;" currentUri="&#63;" currentQueryString="&#63;" protocol="&#63;" /><worker  stage="R" requestProcessingTime="0" requestBytesSent="0" requestBytesReceived="0" remoteAddr="&#63;" virtualHost="&#63;" method="&#63;" currentUri="&#63;" currentQueryString="&#63;" protocol="&#63;" /></workers></connector></status>`

	var p TomcatPlugin
	metrics := make(map[string]float64)

	err := p.parseMetrics(metrics, []byte(xml))
	if err != nil {
		t.Fatal(err)
	}

	if len(metrics) == 0 {
		t.Fatalf("metrics is empty")
	}

	assert.Equal(t, metrics["free"], float64(120588912))
	assert.Equal(t, metrics["total"], float64(167247872))
	assert.Equal(t, metrics["used"], float64(167247872 - 120588912))

	assert.Equal(t, metrics["thread.ajp.currentThreadsBusy"], float64(0))
	assert.Equal(t, metrics["thread.http.currentThreadsBusy"], float64(1))
}