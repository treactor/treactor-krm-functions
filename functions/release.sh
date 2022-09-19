docker tag gcr.io/treactor/kpt-fn/create-atoms:unstable gcr.io/treactor/kpt-fn/create-atoms:$1
docker tag gcr.io/treactor/kpt-fn/ensure-sa:unstable gcr.io/treactor/kpt-fn/ensure-sa:$1
docker tag gcr.io/treactor/kpt-fn/ensure-telemetry:unstable gcr.io/treactor/kpt-fn/ensure-telemetry:$1
docker tag gcr.io/treactor/kpt-fn/gateway-api:unstable gcr.io/treactor/kpt-fn/gateway-api:$1

docker push gcr.io/treactor/kpt-fn/create-atoms:$1
docker push gcr.io/treactor/kpt-fn/ensure-sa:$1
docker push gcr.io/treactor/kpt-fn/ensure-telemetry:$1
docker push gcr.io/treactor/kpt-fn/gateway-api:$1
