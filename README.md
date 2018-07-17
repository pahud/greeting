# greeting - simple greeting service	



### Usage

```bash
$ docker run -d -p 8080:8080 --name greeting pahud/greeting:latest

$ curl localhost:8080/?name=pahud
Hello pahud
```



## Helm Install

```
$ helm install ./greeting-helm/             
NAME:   vested-coral
LAST DEPLOYED: Tue Jul 17 14:25:44 2018
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ServiceAccount
NAME     SECRETS  AGE
ingress  1        1s

==> v1beta1/ClusterRoleBinding
NAME     AGE
ingress  1s

==> v1/Service
NAME                    TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)       AGE
caddy-service           ClusterIP  10.100.166.104  <none>       80/TCP        1s
nginx-service           ClusterIP  10.100.67.78    <none>       80/TCP        1s
greeting-service        ClusterIP  10.100.213.121  <none>       80/TCP        1s
traefik-ingress-lb-svc  NodePort   10.100.44.209   <none>       80:31742/TCP  1s

==> v1beta1/DaemonSet
NAME                DESIRED  CURRENT  READY  UP-TO-DATE  AVAILABLE  NODE SELECTOR  AGE
traefik-ingress-lb  5        5        0      5           0          <none>         0s

==> v1beta1/Deployment
NAME      DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
greeting  2        2        2           0          0s
caddy     2        2        2           0          0s
nginx     2        2        2           0          0s

==> v1beta1/Ingress
NAME             HOSTS  ADDRESS  PORTS  AGE
traefik-ingress  *      80       0s

==> v1/Pod(related)
NAME                      READY  STATUS             RESTARTS  AGE
traefik-ingress-lb-64dn4  0/1    ContainerCreating  0         0s
traefik-ingress-lb-67pc7  0/1    ContainerCreating  0         0s
traefik-ingress-lb-btr7j  0/1    ContainerCreating  0         0s
traefik-ingress-lb-pr2f9  0/1    ContainerCreating  0         0s
traefik-ingress-lb-qpvm7  0/1    ContainerCreating  0         0s
greeting-b4df9d899-nrf8w  0/1    ContainerCreating  0         0s
greeting-b4df9d899-s957j  0/1    ContainerCreating  0         0s
caddy-76c4887dd7-c6hfn    0/1    ContainerCreating  0         0s
caddy-76c4887dd7-l48td    0/1    ContainerCreating  0         0s
nginx-7d86684c7c-5cwbj    0/1    ContainerCreating  0         0s
nginx-7d86684c7c-t8gks    0/1    ContainerCreating  0         0s
```

