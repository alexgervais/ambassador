# Ambassador as an Ingress Controller

Starting with version 0.80.0, Ambassador can act as a Kubernetes
[ingress controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/),
reading configuration data from Kubernetes 
[`Ingress`](https://kubernetes.io/docs/concepts/services-networking/ingress/) resources.
This makes it easier to work with other `Ingress`-oriented tools within the Kubernetes
ecosystem, and it makes it easier for users migrating from other
ingress controllers to try Ambassador.

## When and How to Use the `Ingress` Resource

In order to use the `Ingress` resource effectively, it's important to understand not
just how to use it, but also when to use it, and how it interacts with CRDs and
annotations.

### What is required to use the `Ingress` resource?

- You will need RBAC permissions to create `Ingress` resources.

- Ambassador will need RBAC permissions to get, list, and watch `Ingress` resources. 

  You can see this in the `https://getambassador.io/yaml/ambassador-rbac.yaml`
  file, but this is the critical rule to add to Ambassador's `Role` or `ClusterRole`:

      - apiGroups: [ "extensions" ]
        resources: [ "ingresses" ]
        verbs: ["get", "list", "watch"]

- You must create your `Ingress` resource with the correct `ingress.class`.

  Ambassador will automatically read Ingress resources with the annotation
  `kubernetes.io/ingress.class: ambassador`.

- You may need to set your `Ingress` resources' `ambassador-id`.

  If you're not using the `default` ID, you'll need to add the `getambassador.io/ambassador-id`
  annotation to your `Ingress`. See the examples below.

### When should I use an `Ingress` instead of annotations or CRDs?

As of 0.80.0, Datawire recommends that Ambassador be configured with CRDs. The `Ingress`
resource is available to users who need it for integration with other ecosystem tools, or
who feel that it more closely matches their workflows -- however, it is important to 
recognize that the `Ingress` resource is rather more limited than the Ambassador `Mapping`
is (for example, the `Ingress` spec has no support for rewriting or for TLS origination).
**When in doubt, use CRDs.**

### Can 0.80.0 support using an `Ingress` and CRDs in concert?

Yes. All Ambassador configuration mechanisms are the same under the hood: from
Ambassador's point of view, it doesn't matter if you use CRDs, annotations, or
`Ingress` resources. Even in 0.80.0, it is definitely supported to use an `Ingress`
to define the edge of your cluster, and CRDs for more advanced functionality.

## `Ingress` Support in 0.80.0

Ambassador 0.80.0 supports basic core functionality of the  `Ingress` resource, as
defined by the [`Ingress`](https://kubernetes.io/docs/concepts/services-networking/ingress/)
resource itself:

- Basic routing, including the `route` specification and the default backend
  functionality, is supported.
   - it's particularly easy to use a minimal `Ingress` to the Ambassador diagnostic UI
- TLS termination is supported.
   - you can use multiple `Ingress` resources for SNI
- Using the `Ingress` resource in concert with Ambassador CRDs or annotations is supported.
   - this includes Ambassador annotations on the `Ingress` resource itself

Ambassador 0.80.0 does **not** extend the basic `Ingress` specification except as follows:

- the `getambassador.io/ambassador-id` annotation allows you to set an Ambassador ID for
  the `Ingress` itself; and

- the `getambassador.io/config` annotation can be provided on the `Ingress` resource, just
  as on a `Service`.
     - note that if you need to set `getambassador.io/ambassador-id` on the `Ingress`, you
       will also need to set `ambassador-id` on resources within the annotation!

### `Ingress` routes and `Mapping`s

Ambassador actually creates `Mapping` objects from the `Ingress` route rules. These `Mapping`
objects interact with `Mapping`s defined in CRDs **exactly** as they would if the `Ingress`
route rules had been specified with CRDs originally.

For example, this `Ingress` resource

```yaml
---
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
  name: test-ingress
spec:
  rules:
  - http:
      paths:
      - path: /foo/
        backend:
          serviceName: service1
          servicePort: 80
```

is **exactly equivalent** to a `Mapping` CRD of 

```yaml
---
apiVersion: getambassador.io/v1
kind: Mapping
metadata:
  name: test-ingress-0-0
spec:
  prefix: /foo/
  service: service1:80
```

This means that the following YAML:

```yaml
---
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
  name: test-ingress
spec:
  rules:
  - http:
      paths:
      - path: /foo/
        backend:
          serviceName: service1
          servicePort: 80
---
apiVersion: getambassador.io/v1
kind: Mapping
metadata:
  name: my-mapping
spec:
  prefix: /foo/
  service: service2
```

will set up Ambassador to do canary routing where 50% of the traffic will go to `service1`
and 50% will go to `service2`.

### The Minimal `Ingress`

An `Ingress` resource must provide at least of some routes or a
[default backend](https://kubernetes.io/docs/concepts/services-networking/ingress/#default-backend).
The default backend provides for a simple way to direct all traffic to some upstream
service:

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
  name: test-ingress
spec:
  backend:
    serviceName: exampleservice
    servicePort: 8080
```

This is equivalent to

```yaml
---
apiVersion: getambassador.io/v1
kind: Mapping
metadata:
  name: test-ingress
spec:
  prefix: /
  service: exampleservice:8080
```

### [Name based virtual hosting](https://kubernetes.io/docs/concepts/services-networking/ingress/#name-based-virtual-hosting) with an Ambassador ID

```yaml
---
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
    getambassador.io/ambassador-id: externalid
  name: name-virtual-host-ingress
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - backend:
          serviceName: service1
          servicePort: 80
   - host: bar.foo.com
     http:
       paths:
       - backend:
           serviceName: service2
           servicePort: 80
```

This is equivalent to

```yaml
---
apiVersion: getambassador.io/v1
kind: Mapping
metadata:
  name: host-foo-mapping
spec:
  ambassador_id: externalid
  prefix: /
  host: foo.bar.com
  service: service1
---
apiVersion: getambassador.io/v1
kind: Mapping
metadata:
  name: host-bar-mapping
spec:
  ambassador_id: externalid
  prefix: /
  host: bar.foo.com
  service: service2
```

and will result in all requests to `foo.bar.com` going to `service1`, and requests to `bar.foo.com` going to `service2`.

### [TLS termination](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls)

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
  name: tls-example-ingress
spec:
  tls:
  - hosts:
    - sslexample.foo.com
      secretName: testsecret-tls
  rules:
    - host: sslexample.foo.com
      http:
        paths:
        - path: /
          backend:
            serviceName: service1
            servicePort: 80
```

This is equivalent to

```yaml
---
apiVersion: getambassador.io/v1
kind: TLSContext
metadata:
  name: sslexample-termination-context
spec:
  hosts:
  - sslexample.foo.com
  secret: testsecret-tls
---
apiVersion: getambassador.io/v1
kind: Mapping
metadata:
  name: sslexample-mapping
spec:
  host: sslexample.foo.com
  prefix: /
  service: service1
```

Note that this shows TLS termination, not origination: the `Ingress` spec does not
support origination.
