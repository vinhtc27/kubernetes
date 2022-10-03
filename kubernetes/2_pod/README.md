# Kubernetes Pod

## Tạo file docker và chạy container bằng Docker
+ Tạo file Dockerfile cho file main.go (đã có)
+ Run câu lệnh build image: ```docker build . -t golang-docker-example```
+ Test thử container có chạy đúng hay không, chạy container bằng câu lệnh: ```docker run -d --name hello-kube -p 3000:3000 golang-docker-example```
+ Xóa resources: ```docker rm -f hello-kube```

## Chạy container bằng Pod
+ Tạo file tên là hello-kube.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-kube.yaml```
+ Kiểm tra pod đã chạy hay chưa: ```kubectl get pod```
+ Chạy câu lệnh sau để expose port của pod: ```kubectl port-forward pod/hello-kube 3000:3000```
+ Test reques gửi tới pod: ```curl localhost:3000```
+ Sau khi chạy xong để clear resource thì chúng ta xóa pod bằng câu lệnh: ```kubectl delete pod hello-kube```

## Tổ chức pod bằng cách sử dụng labels 
+ Tạo một file tên là hello-kube-label.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-kube-label.yaml```
+ Ta có thể list pod với labels như sau: ```kubectl get pod --show-labels```
+ Ta có thể chọn chính xác cột label hiển thị với -L options: ```kubectl get pod -L enviroment```
+ Và ta có thể lọc pod theo label với -l options: ```kubectl get pod -l enviroment=production```
+ Để clear resource thì chúng ta xóa pod: ```kubectl delete -f hello-kube-label.yaml```

## Phân chia tài nguyên của kubernetes cluster bằng cách sử dụng namespace
+ Đầu tiên chúng ta list ra toàn bộ namespace: ```kubectl get ns```
+ Ta có thể chỉ định resource của namespace chúng ta muốn bằng cách thêm option --namespace: ```kubectl get pod --namespace kube-system```
+ Cách tổ chức namespace tốt là tạo theo: ```<project_name>:<enviroment>```
+ Ở đây làm nhanh thì mình sẽ không đặt namespace theo cách trên: ```kubectl create ns testing```
+ Tạo file tên hello-kube-ns.yaml (đã có)
+ Tạo pod: ```kubectl apply -f hello-kube-ns.yaml```
+ Để list pod, ta phải chỉ định thêm namespace chúng ta muốn lấy: ```kubectl get pod -n testing```
+ Khi xóa thì ta cũng cần chỉ định namespace chứa resource: ```kubectl delete pod hello-kube-testing -n testing```
+ Có thể xóa namespace bằng cách dùng câu lệnh delete, chú ý là khi xóa namespace thì toàn bộ resource trong đó cũng sẽ bị xóa theo: ```kubectl delete ns testing```