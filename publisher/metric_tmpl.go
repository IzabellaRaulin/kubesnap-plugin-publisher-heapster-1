/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package publisher

const builtinMetricTemplate = `{
"id": "!!",
"name": "!!",
"aliases": [
],
"namespace": "docker",
"labels": {
},
  "subcontainers": [],
"spec": {
 "creation_time": "/creation_time",
 "labels": {
 },
 "has_cpu": true,
 "cpu": {
 },
 "has_memory": true,
 "memory": {
 },
 "has_network": true,
 "has_filesystem": false,
 "has_diskio": true,
 "has_custom_metrics": true,
 "custom_metrics":[],
 "image": "/image_name"
},
"stats": [
 {
  "timestamp": "!!",
  "custom_metrics": {
   "SNAP": [
   ]
  },
  "cpu": {
   "usage": {
    "total": "/cgroups/cpu_stats/cpu_usage/total_usage"
   }
  },
  "memory": {
   "usage": "/cgroups/memory_stats/usage/usage",
   "container_data": {
    "pgfault": "/cgroups/memory_stats/stats/pgfault",
    "pgmajfault": "/cgroups/memory_stats/stats/pgmajfault"
   }
  },
  "network": {
   "rx_bytes": "0",
   "rx_errors": "0",
   "tx_bytes": "0",
   "tx_errors": "0",
   "interfaces":[
   ]
  }
 }
]
}
`
