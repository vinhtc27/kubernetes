# ConfigMap and Secret

## Chỉ định biến môi trường cho container
+ Tạo file pod-hello-env.yaml (đã có)
+ Tạo Pod: ```kubectl apply -f pod-hello-env.yaml```
```kubectl logs hello-env```
+ kiểm tra kĩ hơn bằng cách truy cập vào container và in ra env:
```kubectl exec -it hello-env -- sh```
```/app # env | grep PORT```
==> list env này sẽ không thể update bên trong container khi container đó đã chạy, muốn update thì ta phải xóa Pod đó để nó chạy lại.

## ConfigMap
+ Tạo một ConfigMap. Tạo một file cm-db.yaml (đã có)
```kubectl apply -f cm-db.yaml```
+ Truyền ConfigMap vào bên trong container, tạo file pod-with-cm.yaml (đã có) sử dụng image 080196/hello-cm
+ tạo và kiểm tra thử:
```kubectl apply -f pod-with-cm.yaml -l app=db```
```kubectl get pod```
```kubectl apply -f pod-with-cm.yaml -l app=application```
```kubectl logs hello-cm```

## Dùng ConfigMap để truyền cấu hình dạng file vào trong container thông qua volume config
+ tạo file nginx-config-cm.yaml(da co)
+ Tạo file pod-nginx.yaml(da co)
+ tạo và kiểm tra thử:
```kubectl apply -f nginx-config-cm.yaml```
```kubectl apply -f pod-nginx.yaml```
```kubectl port-forward nginx 8080:80```
+ Mở một terminal khác:
```curl -H "Accept-Encoding: gzip" -I localhost:8080```
+ kiểm tra kỹ hơn bằng cách truy cập vào trong Pod:
```kubectl exec -it nginx -- sh```
```/ # cd /etc/nginx/conf.d/```
```/etc/nginx/conf.d # ls```
```/etc/nginx/conf.d # cat my-nginx-config.conf ```
+ Dổi tên của file: lam theo huong dan tren series

## Secret
### Tạo một Secret
+  nên dùng CLI hơn là tạo file config