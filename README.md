# docker plugin secrets

docker secret plugin.

<https://docs.docker.com/engine/extend/#debugging-plugins>

## how to build docker plugin

```bash
make clean
make create
make push
```

## how to install plugin
 
install and setup enabled by default:

```bash
docker plugin install --grant-all-permissions docker-plugin-secrets:0.0.1
```

setup token if renew:

```bash
docker plugin disable docker-plugin-secrets:1.0.0
docker plugin set docker-plugin-secrets:1.0.0
docker plugin enable docker-plugin-secrets:1.0.0
```

## how to debug plugin

docker以debug模式启动

```json
{
  "debug": true
}
```
    
    
查看log

```bash
journalctl -f -u docker.service

cd /run/docker/plugins/$your_plugin_id
cat < init-stdout
cat < init-stderr
```
    
## how to use 

use it in compose file

```yaml
secrets:
  haproxy:
    driver: docker-plugin-secrets:0.0.1
    labels:
      docker.plugin.secretprovider.vault.path: canux/data/pki
      docker.plugin.secretprovider.vault.field: "*.canuxcheng.com"
```

