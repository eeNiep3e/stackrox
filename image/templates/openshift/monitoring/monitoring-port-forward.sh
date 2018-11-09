#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"

if [[ -z "$1" ]]; then
	echo "usage: bash monitoring-port-forward.sh 3000"
	echo "The above would forward localhost:3000 to monitoring:3000."
	exit 1
fi

until [ "$(oc get pod -n {{.K8sConfig.Namespace}} --selector 'app=monitoring' | grep Running | wc -l)" -eq 1 ]; do
    echo -n .
    sleep 1
done

export MONITORING_POD="$(oc get pod -n {{.K8sConfig.Namespace}} --selector 'app=monitoring' --output=jsonpath='{.items..metadata.name} {.items..status.phase}' | grep Running | cut -f 1 -d ' ')"
oc port-forward -n "{{.K8sConfig.Namespace}}" "${MONITORING_POD}" "$1:3000" > /dev/null &
echo "Access monitoring on localhost:$1"
