# Quản lý và tính toán tài nguyên sử dụng cho Pod

## Requesting resources
+ thêm thuộc tính( chỉ định thuôc tính cho từng container trong pod)
+ resource request-resource limits += những container của nó

## Tạo pod với resource requests
+ tạo file tên requests-pod.yaml
+ tạo Pod và kiểm tra thử cpu của nó : ```kubectl apply -f requests-pod.yaml```
```kubectl exec -it requests-pod top```

## Resource requests có ý nghĩa:
+ ví dụ ta có resource.requests: cpu: 200m, memory:10Mi
+ giá trị mà ta cài cho resource.requests sẽ đc sử dụng trong lúc 1 Pod schedule tới worker node.
+ Scheduler → tìm worker node có CPU unallocated >200m và memory unallocated >10Mi → schedule Pod tới worker node đó. (Scheduler không xem xét gtri CPU và memory chưa unallocated)
=> Nếu ta run Pod cần 200m CPU, nhưng k chỉ định rõ → worker node còn có 100m CPU, Pod đc schedule tới worker node.
+ gtri CPU và memory được Scheduler tính = lấy (resources  của node - tổng resources.requests của Pod) chứ k phải gtri CPU, memory còn free
+ nếu gtri free >= đáp ứng đc Pod request, nma gtri unallocated không đáp ứng đc => không schedule tới.

## Cách Scheduler sử dụng resource requests để chọn node tốt nhất cho Pod
+ lọc ra node tốt nhất = thuật toán rồi sắp xếp
+ resource requests là 1 tiêu chí trong đó
+ 2 phương thức là LeastRequestedPriority, MostRequestedPriority.
+ LeastRequestedPriorit: sẽ chọn node có tổng resource request là ít nhất, nghĩa là resource unallocated là nhiều nhất.
+ MostRequestedPriority: sẽ chọn node có tổng resource request là cao nhất, nghĩa là resource unallocated là ít nhất.

## Scheduler in action
+ xem capacity của node: ```kubectl describe node```
+ chú ý: Capacity: tổng tài nguyên của worker node, Allocatable: tài nguyên mà Pod có thể dùng.
+ tạo 1 pod mà request 3000m xem cpu allocated có tăng không? ```kubectl run requests-pod-2 --image=busybox --restart Never --requests='cpu=3000m,memory=20Mi' -- dd if=/dev/zero of=/dev/null```
+(do em dùng version > 2.1 nên e sẽ tạo file config là requests-pod_test.yaml)
→ tạo và xem thay đổi

## Limiting resources
+ dùng resources limit  để tránh Pod chiếm tài nguyên Pod khác
+ tạo config ```limited-pod.yaml```
+ ta dùng resources.limits, nếu k dùng request → mặc định dùng với limits
+ tạo file và ktra ```kubectl apply -f limited-pod.yaml```
```kubectl exec -it limits-pod top```

## Khi một container vượt quá limit
+ TH1: với CPU → không thể dùng quá số CPU ta chỉ định
+ TH2: với memory → container sẽ bị kill, nếu chỉ định restartPolicylaf Always hoặc OnFailure → container sẽ tự động restart

## QoS classes
+ tổng memory limit khi config có thể > 100% memory của worker node
+ đọc ví dụ trên viblo
+ đánh giá độ ưu tiên của Pod: BestEffort (độ ưu tiên thấp nhất), Burstable, Guaranteed (độ ưu tiên cao nhất)

## Xác định QoS của container
+ BestEffort gán cho container mà không có chỉ định của resource requests. nếu memory của nó > 100%, Pod sẽ bị kill đầu tiên
+ Burstable:  gán cho container mà chỉ định mỗi resource requests( hoặc cả limits nhưng gtri khoong  bằng nhau) → bị kill khi không còn Pod có class BestEffort tồn tại.
+ Guaranteed → bị kill cuối cùng, đc gán cho Pod mà: Chỉ định cả hai thuộc resources requests và limits trong container. Hoặc chỉ định mỗi resources limits, resources requests sẽ mặc định lấy giá trị của resources limits. Hai thuộc tính này giá trị phải bằng nhau.

## Xác định QoS của Pod
+  xem tren viblo

## Cách Pod cùng QoS sẽ bị kill
+ Mỗi process của chúng ta sẽ có một OutOfMemory (OOM) score
→ Pod và QoS cái nào có process và OOM cao hơn sẽ bị kill trước.
+ OOM score = số memory request + memory đang sử dụng.

### LimitRange resource
+ Tạo một resource tên là LimitRange, để chỉ định requests và limits mặc định trong một namespace thay vì cấu hình cho từng container riêng.
LimitRange: chỉ định min, max của requests và limits và gtri mặc định khi không chỉ định thuộc tính trong container.
+ tạo limitRange sẽ config trong Admission control plugin

## Tạo LimitRange
+ tạo LimitRange resource, tạo file limits.yaml
+ chỉ định limit cho 3 thành phần là Pod, container, PersistentVolumeClaim
+ min CPU của Pod = 50m, memory = 5Mi; max là 1cpu và memory = 1Gi
+ tạo và ktra LimitRange: ```kubectl apply -f limits.yaml```
```kubectl run requests-pod-big-cpu --image=busybox –requests='cpu=3'```
+ kiểm tra gtri chỉ định requests và limits
```kubectl run pod-no-setting-resoucre --image=busybox --restart Never -- dd if=/dev/zero of=/dev/null``` , 
```kubectl describe pod pod-no-setting-resoucre```

## Giới hạn tổng số lượng tài nguyên của một namespace
### ResourceQuota
→ để giới hạn tài nguyên của 1 ns
+ tạo 1 ResourceQuota resource → config ở Admission Control plugin
+ khi 1 Pod đc tạo → qua plugin check giới hạn tài nguyên( nếu lớn hơn→ API server trả về lỗi)
### ResourceQuota giới hạn cpu và memory
+ tạo file quota-cpu-memory.yaml
+ Xem resource đc sử dụng bên trong quota ```kubectl apply -f quota-cpu-memory.yaml``` , ```kubectl describe quota cpu-and-mem```
+ tạo 1 pod mới xem có tăng thông số không ```kubectl run quota-pod --image=busybox --restart Never --limits='cpu=300m,memory=200Mi' -- dd if=/dev/zero of=/dev/null```
+ tạo 1 pod request 400m cpu ```kubectl run quota-pod-1 --image=busybox --restart Never --limits='cpu=400m,memory=200Mi' -- dd if=/dev/zero of=/dev/null``` → lỗi
### ResourceQuota giới hạn persistent storage
### ResourceQuota giới hạn số lượng resource có thể tạo
+ Những resource mà ResourceQuota có thể chỉ định là:
Pods
ReplicationControllers
Secrets
ConfigMaps
PersistentVolumeClaims
Services: có thể chỉ định rõ số lượng LoadBalancer Services và NodePort Services
### Quota scope cho Pod
+ Đối với Pod, ta có thể chỉ định Quota có được áp dụng với tới nó hay không dựa vào 4 thuộc tính sau đây:
+ BestEffort: chỉ ảnh hưởng tới Pod với Qos class là BestEffort
+ NotBestEffort: chỉ ảnh hưởng tới Pod với Qos class là Burstable và Guaranteed
+ Terminating: chỉ ảnh hưởng tới Pod có thuộc tính activeDeadlineSeconds
+ NotTerminating: chỉ ảnh hưởng tới Pod không có thuộc tính activeDeadlineSeconds