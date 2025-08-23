# Securing in-cluster Virtual Machines and Pods Communication with mTLS

## Summary

Within Harvester, current in-cluster communication between virtual machines (VMs) and pods is unencrypted, exposing user workloads to potential spoofing and man-in-the-middle attacks. This enhancement proposal introduces the use of a service mesh solution (such as Linkerd or Istio) to provide automatic mutual TLS (mTLS) for all internal communications, significantly improving in-cluster security and moving towards a zero-trust network model. The solution aims to ensure encrypted traffic and authenticated service-to-service communication with manageable disruption to existing workloads.

### Related Issues

- [Harvester GitHub Issues: Networking and Security](https://github.com/harvester/harvester/issues?q=is%3Aissue+label%3Anetworking+label%3Asecurity)

## Motivation

### Goals

- Encrypt all pod and VM communication in the cluster using mTLS.
- Ensure service-to-service authentication and encryption.
- Minimize operational disruption, enable opt-in/opt-out per-namespace or workload.
- Provide certificate management with automatic rotation.

### Non-goals [optional]

- Encryption for communication external to the cluster (north-south traffic) is out of scope.
- Migration of legacy workloads incompatible with sidecar injection is not addressed.

## Proposal

### User Stories

#### Story 1
_As a Harvester user, I want communication between in-cluster VMs and pods to be encrypted automatically so that I do not have to manually provision TLS certificates or handle key management._

#### Story 2
_As a cluster administrator, I want to be able to monitor and audit service-to-service communications and assert compliance with zero-trust networking practices._

### User Experience In Detail

- By enabling the service mesh, administrators can enforce mTLS for namespaces or the entire cluster by setting a policy.
- No application code changes are required; sidecar proxies handle encryption/decryption.
- Certificate issuance and rotation are automatic and transparent.
- Example: Applying a `PeerAuthentication` policy in Istio YAML:

```yaml
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: default
spec:
  mtls:
    mode: STRICT
```

- After deployment, all intra-cluster communication will be transparently encrypted.

### API changes

No new Harvester API objects; integration is achieved via standard Kubernetes Custom Resource Definitions (CRDs) supplied by the service mesh (e.g., Istio, Linkerd).

## Design

### Implementation Overview

- Evaluate Istio and Linkerd service meshes for ease of integration and operational overhead.
- Deploy the chosen mesh in Harvester clusters, integrating with KubeVirt (for VMs) and standard pods.
- Enable automatic mTLS and flexible policy management per namespace/workload.
- Use built-in certificate authorities for issuing x509 certs to sidecar proxies.

**Diagram: Proposed mTLS-enforced Data Plane**

![mTLS Service Mesh Architecture](./mtls-service-mesh-architecture.png)

### Test plan

- Provision Harvester with and without service mesh in test clusters.
- Use network tracing tools to verify encrypted traffic.
- Test with workloads consisting of VMs, pods, and mixed scenarios.
- Validate certificate issuance/rotation and connectivity with/without mTLS enabled.

### Upgrade strategy

- The feature should be opt-in, with a migration guide for workloads requiring the mesh.
- Provide documentation for progressive rollout (per-namespace, per-deployment).
- No breaking changes to existing unencrypted workloads; gradual adoption is supported.

## Note [optional]

- Potential future work: integrate with external certificate authorities or hardware security modules (HSMs).
- Consider performance implications for high-throughput environments.
