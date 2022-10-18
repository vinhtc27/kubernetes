# PersistentVolumeClaims, PersistentVolumes
- 2 role:
    + kubernetes administrator
    kubernetes administrator -> setup storage -> tạo PersistentVolumes.

    + kubernetes developer

## Tạo PersistentVolumes
+ Tạo một file tên pv-gcepd.yaml (đã có)
+ Tạo PV , list ra:
```kubectl apply -f pv-gcepd.yaml```
```kubectl get pv```

## Tạo PersistentVolumeClaim tiêu thụ PersistentVolumes
+ ta là developer, cần deploy Pod, xài volume để lưu trữ persistent data
+ Tạo file mongodb-pvc.yaml (đã có)
+ Tạo PVCs và list ra xem:
```kubectl apply -f mongodb-pvc.yaml```
```kubectl get pvc```

## Tạo Pod sử dụng PersistentVolumeClaim
+ tạo Pod xài PVCs, tạo file mongodb-pod-pvc.yaml (đã có)
+ Tạo một gcePersistentDisk và insert dữ liệu vào trong đó:
```kubectl create -f mongodb-pod-pvc.yaml```
```kubectl exec -it mongodb mongo```

## Lợi ích của việc xài PersistentVolumeClaim

## Recycling PersistentVolumes
+ Xóa thử  PVCs: ```kubectl delete pod mongodb```
```kubectl delete pvc mongodb-pvc```
```kubectl get pv```

## Tự động cấp PersistentVolumes (Dynamic provisioning)

## Tạo StorageClass
+ tự động tạo PV, ta chỉ cần tạo StorageClass một lần
+ Tạo file storageclass-fast-gcepd.yaml (đã có)
+ Tạo file mongodb-pvc-dp.yaml (đã có)
```kubectl apply -f mongodb-pvc-dp.yaml```
```kubectl get pvc mongodb-pvc```
```kubectl get pv```

## Dynamic provisioning mà không cần chỉ định storage class
+ không chỉ định thuộc tính storageClassName trong PVCs -> dungf storage class mặc định
```kubectl get sc```
