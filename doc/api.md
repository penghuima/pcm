# 接口文档

# 目录 

[toc]

# API接口规范 

> HTTP接口遵循API协议规范。返回数据格式统一如下：


```
{
    "code":"", //必选,返回码
    "message":"" ,//可选，返回消息， 
    "other":""//可选
    ...
}
```

> Api returnCode定义(一般采用http状态码)


| code   | value      |
| ------ | ---------- |
| 200    | 正常       |
| 500    | 服务器错误 |
| 自定义 | 自定义     |

 http状态码参考 https://httpstatuses.com/

一、节点管理

## 1、集群节点查询接口

### 应用场景

​      点击节点管理，即可查看当前集群下节点的运行状态。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |             查询当前集群下所有节点以及其运行状态             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /prod-api/kapis/clusters/{cluster}/resources.kubesphere.io/v1alpha3/nodes |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.45.100.73/prod-api/kapis/clusters/host/resources.kubesphere.io/v1alpha3/nodes?sortBy=createTime&page=1&limit=10 |

请求参数

|    参数     |       描述       | 必填 |  类型  |
| :---------: | :--------------: | :--: | :----: |
| **cluster** |     集群角色     |  是  | String |
| **sortBy**  |     排序方式     |  是  | String |
|  **page**   |       页数       |  是  |  int   |
|  **limit**  | 一页最大显示数量 |  是  |  int   |

响应参数 

|          参数           |   描述   |
| :---------------------: | :------: |
|       totalItems        | 节点数量 |
|          name           | 节点名称 |
|        public-ip        |  节点ip  |
| node-role.kubernetes.io |   角色   |
|           CPU           | CPU核数  |
|         memory          | 内存大小 |
|          pods           | pod数量  |
|         status          | 运行状态 |

请求示例

```
{
   sortBy: createTime
   page: 1
   limit: 10
}
```

响应示例：

​       成功情况

```
{ 
  items: [{metadata: {name: "jcc-txy-003", selfLink: "/api/v1/nodes/jcc-txy-003",…},…},…],…}
  items: [{metadata: {name: "jcc-txy-003", selfLink: "/api/v1/nodes/jcc-txy-003",…},…},…]
  0: {metadata: {name: "jcc-txy-003", selfLink: "/api/v1/nodes/jcc-txy-003",…},…}
  metadata: {name: "jcc-txy-003", selfLink: "/api/v1/nodes/jcc-txy-003",…}
  annotations: {flannel.alpha.coreos.com/backend-data: "{"VNI":1,"VtepMAC":"de:fd:98:f3:29:f1"}",…}
  flannel.alpha.coreos.com/backend-data: "{\"VNI\":1,\"VtepMAC\":\"de:fd:98:f3:29:f1\"}"
  flannel.alpha.coreos.com/backend-type: "vxlan"
  flannel.alpha.coreos.com/kube-subnet-manager: "true"
  flannel.alpha.coreos.com/public-ip: "10.206.0.15"
  kubeadm.alpha.kubernetes.io/cri-socket: "/var/run/dockershim.sock"
  node.alpha.kubernetes.io/ttl: "0"
  node.kubesphere.io/cpu-limits: "9200m"
  node.kubesphere.io/cpu-limits-fraction: "459%"
  node.kubesphere.io/cpu-requests: "762m"
  node.kubesphere.io/cpu-requests-fraction: "38%"
  node.kubesphere.io/last-annotated-at: "2021-06-24T10:04:13+08:00"
  node.kubesphere.io/memory-limits: "20325Mi"
  node.kubesphere.io/memory-limits-fraction: "552%"
  node.kubesphere.io/memory-requests: "1167Mi"
  node.kubesphere.io/memory-requests-fraction: "31%"
  volumes.kubernetes.io/controller-managed-attach-detach: "true"
  creationTimestamp: "2021-06-06T11:45:29Z"
  labels: {beta.kubernetes.io/arch: "amd64", beta.kubernetes.io/os: "linux", kubernetes.io/arch: "amd64",…}
  beta.kubernetes.io/arch: "amd64"
  beta.kubernetes.io/os: "linux"
  kubernetes.io/arch: "amd64"
  kubernetes.io/hostname: "jcc-txy-003"
  kubernetes.io/os: "linux"
  managedFields: [{manager: "kube-controller-manager", operation: "Update", apiVersion: "v1",…},…]
  name: "jcc-txy-003"
  resourceVersion: "5855166"
  selfLink: "/api/v1/nodes/jcc-txy-003"
  uid: "180fbbaf-57fd-4dd7-9ace-573bfb879f11"
  spec: {podCIDR: "10.244.2.0/24", podCIDRs: ["10.244.2.0/24"]}
  status: {capacity: {cpu: "2", ephemeral-storage: "51539404Ki", hugepages-1Gi: "0", hugepages-2Mi: "0",…},…}
  1: {metadata: {name: "jcc-txy-002", selfLink: "/api/v1/nodes/jcc-txy-002",…},…}
  2: {metadata: {name: "jcc-txy-001", selfLink: "/api/v1/nodes/jcc-txy-001",…},…}
  totalItems: 3
}
```

## 2、根据节点名查询集群节点详细信息接口

### 应用场景

​      在节点管理列表，选择点击其中一个集群节点，即可查看对应节点详细内容。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                   查询该集群节点的详细信息                   |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /prod-api/kapis/clusters/{cluster}/resources.kubesphere.io/v1alpha3/nodes/{node} |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.45.100.73/prod-api/kapis/clusters/bj-member2/resources.kubesphere.io/v1alpha3/nodes/jcc-bj006 |

请求参数

```
无
```

响应参数 

|                参数                |      描述       |
| :--------------------------------: | :-------------: |
|               status               |    运行状态     |
| flannel.alpha.coreos.com/public-ip |     IP地址      |
|              osImage               |    操作系统     |
|          operatingSystem           |  操作系统类型   |
|      node-role.kubernetes.io       |      角色       |
|           kernelVersion            |    内核版本     |
|      containerRuntimeVersion       |    容器版本     |
|           kubeletVersion           |  Kubelet 版本   |
|          kubeProxyVersion          | Kube-Proxy 版本 |
|            architecture            |    系统架构     |
|         creationTimestamp          |    创建时间     |

请求示例

```
http://119.45.100.73/prod-api/kapis/clusters/host/resources.kubesphere.io/v1alpha3/nodes/jcc-txy-001
```

响应示例：

​       成功情况

```
{metadata: {name: "jcc-txy-001", selfLink: "/api/v1/nodes/jcc-txy-001",…},…}
metadata: {name: "jcc-txy-001", selfLink: "/api/v1/nodes/jcc-txy-001",…}
annotations: {flannel.alpha.coreos.com/backend-data: "{"VNI":1,"VtepMAC":"3a:41:25:2b:f8:58"}",…}
flannel.alpha.coreos.com/backend-data: "{\"VNI\":1,\"VtepMAC\":\"3a:41:25:2b:f8:58\"}"
flannel.alpha.coreos.com/backend-type: "vxlan"
flannel.alpha.coreos.com/kube-subnet-manager: "true"
flannel.alpha.coreos.com/public-ip: "10.206.0.12"
kubeadm.alpha.kubernetes.io/cri-socket: "/var/run/dockershim.sock"
node.alpha.kubernetes.io/ttl: "0"
node.kubesphere.io/cpu-limits: "4400m"
node.kubesphere.io/cpu-limits-fraction: "220%"
node.kubesphere.io/cpu-requests: "1132m"
node.kubesphere.io/cpu-requests-fraction: "56%"
node.kubesphere.io/last-annotated-at: "2021-06-24T14:45:34+08:00"
node.kubesphere.io/memory-limits: "3426Mi"
node.kubesphere.io/memory-limits-fraction: "94%"
node.kubesphere.io/memory-requests: "740Mi"
node.kubesphere.io/memory-requests-fraction: "20%"
volumes.kubernetes.io/controller-managed-attach-detach: "true"
creationTimestamp: "2021-06-06T11:44:41Z"
labels: {beta.kubernetes.io/arch: "amd64", beta.kubernetes.io/os: "linux", kubernetes.io/arch: "amd64",…}
managedFields: [{manager: "kubelet", operation: "Update", apiVersion: "v1", time: "2021-06-06T11:44:41Z",…},…]
name: "jcc-txy-001"
resourceVersion: "5919629"
selfLink: "/api/v1/nodes/jcc-txy-001"
uid: "dc6f1318-a8a4-40e2-8e84-12435046ff5f"
spec: {podCIDR: "10.244.0.0/24", podCIDRs: ["10.244.0.0/24"],…}
status: {capacity: {cpu: "2", ephemeral-storage: "51539404Ki", hugepages-1Gi: "0", hugepages-2Mi: "0",…},…}
addresses: [{type: "InternalIP", address: "10.206.0.12"}, {type: "Hostname", address: "jcc-txy-001"}]
allocatable: {cpu: "2", ephemeral-storage: "47498714648", hugepages-1Gi: "0", hugepages-2Mi: "0",…}
capacity: {cpu: "2", ephemeral-storage: "51539404Ki", hugepages-1Gi: "0", hugepages-2Mi: "0",…}
conditions: [{type: "NetworkUnavailable", status: "False", lastHeartbeatTime: "2021-06-06T11:45:39Z",…},…]
daemonEndpoints: {kubeletEndpoint: {Port: 10250}}
images: [{names: [,…], sizeBytes: 703055917}, {names: ["k8s.gcr.io/etcd:3.4.13-0"], sizeBytes: 253392289},…]
nodeInfo: {machineID: "051dd68ea77342889b036bf27e6b4292", systemUUID: "051dd68e-a773-4288-9b03-6bf27e6b4292",…}
architecture: "amd64"
bootID: "18c82e86-f941-4211-b050-fd0629d56a71"
containerRuntimeVersion: "docker://20.10.7"
kernelVersion: "4.18.0-305.3.1.el8.x86_64"
kubeProxyVersion: "v1.20.7"
kubeletVersion: "v1.20.7"
machineID: "051dd68ea77342889b036bf27e6b4292"
operatingSystem: "linux"
osImage: "CentOS Linux 8"
systemUUID: "051dd68e-a773-4288-9b03-6bf27e6b4292"
}
```


## 3、根据pod名查询pod信息详情接口

### 应用场景

​      场景：在容器组列表，点击单个的pod，即可查看该pod详细信息。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                       查看单个pod详情                        |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |              /api/v1/namespaces/dev/pods/{pods}              |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/dev/pods/cftgh-79f5598df9-js6ld |

请求参数

| 参数 |   描述   | 必填 |  类型  |
| :--: | :------: | :--: | :----: |
| pods | pods名称 |  是  | String |

响应参数 

|     参数     |   描述   |
| :----------: | :------: |
|     name     | pod名称  |
|  namespace   |  项目名  |
|   nodeName   | 节点名称 |
|    podIP     | 容器组IP |
|    hostIP    |  主机ip  |
|    state     | 运行状态 |
| restartCount | 重启次数 |
|  startTime   | 创建时间 |

请求示例

```
 http://119.3.157.144:30880/api/v1/namespaces/dev/pods/cftgh-79f5598df9-js6ld
```

响应示例：

​       成功情况

```
{
  "kind": "Pod",
  "apiVersion": "v1",
  "metadata": {
    "name": "ks-installer-69699b8464-vkm57",
    "generateName": "ks-installer-69699b8464-",
    "namespace": "kubesphere-system",
    "selfLink": "/api/v1/namespaces/kubesphere-system/pods/ks-installer-69699b8464-vkm57",
    "uid": "80f12035-23a4-47ea-84c0-22eefd67f154",
    "resourceVersion": "22660320",
    "creationTimestamp": "2021-06-23T09:26:35Z" 
  },
  "spec": {
    "restartPolicy": "Always",
    "terminationGracePeriodSeconds": 30,
    "dnsPolicy": "ClusterFirst",
    "serviceAccountName": "ks-installer",
    "serviceAccount": "ks-installer",
    "nodeName": "kubex-0003",
    "schedulerName": "default-scheduler",
    "priority": 0,
    "enableServiceLinks": true
  },
  "status": {
    "phase": "Running",
    "hostIP": "192.168.0.227",
    "podIP": "10.244.2.23",
    "startTime": "2021-06-23T09:26:35Z",
    "containerStatuses": [
      {
        "name": "installer",
        "state": {
          "running": {
            "startedAt": "2021-06-23T09:26:52Z"
          }
        },
        "ready": true,
        "restartCount": 0,
        "image": "kubesphere/ks-installer:v3.0.0",
        "imageID": "docker-pullable://kubesphere/ks-installer@sha256:2b6f2efdb865baac08a0cc3e88696c00ac195b5ccdf1f58947819b69cc405aaf",
        "containerID": "docker://d1167e260f8f7faa4783561e95f1c5e7af829aa4b2a59cd722f28655138a9356",
        "started": true
      }
    ],
    "qosClass": "BestEffort"
  }
}
```



## 4、根据pod名查询pod配置文件

### 应用场景

​      场景：点击查看单个的pod配置，即可查看该pod详细配置信息。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                     查看pod详细配置信息                      |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |     /api/v1/namespaces/kubesphere-system/pods/{podsname}     |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/kubesphere-system/pods/ks-installer-69699b8464-vkm57 |

请求参数

| 参数 |   描述   | 必填 |  类型  |
| :--: | :------: | :--: | :----: |
| pods | pods名称 |  是  | String |

响应参数 

|     参数     |   描述   |
| :----------: | :------: |
|     name     | pod名称  |
|  namespace   |  项目名  |
|   nodeName   | 节点名称 |
|    podIP     | 容器组IP |
|    hostIP    |  主机ip  |
|    state     | 运行状态 |
| restartCount | 重启次数 |
|  startTime   | 创建时间 |

请求示例

```
 http://119.3.157.144:30880/api/v1/namespaces/kubesphere-system/pods/ks-installer-69699b8464-vkm57
```

响应示例：

​       成功情况

```
{
  "kind": "Pod",
  "apiVersion": "v1",
  "metadata": {
    "name": "ks-installer-69699b8464-vkm57",
    "generateName": "ks-installer-69699b8464-",
    "namespace": "kubesphere-system",
    "selfLink": "/api/v1/namespaces/kubesphere-system/pods/ks-installer-69699b8464-vkm57",
    "uid": "80f12035-23a4-47ea-84c0-22eefd67f154",
    "resourceVersion": "22660320",
    "creationTimestamp": "2021-06-23T09:26:35Z",
    "labels": {
      "app": "ks-install",
      "pod-template-hash": "69699b8464"
    },
    "annotations": {
      "kubectl.kubernetes.io/restartedAt": "2021-06-23T17:26:35+08:00"
    },
    "ownerReferences": [
      {
        "apiVersion": "apps/v1",
        "kind": "ReplicaSet",
        "name": "ks-installer-69699b8464",
        "uid": "87b0a449-8219-444d-a6f2-e512816d1a5f",
        "controller": true,
        "blockOwnerDeletion": true
      }
    ],
    "managedFields": [
      {
        "manager": "kube-controller-manager",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-06-23T09:26:35Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/restartedAt":{}},"f:generateName":{},"f:labels":{".":{},"f:app":{},"f:pod-template-hash":{}},"f:ownerReferences":{".":{},"k:{\"uid\":\"87b0a449-8219-444d-a6f2-e512816d1a5f\"}":{".":{},"f:apiVersion":{},"f:blockOwnerDeletion":{},"f:controller":{},"f:kind":{},"f:name":{},"f:uid":{}}}},"f:spec":{"f:containers":{"k:{\"name\":\"installer\"}":{".":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{},"f:volumeMounts":{".":{},"k:{\"mountPath\":\"/etc/localtime\"}":{".":{},"f:mountPath":{},"f:name":{}}}}},"f:dnsPolicy":{},"f:enableServiceLinks":{},"f:restartPolicy":{},"f:schedulerName":{},"f:securityContext":{},"f:serviceAccount":{},"f:serviceAccountName":{},"f:terminationGracePeriodSeconds":{},"f:volumes":{".":{},"k:{\"name\":\"host-time\"}":{".":{},"f:hostPath":{".":{},"f:path":{},"f:type":{}},"f:name":{}}}}}
      },
      {
        "manager": "kubelet",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-06-23T09:26:52Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:status":{"f:conditions":{"k:{\"type\":\"ContainersReady\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}},"k:{\"type\":\"Initialized\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}},"k:{\"type\":\"Ready\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}}},"f:containerStatuses":{},"f:hostIP":{},"f:phase":{},"f:podIP":{},"f:podIPs":{".":{},"k:{\"ip\":\"10.244.2.23\"}":{".":{},"f:ip":{}}},"f:startTime":{}}}
      }
    ]
  },
  "spec": {
    "volumes": [
      {
        "name": "host-time",
        "hostPath": {
          "path": "/etc/localtime",
          "type": ""
        }
      },
      {
        "name": "ks-installer-token-7z24k",
        "secret": {
          "secretName": "ks-installer-token-7z24k",
          "defaultMode": 420
        }
      }
    ],
    "containers": [
      {
        "name": "installer",
        "image": "kubesphere/ks-installer:v3.0.0",
        "resources": {
          
        },
        "volumeMounts": [
          {
            "name": "host-time",
            "mountPath": "/etc/localtime"
          },
          {
            "name": "ks-installer-token-7z24k",
            "readOnly": true,
            "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
          }
        ],
        "terminationMessagePath": "/dev/termination-log",
        "terminationMessagePolicy": "File",
        "imagePullPolicy": "Always"
      }
    ],
    "restartPolicy": "Always",
    "terminationGracePeriodSeconds": 30,
    "dnsPolicy": "ClusterFirst",
    "serviceAccountName": "ks-installer",
    "serviceAccount": "ks-installer",
    "nodeName": "kubex-0003",
    "securityContext": {
      
    },
    "schedulerName": "default-scheduler",
    "tolerations": [
      {
        "key": "node.kubernetes.io/not-ready",
        "operator": "Exists",
        "effect": "NoExecute",
        "tolerationSeconds": 300
      },
      {
        "key": "node.kubernetes.io/unreachable",
        "operator": "Exists",
        "effect": "NoExecute",
        "tolerationSeconds": 300
      }
    ],
    "priority": 0,
    "enableServiceLinks": true
  },
  "status": {
    "phase": "Running",
    "conditions": [
      {
        "type": "Initialized",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2021-06-23T09:26:35Z"
      },
      {
        "type": "Ready",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2021-06-23T09:26:52Z"
      },
      {
        "type": "ContainersReady",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2021-06-23T09:26:52Z"
      },
      {
        "type": "PodScheduled",
        "status": "True",
        "lastProbeTime": null,
        "lastTransitionTime": "2021-06-23T09:26:35Z"
      }
    ],
    "hostIP": "192.168.0.227",
    "podIP": "10.244.2.23",
    "podIPs": [
      {
        "ip": "10.244.2.23"
      }
    ],
    "startTime": "2021-06-23T09:26:35Z",
    "containerStatuses": [
      {
        "name": "installer",
        "state": {
          "running": {
            "startedAt": "2021-06-23T09:26:52Z"
          }
        },
        "lastState": {
          
        },
        "ready": true,
        "restartCount": 0,
        "image": "kubesphere/ks-installer:v3.0.0",
        "imageID": "docker-pullable://kubesphere/ks-installer@sha256:2b6f2efdb865baac08a0cc3e88696c00ac195b5ccdf1f58947819b69cc405aaf",
        "containerID": "docker://d1167e260f8f7faa4783561e95f1c5e7af829aa4b2a59cd722f28655138a9356",
        "started": true
      }
    ],
    "qosClass": "BestEffort"
  }
}
```

## 5、删除单个pod（容器组）接口

### 应用场景

​      场景：点击查看单个的pod，点击删除按钮，即可删除pod

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         删除单个pod                          |
|        请求协议        |                             HTTP                             |
|        请求方法        |                            DELETE                            |
|        请求格式        |                             JSON                             |
|        请求url         |       /api/v1/namespaces/{namespaces}/pods/{podsname}        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/xiaoyan/pods/xxxxx-79f44c55f9-wmhjh |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
|    pods    | pods名称 |  是  | String |
| namespaces | 项目名称 |  是  | String |

响应参数 

|     参数     |   描述   |
| :----------: | :------: |
|     name     | pod名称  |
|  namespace   |  项目名  |
|   nodeName   | 节点名称 |
|    podIP     | 容器组IP |
|    hostIP    |  主机ip  |
|    state     | 运行状态 |
| restartCount | 重启次数 |
|  startTime   | 创建时间 |

请求示例

```
http://119.3.157.144:30880/api/v1/namespaces/xiaoyan/pods/xxxxx-79f44c55f9-wmhjh
```

响应示例：

​       成功情况

```
{

}
```

## 6、查询集群级别的指标数据接口

###  应用场景

> 场景1：xxxxxxxx

### 接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 获取集群级别的指标数据                                       |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/monitoring.kubesphere.io/v1alpha3/cluster             |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/cluster?metrics_filter=cluster_cpu_usage%7Ccluster_cpu_total%7Ccluster_memory_usage_wo_cache%7Ccluster_memory_total%7Ccluster_disk_size_usage%7Ccluster_disk_size_capacity%7Ccluster_pod_running_count%7Ccluster_pod_quota%24 |

请求参数

|      参数      |                             描述                             | 必填 |  类型  |
| :------------: | :----------------------------------------------------------: | :--: | :----: |
| metrics_filter | 指标名称过滤器由一个正则表达式模式组成。 它指定要返回的指标数据 |  否  | String |
|     start      | 查询的开始时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如。 1559347200。 |  否  | string |
|      end       | 查询结束时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如。 1561939200。 |  否  | string |
|      step      | 时间间隔。 在开始和结束的时间范围内以固定时间间隔检索指标数据。 它需要同时提供 **start** 和 **end**。 格式为[0-9]+[smhdwy]。 默认为 10m（即 10 分钟）。 |  否  | string |
|      time      | Unix 时间格式的时间戳。 在单个时间点检索指标数据。 默认为现在。 时间和开始、结束、步骤的组合是互斥的。 |  否  | string |

响应参数 

|             参数              | 描述 |
| :---------------------------: | :--: |
|       cluster_cpu_total       |      |
|     cluster_memory_total      |      |
|   cluster_pod_running_count   |      |
|    cluster_disk_size_usage    |      |
|  cluster_disk_size_capacity   |      |
|       cluster_pod_quota       |      |
| cluster_memory_usage_wo_cache |      |
|       cluster_cpu_usage       |      |

请求示例

```
{
"metrics_filter": "cluster_cpu_usage%7Ccluster_cpu_total%7Ccluster_memory_usage_wo_cache%7Ccluster_memory_total%7Ccluster_disk_size_usage%7Ccluster_disk_size_capacity%7Ccluster_pod_running_count%7Ccluster_pod_quota%24"
}
```

响应示例：

​       成功情况

```
{
 "results": [
  {
   "metric_name": "cluster_cpu_total",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "6"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_memory_total",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "12380930048"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_pod_running_count",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "metric": {
       "__name__": "cluster:pod_running:count"
      },
      "value": [
       1624327121.488,
       "52"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_disk_size_usage",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "52054761472"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_disk_size_capacity",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "126015406080"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_pod_quota",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "330"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_memory_usage_wo_cache",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "4675477504"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_cpu_usage",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624327121.488,
       "0.323"
      ]
     }
    ]
   }
  }
 ]
}
```

​      失败情况

```
{"code":401,"kind":"Status","apiVersion":"v1","metadata":{},"status":"Failure","message":"Unauthorized: token is expired by 53m54s","reason":"Unauthorized","statusText":"Unauthorized"}
```



## 7、查询集群资源接口

### 应用场景 

> 场景1：xxxxxxxx

### 接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 集群级资源                                                   |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/resources.kubesphere.io/v1alpha3/{resources}          |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/nodes?sortBy=createTime&limit=10 |

请求参数

|   参数    |                           描述                           | 必填 |  类型  |
| :-------: | :------------------------------------------------------: | :--: | :----: |
| resources | 集群级资源类型，比如pods,jobs,nodes,configmaps,service。 |  是  | String |
|   name    |                      用于过滤的名称                      |  否  | string |
|   page    |         页面，形如"page=%d”，默认情况下"page=1“          |  否  | string |
|   limit   |                           上限                           |  否  | string |
| ascending |                排序参数，例如reverse=true                |  否  | string |
|  sortBy   |             排序参数，例如orderBy=createTime             |  否  | string |

响应参数 

|            参数             | 描述 |
| :-------------------------: | :--: |
|         node_load1          |      |
|       node_cpu_total        |      |
|    node_cpu_utilisation     |      |
|   node_memory_utilisation   |      |
|       node_cpu_usage        |      |
| node_memory_usage_wo_cache  |      |
|      node_memory_total      |      |
| node_disk_size_utilisation  |      |
|   node_pod_running_count    |      |
|   node_disk_size_capacity   |      |
| node_disk_inode_utilisation |      |
|    node_disk_size_usage     |      |
|    node_disk_inode_usage    |      |
|       node_pod_quota        |      |
|    node_pod_utilisation     |      |
|    node_disk_inode_total    |      |
|   node:node_inodes_total:   |      |

请求示例

```
{	"type":"rank",   	"metrics_filter" :"node_cpu_utilisation%7Cnode_cpu_usage%7Cnode_cpu_total%7Cnode_memory_utilisation%7				Cnode_memory_usage_wo_cache%7Cnode_memory_total%7Cnode_disk_size_utilisation%7Cnode_disk_size_				usage%7Cnode_disk_size_capacity%7Cnode_pod_utilisation%7Cnode_pod_running_count%7Cnode_pod_quota%7				Cnode_disk_inode_utilisation%7Cnode_disk_inode_total%7Cnode_disk_inode_usage%7Cnode_				load1%24",	"page":"1",	"limit":"5",	"sort_type":"desc",	"sort_metric":"node_cpu_utilisation"}
```

响应示例：

​       成功情况

```
"responses": {    "200": {    	"description": "OK",    "schema": {    "$ref": "#/definitions/api.ListResult"    }}
```

## 8、检索成员的角色模板接口

### 应用场景 

> 场景1：xxxxxxxx

###接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 在集群中检索用户的角色模板                                   |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/iam.kubesphere.io/v1alpha2/clustermembers/{clustermember}/clusterroles |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/iam.kubesphere.io/v1alpha2/clustermembers/wangqi/clusterroles |

请求参数

|     参数      |       描述       | 必填 |  类型  |
| :-----------: | :--------------: | :--: | :----: |
| clustermember | 集群成员的用户名 |  是  | String |

响应参数

|                  参数                  | 描述 |
| :------------------------------------: | :--: |
|       role-template-view-volumes       |      |
|  role-template-view-volume-snapshots   |      |
|   role-template-view-storageclasses    |      |
|      role-template-view-projects       |      |
|        role-template-view-nodes        |      |
|  role-template-view-network-policies   |      |
|       role-template-view-members       |      |
| role-template-view-cluster-monitoring  |      |
|    role-template-view-app-workloads    |      |
|  role-template-view-alerting-policies  |      |
|  role-template-view-alerting-messages  |      |
|      role-template-manage-volumes      |      |
|  role-template-manage-storageclasses   |      |
|       role-template-manage-roles       |      |
|     role-template-manage-projects      |      |
|       role-template-manage-nodes       |      |
| role-template-manage-network-policies  |      |
|      role-template-manage-members      |      |
|       role-template-manage-crds        |      |
|    role-template-manage-components     |      |
| role-template-manage-cluster-settings  |      |
|   role-template-manage-app-workloads   |      |
| role-template-manage-alerting-policies |      |
| role-template-manage-alerting-messages |      |

响应示例：

​       成功情况

```json
 "responses": {  "200": {      "description": "ok",      "schema": {        "$ref": "#/definitions/api.ListResult"      }  }
```

## 9、查询命名空间级指标数据

### 应用场景 

> 场景1：xxxxxxxx

###接口描述

|                        | 内容                                                |
| ---------------------- | --------------------------------------------------- |
| 接口功能               | 获取所有命名空间的命名空间级指标数据                |
| 请求协议               | HTTP                                                |
| 请求方法               | GET                                                 |
| 请求格式               | json                                                |
| 请求url                | /kapis/monitoring.kubesphere.io/v1alpha3/namespaces |
| 请求头(和请求格式对应) | Content-Type: application/json                      |
| 请求示例               |                                                     |

请求参数

|       参数       |                             描述                             | 必填 |  类型  |
| :--------------: | :----------------------------------------------------------: | :--: | :----: |
|  metrics_filter  | 指标名称过滤器由一个正则表达式模式组成。 它指定要返回的指标数据。 例如，以下过滤器匹配命名空间 CPU 使用率和内存使用率：`namespace_cpu_usage|namespace_memory_usage`。 |  否  | String |
| resources_filter | 命名空间过滤器由一个正则表达式模式组成。 它指定要返回的命名空间数据。 例如，以下过滤器匹配命名空间 test 和 kube-system：`test|kube-system`。 |  是  | string |
|      start       | 查询的开始时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，如1559347200 |  否  | string |
|       end        | 查询结束时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如1561939200。 |  否  | string |
|       step       | 时间间隔。 在开始和结束的时间范围内以固定时间间隔检索指标数据。 它需要同时提供 **start** 和 **end**，格式为[0-9]+[smhdwy]。 默认为 10m（即 10 分钟）。 |  否  | string |
|       time       | Unix 时间格式的时间戳。 在单个时间点检索指标数据。 默认为现在。 时间和开始、结束、步骤的组合是互斥的。 |  否  | string |
|   sort_metric    | 按指定的指标对命名空间进行排序。 如果提供了 **start** 和 **end**，则不适用。 |  否  | string |
|    sort_type     |                    排序。 升序或者降序。                     |  否  | string |
|       page       | 页码。 此字段对每个指标的结果数据进行分页，然后返回特定页面。 例如，将 **page** 设置为 2 将返回第二页。 它仅适用于排序的指标数据。 |  否  | string |

响应参数

|                  参数                  | 描述 | 必填 | 类型 |
| :------------------------------------: | :--: | :--: | :--: |
|       role-template-view-volumes       |      |      |      |
|  role-template-view-volume-snapshots   |      |      |      |
|   role-template-view-storageclasses    |      |      |      |
|      role-template-view-projects       |      |      |      |
|        role-template-view-nodes        |      |      |      |
|  role-template-view-network-policies   |      |      |      |
|       role-template-view-members       |      |      |      |
| role-template-view-cluster-monitoring  |      |      |      |
|    role-template-view-app-workloads    |      |      |      |
|  role-template-view-alerting-policies  |      |      |      |
|  role-template-view-alerting-messages  |      |      |      |
|      role-template-manage-volumes      |      |      |      |
|  role-template-manage-storageclasses   |      |      |      |
|       role-template-manage-roles       |      |      |      |
|     role-template-manage-projects      |      |      |      |
|       role-template-manage-nodes       |      |      |      |
| role-template-manage-network-policies  |      |      |      |
|      role-template-manage-members      |      |      |      |
|       role-template-manage-crds        |      |      |      |
|    role-template-manage-components     |      |      |      |
| role-template-manage-cluster-settings  |      |      |      |
|   role-template-manage-app-workloads   |      |      |      |
| role-template-manage-alerting-policies |      |      |      |
| role-template-manage-alerting-messages |      |      |      |

响应示例：

​       成功情况

```
{  "responses": {  	"200": {  	"description": "ok",  		"schema": {  		"$ref": "#/definitions/monitoring.Metrics"  }}
```



## 10、命名空间级资源查询接口

### 应用场景 

> 根据输入命名空间查询

接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 命名空间级资源查询                                           |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace} |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               |                                                              |

请求参数

|       参数       |                             描述                             | 必填 |  类型  |
| :--------------: | :----------------------------------------------------------: | :--: | :----: |
|    namespace     |                           项目名称                           |  是  | String |
| resources_filter | 命名空间过滤器由一个正则表达式模式组成。 它指定要返回的命名空间数据。 例如，以下过滤器匹配命名空间 test 和 kube-system：`test|kube-system`。 |  是  | string |
|      start       | 查询的开始时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，如1559347200 |  否  | string |
|       end        | 查询结束时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如1561939200。 |  否  | string |
|       step       | 时间间隔。 在开始和结束的时间范围内以固定时间间隔检索指标数据。 它需要同时提供 **start** 和 **end**，格式为[0-9]+[smhdwy]。 默认为 10m（即 10 分钟）。 |  否  | string |
|       time       | Unix 时间格式的时间戳。 在单个时间点检索指标数据。 默认为现在。 时间和开始、结束、步骤的组合是互斥的。 |  否  | string |

响应参数

|                  参数                  | 描述 | 必填 | 类型 |
| :------------------------------------: | :--: | :--: | :--: |
|       role-template-view-volumes       |      |      |      |
|  role-template-view-volume-snapshots   |      |      |      |
|   role-template-view-storageclasses    |      |      |      |
|      role-template-view-projects       |      |      |      |
|        role-template-view-nodes        |      |      |      |
|  role-template-view-network-policies   |      |      |      |
|       role-template-view-members       |      |      |      |
| role-template-view-cluster-monitoring  |      |      |      |
|    role-template-view-app-workloads    |      |      |      |
|  role-template-view-alerting-policies  |      |      |      |
|  role-template-view-alerting-messages  |      |      |      |
|      role-template-manage-volumes      |      |      |      |
|  role-template-manage-storageclasses   |      |      |      |
|       role-template-manage-roles       |      |      |      |
|     role-template-manage-projects      |      |      |      |
|       role-template-manage-nodes       |      |      |      |
| role-template-manage-network-policies  |      |      |      |
|      role-template-manage-members      |      |      |      |
|       role-template-manage-crds        |      |      |      |
|    role-template-manage-components     |      |      |      |
| role-template-manage-cluster-settings  |      |      |      |
|   role-template-manage-app-workloads   |      |      |      |
| role-template-manage-alerting-policies |      |      |      |
| role-template-manage-alerting-messages |      |      |      |

响应示例：

​       成功情况

```
{  "responses": {    "200": {      "description": "OK",      "schema": {          "$ref": "#/definitions/monitoring.Metrics"		}	}        }
```







## 11、命名空间级资源查询接口



### 应用场景 

> 根据命名空间和资源类型查询命名空间级别资源



接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 命名空间级资源查询                                           |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources} |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               |                                                              |

请求参数

|   参数    |                            描述                            | 必填 |  类型  |
| :-------: | :--------------------------------------------------------: | :--: | :----: |
| namespace |                          项目名称                          |  是  | String |
| Resources | 命名空间级别的资源类型，例如 pods,jobs,configmaps,services |  是  | string |
|   name    |                       用于过滤的名称                       |  否  | string |
|   page    |                    页面，默认为"page=1"                    |  否  | string |
|   limit   |                            限制                            |  否  | string |
| ascending |                 排序参数，例如reverse=true                 |  否  | string |
|  sortBy   |              排序参数，例如orderBy=createTime              |  否  | string |

响应参数

|                  参数                  | 描述 | 必填 | 类型 |
| :------------------------------------: | :--: | :--: | :--: |
|       role-template-view-volumes       |      |      |      |
|  role-template-view-volume-snapshots   |      |      |      |
|   role-template-view-storageclasses    |      |      |      |
|      role-template-view-projects       |      |      |      |
|        role-template-view-nodes        |      |      |      |
|  role-template-view-network-policies   |      |      |      |
|       role-template-view-members       |      |      |      |
| role-template-view-cluster-monitoring  |      |      |      |
|    role-template-view-app-workloads    |      |      |      |
|  role-template-view-alerting-policies  |      |      |      |
|  role-template-view-alerting-messages  |      |      |      |
|      role-template-manage-volumes      |      |      |      |
|  role-template-manage-storageclasses   |      |      |      |
|       role-template-manage-roles       |      |      |      |
|     role-template-manage-projects      |      |      |      |
|       role-template-manage-nodes       |      |      |      |
| role-template-manage-network-policies  |      |      |      |
|      role-template-manage-members      |      |      |      |
|       role-template-manage-crds        |      |      |      |
|    role-template-manage-components     |      |      |      |
| role-template-manage-cluster-settings  |      |      |      |
|   role-template-manage-app-workloads   |      |      |      |
| role-template-manage-alerting-policies |      |      |      |
| role-template-manage-alerting-messages |      |      |      |

响应示例：

​       成功情况

```
{  "responses": {    "200": {      "description": "OK",      "schema": {          "$ref": "#/definitions/api.ListResult"		}	}        }
```

## 12、命名空间级资源查询接口

### 应用场景 

> 根据命名空间、服务名、资源名等查询命令空间级别的资源



###接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 命名空间级别获取资源查询                                     |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources}/{name} |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               |                                                              |

请求参数

|   参数    |                            描述                            | 必填 |  类型  |
| :-------: | :--------------------------------------------------------: | :--: | :----: |
| namespace |                          项目名称                          |  是  | String |
| Resources | 命名空间级别的资源类型，例如 pods,jobs,configmaps,services |  是  | string |
|   name    |                       用于过滤的名称                       |  是  | string |

响应参数

|                  参数                  | 描述 | 必填 | 类型 |
| :------------------------------------: | :--: | :--: | :--: |
|       role-template-view-volumes       |      |      |      |
|  role-template-view-volume-snapshots   |      |      |      |
|   role-template-view-storageclasses    |      |      |      |
|      role-template-view-projects       |      |      |      |
|        role-template-view-nodes        |      |      |      |
|  role-template-view-network-policies   |      |      |      |
|       role-template-view-members       |      |      |      |
| role-template-view-cluster-monitoring  |      |      |      |
|    role-template-view-app-workloads    |      |      |      |
|  role-template-view-alerting-policies  |      |      |      |
|  role-template-view-alerting-messages  |      |      |      |
|      role-template-manage-volumes      |      |      |      |
|  role-template-manage-storageclasses   |      |      |      |
|       role-template-manage-roles       |      |      |      |
|     role-template-manage-projects      |      |      |      |
|       role-template-manage-nodes       |      |      |      |
| role-template-manage-network-policies  |      |      |      |
|      role-template-manage-members      |      |      |      |
|       role-template-manage-crds        |      |      |      |
|    role-template-manage-components     |      |      |      |
| role-template-manage-cluster-settings  |      |      |      |
|   role-template-manage-app-workloads   |      |      |      |
| role-template-manage-alerting-policies |      |      |      |
| role-template-manage-alerting-messages |      |      |      |

响应示例：

​       成功情况

```
{  "responses": {    "200": {      "description": "OK",      "schema": {          "$ref": "#/definitions/api.ListResult"		}	}        }
```





## 13、查询节点级指标数据接口

### 应用场景

> 场景1：xxxxxxxx

###接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 获取所有节点的节点级指标数据                                 |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/monitoring.kubesphere.io/v1alpha3/nodes               |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/nodes?type=rank&metrics_filter=node_cpu_utilisation%7Cnode_cpu_usage%7Cnode_cpu_total%7Cnode_memory_utilisation%7Cnode_memory_usage_wo_cache%7Cnode_memory_total%7Cnode_disk_size_utilisation%7Cnode_disk_size_usage%7Cnode_disk_size_capacity%7Cnode_pod_utilisation%7Cnode_pod_running_count%7Cnode_pod_quota%7Cnode_disk_inode_utilisation%7Cnode_disk_inode_total%7Cnode_disk_inode_usage%7Cnode_load1%24&page=1&limit=5&sort_type=desc&sort_metric=node_cpu_utilisation |

请求参数

|       参数       |                             描述                             | 必填 |  类型   |
| :--------------: | :----------------------------------------------------------: | :--: | :-----: |
|  metrics_filter  | 指标名称过滤器由一个正则表达式模式组成。 它指定要返回的指标数据。 例如，以下过滤器匹配节点 CPU 使用率和磁盘使用率：`node_cpu_usage|node_disk_size_usage`。 |  否  | string  |
| resources_filter | 节点过滤器由一个正则表达式模式组成。 它指定要返回的节点数据。 例如，以下过滤器匹配节点 i-caojnter 和 i-cmu82ogj：`i-caojnter|i-cmu82ogj`。 |  否  | string  |
|      start       | 查询的开始时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，如1559347200 |  否  | string  |
|       end        | 查询结束时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如1561939200。 |  否  | string  |
|       step       | 时间间隔。 在开始和结束的时间范围内以固定时间间隔检索指标数据。 它需要同时提供 **start** 和 **end**，格式为[0-9]+[smhdwy]。 默认为 10m（即 10 分钟）。 |  否  | string  |
|       time       | Unix 时间格式的时间戳。 在单个时间点检索指标数据。 默认为现在。 时间和开始、结束、步骤的组合是互斥的。 |  否  | string  |
|   sort_metric    | 按指定的度量标准对节点进行排序。 如果**开始**和**结束**则不适用 |  否  | string  |
|    sort_type     |                   排序。 升序、降序之一。                    |  否  | string  |
|       page       | 页码。 此字段对每个指标的结果数据进行分页，然后返回特定页面。 例如，将 **page** 设置为 2 将返回第二页。 它仅适用于排序的指标数据。 |  否  | string  |
|      limit       |        页面大小，单个页面中的最大结果数。 默认为 5。         |  否  | integer |

响应参数

|            参数             | 描述 |
| :-------------------------: | :--: |
| node_memory_usage_wo_cache  |      |
|   node_memory_utilisation   |      |
|    node_cpu_utilisation     |      |
|      node_memory_total      |      |
|       node_cpu_total        |      |
|       node_cpu_usage        |      |
| node_disk_size_utilisation  |      |
|   node_disk_size_capacity   |      |
|    node_pod_utilisation     |      |
|   node_pod_running_count    |      |
|    node_disk_inode_total    |      |
|    node_disk_inode_usage    |      |
|       node_pod_quota        |      |
|    node_disk_size_usage     |      |
| node_disk_inode_utilisation |      |
|         node_load1          |      |

响应示例：

​       成功情况

```
{  "responses": {    "200": {      "description": "OK",      "schema": {          "$ref": "#/definitions/monitoring.Metrics"		}	}        }
```



##  14、查询命名空间pod级别指标数据接口

### 应用场景 

场景1：根据命名空间查询该命名空间下pod级别的指标数据

###接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 获取特定命名空间 pod 的 pod 级指标数据。                     |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/pods |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-system/pods?cluster=default&labelSelector=app.kubernetes.io%2Finstance%3Dks-openldap%2Capp.kubernetes.io%2Fname%3Dopenldap-ha&resources_filter=openldap-0%24&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%24 |

请求参数

|       参数       |                             描述                             | 必填 |  类型   |
| :--------------: | :----------------------------------------------------------: | :--: | :-----: |
|    namespace     |                           项目名称                           |  是  | String  |
| resources_filter | 节点过滤器由一个正则表达式模式组成。 它指定要返回的节点数据。 例如，以下过滤器匹配节点 i-caojnter 和 i-cmu82ogj：`i-caojnter|i-cmu82ogj`。 |  否  | string  |
|      start       | 查询的开始时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，如1559347200 |  否  | string  |
|       end        | 查询结束时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如1561939200。 |  否  | string  |
|       step       | 时间间隔。 在开始和结束的时间范围内以固定时间间隔检索指标数据。 它需要同时提供 **start** 和 **end**，格式为[0-9]+[smhdwy]。 默认为 10m（即 10 分钟）。 |  否  | string  |
|       time       | Unix 时间格式的时间戳。 在单个时间点检索指标数据。 默认为现在。 时间和开始、结束、步骤的组合是互斥的。 |  否  | string  |
|   sort_metric    | 按指定的度量标准对节点进行排序。 如果**开始**和**结束**则不适用 |  否  | string  |
|    sort_type     |                   排序。 升序、降序之一。                    |  否  | string  |
|       page       | 页码。 此字段对每个指标的结果数据进行分页，然后返回特定页面。 例如，将 **page** 设置为 2 将返回第二页。 它仅适用于排序的指标数据。 |  否  | string  |
|      limit       |        页面大小，单个页面中的最大结果数。 默认为 5。         |  否  | integer |

响应参数

|           参数            | 描述 | 必填 | 类型 |
| :-----------------------: | :--: | :--: | :--: |
|       pod_cpu_usage       |      |      |      |
| pod_memory_usage_wo_cache |      |      |      |

响应示例：

​       成功情况

```
{  "responses": {    "200": {      "description": "OK",      "schema": {          "$ref": "#/definitions/monitoring.Metrics"		}	}        }
```



##  15、系统组件组件状态查询接口

### 应用场景

> 根据系统组件查询组件状态



###接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 描述指定的系统组件。                                         |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/resources.kubesphere.io/v1alpha2/components/{component} |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/redis |

请求参数

|   参数    |  描述  | 必填 |  类型  |
| :-------: | :----: | :--: | :----: |
| component | 组件名 |  是  | String |

响应参数

|       参数        | 描述 |
| :---------------: | :--: |
| kubesphere-system |      |

响应示例：

​       成功情况

```
{ "name": "redis", "namespace": "kubesphere-system", "selfLink": "/api/v1/namespaces/kubesphere-system/services/redis", "label": {  "app": "redis",  "tier": "database" }, "startedAt": "2021-04-12T17:46:16+08:00", "totalBackends": 1, "healthyBackends": 1}
```



##  16、缓存查询接口

### 应用场景 

> 查询redis缓存接口



###接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 查询redis缓存接口                                            |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /api/v1/namespaces/kubesphere-system/services/redis          |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/api/v1/namespaces/kubesphere-system/services/redis |
| 备注                   | 官方文档没有                                                 |

请求参数：无参数

响应参数

| 参数  | 描述 |
| :---: | :--: |
| redis |      |

响应示例：

​       成功情况

```
{  "kind": "Service",  "apiVersion": "v1",  "metadata": {    "name": "redis",    "namespace": "kubesphere-system",    "selfLink": "/api/v1/namespaces/kubesphere-system/services/redis",    "uid": "78a16464-f0ea-43fc-b5fd-b9452d15273a",    "resourceVersion": "17836",    "creationTimestamp": "2021-04-12T09:46:16Z",    "labels": {      "app": "redis",      "tier": "database"    },    "annotations": {      "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"redis\",\"tier\":\"database\"},\"name\":\"redis\",\"namespace\":\"kubesphere-system\"},\"spec\":{\"ports\":[{\"name\":\"redis\",\"port\":6379,\"protocol\":\"TCP\",\"targetPort\":6379}],\"selector\":{\"app\":\"redis\",\"tier\":\"database\"},\"sessionAffinity\":\"None\",\"type\":\"ClusterIP\"}}\n"    },    "managedFields": [      {        "manager": "kubectl",        "operation": "Update",        "apiVersion": "v1",        "time": "2021-04-12T09:46:16Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app":{},"f:tier":{}}},"f:spec":{"f:ports":{".":{},"k:{\"port\":6379,\"protocol\":\"TCP\"}":{".":{},"f:name":{},"f:port":{},"f:protocol":{},"f:targetPort":{}}},"f:selector":{".":{},"f:app":{},"f:tier":{}},"f:sessionAffinity":{},"f:type":{}}}      }    ]  },  "spec": {    "ports": [      {        "name": "redis",        "protocol": "TCP",        "port": 6379,        "targetPort": 6379      }    ],    "selector": {      "app": "redis",      "tier": "database"    },    "clusterIP": "10.96.155.75",    "type": "ClusterIP",    "sessionAffinity": "None"  },  "status": {    "loadBalancer": {          }  }}
```



##  17、版本查询接口

### 应用场景

> 场景1：xxxxxxxx



###接口描述

| 描述                   | 内容                                     |
| ---------------------- | ---------------------------------------- |
| 接口功能               | 获取特定命名空间 pod 的 pod 级指标数据。 |
| 请求协议               | HTTP                                     |
| 请求方法               | GET                                      |
| 请求格式               | json                                     |
| 请求url                | /kapis/version                           |
| 请求头(和请求格式对应) | Content-Type: application/json           |
| 请求示例               | http://119.3.157.144:30880/kapis/version |
| 备注                   | 官方文档没有                             |

请求参数：无参数

响应参数

|           参数            | 描述 | 必填 | 类型 |
| :-----------------------: | :--: | :--: | :--: |
|       pod_cpu_usage       |      |      |      |
| pod_memory_usage_wo_cache |      |      |      |

响应示例：

​       成功情况

```
{ "gitVersion": "v0.0.0", "gitMajor": "", "gitMinor": "", "gitCommit": "d567f438ff62e946d8955eb6d49968971d42097f", "gitTreeState": "dirty", "buildDate": "2020-08-28T10:15:00Z", "goVersion": "go1.13.15", "compiler": "gc", "platform": "linux/amd64", "kubernetes": {  "major": "1",  "minor": "18",  "gitVersion": "v1.18.17",  "gitCommit": "68b4e26caf6ede7af577db4af62fb405b4dd47e6",  "gitTreeState": "clean",  "buildDate": "2021-03-18T00:54:02Z",  "goVersion": "go1.13.15",  "compiler": "gc",  "platform": "linux/amd64" }}
```



##  18、编辑集群节点标签接口

### 应用场景 

> 编辑集群节点标签

###接口描述

| 描述                   | 内容                                               |
| ---------------------- | -------------------------------------------------- |
| 接口功能               | 编辑集群节点标签                                   |
| 请求协议               | HTTP                                               |
| 请求方法               | PATCH                                              |
| 请求格式               | json                                               |
| 请求url                | /api/v1/nodes/{resources}                          |
| 请求头(和请求格式对应) | Content-Type: application/json                     |
| 请求示例               | http://119.3.157.144:30880/api/v1/nodes/kubex-0003 |

请求参数

|   参数    |     描述     | 必填 |  类型  |
| :-------: | :----------: | :--: | :----: |
| resources | 集群节点名称 |  是  | String |

响应参数

| 参数 | 描述 |
| :--: | :--: |

响应示例：

​       成功情况

```
{kind: "Node", apiVersion: "v1",…}apiVersion: "v1"kind: "Node"metadata: {name: "kubex-0003", selfLink: "/api/v1/nodes/kubex-0003", uid: "8cf55c0b-4c2f-4582-b621-5534d9d7b581",…}annotations: {flannel.alpha.coreos.com/backend-data: "{"VNI":1,"VtepMAC":"fe:6b:c7:66:65:e4"}",…}flannel.alpha.coreos.com/backend-data: "{\"VNI\":1,\"VtepMAC\":\"fe:6b:c7:66:65:e4\"}"flannel.alpha.coreos.com/backend-type: "vxlan"flannel.alpha.coreos.com/kube-subnet-manager: "true"flannel.alpha.coreos.com/public-ip: "192.168.0.227"kubeadm.alpha.kubernetes.io/cri-socket: "/var/run/dockershim.sock"node.alpha.kubernetes.io/ttl: "0"volumes.kubernetes.io/controller-managed-attach-detach: "true"creationTimestamp: "2021-04-12T05:56:47Z"labels: {beta.kubernetes.io/arch: "amd64", beta.kubernetes.io/os: "linux", kubernetes.io/arch: "amd64",…}managedFields: [{manager: "kubeadm", operation: "Update", apiVersion: "v1", time: "2021-04-12T05:56:47Z",…},…]name: "kubex-0003"resourceVersion: "22570731"selfLink: "/api/v1/nodes/kubex-0003"uid: "8cf55c0b-4c2f-4582-b621-5534d9d7b581"spec: {podCIDR: "10.244.2.0/24", podCIDRs: ["10.244.2.0/24"]}podCIDR: "10.244.2.0/24"podCIDRs: ["10.244.2.0/24"]status: {capacity: {cpu: "2", ephemeral-storage: "41020640Ki", hugepages-1Gi: "0", hugepages-2Mi: "0",…},…}
```



##  19、查看配置文件接口

### 应用场景 

> 根据命名空间查询相应配置文件



###接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 在容器组中查看配置文件                                       |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /api/v1/namespaces/kubesphere-system/pods/{namespace}        |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/api/v1/namespaces/kubesphere-system/pods/ks-console-b4df86d6f-6fbgx |

请求参数

|   参数    |     描述     | 必填 |  类型  |
| :-------: | :----------: | :--: | :----: |
| namespace | 容器组明名称 |  是  | String |

响应参数

|            参数            | 描述 |
| :------------------------: | :--: |
| ks-console-b4df86d6f-6fbgx |      |

响应示例：

​       成功情况

```
{kind: "Pod", apiVersion: "v1",…}apiVersion: "v1"kind: "Pod"metadata: {name: "ks-console-b4df86d6f-6fbgx", generateName: "ks-console-b4df86d6f-",…}spec: {volumes: [{name: "ks-console-config",…},…],…}status: {phase: "Running", conditions: [,…], hostIP: "192.168.0.243", podIP: "10.244.0.85",…}
```

##  20、查询容器日志

### 应用场景 

> 根据命名空间查询容器日志



###接口描述

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 查询容器日志                                                 |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /api/v1/namespaces/kubesphere-system/pods/{namespace}/log    |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/api/v1/namespaces/kubesphere-system/pods/ks-console-b4df86d6f-6fbgx/log?container=ks-console&tailLines=1000&timestamps=true&follow=false |
|                        |                                                              |

请求参数

|    参数    |      描述      | 必填 |  类型  |
| :--------: | :------------: | :--: | :----: |
| namespace  | 命名空间的名称 |  是  | string |
| container  |    容器名称    |  否  | string |
| tailLines  |                |  否  | string |
| timestamps |                |  否  | string |
|   follow   |                |  否  | string |

响应参数

| 参数 |   描述   |
| :--: | :------: |
|      | 日志文件 |

响应示例：

​     返回日志文件


## 21、工作负载查询接口

### 应用场景

​      场景：查询已经部署的容器。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                      查询已经部署的容器                      |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |     /kapis/resources.kubesphere.io/v1alpha3/deployments      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/deployments?sortBy=updateTime&limit=10 |

请求参数

|       参数        |   描述   | 必填 |  类型  |
| :---------------: | :------: | :--: | :----: |
|    **sortBy**     | 排序时间 |  是  | String |
|     **limit**     | 查询数量 |  是  |  int   |
| **labelSelector** | 选择标签 |  否  | String |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{
    "sortBy": "updateTime",
	"limit": 10
}
```

响应示例：

​       成功情况

```json
{
	"items":[]
}
```



## 22、 创建容器部署接口

### 应用场景



### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           创建容器                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             POST                             |
|        请求格式        |                             JSON                             |
|        请求url         |           apis/apps/v1/namespaces/dev/deployments            |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/apps/v1/namespaces/dev/deployments?dryRun=All |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{
  "apiVersion": "apps/v1",
  "kind": "Deployment",
  "metadata": {
    "namespace": "dev",
    "labels": { "app": "cftgh" },
    "name": "cftgh",
    "annotations": {
      "kubesphere.io/alias-name": "三国杀",
      "kubesphere.io/description": "其他人",
      "kubesphere.io/creator": "wangqi"
    }
  },
  "spec": {
    "replicas": 1,
    "selector": { "matchLabels": { "app": "cftgh" } },
    "template": {
      "metadata": {
        "labels": { "app": "cftgh" },
        "annotations": {
          "kubesphere.io/containerSecrets": null,
          "logging.kubesphere.io/logsidecar-config": "{}"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "container-dhab24",
            "imagePullPolicy": "IfNotPresent",
            "image": "hello-world",
            "ports": [
              { "name": "http-5000", "protocol": "TCP", "containerPort": 5000 }
            ],
            "volumeMounts": [
              {
                "name": "volume-t0gz7u",
                "readOnly": true,
                "mountPath": "/data"
              },
              { "name": "volume-zmvp5n", "readOnly": true, "mountPath": "/tmp" }
            ]
          }
        ],
        "serviceAccount": "default",
        "affinity": {},
        "initContainers": [],
        "volumes": [
          {
            "name": "volume-t0gz7u",
            "persistentVolumeClaim": { "claimName": "minio-j9ekzz" }
          },
          {
            "name": "volume-zmvp5n",
            "configMap": { "name": "rabbitm-2jtg9x-rabbitmq" }
          }
        ],
        "imagePullSecrets": null,
        "nodeSelector": { "qwe": "qwe" }
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": "25%",
        "maxSurge": "25%"
      }
    }
  }
}
```

响应示例：

​       成功情况

```json
{
  "kind": "Deployment",
  "apiVersion": "apps/v1",
  "metadata": {
    "name": "cftgh",
    "namespace": "dev",
    "selfLink": "/apis/apps/v1/namespaces/dev/deployments/cftgh",
    "uid": "0191351c-fd41-4df9-b017-8937318340d9",
    "generation": 1,
    "creationTimestamp": "2021-06-22T08:54:28Z",
    "labels": {
      "app": "cftgh"
    },
    "annotations": {
      "kubesphere.io/alias-name": "三国杀",
      "kubesphere.io/creator": "wangqi",
      "kubesphere.io/description": "其他人"
    },
    "managedFields": [
      {
        "manager": "Mozilla",
        "operation": "Update",
        "apiVersion": "apps/v1",
        "time": "2021-06-22T08:54:28Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {
          "f:metadata": {
            "f:annotations": {
              ".": {},
              "f:kubesphere.io/alias-name": {},
              "f:kubesphere.io/creator": {},
              "f:kubesphere.io/description": {}
            },
            "f:labels": { ".": {}, "f:app": {} }
          },
          "f:spec": {
            "f:progressDeadlineSeconds": {},
            "f:replicas": {},
            "f:revisionHistoryLimit": {},
            "f:selector": { "f:matchLabels": { ".": {}, "f:app": {} } },
            "f:strategy": {
              "f:rollingUpdate": {
                ".": {},
                "f:maxSurge": {},
                "f:maxUnavailable": {}
              },
              "f:type": {}
            },
            "f:template": {
              "f:metadata": {
                "f:annotations": {
                  ".": {},
                  "f:kubesphere.io/containerSecrets": {},
                  "f:logging.kubesphere.io/logsidecar-config": {}
                },
                "f:labels": { ".": {}, "f:app": {} }
              },
              "f:spec": {
                "f:affinity": {},
                "f:containers": {
                  "k:{\"name\":\"container-dhab24\"}": {
                    ".": {},
                    "f:image": {},
                    "f:imagePullPolicy": {},
                    "f:name": {},
                    "f:ports": {
                      ".": {},
                      "k:{\"containerPort\":5000,\"protocol\":\"TCP\"}": {
                        ".": {},
                        "f:containerPort": {},
                        "f:name": {},
                        "f:protocol": {}
                      }
                    },
                    "f:resources": {},
                    "f:terminationMessagePath": {},
                    "f:terminationMessagePolicy": {},
                    "f:volumeMounts": {
                      ".": {},
                      "k:{\"mountPath\":\"/data\"}": {
                        ".": {},
                        "f:mountPath": {},
                        "f:name": {},
                        "f:readOnly": {}
                      },
                      "k:{\"mountPath\":\"/tmp\"}": {
                        ".": {},
                        "f:mountPath": {},
                        "f:name": {},
                        "f:readOnly": {}
                      }
                    }
                  }
                },
                "f:dnsPolicy": {},
                "f:nodeSelector": { ".": {}, "f:qwe": {} },
                "f:restartPolicy": {},
                "f:schedulerName": {},
                "f:securityContext": {},
                "f:serviceAccount": {},
                "f:serviceAccountName": {},
                "f:terminationGracePeriodSeconds": {},
                "f:volumes": {
                  ".": {},
                  "k:{\"name\":\"volume-t0gz7u\"}": {
                    ".": {},
                    "f:name": {},
                    "f:persistentVolumeClaim": { ".": {}, "f:claimName": {} }
                  },
                  "k:{\"name\":\"volume-zmvp5n\"}": {
                    ".": {},
                    "f:configMap": {
                      ".": {},
                      "f:defaultMode": {},
                      "f:name": {}
                    },
                    "f:name": {}
                  }
                }
              }
            }
          }
        }
      }
    ]
  },
  "spec": {
    "replicas": 1,
    "selector": {
      "matchLabels": {
        "app": "cftgh"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app": "cftgh"
        },
        "annotations": {
          "kubesphere.io/containerSecrets": "",
          "logging.kubesphere.io/logsidecar-config": "{}"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "volume-t0gz7u",
            "persistentVolumeClaim": {
              "claimName": "minio-j9ekzz"
            }
          },
          {
            "name": "volume-zmvp5n",
            "configMap": {
              "name": "rabbitm-2jtg9x-rabbitmq",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "container-dhab24",
            "image": "hello-world",
            "ports": [
              {
                "name": "http-5000",
                "containerPort": 5000,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "volume-t0gz7u",
                "readOnly": true,
                "mountPath": "/data"
              },
              {
                "name": "volume-zmvp5n",
                "readOnly": true,
                "mountPath": "/tmp"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeSelector": {
          "qwe": "qwe"
        },
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "securityContext": {},
        "affinity": {},
        "schedulerName": "default-scheduler"
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": "25%",
        "maxSurge": "25%"
      }
    },
    "revisionHistoryLimit": 10,
    "progressDeadlineSeconds": 600
  },
  "status": {}
}

```





##  23、 statefulsets接口描述

### 应用场景

### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         有状态副本集                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |     /kapis/resources.kubesphere.io/v1alpha3/statefulsets     |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/statefulsets?sortBy=createTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** | 排序时间 |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{
    "sortBy": "createTime",
	"limit": 10
}
```

响应示例：

​       成功情况

```json
{
	"items":[],
    "totalItems":5
}
```

##  24、daemonsets接口描述

### 应用场景

### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                          守护进程集                          |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |      /kapis/resources.kubesphere.io/v1alpha3/daemonsets      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/daemonsets?sortBy=createTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** | 排序时间 |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{
    "sortBy": "createTime",
	"limit": 10
}
```

响应示例：

​       成功情况

```json
{
	"items":[],
    "totalItems":3
}
```





## 25、任务查询接口

### 应用场景

​      场景：任务 (Job) 负责批量处理短暂的一次性任务，即仅执行一次的任务，它保证批处理任务的一个或多个容器组成功结束。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           查询任务                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |         /kapis/resources.kubesphere.io/v1alpha3/jobs         |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/jobs?sortBy=updateTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** | 排序时间 |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{
    "sortBy": "updateTime",
	"limit": 10
}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 3}
```







##  26、cronjobs接口

### 应用场景

### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        | 定时任务 (CronJob) 管理基于时间的任务，例如在给定时间点只运行一次，或周期性地在给定时间点运行。 |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/resources.kubesphere.io/v1alpha3/cronjobs       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/cronjobs?sortBy=createTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** | 排序时间 |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{    "sortBy": "createTime",	"limit": 10}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 1}
```

##  27、创建任务接口

### 应用场景

### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           创建任务                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             POST                             |
|        请求格式        |                             JSON                             |
|        请求url         |             pis/batch/v1/namespaces/xiaoyan/jobs             |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/batch/v1/namespaces/xiaoyan/jobs?dryRun=All |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{    "apiVersion": "batch/v1",    "kind": "Job",    "metadata": {        "namespace": "xiaoyan",        "labels": {            "app": "testx"        },        "name": "testx",        "annotations": {            "kubesphere.io/alias-name": "testx",            "kubesphere.io/description": "testx",            "kubesphere.io/creator": "wangqi"        }    },    "spec": {        "template": {            "metadata": {                "labels": {                    "app": "testx"                },                "annotations": {                    "kubesphere.io/containerSecrets": null                }            },            "spec": {                "containers": [                    {                        "name": "container-ionw94",                        "imagePullPolicy": "IfNotPresent",                        "image": "hello-world"                    }                ],                "restartPolicy": "Never",                "serviceAccount": "default",                "initContainers": [],                "volumes": [],                "imagePullSecrets": null            }        },        "backoffLimit": 100,        "completions": 80,        "parallelism": 20,        "activeDeadlineSeconds": 10    }}
```

响应示例：

​       成功情况

```json
{
    "kind": "Job",
    "apiVersion": "batch/v1",
    "metadata": {
        "name": "testx",
        "namespace": "xiaoyan",
        "selfLink": "/apis/batch/v1/namespaces/xiaoyan/jobs/testx",
        "uid": "b5060b04-bf67-4cf1-8b51-a0ce8c363168",
        "creationTimestamp": "2021-06-22T09:04:08Z",
        "labels": {
            "app": "testx"
        },
        "annotations": {
            "kubesphere.io/alias-name": "testx",
            "kubesphere.io/creator": "wangqi",
            "kubesphere.io/description": "testx"
        },
        "managedFields": [
            {
                "manager": "Mozilla",
                "operation": "Update",
                "apiVersion": "batch/v1",
                "time": "2021-06-22T09:04:08Z",
                "fieldsType": "FieldsV1",
                "fieldsV1": {
                    "f:metadata": {
                        "f:annotations": {
                            ".": {},
                            "f:kubesphere.io/alias-name": {},
                            "f:kubesphere.io/creator": {},
                            "f:kubesphere.io/description": {}
                        },
                        "f:labels": {
                            ".": {},
                            "f:app": {}
                        }
                    },
                    "f:spec": {
                        "f:activeDeadlineSeconds": {},
                        "f:backoffLimit": {},
                        "f:completions": {},
                        "f:parallelism": {},
                        "f:template": {
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:kubesphere.io/containerSecrets": {}
                                },
                                "f:labels": {
                                    ".": {},
                                    "f:app": {}
                                }
                            },
                            "f:spec": {
                                "f:containers": {
                                    "k:{\"name\":\"container-ionw94\"}": {
                                        ".": {},
                                        "f:image": {},
                                        "f:imagePullPolicy": {},
                                        "f:name": {},
                                        "f:resources": {},
                                        "f:terminationMessagePath": {},
                                        "f:terminationMessagePolicy": {}
                                    }
                                },
                                "f:dnsPolicy": {},
                                "f:restartPolicy": {},
                                "f:schedulerName": {},
                                "f:securityContext": {},
                                "f:serviceAccount": {},
                                "f:serviceAccountName": {},
                                "f:terminationGracePeriodSeconds": {}
                            }
                        }
                    }
                }
            }
        ]
    },
    "spec": {
        "parallelism": 20,
        "completions": 80,
        "activeDeadlineSeconds": 10,
        "backoffLimit": 100,
        "selector": {
            "matchLabels": {
                "controller-uid": "b5060b04-bf67-4cf1-8b51-a0ce8c363168"
            }
        },
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "app": "testx",
                    "controller-uid": "b5060b04-bf67-4cf1-8b51-a0ce8c363168",
                    "job-name": "testx"
                },
                "annotations": {
                    "kubesphere.io/containerSecrets": ""
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "container-ionw94",
                        "image": "hello-world",
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "imagePullPolicy": "IfNotPresent"
                    }
                ],
                "restartPolicy": "Never",
                "terminationGracePeriodSeconds": 30,
                "dnsPolicy": "ClusterFirst",
                "serviceAccountName": "default",
                "serviceAccount": "default",
                "securityContext": {},
                "schedulerName": "default-scheduler"
            }
        }
    },
    "status": {}
}
```




## 28 、容器组查询接口

### 应用场景

​      场景：容器组 (Pod) 是 Kubernetes 应用程序的基本执行单元，是您创建或部署的 Kubernetes 对象模型中最小和最简单的单元。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                       查询部署的容器组                       |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |         /kapis/resources.kubesphere.io/v1alpha3/pods         |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/pods?sortBy=startTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** |   排序   |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{
    "sortBy": "startTime",
	"limit": 10
}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 40}
```


## 29、服务查询接口

### 应用场景

​      场景：服务定义了一类容器组的逻辑集合和一个用于访问它们的策略。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                      查询容器组服务信息                      |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/resources.kubesphere.io/v1alpha3/servies        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/services?sortBy=createTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** |   排序   |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{    "sortBy": "createTime",	"limit": 10}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 40}
```

##  30、 创建服务接口

### 应用场景

### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           创建服务                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             POST                             |
|        请求格式        |                             JSON                             |
|        请求url         |                api/v1/namespaces/dev/services                |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/dev/services?dryRun=All |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{    "apiVersion": "v1",    "kind": "Service",    "metadata": {        "namespace": "dev",        "labels": {            "app": "twstx"        },        "name": "twstx",        "annotations": {            "kubesphere.io/alias-name": "tesxs",            "kubesphere.io/description": "testxt",            "kubesphere.io/creator": "wangqi"        }    },    "spec": {        "sessionAffinity": "ClientIP",        "selector": {            "asd": "asd"        },        "template": {            "metadata": {                "labels": {                    "app": "twstx"                }            }        },        "ports": [            {                "name": "http-2345",                "protocol": "TCP",                "targetPort": 2345,                "port": 2345            }        ],        "type": "NodePort",        "sessionAffinityConfig": {            "clientIP": {                "timeoutSeconds": 10800            }        }    }}
```

响应示例：

​       成功情况

```json
{    "kind": "Service",    "apiVersion": "v1",    "metadata": {        "name": "twstx",        "namespace": "dev",        "selfLink": "/api/v1/namespaces/dev/services/twstx",        "uid": "d0fd90eb-12f2-4551-b236-9d8547eba833",        "creationTimestamp": "2021-06-22T09:19:22Z",        "labels": {            "app": "twstx"        },        "annotations": {            "kubesphere.io/alias-name": "tesxs",            "kubesphere.io/creator": "wangqi",            "kubesphere.io/description": "testxt"        },        "managedFields": [            {                "manager": "Mozilla",                "operation": "Update",                "apiVersion": "v1",                "time": "2021-06-22T09:19:22Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:metadata": {                        "f:annotations": {                            ".": {},                            "f:kubesphere.io/alias-name": {},                            "f:kubesphere.io/creator": {},                            "f:kubesphere.io/description": {}                        },                        "f:labels": {                            ".": {},                            "f:app": {}                        }                    },                    "f:spec": {                        "f:externalTrafficPolicy": {},                        "f:ports": {                            ".": {},                            "k:{\"port\":2345,\"protocol\":\"TCP\"}": {                                ".": {},                                "f:name": {},                                "f:port": {},                                "f:protocol": {},                                "f:targetPort": {}                            }                        },                        "f:selector": {                            ".": {},                            "f:asd": {}                        },                        "f:sessionAffinity": {},                        "f:sessionAffinityConfig": {                            ".": {},                            "f:clientIP": {                                ".": {},                                "f:timeoutSeconds": {}                            }                        },                        "f:type": {}                    }                }            }        ]    },    "spec": {        "ports": [            {                "name": "http-2345",                "protocol": "TCP",                "port": 2345,                "targetPort": 2345,                "nodePort": 32769            }        ],        "selector": {            "asd": "asd"        },        "type": "NodePort",        "sessionAffinity": "ClientIP",        "externalTrafficPolicy": "Cluster",        "sessionAffinityConfig": {            "clientIP": {                "timeoutSeconds": 10800            }        }    },    "status": {        "loadBalancer": {}    }}
```



## 31、应用路由查询接口

### 应用场景

​      场景：应用路由查询模块。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        | 应用路由提供一种聚合服务的方式，将集群的内部服务通过一个外部可访问的 IP 地址暴露给集群外部。 |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |      /kapis/resources.kubesphere.io/v1alpha3/ingresses       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/ingresses?sortBy=createTime&limit=10 |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{    "sortBy": "createTime",	"limit": 10}
```

响应示例：

​       成功情况

```json
{	"items": [],	"totalItems": 0}
```

## 32、 密钥查询接口

### 应用场景

​      场景：密钥 (Secret) 是一种包含少量敏感信息的资源对象，例如密码、token、秘钥等，以键/值对形式保存并且可以在容器组中使用。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         查询密钥信息                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/resources.kubesphere.io/v1alpha3/secrets        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/secrets?sortBy=createTime&limit=1 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** |   排序   |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |

请求示例

```json
{    "sortBy": "createTime",	"limit": 1}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 40}
```


## 33、 配置接口

### 应用场景

​      场景：配置集 (ConfigMap) 常用于存储工作负载所需的配置信息，许多应用程序会从配置文件、命令行参数或环境变量中读取配置信息。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                    查询工作负载所需的配置                    |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |      /kapis/resources.kubesphere.io/v1alpha3/configmaps      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/configmaps?sortBy=createTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** |   排序   |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |


请求示例

```json
{    "sortBy": "updateTime",	"limit": 10}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 40}
```

##  34、创建配置接口

### 应用场景

### 接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           创建配置                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             POST                             |
|        请求格式        |                             JSON                             |
|        请求url         |      /kapis/resources.kubesphere.io/v1alpha3/configmaps      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/xiaoyan/configmaps |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{    "apiVersion": "v1",    "kind": "ConfigMap",    "metadata": {        "namespace": "xiaoyan",        "labels": {},        "name": "asd",        "annotations": {            "kubesphere.io/alias-name": "asd",            "kubesphere.io/description": "asd",            "kubesphere.io/creator": "wangqi"        }    },    "spec": {        "template": {            "metadata": {                "labels": {}            }        }    },    "data": {        "asd": "adsf"    }}
```

响应示例：

​       成功情况

```json
{    "kind": "ConfigMap",    "apiVersion": "v1",    "metadata": {        "name": "asd",        "namespace": "xiaoyan",        "selfLink": "/api/v1/namespaces/xiaoyan/configmaps/asd",        "uid": "a1a29465-4e9f-412a-bdd5-7ffa632c4c3a",        "resourceVersion": "22365161",        "creationTimestamp": "2021-06-22T12:11:20Z",        "annotations": {            "kubesphere.io/alias-name": "asd",            "kubesphere.io/creator": "wangqi",            "kubesphere.io/description": "asd"        },        "managedFields": [            {                "manager": "Mozilla",                "operation": "Update",                "apiVersion": "v1",                "time": "2021-06-22T12:11:20Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:data": {                        ".": {},                        "f:asd": {}                    },                    "f:metadata": {                        "f:annotations": {                            ".": {},                            "f:kubesphere.io/alias-name": {},                            "f:kubesphere.io/creator": {},                            "f:kubesphere.io/description": {}                        }                    }                }            }        ]    },    "data": {        "asd": "adsf"    }}
```



## 35、自定义资源查询接口

### 应用场景

​      场景：自定义资源 (CRD) 是一种 Kubernetes 实现自定义资源类型的扩展方式，用户可以如同操作内置资源对象一样操作 CRD 对象。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                        查询自定义资源                        |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/customresourcedefinitions |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/customresourcedefinitions?sortBy=createTime&limit=10 |

请求参数

|    参数    |   描述   | 必填 |  类型  |
| :--------: | :------: | :--: | :----: |
| **sortBy** |   排序   |  是  | String |
| **limit**  | 查询数量 |  是  |  int   |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |


请求示例

```json
{    "sortBy": "createTime",	"limit": 10}
```

响应示例：

​       成功情况

```json
{    "items": [],	"totalItems": 40}
```



## 36、部署查询接口

### 应用场景

​      场景：部署 (Deployment) 用来描述期望应用达到的目标状态，主要用来描述无状态应用，副本的数量和状态由其背后的控制器来维护，确保状态与定义的期望状态一致。您可以增加副本数量来满足更高负载；回滚部署的版本来消除程序的错误修改；创建自动伸缩器来弹性应对不同场景下的负载。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                      命名空间级资源查询                      |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources} |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods?limit=10&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dwangqi&sortBy=startTime |

请求参数

|       参数        |     描述     | 必填 |  类型  |
| :---------------: | :----------: | :--: | :----: |
|   **ownerKind**   | 有状态副本集 |  是  | String |
| **labelSelector** |  标签选择器  |  是  | String |
|    **sortBy**     |     排序     |  是  | String |
|     **limit**     |   查询数量   |  是  |  int   |

响应参数 

|     参数     |  描述  |
| :----------: | :----: |
| **nodeName** | 节点名 |
|   **name**   | 集群名 |
|  **podIP**   | 节点IP |
|              |        |
|              |        |


请求示例

```json
http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods?limit=10&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dwangqi&sortBy=startTime
```

响应示例：

​       成功情况

```json
{
 "items": [
  {
   "metadata": {
    "name": "kubectl-wangqi-6db7db9448-d9gqd",
    "generateName": "kubectl-wangqi-6db7db9448-",
    "namespace": "kubesphere-controls-system",
    "selfLink": "/api/v1/namespaces/kubesphere-controls-system/pods/kubectl-wangqi-6db7db9448-d9gqd",
    "uid": "254f3888-87f2-4b4c-a656-c97d1e4f06f7",
    "resourceVersion": "21968600",
    "creationTimestamp": "2021-06-21T07:28:25Z",
    "labels": {
     "kubesphere.io/username": "wangqi",
     "pod-template-hash": "6db7db9448"
    },
    "annotations": {
     "kubesphere.io/restartedAt": "2021-06-21T07:28:28.215Z"
    },
    "ownerReferences": [
     {
      "apiVersion": "apps/v1",
      "kind": "ReplicaSet",
      "name": "kubectl-wangqi-6db7db9448",
      "uid": "47579e7d-6da9-47d1-abf1-ac292668c7c8",
      "controller": true,
      "blockOwnerDeletion": true
     }
    ],
    "managedFields": [
     {
      "manager": "kube-controller-manager",
      "operation": "Update",
      "apiVersion": "v1",
      "time": "2021-06-21T07:28:25Z",
      "fieldsType": "FieldsV1",
      "fieldsV1": {
       "f:metadata": {
        "f:annotations": {
         ".": {},
         "f:kubesphere.io/restartedAt": {}
        },
        "f:generateName": {},
        "f:labels": {
         ".": {},
         "f:kubesphere.io/username": {},
         "f:pod-template-hash": {}
        },
        "f:ownerReferences": {
         ".": {},
         "k:{\"uid\":\"47579e7d-6da9-47d1-abf1-ac292668c7c8\"}": {
          ".": {},
          "f:apiVersion": {},
          "f:blockOwnerDeletion": {},
          "f:controller": {},
          "f:kind": {},
          "f:name": {},
          "f:uid": {}
         }
        }
       },
       "f:spec": {
        "f:containers": {
         "k:{\"name\":\"kubectl\"}": {
          ".": {},
          "f:image": {},
          "f:imagePullPolicy": {},
          "f:name": {},
          "f:resources": {},
          "f:terminationMessagePath": {},
          "f:terminationMessagePolicy": {},
          "f:volumeMounts": {
           ".": {},
           "k:{\"mountPath\":\"/etc/localtime\"}": {
            ".": {},
            "f:mountPath": {},
            "f:name": {}
           }
          }
         }
        },
        "f:dnsPolicy": {},
        "f:enableServiceLinks": {},
        "f:restartPolicy": {},
        "f:schedulerName": {},
        "f:securityContext": {},
        "f:serviceAccount": {},
        "f:serviceAccountName": {},
        "f:terminationGracePeriodSeconds": {},
        "f:volumes": {
         ".": {},
         "k:{\"name\":\"host-time\"}": {
          ".": {},
          "f:hostPath": {
           ".": {},
           "f:path": {},
           "f:type": {}
          },
          "f:name": {}
         }
        }
       }
      }
     },
     {
      "manager": "kubelet",
      "operation": "Update",
      "apiVersion": "v1",
      "time": "2021-06-21T07:28:27Z",
      "fieldsType": "FieldsV1",
      "fieldsV1": {
       "f:status": {
        "f:conditions": {
         "k:{\"type\":\"ContainersReady\"}": {
          ".": {},
          "f:lastProbeTime": {},
          "f:lastTransitionTime": {},
          "f:status": {},
          "f:type": {}
         },
         "k:{\"type\":\"Initialized\"}": {
          ".": {},
          "f:lastProbeTime": {},
          "f:lastTransitionTime": {},
          "f:status": {},
          "f:type": {}
         },
         "k:{\"type\":\"Ready\"}": {
          ".": {},
          "f:lastProbeTime": {},
          "f:lastTransitionTime": {},
          "f:status": {},
          "f:type": {}
         }
        },
        "f:containerStatuses": {},
        "f:hostIP": {},
        "f:phase": {},
        "f:podIP": {},
        "f:podIPs": {
         ".": {},
         "k:{\"ip\":\"10.244.1.115\"}": {
          ".": {},
          "f:ip": {}
         }
        },
        "f:startTime": {}
       }
      }
     }
    ]
   },
   "spec": {
    "volumes": [
     {
      "name": "host-time",
      "hostPath": {
       "path": "/etc/localtime",
       "type": ""
      }
     },
     {
      "name": "kubesphere-cluster-admin-token-zc5l2",
      "secret": {
       "secretName": "kubesphere-cluster-admin-token-zc5l2",
       "defaultMode": 420
      }
     }
    ],
    "containers": [
     {
      "name": "kubectl",
      "image": "kubesphere/kubectl:v1.0.0",
      "resources": {},
      "volumeMounts": [
       {
        "name": "host-time",
        "mountPath": "/etc/localtime"
       },
       {
        "name": "kubesphere-cluster-admin-token-zc5l2",
        "readOnly": true,
        "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
       }
      ],
      "terminationMessagePath": "/dev/termination-log",
      "terminationMessagePolicy": "File",
      "imagePullPolicy": "IfNotPresent"
     }
    ],
    "restartPolicy": "Always",
    "terminationGracePeriodSeconds": 30,
    "dnsPolicy": "ClusterFirst",
    "serviceAccountName": "kubesphere-cluster-admin",
    "serviceAccount": "kubesphere-cluster-admin",
    "nodeName": "kubex-0002",
    "securityContext": {},
    "schedulerName": "default-scheduler",
    "tolerations": [
     {
      "key": "node.kubernetes.io/not-ready",
      "operator": "Exists",
      "effect": "NoExecute",
      "tolerationSeconds": 300
     },
     {
      "key": "node.kubernetes.io/unreachable",
      "operator": "Exists",
      "effect": "NoExecute",
      "tolerationSeconds": 300
     }
    ],
    "priority": 0,
    "enableServiceLinks": true
   },
   "status": {
    "phase": "Running",
    "conditions": [
     {
      "type": "Initialized",
      "status": "True",
      "lastProbeTime": null,
      "lastTransitionTime": "2021-06-21T07:28:25Z"
     },
     {
      "type": "Ready",
      "status": "True",
      "lastProbeTime": null,
      "lastTransitionTime": "2021-06-21T07:28:27Z"
     },
     {
      "type": "ContainersReady",
      "status": "True",
      "lastProbeTime": null,
      "lastTransitionTime": "2021-06-21T07:28:27Z"
     },
     {
      "type": "PodScheduled",
      "status": "True",
      "lastProbeTime": null,
      "lastTransitionTime": "2021-06-21T07:28:25Z"
     }
    ],
    "hostIP": "192.168.0.179",
    "podIP": "10.244.1.115",
    "podIPs": [
     {
      "ip": "10.244.1.115"
     }
    ],
    "startTime": "2021-06-21T07:28:25Z",
    "containerStatuses": [
     {
      "name": "kubectl",
      "state": {
       "running": {
        "startedAt": "2021-06-21T07:28:26Z"
       }
      },
      "lastState": {},
      "ready": true,
      "restartCount": 0,
      "image": "kubesphere/kubectl:v1.0.0",
      "imageID": "docker-pullable://kubesphere/kubectl@sha256:e16a56d8c79e76c775ea2d2d2559e4c5d2e78f750f6e075e183948a8dd6e6ebb",
      "containerID": "docker://9d44c4a2ddaf69d20ddb267ae57f38fe6806176dfc988d0b512642b04ef1f963",
      "started": true
     }
    ],
    "qosClass": "BestEffort"
   }
  }
 ],
 "totalItems": 1
}
```


## 37、 自定义资源查询接口

### 应用场景

​      场景：弹性伸缩

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           弹性伸缩                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | prod-api/kapis/clusters/host/resources.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/customresourcedefinitions?sortBy=createTime&limit=10 |

请求参数

|       参数        |     描述     | 必填 |  类型  |
| :---------------: | :----------: | :--: | :----: |
|    **sortBy**     |     排序     |  是  | String |
|     **limit**     |   查询数量   |  是  |  int   |
|   **ownerKind**   | 有状态副本集 |  是  | String |
| **labelSelector** |  标签选择器  |  是  | String |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |
|      |      |
|      |      |


请求示例

```json
http://119.3.157.144:30880/apis/autoscaling/v2beta2/namespaces/kubesphere-controls-system/horizontalpodautoscalers/kubectl-wangqi
```

响应示例：

​       成功情况

```json
{ "items": [  {   "metadata": {    "name": "kubectl-dev-799cdc5745-hz655",    "generateName": "kubectl-dev-799cdc5745-",    "namespace": "kubesphere-controls-system",    "selfLink": "/api/v1/namespaces/kubesphere-controls-system/pods/kubectl-dev-799cdc5745-hz655",    "uid": "f651e3c5-be7c-40ea-bbe1-72ba6881f826",    "resourceVersion": "846161",    "creationTimestamp": "2021-06-09T01:16:36Z",    "labels": {     "kubesphere.io/username": "dev",     "pod-template-hash": "799cdc5745"    },    "ownerReferences": [     {      "apiVersion": "apps/v1",      "kind": "ReplicaSet",      "name": "kubectl-dev-799cdc5745",      "uid": "53a72173-8cad-4d05-9d65-b0b21af3a004",      "controller": true,      "blockOwnerDeletion": true     }    ],    "managedFields": [     {      "manager": "kube-controller-manager",      "operation": "Update",      "apiVersion": "v1",      "time": "2021-06-09T01:16:36Z",      "fieldsType": "FieldsV1",      "fieldsV1": {       "f:metadata": {        "f:generateName": {},        "f:labels": {         ".": {},         "f:kubesphere.io/username": {},         "f:pod-template-hash": {}        },        "f:ownerReferences": {         ".": {},         "k:{\"uid\":\"53a72173-8cad-4d05-9d65-b0b21af3a004\"}": {          ".": {},          "f:apiVersion": {},          "f:blockOwnerDeletion": {},          "f:controller": {},          "f:kind": {},          "f:name": {},          "f:uid": {}         }        }       },       "f:spec": {        "f:containers": {         "k:{\"name\":\"kubectl\"}": {          ".": {},          "f:image": {},          "f:imagePullPolicy": {},          "f:name": {},          "f:resources": {},          "f:terminationMessagePath": {},          "f:terminationMessagePolicy": {},          "f:volumeMounts": {           ".": {},           "k:{\"mountPath\":\"/etc/localtime\"}": {            ".": {},            "f:mountPath": {},            "f:name": {}           }          }         }        },        "f:dnsPolicy": {},        "f:enableServiceLinks": {},        "f:restartPolicy": {},        "f:schedulerName": {},        "f:securityContext": {},        "f:serviceAccount": {},        "f:serviceAccountName": {},        "f:terminationGracePeriodSeconds": {},        "f:volumes": {         ".": {},         "k:{\"name\":\"host-time\"}": {          ".": {},          "f:hostPath": {           ".": {},           "f:path": {},           "f:type": {}          },          "f:name": {}         }        }       }      }     },     {      "manager": "kubelet",      "operation": "Update",      "apiVersion": "v1",      "time": "2021-06-09T01:16:38Z",      "fieldsType": "FieldsV1",      "fieldsV1": {       "f:status": {        "f:conditions": {         "k:{\"type\":\"ContainersReady\"}": {          ".": {},          "f:lastProbeTime": {},          "f:lastTransitionTime": {},          "f:status": {},          "f:type": {}         },         "k:{\"type\":\"Initialized\"}": {          ".": {},          "f:lastProbeTime": {},          "f:lastTransitionTime": {},          "f:status": {},          "f:type": {}         },         "k:{\"type\":\"Ready\"}": {          ".": {},          "f:lastProbeTime": {},          "f:lastTransitionTime": {},          "f:status": {},          "f:type": {}         }        },        "f:containerStatuses": {},        "f:hostIP": {},        "f:phase": {},        "f:podIP": {},        "f:podIPs": {         ".": {},         "k:{\"ip\":\"10.244.2.14\"}": {          ".": {},          "f:ip": {}         }        },        "f:startTime": {}       }      }     }    ]   },   "spec": {    "volumes": [     {      "name": "host-time",      "hostPath": {       "path": "/etc/localtime",       "type": ""      }     },     {      "name": "kubesphere-cluster-admin-token-pdqp8",      "secret": {       "secretName": "kubesphere-cluster-admin-token-pdqp8",       "defaultMode": 420      }     }    ],    "containers": [     {      "name": "kubectl",      "image": "kubesphere/kubectl:v1.19.0",      "resources": {},      "volumeMounts": [       {        "name": "host-time",        "mountPath": "/etc/localtime"       },       {        "name": "kubesphere-cluster-admin-token-pdqp8",        "readOnly": true,        "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"       }      ],      "terminationMessagePath": "/dev/termination-log",      "terminationMessagePolicy": "File",      "imagePullPolicy": "IfNotPresent"     }    ],    "restartPolicy": "Always",    "terminationGracePeriodSeconds": 30,    "dnsPolicy": "ClusterFirst",    "serviceAccountName": "kubesphere-cluster-admin",    "serviceAccount": "kubesphere-cluster-admin",    "nodeName": "jcc-txy-003",    "securityContext": {},    "schedulerName": "default-scheduler",    "tolerations": [     {      "key": "node.kubernetes.io/not-ready",      "operator": "Exists",      "effect": "NoExecute",      "tolerationSeconds": 300     },     {      "key": "node.kubernetes.io/unreachable",      "operator": "Exists",      "effect": "NoExecute",      "tolerationSeconds": 300     }    ],    "priority": 0,    "enableServiceLinks": true,    "preemptionPolicy": "PreemptLowerPriority"   },   "status": {    "phase": "Running",    "conditions": [     {      "type": "Initialized",      "status": "True",      "lastProbeTime": null,      "lastTransitionTime": "2021-06-09T01:16:36Z"     },     {      "type": "Ready",      "status": "True",      "lastProbeTime": null,      "lastTransitionTime": "2021-06-09T01:16:38Z"     },     {      "type": "ContainersReady",      "status": "True",      "lastProbeTime": null,      "lastTransitionTime": "2021-06-09T01:16:38Z"     },     {      "type": "PodScheduled",      "status": "True",      "lastProbeTime": null,      "lastTransitionTime": "2021-06-09T01:16:36Z"     }    ],    "hostIP": "10.206.0.15",    "podIP": "10.244.2.14",    "podIPs": [     {      "ip": "10.244.2.14"     }    ],    "startTime": "2021-06-09T01:16:36Z",    "containerStatuses": [     {      "name": "kubectl",      "state": {       "running": {        "startedAt": "2021-06-09T01:16:37Z"       }      },      "lastState": {},      "ready": true,      "restartCount": 0,      "image": "kubesphere/kubectl:v1.19.0",      "imageID": "docker-pullable://kubesphere/kubectl@sha256:f6851eae93529f70b3a1e7d4fee00b3a49cf548700a91457a4b8096a466d8efe",      "containerID": "docker://effb3d33b12bf8465925c2d526d7040bfab484bd6e73cc3f5bb337e874dd2f53",      "started": true     }    ],    "qosClass": "BestEffort"   }  } ], "totalItems": 1}
```



## 38、任务详情查询接口

### 应用场景

​      场景：自定义资源 (CRD) 是一种 Kubernetes 实现自定义资源类型的扩展方式，用户可以如同操作内置资源对象一样操作 CRD 对象。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         查询任务详情                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/customresourcedefinitions?sortBy=createTime&limit=10 |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |
|      |      |
|      |      |


请求示例

```json
{    "sortBy": "createTime",	"limit": 10}
```

响应示例：

​       成功情况

```json
{  "kind": "Job",  "apiVersion": "batch/v1",  "metadata": {    "name": "hyperpitrix-generate-kubeconfig",    "namespace": "openpitrix-system",    "selfLink": "/apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig",    "uid": "1a33b967-d138-48df-a74d-31511fe51721",    "resourceVersion": "19923",    "creationTimestamp": "2021-04-12T09:47:32Z",    "labels": {      "app": "hyperpitrix",      "job": "hyperpitrix-generate-kubeconfig",      "version": "v0.5.0"    },    "annotations": {      "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"batch/v1\",\"kind\":\"Job\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"hyperpitrix\",\"job\":\"hyperpitrix-generate-kubeconfig\",\"version\":\"v0.5.0\"},\"name\":\"hyperpitrix-generate-kubeconfig\",\"namespace\":\"openpitrix-system\"},\"spec\":{\"backoffLimit\":100,\"completions\":1,\"parallelism\":1,\"template\":{\"metadata\":{\"labels\":{\"app\":\"hyperpitrix\",\"job\":\"hyperpitrix-generate-kubeconfig\",\"version\":\"v0.5.0\"},\"name\":\"hyperpitrix-generate-kubeconfig-job\"},\"spec\":{\"containers\":[{\"command\":[\"generate-kubeconfig\",\"-srv\",\"https://kubernetes.default\"],\"image\":\"openpitrix/generate-kubeconfig:v0.5.0\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"hyperpitrix-generate-kubeconfig\"},{\"command\":[\"migrate-runtime\"],\"image\":\"openpitrix/generate-kubeconfig:v0.5.0\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"hyperpitrix-migrate-runtime\"}],\"dnsPolicy\":\"ClusterFirst\",\"initContainers\":[{\"command\":[\"sh\",\"-c\",\"until nc -z hyperpitrix.openpitrix-system.svc 9103; do echo \\\"waiting for runtime-manager\\\"; sleep 2; done;\"],\"image\":\"alpine:3.10.4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"wait-runtime-manager\"}],\"restartPolicy\":\"OnFailure\",\"schedulerName\":\"default-scheduler\",\"securityContext\":{},\"terminationGracePeriodSeconds\":30}}}}\n",      "revisions": "{\"1\":{\"status\":\"completed\",\"succeed\":1,\"desire\":1,\"uid\":\"1a33b967-d138-48df-a74d-31511fe51721\",\"start-time\":\"2021-04-12T17:47:32+08:00\",\"completion-time\":\"2021-04-12T17:48:54+08:00\"}}"    },    "managedFields": [      {        "manager": "kubectl",        "operation": "Update",        "apiVersion": "batch/v1",        "time": "2021-04-12T09:47:32Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app":{},"f:job":{},"f:version":{}}},"f:spec":{"f:backoffLimit":{},"f:completions":{},"f:parallelism":{},"f:template":{"f:metadata":{"f:labels":{".":{},"f:app":{},"f:job":{},"f:version":{}},"f:name":{}},"f:spec":{"f:containers":{"k:{\"name\":\"hyperpitrix-generate-kubeconfig\"}":{".":{},"f:command":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}},"k:{\"name\":\"hyperpitrix-migrate-runtime\"}":{".":{},"f:command":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}}},"f:dnsPolicy":{},"f:initContainers":{".":{},"k:{\"name\":\"wait-runtime-manager\"}":{".":{},"f:command":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}}},"f:restartPolicy":{},"f:schedulerName":{},"f:securityContext":{},"f:terminationGracePeriodSeconds":{}}}}}      },      {        "manager": "controller-manager",        "operation": "Update",        "apiVersion": "batch/v1",        "time": "2021-04-12T09:48:38Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:metadata":{"f:annotations":{"f:revisions":{}}}}      },      {        "manager": "kube-controller-manager",        "operation": "Update",        "apiVersion": "batch/v1",        "time": "2021-04-12T09:48:54Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:status":{"f:completionTime":{},"f:conditions":{".":{},"k:{\"type\":\"Complete\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}}},"f:startTime":{},"f:succeeded":{}}}      }    ]  },  "spec": {    "parallelism": 1,    "completions": 1,    "backoffLimit": 100,    "selector": {      "matchLabels": {        "controller-uid": "1a33b967-d138-48df-a74d-31511fe51721"      }    },    "template": {      "metadata": {        "name": "hyperpitrix-generate-kubeconfig-job",        "creationTimestamp": null,        "labels": {          "app": "hyperpitrix",          "controller-uid": "1a33b967-d138-48df-a74d-31511fe51721",          "job": "hyperpitrix-generate-kubeconfig",          "job-name": "hyperpitrix-generate-kubeconfig",          "version": "v0.5.0"        }      },      "spec": {        "initContainers": [          {            "name": "wait-runtime-manager",            "image": "alpine:3.10.4",            "command": [              "sh",              "-c",              "until nc -z hyperpitrix.openpitrix-system.svc 9103; do echo \"waiting for runtime-manager\"; sleep 2; done;"            ],            "resources": {                          },            "terminationMessagePath": "/dev/termination-log",            "terminationMessagePolicy": "File",            "imagePullPolicy": "IfNotPresent"          }        ],        "containers": [          {            "name": "hyperpitrix-generate-kubeconfig",            "image": "openpitrix/generate-kubeconfig:v0.5.0",            "command": [              "generate-kubeconfig",              "-srv",              "https://kubernetes.default"            ],            "resources": {                          },            "terminationMessagePath": "/dev/termination-log",            "terminationMessagePolicy": "File",            "imagePullPolicy": "IfNotPresent"          },          {            "name": "hyperpitrix-migrate-runtime",            "image": "openpitrix/generate-kubeconfig:v0.5.0",            "command": [              "migrate-runtime"            ],            "resources": {                          },            "terminationMessagePath": "/dev/termination-log",            "terminationMessagePolicy": "File",            "imagePullPolicy": "IfNotPresent"          }        ],        "restartPolicy": "OnFailure",        "terminationGracePeriodSeconds": 30,        "dnsPolicy": "ClusterFirst",        "securityContext": {                  },        "schedulerName": "default-scheduler"      }    }  },  "status": {    "conditions": [      {        "type": "Complete",        "status": "True",        "lastProbeTime": "2021-04-12T09:48:54Z",        "lastTransitionTime": "2021-04-12T09:48:54Z"      }    ],    "startTime": "2021-04-12T09:47:32Z",    "completionTime": "2021-04-12T09:48:54Z",    "succeeded": 1  }}
```


## 39、密钥编辑接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           编辑密钥                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |            api/v1/namespaces/gaofei/secrets/asdfg            |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/gaofei/secrets/asdfg |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

|            参数            |     描述     |
| :------------------------: | :----------: |
|            data            |  密钥键值对  |
|         namespace          |   项目名称   |
|   kubesphere.io/creator    |  项目创建者  |
|  kubesphere.io/alias-name  |     名称     |
| "kubesphere.io/description | 项目描述信息 |


请求示例

```json

```

响应示例：

​       成功情况

```json
{  "kind": "Secret",  "apiVersion": "v1",  "metadata": {    "name": "asdfg",    "namespace": "gaofei",    "selfLink": "/api/v1/namespaces/gaofei/secrets/asdfg",    "uid": "b3cd50f9-f514-4503-9149-6cf6d96be6cd",    "resourceVersion": "22362340",    "creationTimestamp": "2021-06-22T11:58:54Z",    "annotations": {      "kubesphere.io/alias-name": "asdfg",      "kubesphere.io/creator": "wangqi",      "kubesphere.io/description": "asdfg"    },    "managedFields": [      {        "manager": "Mozilla",        "operation": "Update",        "apiVersion": "v1",        "time": "2021-06-22T11:58:54Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:data":{".":{},"f:asd":{}},"f:metadata":{"f:annotations":{".":{},"f:kubesphere.io/alias-name":{},"f:kubesphere.io/creator":{},"f:kubesphere.io/description":{}}},"f:type":{}}      }    ]  },  "data": {    "asd": "YXNk"  },  "type": "Opaque"}
```

## 40、配置编辑接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         配置信息编辑                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |           api/v1/namespaces/xiaoyan/configmaps/asd           |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/xiaoyan/configmaps/asd |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

|            参数            |     描述     |
| :------------------------: | :----------: |
|            data            |  密钥键值对  |
|         namespace          |   项目名称   |
|   kubesphere.io/creator    |  项目创建者  |
|  kubesphere.io/alias-name  |     名称     |
| "kubesphere.io/description | 项目描述信息 |


请求示例

```json

```

响应示例：

​       成功情况

```json
{  "kind": "ConfigMap",  "apiVersion": "v1",  "metadata": {    "name": "asd",    "namespace": "xiaoyan",    "selfLink": "/api/v1/namespaces/xiaoyan/configmaps/asd",    "uid": "a1a29465-4e9f-412a-bdd5-7ffa632c4c3a",    "resourceVersion": "22365161",    "creationTimestamp": "2021-06-22T12:11:20Z",    "annotations": {      "kubesphere.io/alias-name": "asd",      "kubesphere.io/creator": "wangqi",      "kubesphere.io/description": "asd"    },    "managedFields": [      {        "manager": "Mozilla",        "operation": "Update",        "apiVersion": "v1",        "time": "2021-06-22T12:11:20Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:data":{".":{},"f:asd":{}},"f:metadata":{"f:annotations":{".":{},"f:kubesphere.io/alias-name":{},"f:kubesphere.io/creator":{},"f:kubesphere.io/description":{}}}}      }    ]  },  "data": {    "asd": "adsf"  }}
```

## 41.配置更新接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         更新配置信息                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                            PATCH                             |
|        请求格式        |                             JSON                             |
|        请求url         |           api/v1/namespaces/xiaoyan/configmaps/asd           |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/xiaoyan/configmaps/asd |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

|            参数            |     描述     |
| :------------------------: | :----------: |
|            data            |  密钥键值对  |
|         namespace          |   项目名称   |
|   kubesphere.io/creator    |  项目创建者  |
|  kubesphere.io/alias-name  |     名称     |
| "kubesphere.io/description | 项目描述信息 |


请求示例

```json
{    "kind": "ConfigMap",    "apiVersion": "v1",    "metadata": {        "name": "asd",        "namespace": "xiaoyan",        "annotations": {            "kubesphere.io/alias-name": "asd",            "kubesphere.io/creator": "wangqi",            "kubesphere.io/description": "asd122"        }    },    "data": {        "asd": "adsf"    }}
```

响应示例：

​       成功情况

```json
{  "kind": "ConfigMap",  "apiVersion": "v1",  "metadata": {    "name": "asd",    "namespace": "xiaoyan",    "selfLink": "/api/v1/namespaces/xiaoyan/configmaps/asd",    "uid": "a1a29465-4e9f-412a-bdd5-7ffa632c4c3a",    "resourceVersion": "22618759",    "creationTimestamp": "2021-06-22T12:11:20Z",    "annotations": {      "kubesphere.io/alias-name": "asd",      "kubesphere.io/creator": "wangqi",      "kubesphere.io/description": "asd122"    },    "managedFields": [      {        "manager": "Mozilla",        "operation": "Update",        "apiVersion": "v1",        "time": "2021-06-22T12:11:20Z",        "fieldsType": "FieldsV1",        "fieldsV1": {"f:data":{".":{},"f:asd":{}},"f:metadata":{"f:annotations":{".":{},"f:kubesphere.io/alias-name":{},"f:kubesphere.io/creator":{},"f:kubesphere.io/description":{}}}}      }    ]  },  "data": {    "asd": "adsf"  }}
```

## 42.工作负载部署详情查询接口

### 应用场景

​      场景：部署 (Deployment) 用来描述期望应用达到的目标状态，主要用来描述无状态应用，副本的数量和状态由其背后的控制器来维护，确保状态与定义的期望状态一致。您可以增加副本数量来满足更高负载；回滚部署的版本来消除程序的错误修改；创建自动伸缩器来弹性应对不同场景下的负载。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                   查询工作负载部署详细信息                   |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods?cluster=default&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dwangqi&start=1624428393&end=1624430193&step=60s&times=30&resources_filter=kubectl-wangqi-6db7db9448-d9gqd%24&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%24 |

请求参数

|       参数       |    描述    | 必填 |  类型  |
| :--------------: | :--------: | :--: | :----: |
|     cluster      |    集群    |  是  | String |
|    ownerKind     |            |      | String |
|  labelSelector   | 标签选择器 |  是  | String |
|      start       |    开始    |  是  |  int   |
|       end        |    结束    |  是  |  int   |
|       step       |    间隔    |  是  |  int   |
|      times       |    时间    |  是  |  int   |
| resources_filter |  资源过滤  |  是  | String |
|  metrics_filter  |            |      | String |

响应参数 

|            参数            |     描述     |
| :------------------------: | :----------: |
|            data            |  密钥键值对  |
|         namespace          |   项目名称   |
|   kubesphere.io/creator    |  项目创建者  |
|  kubesphere.io/alias-name  |     名称     |
| "kubesphere.io/description | 项目描述信息 |


请求示例

```json
cluster=default&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dwangqi&start=1624428393&end=1624430193&step=60s&times=30&resources_filter=kubectl-wangqi-6db7db9448-d9gqd%24&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%24
```

响应示例：

​       成功情况

```json
{
  "kind": "ConfigMap",
  "apiVersion": "v1",
  "metadata": {
    "name": "asd",
    "namespace": "xiaoyan",
    "selfLink": "/api/v1/namespaces/xiaoyan/configmaps/asd",
    "uid": "a1a29465-4e9f-412a-bdd5-7ffa632c4c3a",
    "resourceVersion": "22618759",
    "creationTimestamp": "2021-06-22T12:11:20Z",
    "annotations": {
      "kubesphere.io/alias-name": "asd",
      "kubesphere.io/creator": "wangqi",
      "kubesphere.io/description": "asd122"
    },
    "managedFields": [
      {
        "manager": "Mozilla",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-06-22T12:11:20Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:data":{".":{},"f:asd":{}},"f:metadata":{"f:annotations":{".":{},"f:kubesphere.io/alias-name":{},"f:kubesphere.io/creator":{},"f:kubesphere.io/description":{}}}}
      }
    ]
  },
  "data": {
    "asd": "adsf"
  }
}
```

## 43.工作负载部署编辑信息接口

### 应用场景

​      场景：更新部署信息

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                       编辑更新部署信息                       |
|        请求协议        |                             HTTP                             |
|        请求方法        |                            PATCH                             |
|        请求格式        |                             JSON                             |
|        请求url         | kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods?cluster=default&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dwangqi&start=1624428393&end=1624430193&step=60s&times=30&resources_filter=kubectl-wangqi-6db7db9448-d9gqd%24&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%24 |

请求参数

|           参数            | 描述 | 必填 |  类型  |
| :-----------------------: | :--: | :--: | :----: |
| kubesphere.io/description | 描述 |  是  | String |
| kubesphere.io/alias-name  | 别名 |  是  | String |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |

响应参数 

|            参数            |     描述     |
| :------------------------: | :----------: |
|            data            |  密钥键值对  |
|         namespace          |   项目名称   |
|   kubesphere.io/creator    |  项目创建者  |
|  kubesphere.io/alias-name  |     名称     |
| "kubesphere.io/description | 项目描述信息 |


请求示例

```json
{    "kind": "Deployment",    "apiVersion": "apps/v1",    "metadata": {        "name": "kubectl-wangqi",        "namespace": "kubesphere-controls-system",        "annotations": {            "deployment.kubernetes.io/revision": "11",            "kubesphere.io/alias-name": "1114",            "kubesphere.io/description": "我来测试了1111",            "kubesphere.io/relatedHPA": "kubectl-wangqi"        }    },    "spec": {        "replicas": 1,        "selector": {            "matchLabels": {                "kubesphere.io/username": "wangqi"            }        },        "template": {            "metadata": {                "creationTimestamp": null,                "labels": {                    "kubesphere.io/username": "wangqi"                },                "annotations": {                    "kubesphere.io/restartedAt": "2021-06-21T07:28:28.215Z"                }            },            "spec": {                "volumes": [                    {                        "name": "host-time",                        "hostPath": {                            "path": "/etc/localtime",                            "type": ""                        }                    }                ],                "containers": [                    {                        "name": "kubectl",                        "image": "kubesphere/kubectl:v1.0.0",                        "resources": {},                        "volumeMounts": [                            {                                "name": "host-time",                                "mountPath": "/etc/localtime"                            }                        ],                        "terminationMessagePath": "/dev/termination-log",                        "terminationMessagePolicy": "File",                        "imagePullPolicy": "IfNotPresent"                    }                ],                "restartPolicy": "Always",                "terminationGracePeriodSeconds": 30,                "dnsPolicy": "ClusterFirst",                "serviceAccountName": "kubesphere-cluster-admin",                "serviceAccount": "kubesphere-cluster-admin",                "securityContext": {},                "schedulerName": "default-scheduler"            }        },        "strategy": {            "type": "RollingUpdate",            "rollingUpdate": {                "maxUnavailable": "40%",                "maxSurge": "40%"            }        },        "revisionHistoryLimit": 10,        "progressDeadlineSeconds": 600    }}
```

响应示例：

​       成功情况

```json
{    "kind": "Deployment",    "apiVersion": "apps/v1",    "metadata": {        "name": "kubectl-wangqi",        "namespace": "kubesphere-controls-system",        "selfLink": "/apis/apps/v1/namespaces/kubesphere-controls-system/deployments/kubectl-wangqi",        "uid": "eb2ada80-f857-44ba-945e-94a4a266eaf8",        "resourceVersion": "22621413",        "generation": 24,        "creationTimestamp": "2021-05-19T07:40:55Z",        "annotations": {            "deployment.kubernetes.io/revision": "11",            "kubesphere.io/alias-name": "1114",            "kubesphere.io/description": "我来测试了1111",            "kubesphere.io/relatedHPA": "kubectl-wangqi"        },        "managedFields": [            {                "manager": "controller-manager",                "operation": "Update",                "apiVersion": "apps/v1",                "time": "2021-05-19T07:40:55Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:spec": {                        "f:progressDeadlineSeconds": {},                        "f:revisionHistoryLimit": {},                        "f:selector": {                            "f:matchLabels": {                                ".": {},                                "f:kubesphere.io/username": {}                            }                        },                        "f:strategy": {                            "f:rollingUpdate": {},                            "f:type": {}                        },                        "f:template": {                            "f:metadata": {                                "f:labels": {                                    ".": {},                                    "f:kubesphere.io/username": {}                                }                            },                            "f:spec": {                                "f:containers": {                                    "k:{\"name\":\"kubectl\"}": {                                        ".": {},                                        "f:image": {},                                        "f:imagePullPolicy": {},                                        "f:name": {},                                        "f:resources": {},                                        "f:terminationMessagePath": {},                                        "f:terminationMessagePolicy": {},                                        "f:volumeMounts": {                                            ".": {},                                            "k:{\"mountPath\":\"/etc/localtime\"}": {                                                ".": {},                                                "f:mountPath": {},                                                "f:name": {}                                            }                                        }                                    }                                },                                "f:dnsPolicy": {},                                "f:restartPolicy": {},                                "f:schedulerName": {},                                "f:securityContext": {},                                "f:serviceAccount": {},                                "f:serviceAccountName": {},                                "f:terminationGracePeriodSeconds": {},                                "f:volumes": {                                    ".": {},                                    "k:{\"name\":\"host-time\"}": {                                        ".": {},                                        "f:hostPath": {                                            ".": {},                                            "f:path": {},                                            "f:type": {}                                        },                                        "f:name": {}                                    }                                }                            }                        }                    }                }            },            {                "manager": "Mozilla",                "operation": "Update",                "apiVersion": "apps/v1",                "time": "2021-06-18T02:17:15Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:metadata": {                        "f:annotations": {                            "f:kubesphere.io/alias-name": {},                            "f:kubesphere.io/description": {},                            "f:kubesphere.io/relatedHPA": {}                        }                    },                    "f:spec": {                        "f:replicas": {},                        "f:strategy": {                            "f:rollingUpdate": {                                "f:maxSurge": {},                                "f:maxUnavailable": {}                            }                        },                        "f:template": {                            "f:metadata": {                                "f:annotations": {                                    ".": {},                                    "f:kubesphere.io/restartedAt": {}                                }                            }                        }                    }                }            },            {                "manager": "kube-controller-manager",                "operation": "Update",                "apiVersion": "apps/v1",                "time": "2021-06-21T07:28:27Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:metadata": {                        "f:annotations": {                            ".": {},                            "f:deployment.kubernetes.io/revision": {}                        }                    },                    "f:status": {                        "f:availableReplicas": {},                        "f:conditions": {                            "k:{\"type\":\"Progressing\"}": {                                "f:lastUpdateTime": {},                                "f:message": {},                                "f:reason": {}                            }                        },                        "f:observedGeneration": {},                        "f:readyReplicas": {},                        "f:replicas": {},                        "f:updatedReplicas": {}                    }                }            }        ]    },    "spec": {        "replicas": 1,        "selector": {            "matchLabels": {                "kubesphere.io/username": "wangqi"            }        },        "template": {            "metadata": {                "creationTimestamp": null,                "labels": {                    "kubesphere.io/username": "wangqi"                },                "annotations": {                    "kubesphere.io/restartedAt": "2021-06-21T07:28:28.215Z"                }            },            "spec": {                "volumes": [                    {                        "name": "host-time",                        "hostPath": {                            "path": "/etc/localtime",                            "type": ""                        }                    }                ],                "containers": [                    {                        "name": "kubectl",                        "image": "kubesphere/kubectl:v1.0.0",                        "resources": {},                        "volumeMounts": [                            {                                "name": "host-time",                                "mountPath": "/etc/localtime"                            }                        ],                        "terminationMessagePath": "/dev/termination-log",                        "terminationMessagePolicy": "File",                        "imagePullPolicy": "IfNotPresent"                    }                ],                "restartPolicy": "Always",                "terminationGracePeriodSeconds": 30,                "dnsPolicy": "ClusterFirst",                "serviceAccountName": "kubesphere-cluster-admin",                "serviceAccount": "kubesphere-cluster-admin",                "securityContext": {},                "schedulerName": "default-scheduler"            }        },        "strategy": {            "type": "RollingUpdate",            "rollingUpdate": {                "maxUnavailable": "40%",                "maxSurge": "40%"            }        },        "revisionHistoryLimit": 10,        "progressDeadlineSeconds": 600    },    "status": {        "observedGeneration": 23,        "replicas": 1,        "updatedReplicas": 1,        "readyReplicas": 1,        "availableReplicas": 1,        "conditions": [            {                "type": "Available",                "status": "True",                "lastUpdateTime": "2021-06-03T02:16:53Z",                "lastTransitionTime": "2021-06-03T02:16:53Z",                "reason": "MinimumReplicasAvailable",                "message": "Deployment has minimum availability."            },            {                "type": "Progressing",                "status": "True",                "lastUpdateTime": "2021-06-21T07:28:27Z",                "lastTransitionTime": "2021-05-19T07:40:56Z",                "reason": "NewReplicaSetAvailable",                "message": "ReplicaSet \"kubectl-wangqi-6db7db9448\" has successfully progressed."            }        ]    }}
```

## 44.工作负载_部署\_更多操作\_版本回退接口

### 应用场景

​      场景：更新部署信息

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                       编辑更新部署信息                       |
|        请求协议        |                             HTTP                             |
|        请求方法        |                            PATCH                             |
|        请求格式        |                             JSON                             |
|        请求url         | apis/apps/v1/namespaces/kubesphere-controls-system/deployments/kubectl-shangzhi |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/apps/v1/namespaces/kubesphere-controls-system/deployments/kubectl-shangzhi |

请求参数

| 参数  | 描述 | 必填 |  类型  |
| :---: | :--: | :--: | :----: |
|  op   | 操作 |  是  | String |
| path  | 路径 |  是  | String |
| value |  值  |  是  |  Json  |
|       |      |      |        |
|       |      |      |        |
|       |      |      |        |
|       |      |      |        |
|       |      |      |        |
|       |      |      |        |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
[    {        "op": "replace",        "path": "/spec/template",        "value": {            "metadata": {                "creationTimestamp": null,                "labels": {                    "kubesphere.io/username": "shangzhi",                    "pod-template-hash": "687d6cf767"                }            },            "spec": {                "volumes": [                    {                        "name": "host-time",                        "hostPath": {                            "path": "/etc/localtime",                            "type": ""                        }                    }                ],                "containers": [                    {                        "name": "kubectl",                        "image": "kubesphere/kubectl:v1.0.0",                        "resources": {},                        "volumeMounts": [                            {                                "name": "host-time",                                "mountPath": "/etc/localtime"                            }                        ],                        "terminationMessagePath": "/dev/termination-log",                        "terminationMessagePolicy": "File",                        "imagePullPolicy": "IfNotPresent"                    }                ],                "restartPolicy": "Always",                "terminationGracePeriodSeconds": 30,                "dnsPolicy": "ClusterFirst",                "serviceAccountName": "kubesphere-cluster-admin",                "serviceAccount": "kubesphere-cluster-admin",                "securityContext": {},                "schedulerName": "default-scheduler"            }        }    },    {        "op": "replace",        "path": "/metadata/annotations",        "value": {            "deployment.kubernetes.io/revision": "6",            "kubesphere.io/alias-name": "xiazhi",            "kubesphere.io/description": "I'm editing this message."        }    }]
```

响应示例

```json
{kind: "Deployment", apiVersion: "apps/v1",…}apiVersion: "apps/v1"kind: "Deployment"metadata: {name: "kubectl-shangzhi", namespace: "kubesphere-controls-system",…}spec: {replicas: 1, selector: {matchLabels: {kubesphere.io/username: "shangzhi"}},…}status: {observedGeneration: 13, replicas: 1, updatedReplicas: 1, readyReplicas: 1, availableReplicas: 1,…}
```

## 45.工作负载_部署\_更多操作\_弹性伸缩接口

### 应用场景

​      场景：创建自动伸缩器来弹性应对不同场景下的负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |           创建自动伸缩器来弹性应对不同场景下的负载           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             POST                             |
|        请求格式        |                             JSON                             |
|        请求url         | apis/apps/v1/namespaces/kubesphere-controls-system/deployments/kubectl-shangzhi |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/autoscaling/v2beta2/namespaces/kubesphere-controls-system/horizontalpodautoscalers |

请求参数

|         参数         |      描述      | 必填 |  类型  |
| :------------------: | :------------: | :--: | :----: |
| cpuTargetUtilization | cpu目标使用率  |  是  | String |
|  memoryTargetValue   | 内存目标使用量 |  是  | String |
|     minReplicas      |   最小副本数   |  是  |  Int   |
|     maxReplicas      |   最大副本数   |  是  |  Int   |
|                      |                |      |        |
|                      |                |      |        |
|                      |                |      |        |
|                      |                |      |        |
|                      |                |      |        |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{
    "metadata": {
        "name": "kubectl-shangzhi",
        "namespace": "kubesphere-controls-system",
        "annotations": {
            "cpuCurrentUtilization": "0",
            "cpuTargetUtilization": "20",
            "memoryCurrentValue": "0",
            "memoryTargetValue": "20Mi",
            "kubesphere.io/creator": "wangqi"
        }
    },
    "spec": {
        "scaleTargetRef": {
            "apiVersion": "apps/v1",
            "kind": "Deployment",
            "name": "kubectl-shangzhi"
        },
        "minReplicas": 6,
        "maxReplicas": 10,
        "metrics": [
            {
                "type": "Resource",
                "resource": {
                    "name": "memory",
                    "target": {
                        "type": "AverageValue",
                        "averageValue": "20Mi"
                    }
                }
            },
            {
                "type": "Resource",
                "resource": {
                    "name": "cpu",
                    "target": {
                        "type": "Utilization",
                        "averageUtilization": 20
                    }
                }
            }
        ]
    },
    "apiVersion": "autoscaling/v2beta2",
    "kind": "HorizontalPodAutoscaler"
}
```

响应示例

```json
{kind: "HorizontalPodAutoscaler", apiVersion: "autoscaling/v2beta2",…}apiVersion: "autoscaling/v2beta2"kind: "HorizontalPodAutoscaler"metadata: {name: "kubectl-shangzhi", namespace: "kubesphere-controls-system",…}spec: {scaleTargetRef: {kind: "Deployment", name: "kubectl-shangzhi", apiVersion: "apps/v1"}, minReplicas: 6,…}status: {currentReplicas: 0, desiredReplicas: 0, currentMetrics: null, conditions: null}
```

## 46.工作负载_部署\_更多操作\_编辑配置模板接口

### 应用场景

​      场景：编辑配置模板

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         编辑配置模板                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods?cluster=default&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dshangzhi&resources_filter=kubectl-shangzhi-687d6cf767-wdts7%7Ckubectl-shangzhi-687d6cf767-rv8q5%7Ckubectl-shangzhi-687d6cf767-h4rjd%7Ckubectl-shangzhi-687d6cf767-dm7kw%7Ckubectl-shangzhi-687d6cf767-dhqbc%7Ckubectl-shangzhi-687d6cf767-f5tcz%24&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%24 |

请求参数

|       参数       |    描述    | 必填 |  类型  |
| :--------------: | :--------: | :--: | :----: |
|     cluster      |    集群    |  是  | String |
|    ownerKind     |            |      | String |
|  labelSelector   | 标签选择器 |  是  | String |
|      start       |    开始    |  是  |  int   |
|       end        |    结束    |  是  |  int   |
|       step       |    间隔    |  是  |  int   |
|      times       |    时间    |  是  |  int   |
| resources_filter |  资源过滤  |  是  | String |
|  metrics_filter  |            |      | String |

响应参数 

|  参数   |   描述   |
| :-----: | :------: |
| results | 返回结果 |
|         |          |
|         |          |
|         |          |
|         |          |

请求示例

```json
cluster=default&ownerKind=ReplicaSet&labelSelector=kubesphere.io%2Fusername%3Dshangzhi&resources_filter=kubectl-shangzhi-687d6cf767-wdts7%7Ckubectl-shangzhi-687d6cf767-rv8q5%7Ckubectl-shangzhi-687d6cf767-h4rjd%7Ckubectl-shangzhi-687d6cf767-dm7kw%7Ckubectl-shangzhi-687d6cf767-dhqbc%7Ckubectl-shangzhi-687d6cf767-f5tcz%24&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%24
```

响应示例

```json
{results: [{metric_name: "pod_memory_usage_wo_cache", data: {resultType: "vector", result: [,…]}},…]}results: [{metric_name: "pod_memory_usage_wo_cache", data: {resultType: "vector", result: [,…]}},…]0: {metric_name: "pod_memory_usage_wo_cache", data: {resultType: "vector", result: [,…]}}1: {metric_name: "pod_cpu_usage", data: {resultType: "vector", result: [,…]}}
```

## 47.工作负载_部署\_更多操作\_重新部署接口

### 应用场景

​      场景：重新部署

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           重新部署                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                            PATCH                             |
|        请求格式        |                             JSON                             |
|        请求url         | apis/apps/v1/namespaces/kubesphere-controls-system/deployments/kubectl-shangzhi |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/apps/v1/namespaces/kubesphere-controls-system/deployments/kubectl-shangzhi |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | API版本 |
|    kind    |  种类   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{"spec":{"template":{"metadata":{"annotations":{"kubesphere.io/restartedAt":"2021-06-23T08:23:07.546Z"}}}}}
```

响应示例

```json
apiVersion: "apps/v1"kind: "Deployment"metadata: {name: "kubectl-shangzhi", namespace: "kubesphere-controls-system",…}spec: {replicas: 6, selector: {matchLabels: {kubesphere.io/username: "shangzhi"}},…}status: {observedGeneration: 16, replicas: 6, updatedReplicas: 6, readyReplicas: 6, availableReplicas: 6,…}
```

## 48.工作负载_部署\_更多操作\_删除接口

### 应用场景

​      场景：重新部署

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                             删除                             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                            DELETE                            |
|        请求格式        |                             JSON                             |
|        请求url         |  api/v1/namespaces/kubesphere-system/services/ks-apiserver   |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/kubesphere-system/services/ks-apiserver |

请求参数

|       参数        |   描述   | 必填 |  类型  |
| :---------------: | :------: | :--: | :----: |
|       kind        |   类型   |  是  | String |
|    apiVersion     | api版本  |  是  | String |
| propagationPolicy | 传播策略 |  是  | String |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |

响应参数 

|    参数    |   描述   |
| :--------: | :------: |
| apiVersion | API版本  |
|    kind    |   种类   |
|  metadata  |  元数据  |
|  details   | 详细信息 |
|   status   |   状态   |

请求示例

```json
{"kind":"DeleteOptions","apiVersion":"v1","propagationPolicy":"Background"}
```

响应示例

```json
{  "kind": "Status",  "apiVersion": "v1",  "metadata": {      },  "status": "Success",  "details": {    "name": "ks-apiserver",    "kind": "services",    "uid": "be153820-67d5-45f0-98cf-2a649ec523cb"  }}
```

## 49、服务接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |   服务是定义了一类容器组的逻辑集合和一个用于访问它们的策略   |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/resources.kubesphere.io/v1alpha3/services       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/services?sortBy=createTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{    "sortBy": "createTime",    "limit": 10}
```

响应示例：

​       成功情况

```
{items: [,…], totalItems: 27}items: [,…]0: {metadata: {name: "twstx", namespace: "dev", selfLink: "/api/v1/namespaces/dev/services/twstx",…},…}1: {,…}2: {,…}3: {metadata: {name: "notification-manager-svc", namespace: "kubesphere-monitoring-system",…},…}4: {,…}5: {metadata: {name: "prometheus-operated", namespace: "kubesphere-monitoring-system",…},…}6: {metadata: {name: "kubelet", namespace: "kube-system",…},…}7: {metadata: {name: "alertmanager-operated", namespace: "kubesphere-monitoring-system",…},…}8: {metadata: {name: "alertmanager-main", namespace: "kubesphere-monitoring-system",…},…}9: {metadata: {name: "prometheus-k8s", namespace: "kubesphere-monitoring-system",…},…}totalItems: 27
```

​      失败情况

```
{}
```

### 



## 50、服务_创建服务接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           创建服务                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |      /kapis/resources.kubesphere.io/v1alpha3/namespaces      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces?sortBy=createTime&limit=10&labelSelector=%21kubesphere.io%2Fkubefed-host-namespace%2C%21kubesphere.io%2Fdevopsproject |

请求参数

|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
|    sortBy     |      |  是  | String |
|     limit     |      |  是  | String |
| labelSelector |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{     "sortBy": createTime     "limit": 10     "labelSelector": %21kubesphere.io%2Fkubefed-host-namespace%2C%21kubesphere.io%2Fdevopsproject}
```

响应示例：

​       成功情况

```
{items: [{,…}, {,…}, {,…},…], totalItems: 11}items: [{,…}, {,…}, {,…},…]0: {,…}metadata: {name: "xiaoyan", selfLink: "/api/v1/namespaces/xiaoyan", uid: "00642d6a-8072-4377-b011-a67cdd19f15a",…}spec: {finalizers: ["kubernetes"]}status: {phase: "Active"}1: {,…}2: {,…}3: {metadata: {name: "openpitrix-system", selfLink: "/api/v1/namespaces/openpitrix-system",…},…}4: {,…}5: {,…}6: {metadata: {name: "kubesphere-system", selfLink: "/api/v1/namespaces/kubesphere-system",…},…}7: {,…}8: {metadata: {name: "kube-system", selfLink: "/api/v1/namespaces/kube-system",…},…}9: {metadata: {name: "kube-public", selfLink: "/api/v1/namespaces/kube-public",…},…}totalItems: 11
```

​      失败情况

```
{}
```

### 



## 51、服务_服务设置接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           服务设置                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/default/deployments |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/default/deployments?sortBy=updateTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{    "sortBy": updateTime    "limit": 10}
```

响应示例：

​       成功情况

```
{items: [{metadata: {name: "nfs-client-provisioner", namespace: "default",…},…}], totalItems: 1}items: [{metadata: {name: "nfs-client-provisioner", namespace: "default",…},…}]0: {metadata: {name: "nfs-client-provisioner", namespace: "default",…},…}metadata: {name: "nfs-client-provisioner", namespace: "default",…}spec: {replicas: 1, selector: {matchLabels: {app: "nfs-client-provisioner"}},…}status: {observedGeneration: 1, replicas: 1, updatedReplicas: 1, readyReplicas: 1, availableReplicas: 1,…}totalItems: 1
```

​      失败情况

```
{}
```

### 



## 52、服务_事件接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                        存储卷资源状态                        |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |                /api/v1/namespaces/dev/events                 |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/dev/events?fieldSelector=involvedObject.name%3Dtwstx%2CinvolvedObject.namespace%3Ddev%2CinvolvedObject.kind%3DService%2CinvolvedObject.uid%3Daae49974-561c-42ab-b97d-89d129301f0d |



|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
| fieldSelector |      |  是  | String |

响应参数 

|    参数    |                   描述                   |
| :--------: | :--------------------------------------: |
|    kind    |                                          |
| apiVersion | APIVersion定义了对象的这种表示的版本模式 |

请求示例

```
{    "fieldSelector": involvedObject.name%3Dtwstx%2CinvolvedObject.namespace%3Ddev%2CinvolvedObject.kind%3DService%2CinvolvedObject.uid%3Daae49974-561c-42ab-b97d-89d129301f0d}
```

响应示例：

​       成功情况

```
{kind: "EventList", apiVersion: "v1",…}apiVersion: "v1"items: []kind: "EventList"metadata: {selfLink: "/api/v1/namespaces/dev/events", resourceVersion: "22641170"}resourceVersion: "22641170"selfLink: "/api/v1/namespaces/dev/events"
```

​      失败情况

```
{}
```

### 



## 53、服务_资源状态接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         资源状态查看                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods?limit=10&labelSelector=asd%3Dasd111&sortBy=startTime |



|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
|     limit     |      |  是  | String |
| labelSelector |      |  是  | String |
|    sortBy     |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{
    "limit": 10
    "labelSelector": asd%3Dasd111
    "sortBy": startTime
}
```

响应示例：

​       成功情况

```
{
   "items": []
   "totalItems": 0
}
```

​      失败情况

```
{}
```

### 



## 54、服务_编辑信息接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           编辑信息                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |            /api/v1/namespaces/dev/endpoints/twstx            |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/dev/endpoints/twstx |

响应参数 

|    参数    | 描述 |
| :--------: | :--: |
|    kind    |      |
| apiVersion |      |

响应示例：

​       成功情况

```
{kind: "Endpoints", apiVersion: "v1",…}apiVersion: "v1"kind: "Endpoints"metadata: {name: "twstx", namespace: "dev", selfLink: "/api/v1/namespaces/dev/endpoints/twstx",…}
```

​      失败情况

```
{}
```

### 



## 55、服务_编辑服务接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           编辑服务                           |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/deployments |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/deployments?sortBy=updateTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{    "sortBy": "updateTime",    "limit": 10}
```

响应示例：

​       成功情况

```
{items: [{,…}, {metadata: {name: "minio-j9ekzz", namespace: "dev",…},…}], totalItems: 2}items: [{,…}, {metadata: {name: "minio-j9ekzz", namespace: "dev",…},…}]totalItems: 2
```

​      失败情况

```
{}
```

### 



## 56、服务_编辑外网访问接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         编辑外网访问                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods?limit=10&labelSelector=asd%3Dasd111&sortBy=startTime |

请求参数

|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
|     limit     |      |  是  | String |
| labelSelector |      |  是  | String |
|    sortBy     |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{    "limit": 10    "labelSelector": asd%3Dasd111    "sortBy": startTime}
```

响应示例：

​       成功情况

```
{   "items": [],   "totalItems": 0}
```

​      失败情况

```
{}
```

### 



## 57、服务_编辑配置文件接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         编辑配置文件                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods?limit=10&labelSelector=asd%3Dasd111&sortBy=startTime |

请求参数

|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
|     limit     |      |  是  | String |
| labelSelector |      |  是  | String |
|    sortBy     |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{    "limit": 10    "labelSelector": asd%3Dasd111    "sortBy": startTime}
```

响应示例：

​       成功情况

```
{   "items": [],   "totalItems": 0}
```

​      失败情况

```
{}
```

### 





## 58.任务_执行记录详细信息查询接口

### 应用场景

​      场景：查询容器组监控信息

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                      查询容器组监控信息                      |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

|     参数     |   描述   |
| :----------: | :------: |
| parallelism  | 并行数量 |
| completions  | 完成数量 |
| backoffLimit |          |
|              |          |
|              |          |


请求示例

```json

```

响应示例：

​       成功情况

```json
{
  "kind": "Job",
  "apiVersion": "batch/v1",
  "metadata": {
    "name": "hyperpitrix-generate-kubeconfig",
    "namespace": "openpitrix-system",
    "selfLink": "/apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig",
    "uid": "1a33b967-d138-48df-a74d-31511fe51721",
    "resourceVersion": "19923",
    "creationTimestamp": "2021-04-12T09:47:32Z",
    "labels": {
      "app": "hyperpitrix",
      "job": "hyperpitrix-generate-kubeconfig",
      "version": "v0.5.0"
    },
    "annotations": {
      "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"batch/v1\",\"kind\":\"Job\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"hyperpitrix\",\"job\":\"hyperpitrix-generate-kubeconfig\",\"version\":\"v0.5.0\"},\"name\":\"hyperpitrix-generate-kubeconfig\",\"namespace\":\"openpitrix-system\"},\"spec\":{\"backoffLimit\":100,\"completions\":1,\"parallelism\":1,\"template\":{\"metadata\":{\"labels\":{\"app\":\"hyperpitrix\",\"job\":\"hyperpitrix-generate-kubeconfig\",\"version\":\"v0.5.0\"},\"name\":\"hyperpitrix-generate-kubeconfig-job\"},\"spec\":{\"containers\":[{\"command\":[\"generate-kubeconfig\",\"-srv\",\"https://kubernetes.default\"],\"image\":\"openpitrix/generate-kubeconfig:v0.5.0\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"hyperpitrix-generate-kubeconfig\"},{\"command\":[\"migrate-runtime\"],\"image\":\"openpitrix/generate-kubeconfig:v0.5.0\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"hyperpitrix-migrate-runtime\"}],\"dnsPolicy\":\"ClusterFirst\",\"initContainers\":[{\"command\":[\"sh\",\"-c\",\"until nc -z hyperpitrix.openpitrix-system.svc 9103; do echo \\\"waiting for runtime-manager\\\"; sleep 2; done;\"],\"image\":\"alpine:3.10.4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"wait-runtime-manager\"}],\"restartPolicy\":\"OnFailure\",\"schedulerName\":\"default-scheduler\",\"securityContext\":{},\"terminationGracePeriodSeconds\":30}}}}\n",
      "revisions": "{\"1\":{\"status\":\"completed\",\"succeed\":1,\"desire\":1,\"uid\":\"1a33b967-d138-48df-a74d-31511fe51721\",\"start-time\":\"2021-04-12T17:47:32+08:00\",\"completion-time\":\"2021-04-12T17:48:54+08:00\"}}"
    },
    "managedFields": [
      {
        "manager": "kubectl",
        "operation": "Update",
        "apiVersion": "batch/v1",
        "time": "2021-04-12T09:47:32Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app":{},"f:job":{},"f:version":{}}},"f:spec":{"f:backoffLimit":{},"f:completions":{},"f:parallelism":{},"f:template":{"f:metadata":{"f:labels":{".":{},"f:app":{},"f:job":{},"f:version":{}},"f:name":{}},"f:spec":{"f:containers":{"k:{\"name\":\"hyperpitrix-generate-kubeconfig\"}":{".":{},"f:command":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}},"k:{\"name\":\"hyperpitrix-migrate-runtime\"}":{".":{},"f:command":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}}},"f:dnsPolicy":{},"f:initContainers":{".":{},"k:{\"name\":\"wait-runtime-manager\"}":{".":{},"f:command":{},"f:image":{},"f:imagePullPolicy":{},"f:name":{},"f:resources":{},"f:terminationMessagePath":{},"f:terminationMessagePolicy":{}}},"f:restartPolicy":{},"f:schedulerName":{},"f:securityContext":{},"f:terminationGracePeriodSeconds":{}}}}}
      },
      {
        "manager": "controller-manager",
        "operation": "Update",
        "apiVersion": "batch/v1",
        "time": "2021-04-12T09:48:38Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:annotations":{"f:revisions":{}}}}
      },
      {
        "manager": "kube-controller-manager",
        "operation": "Update",
        "apiVersion": "batch/v1",
        "time": "2021-04-12T09:48:54Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:status":{"f:completionTime":{},"f:conditions":{".":{},"k:{\"type\":\"Complete\"}":{".":{},"f:lastProbeTime":{},"f:lastTransitionTime":{},"f:status":{},"f:type":{}}},"f:startTime":{},"f:succeeded":{}}}
      }
    ]
  },
  "spec": {
    "parallelism": 1,
    "completions": 1,
    "backoffLimit": 100,
    "selector": {
      "matchLabels": {
        "controller-uid": "1a33b967-d138-48df-a74d-31511fe51721"
      }
    },
    "template": {
      "metadata": {
        "name": "hyperpitrix-generate-kubeconfig-job",
        "creationTimestamp": null,
        "labels": {
          "app": "hyperpitrix",
          "controller-uid": "1a33b967-d138-48df-a74d-31511fe51721",
          "job": "hyperpitrix-generate-kubeconfig",
          "job-name": "hyperpitrix-generate-kubeconfig",
          "version": "v0.5.0"
        }
      },
      "spec": {
        "initContainers": [
          {
            "name": "wait-runtime-manager",
            "image": "alpine:3.10.4",
            "command": [
              "sh",
              "-c",
              "until nc -z hyperpitrix.openpitrix-system.svc 9103; do echo \"waiting for runtime-manager\"; sleep 2; done;"
            ],
            "resources": {
              
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "containers": [
          {
            "name": "hyperpitrix-generate-kubeconfig",
            "image": "openpitrix/generate-kubeconfig:v0.5.0",
            "command": [
              "generate-kubeconfig",
              "-srv",
              "https://kubernetes.default"
            ],
            "resources": {
              
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          },
          {
            "name": "hyperpitrix-migrate-runtime",
            "image": "openpitrix/generate-kubeconfig:v0.5.0",
            "command": [
              "migrate-runtime"
            ],
            "resources": {
              
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "OnFailure",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "securityContext": {
          
        },
        "schedulerName": "default-scheduler"
      }
    }
  },
  "status": {
    "conditions": [
      {
        "type": "Complete",
        "status": "True",
        "lastProbeTime": "2021-04-12T09:48:54Z",
        "lastTransitionTime": "2021-04-12T09:48:54Z"
      }
    ],
    "startTime": "2021-04-12T09:47:32Z",
    "completionTime": "2021-04-12T09:48:54Z",
    "succeeded": 1
  }
}
```

## 59.任务_资源状态详细信息查询接口

### 应用场景

​      场景：资源状态查看

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         查看资源状态                         |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | kapis/resources.kubesphere.io/v1alpha3/namespaces/openpitrix-system/pods |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/openpitrix-system/pods?limit=10&ownerKind=Job&labelSelector=controller-uid%3D1a33b967-d138-48df-a74d-31511fe51721&sortBy=startTime |

请求参数

|     参数      |    描述    | 必填 |  类型  |
| :-----------: | :--------: | :--: | :----: |
|     limit     |    限制    |  是  |  Int   |
|   ownerKind   |    功能    |  是  | String |
| labelSelector | 标签选择器 |  是  | String |
|    sortBy     |    排序    |  是  | String |
|               |            |      |        |
|               |            |      |        |
|               |            |      |        |
|               |            |      |        |
|               |            |      |        |

响应参数 

|     参数     |   描述   |
| :----------: | :------: |
| parallelism  | 并行数量 |
| completions  | 完成数量 |
| backoffLimit |          |
|              |          |
|              |          |


请求示例

```json
limit=10&ownerKind=Job&labelSelector=controller-uid%3D1a33b967-d138-48df-a74d-31511fe51721&sortBy=startTime
```

响应示例：

​       成功情况

```json
items: [{,…}]0: {,…}metadata: {name: "hyperpitrix-generate-kubeconfig-sbs84", generateName: "hyperpitrix-generate-kubeconfig-",…}spec: {,…}status: {phase: "Succeeded", conditions: [,…], hostIP: "192.168.0.179", podIP: "10.244.1.7",…}totalItems: 1
```



## 60.任务_事件详细信息查询接口

### 应用场景

​      场景：资源状态查看

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         查看事件详情                         |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | kapis/resources.kubesphere.io/v1alpha3/namespaces/openpitrix-system/pods |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/openpitrix-system/events?fieldSelector=involvedObject.name%3Dhyperpitrix-generate-kubeconfig%2CinvolvedObject.namespace%3Dopenpitrix-system%2CinvolvedObject.kind%3DJob%2CinvolvedObject.uid%3D1a33b967-d138-48df-a74d-31511fe51721 |

请求参数

|     参数      |    描述    | 必填 | 类型 |
| :-----------: | :--------: | :--: | :--: |
| fieldSelector | 字段选择器 |  是  | Int  |
|               |            |      |      |
|               |            |      |      |
|               |            |      |      |
|               |            |      |      |
|               |            |      |      |
|               |            |      |      |
|               |            |      |      |
|               |            |      |      |

响应参数 

| 参数  |     描述     |
| :---: | :----------: |
| items | 事件对象列表 |
|       |              |
|       |              |
|       |              |
|       |              |


请求示例

```json
fieldSelector=involvedObject.name%3Dhyperpitrix-generate-kubeconfig%2CinvolvedObject.namespace%3Dopenpitrix-system%2CinvolvedObject.kind%3DJob%2CinvolvedObject.uid%3D1a33b967-d138-48df-a74d-31511fe51721
```

响应示例：

​       成功情况

```json
{  "kind": "EventList",  "apiVersion": "v1",  "metadata": {    "selfLink": "/api/v1/namespaces/openpitrix-system/events",    "resourceVersion": "22629205"  },  "items": []}
```





## 61.任务_编辑信息接口

### 应用场景

​      场景：编辑信息

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         编辑任务信息                         |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                            PATCH                             |
|        请求格式        |                             JSON                             |
|        请求url         | apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |

请求参数

|           参数            | 描述 | 必填 |  类型  |
| :-----------------------: | :--: | :--: | :----: |
| kubesphere.io/alias-name  | 别名 |  是  | String |
| kubesphere.io/description | 描述 |  是  | String |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |
|                           |      |      |        |

响应参数 

| 参数 | 描述 |
| :--: | :--: |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |


请求示例

```json
{    "kind": "Job",    "apiVersion": "batch/v1",    "metadata": {        "name": "hyperpitrix-generate-kubeconfig",        "namespace": "openpitrix-system",        "labels": {            "app": "hyperpitrix",            "job": "hyperpitrix-generate-kubeconfig",            "version": "v0.5.0"        },        "annotations": {            "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"batch/v1\",\"kind\":\"Job\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"hyperpitrix\",\"job\":\"hyperpitrix-generate-kubeconfig\",\"version\":\"v0.5.0\"},\"name\":\"hyperpitrix-generate-kubeconfig\",\"namespace\":\"openpitrix-system\"},\"spec\":{\"backoffLimit\":100,\"completions\":1,\"parallelism\":1,\"template\":{\"metadata\":{\"labels\":{\"app\":\"hyperpitrix\",\"job\":\"hyperpitrix-generate-kubeconfig\",\"version\":\"v0.5.0\"},\"name\":\"hyperpitrix-generate-kubeconfig-job\"},\"spec\":{\"containers\":[{\"command\":[\"generate-kubeconfig\",\"-srv\",\"https://kubernetes.default\"],\"image\":\"openpitrix/generate-kubeconfig:v0.5.0\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"hyperpitrix-generate-kubeconfig\"},{\"command\":[\"migrate-runtime\"],\"image\":\"openpitrix/generate-kubeconfig:v0.5.0\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"hyperpitrix-migrate-runtime\"}],\"dnsPolicy\":\"ClusterFirst\",\"initContainers\":[{\"command\":[\"sh\",\"-c\",\"until nc -z hyperpitrix.openpitrix-system.svc 9103; do echo \\\"waiting for runtime-manager\\\"; sleep 2; done;\"],\"image\":\"alpine:3.10.4\",\"imagePullPolicy\":\"IfNotPresent\",\"name\":\"wait-runtime-manager\"}],\"restartPolicy\":\"OnFailure\",\"schedulerName\":\"default-scheduler\",\"securityContext\":{},\"terminationGracePeriodSeconds\":30}}}}\n",            "revisions": "{\"1\":{\"status\":\"completed\",\"succeed\":1,\"desire\":1,\"uid\":\"1a33b967-d138-48df-a74d-31511fe51721\",\"start-time\":\"2021-04-12T17:47:32+08:00\",\"completion-time\":\"2021-04-12T17:48:54+08:00\"}}",            "kubesphere.io/alias-name": "求为",            "kubesphere.io/description": "气温"        }    },    "spec": {        "parallelism": 1,        "completions": 1,        "backoffLimit": 100,        "selector": {            "matchLabels": {                "controller-uid": "1a33b967-d138-48df-a74d-31511fe51721"            }        },        "template": {            "metadata": {                "name": "hyperpitrix-generate-kubeconfig-job",                "creationTimestamp": null,                "labels": {                    "app": "hyperpitrix",                    "controller-uid": "1a33b967-d138-48df-a74d-31511fe51721",                    "job": "hyperpitrix-generate-kubeconfig",                    "job-name": "hyperpitrix-generate-kubeconfig",                    "version": "v0.5.0"                }            },            "spec": {                "initContainers": [                    {                        "name": "wait-runtime-manager",                        "image": "alpine:3.10.4",                        "command": [                            "sh",                            "-c",                            "until nc -z hyperpitrix.openpitrix-system.svc 9103; do echo \"waiting for runtime-manager\"; sleep 2; done;"                        ],                        "resources": {},                        "terminationMessagePath": "/dev/termination-log",                        "terminationMessagePolicy": "File",                        "imagePullPolicy": "IfNotPresent"                    }                ],                "containers": [                    {                        "name": "hyperpitrix-generate-kubeconfig",                        "image": "openpitrix/generate-kubeconfig:v0.5.0",                        "command": [                            "generate-kubeconfig",                            "-srv",                            "https://kubernetes.default"                        ],                        "resources": {},                        "terminationMessagePath": "/dev/termination-log",                        "terminationMessagePolicy": "File",                        "imagePullPolicy": "IfNotPresent"                    },                    {                        "name": "hyperpitrix-migrate-runtime",                        "image": "openpitrix/generate-kubeconfig:v0.5.0",                        "command": [                            "migrate-runtime"                        ],                        "resources": {},                        "terminationMessagePath": "/dev/termination-log",                        "terminationMessagePolicy": "File",                        "imagePullPolicy": "IfNotPresent"                    }                ],                "restartPolicy": "OnFailure",                "terminationGracePeriodSeconds": 30,                "dnsPolicy": "ClusterFirst",                "securityContext": {},                "schedulerName": "default-scheduler"            }        }    }}
```

响应示例：

​       成功情况

```json
{kind: "Job", apiVersion: "batch/v1",…}apiVersion: "batch/v1"kind: "Job"metadata: {name: "hyperpitrix-generate-kubeconfig", namespace: "openpitrix-system",…}spec: {parallelism: 1, completions: 1, backoffLimit: 100,…}status: {conditions: [{type: "Complete", status: "True", lastProbeTime: "2021-04-12T09:48:54Z",…}],…}
```



## 62.任务_更多操作\_重新执行接口

### 应用场景

​      场景：任务重新执行

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         重新执行任务                         |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数     |   描述   |
| :---------: | :------: |
| parallelism | 并行数量 |
| completions | 完成数量 |
|             |          |
|             |          |
|             |          |


请求示例

```json

```

响应示例：

​       成功情况

```json
{kind: "Job", apiVersion: "batch/v1",…}apiVersion: "batch/v1"kind: "Job"metadata: {name: "hyperpitrix-generate-kubeconfig", namespace: "openpitrix-system",…}spec: {parallelism: 1, completions: 1, backoffLimit: 100,…}status: {conditions: [{type: "Complete", status: "True", lastProbeTime: "2021-06-23T07:21:45Z",…}],…}
```

## 63.任务_更多操作\_删除接口

### 应用场景

​      场景：删除任务

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                           删除任务                           |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                            DELETE                            |
|        请求格式        |                             JSON                             |
|        请求url         | apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/batch/v1/namespaces/openpitrix-system/jobs/hyperpitrix-generate-kubeconfig |

请求参数

|       参数        |   描述   | 必填 |  类型  |
| :---------------: | :------: | :--: | :----: |
|       kind        |   类型   |  是  | String |
|    apiVersion     | api版本  |  是  | String |
| propagationPolicy | 传播策略 |  是  | String |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |
|                   |          |      |        |

响应参数 

|  参数   |   描述   |
| :-----: | :------: |
| status  |   状态   |
| details | 详细信息 |
|         |          |
|         |          |
|         |          |


请求示例

```json
{"kind":"DeleteOptions","apiVersion":"v1","propagationPolicy":"Background"}
```

响应示例：

​       成功情况

```json
{  "kind": "Status",  "apiVersion": "v1",  "metadata": {      },  "status": "Success",  "details": {    "name": "hyperpitrix-generate-kubeconfig",    "group": "batch",    "kind": "jobs",    "uid": "e44d81de-83ec-42d6-a9fa-5389f8af6742"  }}
```









## 64、应用路由接口

### 应用场景

​      场景：应用负载

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        | 应用路由提供一种聚合服务的方式，您可以将集群的内部服务通过一个外部可访问的 IP 地址暴露给集群外部。 |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |      /kapis/resources.kubesphere.io/v1alpha3/ingresses       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/ingresses?sortBy=createTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例：

```
{
   "sortBy": createTime
   "limit": 10
}
```

响应示例：

​       成功情况

```
{
   "items": []
   "totalItems": 0
}
```

​      失败情况

```
{
}
```

### 





## 65.容器组监控信息查询接口

### 应用场景

​      场景：查询容器组监控信息

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                      查询容器组监控信息                      |
|        请求协议        |                             HTTP                             |
|      GGET请求方法      |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods/kubectl-wangqi-6db7db9448-d9gqd |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-controls-system/pods/kubectl-wangqi-6db7db9448-d9gqd?cluster=default&start=1624401033&end=1624431033&step=600s&times=50&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%7Cpod_net_bytes_transmitted%7Cpod_net_bytes_received%24 |

请求参数

|      参数      |    描述    | 必填 |  类型  |
| :------------: | :--------: | :--: | :----: |
|    cluster     |    集群    |  是  | String |
|     start      | 起始时间戳 |  是  |  int   |
|      end       | 结束时间戳 |  是  |  int   |
|      step      |  间隔时间  |  是  | String |
|     times      |    次数    |  是  |  int   |
| metrics_filter |            |      |        |
|                |            |      |        |
|                |            |      |        |
|                |            |      |        |

响应参数 

|  参数  |   描述   |
| :----: | :------: |
| result | 请求结果 |
|        |          |
|        |          |
|        |          |
|        |          |


请求示例

```json
cluster=default&start=1624401033&end=1624431033&step=600s&times=50&metrics_filter=pod_cpu_usage%7Cpod_memory_usage_wo_cache%7Cpod_net_bytes_transmitted%7Cpod_net_bytes_received%24           
```

响应示例：

​       成功情况

```json
results: [{metric_name: "pod_memory_usage_wo_cache", data: {resultType: "matrix", result: [,…]}},…]
0: {metric_name: "pod_memory_usage_wo_cache", data: {resultType: "matrix", result: [,…]}}
1: {metric_name: "pod_net_bytes_transmitted", data: {resultType: "matrix", result: [,…]}}
2: {metric_name: "pod_net_bytes_received", data: {resultType: "matrix", result: [,…]}}
3: {metric_name: "pod_cpu_usage", data: {resultType: "matrix", result: [,…]}}
```





## 66、存储卷接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        | 存储卷供用户创建的工作负载使用，是将工作负载数据持久化的一种资源对象 |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/persistentvolumeclaims |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/persistentvolumeclaims?sortBy=createTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{
    "sortBy": "createTime",
    "limit": 10
}
```

响应示例：

​       成功情况

```
{
    "items": [{metadata: {name: "data-rabbitm-2jtg9x-rabbitmq-0", namespace: "dev",…},…},…], 
    "totalItems": 10
}
```



## 67、快照信息接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                          存储卷快照                          |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/volumesnapshots |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/volumesnapshots?persistentVolumeClaimName=data-rabbitm-2jtg9x-rabbitmq-0&sortBy=createTime&limit=10 |



|           参数            | 描述 | 必填 |  类型  |
| :-----------------------: | :--: | :--: | :----: |
| persistentVolumeClaimName |      |  是  | String |
|          sortBy           |      |  是  | String |
|           limit           |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{
    "persistentVolumeClaimName": data-rabbitm-2jtg9x-rabbitmq-0
    "sortBy": createTime
    "limit": 10
}
```

响应示例：

​       成功情况

```
{
    "items": []
    "totalItems": 0
}
```

   

## 68、事件接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                          存储卷事件                          |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |                /api/v1/namespaces/dev/events                 |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/dev/events?fieldSelector=involvedObject.name%3Ddata-rabbitm-2jtg9x-rabbitmq-0%2CinvolvedObject.namespace%3Ddev%2CinvolvedObject.kind%3DPersistentVolumeClaim%2CinvolvedObject.uid%3Dba1cfc73-c4ae-4f1b-bc6d-eddb7634940f |



|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
| fieldSelector |      |  是  | String |

响应参数 

|    参数    |                   描述                   |
| :--------: | :--------------------------------------: |
|    kind    |                                          |
| apiVersion | APIVersion定义了对象的这种表示的版本模式 |

请求示例

```
{
    "fieldSelector": involvedObject.name%3Ddata-rabbitm-2jtg9x-rabbitmq-  0%2CinvolvedObject.namespace%3Ddev%2CinvolvedObject.kind%3DPersistentVolumeClaim%2CinvolvedObject.uid%3Dba1cfc73-c4ae-4f1b-bc6d-eddb7634940f
}
```

响应示例：

​       成功情况

```
{kind: "EventList", apiVersion: "v1",…}
   "apiVersion": "v1"
   "items": []
   "kind": "EventList"
   "metadata": {selfLink: "/api/v1/namespaces/dev/events", resourceVersion: "22314304"}
     "resourceVersion": "22314304"
     "selfLink": "/api/v1/namespaces/dev/events"
```



## 69、资源状态接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                        存储卷资源状态                        |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods?limit=10&pvcName=data-rabbitm-2jtg9x-rabbitmq-0&sortBy=startTime |



|  参数   | 描述 | 必填 |  类型  |
| :-----: | :--: | :--: | :----: |
|  limit  |      |  是  | String |
| pvcName |      |  是  | String |
| sortBy  |      |  是  | String |

响应参数 

|    参数    |                   描述                   |
| :--------: | :--------------------------------------: |
|    kind    |                                          |
| apiVersion | APIVersion定义了对象的这种表示的版本模式 |

请求示例

```
{
    "limit": 10
    "pvcName": data-rabbitm-2jtg9x-rabbitmq-0
    "sortBy": startTime
}
```

响应示例：

​       成功情况

```
{
   "items": []
   "totalItems": 0
}
```



## 70、编辑配置文件接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                        存储卷编辑配置                        |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/namespaces/dev/pods?limit=10&pvcName=data-rabbitm-2jtg9x-rabbitmq-0&sortBy=startTime |



|  参数   | 描述 | 必填 |  类型  |
| :-----: | :--: | :--: | :----: |
|  limit  |      |  是  | String |
| pvcName |      |  是  | String |
| sortBy  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{
    "limit": 10
    "pvcName": data-rabbitm-2jtg9x-rabbitmq-0
    "sortBy": startTime
}
```

响应示例：

​       成功情况

```
{
   "items": []
   "totalItems": 0
}
```





## 71、存储卷快照接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        | 存储卷快照表示存储卷的时间点副本。快照可用于配置新卷（预先填充快照数据）或将现有存储卷还原到先前状态（由快照表示） |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha3/volumesnapshots    |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/volumesnapshots?sortBy=createTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{
    "sortBy": "createTime",
    "limit": 10
}
```

响应示例：

​       成功情况

```
{
    "items": []
    "totalItems": 0
}
```



## 72、存储类型接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        | 存储类型 (StorageClass) 是由集群管理员配置存储服务端参数，并按类型提供存储给集群用户使用 |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |    /kapis/resources.kubesphere.io/v1alpha3/storageclasses    |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/storageclasses?sortBy=createTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

响应参数 

|    参数    |  描述  |
| :--------: | :----: |
|   items    |  条目  |
| totalItems | 总条目 |

请求示例

```
{    "sortBy": "createTime",    "limit": 10}
```

响应示例：

​       成功情况

```
{
    "items": [{,…}]
    "totalItems": 1
}
```



## 73、查看配置文件接口

### 应用场景

​      场景：存储管理

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |             存储类型 (StorageClass) 查看配置文件             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |  /apis/storage.k8s.io/v1/storageclasses/managed-nfs-storage  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/apis/storage.k8s.io/v1/storageclasses/managed-nfs-storage |

响应参数 

|    参数    |                   描述                   |
| :--------: | :--------------------------------------: |
|    kind    |                                          |
| apiVersion | APIVersion定义了对象的这种表示的版本模式 |

响应示例：

​       成功情况

```
{kind: "StorageClass", apiVersion: "storage.k8s.io/v1",…}
  apiVersion: "storage.k8s.io/v1"
  kind: "StorageClass"
  metadata: {name: "managed-nfs-storage", selfLink: "/apis/storage.k8s.io/v1/storageclasses/managed-nfs-storage",…}
  parameters: {archiveOnDelete: "false"}
  provisioner: "fuseim.pri/ifs"
  reclaimPolicy: "Delete"
  volumeBindingMode: "Immediate"
```



## 74、集群状态接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                      监控集群的运行状态                      |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha2/componenthealth    |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/componenthealth |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      | 开始 |  是  | String |
|      end       | 结束 |  是  | String |
|      step      | 步骤 |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

响应参数 

|       参数       | 描述 |
| :--------------: | :--: |
| kubesphereStatus | 状态 |

响应示例：

​       成功情况

```
{
 "kubesphereStatus": [
  {
   "name": "etcd",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/etcd",
   "label": {
    "app": "kubesphere",
    "tier": "etcd"
   },
   "startedAt": "2021-04-12T17:46:42+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "redis",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/redis",
   "label": {
    "app": "redis",
    "tier": "database"
   },
   "startedAt": "2021-04-12T17:46:16+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "mysql",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/mysql",
   "label": {
    "app": "kubesphere",
    "tier": "db"
   },
   "startedAt": "2021-04-12T17:46:42+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "ks-console",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-console",
   "label": {
    "app": "ks-console",
    "tier": "frontend",
    "version": "v3.0.0"
   },
   "startedAt": "2021-04-12T17:47:08+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "ks-controller-manager",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-controller-manager",
   "label": {
    "app": "ks-controller-manager",
    "tier": "backend",
    "version": "v3.0.0"
   },
   "startedAt": "2021-04-12T17:47:06+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "openldap",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/openldap",
   "label": {
    "app.kubernetes.io/instance": "ks-openldap",
    "app.kubernetes.io/name": "openldap-ha"
   },
   "startedAt": "2021-04-12T17:46:21+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "ks-apiserver",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-apiserver",
   "label": {
    "app": "ks-apiserver",
    "tier": "backend",
    "version": "v3.0.0"
   },
   "startedAt": "2021-04-12T17:47:06+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "minio",
   "namespace": "kubesphere-system",
   "selfLink": "/api/v1/namespaces/kubesphere-system/services/minio",
   "label": {
    "app": "minio",
    "release": "ks-minio"
   },
   "startedAt": "2021-04-12T17:46:29+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "kube-state-metrics",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/kube-state-metrics",
   "label": {
    "app.kubernetes.io/name": "kube-state-metrics"
   },
   "startedAt": "2021-04-12T17:48:05+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "alertmanager-operated",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/alertmanager-operated",
   "label": {
    "app": "alertmanager"
   },
   "startedAt": "2021-04-12T17:48:12+08:00",
   "totalBackends": 3,
   "healthyBackends": 3
  },
  {
   "name": "prometheus-k8s",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/prometheus-k8s",
   "label": {
    "app": "prometheus",
    "prometheus": "k8s"
   },
   "startedAt": "2021-04-12T17:48:10+08:00",
   "totalBackends": 2,
   "healthyBackends": 2
  },
  {
   "name": "prometheus-operated",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/prometheus-operated",
   "label": {
    "app": "prometheus"
   },
   "startedAt": "2021-04-12T17:48:21+08:00",
   "totalBackends": 2,
   "healthyBackends": 2
  },
  {
   "name": "notification-manager-controller-metrics",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/notification-manager-controller-metrics",
   "label": {
    "control-plane": "controller-manager"
   },
   "startedAt": "2021-04-12T17:48:26+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "notification-manager-svc",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/notification-manager-svc",
   "label": {
    "app": "notification-manager",
    "notification-manager": "notification-manager"
   },
   "startedAt": "2021-04-12T17:48:30+08:00",
   "totalBackends": 2,
   "healthyBackends": 2
  },
  {
   "name": "alertmanager-main",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/alertmanager-main",
   "label": {
    "alertmanager": "main",
    "app": "alertmanager"
   },
   "startedAt": "2021-04-12T17:48:12+08:00",
   "totalBackends": 3,
   "healthyBackends": 3
  },
  {
   "name": "node-exporter",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/node-exporter",
   "label": {
    "app.kubernetes.io/name": "node-exporter"
   },
   "startedAt": "2021-04-12T17:48:04+08:00",
   "totalBackends": 3,
   "healthyBackends": 3
  },
  {
   "name": "prometheus-operator",
   "namespace": "kubesphere-monitoring-system",
   "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/prometheus-operator",
   "label": {
    "app.kubernetes.io/component": "controller",
    "app.kubernetes.io/name": "prometheus-operator"
   },
   "startedAt": "2021-04-12T17:48:03+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "hyperpitrix",
   "namespace": "openpitrix-system",
   "selfLink": "/api/v1/namespaces/openpitrix-system/services/hyperpitrix",
   "label": {
    "app": "openpitrix",
    "component": "openpitrix-hyperpitrix"
   },
   "startedAt": "2021-04-12T17:47:28+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "kube-dns",
   "namespace": "kube-system",
   "selfLink": "/api/v1/namespaces/kube-system/services/kube-dns",
   "label": {
    "k8s-app": "kube-dns"
   },
   "startedAt": "2021-04-12T13:53:32+08:00",
   "totalBackends": 2,
   "healthyBackends": 2
  },
  {
   "name": "kube-scheduler-svc",
   "namespace": "kube-system",
   "selfLink": "/api/v1/namespaces/kube-system/services/kube-scheduler-svc",
   "label": {
    "component": "kube-scheduler"
   },
   "startedAt": "2021-04-12T17:48:10+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  },
  {
   "name": "kube-controller-manager-svc",
   "namespace": "kube-system",
   "selfLink": "/api/v1/namespaces/kube-system/services/kube-controller-manager-svc",
   "label": {
    "component": "kube-controller-manager"
   },
   "startedAt": "2021-04-12T17:48:10+08:00",
   "totalBackends": 1,
   "healthyBackends": 1
  }
 ],
 "nodeStatus": {
  "totalNodes": 3,
  "healthyNodes": 3
 }
}
```



## 75、物理资源监控接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         集群状态监控                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/monitoring.kubesphere.io/v1alpha3/cluster       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/cluster?start=1624346737&end=1624352737&step=120s&times=50&metrics_filter=cluster_cpu_utilisation%7Ccluster_memory_utilisation%7Ccluster_load1%7Ccluster_load5%7Ccluster_load15%7Ccluster_disk_size_usage%7Ccluster_disk_inode_utilisation%7Ccluster_disk_inode_usage%7Ccluster_disk_inode_total%7Ccluster_disk_read_iops%7Ccluster_disk_write_iops%7Ccluster_disk_read_throughput%7Ccluster_disk_write_throughput%7Ccluster_net_bytes_transmitted%7Ccluster_net_bytes_received%7Ccluster_pod_running_count%7Ccluster_pod_abnormal_count%7Ccluster_pod_succeeded_count%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      | 开始 |  是  | String |
|      end       | 结束 |  是  | String |
|      step      | 步骤 |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

请求示例：

```
{
    "start": 1624347483
    "end": 1624353483
    "step": 120s
    "times": 50
    "metrics_filter":cluster_cpu_utilisation%7Ccluster_memory_utilisation%7Ccluster_load1%7Ccluster_load5%7Ccluster_load15%7Ccluster_disk_size_usage%7Ccluster_disk_inode_utilisation%7Ccluster_disk_inode_usage%7Ccluster_disk_inode_total%7Ccluster_disk_read_iops%7Ccluster_disk_write_iops%7Ccluster_disk_read_throughput%7Ccluster_disk_write_throughput%7Ccluster_net_bytes_transmitted%7Ccluster_net_bytes_received%7Ccluster_pod_running_count%7Ccluster_pod_abnormal_count%7Ccluster_pod_succeeded_count%24
}
```

响应示例：

​       成功情况

```
{results: [{metric_name: "cluster_cpu_utilisation",…},…]}
   results: [{metric_name: "cluster_cpu_utilisation",…},…]
```





## 76、APIServer监控接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                        APIServer监控                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/monitoring.kubesphere.io/v1alpha3/components/apiserver |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/components/apiserver?start=1624347013&end=1624353013&step=120s&times=50&metrics_filter=apiserver_request_latencies%7Capiserver_request_by_verb_latencies%7Capiserver_request_rate%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      | 开始 |  是  | String |
|      end       | 结束 |  是  | String |
|      step      | 步骤 |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

请求示例：

```
{    "start": 1624347394    "end": 1624353394    "step": 120s    "times": 50    "metrics_filter":apiserver_request_latencies%7Capiserver_request_by_verb_latencies%7Capiserver_request_rate%24}
```

响应示例：

​       成功情况

```
{   "results": [{metric_name: "apiserver_request_rate",…}, {metric_name: "apiserver_request_latencies",…},…]}
```



## 77、调度器监控接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                          调度器监控                          |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/monitoring.kubesphere.io/v1alpha3/components/apiserver |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/components/apiserver?start=1624347013&end=1624353013&step=120s&times=50&metrics_filter=apiserver_request_latencies%7Capiserver_request_by_verb_latencies%7Capiserver_request_rate%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      | 开始 |  是  | String |
|      end       | 结束 |  是  | String |
|      step      |      |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

请求示例：

```
{    "start": 1624347161    "end": 1624353161    "step": 120s    "times": 50    "metrics_filter":scheduler_schedule_attempts%7Cscheduler_schedule_attempt_rate%7Cscheduler_e2e_scheduling_latency%7Cscheduler_e2e_scheduling_latency_quantile%24}
```

响应示例：

​       成功情况

```
{   "results": [{metric_name: "scheduler_schedule_attempts", data: {resultType: "matrix"}},…]}
```



## 78、节点用量排行接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                         节点用量排行                         |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |        /kapis/monitoring.kubesphere.io/v1alpha3/nodes        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/nodes?type=rank&metrics_filter=node_cpu_utilisation%7Cnode_cpu_usage%7Cnode_cpu_total%7Cnode_memory_utilisation%7Cnode_memory_usage_wo_cache%7Cnode_memory_total%7Cnode_disk_size_utilisation%7Cnode_disk_size_usage%7Cnode_disk_size_capacity%7Cnode_pod_utilisation%7Cnode_pod_running_count%7Cnode_pod_quota%7Cnode_disk_inode_utilisation%7Cnode_disk_inode_total%7Cnode_disk_inode_usage%7Cnode_load1%24&page=1&limit=10&sort_type=desc&sort_metric=node_cpu_utilisation |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|      type      | 类型 |  是  | String |
| metrics_filter |      |  是  | String |
|      page      |      |  是  | String |
|     limit      |      |  是  | String |
|   sort_type    |      |  是  | String |
|  sort_metric   |      |  是  | String |

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

请求示例：

```
{  "type": rank  "metrics_filter": node_cpu_utilisation%7Cnode_cpu_usage%7Cnode_cpu_total%7Cnode_memory_utilisation%7Cnode_memory_usage_wo_cache%7Cnode_memory_total%7Cnode_disk_size_utilisation%7Cnode_disk_size_usage%7Cnode_disk_size_capacity%7Cnode_pod_utilisation%7Cnode_pod_running_count%7Cnode_pod_quota%7Cnode_disk_inode_utilisation%7Cnode_disk_inode_total%7Cnode_disk_inode_usage%7Cnode_load1%24  "page": 1  "limit": 10  "sort_type": desc  "sort_metric": node_cpu_utilisation}
```

响应示例：

​       成功情况

```
{results: [{metric_name: "node_cpu_utilisation", data: {resultType: "vector", result: [,…]}},…],…}page: 1results: [{metric_name: "node_cpu_utilisation", data: {resultType: "vector", result: [,…]}},…]total_item: 3total_page: 1{   "results": [{metric_name: "scheduler_schedule_attempts", data: {resultType: "matrix"}},…]}
```

​    

## 79、服务组件ks-apiserver接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha2/components/ks-apiserver |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/ks-apiserver |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
| healthyBackends |          |
|  totalBackends  |          |

响应示例：

​       成功情况

```
{
 "name": "ks-apiserver",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-apiserver",
 "label": {
  "app": "ks-apiserver",
  "tier": "backend",
  "version": "v3.0.0"
 },
 "startedAt": "2021-04-12T17:47:06+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```

  

## 80、服务组件redis接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha2/components/redis   |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/redis |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|      tier       |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{ "name": "redis", "namespace": "kubesphere-system", "selfLink": "/api/v1/namespaces/kubesphere-system/services/redis", "label": {  "app": "redis",  "tier": "database" }, "startedAt": "2021-04-12T17:46:16+08:00", "totalBackends": 1, "healthyBackends": 1}
```

​    

## 81、服务组件minio接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha2/components/minio   |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/minio |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|     release     |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{
 "name": "minio",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/minio",
 "label": {
  "app": "minio",
  "release": "ks-minio"
 },
 "startedAt": "2021-04-12T17:46:29+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```



## 82、服务组件etcd接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha2/components/etcd    |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/etcd |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|      tier       |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{ "name": "etcd", "namespace": "kubesphere-system", "selfLink": "/api/v1/namespaces/kubesphere-system/services/etcd", "label": {  "app": "kubesphere",  "tier": "etcd" }, "startedAt": "2021-04-12T17:46:42+08:00", "totalBackends": 1, "healthyBackends": 1}
```

​    

## 83、服务组件etcd接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha2/components/etcd    |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/etcd |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|      tier       |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{
 "name": "etcd",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/etcd",
 "label": {
  "app": "kubesphere",
  "tier": "etcd"
 },
 "startedAt": "2021-04-12T17:46:42+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```



## 84、服务组件mysql接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |   /kapis/resources.kubesphere.io/v1alpha2/components/mysql   |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/mysql |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|      tier       |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{
 "name": "mysql",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/mysql",
 "label": {
  "app": "kubesphere",
  "tier": "db"
 },
 "startedAt": "2021-04-12T17:46:42+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```

​    

## 85、服务组件ks-controller-manger接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha2/components/ks-controller-manager |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/ks-controller-manager |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|      tier       |          |
|     version     |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{
 "name": "ks-controller-manager",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-controller-manager",
 "label": {
  "app": "ks-controller-manager",
  "tier": "backend",
  "version": "v3.0.0"
 },
 "startedAt": "2021-04-12T17:47:06+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```

   

## 86、服务组件openldap接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha2/components/openldap  |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/openldap |

响应参数 

|            参数            |   描述   |
| :------------------------: | :------: |
|            name            |   名称   |
|         namespace          | 命名空间 |
|          selfLink          |          |
|           label            |          |
| app.kubernetes.io/instance |          |
|   app.kubernetes.io/name   |          |
|         startedAt          |          |
|       totalBackends        |          |

响应示例：

​       成功情况

```
{
 "name": "openldap",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/openldap",
 "label": {
  "app.kubernetes.io/instance": "ks-openldap",
  "app.kubernetes.io/name": "openldap-ha"
 },
 "startedAt": "2021-04-12T17:46:21+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```

​      

## 87、服务组件ks-console接口

### 应用场景

​      场景：

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                                                              |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/resources.kubesphere.io/v1alpha2/components/ks-console |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components/ks-console |

响应参数 

|      参数       |   描述   |
| :-------------: | :------: |
|      name       |   名称   |
|    namespace    | 命名空间 |
|    selfLink     |          |
|      label      |          |
|       app       |          |
|      tier       |          |
|     version     |          |
|    startedAt    |          |
|  totalBackends  |          |
| healthyBackends |          |

响应示例：

​       成功情况

```
{
 "name": "ks-console",
 "namespace": "kubesphere-system",
 "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-console",
 "label": {
  "app": "ks-console",
  "tier": "frontend",
  "version": "v3.0.0"
 },
 "startedAt": "2021-04-12T17:47:08+08:00",
 "totalBackends": 1,
 "healthyBackends": 1
}
```

​      

## 88、应用资源用量接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                     监控集群应用资源用量                     |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |     /kapis/monitoring.kubesphere.io/v1alpha3/namespaces      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces?type=rank&metrics_filter=namespace_memory_usage_wo_cache%7Cnamespace_memory_limit_hard%7Cnamespace_cpu_usage%7Cnamespace_cpu_limit_hard%7Cnamespace_pod_count%7Cnamespace_pod_count_hard%7Cnamespace_net_bytes_received%7Cnamespace_net_bytes_transmitted&page=1&limit=10&sort_type=desc&sort_metric=namespace_cpu_usage |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|      type      |      |  是  | String |
| metrics_filter |      |  是  | String |
|      page      |      |  是  | String |
|     limit      |      |  是  | String |
|   sort_type    |      |  是  | String |
|  sort_metric   |      |  是  | String |

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

请求示例：

```
{
   "type": rank
   "metrics_filter": namespace_memory_usage_wo_cache%7Cnamespace_memory_limit_hard%7Cnamespace_cpu_usage%7Cnamespace_cpu_limit_hard%7Cnamespace_pod_count%7Cnamespace_pod_count_hard%7Cnamespace_net_bytes_received%7Cnamespace_net_bytes_transmitted
   "page": 1
   "limit": 10
   "sort_type": desc
   "sort_metric": namespace_cpu_usage
}
```

 响应示例：

​       成功情况

```
{
    results: [{metric_name: "namespace_pod_count_hard",…},…], page: 1, total_page: 2, total_item: 11
}
```

​    

## 89、应用资源监控接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                  监控集群应用资源的运行状态                  |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/monitoring.kubesphere.io/v1alpha3/cluster       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/cluster?start=1624260416&end=1624346816&step=3600s&times=24&metrics_filter=cluster_cpu_usage%7Ccluster_memory_usage_wo_cache%7Ccluster_disk_size_usage%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      | 开始 |  是  | String |
|      end       | 结束 |  是  | String |
|      step      | 步骤 |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

请求示例：

```
{
    "start": 1624260416
    "end": 1624346816
    "step": 3600s
    "times": 24
    "metrics_filter": cluster_cpu_usage%7Ccluster_memory_usage_wo_cache%7Ccluster_disk_size_usage%24
}
```

 响应示例：

​       成功情况

```
{
    results: [{metric_name: "cluster_cpu_usage", data: {resultType: "matrix", result: [{,…}]}},…]
}
```

​      



## 90、集群节点接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |   集群节点提供了当前集群下的运行状态，以及可以编辑删除节点   |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |        /kapis/resources.kubesphere.io/v1alpha3/nodes         |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/nodes?sortBy=createTime&limit=10 |

请求参数

|  参数  | 描述 | 必填 |  类型  |
| :----: | :--: | :--: | :----: |
| sortBy |      |  是  | String |
| limit  |      |  是  | String |

请求示例：

```
{
   "sortBy": createTime
   "limit": 10
}
```

响应参数 

| 参数  | 描述 |
| :---: | :--: |
| items | 条目 |

响应示例：

​       成功情况

```
{items: [{,…}, {,…}, {,…}], totalItems: 3}
items: [{,…}, {,…}, {,…}]
totalItems: 3
```

​    

## 91、集群节点事件接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |   集群节点提供了当前集群下的运行状态，以及可以编辑删除节点   |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |                        /api/v1/events                        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/events?fieldSelector=involvedObject.name%3Dkubex-0003%2CinvolvedObject.kind%3DNode |

请求参数

|     参数      | 描述 | 必填 |  类型  |
| :-----------: | :--: | :--: | :----: |
| fieldSelector |      |  是  | String |

请求示例：

```
{
     "fieldSelector": involvedObject.name%3Dkubex-0003%2CinvolvedObject.kind%3DNode
}
```

响应参数 

|    参数    | 描述 |
| :--------: | :--: |
|    kind    |      |
| apiVersion |      |

响应示例：

​       成功情况

```
{kind: "EventList", apiVersion: "v1",…}
  apiVersion: "v1"
  items: []
  kind: "EventList"
metadata: {selfLink: "/api/v1/events", resourceVersion: "22556166"}
```

​       

## 92、集群节点监控接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |   集群节点提供了当前集群下的运行状态，以及可以编辑删除节点   |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |        /kapis/monitoring.kubesphere.io/v1alpha3/nodes        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/nodes?start=1624383731&end=1624413731&step=600s&times=50&resources_filter=kubex-0003%24&metrics_filter=node_cpu_utilisation%7Cnode_load1%7Cnode_load5%7Cnode_load15%7Cnode_memory_utilisation%7Cnode_disk_size_utilisation%7Cnode_disk_inode_utilisation%7Cnode_disk_inode_usage%7Cnode_disk_inode_total%7Cnode_disk_read_iops%7Cnode_disk_write_iops%7Cnode_disk_read_throughput%7Cnode_disk_write_throughput%7Cnode_net_bytes_transmitted%7Cnode_net_bytes_received%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      |      |  是  | String |
|      end       |      |  是  | String |
|      step      |      |  是  | String |
|     times      |      |  是  | String |
|   resources    |      |  是  | String |
| metrics_filter |      |  是  | String |

请求示例：

```
{     "start": 1624383731     "end": 1624413731     "step": 600s     "times": 50     "resources_filter": kubex-0003%24     "metrics_filter": node_cpu_utilisation%7Cnode_load1%7Cnode_load5%7Cnode_load15%7Cnode_memory_utilisation%7Cnode_disk_size_utilisation%7Cnode_disk_inode_utilisation%7Cnode_disk_inode_usage%7Cnode_disk_inode_total%7Cnode_disk_read_iops%7Cnode_disk_write_iops%7Cnode_disk_read_throughput%7Cnode_disk_write_throughput%7Cnode_net_bytes_transmitted%7Cnode_net_bytes_received%24}
```

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

响应示例：

​       成功情况

```
{results: [{metric_name: "node_cpu_utilisation", data: {resultType: "matrix", result: [,…]}},…]}results: [{metric_name: "node_cpu_utilisation", data: {resultType: "matrix", result: [,…]}},…]
```

  

## 93、容器组接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                             监控                             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |         /kapis/resources.kubesphere.io/v1alpha3/pods         |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha3/pods?limit=6&nodeName=kubex-0003&sortBy=startTime |

请求参数

|   参数   | 描述 | 必填 |  类型  |
| :------: | :--: | :--: | :----: |
|  limit   |      |  是  | String |
| nodeName |      |  是  | String |
|  sortBy  |      |  是  | String |

请求示例：

```
{
   "limit": 6
   "nodeName": kubex-0003
   "sortBy": startTime
}
```

响应参数 

| 参数  | 描述 |
| :---: | :--: |
| items | 条目 |

响应示例：

​       成功情况

```
{items: [{metadata: {name: "mysql-dvgv23-test", namespace: "dev",…},…},…], totalItems: 20}
items: [{metadata: {name: "mysql-dvgv23-test", namespace: "dev",…},…},…]
totalItems: 20
```

​     

## 93、节点运行状态接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                             监控                             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |        /kapis/monitoring.kubesphere.io/v1alpha3/nodes        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/nodes?start=1624410663&end=1624414263&step=180s&times=20&resources_filter=kubex-0003%24&metrics_filter=node_cpu_utilisation%7Cnode_memory_utilisation%7Cnode_disk_size_utilisation%7Cnode_pod_utilisation%24 |

请求参数

|       参数       | 描述 | 必填 |  类型  |
| :--------------: | :--: | :--: | :----: |
|      start       |      |  是  | String |
|       end        |      |  是  | String |
|       step       |      |  是  | String |
|      times       |      |  是  | String |
| resources_filter |      |  是  | String |
|  metrics_filter  |      |  是  | String |

请求示例：

```
{    "start": 1624410663    "end": 1624414263    "step": 180s    "times": 20    "resources_filter": kubex-0003%24    "metrics_filter": node_cpu_utilisation%7Cnode_memory_utilisation%7Cnode_disk_size_utilisation%7Cnode_pod_utilisation%24}
```

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

响应示例：

​       成功情况

```
{results: [{metric_name: "node_pod_utilisation", data: {resultType: "matrix",…}},…]}
results: [{metric_name: "node_pod_utilisation", data: {resultType: "matrix",…}},…]
```

​       

## 94、容器日志接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                             监控                             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /api/v1/namespaces/kubesphere-monitoring-system/pods/prometheus-k8s-1/log |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/kubesphere-monitoring-system/pods/prometheus-k8s-1/log?container=prometheus&tailLines=1000&timestamps=true&follow=false |

请求参数

|    参数    | 描述 | 必填 |  类型  |
| :--------: | :--: | :--: | :----: |
| container  |      |  是  | String |
| tailLines  |      |  是  | String |
| timestamps |      |  是  | String |
|   follow   |      |      | String |

请求示例：

```
{
   "container": prometheus
   "tailLines": 1000
   "timestamps": true
   "follow": false
}
```

响应示例：

​       成功情况

```
2021-06-03T07:02:31.803831316Z level=info ts=2021-06-03T07:02:31.803Z caller=main.go:343...
```

​    

## 95、监控接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                             监控                             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         | /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-monitoring-system/pods/prometheus-k8s-1/containers/prometheus |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/kubesphere-monitoring-system/pods/prometheus-k8s-1/containers/prometheus?cluster=default&namespace=kubesphere-monitoring-system&container=prometheus&podName=prometheus-k8s-1&start=1624387818&end=1624417818&step=600s&times=50&metrics_filter=container_cpu_usage%7Ccontainer_memory_usage_wo_cache%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|    cluster     |      |  是  | String |
|   namespace    |      |  是  | String |
|   container    |      |  是  | String |
|    podName     |      |  是  | String |
|     start      |      |  是  | String |
|      end       |      |  是  | String |
|      step      |      |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

请求示例：

```
{
   "cluster": default
   "namespace": kubesphere-monitoring-system
   "container": prometheus
   "podName": prometheus-k8s-1
   "start": 1624387818
   "end": 1624417818
   "step": 600s
   "times": 50
   "metrics_filter": container_cpu_usage%7Ccontainer_memory_usage_wo_cache%24
}
```

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

响应示例：

​       成功情况

```
{
   results: [{metric_name: "container_memory_usage_wo_cache", data: {resultType: "matrix", result: [{,…}]}},…]
}
```

​     

## 96、资源状态接口

### 应用场景

​      场景：监控告警

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                             监控                             |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |        /kapis/monitoring.kubesphere.io/v1alpha3/nodes        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/nodes?start=1624410663&end=1624414263&step=180s&times=20&resources_filter=kubex-0003%24&metrics_filter=node_cpu_utilisation%7Cnode_memory_utilisation%7Cnode_disk_size_utilisation%7Cnode_pod_utilisation%24 |

请求参数

|       参数       | 描述 | 必填 |  类型  |
| :--------------: | :--: | :--: | :----: |
|      start       |      |  是  | String |
|       end        |      |  是  | String |
|       step       |      |  是  | String |
|      times       |      |  是  | String |
| resources_filter |      |  是  | String |
|  metrics_filter  |      |  是  | String |

请求示例：

```
{
    "start": 1624410663
    "end": 1624414263
    "step": 180s
    "times": 20
    "resources_filter": kubex-0003%24
    "metrics_filter": node_cpu_utilisation%7Cnode_memory_utilisation%7Cnode_disk_size_utilisation%7Cnode_pod_utilisation%24
}
```

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

响应示例：

​       成功情况

```
{results: [{metric_name: "node_pod_utilisation", data: {resultType: "matrix",…}},…]}
results: [{metric_name: "node_pod_utilisation", data: {resultType: "matrix",…}},…]
```

​       

## 97、基本信息接口

### 应用场景

​      场景：集群设置

###  接口描述

|          描述          |                   内容                   |
| :--------------------: | :--------------------------------------: |
|        接口功能        |           当前集群基础信息总览           |
|        请求协议        |                   HTTP                   |
|        请求方法        |                   GET                    |
|        请求格式        |                   JSON                   |
|        请求url         |              /kapis/version              |
| 请求头(和请求格式对应) |      Content-Type:application/json       |
|        请求示例        | http://119.3.157.144:30880/kapis/version |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|     start      | 开始 |  是  | String |
|      end       | 结束 |  是  | String |
|      step      | 步骤 |  是  | String |
|     times      |      |  是  | String |
| metrics_filter |      |  是  | String |

响应参数 

|    参数    | 描述 |
| :--------: | :--: |
| getVersion |      |

响应示例：

​       成功情况

```
{
 "gitVersion": "v0.0.0",
 "gitMajor": "",
 "gitMinor": "",
 "gitCommit": "d567f438ff62e946d8955eb6d49968971d42097f",
 "gitTreeState": "dirty",
 "buildDate": "2020-08-28T10:15:00Z",
 "goVersion": "go1.13.15",
 "compiler": "gc",
 "platform": "linux/amd64",
 "kubernetes": {
  "major": "1",
  "minor": "18",
  "gitVersion": "v1.18.17",
  "gitCommit": "68b4e26caf6ede7af577db4af62fb405b4dd47e6",
  "gitTreeState": "clean",
  "buildDate": "2021-03-18T00:54:02Z",
  "goVersion": "go1.13.15",
  "compiler": "gc",
  "platform": "linux/amd64"
 }
}
```



##  98、集群信息接口

### 应用场景

​      场景：集群设置

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                     当前集群基础信息总览                     |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |       /kapis/monitoring.kubesphere.io/v1alpha3/cluster       |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/cluster?cluster=default&metrics_filter=cluster_cpu_total%7Ccluster_memory_total%7Ccluster_disk_size_capacity%7Ccluster_node_total%24 |

请求参数

|      参数      | 描述 | 必填 |  类型  |
| :------------: | :--: | :--: | :----: |
|    cluster     |      |  是  | String |
| metrics_filter |      |  是  | String |

请求示例：

```
{
    "cluster": default
    "metrics_filter": cluster_cpu_total%7Ccluster_memory_total%7Ccluster_disk_size_capacity%7Ccluster_node_total%24
}
```

响应参数 

|  参数   | 描述 |
| :-----: | :--: |
| results | 结果 |

响应示例：

​       成功情况

```
{
 "results": [
  {
   "metric_name": "cluster_memory_total",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624418627.757,
       "12380930048"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_node_total",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624418627.757,
       "3"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_cpu_total",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624418627.757,
       "6"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "cluster_disk_size_capacity",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624418627.757,
       "126015406080"
      ]
     }
    ]
   }
  }
 ]
}
```
## 99、个人账号相关数量查询接口

### 应用场景

​      场景：用账号，密码登录kubesphere平台，查询个人账号下，相关集群数量，企业空间、账号、应用模板数量信息。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                用账号，密码登录kubesphere平台                |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |     /kapis/monitoring.kubesphere.io/v1alpha3/kubesphere      |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/kubesphere?start=1624324570&end=1624336570&step=600s&times=20&metrics_filter=kubesphere_cluser_count%7Ckubesphere_workspace_count%7Ckubesphere_user_count%7Ckubesphere_app_template_count%24 |

请求参数

|        参数        |    描述     | 必填 |  类型  |
| :----------------: | :---------: | :--: | :----: |
|     **start**      |  开始时间   |  是  | String |
|      **end**       |  结束时间   |  是  | String |
|      **step**      |             |      | String |
|     **times**      |    次数     |  是  | String |
| **metrics_filter** | metrics指标 |  是  | String |

响应参数 

|             参数              |   描述   |
| :---------------------------: | :------: |
|    kubesphere_cluser_count    | 集群数量 |
|  kubesphere_workspace_count   | 企业空间 |
|     kubesphere_user_count     |   账号   |
| kubesphere_app_template_count | 应用模板 |

请求示例

```
{
    start:1624324570
    end:1624336570
    step:600s
    times:20
    metrics_filter:kubesphere_cluser_count|kubesphere_workspace_count|kubesphere_user_count|kubesphere_app_template_count$
}
```

响应示例：

​       成功情况

```
{
    {
 "results": [
  {
   "metric_name": "kubesphere_cluser_count",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624336570,
       "1"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "kubesphere_workspace_count",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624336570,
       "2"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "kubesphere_user_count",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624336570,
       "10"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "kubesphere_app_template_count",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "value": [
       1624336570,
       "15"
      ]
     }
    ]
   }
  }
 ]
}
}
```

​       失败情况

```
{
}
```
###  

## 100、查询个人基本信息接口

### 应用场景

​      场景：编辑修改个人信息。

###  接口描述

|          描述          |                             内容                             |
| :--------------------: | :----------------------------------------------------------: |
|        接口功能        |                       编辑修改个人信息                       |
|        请求协议        |                             HTTP                             |
|        请求方法        |                             GET                              |
|        请求格式        |                             JSON                             |
|        请求url         |        /kapis/iam.kubesphere.io/v1alpha2/users/{user}        |
| 请求头(和请求格式对应) |                Content-Type:application/json                 |
|        请求示例        | http://119.3.157.144:30880/kapis/iam.kubesphere.io/v1alpha2/users/wangqi |

请求参数

| 参数 |  描述  | 必填 |  类型  |
| :--: | :----: | :--: | :----: |
| user | 用户名 |  是  | String |

响应参数 

| 参数  |  描述  |
| :---: | :----: |
| name  | 用户名 |
| email |  邮箱  |

请求示例

```
无
```

响应示例：

​       成功情况

```
{
 "metadata": {
  "name": "wangqi",
  "selfLink": "/apis/iam.kubesphere.io/v1alpha2/users/wangqi",
  "uid": "a6d50c8b-d41e-4ac1-ba08-d76f710251be",
  "resourceVersion": "22322757",
  "generation": 51,
  "creationTimestamp": "2021-05-19T07:40:55Z",
 },
 "spec": {
  "email": "wang@qq.com"
 },
 "status": {
  "state": "Active",
  "lastTransitionTime": "2021-05-19T07:40:56Z",
  "lastLoginTime": "2021-06-22T09:07:24Z"
 }
}
```

## 101、系统组件查询接口

### 应用场景 

> 列出系统组件



### 接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 列举出系统中的部件                                           |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | kapis/resources.kubesphere.io/v1alpha2/components            |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/resources.kubesphere.io/v1alpha2/components |

响应参数:

| 参数                                    | 描述       |
| --------------------------------------- | :--------- |
| ks-controller-manager                   | 管理控制器 |
| ks-apiserver                            |            |
| minio                                   |            |
| etcd                                    |            |
| redis                                   |            |
| mysql                                   |            |
| ks_console                              |            |
| kube-state-metrics                      |            |
| alertmanager-operated                   |            |
| prometheus-k8s                          |            |
| prometheus-operated                     |            |
| notification-manager-controller-metrics |            |
| notification-manager-svc                |            |
| alertmanager-main                       |            |
| node-exporter                           |            |
| prometheus-operator                     |            |
| hyperpitrix                             |            |
| kube-scheduler-svc                      |            |
| kube-controller-manager-svc             |            |
| kube-dns                                |            |

响应示例：

> 成功情况：

```
[
 {
  "name": "ks-controller-manager",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-controller-manager",
  "label": {
   "app": "ks-controller-manager",
   "tier": "backend",
   "version": "v3.0.0"
  },
  "startedAt": "2021-04-12T17:47:06+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "openldap",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/openldap",
  "label": {
   "app.kubernetes.io/instance": "ks-openldap",
   "app.kubernetes.io/name": "openldap-ha"
  },
  "startedAt": "2021-04-12T17:46:21+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "ks-apiserver",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-apiserver",
  "label": {
   "app": "ks-apiserver",
   "tier": "backend",
   "version": "v3.0.0"
  },
  "startedAt": "2021-04-12T17:47:06+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "minio",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/minio",
  "label": {
   "app": "minio",
   "release": "ks-minio"
  },
  "startedAt": "2021-04-12T17:46:29+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "etcd",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/etcd",
  "label": {
   "app": "kubesphere",
   "tier": "etcd"
  },
  "startedAt": "2021-04-12T17:46:42+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "redis",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/redis",
  "label": {
   "app": "redis",
   "tier": "database"
  },
  "startedAt": "2021-04-12T17:46:16+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "mysql",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/mysql",
  "label": {
   "app": "kubesphere",
   "tier": "db"
  },
  "startedAt": "2021-04-12T17:46:42+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "ks-console",
  "namespace": "kubesphere-system",
  "selfLink": "/api/v1/namespaces/kubesphere-system/services/ks-console",
  "label": {
   "app": "ks-console",
   "tier": "frontend",
   "version": "v3.0.0"
  },
  "startedAt": "2021-04-12T17:47:08+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "kube-state-metrics",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/kube-state-metrics",
  "label": {
   "app.kubernetes.io/name": "kube-state-metrics"
  },
  "startedAt": "2021-04-12T17:48:05+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "alertmanager-operated",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/alertmanager-operated",
  "label": {
   "app": "alertmanager"
  },
  "startedAt": "2021-04-12T17:48:12+08:00",
  "totalBackends": 3,
  "healthyBackends": 3
 },
 {
  "name": "prometheus-k8s",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/prometheus-k8s",
  "label": {
   "app": "prometheus",
   "prometheus": "k8s"
  },
  "startedAt": "2021-04-12T17:48:10+08:00",
  "totalBackends": 2,
  "healthyBackends": 2
 },
 {
  "name": "prometheus-operated",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/prometheus-operated",
  "label": {
   "app": "prometheus"
  },
  "startedAt": "2021-04-12T17:48:21+08:00",
  "totalBackends": 2,
  "healthyBackends": 2
 },
 {
  "name": "notification-manager-controller-metrics",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/notification-manager-controller-metrics",
  "label": {
   "control-plane": "controller-manager"
  },
  "startedAt": "2021-04-12T17:48:26+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "notification-manager-svc",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/notification-manager-svc",
  "label": {
   "app": "notification-manager",
   "notification-manager": "notification-manager"
  },
  "startedAt": "2021-04-12T17:48:30+08:00",
  "totalBackends": 2,
  "healthyBackends": 2
 },
 {
  "name": "alertmanager-main",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/alertmanager-main",
  "label": {
   "alertmanager": "main",
   "app": "alertmanager"
  },
  "startedAt": "2021-04-12T17:48:12+08:00",
  "totalBackends": 3,
  "healthyBackends": 3
 },
 {
  "name": "node-exporter",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/node-exporter",
  "label": {
   "app.kubernetes.io/name": "node-exporter"
  },
  "startedAt": "2021-04-12T17:48:04+08:00",
  "totalBackends": 3,
  "healthyBackends": 3
 },
 {
  "name": "prometheus-operator",
  "namespace": "kubesphere-monitoring-system",
  "selfLink": "/api/v1/namespaces/kubesphere-monitoring-system/services/prometheus-operator",
  "label": {
   "app.kubernetes.io/component": "controller",
   "app.kubernetes.io/name": "prometheus-operator"
  },
  "startedAt": "2021-04-12T17:48:03+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "hyperpitrix",
  "namespace": "openpitrix-system",
  "selfLink": "/api/v1/namespaces/openpitrix-system/services/hyperpitrix",
  "label": {
   "app": "openpitrix",
   "component": "openpitrix-hyperpitrix"
  },
  "startedAt": "2021-04-12T17:47:28+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "kube-scheduler-svc",
  "namespace": "kube-system",
  "selfLink": "/api/v1/namespaces/kube-system/services/kube-scheduler-svc",
  "label": {
   "component": "kube-scheduler"
  },
  "startedAt": "2021-04-12T17:48:10+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "kube-controller-manager-svc",
  "namespace": "kube-system",
  "selfLink": "/api/v1/namespaces/kube-system/services/kube-controller-manager-svc",
  "label": {
   "component": "kube-controller-manager"
  },
  "startedAt": "2021-04-12T17:48:10+08:00",
  "totalBackends": 1,
  "healthyBackends": 1
 },
 {
  "name": "kube-dns",
  "namespace": "kube-system",
  "selfLink": "/api/v1/namespaces/kube-system/services/kube-dns",
  "label": {
   "k8s-app": "kube-dns"
  },
  "startedAt": "2021-04-12T13:53:32+08:00",
  "totalBackends": 2,
  "healthyBackends": 2
 }
]
```

 

## 102、查询组件级别的指标数据接口

### 应用场景 

> 场景1：xxxxxxxx

### 接口描述 

| 描述                   | 内容                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 接口功能               | 获取特定系统组件的组件级指标数据                             |
| 请求协议               | HTTP                                                         |
| 请求方法               | GET                                                          |
| 请求格式               | json                                                         |
| 请求url                | /kapis/monitoring.kubesphere.io/v1alpha3/components/{component} |
| 请求头(和请求格式对应) | Content-Type: application/json                               |
| 请求示例               | http://119.3.157.144:30880/kapis/monitoring.kubesphere.io/v1alpha3/components/apiserver?metrics_filter=apiserver_request_latencies%7Capiserver_request_rate%24 |

请求参数

|      参数      |                             描述                             | 必填 |  类型  |
| :------------: | :----------------------------------------------------------: | :--: | :----: |
|   component    |    要监控的系统组件， etcd、apiserver、scheduler、之一。     |  是  | String |
| metrics_filter | 指标名称过滤器由一个正则表达式模式组成。 它指定要返回的指标数据。 例如，以下过滤器匹配 etcd 服务器列表和底层数据库的总大小：`etcd_server_list|etcd_mvcc_db_size`。 |  否  | String |
|     start      | 查询的开始时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如。 1559347200。 |  否  | String |
|      end       | 查询结束时间。 使用 **start** 和 **end** 检索一段时间内的指标数据。 它是一个 Unix 时间格式的字符串，例如。 1561939200。 |  否  | string |
|      step      | Unix 时间格式的时间戳。 在单个时间点检索指标数据。 默认为现在。 时间和开始、结束、步骤的组合是互斥的。 |  否  | string |

响应参数 

|            参数             | 描述 |
| :-------------------------: | :--: |
| apiserver_request_latencies |      |
|   apiserver_request_rate    |      |

请求示例

```
{
    "metrics_filter": "apiserver_request_latencies%7Capiserver_request_rate%24"
}
```

响应示例：

​       成功情况

```
{
 "results": [
  {
   "metric_name": "apiserver_request_latencies",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "metric": {
       "__name__": "apiserver:apiserver_request_duration:avg"
      },
      "value": [
       1624327121.493,
       "0.0012848180054974363"
      ]
     }
    ]
   }
  },
  {
   "metric_name": "apiserver_request_rate",
   "data": {
    "resultType": "vector",
    "result": [
     {
      "metric": {
       "__name__": "apiserver:apiserver_request_total:sum_irate"
      },
      "value": [
       1624327121.493,
       "10.150000000000002"
      ]
     }
    ]
   }
  }
 ]
}
```

​      失败情况

```
{
"code":401,"kind":"Status","apiVersion":"v1","metadata":{},"status":"Failure","message":"Unauthorized: token is expired by 45m55s","reason":"Unauthorized","statusText":"Unauthorized"
}
```

##  103、创建密钥接口

### 应用场景

### 接口描述

|          描述          |                            内容                             |
| :--------------------: | :---------------------------------------------------------: |
|        接口功能        |                          创建密钥                           |
|        请求协议        |                            HTTP                             |
|        请求方法        |                            POST                             |
|        请求格式        |                            JSON                             |
|        请求url         |              api/v1/namespaces/gaofei/secrets               |
| 请求头(和请求格式对应) |                Content-Type:application/json                |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/gaofei/secrets |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{    "apiVersion": "v1",    "kind": "Secret",    "metadata": {        "namespace": "gaofei",        "labels": {},        "name": "asdfg",        "annotations": {            "kubesphere.io/alias-name": "asdfg",            "kubesphere.io/description": "asdfg",            "kubesphere.io/creator": "wangqi"        }    },    "type": "Opaque",    "spec": {        "template": {            "metadata": {                "labels": {}            }        }    },    "data": {        "asd": "YXNk"    }}
```

响应示例：

​       成功情况

```json
{    "kind": "Secret",    "apiVersion": "v1",    "metadata": {        "name": "asdfg",        "namespace": "gaofei",        "selfLink": "/api/v1/namespaces/gaofei/secrets/asdfg",        "uid": "b3cd50f9-f514-4503-9149-6cf6d96be6cd",        "resourceVersion": "22362340",        "creationTimestamp": "2021-06-22T11:58:54Z",        "annotations": {            "kubesphere.io/alias-name": "asdfg",            "kubesphere.io/creator": "wangqi",            "kubesphere.io/description": "asdfg"        },        "managedFields": [            {                "manager": "Mozilla",                "operation": "Update",                "apiVersion": "v1",                "time": "2021-06-22T11:58:54Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:data": {                        ".": {},                        "f:asd": {}                    },                    "f:metadata": {                        "f:annotations": {                            ".": {},                            "f:kubesphere.io/alias-name": {},                            "f:kubesphere.io/creator": {},                            "f:kubesphere.io/description": {}                        }                    },                    "f:type": {}                }            }        ]    },    "data": {        "asd": "YXNk"    },    "type": "Opaque"}
```

### 

##  104、大屏整体总负载接口

### 应用场景

### 接口描述

|          描述          |                            内容                             |
| :--------------------: | :---------------------------------------------------------: |
|        接口功能        |                          创建密钥                           |
|        请求协议        |                            HTTP                             |
|        请求方法        |                            POST                             |
|        请求格式        |                            JSON                             |
|        请求url         |              api/v1/namespaces/gaofei/secrets               |
| 请求头(和请求格式对应) |                Content-Type:application/json                |
|        请求示例        | http://119.3.157.144:30880/api/v1/namespaces/gaofei/secrets |

请求参数

| 参数 | 描述 | 必填 | 类型 |
| :--: | :--: | :--: | :--: |
|      |      |      |      |
|      |      |      |      |

响应参数 

|    参数    |  描述   |
| :--------: | :-----: |
| apiVersion | api版本 |
|    kind    |  类型   |
|  metadata  | 元数据  |
|    spec    |  描述   |
|   status   |  状态   |

请求示例

```json
{    "apiVersion": "v1",    "kind": "Secret",    "metadata": {        "namespace": "gaofei",        "labels": {},        "name": "asdfg",        "annotations": {            "kubesphere.io/alias-name": "asdfg",            "kubesphere.io/description": "asdfg",            "kubesphere.io/creator": "wangqi"        }    },    "type": "Opaque",    "spec": {        "template": {            "metadata": {                "labels": {}            }        }    },    "data": {        "asd": "YXNk"    }}
```

响应示例：

​       成功情况

```json
{    "kind": "Secret",    "apiVersion": "v1",    "metadata": {        "name": "asdfg",        "namespace": "gaofei",        "selfLink": "/api/v1/namespaces/gaofei/secrets/asdfg",        "uid": "b3cd50f9-f514-4503-9149-6cf6d96be6cd",        "resourceVersion": "22362340",        "creationTimestamp": "2021-06-22T11:58:54Z",        "annotations": {            "kubesphere.io/alias-name": "asdfg",            "kubesphere.io/creator": "wangqi",            "kubesphere.io/description": "asdfg"        },        "managedFields": [            {                "manager": "Mozilla",                "operation": "Update",                "apiVersion": "v1",                "time": "2021-06-22T11:58:54Z",                "fieldsType": "FieldsV1",                "fieldsV1": {                    "f:data": {                        ".": {},                        "f:asd": {}                    },                    "f:metadata": {                        "f:annotations": {                            ".": {},                            "f:kubesphere.io/alias-name": {},                            "f:kubesphere.io/creator": {},                            "f:kubesphere.io/description": {}                        }                    },                    "f:type": {}                }            }        ]    },    "data": {        "asd": "YXNk"    },    "type": "Opaque"}
```



















































































