# go-pki

Intent to implement different models of PKI based on:

- https://pki-tutorial.readthedocs.io/en/latest/simple/index.html
- https://pki-tutorial.readthedocs.io/en/latest/advanced/index.html
- https://pki-tutorial.readthedocs.io/en/latest/expert/index.html

The intention is create a modular PKI that could fit the 3 above scenarios or
even more complex scenarios, but letting the user to choose the architecture
of PKI that fit their organization. A new way that could scale and work in
cloud workloads as a microservice.

## Stage 1 - Simple

Provide value by letting to create the simple scenario. This scenario will
implement all the functionality in only one CA object.

## Stage 2 - Add ACME

Add ACME protocol for the simple scenario.

## Stage 3 - Add OCSP

Add OCSP protocol for the simple scenario.

## Stage 4 - Advanced

Refactor code to fit properly the Advanced model of PKI.

## Stage 4 - Expert

Refactor code to fit properly the Expert model of PKI. This is the most flexible
model.

## Stage 5 - Cloud workloads

Even if cloud approach is considered from the early stage, face the cloud
scenario can evoke additional refactors and validation of scenarios.

- The workload can be upgraded with no issues.
- HA is provided with minimum of 3 replicas.
- CQRS architecture pattern is used to let scale read and writes independently.

## Stage 6 - Integration with cert manager in kubernetes

- Refactor to fit cert manager integration in Kubernetes.

## Stage 7 - Hardening

Even if the security is considered from the early stage, I think is important
to focus specifically and dedicate one stage only for this.

- Encrypted communications.
- Rate limits applied.
- Ban weird sources (if no nonce is provided, if no URLs mapped, if no signed
  payloads, ...), add them to a temporary ban list, and reject future requests
  for a period of time. The list should persist so if the service is reloaded
  or upgraded, the list is not lost. Let the user to configure parameters as:
  - threshold of bad requests.
  - rate limit.
  - ban time.
- Apply host block list.
- Reject homographic URLs.
- Reject TLD.
